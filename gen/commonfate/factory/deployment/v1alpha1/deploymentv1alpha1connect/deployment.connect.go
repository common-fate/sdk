// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/factory/deployment/v1alpha1/deployment.proto

package deploymentv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/factory/deployment/v1alpha1"
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
	// DeploymentServiceName is the fully-qualified name of the DeploymentService service.
	DeploymentServiceName = "commonfate.factory.deployment.v1alpha1.DeploymentService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DeploymentServiceGetDeploymentProcedure is the fully-qualified name of the DeploymentService's
	// GetDeployment RPC.
	DeploymentServiceGetDeploymentProcedure = "/commonfate.factory.deployment.v1alpha1.DeploymentService/GetDeployment"
	// DeploymentServiceRegisterDeploymentNameserversProcedure is the fully-qualified name of the
	// DeploymentService's RegisterDeploymentNameservers RPC.
	DeploymentServiceRegisterDeploymentNameserversProcedure = "/commonfate.factory.deployment.v1alpha1.DeploymentService/RegisterDeploymentNameservers"
	// DeploymentServiceGetDeploymentNameserversProcedure is the fully-qualified name of the
	// DeploymentService's GetDeploymentNameservers RPC.
	DeploymentServiceGetDeploymentNameserversProcedure = "/commonfate.factory.deployment.v1alpha1.DeploymentService/GetDeploymentNameservers"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	deploymentServiceServiceDescriptor                             = v1alpha1.File_commonfate_factory_deployment_v1alpha1_deployment_proto.Services().ByName("DeploymentService")
	deploymentServiceGetDeploymentMethodDescriptor                 = deploymentServiceServiceDescriptor.Methods().ByName("GetDeployment")
	deploymentServiceRegisterDeploymentNameserversMethodDescriptor = deploymentServiceServiceDescriptor.Methods().ByName("RegisterDeploymentNameservers")
	deploymentServiceGetDeploymentNameserversMethodDescriptor      = deploymentServiceServiceDescriptor.Methods().ByName("GetDeploymentNameservers")
)

// DeploymentServiceClient is a client for the
// commonfate.factory.deployment.v1alpha1.DeploymentService service.
type DeploymentServiceClient interface {
	// Get information about a deployment.
	GetDeployment(context.Context, *connect.Request[v1alpha1.GetDeploymentRequest]) (*connect.Response[v1alpha1.GetDeploymentResponse], error)
	// Register nameservers for the default deployment domain.
	RegisterDeploymentNameservers(context.Context, *connect.Request[v1alpha1.RegisterDeploymentNameserversRequest]) (*connect.Response[v1alpha1.RegisterDeploymentNameserversResponse], error)
	// Get nameservers for the default deployment domain.
	GetDeploymentNameservers(context.Context, *connect.Request[v1alpha1.GetDeploymentNameserversRequest]) (*connect.Response[v1alpha1.GetDeploymentNameserversResponse], error)
}

// NewDeploymentServiceClient constructs a client for the
// commonfate.factory.deployment.v1alpha1.DeploymentService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDeploymentServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DeploymentServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &deploymentServiceClient{
		getDeployment: connect.NewClient[v1alpha1.GetDeploymentRequest, v1alpha1.GetDeploymentResponse](
			httpClient,
			baseURL+DeploymentServiceGetDeploymentProcedure,
			connect.WithSchema(deploymentServiceGetDeploymentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		registerDeploymentNameservers: connect.NewClient[v1alpha1.RegisterDeploymentNameserversRequest, v1alpha1.RegisterDeploymentNameserversResponse](
			httpClient,
			baseURL+DeploymentServiceRegisterDeploymentNameserversProcedure,
			connect.WithSchema(deploymentServiceRegisterDeploymentNameserversMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getDeploymentNameservers: connect.NewClient[v1alpha1.GetDeploymentNameserversRequest, v1alpha1.GetDeploymentNameserversResponse](
			httpClient,
			baseURL+DeploymentServiceGetDeploymentNameserversProcedure,
			connect.WithSchema(deploymentServiceGetDeploymentNameserversMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// deploymentServiceClient implements DeploymentServiceClient.
type deploymentServiceClient struct {
	getDeployment                 *connect.Client[v1alpha1.GetDeploymentRequest, v1alpha1.GetDeploymentResponse]
	registerDeploymentNameservers *connect.Client[v1alpha1.RegisterDeploymentNameserversRequest, v1alpha1.RegisterDeploymentNameserversResponse]
	getDeploymentNameservers      *connect.Client[v1alpha1.GetDeploymentNameserversRequest, v1alpha1.GetDeploymentNameserversResponse]
}

// GetDeployment calls commonfate.factory.deployment.v1alpha1.DeploymentService.GetDeployment.
func (c *deploymentServiceClient) GetDeployment(ctx context.Context, req *connect.Request[v1alpha1.GetDeploymentRequest]) (*connect.Response[v1alpha1.GetDeploymentResponse], error) {
	return c.getDeployment.CallUnary(ctx, req)
}

// RegisterDeploymentNameservers calls
// commonfate.factory.deployment.v1alpha1.DeploymentService.RegisterDeploymentNameservers.
func (c *deploymentServiceClient) RegisterDeploymentNameservers(ctx context.Context, req *connect.Request[v1alpha1.RegisterDeploymentNameserversRequest]) (*connect.Response[v1alpha1.RegisterDeploymentNameserversResponse], error) {
	return c.registerDeploymentNameservers.CallUnary(ctx, req)
}

// GetDeploymentNameservers calls
// commonfate.factory.deployment.v1alpha1.DeploymentService.GetDeploymentNameservers.
func (c *deploymentServiceClient) GetDeploymentNameservers(ctx context.Context, req *connect.Request[v1alpha1.GetDeploymentNameserversRequest]) (*connect.Response[v1alpha1.GetDeploymentNameserversResponse], error) {
	return c.getDeploymentNameservers.CallUnary(ctx, req)
}

// DeploymentServiceHandler is an implementation of the
// commonfate.factory.deployment.v1alpha1.DeploymentService service.
type DeploymentServiceHandler interface {
	// Get information about a deployment.
	GetDeployment(context.Context, *connect.Request[v1alpha1.GetDeploymentRequest]) (*connect.Response[v1alpha1.GetDeploymentResponse], error)
	// Register nameservers for the default deployment domain.
	RegisterDeploymentNameservers(context.Context, *connect.Request[v1alpha1.RegisterDeploymentNameserversRequest]) (*connect.Response[v1alpha1.RegisterDeploymentNameserversResponse], error)
	// Get nameservers for the default deployment domain.
	GetDeploymentNameservers(context.Context, *connect.Request[v1alpha1.GetDeploymentNameserversRequest]) (*connect.Response[v1alpha1.GetDeploymentNameserversResponse], error)
}

// NewDeploymentServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDeploymentServiceHandler(svc DeploymentServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	deploymentServiceGetDeploymentHandler := connect.NewUnaryHandler(
		DeploymentServiceGetDeploymentProcedure,
		svc.GetDeployment,
		connect.WithSchema(deploymentServiceGetDeploymentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	deploymentServiceRegisterDeploymentNameserversHandler := connect.NewUnaryHandler(
		DeploymentServiceRegisterDeploymentNameserversProcedure,
		svc.RegisterDeploymentNameservers,
		connect.WithSchema(deploymentServiceRegisterDeploymentNameserversMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	deploymentServiceGetDeploymentNameserversHandler := connect.NewUnaryHandler(
		DeploymentServiceGetDeploymentNameserversProcedure,
		svc.GetDeploymentNameservers,
		connect.WithSchema(deploymentServiceGetDeploymentNameserversMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.factory.deployment.v1alpha1.DeploymentService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DeploymentServiceGetDeploymentProcedure:
			deploymentServiceGetDeploymentHandler.ServeHTTP(w, r)
		case DeploymentServiceRegisterDeploymentNameserversProcedure:
			deploymentServiceRegisterDeploymentNameserversHandler.ServeHTTP(w, r)
		case DeploymentServiceGetDeploymentNameserversProcedure:
			deploymentServiceGetDeploymentNameserversHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDeploymentServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDeploymentServiceHandler struct{}

func (UnimplementedDeploymentServiceHandler) GetDeployment(context.Context, *connect.Request[v1alpha1.GetDeploymentRequest]) (*connect.Response[v1alpha1.GetDeploymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.factory.deployment.v1alpha1.DeploymentService.GetDeployment is not implemented"))
}

func (UnimplementedDeploymentServiceHandler) RegisterDeploymentNameservers(context.Context, *connect.Request[v1alpha1.RegisterDeploymentNameserversRequest]) (*connect.Response[v1alpha1.RegisterDeploymentNameserversResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.factory.deployment.v1alpha1.DeploymentService.RegisterDeploymentNameservers is not implemented"))
}

func (UnimplementedDeploymentServiceHandler) GetDeploymentNameservers(context.Context, *connect.Request[v1alpha1.GetDeploymentNameserversRequest]) (*connect.Response[v1alpha1.GetDeploymentNameserversResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.factory.deployment.v1alpha1.DeploymentService.GetDeploymentNameservers is not implemented"))
}
