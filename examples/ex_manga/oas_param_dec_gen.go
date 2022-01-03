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

func decodeGetBookParams(args map[string]string, r *http.Request) (GetBookParams, error) {
	var (
		params GetBookParams
	)
	// Decode path: book_id.
	{
		param := args["book_id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "book_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.BookID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: book_id: not specified`)
		}
	}
	return params, nil
}

func decodeGetPageCoverImageParams(args map[string]string, r *http.Request) (GetPageCoverImageParams, error) {
	var (
		params GetPageCoverImageParams
	)
	// Decode path: media_id.
	{
		param := args["media_id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "media_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.MediaID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: media_id: not specified`)
		}
	}
	// Decode path: format.
	{
		param := args["format"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "format",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Format = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: format: not specified`)
		}
	}
	return params, nil
}

func decodeGetPageImageParams(args map[string]string, r *http.Request) (GetPageImageParams, error) {
	var (
		params GetPageImageParams
	)
	// Decode path: media_id.
	{
		param := args["media_id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "media_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.MediaID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: media_id: not specified`)
		}
	}
	// Decode path: page.
	{
		param := args["page"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "page",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.Page = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: page: not specified`)
		}
	}
	// Decode path: format.
	{
		param := args["format"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "format",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Format = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: format: not specified`)
		}
	}
	return params, nil
}

func decodeGetPageThumbnailImageParams(args map[string]string, r *http.Request) (GetPageThumbnailImageParams, error) {
	var (
		params GetPageThumbnailImageParams
	)
	// Decode path: media_id.
	{
		param := args["media_id"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "media_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.MediaID = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: media_id: not specified`)
		}
	}
	// Decode path: page.
	{
		param := args["page"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "page",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.Page = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: page: not specified`)
		}
	}
	// Decode path: format.
	{
		param := args["format"]
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "format",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Format = c
				return nil
			}(); err != nil {
				return params, err
			}
		} else {
			return params, errors.New(`path: format: not specified`)
		}
	}
	return params, nil
}

func decodeSearchParams(args map[string]string, r *http.Request) (SearchParams, error) {
	var (
		params    SearchParams
		queryArgs = r.URL.Query()
	)
	// Decode query: query.
	{
		values, ok := queryArgs["query"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(s)
				if err != nil {
					return err
				}

				params.Query = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `query: query: parse`)
			}
		} else {
			return params, errors.New(`query: query: not specified`)
		}
	}
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `query: page: parse`)
			}
		}
	}
	return params, nil
}

func decodeSearchByTagIDParams(args map[string]string, r *http.Request) (SearchByTagIDParams, error) {
	var (
		params    SearchByTagIDParams
		queryArgs = r.URL.Query()
	)
	// Decode query: tag_id.
	{
		values, ok := queryArgs["tag_id"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				s, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(s)
				if err != nil {
					return err
				}

				params.TagID = c
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `query: tag_id: parse`)
			}
		} else {
			return params, errors.New(`query: tag_id: not specified`)
		}
	}
	// Decode query: page.
	{
		values, ok := queryArgs["page"]
		if ok {
			d := uri.NewQueryDecoder(uri.QueryDecoderConfig{
				Values:  values,
				Style:   uri.QueryStyleForm,
				Explode: true,
			})

			if err := func() error {
				var paramsPageVal int
				if err := func() error {
					s, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(s)
					if err != nil {
						return err
					}

					paramsPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsPageVal)
				return nil
			}(); err != nil {
				return params, errors.Wrap(err, `query: page: parse`)
			}
		}
	}
	return params, nil
}
