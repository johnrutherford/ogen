// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// ComplicatedParameterNameGetParams is parameters of GET /complicatedParameterName operation.
type ComplicatedParameterNameGetParams struct {
	Eq       string
	Plus     string
	Question string
	And      string
	Percent  string
}

func unpackComplicatedParameterNameGetParams(packed middleware.Parameters) (params ComplicatedParameterNameGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "=",
			In:   "query",
		}
		params.Eq = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "+",
			In:   "query",
		}
		params.Plus = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "question?",
			In:   "query",
		}
		params.Question = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "and&",
			In:   "query",
		}
		params.And = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "percent%",
			In:   "query",
		}
		params.Percent = packed[key].(string)
	}
	return params
}

func decodeComplicatedParameterNameGetParams(args [0]string, r *http.Request) (params ComplicatedParameterNameGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: =.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "=",
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

				params.Eq = c
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
			Name: "=",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: +.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "+",
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

				params.Plus = c
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
			Name: "+",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: question?.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "question?",
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

				params.Question = c
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
			Name: "question?",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: and&.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "and&",
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

				params.And = c
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
			Name: "and&",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: percent%.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "percent%",
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

				params.Percent = c
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
			Name: "percent%",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ContentQueryParameterParams is parameters of contentQueryParameter operation.
type ContentQueryParameterParams struct {
	Param OptContentQueryParameterParam
}

func unpackContentQueryParameterParams(packed middleware.Parameters) (params ContentQueryParameterParams) {
	{
		key := middleware.ParameterKey{
			Name: "param",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Param = v.(OptContentQueryParameterParam)
		}
	}
	return params
}

func decodeContentQueryParameterParams(args [0]string, r *http.Request) (params ContentQueryParameterParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: param.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "param",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}
				if err := func(d *jx.Decoder) error {
					params.Param.Reset()
					if err := params.Param.Decode(d); err != nil {
						return err
					}
					return nil
				}(jx.DecodeStr(val)); err != nil {
					return err
				}
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "param",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// HeaderParameterParams is parameters of headerParameter operation.
type HeaderParameterParams struct {
	XAuthToken string
}

func unpackHeaderParameterParams(packed middleware.Parameters) (params HeaderParameterParams) {
	{
		key := middleware.ParameterKey{
			Name: "x-auth-token",
			In:   "header",
		}
		params.XAuthToken = packed[key].(string)
	}
	return params
}

func decodeHeaderParameterParams(args [0]string, r *http.Request) (params HeaderParameterParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: x-auth-token.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "x-auth-token",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XAuthToken = c
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
			Name: "x-auth-token",
			In:   "header",
			Err:  err,
		}
	}
	return params, nil
}

// ObjectQueryParameterParams is parameters of objectQueryParameter operation.
type ObjectQueryParameterParams struct {
	FormObject OptObjectQueryParameterFormObject
	DeepObject OptObjectQueryParameterDeepObject
}

func unpackObjectQueryParameterParams(packed middleware.Parameters) (params ObjectQueryParameterParams) {
	{
		key := middleware.ParameterKey{
			Name: "formObject",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.FormObject = v.(OptObjectQueryParameterFormObject)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "deepObject",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.DeepObject = v.(OptObjectQueryParameterDeepObject)
		}
	}
	return params
}

func decodeObjectQueryParameterParams(args [0]string, r *http.Request) (params ObjectQueryParameterParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: formObject.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "formObject",
			Style:   uri.QueryStyleForm,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{"min", true}, {"max", true}, {"filter", true}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotFormObjectVal ObjectQueryParameterFormObject
				if err := func() error {
					return paramsDotFormObjectVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.FormObject.SetTo(paramsDotFormObjectVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "formObject",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: deepObject.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "deepObject",
			Style:   uri.QueryStyleDeepObject,
			Explode: true,
			Fields:  []uri.QueryParameterObjectField{{"min", true}, {"max", true}, {"filter", true}},
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotDeepObjectVal ObjectQueryParameterDeepObject
				if err := func() error {
					return paramsDotDeepObjectVal.DecodeURI(d)
				}(); err != nil {
					return err
				}
				params.DeepObject.SetTo(paramsDotDeepObjectVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "deepObject",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// PathObjectParameterParams is parameters of pathObjectParameter operation.
type PathObjectParameterParams struct {
	Param User
}

func unpackPathObjectParameterParams(packed middleware.Parameters) (params PathObjectParameterParams) {
	{
		key := middleware.ParameterKey{
			Name: "param",
			In:   "path",
		}
		params.Param = packed[key].(User)
	}
	return params
}

func decodePathObjectParameterParams(args [1]string, r *http.Request) (params PathObjectParameterParams, _ error) {
	// Decode path: param.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "param",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}
				if err := func(d *jx.Decoder) error {
					if err := params.Param.Decode(d); err != nil {
						return err
					}
					return nil
				}(jx.DecodeStr(val)); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := params.Param.Validate(); err != nil {
					return err
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
			Name: "param",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// SameNameParams is parameters of sameName operation.
type SameNameParams struct {
	pathPath  string
	queryPath string
}

func unpackSameNameParams(packed middleware.Parameters) (params SameNameParams) {
	{
		key := middleware.ParameterKey{
			Name: "path",
			In:   "path",
		}
		params.pathPath = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "path",
			In:   "query",
		}
		params.queryPath = packed[key].(string)
	}
	return params
}

func decodeSameNameParams(args [1]string, r *http.Request) (params SameNameParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: path.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "path",
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

				params.pathPath = c
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
			Name: "path",
			In:   "path",
			Err:  err,
		}
	}
	// Decode query: path.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "path",
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

				params.queryPath = c
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
			Name: "path",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
