// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/encoding/json"
	"github.com/ogen-go/ogen/uri"
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
)

func decodeFoobarGetResponse(resp *http.Response) (_ FoobarGetResponse, rerr error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Pet
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				rerr = err
				return
			}
			if err := json.Unmarshal(data, &response); err != nil {
				rerr = err
				return
			}

			return &response, nil
		default:
			rerr = fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
			return
		}
	case 404:
		return &NotFound{}, nil
	default:
		rerr = fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
		return
	}
}

func decodeFoobarPutResponse(resp *http.Response) (_ FoobarPutDefault, rerr error) {
	switch resp.StatusCode {
	default:
		return FoobarPutDefault{StatusCode: resp.StatusCode}, nil
	}
}

func decodeFoobarPostResponse(resp *http.Response) (_ FoobarPostResponse, rerr error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Pet
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				rerr = err
				return
			}
			if err := json.Unmarshal(data, &response); err != nil {
				rerr = err
				return
			}

			return &response, nil
		default:
			rerr = fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
			return
		}
	case 404:
		return &NotFound{}, nil
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response ErrorStatusCode
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				rerr = err
				return
			}
			if err := json.Unmarshal(data, &response.Response); err != nil {
				rerr = err
				return
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			rerr = fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
			return
		}
	}
}

func decodePetGetResponse(resp *http.Response) (_ PetGetResponse, rerr error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Pet
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				rerr = err
				return
			}
			if err := json.Unmarshal(data, &response); err != nil {
				rerr = err
				return
			}

			return &response, nil
		default:
			rerr = fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
			return
		}
	default:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response PetGetDefaultStatusCode
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				rerr = err
				return
			}
			if err := json.Unmarshal(data, &response.Response); err != nil {
				rerr = err
				return
			}

			response.StatusCode = resp.StatusCode
			return &response, nil
		default:
			rerr = fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
			return
		}
	}
}

func decodePetCreateResponse(resp *http.Response) (_ Pet, rerr error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Pet
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				rerr = err
				return
			}
			if err := json.Unmarshal(data, &response); err != nil {
				rerr = err
				return
			}

			return response, nil
		default:
			rerr = fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
			return
		}
	default:
		rerr = fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
		return
	}
}

func decodePetGetByNameResponse(resp *http.Response) (_ Pet, rerr error) {
	switch resp.StatusCode {
	case 200:
		switch resp.Header.Get("Content-Type") {
		case "application/json":
			var response Pet
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				rerr = err
				return
			}
			if err := json.Unmarshal(data, &response); err != nil {
				rerr = err
				return
			}

			return response, nil
		default:
			rerr = fmt.Errorf("unexpected content-type: %s", resp.Header.Get("Content-Type"))
			return
		}
	default:
		rerr = fmt.Errorf("unexpected statusCode: %d", resp.StatusCode)
		return
	}
}
