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

// Server handles operations described by OpenAPI v3 specification.
type Server interface {
	// CreateSnapshot implements createSnapshot operation.
	CreateSnapshot(ctx context.Context, req SnapshotCreateParams) (CreateSnapshotRes, error)
	// CreateSyncAction implements createSyncAction operation.
	CreateSyncAction(ctx context.Context, req InstanceActionInfo) (CreateSyncActionRes, error)
	// DescribeBalloonConfig implements describeBalloonConfig operation.
	DescribeBalloonConfig(ctx context.Context) (DescribeBalloonConfigRes, error)
	// DescribeBalloonStats implements describeBalloonStats operation.
	DescribeBalloonStats(ctx context.Context) (DescribeBalloonStatsRes, error)
	// DescribeInstance implements describeInstance operation.
	DescribeInstance(ctx context.Context) (DescribeInstanceRes, error)
	// GetExportVmConfig implements getExportVmConfig operation.
	GetExportVmConfig(ctx context.Context) (GetExportVmConfigRes, error)
	// GetMachineConfiguration implements getMachineConfiguration operation.
	GetMachineConfiguration(ctx context.Context) (GetMachineConfigurationRes, error)
	// LoadSnapshot implements loadSnapshot operation.
	LoadSnapshot(ctx context.Context, req SnapshotLoadParams) (LoadSnapshotRes, error)
	// MmdsConfigPut implements  operation.
	MmdsConfigPut(ctx context.Context, req MmdsConfig) (MmdsConfigPutRes, error)
	// MmdsGet implements  operation.
	MmdsGet(ctx context.Context) (MmdsGetRes, error)
	// MmdsPatch implements  operation.
	MmdsPatch(ctx context.Context, req MmdsPatchReq) (MmdsPatchRes, error)
	// MmdsPut implements  operation.
	MmdsPut(ctx context.Context, req MmdsPutReq) (MmdsPutRes, error)
	// PatchBalloon implements patchBalloon operation.
	PatchBalloon(ctx context.Context, req BalloonUpdate) (PatchBalloonRes, error)
	// PatchBalloonStatsInterval implements patchBalloonStatsInterval operation.
	PatchBalloonStatsInterval(ctx context.Context, req BalloonStatsUpdate) (PatchBalloonStatsIntervalRes, error)
	// PatchGuestDriveByID implements patchGuestDriveByID operation.
	PatchGuestDriveByID(ctx context.Context, req PartialDrive, params PatchGuestDriveByIDParams) (PatchGuestDriveByIDRes, error)
	// PatchGuestNetworkInterfaceByID implements patchGuestNetworkInterfaceByID operation.
	PatchGuestNetworkInterfaceByID(ctx context.Context, req PartialNetworkInterface, params PatchGuestNetworkInterfaceByIDParams) (PatchGuestNetworkInterfaceByIDRes, error)
	// PatchMachineConfiguration implements patchMachineConfiguration operation.
	PatchMachineConfiguration(ctx context.Context, req MachineConfiguration) (PatchMachineConfigurationRes, error)
	// PatchVm implements patchVm operation.
	PatchVm(ctx context.Context, req VM) (PatchVmRes, error)
	// PutBalloon implements putBalloon operation.
	PutBalloon(ctx context.Context, req Balloon) (PutBalloonRes, error)
	// PutGuestBootSource implements putGuestBootSource operation.
	PutGuestBootSource(ctx context.Context, req BootSource) (PutGuestBootSourceRes, error)
	// PutGuestDriveByID implements putGuestDriveByID operation.
	PutGuestDriveByID(ctx context.Context, req Drive, params PutGuestDriveByIDParams) (PutGuestDriveByIDRes, error)
	// PutGuestNetworkInterfaceByID implements putGuestNetworkInterfaceByID operation.
	PutGuestNetworkInterfaceByID(ctx context.Context, req NetworkInterface, params PutGuestNetworkInterfaceByIDParams) (PutGuestNetworkInterfaceByIDRes, error)
	// PutGuestVsock implements putGuestVsock operation.
	PutGuestVsock(ctx context.Context, req Vsock) (PutGuestVsockRes, error)
	// PutLogger implements putLogger operation.
	PutLogger(ctx context.Context, req Logger) (PutLoggerRes, error)
	// PutMachineConfiguration implements putMachineConfiguration operation.
	PutMachineConfiguration(ctx context.Context, req MachineConfiguration) (PutMachineConfigurationRes, error)
	// PutMetrics implements putMetrics operation.
	PutMetrics(ctx context.Context, req Metrics) (PutMetricsRes, error)
}
