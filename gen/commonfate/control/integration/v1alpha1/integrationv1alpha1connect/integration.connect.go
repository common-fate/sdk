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
	// IntegrationServiceCreateProxyResourceProcedure is the fully-qualified name of the
	// IntegrationService's CreateProxyResource RPC.
	IntegrationServiceCreateProxyResourceProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/CreateProxyResource"
	// IntegrationServiceUpdateProxyResourceProcedure is the fully-qualified name of the
	// IntegrationService's UpdateProxyResource RPC.
	IntegrationServiceUpdateProxyResourceProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/UpdateProxyResource"
	// IntegrationServiceDeleteProxyResourceProcedure is the fully-qualified name of the
	// IntegrationService's DeleteProxyResource RPC.
	IntegrationServiceDeleteProxyResourceProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/DeleteProxyResource"
	// IntegrationServiceGetProxyResourceProcedure is the fully-qualified name of the
	// IntegrationService's GetProxyResource RPC.
	IntegrationServiceGetProxyResourceProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/GetProxyResource"
	// IntegrationServiceCreateProxyProcedure is the fully-qualified name of the IntegrationService's
	// CreateProxy RPC.
	IntegrationServiceCreateProxyProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/CreateProxy"
	// IntegrationServiceUpdateProxyProcedure is the fully-qualified name of the IntegrationService's
	// UpdateProxy RPC.
	IntegrationServiceUpdateProxyProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/UpdateProxy"
	// IntegrationServiceDeleteProxyProcedure is the fully-qualified name of the IntegrationService's
	// DeleteProxy RPC.
	IntegrationServiceDeleteProxyProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/DeleteProxy"
	// IntegrationServiceGetProxyProcedure is the fully-qualified name of the IntegrationService's
	// GetProxy RPC.
	IntegrationServiceGetProxyProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/GetProxy"
	// IntegrationServiceListProxyResourcesProcedure is the fully-qualified name of the
	// IntegrationService's ListProxyResources RPC.
	IntegrationServiceListProxyResourcesProcedure = "/commonfate.control.integration.v1alpha1.IntegrationService/ListProxyResources"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	integrationServiceServiceDescriptor                   = v1alpha1.File_commonfate_control_integration_v1alpha1_integration_proto.Services().ByName("IntegrationService")
	integrationServiceCreateIntegrationMethodDescriptor   = integrationServiceServiceDescriptor.Methods().ByName("CreateIntegration")
	integrationServiceUpdateIntegrationMethodDescriptor   = integrationServiceServiceDescriptor.Methods().ByName("UpdateIntegration")
	integrationServiceGetIntegrationMethodDescriptor      = integrationServiceServiceDescriptor.Methods().ByName("GetIntegration")
	integrationServiceDeleteIntegrationMethodDescriptor   = integrationServiceServiceDescriptor.Methods().ByName("DeleteIntegration")
	integrationServiceListIntegrationsMethodDescriptor    = integrationServiceServiceDescriptor.Methods().ByName("ListIntegrations")
	integrationServiceCreateProxyResourceMethodDescriptor = integrationServiceServiceDescriptor.Methods().ByName("CreateProxyResource")
	integrationServiceUpdateProxyResourceMethodDescriptor = integrationServiceServiceDescriptor.Methods().ByName("UpdateProxyResource")
	integrationServiceDeleteProxyResourceMethodDescriptor = integrationServiceServiceDescriptor.Methods().ByName("DeleteProxyResource")
	integrationServiceGetProxyResourceMethodDescriptor    = integrationServiceServiceDescriptor.Methods().ByName("GetProxyResource")
	integrationServiceCreateProxyMethodDescriptor         = integrationServiceServiceDescriptor.Methods().ByName("CreateProxy")
	integrationServiceUpdateProxyMethodDescriptor         = integrationServiceServiceDescriptor.Methods().ByName("UpdateProxy")
	integrationServiceDeleteProxyMethodDescriptor         = integrationServiceServiceDescriptor.Methods().ByName("DeleteProxy")
	integrationServiceGetProxyMethodDescriptor            = integrationServiceServiceDescriptor.Methods().ByName("GetProxy")
	integrationServiceListProxyResourcesMethodDescriptor  = integrationServiceServiceDescriptor.Methods().ByName("ListProxyResources")
)

// IntegrationServiceClient is a client for the
// commonfate.control.integration.v1alpha1.IntegrationService service.
type IntegrationServiceClient interface {
	CreateIntegration(context.Context, *connect.Request[v1alpha1.CreateIntegrationRequest]) (*connect.Response[v1alpha1.CreateIntegrationResponse], error)
	UpdateIntegration(context.Context, *connect.Request[v1alpha1.UpdateIntegrationRequest]) (*connect.Response[v1alpha1.UpdateIntegrationResponse], error)
	GetIntegration(context.Context, *connect.Request[v1alpha1.GetIntegrationRequest]) (*connect.Response[v1alpha1.GetIntegrationResponse], error)
	DeleteIntegration(context.Context, *connect.Request[v1alpha1.DeleteIntegrationRequest]) (*connect.Response[v1alpha1.DeleteIntegrationResponse], error)
	ListIntegrations(context.Context, *connect.Request[v1alpha1.ListIntegrationsRequest]) (*connect.Response[v1alpha1.ListIntegrationsResponse], error)
	// CRUD operations for proxy resource terraform provider resource
	CreateProxyResource(context.Context, *connect.Request[v1alpha1.CreateProxyResourceRequest]) (*connect.Response[v1alpha1.CreateProxyResourceResponse], error)
	UpdateProxyResource(context.Context, *connect.Request[v1alpha1.UpdateProxyResourceRequest]) (*connect.Response[v1alpha1.UpdateProxyResourceResponse], error)
	DeleteProxyResource(context.Context, *connect.Request[v1alpha1.DeleteProxyResourceRequest]) (*connect.Response[v1alpha1.DeleteProxyResourceResponse], error)
	GetProxyResource(context.Context, *connect.Request[v1alpha1.GetProxyResourceRequest]) (*connect.Response[v1alpha1.GetProxyResourceResponse], error)
	// CRUD operations for proxy terraform provider resource
	CreateProxy(context.Context, *connect.Request[v1alpha1.CreateProxyRequest]) (*connect.Response[v1alpha1.CreateProxyResponse], error)
	UpdateProxy(context.Context, *connect.Request[v1alpha1.UpdateProxyRequest]) (*connect.Response[v1alpha1.UpdateProxyResponse], error)
	DeleteProxy(context.Context, *connect.Request[v1alpha1.DeleteProxyRequest]) (*connect.Response[v1alpha1.DeleteProxyResponse], error)
	GetProxy(context.Context, *connect.Request[v1alpha1.GetProxyRequest]) (*connect.Response[v1alpha1.GetProxyResponse], error)
	// Used by the proxy to get resources
	ListProxyResources(context.Context, *connect.Request[v1alpha1.DescribeProxyResourcesRequest]) (*connect.Response[v1alpha1.DescribeProxyResourcesResponse], error)
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
		createProxyResource: connect.NewClient[v1alpha1.CreateProxyResourceRequest, v1alpha1.CreateProxyResourceResponse](
			httpClient,
			baseURL+IntegrationServiceCreateProxyResourceProcedure,
			connect.WithSchema(integrationServiceCreateProxyResourceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateProxyResource: connect.NewClient[v1alpha1.UpdateProxyResourceRequest, v1alpha1.UpdateProxyResourceResponse](
			httpClient,
			baseURL+IntegrationServiceUpdateProxyResourceProcedure,
			connect.WithSchema(integrationServiceUpdateProxyResourceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteProxyResource: connect.NewClient[v1alpha1.DeleteProxyResourceRequest, v1alpha1.DeleteProxyResourceResponse](
			httpClient,
			baseURL+IntegrationServiceDeleteProxyResourceProcedure,
			connect.WithSchema(integrationServiceDeleteProxyResourceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getProxyResource: connect.NewClient[v1alpha1.GetProxyResourceRequest, v1alpha1.GetProxyResourceResponse](
			httpClient,
			baseURL+IntegrationServiceGetProxyResourceProcedure,
			connect.WithSchema(integrationServiceGetProxyResourceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createProxy: connect.NewClient[v1alpha1.CreateProxyRequest, v1alpha1.CreateProxyResponse](
			httpClient,
			baseURL+IntegrationServiceCreateProxyProcedure,
			connect.WithSchema(integrationServiceCreateProxyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateProxy: connect.NewClient[v1alpha1.UpdateProxyRequest, v1alpha1.UpdateProxyResponse](
			httpClient,
			baseURL+IntegrationServiceUpdateProxyProcedure,
			connect.WithSchema(integrationServiceUpdateProxyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteProxy: connect.NewClient[v1alpha1.DeleteProxyRequest, v1alpha1.DeleteProxyResponse](
			httpClient,
			baseURL+IntegrationServiceDeleteProxyProcedure,
			connect.WithSchema(integrationServiceDeleteProxyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getProxy: connect.NewClient[v1alpha1.GetProxyRequest, v1alpha1.GetProxyResponse](
			httpClient,
			baseURL+IntegrationServiceGetProxyProcedure,
			connect.WithSchema(integrationServiceGetProxyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listProxyResources: connect.NewClient[v1alpha1.DescribeProxyResourcesRequest, v1alpha1.DescribeProxyResourcesResponse](
			httpClient,
			baseURL+IntegrationServiceListProxyResourcesProcedure,
			connect.WithSchema(integrationServiceListProxyResourcesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// integrationServiceClient implements IntegrationServiceClient.
type integrationServiceClient struct {
	createIntegration   *connect.Client[v1alpha1.CreateIntegrationRequest, v1alpha1.CreateIntegrationResponse]
	updateIntegration   *connect.Client[v1alpha1.UpdateIntegrationRequest, v1alpha1.UpdateIntegrationResponse]
	getIntegration      *connect.Client[v1alpha1.GetIntegrationRequest, v1alpha1.GetIntegrationResponse]
	deleteIntegration   *connect.Client[v1alpha1.DeleteIntegrationRequest, v1alpha1.DeleteIntegrationResponse]
	listIntegrations    *connect.Client[v1alpha1.ListIntegrationsRequest, v1alpha1.ListIntegrationsResponse]
	createProxyResource *connect.Client[v1alpha1.CreateProxyResourceRequest, v1alpha1.CreateProxyResourceResponse]
	updateProxyResource *connect.Client[v1alpha1.UpdateProxyResourceRequest, v1alpha1.UpdateProxyResourceResponse]
	deleteProxyResource *connect.Client[v1alpha1.DeleteProxyResourceRequest, v1alpha1.DeleteProxyResourceResponse]
	getProxyResource    *connect.Client[v1alpha1.GetProxyResourceRequest, v1alpha1.GetProxyResourceResponse]
	createProxy         *connect.Client[v1alpha1.CreateProxyRequest, v1alpha1.CreateProxyResponse]
	updateProxy         *connect.Client[v1alpha1.UpdateProxyRequest, v1alpha1.UpdateProxyResponse]
	deleteProxy         *connect.Client[v1alpha1.DeleteProxyRequest, v1alpha1.DeleteProxyResponse]
	getProxy            *connect.Client[v1alpha1.GetProxyRequest, v1alpha1.GetProxyResponse]
	listProxyResources  *connect.Client[v1alpha1.DescribeProxyResourcesRequest, v1alpha1.DescribeProxyResourcesResponse]
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

// CreateProxyResource calls
// commonfate.control.integration.v1alpha1.IntegrationService.CreateProxyResource.
func (c *integrationServiceClient) CreateProxyResource(ctx context.Context, req *connect.Request[v1alpha1.CreateProxyResourceRequest]) (*connect.Response[v1alpha1.CreateProxyResourceResponse], error) {
	return c.createProxyResource.CallUnary(ctx, req)
}

// UpdateProxyResource calls
// commonfate.control.integration.v1alpha1.IntegrationService.UpdateProxyResource.
func (c *integrationServiceClient) UpdateProxyResource(ctx context.Context, req *connect.Request[v1alpha1.UpdateProxyResourceRequest]) (*connect.Response[v1alpha1.UpdateProxyResourceResponse], error) {
	return c.updateProxyResource.CallUnary(ctx, req)
}

// DeleteProxyResource calls
// commonfate.control.integration.v1alpha1.IntegrationService.DeleteProxyResource.
func (c *integrationServiceClient) DeleteProxyResource(ctx context.Context, req *connect.Request[v1alpha1.DeleteProxyResourceRequest]) (*connect.Response[v1alpha1.DeleteProxyResourceResponse], error) {
	return c.deleteProxyResource.CallUnary(ctx, req)
}

// GetProxyResource calls
// commonfate.control.integration.v1alpha1.IntegrationService.GetProxyResource.
func (c *integrationServiceClient) GetProxyResource(ctx context.Context, req *connect.Request[v1alpha1.GetProxyResourceRequest]) (*connect.Response[v1alpha1.GetProxyResourceResponse], error) {
	return c.getProxyResource.CallUnary(ctx, req)
}

// CreateProxy calls commonfate.control.integration.v1alpha1.IntegrationService.CreateProxy.
func (c *integrationServiceClient) CreateProxy(ctx context.Context, req *connect.Request[v1alpha1.CreateProxyRequest]) (*connect.Response[v1alpha1.CreateProxyResponse], error) {
	return c.createProxy.CallUnary(ctx, req)
}

// UpdateProxy calls commonfate.control.integration.v1alpha1.IntegrationService.UpdateProxy.
func (c *integrationServiceClient) UpdateProxy(ctx context.Context, req *connect.Request[v1alpha1.UpdateProxyRequest]) (*connect.Response[v1alpha1.UpdateProxyResponse], error) {
	return c.updateProxy.CallUnary(ctx, req)
}

// DeleteProxy calls commonfate.control.integration.v1alpha1.IntegrationService.DeleteProxy.
func (c *integrationServiceClient) DeleteProxy(ctx context.Context, req *connect.Request[v1alpha1.DeleteProxyRequest]) (*connect.Response[v1alpha1.DeleteProxyResponse], error) {
	return c.deleteProxy.CallUnary(ctx, req)
}

// GetProxy calls commonfate.control.integration.v1alpha1.IntegrationService.GetProxy.
func (c *integrationServiceClient) GetProxy(ctx context.Context, req *connect.Request[v1alpha1.GetProxyRequest]) (*connect.Response[v1alpha1.GetProxyResponse], error) {
	return c.getProxy.CallUnary(ctx, req)
}

// ListProxyResources calls
// commonfate.control.integration.v1alpha1.IntegrationService.ListProxyResources.
func (c *integrationServiceClient) ListProxyResources(ctx context.Context, req *connect.Request[v1alpha1.DescribeProxyResourcesRequest]) (*connect.Response[v1alpha1.DescribeProxyResourcesResponse], error) {
	return c.listProxyResources.CallUnary(ctx, req)
}

// IntegrationServiceHandler is an implementation of the
// commonfate.control.integration.v1alpha1.IntegrationService service.
type IntegrationServiceHandler interface {
	CreateIntegration(context.Context, *connect.Request[v1alpha1.CreateIntegrationRequest]) (*connect.Response[v1alpha1.CreateIntegrationResponse], error)
	UpdateIntegration(context.Context, *connect.Request[v1alpha1.UpdateIntegrationRequest]) (*connect.Response[v1alpha1.UpdateIntegrationResponse], error)
	GetIntegration(context.Context, *connect.Request[v1alpha1.GetIntegrationRequest]) (*connect.Response[v1alpha1.GetIntegrationResponse], error)
	DeleteIntegration(context.Context, *connect.Request[v1alpha1.DeleteIntegrationRequest]) (*connect.Response[v1alpha1.DeleteIntegrationResponse], error)
	ListIntegrations(context.Context, *connect.Request[v1alpha1.ListIntegrationsRequest]) (*connect.Response[v1alpha1.ListIntegrationsResponse], error)
	// CRUD operations for proxy resource terraform provider resource
	CreateProxyResource(context.Context, *connect.Request[v1alpha1.CreateProxyResourceRequest]) (*connect.Response[v1alpha1.CreateProxyResourceResponse], error)
	UpdateProxyResource(context.Context, *connect.Request[v1alpha1.UpdateProxyResourceRequest]) (*connect.Response[v1alpha1.UpdateProxyResourceResponse], error)
	DeleteProxyResource(context.Context, *connect.Request[v1alpha1.DeleteProxyResourceRequest]) (*connect.Response[v1alpha1.DeleteProxyResourceResponse], error)
	GetProxyResource(context.Context, *connect.Request[v1alpha1.GetProxyResourceRequest]) (*connect.Response[v1alpha1.GetProxyResourceResponse], error)
	// CRUD operations for proxy terraform provider resource
	CreateProxy(context.Context, *connect.Request[v1alpha1.CreateProxyRequest]) (*connect.Response[v1alpha1.CreateProxyResponse], error)
	UpdateProxy(context.Context, *connect.Request[v1alpha1.UpdateProxyRequest]) (*connect.Response[v1alpha1.UpdateProxyResponse], error)
	DeleteProxy(context.Context, *connect.Request[v1alpha1.DeleteProxyRequest]) (*connect.Response[v1alpha1.DeleteProxyResponse], error)
	GetProxy(context.Context, *connect.Request[v1alpha1.GetProxyRequest]) (*connect.Response[v1alpha1.GetProxyResponse], error)
	// Used by the proxy to get resources
	ListProxyResources(context.Context, *connect.Request[v1alpha1.DescribeProxyResourcesRequest]) (*connect.Response[v1alpha1.DescribeProxyResourcesResponse], error)
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
	integrationServiceCreateProxyResourceHandler := connect.NewUnaryHandler(
		IntegrationServiceCreateProxyResourceProcedure,
		svc.CreateProxyResource,
		connect.WithSchema(integrationServiceCreateProxyResourceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceUpdateProxyResourceHandler := connect.NewUnaryHandler(
		IntegrationServiceUpdateProxyResourceProcedure,
		svc.UpdateProxyResource,
		connect.WithSchema(integrationServiceUpdateProxyResourceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceDeleteProxyResourceHandler := connect.NewUnaryHandler(
		IntegrationServiceDeleteProxyResourceProcedure,
		svc.DeleteProxyResource,
		connect.WithSchema(integrationServiceDeleteProxyResourceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceGetProxyResourceHandler := connect.NewUnaryHandler(
		IntegrationServiceGetProxyResourceProcedure,
		svc.GetProxyResource,
		connect.WithSchema(integrationServiceGetProxyResourceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceCreateProxyHandler := connect.NewUnaryHandler(
		IntegrationServiceCreateProxyProcedure,
		svc.CreateProxy,
		connect.WithSchema(integrationServiceCreateProxyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceUpdateProxyHandler := connect.NewUnaryHandler(
		IntegrationServiceUpdateProxyProcedure,
		svc.UpdateProxy,
		connect.WithSchema(integrationServiceUpdateProxyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceDeleteProxyHandler := connect.NewUnaryHandler(
		IntegrationServiceDeleteProxyProcedure,
		svc.DeleteProxy,
		connect.WithSchema(integrationServiceDeleteProxyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceGetProxyHandler := connect.NewUnaryHandler(
		IntegrationServiceGetProxyProcedure,
		svc.GetProxy,
		connect.WithSchema(integrationServiceGetProxyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	integrationServiceListProxyResourcesHandler := connect.NewUnaryHandler(
		IntegrationServiceListProxyResourcesProcedure,
		svc.ListProxyResources,
		connect.WithSchema(integrationServiceListProxyResourcesMethodDescriptor),
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
		case IntegrationServiceCreateProxyResourceProcedure:
			integrationServiceCreateProxyResourceHandler.ServeHTTP(w, r)
		case IntegrationServiceUpdateProxyResourceProcedure:
			integrationServiceUpdateProxyResourceHandler.ServeHTTP(w, r)
		case IntegrationServiceDeleteProxyResourceProcedure:
			integrationServiceDeleteProxyResourceHandler.ServeHTTP(w, r)
		case IntegrationServiceGetProxyResourceProcedure:
			integrationServiceGetProxyResourceHandler.ServeHTTP(w, r)
		case IntegrationServiceCreateProxyProcedure:
			integrationServiceCreateProxyHandler.ServeHTTP(w, r)
		case IntegrationServiceUpdateProxyProcedure:
			integrationServiceUpdateProxyHandler.ServeHTTP(w, r)
		case IntegrationServiceDeleteProxyProcedure:
			integrationServiceDeleteProxyHandler.ServeHTTP(w, r)
		case IntegrationServiceGetProxyProcedure:
			integrationServiceGetProxyHandler.ServeHTTP(w, r)
		case IntegrationServiceListProxyResourcesProcedure:
			integrationServiceListProxyResourcesHandler.ServeHTTP(w, r)
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

func (UnimplementedIntegrationServiceHandler) CreateProxyResource(context.Context, *connect.Request[v1alpha1.CreateProxyResourceRequest]) (*connect.Response[v1alpha1.CreateProxyResourceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.CreateProxyResource is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) UpdateProxyResource(context.Context, *connect.Request[v1alpha1.UpdateProxyResourceRequest]) (*connect.Response[v1alpha1.UpdateProxyResourceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.UpdateProxyResource is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) DeleteProxyResource(context.Context, *connect.Request[v1alpha1.DeleteProxyResourceRequest]) (*connect.Response[v1alpha1.DeleteProxyResourceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.DeleteProxyResource is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) GetProxyResource(context.Context, *connect.Request[v1alpha1.GetProxyResourceRequest]) (*connect.Response[v1alpha1.GetProxyResourceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.GetProxyResource is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) CreateProxy(context.Context, *connect.Request[v1alpha1.CreateProxyRequest]) (*connect.Response[v1alpha1.CreateProxyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.CreateProxy is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) UpdateProxy(context.Context, *connect.Request[v1alpha1.UpdateProxyRequest]) (*connect.Response[v1alpha1.UpdateProxyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.UpdateProxy is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) DeleteProxy(context.Context, *connect.Request[v1alpha1.DeleteProxyRequest]) (*connect.Response[v1alpha1.DeleteProxyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.DeleteProxy is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) GetProxy(context.Context, *connect.Request[v1alpha1.GetProxyRequest]) (*connect.Response[v1alpha1.GetProxyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.GetProxy is not implemented"))
}

func (UnimplementedIntegrationServiceHandler) ListProxyResources(context.Context, *connect.Request[v1alpha1.DescribeProxyResourcesRequest]) (*connect.Response[v1alpha1.DescribeProxyResourcesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.integration.v1alpha1.IntegrationService.ListProxyResources is not implemented"))
}
