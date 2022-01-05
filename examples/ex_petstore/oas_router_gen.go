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

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	args := map[string]string{}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "GET":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/pets"
			if prefix := "/pets"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
				elem = elem[len(prefix):]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleListPetsRequest(args, w, r)
				return
			}
			switch elem[0] {
			case '/': // Prefix: "/"
				if prefix := "/"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
					elem = elem[len(prefix):]
				} else {
					break
				}

				// Param: "petId"
				// Leaf parameter
				args["petId"] = elem

				// Leaf: ShowPetById
				s.handleShowPetByIdRequest(args, w, r)
				return
			}
		}
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/pets"
			if prefix := "/pets"; len(elem) >= len(prefix) && elem[0:len(prefix)] == prefix {
				elem = elem[len(prefix):]
			} else {
				break
			}

			// Leaf: CreatePets
			s.handleCreatePetsRequest(args, w, r)
			return
		}
	}
	s.notFound(w, r)
}
