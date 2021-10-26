// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
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
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

func (s *PetGetResOKApplicationJSON) Validate() error {
	var failures []validate.FieldError
	{
		// Validate 'id' property.
		validator := validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       true,
			Max:          100000,
			MinExclusive: false,
			MaxExclusive: false,
		}
		if err := validator.Validate(int64(s.ID)); err != nil {
			failures = append(failures, validate.FieldError{Name: "id", Error: err})
		}
	}
	{
		// Validate 'name' property.
		validator := validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    24,
			MaxLengthSet: true,
		}
		if err := validator.Validate(string(s.Name)); err != nil {
			failures = append(failures, validate.FieldError{Name: "name", Error: err})
		}
	}
	if s.Primary == nil {
		return &validate.Error{
			Fields: append(failures, validate.FieldError{
				Name:  "primary",
				Error: fmt.Errorf("required, can't be nil"),
			}),
		}
	}
	{
		// Validate 'testFloat1' property.
		validator := validate.Int{
			MinSet:       true,
			Min:          15,
			MaxSet:       false,
			Max:          0,
			MinExclusive: false,
			MaxExclusive: false,
		}
		if err := validator.Validate(int64(s.TestFloat1)); err != nil {
			failures = append(failures, validate.FieldError{Name: "testFloat1", Error: err})
		}
	}
	{
		// Validate 'testInteger1' property.
		validator := validate.Int{
			MinSet:       false,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: false,
			MaxExclusive: false,
		}
		if err := validator.Validate(int64(s.TestInteger1)); err != nil {
			failures = append(failures, validate.FieldError{Name: "testInteger1", Error: err})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
