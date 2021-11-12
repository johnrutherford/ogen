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

// HandleDataCreateRequest handles dataCreate operation.
//
// POST /data
func (s *Server) handleDataCreateRequest(args map[string]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `DataCreate`,
		trace.WithAttributes(otelogen.OperationID(`dataCreate`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeDataCreateRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DataCreate(ctx, request)
	if err != nil {
		span.RecordError(err)
		var errRes *ErrorStatusCode
		if errors.As(err, &errRes) {
			encodeErrorResponse(*errRes, w, span)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeDataCreateResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleDataGetRequest handles dataGet operation.
//
// GET /data
func (s *Server) handleDataGetRequest(args map[string]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), `DataGet`,
		trace.WithAttributes(otelogen.OperationID(`dataGet`)),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	response, err := s.h.DataGet(ctx)
	if err != nil {
		span.RecordError(err)
		var errRes *ErrorStatusCode
		if errors.As(err, &errRes) {
			encodeErrorResponse(*errRes, w, span)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeDataGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, writeErr := json.Marshal(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
	if writeErr == nil {
		w.Write(data)
	}
}