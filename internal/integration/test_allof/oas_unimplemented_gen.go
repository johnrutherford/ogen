// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// NullableStrings implements nullableStrings operation.
//
// Nullable strings.
//
// POST /nullableStrings
func (UnimplementedHandler) NullableStrings(ctx context.Context, req NilString) error {
	return ht.ErrNotImplemented
}

// ObjectsWithConflictingArrayProperty implements objectsWithConflictingArrayProperty operation.
//
// Objects with conflicting array property.
//
// POST /objectsWithConflictingArrayProperty
func (UnimplementedHandler) ObjectsWithConflictingArrayProperty(ctx context.Context, req *ObjectsWithConflictingArrayPropertyReq) error {
	return ht.ErrNotImplemented
}

// ObjectsWithConflictingProperties implements objectsWithConflictingProperties operation.
//
// Objects with conflicting properties.
//
// POST /objectsWithConflictingProperties
func (UnimplementedHandler) ObjectsWithConflictingProperties(ctx context.Context, req *ObjectsWithConflictingPropertiesReq) error {
	return ht.ErrNotImplemented
}

// ReferencedAllof implements referencedAllof operation.
//
// Referenced allOf.
//
// POST /referencedAllof
func (UnimplementedHandler) ReferencedAllof(ctx context.Context, req ReferencedAllofReq) error {
	return ht.ErrNotImplemented
}

// ReferencedAllofOptional implements referencedAllofOptional operation.
//
// Referenced allOf, but requestBody is not required.
//
// POST /referencedAllofOptional
func (UnimplementedHandler) ReferencedAllofOptional(ctx context.Context, req ReferencedAllofOptionalReq) error {
	return ht.ErrNotImplemented
}

// SimpleInteger implements simpleInteger operation.
//
// Simple integers with validation.
//
// POST /simpleInteger
func (UnimplementedHandler) SimpleInteger(ctx context.Context, req int) error {
	return ht.ErrNotImplemented
}

// SimpleObjects implements simpleObjects operation.
//
// Simple objects.
//
// POST /simpleObjects
func (UnimplementedHandler) SimpleObjects(ctx context.Context, req *SimpleObjectsReq) error {
	return ht.ErrNotImplemented
}

// StringsNotype implements stringsNotype operation.
//
// POST /stringsNotype
func (UnimplementedHandler) StringsNotype(ctx context.Context, req NilString) error {
	return ht.ErrNotImplemented
}
