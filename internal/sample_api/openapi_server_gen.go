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

type Server interface {
	FoobarGet(ctx context.Context, params FoobarGetParams) (FoobarGetResponse, error)
	FoobarPut(ctx context.Context) (FoobarPutDefault, error)
	FoobarPost(ctx context.Context, req *Pet) (FoobarPostResponse, error)
	PetGet(ctx context.Context, params PetGetParams) (PetGetResponse, error)
	PetCreate(ctx context.Context, req PetCreateRequest) (Pet, error)
	PetGetByName(ctx context.Context, params PetGetByNameParams) (Pet, error)
}
