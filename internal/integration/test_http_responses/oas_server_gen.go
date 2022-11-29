// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AnyContentTypeBinaryStringSchema implements anyContentTypeBinaryStringSchema operation.
	//
	// GET /anyContentTypeBinaryStringSchema
	AnyContentTypeBinaryStringSchema(ctx context.Context) (AnyContentTypeBinaryStringSchemaOK, error)
	// AnyContentTypeBinaryStringSchemaDefault implements anyContentTypeBinaryStringSchemaDefault operation.
	//
	// GET /anyContentTypeBinaryStringSchemaDefault
	AnyContentTypeBinaryStringSchemaDefault(ctx context.Context) (*AnyContentTypeBinaryStringSchemaDefaultDefStatusCode, error)
	// Combined implements combined operation.
	//
	// GET /combined
	Combined(ctx context.Context, params CombinedParams) (CombinedRes, error)
	// Headers200 implements headers200 operation.
	//
	// GET /headers200
	Headers200(ctx context.Context) (*Headers200OK, error)
	// HeadersCombined implements headersCombined operation.
	//
	// GET /headersCombined
	HeadersCombined(ctx context.Context, params HeadersCombinedParams) (HeadersCombinedRes, error)
	// HeadersDefault implements headersDefault operation.
	//
	// GET /headersDefault
	HeadersDefault(ctx context.Context) (*HeadersDefaultDef, error)
	// HeadersPattern implements headersPattern operation.
	//
	// GET /headersPattern
	HeadersPattern(ctx context.Context) (*HeadersPattern4XX, error)
	// IntersectPatternCode implements intersectPatternCode operation.
	//
	// If a response is defined using an explicit code, the explicit code definition takes precedence
	// over the range definition for that code.
	//
	// GET /intersectPatternCode
	IntersectPatternCode(ctx context.Context, params IntersectPatternCodeParams) (IntersectPatternCodeRes, error)
	// MultipleGenericResponses implements multipleGenericResponses operation.
	//
	// GET /multipleGenericResponses
	MultipleGenericResponses(ctx context.Context) (MultipleGenericResponsesRes, error)
	// OctetStreamBinaryStringSchema implements octetStreamBinaryStringSchema operation.
	//
	// GET /octetStreamBinaryStringSchema
	OctetStreamBinaryStringSchema(ctx context.Context) (OctetStreamBinaryStringSchemaOK, error)
	// OctetStreamEmptySchema implements octetStreamEmptySchema operation.
	//
	// GET /octetStreamEmptySchema
	OctetStreamEmptySchema(ctx context.Context) (OctetStreamEmptySchemaOK, error)
	// TextPlainBinaryStringSchema implements textPlainBinaryStringSchema operation.
	//
	// GET /textPlainBinaryStringSchema
	TextPlainBinaryStringSchema(ctx context.Context) (TextPlainBinaryStringSchemaOK, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
