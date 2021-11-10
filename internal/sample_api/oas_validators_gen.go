// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
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
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func (s Data) Validate() error {
	var failures []validate.FieldError
	if err := func() error {

		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        true,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Email)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if err := func() error {

		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        false,
			Hostname:     true,
			Regex:        nil,
		}).Validate(string(s.Hostname)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "hostname",
			Error: err,
		})
	}
	if err := func() error {

		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        false,
			Hostname:     false,
			Regex:        regexp.MustCompile(`^\d-\d$`),
		}).Validate(string(s.Format)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "format",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s Pet) Validate() error {
	var failures []validate.FieldError
	if err := func() error {

		if s.Primary == nil {
			return nil // optional
		}
		if err := func() error {

			if err := s.Primary.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "pointer")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "primary",
			Error: err,
		})
	}
	if err := func() error {

		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       true,
			Max:          100000,
			MinExclusive: false,
			MaxExclusive: false,
		}).Validate(int64(s.ID)); err != nil {
			return errors.Wrap(err, "int")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "id",
			Error: err,
		})
	}
	if err := func() error {

		if err := (validate.String{
			MinLength:    4,
			MinLengthSet: true,
			MaxLength:    24,
			MaxLengthSet: true,
			Email:        false,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Name)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "name",
			Error: err,
		})
	}
	if err := func() error {

		if s.Type.Set {
			if err := func() error {

				if err := s.Type.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil

		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "type",
			Error: err,
		})
	}
	if err := func() error {

		if err := s.Kind.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "kind",
			Error: err,
		})
	}
	if err := func() error {

		if err := (validate.Array{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    255,
			MaxLengthSet: true,
		}).ValidateLength(len(s.Friends)); err != nil {
			return errors.Wrap(err, "array")
		}
		var failures []validate.FieldError
		for i, elem := range s.Friends {
			if err := func() error {

				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "friends",
			Error: err,
		})
	}
	if err := func() error {

		if s.Next.Set {
			if err := func() error {

				if err := s.Next.Value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil

		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "next",
			Error: err,
		})
	}
	if err := func() error {

		if s.TestInteger1.Set {
			if err := func() error {

				if err := (validate.Int{
					MinSet:       false,
					Min:          0,
					MaxSet:       false,
					Max:          0,
					MinExclusive: false,
					MaxExclusive: false,
				}).Validate(int64(s.TestInteger1.Value)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil

		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "testInteger1",
			Error: err,
		})
	}
	if err := func() error {

		if s.TestFloat1.Set {
			if err := func() error {

				if err := (validate.Int{
					MinSet:       true,
					Min:          15,
					MaxSet:       false,
					Max:          0,
					MinExclusive: false,
					MaxExclusive: false,
				}).Validate(int64(s.TestFloat1.Value)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil

		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "testFloat1",
			Error: err,
		})
	}
	if err := func() error {

		var failures []validate.FieldError
		for i, elem := range s.TestArray1 {
			if err := func() error {
				if elem == nil {
					return errors.New("nil is invalid value")
				}

				if err := (validate.Array{
					MinLength:    0,
					MinLengthSet: false,
					MaxLength:    16,
					MaxLengthSet: true,
				}).ValidateLength(len(elem)); err != nil {
					return errors.Wrap(err, "array")
				}
				var failures []validate.FieldError
				for i, elem := range elem {
					if err := func() error {

						if err := (validate.String{
							MinLength:    0,
							MinLengthSet: false,
							MaxLength:    255,
							MaxLengthSet: true,
							Email:        false,
							Hostname:     false,
							Regex:        nil,
						}).Validate(string(elem)); err != nil {
							return errors.Wrap(err, "string")
						}
						return nil
					}(); err != nil {
						failures = append(failures, validate.FieldError{
							Name:  fmt.Sprintf("[%d]", i),
							Error: err,
						})
					}
				}
				if len(failures) > 0 {
					return &validate.Error{Fields: failures}
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "testArray1",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s PetKind) Validate() error {
	switch s {
	case "big":
		return nil
	case "smol":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
func (s PetName) Validate() error {

	if err := (validate.String{
		MinLength:    6,
		MinLengthSet: true,
		MaxLength:    0,
		MaxLengthSet: false,
		Email:        false,
		Hostname:     false,
		Regex:        nil,
	}).Validate(string(s)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}
func (s PetType) Validate() error {
	switch s {
	case "fifa":
		return nil
	case "fofa":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
