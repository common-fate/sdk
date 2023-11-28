// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/access/v1alpha1/audit_logs.proto

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
	// AuditLogServiceName is the fully-qualified name of the AuditLogService service.
	AuditLogServiceName = "commonfate.access.v1alpha1.AuditLogService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuditLogServiceQueryAuditLogsProcedure is the fully-qualified name of the AuditLogService's
	// QueryAuditLogs RPC.
	AuditLogServiceQueryAuditLogsProcedure = "/commonfate.access.v1alpha1.AuditLogService/QueryAuditLogs"
)

// AuditLogServiceClient is a client for the commonfate.access.v1alpha1.AuditLogService service.
type AuditLogServiceClient interface {
	QueryAuditLogs(context.Context, *connect_go.Request[v1alpha1.QueryAuditLogsRequest]) (*connect_go.Response[v1alpha1.QueryAuditLogsResponse], error)
}

// NewAuditLogServiceClient constructs a client for the commonfate.access.v1alpha1.AuditLogService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuditLogServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AuditLogServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &auditLogServiceClient{
		queryAuditLogs: connect_go.NewClient[v1alpha1.QueryAuditLogsRequest, v1alpha1.QueryAuditLogsResponse](
			httpClient,
			baseURL+AuditLogServiceQueryAuditLogsProcedure,
			opts...,
		),
	}
}

// auditLogServiceClient implements AuditLogServiceClient.
type auditLogServiceClient struct {
	queryAuditLogs *connect_go.Client[v1alpha1.QueryAuditLogsRequest, v1alpha1.QueryAuditLogsResponse]
}

// QueryAuditLogs calls commonfate.access.v1alpha1.AuditLogService.QueryAuditLogs.
func (c *auditLogServiceClient) QueryAuditLogs(ctx context.Context, req *connect_go.Request[v1alpha1.QueryAuditLogsRequest]) (*connect_go.Response[v1alpha1.QueryAuditLogsResponse], error) {
	return c.queryAuditLogs.CallUnary(ctx, req)
}

// AuditLogServiceHandler is an implementation of the commonfate.access.v1alpha1.AuditLogService
// service.
type AuditLogServiceHandler interface {
	QueryAuditLogs(context.Context, *connect_go.Request[v1alpha1.QueryAuditLogsRequest]) (*connect_go.Response[v1alpha1.QueryAuditLogsResponse], error)
}

// NewAuditLogServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuditLogServiceHandler(svc AuditLogServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	auditLogServiceQueryAuditLogsHandler := connect_go.NewUnaryHandler(
		AuditLogServiceQueryAuditLogsProcedure,
		svc.QueryAuditLogs,
		opts...,
	)
	return "/commonfate.access.v1alpha1.AuditLogService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuditLogServiceQueryAuditLogsProcedure:
			auditLogServiceQueryAuditLogsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuditLogServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuditLogServiceHandler struct{}

func (UnimplementedAuditLogServiceHandler) QueryAuditLogs(context.Context, *connect_go.Request[v1alpha1.QueryAuditLogsRequest]) (*connect_go.Response[v1alpha1.QueryAuditLogsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AuditLogService.QueryAuditLogs is not implemented"))
}
