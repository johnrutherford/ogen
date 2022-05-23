// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"

	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/validate"
)

func decodeGetBookResponse(resp *http.Response, span trace.Span) (res GetBookRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Book
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		return &GetBookForbidden{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeSearchResponse(resp *http.Response, span trace.Span) (res SearchRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response SearchOKApplicationJSON
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		return &SearchForbidden{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
func decodeSearchByTagIDResponse(resp *http.Response, span trace.Span) (res SearchByTagIDRes, err error) {
	switch resp.StatusCode {
	case 200:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response SearchByTagIDOKApplicationJSON
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 403:
		return &SearchByTagIDForbidden{}, nil
	default:
		return res, validate.UnexpectedStatusCode(resp.StatusCode)
	}
}
