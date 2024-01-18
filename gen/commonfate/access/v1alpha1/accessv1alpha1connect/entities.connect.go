// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/access/v1alpha1/entities.proto

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
	// EntityServiceName is the fully-qualified name of the EntityService service.
	EntityServiceName = "commonfate.access.v1alpha1.EntityService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EntityServiceQueryDescendentsProcedure is the fully-qualified name of the EntityService's
	// QueryDescendents RPC.
	EntityServiceQueryDescendentsProcedure = "/commonfate.access.v1alpha1.EntityService/QueryDescendents"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	entityServiceServiceDescriptor                = v1alpha1.File_commonfate_access_v1alpha1_entities_proto.Services().ByName("EntityService")
	entityServiceQueryDescendentsMethodDescriptor = entityServiceServiceDescriptor.Methods().ByName("QueryDescendents")
)

// EntityServiceClient is a client for the commonfate.access.v1alpha1.EntityService service.
type EntityServiceClient interface {
	QueryDescendents(context.Context, *connect.Request[v1alpha1.QueryDescendentsRequest]) (*connect.Response[v1alpha1.QueryDescendentsResponse], error)
}

// NewEntityServiceClient constructs a client for the commonfate.access.v1alpha1.EntityService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEntityServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) EntityServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &entityServiceClient{
		queryDescendents: connect.NewClient[v1alpha1.QueryDescendentsRequest, v1alpha1.QueryDescendentsResponse](
			httpClient,
			baseURL+EntityServiceQueryDescendentsProcedure,
			connect.WithSchema(entityServiceQueryDescendentsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// entityServiceClient implements EntityServiceClient.
type entityServiceClient struct {
	queryDescendents *connect.Client[v1alpha1.QueryDescendentsRequest, v1alpha1.QueryDescendentsResponse]
}

// QueryDescendents calls commonfate.access.v1alpha1.EntityService.QueryDescendents.
func (c *entityServiceClient) QueryDescendents(ctx context.Context, req *connect.Request[v1alpha1.QueryDescendentsRequest]) (*connect.Response[v1alpha1.QueryDescendentsResponse], error) {
	return c.queryDescendents.CallUnary(ctx, req)
}

// EntityServiceHandler is an implementation of the commonfate.access.v1alpha1.EntityService
// service.
type EntityServiceHandler interface {
	QueryDescendents(context.Context, *connect.Request[v1alpha1.QueryDescendentsRequest]) (*connect.Response[v1alpha1.QueryDescendentsResponse], error)
}

// NewEntityServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEntityServiceHandler(svc EntityServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	entityServiceQueryDescendentsHandler := connect.NewUnaryHandler(
		EntityServiceQueryDescendentsProcedure,
		svc.QueryDescendents,
		connect.WithSchema(entityServiceQueryDescendentsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.access.v1alpha1.EntityService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case EntityServiceQueryDescendentsProcedure:
			entityServiceQueryDescendentsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedEntityServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEntityServiceHandler struct{}

func (UnimplementedEntityServiceHandler) QueryDescendents(context.Context, *connect.Request[v1alpha1.QueryDescendentsRequest]) (*connect.Response[v1alpha1.QueryDescendentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.EntityService.QueryDescendents is not implemented"))
}
