// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/integration/v1alpha1/integration.proto

package integrationv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
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
	// IntegrationServiceName is the fully-qualified name of the IntegrationService service.
	IntegrationServiceName = "commonfate.control.integration.v1alpha1.IntegrationService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IntegrationServiceCreateIntegrationProcedure is the fully-qualified name of the
	// IntegrationService's CreateIntegration RPC.
	IntegrationServiceCreateIntegrationProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/CreateIntegration"
	// IntegrationServiceUpdateIntegrationProcedure is the fully-qualified name of the
	// IntegrationService's UpdateIntegration RPC.
	IntegrationServiceUpdateIntegrationProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/UpdateIntegration"
	// IntegrationServiceGetIntegrationProcedure is the fully-qualified name of the IntegrationService's
	// GetIntegration RPC.
	IntegrationServiceGetIntegrationProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/GetIntegration"
	// IntegrationServiceDeleteIntegrationProcedure is the fully-qualified name of the
	// IntegrationService's DeleteIntegration RPC.
	IntegrationServiceDeleteIntegrationProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/DeleteIntegration"
	// IntegrationServiceListIntegrationsProcedure is the fully-qualified name of the
	// IntegrationService's ListIntegrations RPC.
	IntegrationServiceListIntegrationsProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/ListIntegrations"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	integrationServiceServiceDescriptor                 = v1alpha1.File_commonfate_control_integration_v1alpha1_integration_proto.Services().ByName("IntegrationService")
	integrationServiceCreateIntegrationMethodDescriptor = integrationServiceServiceDescriptor.Methods().ByName("CreateIntegration")
	integrationServiceUpdateIntegrationMethodDescriptor = integrationServiceServiceDescriptor.Methods().ByName("UpdateIntegration")
	integrationServiceGetIntegrationMethodDescriptor    = integrationServiceServiceDescriptor.Methods().ByName("GetIntegration")
	integrationServiceDeleteIntegrationMethodDescriptor = integrationServiceServiceDescriptor.Methods().ByName("DeleteIntegration")
	integrationServiceListIntegrationsMethodDescriptor  = integrationServiceServiceDescriptor.Methods().ByName("ListIntegrations")
)

// IntegrationServiceClient is a client for the
// commonfate.control.integration.v1alpha1.IntegrationService service.
type IntegrationServiceClient interface {
	CreateIntegration(context.Context, *connect.Request[v1alpha1.CreateIntegrationRequest]) (*connect.Response[v1alpha1.CreateIntegrationResponse], error)
	UpdateIntegration(context.Context, *connect.Request[v1alpha1.UpdateIntegrationRequest]) (*connect.Response[v1alpha1.UpdateIntegrationResponse], error)
	GetIntegration(context.Context, *connect.Request[v1alpha1.GetIntegrationRequest]) (*connect.Response[v1alpha1.GetIntegrationResponse], error)
	DeleteIntegration(context.Context, *connect.Request[v1alpha1.DeleteIntegrationRequest]) (*connect.Response[v1alpha1.DeleteIntegrationResponse], error)
	ListIntegrations(context.Context, *connect.Request[v1alpha1.ListIntegrationsRequest]) (*connect.Response[v1alpha1.ListIntegrationsResponse], error)
}

// NewIntegrationServiceClient constructs a client for the
// commonfate.control.integration.v1alpha1.IntegrationService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIntegrationServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) IntegrationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &integrationServiceClient{
		createIntegration: connect.NewClient[v1alpha1.CreateIntegrationRequest, v1alpha1.CreateIntegrationResponse](
			httpClient,
			baseURL+IntegrationServiceCreateIntegrationProcedure,
			connect.WithSchema(integrationServiceCreateIntegrationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateIntegration: connect.NewClient[v1alpha1.UpdateIntegrationRequest, v1alpha1.UpdateIntegrationResponse](
			httpClient,
			baseURL+IntegrationServiceUpdateIntegrationProcedure,
			connect.WithSchema(integrationServiceUpdateIntegrationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getIntegration: connect.NewClient[v1alpha1.GetIntegrationRequest, v1alpha1.GetIntegrationResponse](
			httpClient,
			baseURL+IntegrationServiceGetIntegrationProcedure,
			connect.WithSchema(integrationServiceGetIntegrationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteIntegration: connect.NewClient[v1alpha1.DeleteIntegrationRequest, v1alpha1.DeleteIntegrationResponse](
			httpClient,
			baseURL+IntegrationServiceDeleteIntegrationProcedure,
			connect.WithSchema(integrationServiceDeleteIntegrationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listIntegrations: connect.NewClient[v1alpha1.ListIntegrationsRequest, v1alpha1.ListIntegrationsResponse](
			httpClient,
			baseURL+IntegrationServiceListIntegrationsProcedure,
			connect.WithSchema(integrationServiceListIntegrationsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// integrationServiceClient implements IntegrationServiceClient.
type integrationServiceClient struct {
	createIntegration *connect.Client[v1alpha1.CreateIntegrationRequest, v1alpha1.CreateIntegrationResponse]
	updateIntegration *connect.Client[v1alpha1.UpdateIntegrationRequest, v1alpha1.UpdateIntegrationResponse]
	getIntegration    *connect.Client[v1alpha1.GetIntegrationRequest, v1alpha1.GetIntegrationResponse]
	deleteIntegration *connect.Client[v1alpha1.DeleteIntegrationRequest, v1alpha1.DeleteIntegrationResponse]
	listIntegrations  *connect.Client[v1alpha1.ListIntegrationsRequest, v1alpha1.ListIntegrationsResponse]
}

// CreateIntegration calls
// commonfate.control.integration.v1alpha1.IntegrationService.CreateIntegration.
func (c *integrationServiceClient) CreateIntegration(ctx context.Context, req *connect.Request[v1alpha1.CreateIntegrationRequest]) (*connect.Response[v1alpha1.CreateIntegrationResponse], error) {
	return c.createIntegration.CallUnary(ctx, req)
}

// UpdateIntegration calls
// commonfate.control.integration.v1alpha1.IntegrationService.UpdateIntegration.
func (c *integrationServiceClient) UpdateIntegration(ctx context.Context, req *connect.Request[v1alpha1.UpdateIntegrationRequest]) (*connect.Response[v1alpha1.UpdateIntegrationResponse], error) {
	return c.updateIntegration.CallUnary(ctx, req)
}

// GetIntegration calls commonfate.control.integration.v1alpha1.IntegrationService.GetIntegration.
func (c *integrationServiceClient) GetIntegration(ctx context.Context, req *connect.Request[v1alpha1.GetIntegrationRequest]) (*connect.Response[v1alpha1.GetIntegrationResponse], error) {
	return c.getIntegration.CallUnary(ctx, req)
}

// DeleteIntegration calls
// commonfate.control.integration.v1alpha1.IntegrationService.DeleteIntegration.
func (c *integrationServiceClient) DeleteIntegration(ctx context.Context, req *connect.Request[v1alpha1.DeleteIntegrationRequest]) (*connect.Response[v1alpha1.DeleteIntegrationResponse], error) {
	return c.deleteIntegration.CallUnary(ctx, req)
}

// ListIntegrations calls
// commonfate.control.integration.v1alpha1.IntegrationService.ListIntegrations.
func (c *integrationServiceClient) ListIntegrations(ctx context.Context, req *connect.Request[v1alpha1.ListIntegrationsRequest]) (*connect.Response[v1alpha1.ListIntegrationsResponse], error) {
	return c.listIntegrations.CallUnary(ctx, req)
}

// IntegrationServiceHandler is an implementation of the
// commonfate.control.integration.v1alpha1.IntegrationService service.
type IntegrationServiceHandler interface {
	CreateIntegration(context.Context, *connect.Request[v1alpha1.CreateIntegrationRequest]) (*connect.Response[v1alpha1.CreateIntegrationResponse], error)
	UpdateIntegration(context.Context, *connect.Request[v1alpha1.UpdateIntegrationRequest]) (*connect.Response[v1alpha1.UpdateIntegrationResponse], error)
	GetIntegration(context.Context, *connect.Request[v1alpha1.GetIntegrationRequest]) (*connect.Response[v1alpha1.GetIntegrationResponse], error)
	DeleteIntegration(context.Context, *connect.Request[v1alpha1.DeleteIntegrationRequest]) (*connect.Response[v1alpha1.DeleteIntegrationResponse], error)
	ListIntegrations(context.Context, *connect.Request[v1alpha1.ListIntegrationsRequest]) (*connect.Response[v1alpha1.ListIntegrationsResponse], error)
}

// NewIntegrationServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIntegrationServiceHandler(svc IntegrationServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	integrationServiceCreateIntegrationHandler := connect.NewUnaryHandler(
		IntegrationServiceCreateIntegrationProcedure,
		svc.CreateIntegration,
		connect.WithSchema(integrationServiceCreateIntegrationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceUpdateIntegrationHandler := connect.NewUnaryHandler(
		IntegrationServiceUpdateIntegrationProcedure,
		svc.UpdateIntegration,
		connect.WithSchema(integrationServiceUpdateIntegrationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceGetIntegrationHandler := connect.NewUnaryHandler(
		IntegrationServiceGetIntegrationProcedure,
		svc.GetIntegration,
		connect.WithSchema(integrationServiceGetIntegrationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceDeleteIntegrationHandler := connect.NewUnaryHandler(
		IntegrationServiceDeleteIntegrationProcedure,
		svc.DeleteIntegration,
		connect.WithSchema(integrationServiceDeleteIntegrationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceListIntegrationsHandler := connect.NewUnaryHandler(
		IntegrationServiceListIntegrationsProcedure,
		svc.ListIntegrations,
		connect.WithSchema(integrationServiceListIntegrationsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.control.integration.v1alpha1.IntegrationService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IntegrationServiceCreateIntegrationProcedure:
			integrationServiceCreateIntegrationHandler.ServeHTTP(w, r)
		case IntegrationServiceUpdateIntegrationProcedure:
			integrationServiceUpdateIntegrationHandler.ServeHTTP(w, r)
		case IntegrationServiceGetIntegrationProcedure:
			integrationServiceGetIntegrationHandler.ServeHTTP(w, r)
		case IntegrationServiceDeleteIntegrationProcedure:
			integrationServiceDeleteIntegrationHandler.ServeHTTP(w, r)
		case IntegrationServiceListIntegrationsProcedure:
			integrationServiceListIntegrationsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIntegrationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIntegrationServiceHandler struct{}

func (UnimplementedIntegrationServiceHandler) CreateIntegration(context.Context, *connect.Request[v1alpha1.CreateIntegrationRequest]) (*connect.Response[v1alpha1.CreateIntegrationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.CreateIntegration is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) UpdateIntegration(context.Context, *connect.Request[v1alpha1.UpdateIntegrationRequest]) (*connect.Response[v1alpha1.UpdateIntegrationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.UpdateIntegration is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) GetIntegration(context.Context, *connect.Request[v1alpha1.GetIntegrationRequest]) (*connect.Response[v1alpha1.GetIntegrationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.GetIntegration is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) DeleteIntegration(context.Context, *connect.Request[v1alpha1.DeleteIntegrationRequest]) (*connect.Response[v1alpha1.DeleteIntegrationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.DeleteIntegration is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) ListIntegrations(context.Context, *connect.Request[v1alpha1.ListIntegrationsRequest]) (*connect.Response[v1alpha1.ListIntegrationsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.ListIntegrations is not implemented"))
}
