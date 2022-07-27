// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/ogenerrors"
)

// HandleAPICaptcha2chcaptchaIDGetRequest handles GET /api/captcha/2chcaptcha/id operation.
//
// GET /api/captcha/2chcaptcha/id
func (s *Server) handleAPICaptcha2chcaptchaIDGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APICaptcha2chcaptchaIDGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "APICaptcha2chcaptchaIDGet",
			ID:   "",
		}
	)
	params, err := decodeAPICaptcha2chcaptchaIDGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	response, err := s.h.APICaptcha2chcaptchaIDGet(ctx, params)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPICaptcha2chcaptchaIDGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPICaptcha2chcaptchaShowGetRequest handles GET /api/captcha/2chcaptcha/show operation.
//
// GET /api/captcha/2chcaptcha/show
func (s *Server) handleAPICaptcha2chcaptchaShowGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APICaptcha2chcaptchaShowGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "APICaptcha2chcaptchaShowGet",
			ID:   "",
		}
	)
	params, err := decodeAPICaptcha2chcaptchaShowGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	response, err := s.h.APICaptcha2chcaptchaShowGet(ctx, params)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPICaptcha2chcaptchaShowGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPICaptchaAppIDPublicKeyGetRequest handles GET /api/captcha/app/id/{public_key} operation.
//
// GET /api/captcha/app/id/{public_key}
func (s *Server) handleAPICaptchaAppIDPublicKeyGetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APICaptchaAppIDPublicKeyGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "APICaptchaAppIDPublicKeyGet",
			ID:   "",
		}
	)
	params, err := decodeAPICaptchaAppIDPublicKeyGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	response, err := s.h.APICaptchaAppIDPublicKeyGet(ctx, params)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPICaptchaAppIDPublicKeyGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPICaptchaInvisibleRecaptchaIDGetRequest handles GET /api/captcha/invisible_recaptcha/id operation.
//
// GET /api/captcha/invisible_recaptcha/id
func (s *Server) handleAPICaptchaInvisibleRecaptchaIDGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APICaptchaInvisibleRecaptchaIDGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "APICaptchaInvisibleRecaptchaIDGet",
			ID:   "",
		}
	)
	params, err := decodeAPICaptchaInvisibleRecaptchaIDGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	response, err := s.h.APICaptchaInvisibleRecaptchaIDGet(ctx, params)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPICaptchaInvisibleRecaptchaIDGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPICaptchaInvisibleRecaptchaMobileGetRequest handles GET /api/captcha/invisible_recaptcha/mobile operation.
//
// GET /api/captcha/invisible_recaptcha/mobile
func (s *Server) handleAPICaptchaInvisibleRecaptchaMobileGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APICaptchaInvisibleRecaptchaMobileGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
		err error
	)

	response, err := s.h.APICaptchaInvisibleRecaptchaMobileGet(ctx)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPICaptchaInvisibleRecaptchaMobileGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPICaptchaRecaptchaIDGetRequest handles GET /api/captcha/recaptcha/id operation.
//
// GET /api/captcha/recaptcha/id
func (s *Server) handleAPICaptchaRecaptchaIDGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APICaptchaRecaptchaIDGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "APICaptchaRecaptchaIDGet",
			ID:   "",
		}
	)
	params, err := decodeAPICaptchaRecaptchaIDGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	response, err := s.h.APICaptchaRecaptchaIDGet(ctx, params)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPICaptchaRecaptchaIDGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPICaptchaRecaptchaMobileGetRequest handles GET /api/captcha/recaptcha/mobile operation.
//
// GET /api/captcha/recaptcha/mobile
func (s *Server) handleAPICaptchaRecaptchaMobileGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APICaptchaRecaptchaMobileGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
		err error
	)

	response, err := s.h.APICaptchaRecaptchaMobileGet(ctx)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPICaptchaRecaptchaMobileGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPIDislikeGetRequest handles GET /api/dislike operation.
//
// GET /api/dislike
func (s *Server) handleAPIDislikeGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIDislikeGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "APIDislikeGet",
			ID:   "",
		}
	)
	params, err := decodeAPIDislikeGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	response, err := s.h.APIDislikeGet(ctx, params)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIDislikeGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPILikeGetRequest handles GET /api/like operation.
//
// GET /api/like
func (s *Server) handleAPILikeGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APILikeGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "APILikeGet",
			ID:   "",
		}
	)
	params, err := decodeAPILikeGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	response, err := s.h.APILikeGet(ctx, params)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPILikeGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPIMobileV2AfterBoardThreadNumGetRequest handles GET /api/mobile/v2/after/{board}/{thread}/{num} operation.
//
// GET /api/mobile/v2/after/{board}/{thread}/{num}
func (s *Server) handleAPIMobileV2AfterBoardThreadNumGetRequest(args [3]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIMobileV2AfterBoardThreadNumGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "APIMobileV2AfterBoardThreadNumGet",
			ID:   "",
		}
	)
	params, err := decodeAPIMobileV2AfterBoardThreadNumGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	response, err := s.h.APIMobileV2AfterBoardThreadNumGet(ctx, params)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIMobileV2AfterBoardThreadNumGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPIMobileV2BoardsGetRequest handles GET /api/mobile/v2/boards operation.
//
// GET /api/mobile/v2/boards
func (s *Server) handleAPIMobileV2BoardsGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIMobileV2BoardsGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
		err error
	)

	response, err := s.h.APIMobileV2BoardsGet(ctx)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIMobileV2BoardsGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPIMobileV2InfoBoardThreadGetRequest handles GET /api/mobile/v2/info/{board}/{thread} operation.
//
// GET /api/mobile/v2/info/{board}/{thread}
func (s *Server) handleAPIMobileV2InfoBoardThreadGetRequest(args [2]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIMobileV2InfoBoardThreadGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "APIMobileV2InfoBoardThreadGet",
			ID:   "",
		}
	)
	params, err := decodeAPIMobileV2InfoBoardThreadGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	response, err := s.h.APIMobileV2InfoBoardThreadGet(ctx, params)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIMobileV2InfoBoardThreadGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleAPIMobileV2PostBoardNumGetRequest handles GET /api/mobile/v2/post/{board}/{num} operation.
//
// GET /api/mobile/v2/post/{board}/{num}
func (s *Server) handleAPIMobileV2PostBoardNumGetRequest(args [2]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "APIMobileV2PostBoardNumGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "APIMobileV2PostBoardNumGet",
			ID:   "",
		}
	)
	params, err := decodeAPIMobileV2PostBoardNumGetParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	response, err := s.h.APIMobileV2PostBoardNumGet(ctx, params)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeAPIMobileV2PostBoardNumGetResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleUserPassloginPostRequest handles POST /user/passlogin operation.
//
// POST /user/passlogin
func (s *Server) handleUserPassloginPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UserPassloginPost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "UserPassloginPost",
			ID:   "",
		}
	)
	params, err := decodeUserPassloginPostParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeUserPassloginPostRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer close()

	response, err := s.h.UserPassloginPost(ctx, request, params)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUserPassloginPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleUserPostingPostRequest handles POST /user/posting operation.
//
// POST /user/posting
func (s *Server) handleUserPostingPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UserPostingPost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "UserPostingPost",
			ID:   "",
		}
	)
	request, close, err := s.decodeUserPostingPostRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer close()

	response, err := s.h.UserPostingPost(ctx, request)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUserPostingPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}

// HandleUserReportPostRequest handles POST /user/report operation.
//
// POST /user/report
func (s *Server) handleUserReportPostRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UserReportPost",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
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
			Name: "UserReportPost",
			ID:   "",
		}
	)
	request, close, err := s.decodeUserReportPostRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer close()

	response, err := s.h.UserReportPost(ctx, request)
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUserReportPostResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
}
