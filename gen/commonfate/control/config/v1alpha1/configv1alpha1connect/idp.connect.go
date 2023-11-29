// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/config/v1alpha1/idp.proto

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
	// IDPServiceName is the fully-qualified name of the IDPService service.
	IDPServiceName = "commonfate.control.config.v1alpha1.IDPService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IDPServiceCreateIDPProcedure is the fully-qualified name of the IDPService's CreateIDP RPC.
	IDPServiceCreateIDPProcedure = "/commonfate.control.config.v1alpha1.IDPService/CreateIDP"
	// IDPServiceReadIDPProcedure is the fully-qualified name of the IDPService's ReadIDP RPC.
	IDPServiceReadIDPProcedure = "/commonfate.control.config.v1alpha1.IDPService/ReadIDP"
	// IDPServiceUpdateIDPProcedure is the fully-qualified name of the IDPService's UpdateIDP RPC.
	IDPServiceUpdateIDPProcedure = "/commonfate.control.config.v1alpha1.IDPService/UpdateIDP"
	// IDPServiceDeleteIDPProcedure is the fully-qualified name of the IDPService's DeleteIDP RPC.
	IDPServiceDeleteIDPProcedure = "/commonfate.control.config.v1alpha1.IDPService/DeleteIDP"
)

// IDPServiceClient is a client for the commonfate.control.config.v1alpha1.IDPService service.
type IDPServiceClient interface {
	CreateIDP(context.Context, *connect_go.Request[v1alpha1.CreateIDPRequest]) (*connect_go.Response[v1alpha1.CreateIDPResponse], error)
	ReadIDP(context.Context, *connect_go.Request[v1alpha1.ReadIDPRequest]) (*connect_go.Response[v1alpha1.ReadIDPResponse], error)
	UpdateIDP(context.Context, *connect_go.Request[v1alpha1.UpdateIDPRequest]) (*connect_go.Response[v1alpha1.UpdateIDPResponse], error)
	DeleteIDP(context.Context, *connect_go.Request[v1alpha1.DeleteIDPRequest]) (*connect_go.Response[v1alpha1.DeleteIDPResponse], error)
}

// NewIDPServiceClient constructs a client for the commonfate.control.config.v1alpha1.IDPService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIDPServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IDPServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &iDPServiceClient{
		createIDP: connect_go.NewClient[v1alpha1.CreateIDPRequest, v1alpha1.CreateIDPResponse](
			httpClient,
			baseURL+IDPServiceCreateIDPProcedure,
			opts...,
		),
		readIDP: connect_go.NewClient[v1alpha1.ReadIDPRequest, v1alpha1.ReadIDPResponse](
			httpClient,
			baseURL+IDPServiceReadIDPProcedure,
			opts...,
		),
		updateIDP: connect_go.NewClient[v1alpha1.UpdateIDPRequest, v1alpha1.UpdateIDPResponse](
			httpClient,
			baseURL+IDPServiceUpdateIDPProcedure,
			opts...,
		),
		deleteIDP: connect_go.NewClient[v1alpha1.DeleteIDPRequest, v1alpha1.DeleteIDPResponse](
			httpClient,
			baseURL+IDPServiceDeleteIDPProcedure,
			opts...,
		),
	}
}

// iDPServiceClient implements IDPServiceClient.
type iDPServiceClient struct {
	createIDP *connect_go.Client[v1alpha1.CreateIDPRequest, v1alpha1.CreateIDPResponse]
	readIDP   *connect_go.Client[v1alpha1.ReadIDPRequest, v1alpha1.ReadIDPResponse]
	updateIDP *connect_go.Client[v1alpha1.UpdateIDPRequest, v1alpha1.UpdateIDPResponse]
	deleteIDP *connect_go.Client[v1alpha1.DeleteIDPRequest, v1alpha1.DeleteIDPResponse]
}

// CreateIDP calls commonfate.control.config.v1alpha1.IDPService.CreateIDP.
func (c *iDPServiceClient) CreateIDP(ctx context.Context, req *connect_go.Request[v1alpha1.CreateIDPRequest]) (*connect_go.Response[v1alpha1.CreateIDPResponse], error) {
	return c.createIDP.CallUnary(ctx, req)
}

// ReadIDP calls commonfate.control.config.v1alpha1.IDPService.ReadIDP.
func (c *iDPServiceClient) ReadIDP(ctx context.Context, req *connect_go.Request[v1alpha1.ReadIDPRequest]) (*connect_go.Response[v1alpha1.ReadIDPResponse], error) {
	return c.readIDP.CallUnary(ctx, req)
}

// UpdateIDP calls commonfate.control.config.v1alpha1.IDPService.UpdateIDP.
func (c *iDPServiceClient) UpdateIDP(ctx context.Context, req *connect_go.Request[v1alpha1.UpdateIDPRequest]) (*connect_go.Response[v1alpha1.UpdateIDPResponse], error) {
	return c.updateIDP.CallUnary(ctx, req)
}

// DeleteIDP calls commonfate.control.config.v1alpha1.IDPService.DeleteIDP.
func (c *iDPServiceClient) DeleteIDP(ctx context.Context, req *connect_go.Request[v1alpha1.DeleteIDPRequest]) (*connect_go.Response[v1alpha1.DeleteIDPResponse], error) {
	return c.deleteIDP.CallUnary(ctx, req)
}

// IDPServiceHandler is an implementation of the commonfate.control.config.v1alpha1.IDPService
// service.
type IDPServiceHandler interface {
	CreateIDP(context.Context, *connect_go.Request[v1alpha1.CreateIDPRequest]) (*connect_go.Response[v1alpha1.CreateIDPResponse], error)
	ReadIDP(context.Context, *connect_go.Request[v1alpha1.ReadIDPRequest]) (*connect_go.Response[v1alpha1.ReadIDPResponse], error)
	UpdateIDP(context.Context, *connect_go.Request[v1alpha1.UpdateIDPRequest]) (*connect_go.Response[v1alpha1.UpdateIDPResponse], error)
	DeleteIDP(context.Context, *connect_go.Request[v1alpha1.DeleteIDPRequest]) (*connect_go.Response[v1alpha1.DeleteIDPResponse], error)
}

// NewIDPServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIDPServiceHandler(svc IDPServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	iDPServiceCreateIDPHandler := connect_go.NewUnaryHandler(
		IDPServiceCreateIDPProcedure,
		svc.CreateIDP,
		opts...,
	)
	iDPServiceReadIDPHandler := connect_go.NewUnaryHandler(
		IDPServiceReadIDPProcedure,
		svc.ReadIDP,
		opts...,
	)
	iDPServiceUpdateIDPHandler := connect_go.NewUnaryHandler(
		IDPServiceUpdateIDPProcedure,
		svc.UpdateIDP,
		opts...,
	)
	iDPServiceDeleteIDPHandler := connect_go.NewUnaryHandler(
		IDPServiceDeleteIDPProcedure,
		svc.DeleteIDP,
		opts...,
	)
	return "/commonfate.control.config.v1alpha1.IDPService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IDPServiceCreateIDPProcedure:
			iDPServiceCreateIDPHandler.ServeHTTP(w, r)
		case IDPServiceReadIDPProcedure:
			iDPServiceReadIDPHandler.ServeHTTP(w, r)
		case IDPServiceUpdateIDPProcedure:
			iDPServiceUpdateIDPHandler.ServeHTTP(w, r)
		case IDPServiceDeleteIDPProcedure:
			iDPServiceDeleteIDPHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIDPServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIDPServiceHandler struct{}

func (UnimplementedIDPServiceHandler) CreateIDP(context.Context, *connect_go.Request[v1alpha1.CreateIDPRequest]) (*connect_go.Response[v1alpha1.CreateIDPResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.IDPService.CreateIDP is not implemented"))
}

func (UnimplementedIDPServiceHandler) ReadIDP(context.Context, *connect_go.Request[v1alpha1.ReadIDPRequest]) (*connect_go.Response[v1alpha1.ReadIDPResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.IDPService.ReadIDP is not implemented"))
}

func (UnimplementedIDPServiceHandler) UpdateIDP(context.Context, *connect_go.Request[v1alpha1.UpdateIDPRequest]) (*connect_go.Response[v1alpha1.UpdateIDPResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.IDPService.UpdateIDP is not implemented"))
}

func (UnimplementedIDPServiceHandler) DeleteIDP(context.Context, *connect_go.Request[v1alpha1.DeleteIDPRequest]) (*connect_go.Response[v1alpha1.DeleteIDPResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.IDPService.DeleteIDP is not implemented"))
}
