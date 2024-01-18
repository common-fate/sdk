// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/config/v1alpha1/selector.proto

package configv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SelectorServiceName is the fully-qualified name of the SelectorService service.
	SelectorServiceName = "commonfate.control.config.v1alpha1.SelectorService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SelectorServiceCreateSelectorProcedure is the fully-qualified name of the SelectorService's
	// CreateSelector RPC.
	SelectorServiceCreateSelectorProcedure = "/commonfate.control.config.v1alpha1.SelectorService/CreateSelector"
	// SelectorServiceGetSelectorProcedure is the fully-qualified name of the SelectorService's
	// GetSelector RPC.
	SelectorServiceGetSelectorProcedure = "/commonfate.control.config.v1alpha1.SelectorService/GetSelector"
	// SelectorServiceUpdateSelectorProcedure is the fully-qualified name of the SelectorService's
	// UpdateSelector RPC.
	SelectorServiceUpdateSelectorProcedure = "/commonfate.control.config.v1alpha1.SelectorService/UpdateSelector"
	// SelectorServiceDeleteSelectorProcedure is the fully-qualified name of the SelectorService's
	// DeleteSelector RPC.
	SelectorServiceDeleteSelectorProcedure = "/commonfate.control.config.v1alpha1.SelectorService/DeleteSelector"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	selectorServiceServiceDescriptor              = v1alpha1.File_commonfate_control_config_v1alpha1_selector_proto.Services().ByName("SelectorService")
	selectorServiceCreateSelectorMethodDescriptor = selectorServiceServiceDescriptor.Methods().ByName("CreateSelector")
	selectorServiceGetSelectorMethodDescriptor    = selectorServiceServiceDescriptor.Methods().ByName("GetSelector")
	selectorServiceUpdateSelectorMethodDescriptor = selectorServiceServiceDescriptor.Methods().ByName("UpdateSelector")
	selectorServiceDeleteSelectorMethodDescriptor = selectorServiceServiceDescriptor.Methods().ByName("DeleteSelector")
)

// SelectorServiceClient is a client for the commonfate.control.config.v1alpha1.SelectorService
// service.
type SelectorServiceClient interface {
	CreateSelector(context.Context, *connect.Request[v1alpha1.CreateSelectorRequest]) (*connect.Response[v1alpha1.CreateSelectorResponse], error)
	GetSelector(context.Context, *connect.Request[v1alpha1.GetSelectorRequest]) (*connect.Response[v1alpha1.GetSelectorResponse], error)
	UpdateSelector(context.Context, *connect.Request[v1alpha1.UpdateSelectorRequest]) (*connect.Response[v1alpha1.UpdateSelectorResponse], error)
	DeleteSelector(context.Context, *connect.Request[v1alpha1.DeleteSelectorRequest]) (*connect.Response[v1alpha1.DeleteSelectorResponse], error)
}

// NewSelectorServiceClient constructs a client for the
// commonfate.control.config.v1alpha1.SelectorService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSelectorServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SelectorServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &selectorServiceClient{
		createSelector: connect.NewClient[v1alpha1.CreateSelectorRequest, v1alpha1.CreateSelectorResponse](
			httpClient,
			baseURL+SelectorServiceCreateSelectorProcedure,
			connect.WithSchema(selectorServiceCreateSelectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getSelector: connect.NewClient[v1alpha1.GetSelectorRequest, v1alpha1.GetSelectorResponse](
			httpClient,
			baseURL+SelectorServiceGetSelectorProcedure,
			connect.WithSchema(selectorServiceGetSelectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateSelector: connect.NewClient[v1alpha1.UpdateSelectorRequest, v1alpha1.UpdateSelectorResponse](
			httpClient,
			baseURL+SelectorServiceUpdateSelectorProcedure,
			connect.WithSchema(selectorServiceUpdateSelectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteSelector: connect.NewClient[v1alpha1.DeleteSelectorRequest, v1alpha1.DeleteSelectorResponse](
			httpClient,
			baseURL+SelectorServiceDeleteSelectorProcedure,
			connect.WithSchema(selectorServiceDeleteSelectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// selectorServiceClient implements SelectorServiceClient.
type selectorServiceClient struct {
	createSelector *connect.Client[v1alpha1.CreateSelectorRequest, v1alpha1.CreateSelectorResponse]
	getSelector    *connect.Client[v1alpha1.GetSelectorRequest, v1alpha1.GetSelectorResponse]
	updateSelector *connect.Client[v1alpha1.UpdateSelectorRequest, v1alpha1.UpdateSelectorResponse]
	deleteSelector *connect.Client[v1alpha1.DeleteSelectorRequest, v1alpha1.DeleteSelectorResponse]
}

// CreateSelector calls commonfate.control.config.v1alpha1.SelectorService.CreateSelector.
func (c *selectorServiceClient) CreateSelector(ctx context.Context, req *connect.Request[v1alpha1.CreateSelectorRequest]) (*connect.Response[v1alpha1.CreateSelectorResponse], error) {
	return c.createSelector.CallUnary(ctx, req)
}

// GetSelector calls commonfate.control.config.v1alpha1.SelectorService.GetSelector.
func (c *selectorServiceClient) GetSelector(ctx context.Context, req *connect.Request[v1alpha1.GetSelectorRequest]) (*connect.Response[v1alpha1.GetSelectorResponse], error) {
	return c.getSelector.CallUnary(ctx, req)
}

// UpdateSelector calls commonfate.control.config.v1alpha1.SelectorService.UpdateSelector.
func (c *selectorServiceClient) UpdateSelector(ctx context.Context, req *connect.Request[v1alpha1.UpdateSelectorRequest]) (*connect.Response[v1alpha1.UpdateSelectorResponse], error) {
	return c.updateSelector.CallUnary(ctx, req)
}

// DeleteSelector calls commonfate.control.config.v1alpha1.SelectorService.DeleteSelector.
func (c *selectorServiceClient) DeleteSelector(ctx context.Context, req *connect.Request[v1alpha1.DeleteSelectorRequest]) (*connect.Response[v1alpha1.DeleteSelectorResponse], error) {
	return c.deleteSelector.CallUnary(ctx, req)
}

// SelectorServiceHandler is an implementation of the
// commonfate.control.config.v1alpha1.SelectorService service.
type SelectorServiceHandler interface {
	CreateSelector(context.Context, *connect.Request[v1alpha1.CreateSelectorRequest]) (*connect.Response[v1alpha1.CreateSelectorResponse], error)
	GetSelector(context.Context, *connect.Request[v1alpha1.GetSelectorRequest]) (*connect.Response[v1alpha1.GetSelectorResponse], error)
	UpdateSelector(context.Context, *connect.Request[v1alpha1.UpdateSelectorRequest]) (*connect.Response[v1alpha1.UpdateSelectorResponse], error)
	DeleteSelector(context.Context, *connect.Request[v1alpha1.DeleteSelectorRequest]) (*connect.Response[v1alpha1.DeleteSelectorResponse], error)
}

// NewSelectorServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSelectorServiceHandler(svc SelectorServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	selectorServiceCreateSelectorHandler := connect.NewUnaryHandler(
		SelectorServiceCreateSelectorProcedure,
		svc.CreateSelector,
		connect.WithSchema(selectorServiceCreateSelectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	selectorServiceGetSelectorHandler := connect.NewUnaryHandler(
		SelectorServiceGetSelectorProcedure,
		svc.GetSelector,
		connect.WithSchema(selectorServiceGetSelectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	selectorServiceUpdateSelectorHandler := connect.NewUnaryHandler(
		SelectorServiceUpdateSelectorProcedure,
		svc.UpdateSelector,
		connect.WithSchema(selectorServiceUpdateSelectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	selectorServiceDeleteSelectorHandler := connect.NewUnaryHandler(
		SelectorServiceDeleteSelectorProcedure,
		svc.DeleteSelector,
		connect.WithSchema(selectorServiceDeleteSelectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.control.config.v1alpha1.SelectorService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SelectorServiceCreateSelectorProcedure:
			selectorServiceCreateSelectorHandler.ServeHTTP(w, r)
		case SelectorServiceGetSelectorProcedure:
			selectorServiceGetSelectorHandler.ServeHTTP(w, r)
		case SelectorServiceUpdateSelectorProcedure:
			selectorServiceUpdateSelectorHandler.ServeHTTP(w, r)
		case SelectorServiceDeleteSelectorProcedure:
			selectorServiceDeleteSelectorHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSelectorServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSelectorServiceHandler struct{}

func (UnimplementedSelectorServiceHandler) CreateSelector(context.Context, *connect.Request[v1alpha1.CreateSelectorRequest]) (*connect.Response[v1alpha1.CreateSelectorResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.SelectorService.CreateSelector is not implemented"))
}

func (UnimplementedSelectorServiceHandler) GetSelector(context.Context, *connect.Request[v1alpha1.GetSelectorRequest]) (*connect.Response[v1alpha1.GetSelectorResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.SelectorService.GetSelector is not implemented"))
}

func (UnimplementedSelectorServiceHandler) UpdateSelector(context.Context, *connect.Request[v1alpha1.UpdateSelectorRequest]) (*connect.Response[v1alpha1.UpdateSelectorResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.SelectorService.UpdateSelector is not implemented"))
}

func (UnimplementedSelectorServiceHandler) DeleteSelector(context.Context, *connect.Request[v1alpha1.DeleteSelectorRequest]) (*connect.Response[v1alpha1.DeleteSelectorResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.SelectorService.DeleteSelector is not implemented"))
}
