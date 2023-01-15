// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleComplicatedParameterNameGetRequest handles GET /complicatedParameterName operation.
//
// GET /complicatedParameterName
func (s *Server) handleComplicatedParameterNameGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	var otelAttrs []attribute.KeyValue

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ComplicatedParameterNameGet",
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "ComplicatedParameterNameGet",
			ID:   "",
		}
	)
	params, err := decodeComplicatedParameterNameGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *ComplicatedParameterNameGetOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "ComplicatedParameterNameGet",
			OperationID:   "",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "=",
					In:   "query",
				}: params.Eq,
				{
					Name: "+",
					In:   "query",
				}: params.Plus,
				{
					Name: "question?",
					In:   "query",
				}: params.Question,
				{
					Name: "and&",
					In:   "query",
				}: params.And,
				{
					Name: "percent%",
					In:   "query",
				}: params.Percent,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ComplicatedParameterNameGetParams
			Response = *ComplicatedParameterNameGetOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackComplicatedParameterNameGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.ComplicatedParameterNameGet(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.ComplicatedParameterNameGet(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeComplicatedParameterNameGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleContentQueryParameterRequest handles contentQueryParameter operation.
//
// GET /contentQueryParameter
func (s *Server) handleContentQueryParameterRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("contentQueryParameter"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ContentQueryParameter",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "ContentQueryParameter",
			ID:   "contentQueryParameter",
		}
	)
	params, err := decodeContentQueryParameterParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response string
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "ContentQueryParameter",
			OperationID:   "contentQueryParameter",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "param",
					In:   "query",
				}: params.Param,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ContentQueryParameterParams
			Response = string
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackContentQueryParameterParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ContentQueryParameter(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ContentQueryParameter(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeContentQueryParameterResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleHeaderParameterRequest handles headerParameter operation.
//
// Test for header param.
//
// GET /headerParameter
func (s *Server) handleHeaderParameterRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("headerParameter"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "HeaderParameter",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "HeaderParameter",
			ID:   "headerParameter",
		}
	)
	params, err := decodeHeaderParameterParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Hash
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "HeaderParameter",
			OperationID:   "headerParameter",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "X-Auth-Token",
					In:   "header",
				}: params.XAuthToken,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = HeaderParameterParams
			Response = *Hash
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackHeaderParameterParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.HeaderParameter(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.HeaderParameter(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeHeaderParameterResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleObjectQueryParameterRequest handles objectQueryParameter operation.
//
// GET /objectQueryParameter
func (s *Server) handleObjectQueryParameterRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("objectQueryParameter"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ObjectQueryParameter",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "ObjectQueryParameter",
			ID:   "objectQueryParameter",
		}
	)
	params, err := decodeObjectQueryParameterParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *ObjectQueryParameterOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "ObjectQueryParameter",
			OperationID:   "objectQueryParameter",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "formObject",
					In:   "query",
				}: params.FormObject,
				{
					Name: "deepObject",
					In:   "query",
				}: params.DeepObject,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = ObjectQueryParameterParams
			Response = *ObjectQueryParameterOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackObjectQueryParameterParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.ObjectQueryParameter(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.ObjectQueryParameter(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeObjectQueryParameterResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handlePathObjectParameterRequest handles pathObjectParameter operation.
//
// GET /pathObjectParameter/{param}
func (s *Server) handlePathObjectParameterRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pathObjectParameter"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PathObjectParameter",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "PathObjectParameter",
			ID:   "pathObjectParameter",
		}
	)
	params, err := decodePathObjectParameterParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *User
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "PathObjectParameter",
			OperationID:   "pathObjectParameter",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "param",
					In:   "path",
				}: params.Param,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = PathObjectParameterParams
			Response = *User
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackPathObjectParameterParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.PathObjectParameter(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.PathObjectParameter(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePathObjectParameterResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// handleSameNameRequest handles sameName operation.
//
// Parameter with different location, but the same name.
//
// GET /same_name/{path}
func (s *Server) handleSameNameRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("sameName"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SameName",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, otelAttrs...)

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, otelAttrs...)
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SameName",
			ID:   "sameName",
		}
	)
	params, err := decodeSameNameParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *SameNameOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:       ctx,
			OperationName: "SameName",
			OperationID:   "sameName",
			Body:          nil,
			Params: middleware.Parameters{
				{
					Name: "path",
					In:   "path",
				}: params.pathPath,
				{
					Name: "path",
					In:   "query",
				}: params.queryPath,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = SameNameParams
			Response = *SameNameOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackSameNameParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.SameName(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.SameName(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeSameNameResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
