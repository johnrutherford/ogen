package ast

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"unicode"

	"golang.org/x/xerrors"

	"github.com/ogen-go/ogen/validate"
)

type SchemaKind = string

const (
	KindStruct    SchemaKind = "struct"
	KindAlias     SchemaKind = "alias"
	KindPrimitive SchemaKind = "primitive"
	KindArray     SchemaKind = "array"
	KindEnum      SchemaKind = "enum"
	KindPointer   SchemaKind = "pointer"
	KindGeneric   SchemaKind = "generic"
)

type Validators struct {
	String validate.String
	Int    validate.Int
	Array  validate.Array
}

func (s Schema) FormatCustom() bool {
	switch s.Primitive {
	case Time:
		return true
	default:
		return false
	}
}

// JSONFields returns set of fields that should be encoded or decoded in json.
func (s Schema) JSONFields() []SchemaField {
	var fields []SchemaField
	for _, f := range s.Fields {
		if f.Tag == "-" {
			continue
		}
		if f.Type.BlankStruct() {
			continue
		}
		fields = append(fields, f)
	}
	return fields
}

func (s Schema) CanGeneric() bool {
	switch s.Kind {
	case KindPrimitive:
		if s.Primitive == EmptyStruct {
			return false
		}
	case KindArray:
		if s.Item.Primitive == Byte {
			return false
		}
	}
	return s.Is(KindPrimitive, KindEnum, KindStruct)
}

// NilSemantic specifies nil value semantics.
type NilSemantic string

// Possible nil value semantics.
const (
	NilInvalid  NilSemantic = "invalid"  // nil is invalid
	NilOptional NilSemantic = "optional" // nil is "no value"
	NilNull     NilSemantic = "null"     // nil is null
)

func (s Schema) BlankStruct() bool {
	if s.Primitive == EmptyStruct {
		return true
	}
	if s.PointerTo != nil && s.PointerTo.BlankStruct() {
		return true
	}
	return false
}

type Schema struct {
	Kind        SchemaKind
	Name        string
	Description string
	Doc         string
	Format      string
	NilSemantic NilSemantic

	GenericOf      *Schema
	GenericVariant GenericVariant

	AliasTo   *Schema
	PointerTo *Schema
	Primitive PrimitiveType
	Item      *Schema

	EnumValues []interface{}
	Fields     []SchemaField

	Implements map[*Interface]struct{}

	// Numeric validation.
	Validators Validators

	// String validation.
	// Pattern   string

	// Array validation.
	// UniqueItems bool

	// Struct validation.
	// MaxProperties *uint64
	// MinProperties *uint64
}

func (s Schema) canRawJSON() bool {
	if !s.Is(KindPrimitive, KindEnum) {
		return false
	}

	if s.IsNumeric() {
		return true
	}
	switch s.Primitive {
	case Bool, String:
		return true
	default:
		return false
	}
}

func (s Schema) JSONType() string {
	if s.IsNumeric() {
		return "NumberValue"
	}
	if s.IsArray() {
		return "ArrayValue"
	}
	if s.IsStruct() {
		return "ObjectValue"
	}
	switch s.Primitive {
	case Bool:
		return "BoolValue"
	case String, Time, Duration, UUID, IP, URL:
		return "StringValue"
	default:
		return ""
	}
}

// JSONHelper returns format name for handling json encoding or decoding.
//
// Mostly used for encoding or decoding of generics, like "json.WriteUUID",
// where UUID is JSONHelper.
func (s Schema) JSONHelper() string {
	if !s.Is(KindPrimitive, KindEnum) {
		return ""
	}

	switch s.Format {
	case "uuid":
		return "UUID"
	case "date":
		return "Date"
	case "time":
		return "Time"
	case "date-time":
		return "DateTime"
	case "duration":
		return "Duration"
	case "ip", "ipv4", "ipv6":
		return "IP"
	case "uri":
		return "URI"
	default:
		return ""
	}
}

func capitalize(s string) string {
	var v []rune
	for i, c := range s {
		if i == 0 {
			c = unicode.ToUpper(c)
		}
		v = append(v, c)
	}
	return string(v)
}

func (s Schema) jsonFn() string {
	if !s.canRawJSON() {
		return ""
	}
	return capitalize(s.Primitive.String())
}

// JSONWrite returns function name from json package that writes value.
func (s Schema) JSONWrite() string {
	if s.jsonFn() == "" {
		return ""
	}
	return "Write" + s.jsonFn()
}

// JSONRead returns function name from json package that reads value.
func (s Schema) JSONRead() string {
	if s.jsonFn() == "" {
		return ""
	}
	return "Read" + s.jsonFn()
}

func (s Schema) IsArray() bool   { return s.Is(KindArray) }
func (s Schema) IsEnum() bool    { return s.Is(KindEnum) }
func (s Schema) IsPointer() bool { return s.Is(KindPointer) }
func (s Schema) IsStruct() bool  { return s.Is(KindStruct) }
func (s Schema) IsAlias() bool   { return s.Is(KindAlias) }
func (s Schema) IsGeneric() bool { return s.Is(KindGeneric) }

// RecursiveTo reports whether target type is recursive to current.
func (s *Schema) RecursiveTo(to *Schema) bool {
	if to.Is(KindPrimitive, KindArray) {
		return false
	}
	return s.recurse(s, to)
}

// Equal reports whether two types are equal.
func (s *Schema) Equal(target *Schema) bool {
	if s.Kind != target.Kind {
		return false
	}
	if s.Primitive != target.Primitive {
		return false
	}
	return s.Type() == target.Type()
}

func (s *Schema) recurse(parent, target *Schema) bool {
	if target.Equal(parent) {
		return true
	}
	for _, f := range target.Fields {
		if parent.RecursiveTo(f.Type) {
			return true
		}
	}
	if s.GenericOf != nil {
		return s.GenericOf.RecursiveTo(target)
	}
	if target.GenericOf != nil {
		return s.recurse(parent, target.GenericOf)
	}
	return false
}

func (s *Schema) IsInteger() bool {
	switch s.Primitive {
	case Int, Int8, Int16, Int32, Int64,
		Uint, Uint8, Uint16, Uint32, Uint64:
		return true
	default:
		return false
	}
}

func (s *Schema) IsFloat() bool {
	switch s.Primitive {
	case Float32, Float64:
		return true
	default:
		return false
	}
}

func (s *Schema) IsNumeric() bool { return s.IsInteger() || s.IsFloat() }

func (s *Schema) NeedValidation() bool {
	return s.needValidation(map[*Schema]struct{}{})
}

func (s *Schema) needValidation(visited map[*Schema]struct{}) (result bool) {
	if s == nil {
		return false
	}

	if _, ok := visited[s]; ok {
		return false
	}

	visited[s] = struct{}{}

	switch s.Kind {
	case KindPrimitive:
		if s.IsNumeric() && s.Validators.Int.Set() {
			return true
		}
		if s.Validators.String.Set() {
			return true
		}
		return false
	case KindEnum:
		return true
	case KindAlias:
		return s.AliasTo.needValidation(visited)
	case KindPointer:
		if s.NilSemantic == NilInvalid {
			return true
		}
		return s.PointerTo.needValidation(visited)
	case KindGeneric:
		return s.GenericOf.needValidation(visited)
	case KindArray:
		if s.NilSemantic == NilInvalid {
			return true
		}
		if s.Validators.Array.Set() {
			return true
		}
		// Prevent infinite recursion.
		if s.Item == s {
			return false
		}
		return s.Item.needValidation(visited)
	case KindStruct:
		for _, f := range s.Fields {
			if f.Type.needValidation(visited) {
				return true
			}
		}
		return false
	default:
		panic("unreachable")
	}
}

type SchemaField struct {
	Name string
	Type *Schema
	Tag  string
}

func afterDot(v string) string {
	idx := strings.Index(v, ".")
	if idx > 0 {
		return v[idx+1:]
	}
	return v
}

func (s Schema) EncodeFn() string {
	if s.IsArray() && s.Item.EncodeFn() != "" {
		return s.Item.EncodeFn() + "Array"
	}
	switch s.Primitive {
	case Int, Int64, Int32, String, Bool, Float32, Float64:
		return capitalize(s.Primitive.String())
	case UUID, Time:
		return afterDot(s.Primitive.String())
	default:
		return ""
	}
}

func (s Schema) ToString() string {
	if s.EncodeFn() == "" {
		return ""
	}
	return s.EncodeFn() + "ToString"
}

func (s Schema) FromString() string {
	if s.EncodeFn() == "" {
		return ""
	}
	return "To" + s.EncodeFn()
}

func (s Schema) Type() string {
	switch s.Kind {
	case KindStruct:
		return s.Name
	case KindAlias:
		return s.Name
	case KindPrimitive:
		return s.Primitive.String()
	case KindGeneric:
		return s.Name
	case KindArray:
		return "[]" + s.Item.Type()
	case KindEnum:
		return s.Name
	case KindPointer:
		return "*" + s.PointerTo.Type()
	default:
		panic(fmt.Errorf("unexpected SchemaKind: %s", s.Kind))
	}
}

func (s Schema) Is(vs ...SchemaKind) bool {
	for _, v := range vs {
		if s.Kind == v {
			return true
		}
	}

	return false
}

func (s *Schema) Implement(iface *Interface) {
	if s.Is(KindPrimitive, KindArray, KindPointer) {
		panic("unreachable")
	}

	if s.Implements == nil {
		s.Implements = map[*Interface]struct{}{}
	}
	if iface.Implementations == nil {
		iface.Implementations = map[*Schema]struct{}{}
	}

	iface.Implementations[s] = struct{}{}
	s.Implements[iface] = struct{}{}
}

func (s *Schema) Unimplement(iface *Interface) {
	delete(iface.Implementations, s)
	delete(s.Implements, iface)
}

func (s *Schema) Methods() []string {
	ms := make(map[string]struct{})
	for iface := range s.Implements {
		for m := range iface.Methods {
			ms[m] = struct{}{}
		}
	}

	var result []string
	for m := range ms {
		result = append(result, m)
	}
	sort.Strings(result)
	return result
}

func Struct(name string) *Schema {
	return &Schema{
		Kind: KindStruct,
		Name: name,
	}
}

func Primitive(typ PrimitiveType) *Schema {
	return &Schema{
		Kind:      KindPrimitive,
		Primitive: typ,
	}
}

func Alias(name string, typ *Schema) *Schema {
	return &Schema{
		Kind:    KindAlias,
		Name:    name,
		AliasTo: typ,
	}
}

// Pointer makes new pointer type.
func Pointer(to *Schema, semantic NilSemantic) *Schema {
	return &Schema{
		Kind:        KindPointer,
		NilSemantic: semantic,
		PointerTo:   to,
	}
}

type GenericVariant struct {
	Nullable bool
	Optional bool
}

func (v GenericVariant) OnlyOptional() bool {
	return v.Optional && !v.Nullable
}

func (v GenericVariant) OnlyNullable() bool {
	return v.Nullable && !v.Optional
}

func (v GenericVariant) Name() string {
	var b strings.Builder
	if v.Optional {
		b.WriteString("Opt")
	}
	if v.Nullable {
		b.WriteString("Nil")
	}
	return b.String()
}

func (v GenericVariant) Any() bool {
	return v.Nullable || v.Optional
}

func Generic(name string, of *Schema, v GenericVariant) *Schema {
	name = v.Name() + name
	if of.IsArray() {
		name = name + "Array"
	}
	return &Schema{
		Name:           name,
		Kind:           KindGeneric,
		GenericOf:      of,
		GenericVariant: v,
	}
}

func Array(item *Schema) *Schema {
	return &Schema{
		Kind:        KindArray,
		Item:        item,
		NilSemantic: NilInvalid,
	}
}

func Enum(name string, typ PrimitiveType, rawValues []json.RawMessage) (*Schema, error) {
	var (
		values []interface{}
		uniq   = map[interface{}]struct{}{}
	)
	for _, raw := range rawValues {
		val, err := parseJSONValue(typ, raw)
		if err != nil {
			if xerrors.Is(err, errNullValue) {
				continue
			}
			return nil, xerrors.Errorf("parse value '%s': %w", raw, err)
		}

		if _, found := uniq[val]; found {
			return nil, xerrors.Errorf("duplicate enum value: '%v'", val)
		}

		uniq[val] = struct{}{}
		values = append(values, val)
	}

	return &Schema{
		Kind:       KindEnum,
		Name:       name,
		Primitive:  typ,
		EnumValues: values,
	}, nil
}

func Iface(name string) *Interface {
	return &Interface{
		Name:            name,
		Methods:         map[string]struct{}{},
		Implementations: map[*Schema]struct{}{},
	}
}

func CreateRequestBody() *RequestBody {
	return &RequestBody{
		Contents: map[string]*Schema{},
	}
}

func CreateResponse() *Response {
	return &Response{
		Contents: map[string]*Schema{},
	}
}
