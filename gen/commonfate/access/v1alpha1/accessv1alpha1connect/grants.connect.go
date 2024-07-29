// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/access/v1alpha1/grants.proto

package accessv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
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
	// GrantsServiceName is the fully-qualified name of the GrantsService service.
	GrantsServiceName = "commonfate.access.v1alpha1.GrantsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GrantsServiceQueryGrantsProcedure is the fully-qualified name of the GrantsService's QueryGrants
	// RPC.
	GrantsServiceQueryGrantsProcedure = "/commonfate.access.v1alpha1.GrantsService/QueryGrants"
	// GrantsServiceGetGrantProcedure is the fully-qualified name of the GrantsService's GetGrant RPC.
	GrantsServiceGetGrantProcedure = "/commonfate.access.v1alpha1.GrantsService/GetGrant"
	// GrantsServiceGetGrantOutputProcedure is the fully-qualified name of the GrantsService's
	// GetGrantOutput RPC.
	GrantsServiceGetGrantOutputProcedure = "/commonfate.access.v1alpha1.GrantsService/GetGrantOutput"
	// GrantsServiceQueryGrantChildrenProcedure is the fully-qualified name of the GrantsService's
	// QueryGrantChildren RPC.
	GrantsServiceQueryGrantChildrenProcedure = "/commonfate.access.v1alpha1.GrantsService/QueryGrantChildren"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	grantsServiceServiceDescriptor                  = v1alpha1.File_commonfate_access_v1alpha1_grants_proto.Services().ByName("GrantsService")
	grantsServiceQueryGrantsMethodDescriptor        = grantsServiceServiceDescriptor.Methods().ByName("QueryGrants")
	grantsServiceGetGrantMethodDescriptor           = grantsServiceServiceDescriptor.Methods().ByName("GetGrant")
	grantsServiceGetGrantOutputMethodDescriptor     = grantsServiceServiceDescriptor.Methods().ByName("GetGrantOutput")
	grantsServiceQueryGrantChildrenMethodDescriptor = grantsServiceServiceDescriptor.Methods().ByName("QueryGrantChildren")
)

// GrantsServiceClient is a client for the commonfate.access.v1alpha1.GrantsService service.
type GrantsServiceClient interface {
	QueryGrants(context.Context, *connect.Request[v1alpha1.QueryGrantsRequest]) (*connect.Response[v1alpha1.QueryGrantsResponse], error)
	GetGrant(context.Context, *connect.Request[v1alpha1.GetGrantRequest]) (*connect.Response[v1alpha1.GetGrantResponse], error)
	GetGrantOutput(context.Context, *connect.Request[v1alpha1.GetGrantOutputRequest]) (*connect.Response[v1alpha1.GetGrantOutputResponse], error)
	QueryGrantChildren(context.Context, *connect.Request[v1alpha1.QueryGrantChildrenRequest]) (*connect.Response[v1alpha1.QueryGrantChildrenResponse], error)
}

// NewGrantsServiceClient constructs a client for the commonfate.access.v1alpha1.GrantsService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGrantsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GrantsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &grantsServiceClient{
		queryGrants: connect.NewClient[v1alpha1.QueryGrantsRequest, v1alpha1.QueryGrantsResponse](
			httpClient,
			baseURL+GrantsServiceQueryGrantsProcedure,
			connect.WithSchema(grantsServiceQueryGrantsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getGrant: connect.NewClient[v1alpha1.GetGrantRequest, v1alpha1.GetGrantResponse](
			httpClient,
			baseURL+GrantsServiceGetGrantProcedure,
			connect.WithSchema(grantsServiceGetGrantMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getGrantOutput: connect.NewClient[v1alpha1.GetGrantOutputRequest, v1alpha1.GetGrantOutputResponse](
			httpClient,
			baseURL+GrantsServiceGetGrantOutputProcedure,
			connect.WithSchema(grantsServiceGetGrantOutputMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		queryGrantChildren: connect.NewClient[v1alpha1.QueryGrantChildrenRequest, v1alpha1.QueryGrantChildrenResponse](
			httpClient,
			baseURL+GrantsServiceQueryGrantChildrenProcedure,
			connect.WithSchema(grantsServiceQueryGrantChildrenMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// grantsServiceClient implements GrantsServiceClient.
type grantsServiceClient struct {
	queryGrants        *connect.Client[v1alpha1.QueryGrantsRequest, v1alpha1.QueryGrantsResponse]
	getGrant           *connect.Client[v1alpha1.GetGrantRequest, v1alpha1.GetGrantResponse]
	getGrantOutput     *connect.Client[v1alpha1.GetGrantOutputRequest, v1alpha1.GetGrantOutputResponse]
	queryGrantChildren *connect.Client[v1alpha1.QueryGrantChildrenRequest, v1alpha1.QueryGrantChildrenResponse]
}

// QueryGrants calls commonfate.access.v1alpha1.GrantsService.QueryGrants.
func (c *grantsServiceClient) QueryGrants(ctx context.Context, req *connect.Request[v1alpha1.QueryGrantsRequest]) (*connect.Response[v1alpha1.QueryGrantsResponse], error) {
	return c.queryGrants.CallUnary(ctx, req)
}

// GetGrant calls commonfate.access.v1alpha1.GrantsService.GetGrant.
func (c *grantsServiceClient) GetGrant(ctx context.Context, req *connect.Request[v1alpha1.GetGrantRequest]) (*connect.Response[v1alpha1.GetGrantResponse], error) {
	return c.getGrant.CallUnary(ctx, req)
}

// GetGrantOutput calls commonfate.access.v1alpha1.GrantsService.GetGrantOutput.
func (c *grantsServiceClient) GetGrantOutput(ctx context.Context, req *connect.Request[v1alpha1.GetGrantOutputRequest]) (*connect.Response[v1alpha1.GetGrantOutputResponse], error) {
	return c.getGrantOutput.CallUnary(ctx, req)
}

// QueryGrantChildren calls commonfate.access.v1alpha1.GrantsService.QueryGrantChildren.
func (c *grantsServiceClient) QueryGrantChildren(ctx context.Context, req *connect.Request[v1alpha1.QueryGrantChildrenRequest]) (*connect.Response[v1alpha1.QueryGrantChildrenResponse], error) {
	return c.queryGrantChildren.CallUnary(ctx, req)
}

// GrantsServiceHandler is an implementation of the commonfate.access.v1alpha1.GrantsService
// service.
type GrantsServiceHandler interface {
	QueryGrants(context.Context, *connect.Request[v1alpha1.QueryGrantsRequest]) (*connect.Response[v1alpha1.QueryGrantsResponse], error)
	GetGrant(context.Context, *connect.Request[v1alpha1.GetGrantRequest]) (*connect.Response[v1alpha1.GetGrantResponse], error)
	GetGrantOutput(context.Context, *connect.Request[v1alpha1.GetGrantOutputRequest]) (*connect.Response[v1alpha1.GetGrantOutputResponse], error)
	QueryGrantChildren(context.Context, *connect.Request[v1alpha1.QueryGrantChildrenRequest]) (*connect.Response[v1alpha1.QueryGrantChildrenResponse], error)
}

// NewGrantsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGrantsServiceHandler(svc GrantsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	grantsServiceQueryGrantsHandler := connect.NewUnaryHandler(
		GrantsServiceQueryGrantsProcedure,
		svc.QueryGrants,
		connect.WithSchema(grantsServiceQueryGrantsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	grantsServiceGetGrantHandler := connect.NewUnaryHandler(
		GrantsServiceGetGrantProcedure,
		svc.GetGrant,
		connect.WithSchema(grantsServiceGetGrantMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	grantsServiceGetGrantOutputHandler := connect.NewUnaryHandler(
		GrantsServiceGetGrantOutputProcedure,
		svc.GetGrantOutput,
		connect.WithSchema(grantsServiceGetGrantOutputMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	grantsServiceQueryGrantChildrenHandler := connect.NewUnaryHandler(
		GrantsServiceQueryGrantChildrenProcedure,
		svc.QueryGrantChildren,
		connect.WithSchema(grantsServiceQueryGrantChildrenMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.access.v1alpha1.GrantsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GrantsServiceQueryGrantsProcedure:
			grantsServiceQueryGrantsHandler.ServeHTTP(w, r)
		case GrantsServiceGetGrantProcedure:
			grantsServiceGetGrantHandler.ServeHTTP(w, r)
		case GrantsServiceGetGrantOutputProcedure:
			grantsServiceGetGrantOutputHandler.ServeHTTP(w, r)
		case GrantsServiceQueryGrantChildrenProcedure:
			grantsServiceQueryGrantChildrenHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGrantsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGrantsServiceHandler struct{}

func (UnimplementedGrantsServiceHandler) QueryGrants(context.Context, *connect.Request[v1alpha1.QueryGrantsRequest]) (*connect.Response[v1alpha1.QueryGrantsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.GrantsService.QueryGrants is not implemented"))
}

func (UnimplementedGrantsServiceHandler) GetGrant(context.Context, *connect.Request[v1alpha1.GetGrantRequest]) (*connect.Response[v1alpha1.GetGrantResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.GrantsService.GetGrant is not implemented"))
}

func (UnimplementedGrantsServiceHandler) GetGrantOutput(context.Context, *connect.Request[v1alpha1.GetGrantOutputRequest]) (*connect.Response[v1alpha1.GetGrantOutputResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.GrantsService.GetGrantOutput is not implemented"))
}

func (UnimplementedGrantsServiceHandler) QueryGrantChildren(context.Context, *connect.Request[v1alpha1.QueryGrantChildrenRequest]) (*connect.Response[v1alpha1.QueryGrantChildrenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.GrantsService.QueryGrantChildren is not implemented"))
}
