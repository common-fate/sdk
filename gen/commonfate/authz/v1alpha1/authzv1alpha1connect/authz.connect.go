// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/authz/v1alpha1/authz.proto

package authzv1alpha1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
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
	// AuthzServiceName is the fully-qualified name of the AuthzService service.
	AuthzServiceName = "commonfate.authz.v1alpha1.AuthzService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthzServiceBatchPutPolicySetProcedure is the fully-qualified name of the AuthzService's
	// BatchPutPolicySet RPC.
	AuthzServiceBatchPutPolicySetProcedure = "/commonfate.authz.v1alpha1.AuthzService/BatchPutPolicySet"
	// AuthzServiceBatchAuthorizeProcedure is the fully-qualified name of the AuthzService's
	// BatchAuthorize RPC.
	AuthzServiceBatchAuthorizeProcedure = "/commonfate.authz.v1alpha1.AuthzService/BatchAuthorize"
	// AuthzServiceListPolicySetsProcedure is the fully-qualified name of the AuthzService's
	// ListPolicySets RPC.
	AuthzServiceListPolicySetsProcedure = "/commonfate.authz.v1alpha1.AuthzService/ListPolicySets"
)

// AuthzServiceClient is a client for the commonfate.authz.v1alpha1.AuthzService service.
type AuthzServiceClient interface {
	// adds Cedar policies for a particular policy store
	BatchPutPolicySet(context.Context, *connect_go.Request[v1alpha1.BatchPutPolicySetRequest]) (*connect_go.Response[v1alpha1.BatchPutPolicySetResponse], error)
	// run multiple authorization evaluations and returns allow + deny for each.
	BatchAuthorize(context.Context, *connect_go.Request[v1alpha1.BatchAuthorizeRequest]) (*connect_go.Response[v1alpha1.BatchAuthorizeResponse], error)
	ListPolicySets(context.Context, *connect_go.Request[v1alpha1.ListPolicySetsRequest]) (*connect_go.Response[v1alpha1.ListPolicySetsResponse], error)
}

// NewAuthzServiceClient constructs a client for the commonfate.authz.v1alpha1.AuthzService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthzServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AuthzServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authzServiceClient{
		batchPutPolicySet: connect_go.NewClient[v1alpha1.BatchPutPolicySetRequest, v1alpha1.BatchPutPolicySetResponse](
			httpClient,
			baseURL+AuthzServiceBatchPutPolicySetProcedure,
			opts...,
		),
		batchAuthorize: connect_go.NewClient[v1alpha1.BatchAuthorizeRequest, v1alpha1.BatchAuthorizeResponse](
			httpClient,
			baseURL+AuthzServiceBatchAuthorizeProcedure,
			opts...,
		),
		listPolicySets: connect_go.NewClient[v1alpha1.ListPolicySetsRequest, v1alpha1.ListPolicySetsResponse](
			httpClient,
			baseURL+AuthzServiceListPolicySetsProcedure,
			opts...,
		),
	}
}

// authzServiceClient implements AuthzServiceClient.
type authzServiceClient struct {
	batchPutPolicySet *connect_go.Client[v1alpha1.BatchPutPolicySetRequest, v1alpha1.BatchPutPolicySetResponse]
	batchAuthorize    *connect_go.Client[v1alpha1.BatchAuthorizeRequest, v1alpha1.BatchAuthorizeResponse]
	listPolicySets    *connect_go.Client[v1alpha1.ListPolicySetsRequest, v1alpha1.ListPolicySetsResponse]
}

// BatchPutPolicySet calls commonfate.authz.v1alpha1.AuthzService.BatchPutPolicySet.
func (c *authzServiceClient) BatchPutPolicySet(ctx context.Context, req *connect_go.Request[v1alpha1.BatchPutPolicySetRequest]) (*connect_go.Response[v1alpha1.BatchPutPolicySetResponse], error) {
	return c.batchPutPolicySet.CallUnary(ctx, req)
}

// BatchAuthorize calls commonfate.authz.v1alpha1.AuthzService.BatchAuthorize.
func (c *authzServiceClient) BatchAuthorize(ctx context.Context, req *connect_go.Request[v1alpha1.BatchAuthorizeRequest]) (*connect_go.Response[v1alpha1.BatchAuthorizeResponse], error) {
	return c.batchAuthorize.CallUnary(ctx, req)
}

// ListPolicySets calls commonfate.authz.v1alpha1.AuthzService.ListPolicySets.
func (c *authzServiceClient) ListPolicySets(ctx context.Context, req *connect_go.Request[v1alpha1.ListPolicySetsRequest]) (*connect_go.Response[v1alpha1.ListPolicySetsResponse], error) {
	return c.listPolicySets.CallUnary(ctx, req)
}

// AuthzServiceHandler is an implementation of the commonfate.authz.v1alpha1.AuthzService service.
type AuthzServiceHandler interface {
	// adds Cedar policies for a particular policy store
	BatchPutPolicySet(context.Context, *connect_go.Request[v1alpha1.BatchPutPolicySetRequest]) (*connect_go.Response[v1alpha1.BatchPutPolicySetResponse], error)
	// run multiple authorization evaluations and returns allow + deny for each.
	BatchAuthorize(context.Context, *connect_go.Request[v1alpha1.BatchAuthorizeRequest]) (*connect_go.Response[v1alpha1.BatchAuthorizeResponse], error)
	ListPolicySets(context.Context, *connect_go.Request[v1alpha1.ListPolicySetsRequest]) (*connect_go.Response[v1alpha1.ListPolicySetsResponse], error)
}

// NewAuthzServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthzServiceHandler(svc AuthzServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	authzServiceBatchPutPolicySetHandler := connect_go.NewUnaryHandler(
		AuthzServiceBatchPutPolicySetProcedure,
		svc.BatchPutPolicySet,
		opts...,
	)
	authzServiceBatchAuthorizeHandler := connect_go.NewUnaryHandler(
		AuthzServiceBatchAuthorizeProcedure,
		svc.BatchAuthorize,
		opts...,
	)
	authzServiceListPolicySetsHandler := connect_go.NewUnaryHandler(
		AuthzServiceListPolicySetsProcedure,
		svc.ListPolicySets,
		opts...,
	)
	return "/commonfate.authz.v1alpha1.AuthzService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthzServiceBatchPutPolicySetProcedure:
			authzServiceBatchPutPolicySetHandler.ServeHTTP(w, r)
		case AuthzServiceBatchAuthorizeProcedure:
			authzServiceBatchAuthorizeHandler.ServeHTTP(w, r)
		case AuthzServiceListPolicySetsProcedure:
			authzServiceListPolicySetsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthzServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthzServiceHandler struct{}

func (UnimplementedAuthzServiceHandler) BatchPutPolicySet(context.Context, *connect_go.Request[v1alpha1.BatchPutPolicySetRequest]) (*connect_go.Response[v1alpha1.BatchPutPolicySetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.authz.v1alpha1.AuthzService.BatchPutPolicySet is not implemented"))
}

func (UnimplementedAuthzServiceHandler) BatchAuthorize(context.Context, *connect_go.Request[v1alpha1.BatchAuthorizeRequest]) (*connect_go.Response[v1alpha1.BatchAuthorizeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.authz.v1alpha1.AuthzService.BatchAuthorize is not implemented"))
}

func (UnimplementedAuthzServiceHandler) ListPolicySets(context.Context, *connect_go.Request[v1alpha1.ListPolicySetsRequest]) (*connect_go.Response[v1alpha1.ListPolicySetsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.authz.v1alpha1.AuthzService.ListPolicySets is not implemented"))
}
