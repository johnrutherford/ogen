// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/google/uuid"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DataGetFormatParams is parameters of dataGetFormat operation.
type DataGetFormatParams struct {
	ID  int
	Foo string
	Bar string
	Baz string
	Kek string
}

func unpackDataGetFormatParams(packed middleware.Parameters) (params DataGetFormatParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	{
		key := middleware.ParameterKey{
			Name: "foo",
			In:   "path",
		}
		params.Foo = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "bar",
			In:   "path",
		}
		params.Bar = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "baz",
			In:   "path",
		}
		params.Baz = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "kek",
			In:   "path",
		}
		params.Kek = packed[key].(string)
	}
	return params
}

func decodeDataGetFormatParams(args [5]string, r *http.Request) (params DataGetFormatParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
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
				}).Validate(int64(params.ID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: foo.
	if err := func() error {
		param, err := url.PathUnescape(args[1])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "foo",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Foo = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Foo)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "foo",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: bar.
	if err := func() error {
		param, err := url.PathUnescape(args[2])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "bar",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Bar = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Bar)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "bar",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: baz.
	if err := func() error {
		param, err := url.PathUnescape(args[3])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "baz",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Baz = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Baz)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "baz",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: kek.
	if err := func() error {
		param, err := url.PathUnescape(args[4])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "kek",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Kek = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.String{
					MinLength:    1,
					MinLengthSet: true,
					MaxLength:    0,
					MaxLengthSet: false,
					Email:        false,
					Hostname:     false,
					Regex:        nil,
				}).Validate(string(params.Kek)); err != nil {
					return errors.Wrap(err, "string")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "kek",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DefaultTestParams is parameters of defaultTest operation.
type DefaultTestParams struct {
	Default OptInt32
}

func unpackDefaultTestParams(packed middleware.Parameters) (params DefaultTestParams) {
	{
		key := middleware.ParameterKey{
			Name: "default",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Default = v.(OptInt32)
		}
	}
	return params
}

func decodeDefaultTestParams(args [0]string, r *http.Request) (params DefaultTestParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: default.
	{
		val := int32(10)
		params.Default.SetTo(val)
	}
	// Decode query: default.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "default",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotDefaultVal int32
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt32(val)
					if err != nil {
						return err
					}

					paramsDotDefaultVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Default.SetTo(paramsDotDefaultVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "default",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// FoobarGetParams is parameters of foobarGet operation.
type FoobarGetParams struct {
	// InlinedParam.
	InlinedParam int64
	// Number of items to skip.
	Skip int32
}

func unpackFoobarGetParams(packed middleware.Parameters) (params FoobarGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "inlinedParam",
			In:   "query",
		}
		params.InlinedParam = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "skip",
			In:   "query",
		}
		params.Skip = packed[key].(int32)
	}
	return params
}

func decodeFoobarGetParams(args [0]string, r *http.Request) (params FoobarGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: inlinedParam.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "inlinedParam",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.InlinedParam = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "inlinedParam",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: skip.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "skip",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt32(val)
				if err != nil {
					return err
				}

				params.Skip = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "skip",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// PetFriendsNamesByIDParams is parameters of petFriendsNamesByID operation.
type PetFriendsNamesByIDParams struct {
	// Pet ID.
	ID int
}

func unpackPetFriendsNamesByIDParams(packed middleware.Parameters) (params PetFriendsNamesByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodePetFriendsNamesByIDParams(args [1]string, r *http.Request) (params PetFriendsNamesByIDParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PetGetParams is parameters of petGet operation.
type PetGetParams struct {
	// ID of pet.
	PetID int64
	// Tags of pets.
	XTags []uuid.UUID
	// Pet scopes.
	XScope []string
	// Token.
	Token string
}

func unpackPetGetParams(packed middleware.Parameters) (params PetGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "petID",
			In:   "query",
		}
		params.PetID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "X-Tags",
			In:   "header",
		}
		params.XTags = packed[key].([]uuid.UUID)
	}
	{
		key := middleware.ParameterKey{
			Name: "X-Scope",
			In:   "header",
		}
		params.XScope = packed[key].([]string)
	}
	{
		key := middleware.ParameterKey{
			Name: "token",
			In:   "query",
		}
		params.Token = packed[key].(string)
	}
	return params
}

func decodePetGetParams(args [0]string, r *http.Request) (params PetGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	h := uri.NewHeaderDecoder(r.Header)
	// Decode query: petID.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "petID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PetID = c
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1337,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.PetID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "petID",
			In:   "query",
			Err:  err,
		}
	}
	// Decode header: X-Tags.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Tags",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotXTagsVal uuid.UUID
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(val)
						if err != nil {
							return err
						}

						paramsDotXTagsVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.XTags = append(params.XTags, paramsDotXTagsVal)
					return nil
				})
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.XTags == nil {
					return errors.New("nil is invalid value")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Tags",
			In:   "header",
			Err:  err,
		}
	}
	// Decode header: X-Scope.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Scope",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				return d.DecodeArray(func(d uri.Decoder) error {
					var paramsDotXScopeVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						paramsDotXScopeVal = c
						return nil
					}(); err != nil {
						return err
					}
					params.XScope = append(params.XScope, paramsDotXScopeVal)
					return nil
				})
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.XScope == nil {
					return errors.New("nil is invalid value")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Scope",
			In:   "header",
			Err:  err,
		}
	}
	// Decode query: token.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "token",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Token = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "token",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// PetGetAvatarByIDParams is parameters of petGetAvatarByID operation.
type PetGetAvatarByIDParams struct {
	// ID of pet.
	PetID int64
}

func unpackPetGetAvatarByIDParams(packed middleware.Parameters) (params PetGetAvatarByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "petID",
			In:   "query",
		}
		params.PetID = packed[key].(int64)
	}
	return params
}

func decodePetGetAvatarByIDParams(args [0]string, r *http.Request) (params PetGetAvatarByIDParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: petID.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "petID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PetID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "petID",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// PetGetAvatarByNameParams is parameters of petGetAvatarByName operation.
type PetGetAvatarByNameParams struct {
	// Name of pet.
	Name string
}

func unpackPetGetAvatarByNameParams(packed middleware.Parameters) (params PetGetAvatarByNameParams) {
	{
		key := middleware.ParameterKey{
			Name: "name",
			In:   "path",
		}
		params.Name = packed[key].(string)
	}
	return params
}

func decodePetGetAvatarByNameParams(args [1]string, r *http.Request) (params PetGetAvatarByNameParams, _ error) {
	// Decode path: name.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "name",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Name = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "name",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PetGetByNameParams is parameters of petGetByName operation.
type PetGetByNameParams struct {
	// Name of pet.
	Name string
}

func unpackPetGetByNameParams(packed middleware.Parameters) (params PetGetByNameParams) {
	{
		key := middleware.ParameterKey{
			Name: "name",
			In:   "path",
		}
		params.Name = packed[key].(string)
	}
	return params
}

func decodePetGetByNameParams(args [1]string, r *http.Request) (params PetGetByNameParams, _ error) {
	// Decode path: name.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "name",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Name = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "name",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PetNameByIDParams is parameters of petNameByID operation.
type PetNameByIDParams struct {
	// Pet ID.
	ID int
}

func unpackPetNameByIDParams(packed middleware.Parameters) (params PetNameByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodePetNameByIDParams(args [1]string, r *http.Request) (params PetNameByIDParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PetUploadAvatarByIDParams is parameters of petUploadAvatarByID operation.
type PetUploadAvatarByIDParams struct {
	// ID of pet.
	PetID int64
}

func unpackPetUploadAvatarByIDParams(packed middleware.Parameters) (params PetUploadAvatarByIDParams) {
	{
		key := middleware.ParameterKey{
			Name: "petID",
			In:   "query",
		}
		params.PetID = packed[key].(int64)
	}
	return params
}

func decodePetUploadAvatarByIDParams(args [0]string, r *http.Request) (params PetUploadAvatarByIDParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: petID.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "petID",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.PetID = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "petID",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
