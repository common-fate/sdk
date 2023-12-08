// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/config/v1alpha1/access_workflow.proto

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
	// AccessWorkflowServiceName is the fully-qualified name of the AccessWorkflowService service.
	AccessWorkflowServiceName = "commonfate.control.config.v1alpha1.AccessWorkflowService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AccessWorkflowServiceCreateAccessWorkflowProcedure is the fully-qualified name of the
	// AccessWorkflowService's CreateAccessWorkflow RPC.
	AccessWorkflowServiceCreateAccessWorkflowProcedure = "/commonfate.control.config.v1alpha1.AccessWorkflowService/CreateAccessWorkflow"
	// AccessWorkflowServiceReadAccessWorkflowProcedure is the fully-qualified name of the
	// AccessWorkflowService's ReadAccessWorkflow RPC.
	AccessWorkflowServiceReadAccessWorkflowProcedure = "/commonfate.control.config.v1alpha1.AccessWorkflowService/ReadAccessWorkflow"
	// AccessWorkflowServiceUpdateAccessWorkflowProcedure is the fully-qualified name of the
	// AccessWorkflowService's UpdateAccessWorkflow RPC.
	AccessWorkflowServiceUpdateAccessWorkflowProcedure = "/commonfate.control.config.v1alpha1.AccessWorkflowService/UpdateAccessWorkflow"
	// AccessWorkflowServiceDeleteAccessWorkflowProcedure is the fully-qualified name of the
	// AccessWorkflowService's DeleteAccessWorkflow RPC.
	AccessWorkflowServiceDeleteAccessWorkflowProcedure = "/commonfate.control.config.v1alpha1.AccessWorkflowService/DeleteAccessWorkflow"
)

// AccessWorkflowServiceClient is a client for the
// commonfate.control.config.v1alpha1.AccessWorkflowService service.
type AccessWorkflowServiceClient interface {
	CreateAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.CreateAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.CreateAccessWorkflowResponse], error)
	ReadAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.ReadAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.ReadAccessWorkflowResponse], error)
	UpdateAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.UpdateAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.UpdateAccessWorkflowResponse], error)
	DeleteAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.DeleteAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.DeleteAccessWorkflowResponse], error)
}

// NewAccessWorkflowServiceClient constructs a client for the
// commonfate.control.config.v1alpha1.AccessWorkflowService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAccessWorkflowServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AccessWorkflowServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &accessWorkflowServiceClient{
		createAccessWorkflow: connect_go.NewClient[v1alpha1.CreateAccessWorkflowRequest, v1alpha1.CreateAccessWorkflowResponse](
			httpClient,
			baseURL+AccessWorkflowServiceCreateAccessWorkflowProcedure,
			opts...,
		),
		readAccessWorkflow: connect_go.NewClient[v1alpha1.ReadAccessWorkflowRequest, v1alpha1.ReadAccessWorkflowResponse](
			httpClient,
			baseURL+AccessWorkflowServiceReadAccessWorkflowProcedure,
			opts...,
		),
		updateAccessWorkflow: connect_go.NewClient[v1alpha1.UpdateAccessWorkflowRequest, v1alpha1.UpdateAccessWorkflowResponse](
			httpClient,
			baseURL+AccessWorkflowServiceUpdateAccessWorkflowProcedure,
			opts...,
		),
		deleteAccessWorkflow: connect_go.NewClient[v1alpha1.DeleteAccessWorkflowRequest, v1alpha1.DeleteAccessWorkflowResponse](
			httpClient,
			baseURL+AccessWorkflowServiceDeleteAccessWorkflowProcedure,
			opts...,
		),
	}
}

// accessWorkflowServiceClient implements AccessWorkflowServiceClient.
type accessWorkflowServiceClient struct {
	createAccessWorkflow *connect_go.Client[v1alpha1.CreateAccessWorkflowRequest, v1alpha1.CreateAccessWorkflowResponse]
	readAccessWorkflow   *connect_go.Client[v1alpha1.ReadAccessWorkflowRequest, v1alpha1.ReadAccessWorkflowResponse]
	updateAccessWorkflow *connect_go.Client[v1alpha1.UpdateAccessWorkflowRequest, v1alpha1.UpdateAccessWorkflowResponse]
	deleteAccessWorkflow *connect_go.Client[v1alpha1.DeleteAccessWorkflowRequest, v1alpha1.DeleteAccessWorkflowResponse]
}

// CreateAccessWorkflow calls
// commonfate.control.config.v1alpha1.AccessWorkflowService.CreateAccessWorkflow.
func (c *accessWorkflowServiceClient) CreateAccessWorkflow(ctx context.Context, req *connect_go.Request[v1alpha1.CreateAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.CreateAccessWorkflowResponse], error) {
	return c.createAccessWorkflow.CallUnary(ctx, req)
}

// ReadAccessWorkflow calls
// commonfate.control.config.v1alpha1.AccessWorkflowService.ReadAccessWorkflow.
func (c *accessWorkflowServiceClient) ReadAccessWorkflow(ctx context.Context, req *connect_go.Request[v1alpha1.ReadAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.ReadAccessWorkflowResponse], error) {
	return c.readAccessWorkflow.CallUnary(ctx, req)
}

// UpdateAccessWorkflow calls
// commonfate.control.config.v1alpha1.AccessWorkflowService.UpdateAccessWorkflow.
func (c *accessWorkflowServiceClient) UpdateAccessWorkflow(ctx context.Context, req *connect_go.Request[v1alpha1.UpdateAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.UpdateAccessWorkflowResponse], error) {
	return c.updateAccessWorkflow.CallUnary(ctx, req)
}

// DeleteAccessWorkflow calls
// commonfate.control.config.v1alpha1.AccessWorkflowService.DeleteAccessWorkflow.
func (c *accessWorkflowServiceClient) DeleteAccessWorkflow(ctx context.Context, req *connect_go.Request[v1alpha1.DeleteAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.DeleteAccessWorkflowResponse], error) {
	return c.deleteAccessWorkflow.CallUnary(ctx, req)
}

// AccessWorkflowServiceHandler is an implementation of the
// commonfate.control.config.v1alpha1.AccessWorkflowService service.
type AccessWorkflowServiceHandler interface {
	CreateAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.CreateAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.CreateAccessWorkflowResponse], error)
	ReadAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.ReadAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.ReadAccessWorkflowResponse], error)
	UpdateAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.UpdateAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.UpdateAccessWorkflowResponse], error)
	DeleteAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.DeleteAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.DeleteAccessWorkflowResponse], error)
}

// NewAccessWorkflowServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAccessWorkflowServiceHandler(svc AccessWorkflowServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	accessWorkflowServiceCreateAccessWorkflowHandler := connect_go.NewUnaryHandler(
		AccessWorkflowServiceCreateAccessWorkflowProcedure,
		svc.CreateAccessWorkflow,
		opts...,
	)
	accessWorkflowServiceReadAccessWorkflowHandler := connect_go.NewUnaryHandler(
		AccessWorkflowServiceReadAccessWorkflowProcedure,
		svc.ReadAccessWorkflow,
		opts...,
	)
	accessWorkflowServiceUpdateAccessWorkflowHandler := connect_go.NewUnaryHandler(
		AccessWorkflowServiceUpdateAccessWorkflowProcedure,
		svc.UpdateAccessWorkflow,
		opts...,
	)
	accessWorkflowServiceDeleteAccessWorkflowHandler := connect_go.NewUnaryHandler(
		AccessWorkflowServiceDeleteAccessWorkflowProcedure,
		svc.DeleteAccessWorkflow,
		opts...,
	)
	return "/commonfate.control.config.v1alpha1.AccessWorkflowService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AccessWorkflowServiceCreateAccessWorkflowProcedure:
			accessWorkflowServiceCreateAccessWorkflowHandler.ServeHTTP(w, r)
		case AccessWorkflowServiceReadAccessWorkflowProcedure:
			accessWorkflowServiceReadAccessWorkflowHandler.ServeHTTP(w, r)
		case AccessWorkflowServiceUpdateAccessWorkflowProcedure:
			accessWorkflowServiceUpdateAccessWorkflowHandler.ServeHTTP(w, r)
		case AccessWorkflowServiceDeleteAccessWorkflowProcedure:
			accessWorkflowServiceDeleteAccessWorkflowHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAccessWorkflowServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAccessWorkflowServiceHandler struct{}

func (UnimplementedAccessWorkflowServiceHandler) CreateAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.CreateAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.CreateAccessWorkflowResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.AccessWorkflowService.CreateAccessWorkflow is not implemented"))
}

func (UnimplementedAccessWorkflowServiceHandler) ReadAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.ReadAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.ReadAccessWorkflowResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.AccessWorkflowService.ReadAccessWorkflow is not implemented"))
}

func (UnimplementedAccessWorkflowServiceHandler) UpdateAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.UpdateAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.UpdateAccessWorkflowResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.AccessWorkflowService.UpdateAccessWorkflow is not implemented"))
}

func (UnimplementedAccessWorkflowServiceHandler) DeleteAccessWorkflow(context.Context, *connect_go.Request[v1alpha1.DeleteAccessWorkflowRequest]) (*connect_go.Response[v1alpha1.DeleteAccessWorkflowResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.AccessWorkflowService.DeleteAccessWorkflow is not implemented"))
}