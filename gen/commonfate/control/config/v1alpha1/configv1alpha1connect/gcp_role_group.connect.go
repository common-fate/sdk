// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/config/v1alpha1/gcp_role_group.proto

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
	// GCPRoleGroupServiceName is the fully-qualified name of the GCPRoleGroupService service.
	GCPRoleGroupServiceName = "commonfate.control.config.v1alpha1.GCPRoleGroupService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GCPRoleGroupServiceCreateGCPRoleGroupProcedure is the fully-qualified name of the
	// GCPRoleGroupService's CreateGCPRoleGroup RPC.
	GCPRoleGroupServiceCreateGCPRoleGroupProcedure = "/commonfate.control.config.v1alpha1.GCPRoleGroupService/CreateGCPRoleGroup"
	// GCPRoleGroupServiceGetGCPRoleGroupProcedure is the fully-qualified name of the
	// GCPRoleGroupService's GetGCPRoleGroup RPC.
	GCPRoleGroupServiceGetGCPRoleGroupProcedure = "/commonfate.control.config.v1alpha1.GCPRoleGroupService/GetGCPRoleGroup"
	// GCPRoleGroupServiceUpdateGCPRoleGroupProcedure is the fully-qualified name of the
	// GCPRoleGroupService's UpdateGCPRoleGroup RPC.
	GCPRoleGroupServiceUpdateGCPRoleGroupProcedure = "/commonfate.control.config.v1alpha1.GCPRoleGroupService/UpdateGCPRoleGroup"
	// GCPRoleGroupServiceDeleteGCPRoleGroupProcedure is the fully-qualified name of the
	// GCPRoleGroupService's DeleteGCPRoleGroup RPC.
	GCPRoleGroupServiceDeleteGCPRoleGroupProcedure = "/commonfate.control.config.v1alpha1.GCPRoleGroupService/DeleteGCPRoleGroup"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	gCPRoleGroupServiceServiceDescriptor                  = v1alpha1.File_commonfate_control_config_v1alpha1_gcp_role_group_proto.Services().ByName("GCPRoleGroupService")
	gCPRoleGroupServiceCreateGCPRoleGroupMethodDescriptor = gCPRoleGroupServiceServiceDescriptor.Methods().ByName("CreateGCPRoleGroup")
	gCPRoleGroupServiceGetGCPRoleGroupMethodDescriptor    = gCPRoleGroupServiceServiceDescriptor.Methods().ByName("GetGCPRoleGroup")
	gCPRoleGroupServiceUpdateGCPRoleGroupMethodDescriptor = gCPRoleGroupServiceServiceDescriptor.Methods().ByName("UpdateGCPRoleGroup")
	gCPRoleGroupServiceDeleteGCPRoleGroupMethodDescriptor = gCPRoleGroupServiceServiceDescriptor.Methods().ByName("DeleteGCPRoleGroup")
)

// GCPRoleGroupServiceClient is a client for the
// commonfate.control.config.v1alpha1.GCPRoleGroupService service.
type GCPRoleGroupServiceClient interface {
	CreateGCPRoleGroup(context.Context, *connect.Request[v1alpha1.CreateGCPRoleGroupRequest]) (*connect.Response[v1alpha1.CreateGCPRoleGroupResponse], error)
	GetGCPRoleGroup(context.Context, *connect.Request[v1alpha1.GetGCPRoleGroupRequest]) (*connect.Response[v1alpha1.GetGCPRoleGroupResponse], error)
	UpdateGCPRoleGroup(context.Context, *connect.Request[v1alpha1.UpdateGCPRoleGroupRequest]) (*connect.Response[v1alpha1.UpdateGCPRoleGroupResponse], error)
	DeleteGCPRoleGroup(context.Context, *connect.Request[v1alpha1.DeleteGCPRoleGroupRequest]) (*connect.Response[v1alpha1.DeleteGCPRoleGroupResponse], error)
}

// NewGCPRoleGroupServiceClient constructs a client for the
// commonfate.control.config.v1alpha1.GCPRoleGroupService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGCPRoleGroupServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GCPRoleGroupServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &gCPRoleGroupServiceClient{
		createGCPRoleGroup: connect.NewClient[v1alpha1.CreateGCPRoleGroupRequest, v1alpha1.CreateGCPRoleGroupResponse](
			httpClient,
			baseURL+GCPRoleGroupServiceCreateGCPRoleGroupProcedure,
			connect.WithSchema(gCPRoleGroupServiceCreateGCPRoleGroupMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getGCPRoleGroup: connect.NewClient[v1alpha1.GetGCPRoleGroupRequest, v1alpha1.GetGCPRoleGroupResponse](
			httpClient,
			baseURL+GCPRoleGroupServiceGetGCPRoleGroupProcedure,
			connect.WithSchema(gCPRoleGroupServiceGetGCPRoleGroupMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateGCPRoleGroup: connect.NewClient[v1alpha1.UpdateGCPRoleGroupRequest, v1alpha1.UpdateGCPRoleGroupResponse](
			httpClient,
			baseURL+GCPRoleGroupServiceUpdateGCPRoleGroupProcedure,
			connect.WithSchema(gCPRoleGroupServiceUpdateGCPRoleGroupMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteGCPRoleGroup: connect.NewClient[v1alpha1.DeleteGCPRoleGroupRequest, v1alpha1.DeleteGCPRoleGroupResponse](
			httpClient,
			baseURL+GCPRoleGroupServiceDeleteGCPRoleGroupProcedure,
			connect.WithSchema(gCPRoleGroupServiceDeleteGCPRoleGroupMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// gCPRoleGroupServiceClient implements GCPRoleGroupServiceClient.
type gCPRoleGroupServiceClient struct {
	createGCPRoleGroup *connect.Client[v1alpha1.CreateGCPRoleGroupRequest, v1alpha1.CreateGCPRoleGroupResponse]
	getGCPRoleGroup    *connect.Client[v1alpha1.GetGCPRoleGroupRequest, v1alpha1.GetGCPRoleGroupResponse]
	updateGCPRoleGroup *connect.Client[v1alpha1.UpdateGCPRoleGroupRequest, v1alpha1.UpdateGCPRoleGroupResponse]
	deleteGCPRoleGroup *connect.Client[v1alpha1.DeleteGCPRoleGroupRequest, v1alpha1.DeleteGCPRoleGroupResponse]
}

// CreateGCPRoleGroup calls
// commonfate.control.config.v1alpha1.GCPRoleGroupService.CreateGCPRoleGroup.
func (c *gCPRoleGroupServiceClient) CreateGCPRoleGroup(ctx context.Context, req *connect.Request[v1alpha1.CreateGCPRoleGroupRequest]) (*connect.Response[v1alpha1.CreateGCPRoleGroupResponse], error) {
	return c.createGCPRoleGroup.CallUnary(ctx, req)
}

// GetGCPRoleGroup calls commonfate.control.config.v1alpha1.GCPRoleGroupService.GetGCPRoleGroup.
func (c *gCPRoleGroupServiceClient) GetGCPRoleGroup(ctx context.Context, req *connect.Request[v1alpha1.GetGCPRoleGroupRequest]) (*connect.Response[v1alpha1.GetGCPRoleGroupResponse], error) {
	return c.getGCPRoleGroup.CallUnary(ctx, req)
}

// UpdateGCPRoleGroup calls
// commonfate.control.config.v1alpha1.GCPRoleGroupService.UpdateGCPRoleGroup.
func (c *gCPRoleGroupServiceClient) UpdateGCPRoleGroup(ctx context.Context, req *connect.Request[v1alpha1.UpdateGCPRoleGroupRequest]) (*connect.Response[v1alpha1.UpdateGCPRoleGroupResponse], error) {
	return c.updateGCPRoleGroup.CallUnary(ctx, req)
}

// DeleteGCPRoleGroup calls
// commonfate.control.config.v1alpha1.GCPRoleGroupService.DeleteGCPRoleGroup.
func (c *gCPRoleGroupServiceClient) DeleteGCPRoleGroup(ctx context.Context, req *connect.Request[v1alpha1.DeleteGCPRoleGroupRequest]) (*connect.Response[v1alpha1.DeleteGCPRoleGroupResponse], error) {
	return c.deleteGCPRoleGroup.CallUnary(ctx, req)
}

// GCPRoleGroupServiceHandler is an implementation of the
// commonfate.control.config.v1alpha1.GCPRoleGroupService service.
type GCPRoleGroupServiceHandler interface {
	CreateGCPRoleGroup(context.Context, *connect.Request[v1alpha1.CreateGCPRoleGroupRequest]) (*connect.Response[v1alpha1.CreateGCPRoleGroupResponse], error)
	GetGCPRoleGroup(context.Context, *connect.Request[v1alpha1.GetGCPRoleGroupRequest]) (*connect.Response[v1alpha1.GetGCPRoleGroupResponse], error)
	UpdateGCPRoleGroup(context.Context, *connect.Request[v1alpha1.UpdateGCPRoleGroupRequest]) (*connect.Response[v1alpha1.UpdateGCPRoleGroupResponse], error)
	DeleteGCPRoleGroup(context.Context, *connect.Request[v1alpha1.DeleteGCPRoleGroupRequest]) (*connect.Response[v1alpha1.DeleteGCPRoleGroupResponse], error)
}

// NewGCPRoleGroupServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGCPRoleGroupServiceHandler(svc GCPRoleGroupServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	gCPRoleGroupServiceCreateGCPRoleGroupHandler := connect.NewUnaryHandler(
		GCPRoleGroupServiceCreateGCPRoleGroupProcedure,
		svc.CreateGCPRoleGroup,
		connect.WithSchema(gCPRoleGroupServiceCreateGCPRoleGroupMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	gCPRoleGroupServiceGetGCPRoleGroupHandler := connect.NewUnaryHandler(
		GCPRoleGroupServiceGetGCPRoleGroupProcedure,
		svc.GetGCPRoleGroup,
		connect.WithSchema(gCPRoleGroupServiceGetGCPRoleGroupMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	gCPRoleGroupServiceUpdateGCPRoleGroupHandler := connect.NewUnaryHandler(
		GCPRoleGroupServiceUpdateGCPRoleGroupProcedure,
		svc.UpdateGCPRoleGroup,
		connect.WithSchema(gCPRoleGroupServiceUpdateGCPRoleGroupMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	gCPRoleGroupServiceDeleteGCPRoleGroupHandler := connect.NewUnaryHandler(
		GCPRoleGroupServiceDeleteGCPRoleGroupProcedure,
		svc.DeleteGCPRoleGroup,
		connect.WithSchema(gCPRoleGroupServiceDeleteGCPRoleGroupMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.control.config.v1alpha1.GCPRoleGroupService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GCPRoleGroupServiceCreateGCPRoleGroupProcedure:
			gCPRoleGroupServiceCreateGCPRoleGroupHandler.ServeHTTP(w, r)
		case GCPRoleGroupServiceGetGCPRoleGroupProcedure:
			gCPRoleGroupServiceGetGCPRoleGroupHandler.ServeHTTP(w, r)
		case GCPRoleGroupServiceUpdateGCPRoleGroupProcedure:
			gCPRoleGroupServiceUpdateGCPRoleGroupHandler.ServeHTTP(w, r)
		case GCPRoleGroupServiceDeleteGCPRoleGroupProcedure:
			gCPRoleGroupServiceDeleteGCPRoleGroupHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGCPRoleGroupServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGCPRoleGroupServiceHandler struct{}

func (UnimplementedGCPRoleGroupServiceHandler) CreateGCPRoleGroup(context.Context, *connect.Request[v1alpha1.CreateGCPRoleGroupRequest]) (*connect.Response[v1alpha1.CreateGCPRoleGroupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.GCPRoleGroupService.CreateGCPRoleGroup is not implemented"))
}

func (UnimplementedGCPRoleGroupServiceHandler) GetGCPRoleGroup(context.Context, *connect.Request[v1alpha1.GetGCPRoleGroupRequest]) (*connect.Response[v1alpha1.GetGCPRoleGroupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.GCPRoleGroupService.GetGCPRoleGroup is not implemented"))
}

func (UnimplementedGCPRoleGroupServiceHandler) UpdateGCPRoleGroup(context.Context, *connect.Request[v1alpha1.UpdateGCPRoleGroupRequest]) (*connect.Response[v1alpha1.UpdateGCPRoleGroupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.GCPRoleGroupService.UpdateGCPRoleGroup is not implemented"))
}

func (UnimplementedGCPRoleGroupServiceHandler) DeleteGCPRoleGroup(context.Context, *connect.Request[v1alpha1.DeleteGCPRoleGroupRequest]) (*connect.Response[v1alpha1.DeleteGCPRoleGroupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.GCPRoleGroupService.DeleteGCPRoleGroup is not implemented"))
}