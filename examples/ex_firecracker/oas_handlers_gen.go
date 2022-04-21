// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// HandleCreateSnapshotRequest handles createSnapshot operation.
//
// PUT /snapshot/create
func (s *Server) handleCreateSnapshotRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createSnapshot"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateSnapshot",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodeCreateSnapshotRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"CreateSnapshot",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.CreateSnapshot(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateSnapshotResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleCreateSyncActionRequest handles createSyncAction operation.
//
// PUT /actions
func (s *Server) handleCreateSyncActionRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createSyncAction"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateSyncAction",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodeCreateSyncActionRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"CreateSyncAction",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.CreateSyncAction(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateSyncActionResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleDescribeBalloonConfigRequest handles describeBalloonConfig operation.
//
// GET /balloon
func (s *Server) handleDescribeBalloonConfigRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("describeBalloonConfig"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DescribeBalloonConfig",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.DescribeBalloonConfig(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDescribeBalloonConfigResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleDescribeBalloonStatsRequest handles describeBalloonStats operation.
//
// GET /balloon/statistics
func (s *Server) handleDescribeBalloonStatsRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("describeBalloonStats"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DescribeBalloonStats",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.DescribeBalloonStats(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDescribeBalloonStatsResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleDescribeInstanceRequest handles describeInstance operation.
//
// GET /
func (s *Server) handleDescribeInstanceRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("describeInstance"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DescribeInstance",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.DescribeInstance(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDescribeInstanceResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleGetExportVmConfigRequest handles getExportVmConfig operation.
//
// GET /vm/config
func (s *Server) handleGetExportVmConfigRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getExportVmConfig"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetExportVmConfig",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.GetExportVmConfig(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetExportVmConfigResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleGetMachineConfigurationRequest handles getMachineConfiguration operation.
//
// GET /machine-config
func (s *Server) handleGetMachineConfigurationRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getMachineConfiguration"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetMachineConfiguration",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.GetMachineConfiguration(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetMachineConfigurationResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleLoadSnapshotRequest handles loadSnapshot operation.
//
// PUT /snapshot/load
func (s *Server) handleLoadSnapshotRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("loadSnapshot"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "LoadSnapshot",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodeLoadSnapshotRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"LoadSnapshot",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.LoadSnapshot(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeLoadSnapshotResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleMmdsConfigPutRequest handles  operation.
//
// PUT /mmds/config
func (s *Server) handleMmdsConfigPutRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MmdsConfigPut",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodeMmdsConfigPutRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"MmdsConfigPut",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.MmdsConfigPut(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMmdsConfigPutResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleMmdsGetRequest handles  operation.
//
// GET /mmds
func (s *Server) handleMmdsGetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MmdsGet",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error

	response, err := s.h.MmdsGet(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMmdsGetResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleMmdsPatchRequest handles  operation.
//
// PATCH /mmds
func (s *Server) handleMmdsPatchRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MmdsPatch",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodeMmdsPatchRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"MmdsPatch",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.MmdsPatch(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMmdsPatchResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandleMmdsPutRequest handles  operation.
//
// PUT /mmds
func (s *Server) handleMmdsPutRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MmdsPut",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodeMmdsPutRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"MmdsPut",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.MmdsPut(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeMmdsPutResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePatchBalloonRequest handles patchBalloon operation.
//
// PATCH /balloon
func (s *Server) handlePatchBalloonRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("patchBalloon"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchBalloon",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodePatchBalloonRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PatchBalloon",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PatchBalloon(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePatchBalloonResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePatchBalloonStatsIntervalRequest handles patchBalloonStatsInterval operation.
//
// PATCH /balloon/statistics
func (s *Server) handlePatchBalloonStatsIntervalRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("patchBalloonStatsInterval"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchBalloonStatsInterval",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodePatchBalloonStatsIntervalRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PatchBalloonStatsInterval",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PatchBalloonStatsInterval(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePatchBalloonStatsIntervalResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePatchGuestDriveByIDRequest handles patchGuestDriveByID operation.
//
// PATCH /drives/{drive_id}
func (s *Server) handlePatchGuestDriveByIDRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("patchGuestDriveByID"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchGuestDriveByID",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodePatchGuestDriveByIDParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			"PatchGuestDriveByID",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, err := decodePatchGuestDriveByIDRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PatchGuestDriveByID",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PatchGuestDriveByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePatchGuestDriveByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePatchGuestNetworkInterfaceByIDRequest handles patchGuestNetworkInterfaceByID operation.
//
// PATCH /network-interfaces/{iface_id}
func (s *Server) handlePatchGuestNetworkInterfaceByIDRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("patchGuestNetworkInterfaceByID"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchGuestNetworkInterfaceByID",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodePatchGuestNetworkInterfaceByIDParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			"PatchGuestNetworkInterfaceByID",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, err := decodePatchGuestNetworkInterfaceByIDRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PatchGuestNetworkInterfaceByID",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PatchGuestNetworkInterfaceByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePatchGuestNetworkInterfaceByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePatchMachineConfigurationRequest handles patchMachineConfiguration operation.
//
// PATCH /machine-config
func (s *Server) handlePatchMachineConfigurationRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("patchMachineConfiguration"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchMachineConfiguration",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodePatchMachineConfigurationRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PatchMachineConfiguration",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PatchMachineConfiguration(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePatchMachineConfigurationResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePatchVmRequest handles patchVm operation.
//
// PATCH /vm
func (s *Server) handlePatchVmRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("patchVm"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PatchVm",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodePatchVmRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PatchVm",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PatchVm(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePatchVmResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePutBalloonRequest handles putBalloon operation.
//
// PUT /balloon
func (s *Server) handlePutBalloonRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putBalloon"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutBalloon",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodePutBalloonRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PutBalloon",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PutBalloon(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePutBalloonResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePutGuestBootSourceRequest handles putGuestBootSource operation.
//
// PUT /boot-source
func (s *Server) handlePutGuestBootSourceRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putGuestBootSource"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutGuestBootSource",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodePutGuestBootSourceRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PutGuestBootSource",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PutGuestBootSource(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePutGuestBootSourceResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePutGuestDriveByIDRequest handles putGuestDriveByID operation.
//
// PUT /drives/{drive_id}
func (s *Server) handlePutGuestDriveByIDRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putGuestDriveByID"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutGuestDriveByID",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodePutGuestDriveByIDParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			"PutGuestDriveByID",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, err := decodePutGuestDriveByIDRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PutGuestDriveByID",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PutGuestDriveByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePutGuestDriveByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePutGuestNetworkInterfaceByIDRequest handles putGuestNetworkInterfaceByID operation.
//
// PUT /network-interfaces/{iface_id}
func (s *Server) handlePutGuestNetworkInterfaceByIDRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putGuestNetworkInterfaceByID"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutGuestNetworkInterfaceByID",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	params, err := decodePutGuestNetworkInterfaceByIDParams(args, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			"PutGuestNetworkInterfaceByID",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}
	request, err := decodePutGuestNetworkInterfaceByIDRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PutGuestNetworkInterfaceByID",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PutGuestNetworkInterfaceByID(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePutGuestNetworkInterfaceByIDResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePutGuestVsockRequest handles putGuestVsock operation.
//
// PUT /vsock
func (s *Server) handlePutGuestVsockRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putGuestVsock"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutGuestVsock",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodePutGuestVsockRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PutGuestVsock",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PutGuestVsock(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePutGuestVsockResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePutLoggerRequest handles putLogger operation.
//
// PUT /logger
func (s *Server) handlePutLoggerRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putLogger"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutLogger",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodePutLoggerRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PutLogger",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PutLogger(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePutLoggerResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePutMachineConfigurationRequest handles putMachineConfiguration operation.
//
// PUT /machine-config
func (s *Server) handlePutMachineConfigurationRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putMachineConfiguration"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutMachineConfiguration",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodePutMachineConfigurationRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PutMachineConfiguration",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PutMachineConfiguration(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePutMachineConfigurationResponse(response, w, span); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Response")
		s.errors.Add(ctx, 1, otelAttrs...)
		return
	}
	elapsedDuration := time.Since(startTime)
	s.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
}

// HandlePutMetricsRequest handles putMetrics operation.
//
// PUT /metrics
func (s *Server) handlePutMetricsRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("putMetrics"),
	}
	ctx, span := s.cfg.Tracer.Start(r.Context(), "PutMetrics",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	s.requests.Add(ctx, 1, otelAttrs...)
	defer span.End()

	var err error
	request, err := decodePutMetricsRequest(r, span)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			"PutMetrics",
			err,
		}
		s.badRequest(ctx, w, r, span, otelAttrs, err)
		return
	}

	response, err := s.h.PutMetrics(ctx, request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "Internal")
		s.errors.Add(ctx, 1, otelAttrs...)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodePutMetricsResponse(response, w, span); err != nil {
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
