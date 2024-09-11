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
	// AccessServiceQueryApproversProcedure is the fully-qualified name of the AccessService's
	// QueryApprovers RPC.
	AccessServiceQueryApproversProcedure = "/commonfate.access.v1alpha1.AccessService/QueryApprovers"
	// AccessServicePreviewUserAccessProcedure is the fully-qualified name of the AccessService's
	// PreviewUserAccess RPC.
	AccessServicePreviewUserAccessProcedure = "/commonfate.access.v1alpha1.AccessService/PreviewUserAccess"
	// AccessServicePreviewEntitlementAccessProcedure is the fully-qualified name of the AccessService's
	// PreviewEntitlementAccess RPC.
	AccessServicePreviewEntitlementAccessProcedure = "/commonfate.access.v1alpha1.AccessService/PreviewEntitlementAccess"
	// AccessServiceDebugEntitlementAccessProcedure is the fully-qualified name of the AccessService's
	// DebugEntitlementAccess RPC.
	AccessServiceDebugEntitlementAccessProcedure = "/commonfate.access.v1alpha1.AccessService/DebugEntitlementAccess"
	// AccessServiceQueryEntitlementsTreeProcedure is the fully-qualified name of the AccessService's
	// QueryEntitlementsTree RPC.
	AccessServiceQueryEntitlementsTreeProcedure = "/commonfate.access.v1alpha1.AccessService/QueryEntitlementsTree"
	// AccessServicePreviewPolicyProcedure is the fully-qualified name of the AccessService's
	// PreviewPolicy RPC.
	AccessServicePreviewPolicyProcedure = "/commonfate.access.v1alpha1.AccessService/PreviewPolicy"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	accessServiceServiceDescriptor                        = v1alpha1.File_commonfate_access_v1alpha1_access_proto.Services().ByName("AccessService")
	accessServiceBatchEnsureMethodDescriptor              = accessServiceServiceDescriptor.Methods().ByName("BatchEnsure")
	accessServiceQueryAvailabilitiesMethodDescriptor      = accessServiceServiceDescriptor.Methods().ByName("QueryAvailabilities")
	accessServiceQueryEntitlementsMethodDescriptor        = accessServiceServiceDescriptor.Methods().ByName("QueryEntitlements")
	accessServiceQueryApproversMethodDescriptor           = accessServiceServiceDescriptor.Methods().ByName("QueryApprovers")
	accessServicePreviewUserAccessMethodDescriptor        = accessServiceServiceDescriptor.Methods().ByName("PreviewUserAccess")
	accessServicePreviewEntitlementAccessMethodDescriptor = accessServiceServiceDescriptor.Methods().ByName("PreviewEntitlementAccess")
	accessServiceDebugEntitlementAccessMethodDescriptor   = accessServiceServiceDescriptor.Methods().ByName("DebugEntitlementAccess")
	accessServiceQueryEntitlementsTreeMethodDescriptor    = accessServiceServiceDescriptor.Methods().ByName("QueryEntitlementsTree")
	accessServicePreviewPolicyMethodDescriptor            = accessServiceServiceDescriptor.Methods().ByName("PreviewPolicy")
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
	QueryApprovers(context.Context, *connect.Request[v1alpha1.QueryApproversRequest]) (*connect.Response[v1alpha1.QueryApproversResponse], error)
	PreviewUserAccess(context.Context, *connect.Request[v1alpha1.PreviewUserAccessRequest]) (*connect.Response[v1alpha1.PreviewUserAccessResponse], error)
	PreviewEntitlementAccess(context.Context, *connect.Request[v1alpha1.PreviewEntitlementAccessRequest]) (*connect.Response[v1alpha1.PreviewEntitlementAccessResponse], error)
	DebugEntitlementAccess(context.Context, *connect.Request[v1alpha1.DebugEntitlementAccessRequest]) (*connect.Response[v1alpha1.DebugEntitlementAccessResponse], error)
	QueryEntitlementsTree(context.Context, *connect.Request[v1alpha1.QueryEntitlementsTreeRequest]) (*connect.Response[v1alpha1.QueryEntitlementsTreeResponse], error)
	// returns all the users
	PreviewPolicy(context.Context, *connect.Request[v1alpha1.PreviewPolicyRequest]) (*connect.Response[v1alpha1.PreviewPolicyResponse], error)
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
		queryApprovers: connect.NewClient[v1alpha1.QueryApproversRequest, v1alpha1.QueryApproversResponse](
			httpClient,
			baseURL+AccessServiceQueryApproversProcedure,
			connect.WithSchema(accessServiceQueryApproversMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		previewUserAccess: connect.NewClient[v1alpha1.PreviewUserAccessRequest, v1alpha1.PreviewUserAccessResponse](
			httpClient,
			baseURL+AccessServicePreviewUserAccessProcedure,
			connect.WithSchema(accessServicePreviewUserAccessMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		previewEntitlementAccess: connect.NewClient[v1alpha1.PreviewEntitlementAccessRequest, v1alpha1.PreviewEntitlementAccessResponse](
			httpClient,
			baseURL+AccessServicePreviewEntitlementAccessProcedure,
			connect.WithSchema(accessServicePreviewEntitlementAccessMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		debugEntitlementAccess: connect.NewClient[v1alpha1.DebugEntitlementAccessRequest, v1alpha1.DebugEntitlementAccessResponse](
			httpClient,
			baseURL+AccessServiceDebugEntitlementAccessProcedure,
			connect.WithSchema(accessServiceDebugEntitlementAccessMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		queryEntitlementsTree: connect.NewClient[v1alpha1.QueryEntitlementsTreeRequest, v1alpha1.QueryEntitlementsTreeResponse](
			httpClient,
			baseURL+AccessServiceQueryEntitlementsTreeProcedure,
			connect.WithSchema(accessServiceQueryEntitlementsTreeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		previewPolicy: connect.NewClient[v1alpha1.PreviewPolicyRequest, v1alpha1.PreviewPolicyResponse](
			httpClient,
			baseURL+AccessServicePreviewPolicyProcedure,
			connect.WithSchema(accessServicePreviewPolicyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// accessServiceClient implements AccessServiceClient.
type accessServiceClient struct {
	batchEnsure              *connect.Client[v1alpha1.BatchEnsureRequest, v1alpha1.BatchEnsureResponse]
	queryAvailabilities      *connect.Client[v1alpha1.QueryAvailabilitiesRequest, v1alpha1.QueryAvailabilitiesResponse]
	queryEntitlements        *connect.Client[v1alpha1.QueryEntitlementsRequest, v1alpha1.QueryEntitlementsResponse]
	queryApprovers           *connect.Client[v1alpha1.QueryApproversRequest, v1alpha1.QueryApproversResponse]
	previewUserAccess        *connect.Client[v1alpha1.PreviewUserAccessRequest, v1alpha1.PreviewUserAccessResponse]
	previewEntitlementAccess *connect.Client[v1alpha1.PreviewEntitlementAccessRequest, v1alpha1.PreviewEntitlementAccessResponse]
	debugEntitlementAccess   *connect.Client[v1alpha1.DebugEntitlementAccessRequest, v1alpha1.DebugEntitlementAccessResponse]
	queryEntitlementsTree    *connect.Client[v1alpha1.QueryEntitlementsTreeRequest, v1alpha1.QueryEntitlementsTreeResponse]
	previewPolicy            *connect.Client[v1alpha1.PreviewPolicyRequest, v1alpha1.PreviewPolicyResponse]
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

// QueryApprovers calls commonfate.access.v1alpha1.AccessService.QueryApprovers.
func (c *accessServiceClient) QueryApprovers(ctx context.Context, req *connect.Request[v1alpha1.QueryApproversRequest]) (*connect.Response[v1alpha1.QueryApproversResponse], error) {
	return c.queryApprovers.CallUnary(ctx, req)
}

// PreviewUserAccess calls commonfate.access.v1alpha1.AccessService.PreviewUserAccess.
func (c *accessServiceClient) PreviewUserAccess(ctx context.Context, req *connect.Request[v1alpha1.PreviewUserAccessRequest]) (*connect.Response[v1alpha1.PreviewUserAccessResponse], error) {
	return c.previewUserAccess.CallUnary(ctx, req)
}

// PreviewEntitlementAccess calls commonfate.access.v1alpha1.AccessService.PreviewEntitlementAccess.
func (c *accessServiceClient) PreviewEntitlementAccess(ctx context.Context, req *connect.Request[v1alpha1.PreviewEntitlementAccessRequest]) (*connect.Response[v1alpha1.PreviewEntitlementAccessResponse], error) {
	return c.previewEntitlementAccess.CallUnary(ctx, req)
}

// DebugEntitlementAccess calls commonfate.access.v1alpha1.AccessService.DebugEntitlementAccess.
func (c *accessServiceClient) DebugEntitlementAccess(ctx context.Context, req *connect.Request[v1alpha1.DebugEntitlementAccessRequest]) (*connect.Response[v1alpha1.DebugEntitlementAccessResponse], error) {
	return c.debugEntitlementAccess.CallUnary(ctx, req)
}

// QueryEntitlementsTree calls commonfate.access.v1alpha1.AccessService.QueryEntitlementsTree.
func (c *accessServiceClient) QueryEntitlementsTree(ctx context.Context, req *connect.Request[v1alpha1.QueryEntitlementsTreeRequest]) (*connect.Response[v1alpha1.QueryEntitlementsTreeResponse], error) {
	return c.queryEntitlementsTree.CallUnary(ctx, req)
}

// PreviewPolicy calls commonfate.access.v1alpha1.AccessService.PreviewPolicy.
func (c *accessServiceClient) PreviewPolicy(ctx context.Context, req *connect.Request[v1alpha1.PreviewPolicyRequest]) (*connect.Response[v1alpha1.PreviewPolicyResponse], error) {
	return c.previewPolicy.CallUnary(ctx, req)
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
	QueryApprovers(context.Context, *connect.Request[v1alpha1.QueryApproversRequest]) (*connect.Response[v1alpha1.QueryApproversResponse], error)
	PreviewUserAccess(context.Context, *connect.Request[v1alpha1.PreviewUserAccessRequest]) (*connect.Response[v1alpha1.PreviewUserAccessResponse], error)
	PreviewEntitlementAccess(context.Context, *connect.Request[v1alpha1.PreviewEntitlementAccessRequest]) (*connect.Response[v1alpha1.PreviewEntitlementAccessResponse], error)
	DebugEntitlementAccess(context.Context, *connect.Request[v1alpha1.DebugEntitlementAccessRequest]) (*connect.Response[v1alpha1.DebugEntitlementAccessResponse], error)
	QueryEntitlementsTree(context.Context, *connect.Request[v1alpha1.QueryEntitlementsTreeRequest]) (*connect.Response[v1alpha1.QueryEntitlementsTreeResponse], error)
	// returns all the users
	PreviewPolicy(context.Context, *connect.Request[v1alpha1.PreviewPolicyRequest]) (*connect.Response[v1alpha1.PreviewPolicyResponse], error)
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
	accessServiceQueryApproversHandler := connect.NewUnaryHandler(
		AccessServiceQueryApproversProcedure,
		svc.QueryApprovers,
		connect.WithSchema(accessServiceQueryApproversMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accessServicePreviewUserAccessHandler := connect.NewUnaryHandler(
		AccessServicePreviewUserAccessProcedure,
		svc.PreviewUserAccess,
		connect.WithSchema(accessServicePreviewUserAccessMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accessServicePreviewEntitlementAccessHandler := connect.NewUnaryHandler(
		AccessServicePreviewEntitlementAccessProcedure,
		svc.PreviewEntitlementAccess,
		connect.WithSchema(accessServicePreviewEntitlementAccessMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accessServiceDebugEntitlementAccessHandler := connect.NewUnaryHandler(
		AccessServiceDebugEntitlementAccessProcedure,
		svc.DebugEntitlementAccess,
		connect.WithSchema(accessServiceDebugEntitlementAccessMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accessServiceQueryEntitlementsTreeHandler := connect.NewUnaryHandler(
		AccessServiceQueryEntitlementsTreeProcedure,
		svc.QueryEntitlementsTree,
		connect.WithSchema(accessServiceQueryEntitlementsTreeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	accessServicePreviewPolicyHandler := connect.NewUnaryHandler(
		AccessServicePreviewPolicyProcedure,
		svc.PreviewPolicy,
		connect.WithSchema(accessServicePreviewPolicyMethodDescriptor),
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
		case AccessServiceQueryApproversProcedure:
			accessServiceQueryApproversHandler.ServeHTTP(w, r)
		case AccessServicePreviewUserAccessProcedure:
			accessServicePreviewUserAccessHandler.ServeHTTP(w, r)
		case AccessServicePreviewEntitlementAccessProcedure:
			accessServicePreviewEntitlementAccessHandler.ServeHTTP(w, r)
		case AccessServiceDebugEntitlementAccessProcedure:
			accessServiceDebugEntitlementAccessHandler.ServeHTTP(w, r)
		case AccessServiceQueryEntitlementsTreeProcedure:
			accessServiceQueryEntitlementsTreeHandler.ServeHTTP(w, r)
		case AccessServicePreviewPolicyProcedure:
			accessServicePreviewPolicyHandler.ServeHTTP(w, r)
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

func (UnimplementedAccessServiceHandler) QueryApprovers(context.Context, *connect.Request[v1alpha1.QueryApproversRequest]) (*connect.Response[v1alpha1.QueryApproversResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessService.QueryApprovers is not implemented"))
}

func (UnimplementedAccessServiceHandler) PreviewUserAccess(context.Context, *connect.Request[v1alpha1.PreviewUserAccessRequest]) (*connect.Response[v1alpha1.PreviewUserAccessResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessService.PreviewUserAccess is not implemented"))
}

func (UnimplementedAccessServiceHandler) PreviewEntitlementAccess(context.Context, *connect.Request[v1alpha1.PreviewEntitlementAccessRequest]) (*connect.Response[v1alpha1.PreviewEntitlementAccessResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessService.PreviewEntitlementAccess is not implemented"))
}

func (UnimplementedAccessServiceHandler) DebugEntitlementAccess(context.Context, *connect.Request[v1alpha1.DebugEntitlementAccessRequest]) (*connect.Response[v1alpha1.DebugEntitlementAccessResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessService.DebugEntitlementAccess is not implemented"))
}

func (UnimplementedAccessServiceHandler) QueryEntitlementsTree(context.Context, *connect.Request[v1alpha1.QueryEntitlementsTreeRequest]) (*connect.Response[v1alpha1.QueryEntitlementsTreeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessService.QueryEntitlementsTree is not implemented"))
}

func (UnimplementedAccessServiceHandler) PreviewPolicy(context.Context, *connect.Request[v1alpha1.PreviewPolicyRequest]) (*connect.Response[v1alpha1.PreviewPolicyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.access.v1alpha1.AccessService.PreviewPolicy is not implemented"))
}
