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

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ErrorStatusCode struct {
	StatusCode int   `json:"-"`
	Response   Error `json:"-"`
}

func (*ErrorStatusCode) foobarPostResponse() {}

type FoobarPutDefault struct {
	StatusCode int `json:"-"`
}

type NotFound struct{}

func (*NotFound) foobarGetResponse()  {}
func (*NotFound) foobarPostResponse() {}

type Pet struct {
	Birthday time.Time  `json:"birthday"`
	Friends  *[]Pet     `json:"friends"`
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	Tag      *uuid.UUID `json:"tag"`
}

func (*Pet) foobarGetResponse()  {}
func (*Pet) foobarPostResponse() {}
func (*Pet) petCreateRequest()   {}
func (*Pet) petGetResponse()     {}

type PetCreateTextPlainRequest string

func (*PetCreateTextPlainRequest) petCreateRequest() {}

type PetGetDefault struct {
	Message string `json:"message"`
}

type PetGetDefaultStatusCode struct {
	StatusCode int           `json:"-"`
	Response   PetGetDefault `json:"-"`
}

func (*PetGetDefaultStatusCode) petGetResponse() {}
