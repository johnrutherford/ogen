// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

// Ref: #/components/schemas/Location
type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// GetLat returns the value of Lat.
func (s *Location) GetLat() float64 {
	return s.Lat
}

// GetLon returns the value of Lon.
func (s *Location) GetLon() float64 {
	return s.Lon
}

// SetLat sets the value of Lat.
func (s *Location) SetLat(val float64) {
	s.Lat = val
}

// SetLon sets the value of Lon.
func (s *Location) SetLon(val float64) {
	s.Lon = val
}

// NewNilString returns new NilString with value set to v.
func NewNilString(v string) NilString {
	return NilString{
		Value: v,
	}
}

// NilString is nullable string.
type NilString struct {
	Value string
	Null  bool
}

// SetTo sets value to v.
func (o *NilString) SetTo(v string) {
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o NilString) IsNull() bool { return o.Null }

// Get returns value and boolean that denotes whether value was set.
func (o NilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NullableStringsOK is response for NullableStrings operation.
type NullableStringsOK struct{}

// ObjectsWithConflictingArrayPropertyOK is response for ObjectsWithConflictingArrayProperty operation.
type ObjectsWithConflictingArrayPropertyOK struct{}

// Merged schema.
type ObjectsWithConflictingArrayPropertyReq struct {
	// Merged property.
	Foo []int `json:"foo"`
	Bar int   `json:"bar"`
}

// GetFoo returns the value of Foo.
func (s *ObjectsWithConflictingArrayPropertyReq) GetFoo() []int {
	return s.Foo
}

// GetBar returns the value of Bar.
func (s *ObjectsWithConflictingArrayPropertyReq) GetBar() int {
	return s.Bar
}

// SetFoo sets the value of Foo.
func (s *ObjectsWithConflictingArrayPropertyReq) SetFoo(val []int) {
	s.Foo = val
}

// SetBar sets the value of Bar.
func (s *ObjectsWithConflictingArrayPropertyReq) SetBar(val int) {
	s.Bar = val
}

// ObjectsWithConflictingPropertiesOK is response for ObjectsWithConflictingProperties operation.
type ObjectsWithConflictingPropertiesOK struct{}

// Merged schema.
type ObjectsWithConflictingPropertiesReq struct {
	// Merged property.
	Foo string `json:"foo"`
	Bar OptInt `json:"bar"`
}

// GetFoo returns the value of Foo.
func (s *ObjectsWithConflictingPropertiesReq) GetFoo() string {
	return s.Foo
}

// GetBar returns the value of Bar.
func (s *ObjectsWithConflictingPropertiesReq) GetBar() OptInt {
	return s.Bar
}

// SetFoo sets the value of Foo.
func (s *ObjectsWithConflictingPropertiesReq) SetFoo(val string) {
	s.Foo = val
}

// SetBar sets the value of Bar.
func (s *ObjectsWithConflictingPropertiesReq) SetBar(val OptInt) {
	s.Bar = val
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type ReferencedAllofApplicationJSON Robot

func (*ReferencedAllofApplicationJSON) referencedAllofReq() {}

type ReferencedAllofMultipartFormData Robot

func (*ReferencedAllofMultipartFormData) referencedAllofReq() {}

// ReferencedAllofOK is response for ReferencedAllof operation.
type ReferencedAllofOK struct{}

type ReferencedAllofOptionalApplicationJSON Robot

func (*ReferencedAllofOptionalApplicationJSON) referencedAllofOptionalReq() {}

type ReferencedAllofOptionalMultipartFormData Robot

func (*ReferencedAllofOptionalMultipartFormData) referencedAllofOptionalReq() {}

// ReferencedAllofOptionalOK is response for ReferencedAllofOptional operation.
type ReferencedAllofOptionalOK struct{}

type ReferencedAllofOptionalReqEmptyBody struct{}

func (*ReferencedAllofOptionalReqEmptyBody) referencedAllofOptionalReq() {}

// Merged schema.
// Ref: #/components/schemas/Robot
type Robot struct {
	State    RobotState `json:"state"`
	ID       uuid.UUID  `json:"id"`
	Location Location   `json:"location"`
}

// GetState returns the value of State.
func (s *Robot) GetState() RobotState {
	return s.State
}

// GetID returns the value of ID.
func (s *Robot) GetID() uuid.UUID {
	return s.ID
}

// GetLocation returns the value of Location.
func (s *Robot) GetLocation() Location {
	return s.Location
}

// SetState sets the value of State.
func (s *Robot) SetState(val RobotState) {
	s.State = val
}

// SetID sets the value of ID.
func (s *Robot) SetID(val uuid.UUID) {
	s.ID = val
}

// SetLocation sets the value of Location.
func (s *Robot) SetLocation(val Location) {
	s.Location = val
}

type RobotState string

const (
	RobotStateOn  RobotState = "on"
	RobotStateOff RobotState = "off"
)

// MarshalText implements encoding.TextMarshaler.
func (s RobotState) MarshalText() ([]byte, error) {
	switch s {
	case RobotStateOn:
		return []byte(s), nil
	case RobotStateOff:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *RobotState) UnmarshalText(data []byte) error {
	switch RobotState(data) {
	case RobotStateOn:
		*s = RobotStateOn
		return nil
	case RobotStateOff:
		*s = RobotStateOff
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// SimpleIntegerOK is response for SimpleInteger operation.
type SimpleIntegerOK struct{}

// SimpleObjectsOK is response for SimpleObjects operation.
type SimpleObjectsOK struct{}

// Merged schema.
type SimpleObjectsReq struct {
	Foo OptString `json:"foo"`
	Bar OptBool   `json:"bar"`
}

// GetFoo returns the value of Foo.
func (s *SimpleObjectsReq) GetFoo() OptString {
	return s.Foo
}

// GetBar returns the value of Bar.
func (s *SimpleObjectsReq) GetBar() OptBool {
	return s.Bar
}

// SetFoo sets the value of Foo.
func (s *SimpleObjectsReq) SetFoo(val OptString) {
	s.Foo = val
}

// SetBar sets the value of Bar.
func (s *SimpleObjectsReq) SetBar(val OptBool) {
	s.Bar = val
}
