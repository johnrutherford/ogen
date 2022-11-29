// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// TestFormURLEncoded implements testFormURLEncoded operation.
//
// POST /testFormURLEncoded
func (UnimplementedHandler) TestFormURLEncoded(ctx context.Context, req *TestForm) (r *TestFormURLEncodedOK, _ error) {
	return r, ht.ErrNotImplemented
}

// TestMultipart implements testMultipart operation.
//
// POST /testMultipart
func (UnimplementedHandler) TestMultipart(ctx context.Context, req *TestForm) (r *TestMultipartOK, _ error) {
	return r, ht.ErrNotImplemented
}

// TestMultipartUpload implements testMultipartUpload operation.
//
// POST /testMultipartUpload
func (UnimplementedHandler) TestMultipartUpload(ctx context.Context, req *TestMultipartUploadReqForm) (r *TestMultipartUploadOK, _ error) {
	return r, ht.ErrNotImplemented
}

// TestShareFormSchema implements testShareFormSchema operation.
//
// POST /testShareFormSchema
func (UnimplementedHandler) TestShareFormSchema(ctx context.Context, req TestShareFormSchemaReq) (r *TestShareFormSchemaOK, _ error) {
	return r, ht.ErrNotImplemented
}
