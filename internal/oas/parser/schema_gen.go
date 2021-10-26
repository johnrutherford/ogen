package parser

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/xerrors"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/internal/oas"
)

// schemaGen is used to convert openapi schemas into oas representation.
type schemaGen struct {
	// Spec is used to lookup for schema components
	// which is not in the 'refs' cache.
	spec *ogen.Spec

	// Global schema reference cache from *Generator (readonly).
	globalRefs map[string]*oas.Schema

	// Local schema reference cache.
	localRefs map[string]*oas.Schema
}

// Generate converts ogen.Schema into *oas.Schema.
//
// If ogen.Schema contains references to schema components,
// these referenced schemas will be saved in g.localRefs.
func (g *schemaGen) Generate(schema ogen.Schema) (*oas.Schema, error) {
	s, err := g.generate(schema, "")
	if err != nil {
		return nil, xerrors.Errorf("gen: %w", err)
	}

	return s, nil
}

func (g *schemaGen) generate(schema ogen.Schema, ref string) (*oas.Schema, error) {
	if ref := schema.Ref; ref != "" {
		s, err := g.ref(ref)
		if err != nil {
			return nil, xerrors.Errorf("ref %q: %w", ref, err)
		}
		return s, nil
	}

	onret := func(s *oas.Schema) *oas.Schema {
		if ref != "" {
			g.localRefs[ref] = s
		}
		return s
	}

	switch {
	case len(schema.Enum) > 0:
		values, err := g.parseEnumValues(oas.SchemaType(schema.Type), schema.Enum)
		if err != nil {
			return nil, xerrors.Errorf("parse enum: %w", err)
		}

		return onret(&oas.Schema{
			Type:        oas.SchemaType(schema.Type),
			Format:      oas.Format(schema.Format),
			Description: schema.Description,
			Enum:        values,
		}), nil
	case len(schema.OneOf) > 0:
		var schemas []*oas.Schema
		for i, s := range schema.OneOf {
			schema, err := g.generate(s, "")
			if err != nil {
				return nil, xerrors.Errorf("oneOf: %d: %w", i, err)
			}

			schemas = append(schemas, schema)
		}

		return onret(&oas.Schema{
			OneOf:       schemas,
			Ref:         schema.Ref,
			Description: schema.Description,
		}), nil
	case len(schema.AnyOf) > 0:
		var schemas []*oas.Schema
		for i, s := range schema.AnyOf {
			schema, err := g.generate(s, "")
			if err != nil {
				return nil, xerrors.Errorf("anyOf: %d: %w", i, err)
			}

			schemas = append(schemas, schema)
		}

		return onret(&oas.Schema{
			AnyOf:       schemas,
			Ref:         schema.Ref,
			Description: schema.Description,
		}), nil
	case len(schema.AllOf) > 0:
		var schemas []*oas.Schema
		for i, s := range schema.AllOf {
			schema, err := g.generate(s, "")
			if err != nil {
				return nil, xerrors.Errorf("allOf: %d: %w", i, err)
			}

			schemas = append(schemas, schema)
		}

		return onret(&oas.Schema{
			AllOf:       schemas,
			Ref:         schema.Ref,
			Description: schema.Description,
		}), nil
	}

	if err := g.validateTypeFormat(schema.Type, schema.Format); err != nil {
		return nil, xerrors.Errorf("validate format: %w", err)
	}

	switch schema.Type {
	case "object":
		if schema.Items != nil {
			return nil, xerrors.New("object cannot contain 'items' field")
		}
		optional := func(name string) bool {
			for _, p := range schema.Required {
				if p == name {
					return false
				}
			}
			return true
		}
		s := &oas.Schema{
			Type:          oas.Object,
			Description:   schema.Description,
			Ref:           ref,
			MinProperties: schema.MinProperties,
			MaxProperties: schema.MaxProperties,
		}
		if ref != "" {
			g.localRefs[ref] = s
		}

		for propName, propSchema := range schema.Properties {
			prop, err := g.generate(propSchema, "")
			if err != nil {
				return nil, xerrors.Errorf("%s: %w", propName, err)
			}

			s.Properties = append(s.Properties, oas.Property{
				Name:     propName,
				Schema:   prop,
				Optional: optional(propName),
				Nullable: propSchema.Nullable,
			})
		}
		sort.SliceStable(s.Properties, func(i, j int) bool {
			return strings.Compare(s.Properties[i].Name, s.Properties[j].Name) < 0
		})
		return s, nil

	case "array":
		array := &oas.Schema{
			Type:        oas.Array,
			Description: schema.Description,
			Ref:         ref,
			MinItems:    schema.MinItems,
			MaxItems:    schema.MaxItems,
			UniqueItems: schema.UniqueItems,
		}
		if schema.Items == nil {
			// Fallback to string.
			array.Item = &oas.Schema{Type: oas.String}
			return array, nil
		}
		if len(schema.Properties) > 0 {
			return nil, xerrors.New("array cannot contain properties")
		}

		if ref != "" {
			g.localRefs[ref] = array
		}

		item, err := g.generate(*schema.Items, "")
		if err != nil {
			return nil, xerrors.Errorf("item: %w", err)
		}

		array.Item = item
		return array, nil

	case "number", "integer":
		return onret(&oas.Schema{
			Type:             oas.SchemaType(schema.Type),
			Format:           oas.Format(schema.Format),
			Description:      schema.Description,
			Ref:              ref,
			Minimum:          schema.Minimum,
			Maximum:          schema.Maximum,
			ExclusiveMinimum: schema.ExclusiveMinimum,
			ExclusiveMaximum: schema.ExclusiveMaximum,
			MultipleOf:       schema.MultipleOf,
		}), nil

	case "boolean":
		return onret(&oas.Schema{
			Type:        oas.Boolean,
			Format:      oas.Format(schema.Format),
			Description: schema.Description,
			Ref:         ref,
		}), nil

	case "string":
		return onret(&oas.Schema{
			Type:        oas.String,
			Format:      oas.Format(schema.Format),
			Description: schema.Description,
			Ref:         ref,
			MaxLength:   schema.MaxLength,
			MinLength:   schema.MinLength,
			Pattern:     schema.Pattern,
		}), nil

	case "":
		return onret(&oas.Schema{Type: oas.String}), nil

	default:
		return nil, xerrors.Errorf("unexpected schema type: '%s'", schema.Type)
	}
}

func (g *schemaGen) ref(ref string) (*oas.Schema, error) {
	const prefix = "#/components/schemas/"
	if !strings.HasPrefix(ref, prefix) {
		return nil, fmt.Errorf("invalid schema reference %q", ref)
	}

	if s, ok := g.globalRefs[ref]; ok {
		return s, nil
	}

	if s, ok := g.localRefs[ref]; ok {
		return s, nil
	}

	name := strings.TrimPrefix(ref, prefix)
	component, found := g.spec.Components.Schemas[name]
	if !found {
		return nil, xerrors.Errorf("component by reference %q not found", ref)
	}

	return g.generate(component, ref)
}

func (g *schemaGen) validateTypeFormat(typ, format string) error {
	switch typ {
	case "object":
		switch format {
		case "":
			return nil
		default:
			return xerrors.Errorf("unexpected object format: %q", format)
		}
	case "array":
		switch format {
		case "":
			return nil
		default:
			return xerrors.Errorf("unexpected array format: %q", format)
		}
	case "integer":
		switch format {
		case "int32", "int64", "":
			return nil
		default:
			return xerrors.Errorf("unexpected integer format: %q", format)
		}
	case "number":
		switch format {
		case "float", "double", "":
			return nil
		default:
			return xerrors.Errorf("unexpected number format: %q", format)
		}
	case "string":
		switch format {
		case "byte":
			return nil
		case "date-time", "time", "date":
			return nil
		case "duration":
			return nil
		case "uuid":
			return nil
		case "ipv4", "ipv6", "ip":
			return nil
		case "uri":
			return nil
		case "password", "":
			return nil
		default:
			// return nil, xerrors.Errorf("unexpected string format: %q", format)
			return nil
		}
	case "boolean":
		switch format {
		case "":
			return nil
		default:
			return xerrors.Errorf("unexpected bool format: %q", format)
		}
	default:
		return xerrors.Errorf("unexpected type: %q", typ)
	}
}
