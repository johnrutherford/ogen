// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s JaegerAnyOf) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.SizeLimit.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sizeLimit",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s JaegerAnyOfSizeLimit) Validate() error {
	switch s.Type {
	case IntJaegerAnyOfSizeLimit:
		return nil // no validation needed
	case StringJaegerAnyOfSizeLimit:
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        false,
			Hostname:     false,
			Regex:        regexMap["^(\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\\+|-)?(([0-9]+(\\.[0-9]*)?)|(\\.[0-9]+))))?$"],
		}).Validate(string(s.String)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s OneUUID) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:        true,
			Min:           1,
			MaxSet:        false,
			Max:           0,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: false,
			MultipleOf:    0,
		}).Validate(int64(s.Version)); err != nil {
			return errors.Wrap(err, "int")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "version",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
