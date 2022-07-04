// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"
	"net/url"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func (s *Server) decodeTestFormURLEncodedRequest(r *http.Request, span trace.Span) (
	req TestForm,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/x-www-form-urlencoded":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseForm(); err != nil {
			return req, close, errors.Wrap(err, "parse form")
		}
		form := r.PostForm

		var request TestForm
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIDVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						requestDotIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ID.SetTo(requestDotIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal uuid.UUID
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "description",
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

					request.Description = c
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"description\"")
				}
			} else {
				return req, close, errors.Wrap(err, "query")
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "array",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotArrayVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							requestDotArrayVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.Array = append(request.Array, requestDotArrayVal)
						return nil
					})
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"array\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "object",
				Style:   uri.QueryStyleForm,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotObjectVal TestFormObject
					if err := func() error {
						return requestDotObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					request.Object.SetTo(requestDotObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"object\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "deepObject",
				Style:   uri.QueryStyleDeepObject,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotDeepObjectVal TestFormDeepObject
					if err := func() error {
						return requestDotDeepObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					request.DeepObject.SetTo(requestDotDeepObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"deepObject\"")
				}
			}
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestMultipartRequest(r *http.Request, span trace.Span) (
	req TestForm,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		form := url.Values(r.MultipartForm.Value)

		var request TestForm
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "id",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotIDVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						requestDotIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.ID.SetTo(requestDotIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"id\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "uuid",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUUIDVal uuid.UUID
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToUUID(val)
						if err != nil {
							return err
						}

						requestDotUUIDVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UUID.SetTo(requestDotUUIDVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"uuid\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "description",
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

					request.Description = c
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"description\"")
				}
			} else {
				return req, close, errors.Wrap(err, "query")
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "array",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					return d.DecodeArray(func(d uri.Decoder) error {
						var requestDotArrayVal string
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToString(val)
							if err != nil {
								return err
							}

							requestDotArrayVal = c
							return nil
						}(); err != nil {
							return err
						}
						request.Array = append(request.Array, requestDotArrayVal)
						return nil
					})
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"array\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "object",
				Style:   uri.QueryStyleForm,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotObjectVal TestFormObject
					if err := func() error {
						return requestDotObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					request.Object.SetTo(requestDotObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"object\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "deepObject",
				Style:   uri.QueryStyleDeepObject,
				Explode: true,
				Fields:  []uri.QueryParameterObjectField{{"min", false}, {"max", true}},
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotDeepObjectVal TestFormDeepObject
					if err := func() error {
						return requestDotDeepObjectVal.DecodeURI(d)
					}(); err != nil {
						return err
					}
					request.DeepObject.SetTo(requestDotDeepObjectVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"deepObject\"")
				}
			}
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestMultipartUploadRequest(r *http.Request, span trace.Span) (
	req TestMultipartUploadReqForm,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		form := url.Values(r.MultipartForm.Value)

		var request TestMultipartUploadReqForm
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "orderId",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotOrderIdVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						requestDotOrderIdVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.OrderId.SetTo(requestDotOrderIdVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"orderId\"")
				}
			}
		}
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "userId",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotUserIdVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						requestDotUserIdVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.UserId.SetTo(requestDotUserIdVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"userId\"")
				}
			}
		}
		{
			if err := func() error {
				files, ok := r.MultipartForm.File["file"]
				if !ok || len(files) < 1 {
					return validate.ErrFieldRequired
				}
				fh := files[0]

				f, err := fh.Open()
				if err != nil {
					return errors.Wrap(err, "open")
				}
				closers = append(closers, f)
				request.File = ht.MultipartFile{
					Name:   fh.Filename,
					File:   f,
					Header: fh.Header,
				}
				return nil
			}(); err != nil {
				return req, close, errors.Wrap(err, "decode \"file\"")
			}
		}
		{
			if err := func() error {
				files, ok := r.MultipartForm.File["optional_file"]
				if !ok || len(files) < 1 {
					return nil
				}
				fh := files[0]

				f, err := fh.Open()
				if err != nil {
					return errors.Wrap(err, "open")
				}
				closers = append(closers, f)
				request.OptionalFile.SetTo(ht.MultipartFile{
					Name:   fh.Filename,
					File:   f,
					Header: fh.Header,
				})
				return nil
			}(); err != nil {
				return req, close, errors.Wrap(err, "decode \"optional_file\"")
			}
		}
		{
			if err := func() error {
				files, ok := r.MultipartForm.File["files"]
				_ = ok
				request.Files = make([]ht.MultipartFile, 0, len(files))
				for _, fh := range files {
					f, err := fh.Open()
					if err != nil {
						return errors.Wrap(err, "open")
					}
					closers = append(closers, f)

					request.Files = append(request.Files, ht.MultipartFile{
						Name:   fh.Filename,
						File:   f,
						Header: fh.Header,
					})
				}
				if err := func() error {
					if err := (validate.Array{
						MinLength:    0,
						MinLengthSet: false,
						MaxLength:    5,
						MaxLengthSet: true,
					}).ValidateLength(len(request.Files)); err != nil {
						return errors.Wrap(err, "array")
					}
					return nil
				}(); err != nil {
					return errors.Wrap(err, "validate")
				}
				return nil
			}(); err != nil {
				return req, close, errors.Wrap(err, "decode \"files\"")
			}
		}
		return request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}

func (s *Server) decodeTestShareFormSchemaRequest(r *http.Request, span trace.Span) (
	req TestShareFormSchemaReq,
	close func() error,
	rerr error,
) {
	var closers []io.Closer
	close = func() error {
		var merr error
		for _, c := range closers {
			merr = multierr.Append(merr, c.Close())
		}
		return merr
	}
	defer func() {
		if rerr != nil {
			rerr = multierr.Append(rerr, close())
		}
	}()

	ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return req, close, errors.Wrap(err, "parse media type")
	}
	switch ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}

		var request SharedRequest
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return req, close, err
		}

		if len(buf) == 0 {
			return req, close, validate.ErrBodyRequired
		}

		d := jx.DecodeBytes(buf)
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, close, errors.Wrap(err, "decode \"application/json\"")
		}
		return &request, close, nil
	case "multipart/form-data":
		if r.ContentLength == 0 {
			return req, close, validate.ErrBodyRequired
		}
		if err := r.ParseMultipartForm(s.cfg.MaxMultipartMemory); err != nil {
			return req, close, errors.Wrap(err, "parse multipart form")
		}
		form := url.Values(r.MultipartForm.Value)

		var request SharedRequestForm
		q := uri.NewQueryDecoder(form)
		{
			cfg := uri.QueryParameterDecodingConfig{
				Name:    "filename",
				Style:   uri.QueryStyleForm,
				Explode: true,
			}
			if err := q.HasParam(cfg); err == nil {
				if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
					var requestDotFilenameVal string
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToString(val)
						if err != nil {
							return err
						}

						requestDotFilenameVal = c
						return nil
					}(); err != nil {
						return err
					}
					request.Filename.SetTo(requestDotFilenameVal)
					return nil
				}); err != nil {
					return req, close, errors.Wrap(err, "decode \"filename\"")
				}
			}
		}
		{
			if err := func() error {
				files, ok := r.MultipartForm.File["file"]
				if !ok || len(files) < 1 {
					return nil
				}
				fh := files[0]

				f, err := fh.Open()
				if err != nil {
					return errors.Wrap(err, "open")
				}
				closers = append(closers, f)
				request.File.SetTo(ht.MultipartFile{
					Name:   fh.Filename,
					File:   f,
					Header: fh.Header,
				})
				return nil
			}(); err != nil {
				return req, close, errors.Wrap(err, "decode \"file\"")
			}
		}
		return &request, close, nil
	default:
		return req, close, validate.InvalidContentType(ct)
	}
}
