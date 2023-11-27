// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/access/v1alpha1/access_request.proto

package accessv1alpha1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
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
	// AccessRequestServiceName is the fully-qualified name of the AccessRequestService service.
	AccessRequestServiceName = "commonfate.access.v1alpha1.AccessRequestService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AccessRequestServiceQueryAccessRequestsProcedure is the fully-qualified name of the
	// AccessRequestService's QueryAccessRequests RPC.
	AccessRequestServiceQueryAccessRequestsProcedure = "/commonfate.access.v1alpha1.AccessRequestService/QueryAccessRequests"
	// AccessRequestServiceGetAccessRequestProcedure is the fully-qualified name of the
	// AccessRequestService's GetAccessRequest RPC.
	AccessRequestServiceGetAccessRequestProcedure = "/commonfate.access.v1alpha1.AccessRequestService/GetAccessRequest"
	// AccessRequestServiceRevokeAccessRequestProcedure is the fully-qualified name of the
	// AccessRequestService's RevokeAccessRequest RPC.
	AccessRequestServiceRevokeAccessRequestProcedure = "/commonfate.access.v1alpha1.AccessRequestService/RevokeAccessRequest"
	// AccessRequestServiceCancelAccessRequestProcedure is the fully-qualified name of the
	// AccessRequestService's CancelAccessRequest RPC.
	AccessRequestServiceCancelAccessRequestProcedure = "/commonfate.access.v1alpha1.AccessRequestService/CancelAccessRequest"
	// AccessRequestServiceReviewAccessRequestProcedure is the fully-qualified name of the
	// AccessRequestService's ReviewAccessRequest RPC.
	AccessRequestServiceReviewAccessRequestProcedure = "/commonfate.access.v1alpha1.AccessRequestService/ReviewAccessRequest"
)

// AccessRequestServiceClient is a client for the commonfate.access.v1alpha1.AccessRequestService
// service.
type AccessRequestServiceClient interface {
	QueryAccessRequests(context.Context, *connect_go.Request[v1alpha1.QueryAccessRequestsRequest]) (*connect_go.Response[v1alpha1.QueryAccessRequestsResponse], error)
	GetAccessRequest(context.Context, *connect_go.Request[v1alpha1.GetAccessRequestRequest]) (*connect_go.Response[v1alpha1.GetAccessRequestResponse], error)
	RevokeAccessRequest(context.Context, *connect_go.Request[v1alpha1.RevokeAccessRequestRequest]) (*connect_go.Response[v1alpha1.RevokeAccessRequestResponse], error)
	CancelAccessRequest(context.Context, *connect_go.Request[v1alpha1.CancelAccessRequestRequest]) (*connect_go.Response[v1alpha1.CancelAccessRequestResponse], error)
	ReviewAccessRequest(context.Context, *connect_go.Request[v1alpha1.ReviewAccessRequestRequest]) (*connect_go.Response[v1alpha1.ReviewAccessRequestResponse], error)
}

// NewAccessRequestServiceClient constructs a client for the
// commonfate.access.v1alpha1.AccessRequestService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAccessRequestServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AccessRequestServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &accessRequestServiceClient{
		queryAccessRequests: connect_go.NewClient[v1alpha1.QueryAccessRequestsRequest, v1alpha1.QueryAccessRequestsResponse](
			httpClient,
			baseURL+AccessRequestServiceQueryAccessRequestsProcedure,
			opts...,
		),
		getAccessRequest: connect_go.NewClient[v1alpha1.GetAccessRequestRequest, v1alpha1.GetAccessRequestResponse](
			httpClient,
			baseURL+AccessRequestServiceGetAccessRequestProcedure,
			opts...,
		),
		revokeAccessRequest: connect_go.NewClient[v1alpha1.RevokeAccessRequestRequest, v1alpha1.RevokeAccessRequestResponse](
			httpClient,
			baseURL+AccessRequestServiceRevokeAccessRequestProcedure,
			opts...,
		),
		cancelAccessRequest: connect_go.NewClient[v1alpha1.CancelAccessRequestRequest, v1alpha1.CancelAccessRequestResponse](
			httpClient,
			baseURL+AccessRequestServiceCancelAccessRequestProcedure,
			opts...,
		),
		reviewAccessRequest: connect_go.NewClient[v1alpha1.ReviewAccessRequestRequest, v1alpha1.ReviewAccessRequestResponse](
			httpClient,
			baseURL+AccessRequestServiceReviewAccessRequestProcedure,
			opts...,
		),
	}
}

// accessRequestServiceClient implements AccessRequestServiceClient.
type accessRequestServiceClient struct {
	queryAccessRequests *connect_go.Client[v1alpha1.QueryAccessRequestsRequest, v1alpha1.QueryAccessRequestsResponse]
	getAccessRequest    *connect_go.Client[v1alpha1.GetAccessRequestRequest, v1alpha1.GetAccessRequestResponse]
	revokeAccessRequest *connect_go.Client[v1alpha1.RevokeAccessRequestRequest, v1alpha1.RevokeAccessRequestResponse]
	cancelAccessRequest *connect_go.Client[v1alpha1.CancelAccessRequestRequest, v1alpha1.CancelAccessRequestResponse]
	reviewAccessRequest *connect_go.Client[v1alpha1.ReviewAccessRequestRequest, v1alpha1.ReviewAccessRequestResponse]
}

// QueryAccessRequests calls commonfate.access.v1alpha1.AccessRequestService.QueryAccessRequests.
func (c *accessRequestServiceClient) QueryAccessRequests(ctx context.Context, req *connect_go.Request[v1alpha1.QueryAccessRequestsRequest]) (*connect_go.Response[v1alpha1.QueryAccessRequestsResponse], error) {
	return c.queryAccessRequests.CallUnary(ctx, req)
}

// GetAccessRequest calls commonfate.access.v1alpha1.AccessRequestService.GetAccessRequest.
func (c *accessRequestServiceClient) GetAccessRequest(ctx context.Context, req *connect_go.Request[v1alpha1.GetAccessRequestRequest]) (*connect_go.Response[v1alpha1.GetAccessRequestResponse], error) {
	return c.getAccessRequest.CallUnary(ctx, req)
}

// RevokeAccessRequest calls commonfate.access.v1alpha1.AccessRequestService.RevokeAccessRequest.
func (c *accessRequestServiceClient) RevokeAccessRequest(ctx context.Context, req *connect_go.Request[v1alpha1.RevokeAccessRequestRequest]) (*connect_go.Response[v1alpha1.RevokeAccessRequestResponse], error) {
	return c.revokeAccessRequest.CallUnary(ctx, req)
}

// CancelAccessRequest calls commonfate.access.v1alpha1.AccessRequestService.CancelAccessRequest.
func (c *accessRequestServiceClient) CancelAccessRequest(ctx context.Context, req *connect_go.Request[v1alpha1.CancelAccessRequestRequest]) (*connect_go.Response[v1alpha1.CancelAccessRequestResponse], error) {
	return c.cancelAccessRequest.CallUnary(ctx, req)
}

// ReviewAccessRequest calls commonfate.access.v1alpha1.AccessRequestService.ReviewAccessRequest.
func (c *accessRequestServiceClient) ReviewAccessRequest(ctx context.Context, req *connect_go.Request[v1alpha1.ReviewAccessRequestRequest]) (*connect_go.Response[v1alpha1.ReviewAccessRequestResponse], error) {
	return c.reviewAccessRequest.CallUnary(ctx, req)
}

// AccessRequestServiceHandler is an implementation of the
// commonfate.access.v1alpha1.AccessRequestService service.
type AccessRequestServiceHandler interface {
	QueryAccessRequests(context.Context, *connect_go.Request[v1alpha1.QueryAccessRequestsRequest]) (*connect_go.Response[v1alpha1.QueryAccessRequestsResponse], error)
	GetAccessRequest(context.Context, *connect_go.Request[v1alpha1.GetAccessRequestRequest]) (*connect_go.Response[v1alpha1.GetAccessRequestResponse], error)
	RevokeAccessRequest(context.Context, *connect_go.Request[v1alpha1.RevokeAccessRequestRequest]) (*connect_go.Response[v1alpha1.RevokeAccessRequestResponse], error)
	CancelAccessRequest(context.Context, *connect_go.Request[v1alpha1.CancelAccessRequestRequest]) (*connect_go.Response[v1alpha1.CancelAccessRequestResponse], error)
	ReviewAccessRequest(context.Context, *connect_go.Request[v1alpha1.ReviewAccessRequestRequest]) (*connect_go.Response[v1alpha1.ReviewAccessRequestResponse], error)
}

// NewAccessRequestServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAccessRequestServiceHandler(svc AccessRequestServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	accessRequestServiceQueryAccessRequestsHandler := connect_go.NewUnaryHandler(
		AccessRequestServiceQueryAccessRequestsProcedure,
		svc.QueryAccessRequests,
		opts...,
	)
	accessRequestServiceGetAccessRequestHandler := connect_go.NewUnaryHandler(
		AccessRequestServiceGetAccessRequestProcedure,
		svc.GetAccessRequest,
		opts...,
	)
	accessRequestServiceRevokeAccessRequestHandler := connect_go.NewUnaryHandler(
		AccessRequestServiceRevokeAccessRequestProcedure,
		svc.RevokeAccessRequest,
		opts...,
	)
	accessRequestServiceCancelAccessRequestHandler := connect_go.NewUnaryHandler(
		AccessRequestServiceCancelAccessRequestProcedure,
		svc.CancelAccessRequest,
		opts...,
	)
	accessRequestServiceReviewAccessRequestHandler := connect_go.NewUnaryHandler(
		AccessRequestServiceReviewAccessRequestProcedure,
		svc.ReviewAccessRequest,
		opts...,
	)
	return "/commonfate.access.v1alpha1.AccessRequestService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AccessRequestServiceQueryAccessRequestsProcedure:
			accessRequestServiceQueryAccessRequestsHandler.ServeHTTP(w, r)
		case AccessRequestServiceGetAccessRequestProcedure:
			accessRequestServiceGetAccessRequestHandler.ServeHTTP(w, r)
		case AccessRequestServiceRevokeAccessRequestProcedure:
			accessRequestServiceRevokeAccessRequestHandler.ServeHTTP(w, r)
		case AccessRequestServiceCancelAccessRequestProcedure:
			accessRequestServiceCancelAccessRequestHandler.ServeHTTP(w, r)
		case AccessRequestServiceReviewAccessRequestProcedure:
			accessRequestServiceReviewAccessRequestHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAccessRequestServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAccessRequestServiceHandler struct{}

func (UnimplementedAccessRequestServiceHandler) QueryAccessRequests(context.Context, *connect_go.Request[v1alpha1.QueryAccessRequestsRequest]) (*connect_go.Response[v1alpha1.QueryAccessRequestsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessRequestService.QueryAccessRequests is not implemented"))
}

func (UnimplementedAccessRequestServiceHandler) GetAccessRequest(context.Context, *connect_go.Request[v1alpha1.GetAccessRequestRequest]) (*connect_go.Response[v1alpha1.GetAccessRequestResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessRequestService.GetAccessRequest is not implemented"))
}

func (UnimplementedAccessRequestServiceHandler) RevokeAccessRequest(context.Context, *connect_go.Request[v1alpha1.RevokeAccessRequestRequest]) (*connect_go.Response[v1alpha1.RevokeAccessRequestResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessRequestService.RevokeAccessRequest is not implemented"))
}

func (UnimplementedAccessRequestServiceHandler) CancelAccessRequest(context.Context, *connect_go.Request[v1alpha1.CancelAccessRequestRequest]) (*connect_go.Response[v1alpha1.CancelAccessRequestResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessRequestService.CancelAccessRequest is not implemented"))
}

func (UnimplementedAccessRequestServiceHandler) ReviewAccessRequest(context.Context, *connect_go.Request[v1alpha1.ReviewAccessRequestRequest]) (*connect_go.Response[v1alpha1.ReviewAccessRequestResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessRequestService.ReviewAccessRequest is not implemented"))
}
