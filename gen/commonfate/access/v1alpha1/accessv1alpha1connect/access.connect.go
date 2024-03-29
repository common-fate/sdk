// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/access/v1alpha1/access.proto

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
	// AccessServiceName is the fully-qualified name of the AccessService service.
	AccessServiceName = "commonfate.access.v1alpha1.AccessService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AccessServiceBatchEnsureProcedure is the fully-qualified name of the AccessService's BatchEnsure
	// RPC.
	AccessServiceBatchEnsureProcedure = "/commonfate.access.v1alpha1.AccessService/BatchEnsure"
	// AccessServiceQueryAvailabilitiesProcedure is the fully-qualified name of the AccessService's
	// QueryAvailabilities RPC.
	AccessServiceQueryAvailabilitiesProcedure = "/commonfate.access.v1alpha1.AccessService/QueryAvailabilities"
	// AccessServiceQueryEntitlementsProcedure is the fully-qualified name of the AccessService's
	// QueryEntitlements RPC.
	AccessServiceQueryEntitlementsProcedure = "/commonfate.access.v1alpha1.AccessService/QueryEntitlements"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	accessServiceServiceDescriptor                   = v1alpha1.File_commonfate_access_v1alpha1_access_proto.Services().ByName("AccessService")
	accessServiceBatchEnsureMethodDescriptor         = accessServiceServiceDescriptor.Methods().ByName("BatchEnsure")
	accessServiceQueryAvailabilitiesMethodDescriptor = accessServiceServiceDescriptor.Methods().ByName("QueryAvailabilities")
	accessServiceQueryEntitlementsMethodDescriptor   = accessServiceServiceDescriptor.Methods().ByName("QueryEntitlements")
)

// AccessServiceClient is a client for the commonfate.access.v1alpha1.AccessService service.
type AccessServiceClient interface {
	// BatchEnsure is a high-level declarative API which can be called to ensure access has been provisioned to multiple entitlements.
	//
	// The method checks whether the entitlement has been provisioned to the user.
	// If the entitlement has not been provisioned, an Access Request will be created for the entitlement.
	// If a pending Access Request exists for the entitlement, this request is returned.
	//
	// In future, this method may trigger an extension to any Access Requests which are due to expire.
	BatchEnsure(context.Context, *connect.Request[v1alpha1.BatchEnsureRequest]) (*connect.Response[v1alpha1.BatchEnsureResponse], error)
	// Query for JIT availabilities.
	QueryAvailabilities(context.Context, *connect.Request[v1alpha1.QueryAvailabilitiesRequest]) (*connect.Response[v1alpha1.QueryAvailabilitiesResponse], error)
	QueryEntitlements(context.Context, *connect.Request[v1alpha1.QueryEntitlementsRequest]) (*connect.Response[v1alpha1.QueryEntitlementsResponse], error)
}

// NewAccessServiceClient constructs a client for the commonfate.access.v1alpha1.AccessService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAccessServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AccessServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &accessServiceClient{
		batchEnsure: connect.NewClient[v1alpha1.BatchEnsureRequest, v1alpha1.BatchEnsureResponse](
			httpClient,
			baseURL+AccessServiceBatchEnsureProcedure,
			connect.WithSchema(accessServiceBatchEnsureMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		queryAvailabilities: connect.NewClient[v1alpha1.QueryAvailabilitiesRequest, v1alpha1.QueryAvailabilitiesResponse](
			httpClient,
			baseURL+AccessServiceQueryAvailabilitiesProcedure,
			connect.WithSchema(accessServiceQueryAvailabilitiesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		queryEntitlements: connect.NewClient[v1alpha1.QueryEntitlementsRequest, v1alpha1.QueryEntitlementsResponse](
			httpClient,
			baseURL+AccessServiceQueryEntitlementsProcedure,
			connect.WithSchema(accessServiceQueryEntitlementsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// accessServiceClient implements AccessServiceClient.
type accessServiceClient struct {
	batchEnsure         *connect.Client[v1alpha1.BatchEnsureRequest, v1alpha1.BatchEnsureResponse]
	queryAvailabilities *connect.Client[v1alpha1.QueryAvailabilitiesRequest, v1alpha1.QueryAvailabilitiesResponse]
	queryEntitlements   *connect.Client[v1alpha1.QueryEntitlementsRequest, v1alpha1.QueryEntitlementsResponse]
}

// BatchEnsure calls commonfate.access.v1alpha1.AccessService.BatchEnsure.
func (c *accessServiceClient) BatchEnsure(ctx context.Context, req *connect.Request[v1alpha1.BatchEnsureRequest]) (*connect.Response[v1alpha1.BatchEnsureResponse], error) {
	return c.batchEnsure.CallUnary(ctx, req)
}

// QueryAvailabilities calls commonfate.access.v1alpha1.AccessService.QueryAvailabilities.
func (c *accessServiceClient) QueryAvailabilities(ctx context.Context, req *connect.Request[v1alpha1.QueryAvailabilitiesRequest]) (*connect.Response[v1alpha1.QueryAvailabilitiesResponse], error) {
	return c.queryAvailabilities.CallUnary(ctx, req)
}

// QueryEntitlements calls commonfate.access.v1alpha1.AccessService.QueryEntitlements.
func (c *accessServiceClient) QueryEntitlements(ctx context.Context, req *connect.Request[v1alpha1.QueryEntitlementsRequest]) (*connect.Response[v1alpha1.QueryEntitlementsResponse], error) {
	return c.queryEntitlements.CallUnary(ctx, req)
}

// AccessServiceHandler is an implementation of the commonfate.access.v1alpha1.AccessService
// service.
type AccessServiceHandler interface {
	// BatchEnsure is a high-level declarative API which can be called to ensure access has been provisioned to multiple entitlements.
	//
	// The method checks whether the entitlement has been provisioned to the user.
	// If the entitlement has not been provisioned, an Access Request will be created for the entitlement.
	// If a pending Access Request exists for the entitlement, this request is returned.
	//
	// In future, this method may trigger an extension to any Access Requests which are due to expire.
	BatchEnsure(context.Context, *connect.Request[v1alpha1.BatchEnsureRequest]) (*connect.Response[v1alpha1.BatchEnsureResponse], error)
	// Query for JIT availabilities.
	QueryAvailabilities(context.Context, *connect.Request[v1alpha1.QueryAvailabilitiesRequest]) (*connect.Response[v1alpha1.QueryAvailabilitiesResponse], error)
	QueryEntitlements(context.Context, *connect.Request[v1alpha1.QueryEntitlementsRequest]) (*connect.Response[v1alpha1.QueryEntitlementsResponse], error)
}

// NewAccessServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAccessServiceHandler(svc AccessServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	accessServiceBatchEnsureHandler := connect.NewUnaryHandler(
		AccessServiceBatchEnsureProcedure,
		svc.BatchEnsure,
		connect.WithSchema(accessServiceBatchEnsureMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accessServiceQueryAvailabilitiesHandler := connect.NewUnaryHandler(
		AccessServiceQueryAvailabilitiesProcedure,
		svc.QueryAvailabilities,
		connect.WithSchema(accessServiceQueryAvailabilitiesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accessServiceQueryEntitlementsHandler := connect.NewUnaryHandler(
		AccessServiceQueryEntitlementsProcedure,
		svc.QueryEntitlements,
		connect.WithSchema(accessServiceQueryEntitlementsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.access.v1alpha1.AccessService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AccessServiceBatchEnsureProcedure:
			accessServiceBatchEnsureHandler.ServeHTTP(w, r)
		case AccessServiceQueryAvailabilitiesProcedure:
			accessServiceQueryAvailabilitiesHandler.ServeHTTP(w, r)
		case AccessServiceQueryEntitlementsProcedure:
			accessServiceQueryEntitlementsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAccessServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAccessServiceHandler struct{}

func (UnimplementedAccessServiceHandler) BatchEnsure(context.Context, *connect.Request[v1alpha1.BatchEnsureRequest]) (*connect.Response[v1alpha1.BatchEnsureResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessService.BatchEnsure is not implemented"))
}

func (UnimplementedAccessServiceHandler) QueryAvailabilities(context.Context, *connect.Request[v1alpha1.QueryAvailabilitiesRequest]) (*connect.Response[v1alpha1.QueryAvailabilitiesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessService.QueryAvailabilities is not implemented"))
}

func (UnimplementedAccessServiceHandler) QueryEntitlements(context.Context, *connect.Request[v1alpha1.QueryEntitlementsRequest]) (*connect.Response[v1alpha1.QueryEntitlementsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessService.QueryEntitlements is not implemented"))
}
