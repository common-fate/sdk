// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/directory/v1alpha1/directory.proto

package directoryv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/control/directory/v1alpha1"
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
	// DirectoryServiceName is the fully-qualified name of the DirectoryService service.
	DirectoryServiceName = "commonfate.control.directory.v1alpha1.DirectoryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DirectoryServiceQueryUsersProcedure is the fully-qualified name of the DirectoryService's
	// QueryUsers RPC.
	DirectoryServiceQueryUsersProcedure = "/commonfate.control.directory.v1alpha1.DirectoryService/QueryUsers"
	// DirectoryServiceGetUserProcedure is the fully-qualified name of the DirectoryService's GetUser
	// RPC.
	DirectoryServiceGetUserProcedure = "/commonfate.control.directory.v1alpha1.DirectoryService/GetUser"
	// DirectoryServiceQueryGroupsProcedure is the fully-qualified name of the DirectoryService's
	// QueryGroups RPC.
	DirectoryServiceQueryGroupsProcedure = "/commonfate.control.directory.v1alpha1.DirectoryService/QueryGroups"
	// DirectoryServiceGetGroupProcedure is the fully-qualified name of the DirectoryService's GetGroup
	// RPC.
	DirectoryServiceGetGroupProcedure = "/commonfate.control.directory.v1alpha1.DirectoryService/GetGroup"
	// DirectoryServiceQueryGroupMembersProcedure is the fully-qualified name of the DirectoryService's
	// QueryGroupMembers RPC.
	DirectoryServiceQueryGroupMembersProcedure = "/commonfate.control.directory.v1alpha1.DirectoryService/QueryGroupMembers"
	// DirectoryServiceQueryChildGroupsProcedure is the fully-qualified name of the DirectoryService's
	// QueryChildGroups RPC.
	DirectoryServiceQueryChildGroupsProcedure = "/commonfate.control.directory.v1alpha1.DirectoryService/QueryChildGroups"
	// DirectoryServiceQueryGroupsForUserProcedure is the fully-qualified name of the DirectoryService's
	// QueryGroupsForUser RPC.
	DirectoryServiceQueryGroupsForUserProcedure = "/commonfate.control.directory.v1alpha1.DirectoryService/QueryGroupsForUser"
	// DirectoryServiceLookupUserAccountProcedure is the fully-qualified name of the DirectoryService's
	// LookupUserAccount RPC.
	DirectoryServiceLookupUserAccountProcedure = "/commonfate.control.directory.v1alpha1.DirectoryService/LookupUserAccount"
	// DirectoryServiceDeleteUserProcedure is the fully-qualified name of the DirectoryService's
	// DeleteUser RPC.
	DirectoryServiceDeleteUserProcedure = "/commonfate.control.directory.v1alpha1.DirectoryService/DeleteUser"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	directoryServiceServiceDescriptor                  = v1alpha1.File_commonfate_control_directory_v1alpha1_directory_proto.Services().ByName("DirectoryService")
	directoryServiceQueryUsersMethodDescriptor         = directoryServiceServiceDescriptor.Methods().ByName("QueryUsers")
	directoryServiceGetUserMethodDescriptor            = directoryServiceServiceDescriptor.Methods().ByName("GetUser")
	directoryServiceQueryGroupsMethodDescriptor        = directoryServiceServiceDescriptor.Methods().ByName("QueryGroups")
	directoryServiceGetGroupMethodDescriptor           = directoryServiceServiceDescriptor.Methods().ByName("GetGroup")
	directoryServiceQueryGroupMembersMethodDescriptor  = directoryServiceServiceDescriptor.Methods().ByName("QueryGroupMembers")
	directoryServiceQueryChildGroupsMethodDescriptor   = directoryServiceServiceDescriptor.Methods().ByName("QueryChildGroups")
	directoryServiceQueryGroupsForUserMethodDescriptor = directoryServiceServiceDescriptor.Methods().ByName("QueryGroupsForUser")
	directoryServiceLookupUserAccountMethodDescriptor  = directoryServiceServiceDescriptor.Methods().ByName("LookupUserAccount")
	directoryServiceDeleteUserMethodDescriptor         = directoryServiceServiceDescriptor.Methods().ByName("DeleteUser")
)

// DirectoryServiceClient is a client for the commonfate.control.directory.v1alpha1.DirectoryService
// service.
type DirectoryServiceClient interface {
	QueryUsers(context.Context, *connect.Request[v1alpha1.QueryUsersRequest]) (*connect.Response[v1alpha1.QueryUsersResponse], error)
	GetUser(context.Context, *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error)
	QueryGroups(context.Context, *connect.Request[v1alpha1.QueryGroupsRequest]) (*connect.Response[v1alpha1.QueryGroupsResponse], error)
	GetGroup(context.Context, *connect.Request[v1alpha1.GetGroupRequest]) (*connect.Response[v1alpha1.GetGroupResponse], error)
	QueryGroupMembers(context.Context, *connect.Request[v1alpha1.QueryGroupMembersRequest]) (*connect.Response[v1alpha1.QueryGroupMembersResponse], error)
	QueryChildGroups(context.Context, *connect.Request[v1alpha1.QueryChildGroupsRequest]) (*connect.Response[v1alpha1.QueryChildGroupsResponse], error)
	// Lists the groups that a user is a member of.
	QueryGroupsForUser(context.Context, *connect.Request[v1alpha1.QueryGroupsForUserRequest]) (*connect.Response[v1alpha1.QueryGroupsForUserResponse], error)
	// Looks up the matching User for a particular
	// connected user account in an integration.
	LookupUserAccount(context.Context, *connect.Request[v1alpha1.LookupUserAccountRequest]) (*connect.Response[v1alpha1.LookupUserAccountResponse], error)
	DeleteUser(context.Context, *connect.Request[v1alpha1.DeleteUserRequest]) (*connect.Response[v1alpha1.DeleteUserResponse], error)
}

// NewDirectoryServiceClient constructs a client for the
// commonfate.control.directory.v1alpha1.DirectoryService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDirectoryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DirectoryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &directoryServiceClient{
		queryUsers: connect.NewClient[v1alpha1.QueryUsersRequest, v1alpha1.QueryUsersResponse](
			httpClient,
			baseURL+DirectoryServiceQueryUsersProcedure,
			connect.WithSchema(directoryServiceQueryUsersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getUser: connect.NewClient[v1alpha1.GetUserRequest, v1alpha1.GetUserResponse](
			httpClient,
			baseURL+DirectoryServiceGetUserProcedure,
			connect.WithSchema(directoryServiceGetUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		queryGroups: connect.NewClient[v1alpha1.QueryGroupsRequest, v1alpha1.QueryGroupsResponse](
			httpClient,
			baseURL+DirectoryServiceQueryGroupsProcedure,
			connect.WithSchema(directoryServiceQueryGroupsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getGroup: connect.NewClient[v1alpha1.GetGroupRequest, v1alpha1.GetGroupResponse](
			httpClient,
			baseURL+DirectoryServiceGetGroupProcedure,
			connect.WithSchema(directoryServiceGetGroupMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		queryGroupMembers: connect.NewClient[v1alpha1.QueryGroupMembersRequest, v1alpha1.QueryGroupMembersResponse](
			httpClient,
			baseURL+DirectoryServiceQueryGroupMembersProcedure,
			connect.WithSchema(directoryServiceQueryGroupMembersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		queryChildGroups: connect.NewClient[v1alpha1.QueryChildGroupsRequest, v1alpha1.QueryChildGroupsResponse](
			httpClient,
			baseURL+DirectoryServiceQueryChildGroupsProcedure,
			connect.WithSchema(directoryServiceQueryChildGroupsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		queryGroupsForUser: connect.NewClient[v1alpha1.QueryGroupsForUserRequest, v1alpha1.QueryGroupsForUserResponse](
			httpClient,
			baseURL+DirectoryServiceQueryGroupsForUserProcedure,
			connect.WithSchema(directoryServiceQueryGroupsForUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		lookupUserAccount: connect.NewClient[v1alpha1.LookupUserAccountRequest, v1alpha1.LookupUserAccountResponse](
			httpClient,
			baseURL+DirectoryServiceLookupUserAccountProcedure,
			connect.WithSchema(directoryServiceLookupUserAccountMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteUser: connect.NewClient[v1alpha1.DeleteUserRequest, v1alpha1.DeleteUserResponse](
			httpClient,
			baseURL+DirectoryServiceDeleteUserProcedure,
			connect.WithSchema(directoryServiceDeleteUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// directoryServiceClient implements DirectoryServiceClient.
type directoryServiceClient struct {
	queryUsers         *connect.Client[v1alpha1.QueryUsersRequest, v1alpha1.QueryUsersResponse]
	getUser            *connect.Client[v1alpha1.GetUserRequest, v1alpha1.GetUserResponse]
	queryGroups        *connect.Client[v1alpha1.QueryGroupsRequest, v1alpha1.QueryGroupsResponse]
	getGroup           *connect.Client[v1alpha1.GetGroupRequest, v1alpha1.GetGroupResponse]
	queryGroupMembers  *connect.Client[v1alpha1.QueryGroupMembersRequest, v1alpha1.QueryGroupMembersResponse]
	queryChildGroups   *connect.Client[v1alpha1.QueryChildGroupsRequest, v1alpha1.QueryChildGroupsResponse]
	queryGroupsForUser *connect.Client[v1alpha1.QueryGroupsForUserRequest, v1alpha1.QueryGroupsForUserResponse]
	lookupUserAccount  *connect.Client[v1alpha1.LookupUserAccountRequest, v1alpha1.LookupUserAccountResponse]
	deleteUser         *connect.Client[v1alpha1.DeleteUserRequest, v1alpha1.DeleteUserResponse]
}

// QueryUsers calls commonfate.control.directory.v1alpha1.DirectoryService.QueryUsers.
func (c *directoryServiceClient) QueryUsers(ctx context.Context, req *connect.Request[v1alpha1.QueryUsersRequest]) (*connect.Response[v1alpha1.QueryUsersResponse], error) {
	return c.queryUsers.CallUnary(ctx, req)
}

// GetUser calls commonfate.control.directory.v1alpha1.DirectoryService.GetUser.
func (c *directoryServiceClient) GetUser(ctx context.Context, req *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// QueryGroups calls commonfate.control.directory.v1alpha1.DirectoryService.QueryGroups.
func (c *directoryServiceClient) QueryGroups(ctx context.Context, req *connect.Request[v1alpha1.QueryGroupsRequest]) (*connect.Response[v1alpha1.QueryGroupsResponse], error) {
	return c.queryGroups.CallUnary(ctx, req)
}

// GetGroup calls commonfate.control.directory.v1alpha1.DirectoryService.GetGroup.
func (c *directoryServiceClient) GetGroup(ctx context.Context, req *connect.Request[v1alpha1.GetGroupRequest]) (*connect.Response[v1alpha1.GetGroupResponse], error) {
	return c.getGroup.CallUnary(ctx, req)
}

// QueryGroupMembers calls commonfate.control.directory.v1alpha1.DirectoryService.QueryGroupMembers.
func (c *directoryServiceClient) QueryGroupMembers(ctx context.Context, req *connect.Request[v1alpha1.QueryGroupMembersRequest]) (*connect.Response[v1alpha1.QueryGroupMembersResponse], error) {
	return c.queryGroupMembers.CallUnary(ctx, req)
}

// QueryChildGroups calls commonfate.control.directory.v1alpha1.DirectoryService.QueryChildGroups.
func (c *directoryServiceClient) QueryChildGroups(ctx context.Context, req *connect.Request[v1alpha1.QueryChildGroupsRequest]) (*connect.Response[v1alpha1.QueryChildGroupsResponse], error) {
	return c.queryChildGroups.CallUnary(ctx, req)
}

// QueryGroupsForUser calls
// commonfate.control.directory.v1alpha1.DirectoryService.QueryGroupsForUser.
func (c *directoryServiceClient) QueryGroupsForUser(ctx context.Context, req *connect.Request[v1alpha1.QueryGroupsForUserRequest]) (*connect.Response[v1alpha1.QueryGroupsForUserResponse], error) {
	return c.queryGroupsForUser.CallUnary(ctx, req)
}

// LookupUserAccount calls commonfate.control.directory.v1alpha1.DirectoryService.LookupUserAccount.
func (c *directoryServiceClient) LookupUserAccount(ctx context.Context, req *connect.Request[v1alpha1.LookupUserAccountRequest]) (*connect.Response[v1alpha1.LookupUserAccountResponse], error) {
	return c.lookupUserAccount.CallUnary(ctx, req)
}

// DeleteUser calls commonfate.control.directory.v1alpha1.DirectoryService.DeleteUser.
func (c *directoryServiceClient) DeleteUser(ctx context.Context, req *connect.Request[v1alpha1.DeleteUserRequest]) (*connect.Response[v1alpha1.DeleteUserResponse], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// DirectoryServiceHandler is an implementation of the
// commonfate.control.directory.v1alpha1.DirectoryService service.
type DirectoryServiceHandler interface {
	QueryUsers(context.Context, *connect.Request[v1alpha1.QueryUsersRequest]) (*connect.Response[v1alpha1.QueryUsersResponse], error)
	GetUser(context.Context, *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error)
	QueryGroups(context.Context, *connect.Request[v1alpha1.QueryGroupsRequest]) (*connect.Response[v1alpha1.QueryGroupsResponse], error)
	GetGroup(context.Context, *connect.Request[v1alpha1.GetGroupRequest]) (*connect.Response[v1alpha1.GetGroupResponse], error)
	QueryGroupMembers(context.Context, *connect.Request[v1alpha1.QueryGroupMembersRequest]) (*connect.Response[v1alpha1.QueryGroupMembersResponse], error)
	QueryChildGroups(context.Context, *connect.Request[v1alpha1.QueryChildGroupsRequest]) (*connect.Response[v1alpha1.QueryChildGroupsResponse], error)
	// Lists the groups that a user is a member of.
	QueryGroupsForUser(context.Context, *connect.Request[v1alpha1.QueryGroupsForUserRequest]) (*connect.Response[v1alpha1.QueryGroupsForUserResponse], error)
	// Looks up the matching User for a particular
	// connected user account in an integration.
	LookupUserAccount(context.Context, *connect.Request[v1alpha1.LookupUserAccountRequest]) (*connect.Response[v1alpha1.LookupUserAccountResponse], error)
	DeleteUser(context.Context, *connect.Request[v1alpha1.DeleteUserRequest]) (*connect.Response[v1alpha1.DeleteUserResponse], error)
}

// NewDirectoryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDirectoryServiceHandler(svc DirectoryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	directoryServiceQueryUsersHandler := connect.NewUnaryHandler(
		DirectoryServiceQueryUsersProcedure,
		svc.QueryUsers,
		connect.WithSchema(directoryServiceQueryUsersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	directoryServiceGetUserHandler := connect.NewUnaryHandler(
		DirectoryServiceGetUserProcedure,
		svc.GetUser,
		connect.WithSchema(directoryServiceGetUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	directoryServiceQueryGroupsHandler := connect.NewUnaryHandler(
		DirectoryServiceQueryGroupsProcedure,
		svc.QueryGroups,
		connect.WithSchema(directoryServiceQueryGroupsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	directoryServiceGetGroupHandler := connect.NewUnaryHandler(
		DirectoryServiceGetGroupProcedure,
		svc.GetGroup,
		connect.WithSchema(directoryServiceGetGroupMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	directoryServiceQueryGroupMembersHandler := connect.NewUnaryHandler(
		DirectoryServiceQueryGroupMembersProcedure,
		svc.QueryGroupMembers,
		connect.WithSchema(directoryServiceQueryGroupMembersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	directoryServiceQueryChildGroupsHandler := connect.NewUnaryHandler(
		DirectoryServiceQueryChildGroupsProcedure,
		svc.QueryChildGroups,
		connect.WithSchema(directoryServiceQueryChildGroupsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	directoryServiceQueryGroupsForUserHandler := connect.NewUnaryHandler(
		DirectoryServiceQueryGroupsForUserProcedure,
		svc.QueryGroupsForUser,
		connect.WithSchema(directoryServiceQueryGroupsForUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	directoryServiceLookupUserAccountHandler := connect.NewUnaryHandler(
		DirectoryServiceLookupUserAccountProcedure,
		svc.LookupUserAccount,
		connect.WithSchema(directoryServiceLookupUserAccountMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	directoryServiceDeleteUserHandler := connect.NewUnaryHandler(
		DirectoryServiceDeleteUserProcedure,
		svc.DeleteUser,
		connect.WithSchema(directoryServiceDeleteUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.control.directory.v1alpha1.DirectoryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DirectoryServiceQueryUsersProcedure:
			directoryServiceQueryUsersHandler.ServeHTTP(w, r)
		case DirectoryServiceGetUserProcedure:
			directoryServiceGetUserHandler.ServeHTTP(w, r)
		case DirectoryServiceQueryGroupsProcedure:
			directoryServiceQueryGroupsHandler.ServeHTTP(w, r)
		case DirectoryServiceGetGroupProcedure:
			directoryServiceGetGroupHandler.ServeHTTP(w, r)
		case DirectoryServiceQueryGroupMembersProcedure:
			directoryServiceQueryGroupMembersHandler.ServeHTTP(w, r)
		case DirectoryServiceQueryChildGroupsProcedure:
			directoryServiceQueryChildGroupsHandler.ServeHTTP(w, r)
		case DirectoryServiceQueryGroupsForUserProcedure:
			directoryServiceQueryGroupsForUserHandler.ServeHTTP(w, r)
		case DirectoryServiceLookupUserAccountProcedure:
			directoryServiceLookupUserAccountHandler.ServeHTTP(w, r)
		case DirectoryServiceDeleteUserProcedure:
			directoryServiceDeleteUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDirectoryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDirectoryServiceHandler struct{}

func (UnimplementedDirectoryServiceHandler) QueryUsers(context.Context, *connect.Request[v1alpha1.QueryUsersRequest]) (*connect.Response[v1alpha1.QueryUsersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.directory.v1alpha1.DirectoryService.QueryUsers is not implemented"))
}

func (UnimplementedDirectoryServiceHandler) GetUser(context.Context, *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.directory.v1alpha1.DirectoryService.GetUser is not implemented"))
}

func (UnimplementedDirectoryServiceHandler) QueryGroups(context.Context, *connect.Request[v1alpha1.QueryGroupsRequest]) (*connect.Response[v1alpha1.QueryGroupsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.directory.v1alpha1.DirectoryService.QueryGroups is not implemented"))
}

func (UnimplementedDirectoryServiceHandler) GetGroup(context.Context, *connect.Request[v1alpha1.GetGroupRequest]) (*connect.Response[v1alpha1.GetGroupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.directory.v1alpha1.DirectoryService.GetGroup is not implemented"))
}

func (UnimplementedDirectoryServiceHandler) QueryGroupMembers(context.Context, *connect.Request[v1alpha1.QueryGroupMembersRequest]) (*connect.Response[v1alpha1.QueryGroupMembersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.directory.v1alpha1.DirectoryService.QueryGroupMembers is not implemented"))
}

func (UnimplementedDirectoryServiceHandler) QueryChildGroups(context.Context, *connect.Request[v1alpha1.QueryChildGroupsRequest]) (*connect.Response[v1alpha1.QueryChildGroupsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.directory.v1alpha1.DirectoryService.QueryChildGroups is not implemented"))
}

func (UnimplementedDirectoryServiceHandler) QueryGroupsForUser(context.Context, *connect.Request[v1alpha1.QueryGroupsForUserRequest]) (*connect.Response[v1alpha1.QueryGroupsForUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.directory.v1alpha1.DirectoryService.QueryGroupsForUser is not implemented"))
}

func (UnimplementedDirectoryServiceHandler) LookupUserAccount(context.Context, *connect.Request[v1alpha1.LookupUserAccountRequest]) (*connect.Response[v1alpha1.LookupUserAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.directory.v1alpha1.DirectoryService.LookupUserAccount is not implemented"))
}

func (UnimplementedDirectoryServiceHandler) DeleteUser(context.Context, *connect.Request[v1alpha1.DeleteUserRequest]) (*connect.Response[v1alpha1.DeleteUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.directory.v1alpha1.DirectoryService.DeleteUser is not implemented"))
}
