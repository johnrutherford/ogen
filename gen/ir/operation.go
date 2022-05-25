package ir

import (
	"fmt"
	"sort"

	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/openapi"
)

type Operation struct {
	Name        string
	Description string
	PathParts   []*PathPart
	Params      []*Parameter
	Request     *Request
	Responses   *Responses
	Security    []*SecurityRequirement
	Spec        *openapi.Operation
}

func (op Operation) GoDoc() []string {
	return prettyDoc(op.Description)
}

type PathPart struct {
	Raw   string
	Param *Parameter
}

func (p PathPart) String() string {
	if p.Param != nil {
		return fmt.Sprintf("{%s}", p.Param.Spec.Name)
	}
	return p.Raw
}

type Parameter struct {
	Name string
	Type *Type
	Spec *openapi.Parameter
}

// Default returns default value of this field, if it is set.
func (op Parameter) Default() Default {
	var schema *jsonschema.Schema
	if spec := op.Spec; spec != nil {
		schema = spec.Schema
	}
	if schema != nil {
		return Default{
			Value: schema.Default,
			Set:   schema.DefaultSet,
		}
	}
	if typ := op.Type; typ != nil {
		return typ.Default()
	}
	return Default{}
}

func (op Parameter) GoDoc() []string {
	if op.Spec == nil {
		return nil
	}
	return prettyDoc(op.Spec.Description)
}

type Request struct {
	Type     *Type
	Contents map[ContentType]*Type
	Spec     *openapi.RequestBody
}

func (r Request) FormParameters() (params []Parameter) {
	ct := ContentTypeFormURLEncoded
	t, ok := r.Contents[ct]
	if !ok {
		panic(ct)
	}

	for _, f := range t.Fields {
		params = append(params, Parameter{
			Name: f.Name,
			Type: f.Type,
			Spec: f.Tag.Form,
		})
	}
	return params
}

type Content struct {
	ContentType ContentType
	Type        *Type
}

// ContentType of body.
type ContentType string

const (
	// ContentTypeJSON is ContentType for json.
	ContentTypeJSON ContentType = "application/json"
	// ContentTypeFormURLEncoded is ContentType for URL-encoded form.
	ContentTypeFormURLEncoded ContentType = "application/x-www-form-urlencoded"
	// ContentTypeOctetStream is ContentType for binary.
	ContentTypeOctetStream ContentType = "application/octet-stream"
)

func (t ContentType) String() string { return string(t) }

func (t ContentType) JSON() bool { return t == ContentTypeJSON }

func (t ContentType) FormURLEncode() bool { return t == ContentTypeFormURLEncoded }

func (t ContentType) OctetStream() bool { return t == ContentTypeOctetStream }

func (t ContentType) EncodedDataTypeGo() string {
	switch t {
	case ContentTypeJSON:
		return "*jx.Encoder"
	case ContentTypeOctetStream:
		return "io.Reader"
	default:
		return "io.Reader"
	}
}

func (t ContentType) Name() string {
	switch t {
	case ContentTypeJSON:
		return "JSON"
	case ContentTypeOctetStream:
		return "OctetStream"
	case ContentTypeFormURLEncoded:
		return "FormURLEncoded"
	default:
		return ""
	}
}

type Responses struct {
	Type       *Type
	StatusCode map[int]*Response
	Default    *Response
}

type Response struct {
	Wrapped   bool
	NoContent *Type
	Contents  map[ContentType]*Type
	Spec      *openapi.Response
}

func (s Response) ResponseInfo() []ResponseInfo {
	var result []ResponseInfo

	if noc := s.NoContent; noc != nil {
		result = append(result, ResponseInfo{
			Type:      noc,
			Default:   true,
			NoContent: true,
		})
	}
	for contentType, typ := range s.Contents {
		result = append(result, ResponseInfo{
			Type:        typ,
			Default:     true,
			ContentType: contentType,
		})
	}

	sort.SliceStable(result, func(i, j int) bool {
		l, r := result[i], result[j]
		// Default responses has zero status code.
		if l.Default {
			l.StatusCode = 999
		}
		if r.Default {
			r.StatusCode = 999
		}
		if l.StatusCode != r.StatusCode {
			return l.StatusCode < r.StatusCode
		}
		return string(l.ContentType) < string(r.ContentType)
	})

	return result
}
