// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ogen-go/ogen/conv"
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
)

type Server interface {
	FoobarGet(ctx context.Context, params FoobarGetParams) (FoobarGetResponder, error)
	FoobarPost(ctx context.Context, req *Pet) (FoobarPostResponder, error)
	PetGet(ctx context.Context, params PetGetParams) (PetGetResponder, error)
	PetCreate(ctx context.Context, req PetCreateRequester) (*Pet, error)
	PetGetByName(ctx context.Context, params PetGetByNameParams) (*Pet, error)
}
