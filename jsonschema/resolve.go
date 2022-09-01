package jsonschema

import (
	"context"

	"github.com/go-faster/errors"
	yaml "github.com/go-faster/yamlx"

	"github.com/ogen-go/ogen/internal/jsonpointer"
)

// ReferenceResolver resolves JSON schema references.
type ReferenceResolver interface {
	ResolveReference(ref string) (*RawSchema, error)
}

func (p *Parser) getResolver(ctx *jsonpointer.ResolveCtx) (ReferenceResolver, error) {
	loc := ctx.LastLoc()

	r, ok := p.schemas[loc]
	if ok {
		return r, nil
	}

	var node yaml.Node
	if err := func() error {
		raw, err := p.external.Get(context.TODO(), loc)
		if err != nil {
			return errors.Wrap(err, "get")
		}

		if err := yaml.Unmarshal(raw, &node); err != nil {
			return errors.Wrap(err, "unmarshal")
		}

		return nil
	}(); err != nil {
		return nil, errors.Wrapf(err, "external %q", loc)
	}

	r = NewRootResolver(&node)
	p.schemas[loc] = r

	return r, nil
}

func (p *Parser) resolve(ref string, ctx *jsonpointer.ResolveCtx) (*Schema, error) {
	if s, ok := p.refcache[ref]; ok {
		return s, nil
	}

	key, err := ctx.Add(ref)
	if err != nil {
		return nil, err
	}
	defer func() {
		// Drop the resolved ref to prevent false-positive infinite recursion detection.
		ctx.Delete(key)
	}()

	resolver, err := p.getResolver(ctx)
	if err != nil {
		return nil, err
	}

	raw, err := resolver.ResolveReference(key.Ref)
	if err != nil {
		return nil, errors.Wrap(err, "find schema")
	}

	return p.parse1(raw, ctx, func(s *Schema) *Schema {
		s.Ref = ref
		p.refcache[ref] = s
		return p.extendInfo(raw, s)
	})
}
