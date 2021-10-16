// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/encoding/json"
	"github.com/ogen-go/ogen/types"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = math.Mod
	_ = types.Date{}
	_ = jsoniter.Config{}
	_ = validate.Int{}
)

// New returns new OptionalString with value set to v.
func NewOptionalString(v string) OptionalString {
	return OptionalString{
		Value: v,
		Set:   true,
	}
}

// OptionalString is optional string.
type OptionalString struct {
	Value string
	Set   bool
}

// IsSet returns true if string was set.
func (o OptionalString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptionalString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptionalString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of string to json stream.
func (o OptionalString) WriteJSON(js *jsoniter.Stream) {
	js.WriteString(o.Value)
}

// ReadJSON writes json value of string from json iterator.
func (o *OptionalString) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.StringValue:
		o.Set = true
		o.Value = i.ReadString()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalString", i.WhatIsNext())
	}
	return nil
}

// New returns new NilString with value set to v.
func NewNilString(v string) NilString {
	return NilString{
		Value: v,
	}
}

// NilString is nillable string.
type NilString struct {
	Value string
	Nil   bool
}

// SetTo sets value to v.
func (o *NilString) SetTo(v string) {
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o NilString) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o NilString) Get() (v string, ok bool) {
	if o.Nil {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of string to json stream.
func (o NilString) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteString(o.Value)
}

// ReadJSON writes json value of string from json iterator.
func (o *NilString) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.StringValue:
		o.Nil = false
		o.Value = i.ReadString()
		return i.Error
	case jsoniter.NilValue:
		var v string
		o.Value = v
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilString", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalNilString with value set to v.
func NewOptionalNilString(v string) OptionalNilString {
	return OptionalNilString{
		Value: v,
		Set:   true,
	}
}

// OptionalNilString is optional nillable string.
type OptionalNilString struct {
	Value string
	Set   bool
	Nil   bool
}

// IsSet returns true if string was set.
func (o OptionalNilString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalNilString) Reset() {
	var v string
	o.Value = v
	o.Set = false
	o.Nil = false
}

// SetTo sets value to v.
func (o *OptionalNilString) SetTo(v string) {
	o.Set = true
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o OptionalNilString) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o OptionalNilString) Get() (v string, ok bool) {
	if o.Nil {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of string to json stream.
func (o OptionalNilString) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteString(o.Value)
}

// ReadJSON writes json value of string from json iterator.
func (o *OptionalNilString) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.StringValue:
		o.Set = true
		o.Nil = false
		o.Value = i.ReadString()
		return i.Error
	case jsoniter.NilValue:
		var v string
		o.Value = v
		o.Set = true
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilString", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalInt with value set to v.
func NewOptionalInt(v int) OptionalInt {
	return OptionalInt{
		Value: v,
		Set:   true,
	}
}

// OptionalInt is optional int.
type OptionalInt struct {
	Value int
	Set   bool
}

// IsSet returns true if int was set.
func (o OptionalInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptionalInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptionalInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of int to json stream.
func (o OptionalInt) WriteJSON(js *jsoniter.Stream) {
	js.WriteInt(o.Value)
}

// ReadJSON writes json value of int from json iterator.
func (o *OptionalInt) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Set = true
		o.Value = i.ReadInt()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalInt", i.WhatIsNext())
	}
	return nil
}

// New returns new NilInt with value set to v.
func NewNilInt(v int) NilInt {
	return NilInt{
		Value: v,
	}
}

// NilInt is nillable int.
type NilInt struct {
	Value int
	Nil   bool
}

// SetTo sets value to v.
func (o *NilInt) SetTo(v int) {
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o NilInt) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o NilInt) Get() (v int, ok bool) {
	if o.Nil {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of int to json stream.
func (o NilInt) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteInt(o.Value)
}

// ReadJSON writes json value of int from json iterator.
func (o *NilInt) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Nil = false
		o.Value = i.ReadInt()
		return i.Error
	case jsoniter.NilValue:
		var v int
		o.Value = v
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilInt", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalNilInt with value set to v.
func NewOptionalNilInt(v int) OptionalNilInt {
	return OptionalNilInt{
		Value: v,
		Set:   true,
	}
}

// OptionalNilInt is optional nillable int.
type OptionalNilInt struct {
	Value int
	Set   bool
	Nil   bool
}

// IsSet returns true if int was set.
func (o OptionalNilInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalNilInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
	o.Nil = false
}

// SetTo sets value to v.
func (o *OptionalNilInt) SetTo(v int) {
	o.Set = true
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o OptionalNilInt) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o OptionalNilInt) Get() (v int, ok bool) {
	if o.Nil {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of int to json stream.
func (o OptionalNilInt) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteInt(o.Value)
}

// ReadJSON writes json value of int from json iterator.
func (o *OptionalNilInt) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Set = true
		o.Nil = false
		o.Value = i.ReadInt()
		return i.Error
	case jsoniter.NilValue:
		var v int
		o.Value = v
		o.Set = true
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilInt", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalInt32 with value set to v.
func NewOptionalInt32(v int32) OptionalInt32 {
	return OptionalInt32{
		Value: v,
		Set:   true,
	}
}

// OptionalInt32 is optional int32.
type OptionalInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if int32 was set.
func (o OptionalInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptionalInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptionalInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of int32 to json stream.
func (o OptionalInt32) WriteJSON(js *jsoniter.Stream) {
	js.WriteInt32(o.Value)
}

// ReadJSON writes json value of int32 from json iterator.
func (o *OptionalInt32) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Set = true
		o.Value = i.ReadInt32()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalInt32", i.WhatIsNext())
	}
	return nil
}

// New returns new NilInt32 with value set to v.
func NewNilInt32(v int32) NilInt32 {
	return NilInt32{
		Value: v,
	}
}

// NilInt32 is nillable int32.
type NilInt32 struct {
	Value int32
	Nil   bool
}

// SetTo sets value to v.
func (o *NilInt32) SetTo(v int32) {
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o NilInt32) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o NilInt32) Get() (v int32, ok bool) {
	if o.Nil {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of int32 to json stream.
func (o NilInt32) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteInt32(o.Value)
}

// ReadJSON writes json value of int32 from json iterator.
func (o *NilInt32) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Nil = false
		o.Value = i.ReadInt32()
		return i.Error
	case jsoniter.NilValue:
		var v int32
		o.Value = v
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilInt32", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalNilInt32 with value set to v.
func NewOptionalNilInt32(v int32) OptionalNilInt32 {
	return OptionalNilInt32{
		Value: v,
		Set:   true,
	}
}

// OptionalNilInt32 is optional nillable int32.
type OptionalNilInt32 struct {
	Value int32
	Set   bool
	Nil   bool
}

// IsSet returns true if int32 was set.
func (o OptionalNilInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalNilInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
	o.Nil = false
}

// SetTo sets value to v.
func (o *OptionalNilInt32) SetTo(v int32) {
	o.Set = true
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o OptionalNilInt32) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o OptionalNilInt32) Get() (v int32, ok bool) {
	if o.Nil {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of int32 to json stream.
func (o OptionalNilInt32) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteInt32(o.Value)
}

// ReadJSON writes json value of int32 from json iterator.
func (o *OptionalNilInt32) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Set = true
		o.Nil = false
		o.Value = i.ReadInt32()
		return i.Error
	case jsoniter.NilValue:
		var v int32
		o.Value = v
		o.Set = true
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilInt32", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalInt64 with value set to v.
func NewOptionalInt64(v int64) OptionalInt64 {
	return OptionalInt64{
		Value: v,
		Set:   true,
	}
}

// OptionalInt64 is optional int64.
type OptionalInt64 struct {
	Value int64
	Set   bool
}

// IsSet returns true if int64 was set.
func (o OptionalInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptionalInt64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptionalInt64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of int64 to json stream.
func (o OptionalInt64) WriteJSON(js *jsoniter.Stream) {
	js.WriteInt64(o.Value)
}

// ReadJSON writes json value of int64 from json iterator.
func (o *OptionalInt64) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Set = true
		o.Value = i.ReadInt64()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalInt64", i.WhatIsNext())
	}
	return nil
}

// New returns new NilInt64 with value set to v.
func NewNilInt64(v int64) NilInt64 {
	return NilInt64{
		Value: v,
	}
}

// NilInt64 is nillable int64.
type NilInt64 struct {
	Value int64
	Nil   bool
}

// SetTo sets value to v.
func (o *NilInt64) SetTo(v int64) {
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o NilInt64) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o NilInt64) Get() (v int64, ok bool) {
	if o.Nil {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of int64 to json stream.
func (o NilInt64) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteInt64(o.Value)
}

// ReadJSON writes json value of int64 from json iterator.
func (o *NilInt64) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Nil = false
		o.Value = i.ReadInt64()
		return i.Error
	case jsoniter.NilValue:
		var v int64
		o.Value = v
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilInt64", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalNilInt64 with value set to v.
func NewOptionalNilInt64(v int64) OptionalNilInt64 {
	return OptionalNilInt64{
		Value: v,
		Set:   true,
	}
}

// OptionalNilInt64 is optional nillable int64.
type OptionalNilInt64 struct {
	Value int64
	Set   bool
	Nil   bool
}

// IsSet returns true if int64 was set.
func (o OptionalNilInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalNilInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
	o.Nil = false
}

// SetTo sets value to v.
func (o *OptionalNilInt64) SetTo(v int64) {
	o.Set = true
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o OptionalNilInt64) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o OptionalNilInt64) Get() (v int64, ok bool) {
	if o.Nil {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of int64 to json stream.
func (o OptionalNilInt64) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteInt64(o.Value)
}

// ReadJSON writes json value of int64 from json iterator.
func (o *OptionalNilInt64) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Set = true
		o.Nil = false
		o.Value = i.ReadInt64()
		return i.Error
	case jsoniter.NilValue:
		var v int64
		o.Value = v
		o.Set = true
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilInt64", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalFloat32 with value set to v.
func NewOptionalFloat32(v float32) OptionalFloat32 {
	return OptionalFloat32{
		Value: v,
		Set:   true,
	}
}

// OptionalFloat32 is optional float32.
type OptionalFloat32 struct {
	Value float32
	Set   bool
}

// IsSet returns true if float32 was set.
func (o OptionalFloat32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalFloat32) Reset() {
	var v float32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptionalFloat32) SetTo(v float32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptionalFloat32) Get() (v float32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of float32 to json stream.
func (o OptionalFloat32) WriteJSON(js *jsoniter.Stream) {
	js.WriteFloat32(o.Value)
}

// ReadJSON writes json value of float32 from json iterator.
func (o *OptionalFloat32) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Set = true
		o.Value = i.ReadFloat32()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalFloat32", i.WhatIsNext())
	}
	return nil
}

// New returns new NilFloat32 with value set to v.
func NewNilFloat32(v float32) NilFloat32 {
	return NilFloat32{
		Value: v,
	}
}

// NilFloat32 is nillable float32.
type NilFloat32 struct {
	Value float32
	Nil   bool
}

// SetTo sets value to v.
func (o *NilFloat32) SetTo(v float32) {
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o NilFloat32) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o NilFloat32) Get() (v float32, ok bool) {
	if o.Nil {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of float32 to json stream.
func (o NilFloat32) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteFloat32(o.Value)
}

// ReadJSON writes json value of float32 from json iterator.
func (o *NilFloat32) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Nil = false
		o.Value = i.ReadFloat32()
		return i.Error
	case jsoniter.NilValue:
		var v float32
		o.Value = v
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilFloat32", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalNilFloat32 with value set to v.
func NewOptionalNilFloat32(v float32) OptionalNilFloat32 {
	return OptionalNilFloat32{
		Value: v,
		Set:   true,
	}
}

// OptionalNilFloat32 is optional nillable float32.
type OptionalNilFloat32 struct {
	Value float32
	Set   bool
	Nil   bool
}

// IsSet returns true if float32 was set.
func (o OptionalNilFloat32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalNilFloat32) Reset() {
	var v float32
	o.Value = v
	o.Set = false
	o.Nil = false
}

// SetTo sets value to v.
func (o *OptionalNilFloat32) SetTo(v float32) {
	o.Set = true
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o OptionalNilFloat32) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o OptionalNilFloat32) Get() (v float32, ok bool) {
	if o.Nil {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of float32 to json stream.
func (o OptionalNilFloat32) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteFloat32(o.Value)
}

// ReadJSON writes json value of float32 from json iterator.
func (o *OptionalNilFloat32) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Set = true
		o.Nil = false
		o.Value = i.ReadFloat32()
		return i.Error
	case jsoniter.NilValue:
		var v float32
		o.Value = v
		o.Set = true
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilFloat32", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalFloat64 with value set to v.
func NewOptionalFloat64(v float64) OptionalFloat64 {
	return OptionalFloat64{
		Value: v,
		Set:   true,
	}
}

// OptionalFloat64 is optional float64.
type OptionalFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if float64 was set.
func (o OptionalFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptionalFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptionalFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of float64 to json stream.
func (o OptionalFloat64) WriteJSON(js *jsoniter.Stream) {
	js.WriteFloat64(o.Value)
}

// ReadJSON writes json value of float64 from json iterator.
func (o *OptionalFloat64) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Set = true
		o.Value = i.ReadFloat64()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalFloat64", i.WhatIsNext())
	}
	return nil
}

// New returns new NilFloat64 with value set to v.
func NewNilFloat64(v float64) NilFloat64 {
	return NilFloat64{
		Value: v,
	}
}

// NilFloat64 is nillable float64.
type NilFloat64 struct {
	Value float64
	Nil   bool
}

// SetTo sets value to v.
func (o *NilFloat64) SetTo(v float64) {
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o NilFloat64) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o NilFloat64) Get() (v float64, ok bool) {
	if o.Nil {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of float64 to json stream.
func (o NilFloat64) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteFloat64(o.Value)
}

// ReadJSON writes json value of float64 from json iterator.
func (o *NilFloat64) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Nil = false
		o.Value = i.ReadFloat64()
		return i.Error
	case jsoniter.NilValue:
		var v float64
		o.Value = v
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilFloat64", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalNilFloat64 with value set to v.
func NewOptionalNilFloat64(v float64) OptionalNilFloat64 {
	return OptionalNilFloat64{
		Value: v,
		Set:   true,
	}
}

// OptionalNilFloat64 is optional nillable float64.
type OptionalNilFloat64 struct {
	Value float64
	Set   bool
	Nil   bool
}

// IsSet returns true if float64 was set.
func (o OptionalNilFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalNilFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
	o.Nil = false
}

// SetTo sets value to v.
func (o *OptionalNilFloat64) SetTo(v float64) {
	o.Set = true
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o OptionalNilFloat64) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o OptionalNilFloat64) Get() (v float64, ok bool) {
	if o.Nil {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of float64 to json stream.
func (o OptionalNilFloat64) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteFloat64(o.Value)
}

// ReadJSON writes json value of float64 from json iterator.
func (o *OptionalNilFloat64) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.NumberValue:
		o.Set = true
		o.Nil = false
		o.Value = i.ReadFloat64()
		return i.Error
	case jsoniter.NilValue:
		var v float64
		o.Value = v
		o.Set = true
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilFloat64", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalBool with value set to v.
func NewOptionalBool(v bool) OptionalBool {
	return OptionalBool{
		Value: v,
		Set:   true,
	}
}

// OptionalBool is optional bool.
type OptionalBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if bool was set.
func (o OptionalBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptionalBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptionalBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of bool to json stream.
func (o OptionalBool) WriteJSON(js *jsoniter.Stream) {
	js.WriteBool(o.Value)
}

// ReadJSON writes json value of bool from json iterator.
func (o *OptionalBool) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.BoolValue:
		o.Set = true
		o.Value = i.ReadBool()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalBool", i.WhatIsNext())
	}
	return nil
}

// New returns new NilBool with value set to v.
func NewNilBool(v bool) NilBool {
	return NilBool{
		Value: v,
	}
}

// NilBool is nillable bool.
type NilBool struct {
	Value bool
	Nil   bool
}

// SetTo sets value to v.
func (o *NilBool) SetTo(v bool) {
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o NilBool) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o NilBool) Get() (v bool, ok bool) {
	if o.Nil {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of bool to json stream.
func (o NilBool) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteBool(o.Value)
}

// ReadJSON writes json value of bool from json iterator.
func (o *NilBool) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.BoolValue:
		o.Nil = false
		o.Value = i.ReadBool()
		return i.Error
	case jsoniter.NilValue:
		var v bool
		o.Value = v
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading NilBool", i.WhatIsNext())
	}
	return nil
}

// New returns new OptionalNilBool with value set to v.
func NewOptionalNilBool(v bool) OptionalNilBool {
	return OptionalNilBool{
		Value: v,
		Set:   true,
	}
}

// OptionalNilBool is optional nillable bool.
type OptionalNilBool struct {
	Value bool
	Set   bool
	Nil   bool
}

// IsSet returns true if bool was set.
func (o OptionalNilBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptionalNilBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
	o.Nil = false
}

// SetTo sets value to v.
func (o *OptionalNilBool) SetTo(v bool) {
	o.Set = true
	o.Nil = false
	o.Value = v
}

// IsSet returns true if value is nil.
func (o OptionalNilBool) IsNil() bool { return o.Nil }

// Get returns value and boolean that denotes whether value was set.
func (o OptionalNilBool) Get() (v bool, ok bool) {
	if o.Nil {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// WriteJSON writes json value of bool to json stream.
func (o OptionalNilBool) WriteJSON(js *jsoniter.Stream) {
	if o.Nil {
		js.WriteNil()
		return
	}
	js.WriteBool(o.Value)
}

// ReadJSON writes json value of bool from json iterator.
func (o *OptionalNilBool) ReadJSON(i *jsoniter.Iterator) error {
	switch i.WhatIsNext() {
	case jsoniter.BoolValue:
		o.Set = true
		o.Nil = false
		o.Value = i.ReadBool()
		return i.Error
	case jsoniter.NilValue:
		var v bool
		o.Value = v
		o.Set = true
		o.Nil = true
		i.Skip()
		return i.Error
	default:
		return fmt.Errorf("unexpected type %d while reading OptionalNilBool", i.WhatIsNext())
	}
	return nil
}

func (s HelloWorld) WriteJSON(js *jsoniter.Stream) {
	js.WriteObjectStart()
	js.WriteObjectField("message")
	js.WriteString(s.Message)
	js.WriteObjectEnd()
}

func (s HelloWorld) encodeJSON() []byte {
	buf := new(bytes.Buffer)
	js := jsoniter.NewStream(jsoniter.ConfigDefault, buf, 1024)
	s.WriteJSON(js)
	_ = js.Flush()
	return buf.Bytes()
}

func (s *HelloWorld) decodeJSON(data []byte) error {
	i := jsoniter.NewIterator(jsoniter.ConfigDefault)
	i.ResetBytes(data)
	return s.ReadJSON(i)
}

// ReadJSON reads HelloWorld from json stream.
func (s *HelloWorld) ReadJSON(i *jsoniter.Iterator) error {
	return nil
}

func (s WorldObject) WriteJSON(js *jsoniter.Stream) {
	js.WriteObjectStart()
	js.WriteObjectField("id")
	js.WriteInt64(s.ID)
	js.WriteObjectField("randomNumber")
	js.WriteInt64(s.RandomNumber)
	js.WriteObjectEnd()
}

func (s WorldObject) encodeJSON() []byte {
	buf := new(bytes.Buffer)
	js := jsoniter.NewStream(jsoniter.ConfigDefault, buf, 1024)
	s.WriteJSON(js)
	_ = js.Flush()
	return buf.Bytes()
}

func (s *WorldObject) decodeJSON(data []byte) error {
	i := jsoniter.NewIterator(jsoniter.ConfigDefault)
	i.ResetBytes(data)
	return s.ReadJSON(i)
}

// ReadJSON reads WorldObject from json stream.
func (s *WorldObject) ReadJSON(i *jsoniter.Iterator) error {
	return nil
}
