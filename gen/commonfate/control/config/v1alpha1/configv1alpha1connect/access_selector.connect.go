// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/config/v1alpha1/access_selector.proto

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
	// AccessSelectorServiceName is the fully-qualified name of the AccessSelectorService service.
	AccessSelectorServiceName = "commonfate.control.config.v1alpha1.AccessSelectorService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AccessSelectorServiceCreateAccessSelectorProcedure is the fully-qualified name of the
	// AccessSelectorService's CreateAccessSelector RPC.
	AccessSelectorServiceCreateAccessSelectorProcedure = "/commonfate.control.config.v1alpha1.AccessSelectorService/CreateAccessSelector"
	// AccessSelectorServiceGetAccessSelectorProcedure is the fully-qualified name of the
	// AccessSelectorService's GetAccessSelector RPC.
	AccessSelectorServiceGetAccessSelectorProcedure = "/commonfate.control.config.v1alpha1.AccessSelectorService/GetAccessSelector"
	// AccessSelectorServiceUpdateAccessSelectorProcedure is the fully-qualified name of the
	// AccessSelectorService's UpdateAccessSelector RPC.
	AccessSelectorServiceUpdateAccessSelectorProcedure = "/commonfate.control.config.v1alpha1.AccessSelectorService/UpdateAccessSelector"
	// AccessSelectorServiceDeleteAccessSelectorProcedure is the fully-qualified name of the
	// AccessSelectorService's DeleteAccessSelector RPC.
	AccessSelectorServiceDeleteAccessSelectorProcedure = "/commonfate.control.config.v1alpha1.AccessSelectorService/DeleteAccessSelector"
)

// AccessSelectorServiceClient is a client for the
// commonfate.control.config.v1alpha1.AccessSelectorService service.
type AccessSelectorServiceClient interface {
	CreateAccessSelector(context.Context, *connect_go.Request[v1alpha1.CreateAccessSelectorRequest]) (*connect_go.Response[v1alpha1.CreateAccessSelectorResponse], error)
	GetAccessSelector(context.Context, *connect_go.Request[v1alpha1.GetAccessSelectorRequest]) (*connect_go.Response[v1alpha1.GetAccessSelectorResponse], error)
	UpdateAccessSelector(context.Context, *connect_go.Request[v1alpha1.UpdateAccessSelectorRequest]) (*connect_go.Response[v1alpha1.UpdateAccessSelectorResponse], error)
	DeleteAccessSelector(context.Context, *connect_go.Request[v1alpha1.DeleteAccessSelectorRequest]) (*connect_go.Response[v1alpha1.DeleteAccessSelectorResponse], error)
}

// NewAccessSelectorServiceClient constructs a client for the
// commonfate.control.config.v1alpha1.AccessSelectorService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAccessSelectorServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AccessSelectorServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &accessSelectorServiceClient{
		createAccessSelector: connect_go.NewClient[v1alpha1.CreateAccessSelectorRequest, v1alpha1.CreateAccessSelectorResponse](
			httpClient,
			baseURL+AccessSelectorServiceCreateAccessSelectorProcedure,
			opts...,
		),
		getAccessSelector: connect_go.NewClient[v1alpha1.GetAccessSelectorRequest, v1alpha1.GetAccessSelectorResponse](
			httpClient,
			baseURL+AccessSelectorServiceGetAccessSelectorProcedure,
			opts...,
		),
		updateAccessSelector: connect_go.NewClient[v1alpha1.UpdateAccessSelectorRequest, v1alpha1.UpdateAccessSelectorResponse](
			httpClient,
			baseURL+AccessSelectorServiceUpdateAccessSelectorProcedure,
			opts...,
		),
		deleteAccessSelector: connect_go.NewClient[v1alpha1.DeleteAccessSelectorRequest, v1alpha1.DeleteAccessSelectorResponse](
			httpClient,
			baseURL+AccessSelectorServiceDeleteAccessSelectorProcedure,
			opts...,
		),
	}
}

// accessSelectorServiceClient implements AccessSelectorServiceClient.
type accessSelectorServiceClient struct {
	createAccessSelector *connect_go.Client[v1alpha1.CreateAccessSelectorRequest, v1alpha1.CreateAccessSelectorResponse]
	getAccessSelector    *connect_go.Client[v1alpha1.GetAccessSelectorRequest, v1alpha1.GetAccessSelectorResponse]
	updateAccessSelector *connect_go.Client[v1alpha1.UpdateAccessSelectorRequest, v1alpha1.UpdateAccessSelectorResponse]
	deleteAccessSelector *connect_go.Client[v1alpha1.DeleteAccessSelectorRequest, v1alpha1.DeleteAccessSelectorResponse]
}

// CreateAccessSelector calls
// commonfate.control.config.v1alpha1.AccessSelectorService.CreateAccessSelector.
func (c *accessSelectorServiceClient) CreateAccessSelector(ctx context.Context, req *connect_go.Request[v1alpha1.CreateAccessSelectorRequest]) (*connect_go.Response[v1alpha1.CreateAccessSelectorResponse], error) {
	return c.createAccessSelector.CallUnary(ctx, req)
}

// GetAccessSelector calls
// commonfate.control.config.v1alpha1.AccessSelectorService.GetAccessSelector.
func (c *accessSelectorServiceClient) GetAccessSelector(ctx context.Context, req *connect_go.Request[v1alpha1.GetAccessSelectorRequest]) (*connect_go.Response[v1alpha1.GetAccessSelectorResponse], error) {
	return c.getAccessSelector.CallUnary(ctx, req)
}

// UpdateAccessSelector calls
// commonfate.control.config.v1alpha1.AccessSelectorService.UpdateAccessSelector.
func (c *accessSelectorServiceClient) UpdateAccessSelector(ctx context.Context, req *connect_go.Request[v1alpha1.UpdateAccessSelectorRequest]) (*connect_go.Response[v1alpha1.UpdateAccessSelectorResponse], error) {
	return c.updateAccessSelector.CallUnary(ctx, req)
}

// DeleteAccessSelector calls
// commonfate.control.config.v1alpha1.AccessSelectorService.DeleteAccessSelector.
func (c *accessSelectorServiceClient) DeleteAccessSelector(ctx context.Context, req *connect_go.Request[v1alpha1.DeleteAccessSelectorRequest]) (*connect_go.Response[v1alpha1.DeleteAccessSelectorResponse], error) {
	return c.deleteAccessSelector.CallUnary(ctx, req)
}

// AccessSelectorServiceHandler is an implementation of the
// commonfate.control.config.v1alpha1.AccessSelectorService service.
type AccessSelectorServiceHandler interface {
	CreateAccessSelector(context.Context, *connect_go.Request[v1alpha1.CreateAccessSelectorRequest]) (*connect_go.Response[v1alpha1.CreateAccessSelectorResponse], error)
	GetAccessSelector(context.Context, *connect_go.Request[v1alpha1.GetAccessSelectorRequest]) (*connect_go.Response[v1alpha1.GetAccessSelectorResponse], error)
	UpdateAccessSelector(context.Context, *connect_go.Request[v1alpha1.UpdateAccessSelectorRequest]) (*connect_go.Response[v1alpha1.UpdateAccessSelectorResponse], error)
	DeleteAccessSelector(context.Context, *connect_go.Request[v1alpha1.DeleteAccessSelectorRequest]) (*connect_go.Response[v1alpha1.DeleteAccessSelectorResponse], error)
}

// NewAccessSelectorServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAccessSelectorServiceHandler(svc AccessSelectorServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	accessSelectorServiceCreateAccessSelectorHandler := connect_go.NewUnaryHandler(
		AccessSelectorServiceCreateAccessSelectorProcedure,
		svc.CreateAccessSelector,
		opts...,
	)
	accessSelectorServiceGetAccessSelectorHandler := connect_go.NewUnaryHandler(
		AccessSelectorServiceGetAccessSelectorProcedure,
		svc.GetAccessSelector,
		opts...,
	)
	accessSelectorServiceUpdateAccessSelectorHandler := connect_go.NewUnaryHandler(
		AccessSelectorServiceUpdateAccessSelectorProcedure,
		svc.UpdateAccessSelector,
		opts...,
	)
	accessSelectorServiceDeleteAccessSelectorHandler := connect_go.NewUnaryHandler(
		AccessSelectorServiceDeleteAccessSelectorProcedure,
		svc.DeleteAccessSelector,
		opts...,
	)
	return "/commonfate.control.config.v1alpha1.AccessSelectorService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AccessSelectorServiceCreateAccessSelectorProcedure:
			accessSelectorServiceCreateAccessSelectorHandler.ServeHTTP(w, r)
		case AccessSelectorServiceGetAccessSelectorProcedure:
			accessSelectorServiceGetAccessSelectorHandler.ServeHTTP(w, r)
		case AccessSelectorServiceUpdateAccessSelectorProcedure:
			accessSelectorServiceUpdateAccessSelectorHandler.ServeHTTP(w, r)
		case AccessSelectorServiceDeleteAccessSelectorProcedure:
			accessSelectorServiceDeleteAccessSelectorHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAccessSelectorServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAccessSelectorServiceHandler struct{}

func (UnimplementedAccessSelectorServiceHandler) CreateAccessSelector(context.Context, *connect_go.Request[v1alpha1.CreateAccessSelectorRequest]) (*connect_go.Response[v1alpha1.CreateAccessSelectorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.AccessSelectorService.CreateAccessSelector is not implemented"))
}

func (UnimplementedAccessSelectorServiceHandler) GetAccessSelector(context.Context, *connect_go.Request[v1alpha1.GetAccessSelectorRequest]) (*connect_go.Response[v1alpha1.GetAccessSelectorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.AccessSelectorService.GetAccessSelector is not implemented"))
}

func (UnimplementedAccessSelectorServiceHandler) UpdateAccessSelector(context.Context, *connect_go.Request[v1alpha1.UpdateAccessSelectorRequest]) (*connect_go.Response[v1alpha1.UpdateAccessSelectorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.AccessSelectorService.UpdateAccessSelector is not implemented"))
}

func (UnimplementedAccessSelectorServiceHandler) DeleteAccessSelector(context.Context, *connect_go.Request[v1alpha1.DeleteAccessSelectorRequest]) (*connect_go.Response[v1alpha1.DeleteAccessSelectorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.AccessSelectorService.DeleteAccessSelector is not implemented"))
}
