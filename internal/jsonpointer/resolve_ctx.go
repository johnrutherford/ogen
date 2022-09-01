package jsonpointer

import (
	"net/url"

	"github.com/go-faster/errors"
)

// RefKey is JSON reference key.
type RefKey struct {
	Loc string
	Ref string
}

// FromURL sets RefKey from URL.
func (r *RefKey) FromURL(u *url.URL) {
	{
		// Make copy.
		u2 := *u
		u2.Fragment = ""
		r.Loc = u2.String()
	}
	r.Ref = "#" + u.Fragment
}

// ResolveCtx is JSON pointer resolve context.
type ResolveCtx struct {
	// Location stack. Used for context-depending resolving.
	//
	// For resolve trace like
	//
	// 	"#/components/schemas/Schema" ->
	// 	"https://example.com/schema#Schema" ->
	//	"#/definitions/SchemaProperty"
	//
	// "#/definitions/SchemaProperty" should be resolved against "https://example.com/schema".
	locstack []string
	// Store references to detect infinite recursive references.
	refs       map[RefKey]struct{}
	depthLimit int
}

// DefaultDepthLimit is default depth limit for ResolveCtx.
const DefaultDepthLimit = 1000

// DefaultCtx creates new ResolveCtx with default depth limit.
func DefaultCtx() *ResolveCtx {
	return NewResolveCtx(DefaultDepthLimit)
}

// NewResolveCtx creates new ResolveCtx.
func NewResolveCtx(depthLimit int) *ResolveCtx {
	return &ResolveCtx{
		locstack:   nil,
		refs:       map[RefKey]struct{}{},
		depthLimit: depthLimit,
	}
}

// Add adds reference to context and returns key.
func (r *ResolveCtx) Add(ref string) (key RefKey, _ error) {
	parser := url.Parse
	if loc := r.LastLoc(); loc != "" {
		base, err := url.Parse(loc)
		if err != nil {
			return key, err
		}
		parser = base.Parse
	}

	u, err := parser(ref)
	if err != nil {
		return RefKey{}, err
	}
	key.FromURL(u)

	if r.depthLimit <= 0 {
		return RefKey{}, errors.New("depth limit exceeded")
	}
	if _, ok := r.refs[key]; ok {
		return key, errors.New("infinite recursion")
	}
	r.refs[key] = struct{}{}
	r.depthLimit--

	if key.Loc != "" {
		r.locstack = append(r.locstack, key.Loc)
	}
	return key, nil
}

// Delete removes reference from context.
func (r *ResolveCtx) Delete(key RefKey) {
	r.depthLimit++
	delete(r.refs, key)
	if key.Loc != "" && len(r.locstack) > 0 {
		r.locstack = r.locstack[:len(r.locstack)-1]
	}
}

// LastLoc returns last location from stack.
func (r *ResolveCtx) LastLoc() string {
	s := r.locstack
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}