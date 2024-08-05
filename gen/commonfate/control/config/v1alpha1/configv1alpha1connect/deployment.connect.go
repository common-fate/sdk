// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/config/v1alpha1/deployment.proto

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
	// DeploymentServiceName is the fully-qualified name of the DeploymentService service.
	DeploymentServiceName = "commonfate.control.config.v1alpha1.DeploymentService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DeploymentServiceGetDeploymentConfigProcedure is the fully-qualified name of the
	// DeploymentService's GetDeploymentConfig RPC.
	DeploymentServiceGetDeploymentConfigProcedure = "/commonfate.control.config.v1alpha1.DeploymentService/GetDeploymentConfig"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	deploymentServiceServiceDescriptor                   = v1alpha1.File_commonfate_control_config_v1alpha1_deployment_proto.Services().ByName("DeploymentService")
	deploymentServiceGetDeploymentConfigMethodDescriptor = deploymentServiceServiceDescriptor.Methods().ByName("GetDeploymentConfig")
)

// DeploymentServiceClient is a client for the commonfate.control.config.v1alpha1.DeploymentService
// service.
type DeploymentServiceClient interface {
	GetDeploymentConfig(context.Context, *connect.Request[v1alpha1.GetDeploymentConfigRequest]) (*connect.Response[v1alpha1.GetDeploymentConfigResponse], error)
}

// NewDeploymentServiceClient constructs a client for the
// commonfate.control.config.v1alpha1.DeploymentService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDeploymentServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DeploymentServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &deploymentServiceClient{
		getDeploymentConfig: connect.NewClient[v1alpha1.GetDeploymentConfigRequest, v1alpha1.GetDeploymentConfigResponse](
			httpClient,
			baseURL+DeploymentServiceGetDeploymentConfigProcedure,
			connect.WithSchema(deploymentServiceGetDeploymentConfigMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// deploymentServiceClient implements DeploymentServiceClient.
type deploymentServiceClient struct {
	getDeploymentConfig *connect.Client[v1alpha1.GetDeploymentConfigRequest, v1alpha1.GetDeploymentConfigResponse]
}

// GetDeploymentConfig calls
// commonfate.control.config.v1alpha1.DeploymentService.GetDeploymentConfig.
func (c *deploymentServiceClient) GetDeploymentConfig(ctx context.Context, req *connect.Request[v1alpha1.GetDeploymentConfigRequest]) (*connect.Response[v1alpha1.GetDeploymentConfigResponse], error) {
	return c.getDeploymentConfig.CallUnary(ctx, req)
}

// DeploymentServiceHandler is an implementation of the
// commonfate.control.config.v1alpha1.DeploymentService service.
type DeploymentServiceHandler interface {
	GetDeploymentConfig(context.Context, *connect.Request[v1alpha1.GetDeploymentConfigRequest]) (*connect.Response[v1alpha1.GetDeploymentConfigResponse], error)
}

// NewDeploymentServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDeploymentServiceHandler(svc DeploymentServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	deploymentServiceGetDeploymentConfigHandler := connect.NewUnaryHandler(
		DeploymentServiceGetDeploymentConfigProcedure,
		svc.GetDeploymentConfig,
		connect.WithSchema(deploymentServiceGetDeploymentConfigMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.control.config.v1alpha1.DeploymentService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DeploymentServiceGetDeploymentConfigProcedure:
			deploymentServiceGetDeploymentConfigHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDeploymentServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDeploymentServiceHandler struct{}

func (UnimplementedDeploymentServiceHandler) GetDeploymentConfig(context.Context, *connect.Request[v1alpha1.GetDeploymentConfigRequest]) (*connect.Response[v1alpha1.GetDeploymentConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.DeploymentService.GetDeploymentConfig is not implemented"))
}