// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/config/v1alpha1/webhook_provisioner.proto

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
	// WebhookProvisionerServiceName is the fully-qualified name of the WebhookProvisionerService
	// service.
	WebhookProvisionerServiceName = "commonfate.control.config.v1alpha1.WebhookProvisionerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// WebhookProvisionerServiceCreateWebhookProvisionerProcedure is the fully-qualified name of the
	// WebhookProvisionerService's CreateWebhookProvisioner RPC.
	WebhookProvisionerServiceCreateWebhookProvisionerProcedure = "/commonfate.control.config.v1alpha1.WebhookProvisionerService/CreateWebhookProvisioner"
	// WebhookProvisionerServiceGetWebhookProvisionerProcedure is the fully-qualified name of the
	// WebhookProvisionerService's GetWebhookProvisioner RPC.
	WebhookProvisionerServiceGetWebhookProvisionerProcedure = "/commonfate.control.config.v1alpha1.WebhookProvisionerService/GetWebhookProvisioner"
	// WebhookProvisionerServiceUpdateWebhookProvisionerProcedure is the fully-qualified name of the
	// WebhookProvisionerService's UpdateWebhookProvisioner RPC.
	WebhookProvisionerServiceUpdateWebhookProvisionerProcedure = "/commonfate.control.config.v1alpha1.WebhookProvisionerService/UpdateWebhookProvisioner"
	// WebhookProvisionerServiceDeleteWebhookProvisionerProcedure is the fully-qualified name of the
	// WebhookProvisionerService's DeleteWebhookProvisioner RPC.
	WebhookProvisionerServiceDeleteWebhookProvisionerProcedure = "/commonfate.control.config.v1alpha1.WebhookProvisionerService/DeleteWebhookProvisioner"
)

// WebhookProvisionerServiceClient is a client for the
// commonfate.control.config.v1alpha1.WebhookProvisionerService service.
type WebhookProvisionerServiceClient interface {
	CreateWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.CreateWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.CreateWebhookProvisionerResponse], error)
	GetWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.GetWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.GetWebhookProvisionerResponse], error)
	UpdateWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.UpdateWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.UpdateWebhookProvisionerResponse], error)
	DeleteWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.DeleteWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.DeleteWebhookProvisionerResponse], error)
}

// NewWebhookProvisionerServiceClient constructs a client for the
// commonfate.control.config.v1alpha1.WebhookProvisionerService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWebhookProvisionerServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) WebhookProvisionerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &webhookProvisionerServiceClient{
		createWebhookProvisioner: connect_go.NewClient[v1alpha1.CreateWebhookProvisionerRequest, v1alpha1.CreateWebhookProvisionerResponse](
			httpClient,
			baseURL+WebhookProvisionerServiceCreateWebhookProvisionerProcedure,
			opts...,
		),
		getWebhookProvisioner: connect_go.NewClient[v1alpha1.GetWebhookProvisionerRequest, v1alpha1.GetWebhookProvisionerResponse](
			httpClient,
			baseURL+WebhookProvisionerServiceGetWebhookProvisionerProcedure,
			opts...,
		),
		updateWebhookProvisioner: connect_go.NewClient[v1alpha1.UpdateWebhookProvisionerRequest, v1alpha1.UpdateWebhookProvisionerResponse](
			httpClient,
			baseURL+WebhookProvisionerServiceUpdateWebhookProvisionerProcedure,
			opts...,
		),
		deleteWebhookProvisioner: connect_go.NewClient[v1alpha1.DeleteWebhookProvisionerRequest, v1alpha1.DeleteWebhookProvisionerResponse](
			httpClient,
			baseURL+WebhookProvisionerServiceDeleteWebhookProvisionerProcedure,
			opts...,
		),
	}
}

// webhookProvisionerServiceClient implements WebhookProvisionerServiceClient.
type webhookProvisionerServiceClient struct {
	createWebhookProvisioner *connect_go.Client[v1alpha1.CreateWebhookProvisionerRequest, v1alpha1.CreateWebhookProvisionerResponse]
	getWebhookProvisioner    *connect_go.Client[v1alpha1.GetWebhookProvisionerRequest, v1alpha1.GetWebhookProvisionerResponse]
	updateWebhookProvisioner *connect_go.Client[v1alpha1.UpdateWebhookProvisionerRequest, v1alpha1.UpdateWebhookProvisionerResponse]
	deleteWebhookProvisioner *connect_go.Client[v1alpha1.DeleteWebhookProvisionerRequest, v1alpha1.DeleteWebhookProvisionerResponse]
}

// CreateWebhookProvisioner calls
// commonfate.control.config.v1alpha1.WebhookProvisionerService.CreateWebhookProvisioner.
func (c *webhookProvisionerServiceClient) CreateWebhookProvisioner(ctx context.Context, req *connect_go.Request[v1alpha1.CreateWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.CreateWebhookProvisionerResponse], error) {
	return c.createWebhookProvisioner.CallUnary(ctx, req)
}

// GetWebhookProvisioner calls
// commonfate.control.config.v1alpha1.WebhookProvisionerService.GetWebhookProvisioner.
func (c *webhookProvisionerServiceClient) GetWebhookProvisioner(ctx context.Context, req *connect_go.Request[v1alpha1.GetWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.GetWebhookProvisionerResponse], error) {
	return c.getWebhookProvisioner.CallUnary(ctx, req)
}

// UpdateWebhookProvisioner calls
// commonfate.control.config.v1alpha1.WebhookProvisionerService.UpdateWebhookProvisioner.
func (c *webhookProvisionerServiceClient) UpdateWebhookProvisioner(ctx context.Context, req *connect_go.Request[v1alpha1.UpdateWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.UpdateWebhookProvisionerResponse], error) {
	return c.updateWebhookProvisioner.CallUnary(ctx, req)
}

// DeleteWebhookProvisioner calls
// commonfate.control.config.v1alpha1.WebhookProvisionerService.DeleteWebhookProvisioner.
func (c *webhookProvisionerServiceClient) DeleteWebhookProvisioner(ctx context.Context, req *connect_go.Request[v1alpha1.DeleteWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.DeleteWebhookProvisionerResponse], error) {
	return c.deleteWebhookProvisioner.CallUnary(ctx, req)
}

// WebhookProvisionerServiceHandler is an implementation of the
// commonfate.control.config.v1alpha1.WebhookProvisionerService service.
type WebhookProvisionerServiceHandler interface {
	CreateWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.CreateWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.CreateWebhookProvisionerResponse], error)
	GetWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.GetWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.GetWebhookProvisionerResponse], error)
	UpdateWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.UpdateWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.UpdateWebhookProvisionerResponse], error)
	DeleteWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.DeleteWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.DeleteWebhookProvisionerResponse], error)
}

// NewWebhookProvisionerServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWebhookProvisionerServiceHandler(svc WebhookProvisionerServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	webhookProvisionerServiceCreateWebhookProvisionerHandler := connect_go.NewUnaryHandler(
		WebhookProvisionerServiceCreateWebhookProvisionerProcedure,
		svc.CreateWebhookProvisioner,
		opts...,
	)
	webhookProvisionerServiceGetWebhookProvisionerHandler := connect_go.NewUnaryHandler(
		WebhookProvisionerServiceGetWebhookProvisionerProcedure,
		svc.GetWebhookProvisioner,
		opts...,
	)
	webhookProvisionerServiceUpdateWebhookProvisionerHandler := connect_go.NewUnaryHandler(
		WebhookProvisionerServiceUpdateWebhookProvisionerProcedure,
		svc.UpdateWebhookProvisioner,
		opts...,
	)
	webhookProvisionerServiceDeleteWebhookProvisionerHandler := connect_go.NewUnaryHandler(
		WebhookProvisionerServiceDeleteWebhookProvisionerProcedure,
		svc.DeleteWebhookProvisioner,
		opts...,
	)
	return "/commonfate.control.config.v1alpha1.WebhookProvisionerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case WebhookProvisionerServiceCreateWebhookProvisionerProcedure:
			webhookProvisionerServiceCreateWebhookProvisionerHandler.ServeHTTP(w, r)
		case WebhookProvisionerServiceGetWebhookProvisionerProcedure:
			webhookProvisionerServiceGetWebhookProvisionerHandler.ServeHTTP(w, r)
		case WebhookProvisionerServiceUpdateWebhookProvisionerProcedure:
			webhookProvisionerServiceUpdateWebhookProvisionerHandler.ServeHTTP(w, r)
		case WebhookProvisionerServiceDeleteWebhookProvisionerProcedure:
			webhookProvisionerServiceDeleteWebhookProvisionerHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedWebhookProvisionerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWebhookProvisionerServiceHandler struct{}

func (UnimplementedWebhookProvisionerServiceHandler) CreateWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.CreateWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.CreateWebhookProvisionerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.WebhookProvisionerService.CreateWebhookProvisioner is not implemented"))
}

func (UnimplementedWebhookProvisionerServiceHandler) GetWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.GetWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.GetWebhookProvisionerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.WebhookProvisionerService.GetWebhookProvisioner is not implemented"))
}

func (UnimplementedWebhookProvisionerServiceHandler) UpdateWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.UpdateWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.UpdateWebhookProvisionerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.WebhookProvisionerService.UpdateWebhookProvisioner is not implemented"))
}

func (UnimplementedWebhookProvisionerServiceHandler) DeleteWebhookProvisioner(context.Context, *connect_go.Request[v1alpha1.DeleteWebhookProvisionerRequest]) (*connect_go.Response[v1alpha1.DeleteWebhookProvisionerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.WebhookProvisionerService.DeleteWebhookProvisioner is not implemented"))
}
