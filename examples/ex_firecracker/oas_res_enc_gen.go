// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
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
)

func encodeDescribeInstanceResponse(response DescribeInstanceRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *InstanceInfo:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/: unexpected response type: %T", response)
	}
}

func encodeCreateSyncActionResponse(response CreateSyncActionRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *CreateSyncActionNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/actions: unexpected response type: %T", response)
	}
}

func encodeDescribeBalloonConfigResponse(response DescribeBalloonConfigRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *Balloon:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/balloon: unexpected response type: %T", response)
	}
}

func encodePutBalloonResponse(response PutBalloonRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PutBalloonNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/balloon: unexpected response type: %T", response)
	}
}

func encodePatchBalloonResponse(response PatchBalloonRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PatchBalloonNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/balloon: unexpected response type: %T", response)
	}
}

func encodeDescribeBalloonStatsResponse(response DescribeBalloonStatsRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *BalloonStats:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/balloon/statistics: unexpected response type: %T", response)
	}
}

func encodePatchBalloonStatsIntervalResponse(response PatchBalloonStatsIntervalRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PatchBalloonStatsIntervalNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/balloon/statistics: unexpected response type: %T", response)
	}
}

func encodePutGuestBootSourceResponse(response PutGuestBootSourceRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PutGuestBootSourceNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/boot-source: unexpected response type: %T", response)
	}
}

func encodePutGuestDriveByIDResponse(response PutGuestDriveByIDRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PutGuestDriveByIDNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/drives/{drive_id}: unexpected response type: %T", response)
	}
}

func encodePatchGuestDriveByIDResponse(response PatchGuestDriveByIDRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PatchGuestDriveByIDNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/drives/{drive_id}: unexpected response type: %T", response)
	}
}

func encodePutLoggerResponse(response PutLoggerRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PutLoggerNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/logger: unexpected response type: %T", response)
	}
}

func encodeGetMachineConfigurationResponse(response GetMachineConfigurationRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *MachineConfiguration:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/machine-config: unexpected response type: %T", response)
	}
}

func encodePutMachineConfigurationResponse(response PutMachineConfigurationRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PutMachineConfigurationNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/machine-config: unexpected response type: %T", response)
	}
}

func encodePatchMachineConfigurationResponse(response PatchMachineConfigurationRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PatchMachineConfigurationNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/machine-config: unexpected response type: %T", response)
	}
}

func encodePutMetricsResponse(response PutMetricsRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PutMetricsNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/metrics: unexpected response type: %T", response)
	}
}

func encodeMmdsGetResponse(response MmdsGetRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *MmdsGetOK:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/mmds: unexpected response type: %T", response)
	}
}

func encodeMmdsPutResponse(response MmdsPutRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *MmdsPutNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/mmds: unexpected response type: %T", response)
	}
}

func encodeMmdsPatchResponse(response MmdsPatchRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *MmdsPatchNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/mmds: unexpected response type: %T", response)
	}
}

func encodeMmdsConfigPutResponse(response MmdsConfigPutRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *MmdsConfigPutNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/mmds/config: unexpected response type: %T", response)
	}
}

func encodePutGuestNetworkInterfaceByIDResponse(response PutGuestNetworkInterfaceByIDRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PutGuestNetworkInterfaceByIDNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/network-interfaces/{iface_id}: unexpected response type: %T", response)
	}
}

func encodePatchGuestNetworkInterfaceByIDResponse(response PatchGuestNetworkInterfaceByIDRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PatchGuestNetworkInterfaceByIDNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/network-interfaces/{iface_id}: unexpected response type: %T", response)
	}
}

func encodeCreateSnapshotResponse(response CreateSnapshotRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *CreateSnapshotNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/snapshot/create: unexpected response type: %T", response)
	}
}

func encodeLoadSnapshotResponse(response LoadSnapshotRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *LoadSnapshotNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/snapshot/load: unexpected response type: %T", response)
	}
}

func encodePatchVmResponse(response PatchVmRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PatchVmNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/vm: unexpected response type: %T", response)
	}
}

func encodeGetExportVmConfigResponse(response GetExportVmConfigRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *FullVmConfiguration:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/vm/config: unexpected response type: %T", response)
	}
}

func encodePutGuestVsockResponse(response PutGuestVsockRes, w http.ResponseWriter) error {
	switch response := response.(type) {
	case *PutGuestVsockNoContent:
		w.WriteHeader(204)
		return nil
	case *Error:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		j := json.NewStream(w)
		more := json.NewMore(j)
		defer more.Reset()
		more.More()
		response.Response.WriteJSON(j)
		if err := j.Flush(); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("/vsock: unexpected response type: %T", response)
	}
}
