// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/config/v1alpha1/pagerduty.proto

package configv1alpha1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// PagerDutyServiceName is the fully-qualified name of the PagerDutyService service.
	PagerDutyServiceName = "commonfate.control.config.v1alpha1.PagerDutyService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PagerDutyServiceCreateScheduleProcedure is the fully-qualified name of the PagerDutyService's
	// CreateSchedule RPC.
	PagerDutyServiceCreateScheduleProcedure = "/commonfate.control.config.v1alpha1.PagerDutyService/CreateSchedule"
	// PagerDutyServiceReadScheduleProcedure is the fully-qualified name of the PagerDutyService's
	// ReadSchedule RPC.
	PagerDutyServiceReadScheduleProcedure = "/commonfate.control.config.v1alpha1.PagerDutyService/ReadSchedule"
	// PagerDutyServiceUpdateScheduleProcedure is the fully-qualified name of the PagerDutyService's
	// UpdateSchedule RPC.
	PagerDutyServiceUpdateScheduleProcedure = "/commonfate.control.config.v1alpha1.PagerDutyService/UpdateSchedule"
	// PagerDutyServiceDeleteScheduleProcedure is the fully-qualified name of the PagerDutyService's
	// DeleteSchedule RPC.
	PagerDutyServiceDeleteScheduleProcedure = "/commonfate.control.config.v1alpha1.PagerDutyService/DeleteSchedule"
	// PagerDutyServiceReadPagerDutySchedulesProcedure is the fully-qualified name of the
	// PagerDutyService's ReadPagerDutySchedules RPC.
	PagerDutyServiceReadPagerDutySchedulesProcedure = "/commonfate.control.config.v1alpha1.PagerDutyService/ReadPagerDutySchedules"
)

// PagerDutyServiceClient is a client for the commonfate.control.config.v1alpha1.PagerDutyService
// service.
type PagerDutyServiceClient interface {
	CreateSchedule(context.Context, *connect_go.Request[v1alpha1.CreateScheduleRequest]) (*connect_go.Response[v1alpha1.CreateScheduleResponse], error)
	ReadSchedule(context.Context, *connect_go.Request[v1alpha1.ReadScheduleRequest]) (*connect_go.Response[v1alpha1.ReadScheduleResponse], error)
	UpdateSchedule(context.Context, *connect_go.Request[v1alpha1.UpdateScheduleRequest]) (*connect_go.Response[v1alpha1.UpdateScheduleResponse], error)
	DeleteSchedule(context.Context, *connect_go.Request[v1alpha1.DeleteScheduleRequest]) (*connect_go.Response[v1alpha1.DeleteScheduleResponse], error)
	ReadPagerDutySchedules(context.Context, *connect_go.Request[v1alpha1.ReadPagerDutySchedulesRequest]) (*connect_go.Response[v1alpha1.ReadPagerDutySchedulesResponse], error)
}

// NewPagerDutyServiceClient constructs a client for the
// commonfate.control.config.v1alpha1.PagerDutyService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPagerDutyServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PagerDutyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &pagerDutyServiceClient{
		createSchedule: connect_go.NewClient[v1alpha1.CreateScheduleRequest, v1alpha1.CreateScheduleResponse](
			httpClient,
			baseURL+PagerDutyServiceCreateScheduleProcedure,
			opts...,
		),
		readSchedule: connect_go.NewClient[v1alpha1.ReadScheduleRequest, v1alpha1.ReadScheduleResponse](
			httpClient,
			baseURL+PagerDutyServiceReadScheduleProcedure,
			opts...,
		),
		updateSchedule: connect_go.NewClient[v1alpha1.UpdateScheduleRequest, v1alpha1.UpdateScheduleResponse](
			httpClient,
			baseURL+PagerDutyServiceUpdateScheduleProcedure,
			opts...,
		),
		deleteSchedule: connect_go.NewClient[v1alpha1.DeleteScheduleRequest, v1alpha1.DeleteScheduleResponse](
			httpClient,
			baseURL+PagerDutyServiceDeleteScheduleProcedure,
			opts...,
		),
		readPagerDutySchedules: connect_go.NewClient[v1alpha1.ReadPagerDutySchedulesRequest, v1alpha1.ReadPagerDutySchedulesResponse](
			httpClient,
			baseURL+PagerDutyServiceReadPagerDutySchedulesProcedure,
			opts...,
		),
	}
}

// pagerDutyServiceClient implements PagerDutyServiceClient.
type pagerDutyServiceClient struct {
	createSchedule         *connect_go.Client[v1alpha1.CreateScheduleRequest, v1alpha1.CreateScheduleResponse]
	readSchedule           *connect_go.Client[v1alpha1.ReadScheduleRequest, v1alpha1.ReadScheduleResponse]
	updateSchedule         *connect_go.Client[v1alpha1.UpdateScheduleRequest, v1alpha1.UpdateScheduleResponse]
	deleteSchedule         *connect_go.Client[v1alpha1.DeleteScheduleRequest, v1alpha1.DeleteScheduleResponse]
	readPagerDutySchedules *connect_go.Client[v1alpha1.ReadPagerDutySchedulesRequest, v1alpha1.ReadPagerDutySchedulesResponse]
}

// CreateSchedule calls commonfate.control.config.v1alpha1.PagerDutyService.CreateSchedule.
func (c *pagerDutyServiceClient) CreateSchedule(ctx context.Context, req *connect_go.Request[v1alpha1.CreateScheduleRequest]) (*connect_go.Response[v1alpha1.CreateScheduleResponse], error) {
	return c.createSchedule.CallUnary(ctx, req)
}

// ReadSchedule calls commonfate.control.config.v1alpha1.PagerDutyService.ReadSchedule.
func (c *pagerDutyServiceClient) ReadSchedule(ctx context.Context, req *connect_go.Request[v1alpha1.ReadScheduleRequest]) (*connect_go.Response[v1alpha1.ReadScheduleResponse], error) {
	return c.readSchedule.CallUnary(ctx, req)
}

// UpdateSchedule calls commonfate.control.config.v1alpha1.PagerDutyService.UpdateSchedule.
func (c *pagerDutyServiceClient) UpdateSchedule(ctx context.Context, req *connect_go.Request[v1alpha1.UpdateScheduleRequest]) (*connect_go.Response[v1alpha1.UpdateScheduleResponse], error) {
	return c.updateSchedule.CallUnary(ctx, req)
}

// DeleteSchedule calls commonfate.control.config.v1alpha1.PagerDutyService.DeleteSchedule.
func (c *pagerDutyServiceClient) DeleteSchedule(ctx context.Context, req *connect_go.Request[v1alpha1.DeleteScheduleRequest]) (*connect_go.Response[v1alpha1.DeleteScheduleResponse], error) {
	return c.deleteSchedule.CallUnary(ctx, req)
}

// ReadPagerDutySchedules calls
// commonfate.control.config.v1alpha1.PagerDutyService.ReadPagerDutySchedules.
func (c *pagerDutyServiceClient) ReadPagerDutySchedules(ctx context.Context, req *connect_go.Request[v1alpha1.ReadPagerDutySchedulesRequest]) (*connect_go.Response[v1alpha1.ReadPagerDutySchedulesResponse], error) {
	return c.readPagerDutySchedules.CallUnary(ctx, req)
}

// PagerDutyServiceHandler is an implementation of the
// commonfate.control.config.v1alpha1.PagerDutyService service.
type PagerDutyServiceHandler interface {
	CreateSchedule(context.Context, *connect_go.Request[v1alpha1.CreateScheduleRequest]) (*connect_go.Response[v1alpha1.CreateScheduleResponse], error)
	ReadSchedule(context.Context, *connect_go.Request[v1alpha1.ReadScheduleRequest]) (*connect_go.Response[v1alpha1.ReadScheduleResponse], error)
	UpdateSchedule(context.Context, *connect_go.Request[v1alpha1.UpdateScheduleRequest]) (*connect_go.Response[v1alpha1.UpdateScheduleResponse], error)
	DeleteSchedule(context.Context, *connect_go.Request[v1alpha1.DeleteScheduleRequest]) (*connect_go.Response[v1alpha1.DeleteScheduleResponse], error)
	ReadPagerDutySchedules(context.Context, *connect_go.Request[v1alpha1.ReadPagerDutySchedulesRequest]) (*connect_go.Response[v1alpha1.ReadPagerDutySchedulesResponse], error)
}

// NewPagerDutyServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPagerDutyServiceHandler(svc PagerDutyServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	pagerDutyServiceCreateScheduleHandler := connect_go.NewUnaryHandler(
		PagerDutyServiceCreateScheduleProcedure,
		svc.CreateSchedule,
		opts...,
	)
	pagerDutyServiceReadScheduleHandler := connect_go.NewUnaryHandler(
		PagerDutyServiceReadScheduleProcedure,
		svc.ReadSchedule,
		opts...,
	)
	pagerDutyServiceUpdateScheduleHandler := connect_go.NewUnaryHandler(
		PagerDutyServiceUpdateScheduleProcedure,
		svc.UpdateSchedule,
		opts...,
	)
	pagerDutyServiceDeleteScheduleHandler := connect_go.NewUnaryHandler(
		PagerDutyServiceDeleteScheduleProcedure,
		svc.DeleteSchedule,
		opts...,
	)
	pagerDutyServiceReadPagerDutySchedulesHandler := connect_go.NewUnaryHandler(
		PagerDutyServiceReadPagerDutySchedulesProcedure,
		svc.ReadPagerDutySchedules,
		opts...,
	)
	return "/commonfate.control.config.v1alpha1.PagerDutyService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PagerDutyServiceCreateScheduleProcedure:
			pagerDutyServiceCreateScheduleHandler.ServeHTTP(w, r)
		case PagerDutyServiceReadScheduleProcedure:
			pagerDutyServiceReadScheduleHandler.ServeHTTP(w, r)
		case PagerDutyServiceUpdateScheduleProcedure:
			pagerDutyServiceUpdateScheduleHandler.ServeHTTP(w, r)
		case PagerDutyServiceDeleteScheduleProcedure:
			pagerDutyServiceDeleteScheduleHandler.ServeHTTP(w, r)
		case PagerDutyServiceReadPagerDutySchedulesProcedure:
			pagerDutyServiceReadPagerDutySchedulesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPagerDutyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPagerDutyServiceHandler struct{}

func (UnimplementedPagerDutyServiceHandler) CreateSchedule(context.Context, *connect_go.Request[v1alpha1.CreateScheduleRequest]) (*connect_go.Response[v1alpha1.CreateScheduleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.PagerDutyService.CreateSchedule is not implemented"))
}

func (UnimplementedPagerDutyServiceHandler) ReadSchedule(context.Context, *connect_go.Request[v1alpha1.ReadScheduleRequest]) (*connect_go.Response[v1alpha1.ReadScheduleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.PagerDutyService.ReadSchedule is not implemented"))
}

func (UnimplementedPagerDutyServiceHandler) UpdateSchedule(context.Context, *connect_go.Request[v1alpha1.UpdateScheduleRequest]) (*connect_go.Response[v1alpha1.UpdateScheduleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.PagerDutyService.UpdateSchedule is not implemented"))
}

func (UnimplementedPagerDutyServiceHandler) DeleteSchedule(context.Context, *connect_go.Request[v1alpha1.DeleteScheduleRequest]) (*connect_go.Response[v1alpha1.DeleteScheduleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.PagerDutyService.DeleteSchedule is not implemented"))
}

func (UnimplementedPagerDutyServiceHandler) ReadPagerDutySchedules(context.Context, *connect_go.Request[v1alpha1.ReadPagerDutySchedulesRequest]) (*connect_go.Response[v1alpha1.ReadPagerDutySchedulesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.PagerDutyService.ReadPagerDutySchedules is not implemented"))
}