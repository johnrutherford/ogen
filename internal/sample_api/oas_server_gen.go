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

	"github.com/go-chi/chi/v5"
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

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// ErrorGet implements errorGet operation.
	ErrorGet(ctx context.Context) (ErrorStatusCode, error)
	// FoobarGet implements foobarGet operation.
	FoobarGet(ctx context.Context, params FoobarGetParams) (FoobarGetRes, error)
	// FoobarPost implements foobarPost operation.
	FoobarPost(ctx context.Context, req Pet) (FoobarPostRes, error)
	// FoobarPut implements  operation.
	FoobarPut(ctx context.Context) (FoobarPutDefStatusCode, error)
	// PetCreate implements petCreate operation.
	PetCreate(ctx context.Context, req Pet) (Pet, error)
	// PetFriendsNamesByID implements petFriendsNamesByID operation.
	PetFriendsNamesByID(ctx context.Context, params PetFriendsNamesByIDParams) ([]string, error)
	// PetGet implements petGet operation.
	PetGet(ctx context.Context, params PetGetParams) (PetGetRes, error)
	// PetGetAvatarByID implements petGetAvatarByID operation.
	PetGetAvatarByID(ctx context.Context, params PetGetAvatarByIDParams) (PetGetAvatarByIDRes, error)
	// PetGetByName implements petGetByName operation.
	PetGetByName(ctx context.Context, params PetGetByNameParams) (Pet, error)
	// PetNameByID implements petNameByID operation.
	PetNameByID(ctx context.Context, params PetNameByIDParams) (string, error)
	// PetUpdateNameAliasPost implements  operation.
	PetUpdateNameAliasPost(ctx context.Context, req PetName) (PetUpdateNameAliasPostDefStatusCode, error)
	// PetUpdateNamePost implements  operation.
	PetUpdateNamePost(ctx context.Context, req string) (PetUpdateNamePostDefStatusCode, error)
	// PetUploadAvatarByID implements petUploadAvatarByID operation.
	PetUploadAvatarByID(ctx context.Context, req Stream, params PetUploadAvatarByIDParams) (PetUploadAvatarByIDRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	mux *chi.Mux
	cfg config
}

func NewServer(h Handler, opts ...Option) *Server {
	srv := &Server{
		h:   h,
		mux: chi.NewMux(),
		cfg: newConfig(opts...),
	}
	srv.setupRoutes()
	return srv
}

func (s *Server) setupRoutes() {
	s.mux.MethodFunc("GET", "/error", s.HandleErrorGetRequest)
	s.mux.MethodFunc("GET", "/foobar", s.HandleFoobarGetRequest)
	s.mux.MethodFunc("POST", "/foobar", s.HandleFoobarPostRequest)
	s.mux.MethodFunc("PUT", "/foobar", s.HandleFoobarPutRequest)
	s.mux.MethodFunc("POST", "/pet", s.HandlePetCreateRequest)
	s.mux.MethodFunc("GET", "/pet/friendNames/{id}", s.HandlePetFriendsNamesByIDRequest)
	s.mux.MethodFunc("GET", "/pet", s.HandlePetGetRequest)
	s.mux.MethodFunc("GET", "/pet/avatar", s.HandlePetGetAvatarByIDRequest)
	s.mux.MethodFunc("GET", "/pet/{name}", s.HandlePetGetByNameRequest)
	s.mux.MethodFunc("GET", "/pet/name/{id}", s.HandlePetNameByIDRequest)
	s.mux.MethodFunc("POST", "/pet/updateNameAlias", s.HandlePetUpdateNameAliasPostRequest)
	s.mux.MethodFunc("POST", "/pet/updateName", s.HandlePetUpdateNamePostRequest)
	s.mux.MethodFunc("POST", "/pet/avatar", s.HandlePetUploadAvatarByIDRequest)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
