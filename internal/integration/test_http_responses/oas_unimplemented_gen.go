// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AnyContentTypeBinaryStringSchema implements anyContentTypeBinaryStringSchema operation.
//
// GET /anyContentTypeBinaryStringSchema
func (UnimplementedHandler) AnyContentTypeBinaryStringSchema(ctx context.Context) (r AnyContentTypeBinaryStringSchemaOK, _ error) {
	return r, ht.ErrNotImplemented
}

// AnyContentTypeBinaryStringSchemaDefault implements anyContentTypeBinaryStringSchemaDefault operation.
//
// GET /anyContentTypeBinaryStringSchemaDefault
func (UnimplementedHandler) AnyContentTypeBinaryStringSchemaDefault(ctx context.Context) (r *AnyContentTypeBinaryStringSchemaDefaultDefStatusCode, _ error) {
	return r, ht.ErrNotImplemented
}

// Combined implements combined operation.
//
// GET /combined
func (UnimplementedHandler) Combined(ctx context.Context, params CombinedParams) (r CombinedRes, _ error) {
	return r, ht.ErrNotImplemented
}

// Headers200 implements headers200 operation.
//
// GET /headers200
func (UnimplementedHandler) Headers200(ctx context.Context) (r *Headers200OK, _ error) {
	return r, ht.ErrNotImplemented
}

// HeadersCombined implements headersCombined operation.
//
// GET /headersCombined
func (UnimplementedHandler) HeadersCombined(ctx context.Context, params HeadersCombinedParams) (r HeadersCombinedRes, _ error) {
	return r, ht.ErrNotImplemented
}

// HeadersDefault implements headersDefault operation.
//
// GET /headersDefault
func (UnimplementedHandler) HeadersDefault(ctx context.Context) (r *HeadersDefaultDef, _ error) {
	return r, ht.ErrNotImplemented
}

// HeadersJSON implements headersJSON operation.
//
// GET /headersJSON
func (UnimplementedHandler) HeadersJSON(ctx context.Context) (r *HeadersJSONOK, _ error) {
	return r, ht.ErrNotImplemented
}

// HeadersPattern implements headersPattern operation.
//
// GET /headersPattern
func (UnimplementedHandler) HeadersPattern(ctx context.Context) (r *HeadersPattern4XX, _ error) {
	return r, ht.ErrNotImplemented
}

// IntersectPatternCode implements intersectPatternCode operation.
//
// If a response is defined using an explicit code, the explicit code definition takes precedence
// over the range definition for that code.
//
// GET /intersectPatternCode
func (UnimplementedHandler) IntersectPatternCode(ctx context.Context, params IntersectPatternCodeParams) (r IntersectPatternCodeRes, _ error) {
	return r, ht.ErrNotImplemented
}

// MultipleGenericResponses implements multipleGenericResponses operation.
//
// GET /multipleGenericResponses
func (UnimplementedHandler) MultipleGenericResponses(ctx context.Context) (r MultipleGenericResponsesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// OctetStreamBinaryStringSchema implements octetStreamBinaryStringSchema operation.
//
// GET /octetStreamBinaryStringSchema
func (UnimplementedHandler) OctetStreamBinaryStringSchema(ctx context.Context) (r OctetStreamBinaryStringSchemaOK, _ error) {
	return r, ht.ErrNotImplemented
}

// OctetStreamEmptySchema implements octetStreamEmptySchema operation.
//
// GET /octetStreamEmptySchema
func (UnimplementedHandler) OctetStreamEmptySchema(ctx context.Context) (r OctetStreamEmptySchemaOK, _ error) {
	return r, ht.ErrNotImplemented
}

// TextPlainBinaryStringSchema implements textPlainBinaryStringSchema operation.
//
// GET /textPlainBinaryStringSchema
func (UnimplementedHandler) TextPlainBinaryStringSchema(ctx context.Context) (r TextPlainBinaryStringSchemaOK, _ error) {
	return r, ht.ErrNotImplemented
}
