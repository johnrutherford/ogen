// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// HandleDataCreateRequest handles dataCreate operation.
//
// POST /data
func (s *Server) handleDataCreateRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("dataCreate"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DataCreate",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodeDataCreateRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"DataCreate",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.DataCreate(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		var errRes *ErrorStatusCode
		if errors.As(err, &errRes) {
			encodeErrorResponse(*errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeDataCreateResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleDataGetRequest handles dataGet operation.
//
// GET /data
func (s *Server) handleDataGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("dataGet"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DataGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.DataGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		var errRes *ErrorStatusCode
		if errors.As(err, &errRes) {
			encodeErrorResponse(*errRes, w, span)
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		encodeErrorResponse(s.h.NewError(ctx, err), w, span)
		return
	}

	if err := encodeDataGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

func (s *Server) badRequest(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
	span trace.Span,
	otelAttrs []attribute.KeyValue,
	err error,
) {
	span.RecordError(err)
	span.SetStatus(codes.Error, "BadRequest")
	s.errors.Add(ctx, 1, otelAttrs...)
	s.cfg.ErrorHandler(ctx, w, r, err)
}
