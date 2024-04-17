// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: granted/registry/aws/v1alpha1/aws.proto

package awsv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/granted/registry/aws/v1alpha1"
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
	// ProfileRegistryServiceName is the fully-qualified name of the ProfileRegistryService service.
	ProfileRegistryServiceName = "granted.registry.aws.v1alpha1.ProfileRegistryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ProfileRegistryServiceGetProfileForAccountAndRoleProcedure is the fully-qualified name of the
	// ProfileRegistryService's GetProfileForAccountAndRole RPC.
	ProfileRegistryServiceGetProfileForAccountAndRoleProcedure = "/granted.registry.aws.v1alpha1.ProfileRegistryService/GetProfileForAccountAndRole"
	// ProfileRegistryServiceListProfilesProcedure is the fully-qualified name of the
	// ProfileRegistryService's ListProfiles RPC.
	ProfileRegistryServiceListProfilesProcedure = "/granted.registry.aws.v1alpha1.ProfileRegistryService/ListProfiles"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	profileRegistryServiceServiceDescriptor                           = v1alpha1.File_granted_registry_aws_v1alpha1_aws_proto.Services().ByName("ProfileRegistryService")
	profileRegistryServiceGetProfileForAccountAndRoleMethodDescriptor = profileRegistryServiceServiceDescriptor.Methods().ByName("GetProfileForAccountAndRole")
	profileRegistryServiceListProfilesMethodDescriptor                = profileRegistryServiceServiceDescriptor.Methods().ByName("ListProfiles")
)

// ProfileRegistryServiceClient is a client for the
// granted.registry.aws.v1alpha1.ProfileRegistryService service.
type ProfileRegistryServiceClient interface {
	// Looks up a profile based on a provided AWS account and role.
	GetProfileForAccountAndRole(context.Context, *connect.Request[v1alpha1.GetProfileForAccountAndRoleRequest]) (*connect.Response[v1alpha1.GetProfileForAccountAndRoleResponse], error)
	// Lists all profiles available to the user.
	ListProfiles(context.Context, *connect.Request[v1alpha1.ListProfilesRequest]) (*connect.Response[v1alpha1.ListProfilesResponse], error)
}

// NewProfileRegistryServiceClient constructs a client for the
// granted.registry.aws.v1alpha1.ProfileRegistryService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProfileRegistryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ProfileRegistryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &profileRegistryServiceClient{
		getProfileForAccountAndRole: connect.NewClient[v1alpha1.GetProfileForAccountAndRoleRequest, v1alpha1.GetProfileForAccountAndRoleResponse](
			httpClient,
			baseURL+ProfileRegistryServiceGetProfileForAccountAndRoleProcedure,
			connect.WithSchema(profileRegistryServiceGetProfileForAccountAndRoleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listProfiles: connect.NewClient[v1alpha1.ListProfilesRequest, v1alpha1.ListProfilesResponse](
			httpClient,
			baseURL+ProfileRegistryServiceListProfilesProcedure,
			connect.WithSchema(profileRegistryServiceListProfilesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// profileRegistryServiceClient implements ProfileRegistryServiceClient.
type profileRegistryServiceClient struct {
	getProfileForAccountAndRole *connect.Client[v1alpha1.GetProfileForAccountAndRoleRequest, v1alpha1.GetProfileForAccountAndRoleResponse]
	listProfiles                *connect.Client[v1alpha1.ListProfilesRequest, v1alpha1.ListProfilesResponse]
}

// GetProfileForAccountAndRole calls
// granted.registry.aws.v1alpha1.ProfileRegistryService.GetProfileForAccountAndRole.
func (c *profileRegistryServiceClient) GetProfileForAccountAndRole(ctx context.Context, req *connect.Request[v1alpha1.GetProfileForAccountAndRoleRequest]) (*connect.Response[v1alpha1.GetProfileForAccountAndRoleResponse], error) {
	return c.getProfileForAccountAndRole.CallUnary(ctx, req)
}

// ListProfiles calls granted.registry.aws.v1alpha1.ProfileRegistryService.ListProfiles.
func (c *profileRegistryServiceClient) ListProfiles(ctx context.Context, req *connect.Request[v1alpha1.ListProfilesRequest]) (*connect.Response[v1alpha1.ListProfilesResponse], error) {
	return c.listProfiles.CallUnary(ctx, req)
}

// ProfileRegistryServiceHandler is an implementation of the
// granted.registry.aws.v1alpha1.ProfileRegistryService service.
type ProfileRegistryServiceHandler interface {
	// Looks up a profile based on a provided AWS account and role.
	GetProfileForAccountAndRole(context.Context, *connect.Request[v1alpha1.GetProfileForAccountAndRoleRequest]) (*connect.Response[v1alpha1.GetProfileForAccountAndRoleResponse], error)
	// Lists all profiles available to the user.
	ListProfiles(context.Context, *connect.Request[v1alpha1.ListProfilesRequest]) (*connect.Response[v1alpha1.ListProfilesResponse], error)
}

// NewProfileRegistryServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProfileRegistryServiceHandler(svc ProfileRegistryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	profileRegistryServiceGetProfileForAccountAndRoleHandler := connect.NewUnaryHandler(
		ProfileRegistryServiceGetProfileForAccountAndRoleProcedure,
		svc.GetProfileForAccountAndRole,
		connect.WithSchema(profileRegistryServiceGetProfileForAccountAndRoleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	profileRegistryServiceListProfilesHandler := connect.NewUnaryHandler(
		ProfileRegistryServiceListProfilesProcedure,
		svc.ListProfiles,
		connect.WithSchema(profileRegistryServiceListProfilesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/granted.registry.aws.v1alpha1.ProfileRegistryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ProfileRegistryServiceGetProfileForAccountAndRoleProcedure:
			profileRegistryServiceGetProfileForAccountAndRoleHandler.ServeHTTP(w, r)
		case ProfileRegistryServiceListProfilesProcedure:
			profileRegistryServiceListProfilesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedProfileRegistryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedProfileRegistryServiceHandler struct{}

func (UnimplementedProfileRegistryServiceHandler) GetProfileForAccountAndRole(context.Context, *connect.Request[v1alpha1.GetProfileForAccountAndRoleRequest]) (*connect.Response[v1alpha1.GetProfileForAccountAndRoleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("granted.registry.aws.v1alpha1.ProfileRegistryService.GetProfileForAccountAndRole is not implemented"))
}

func (UnimplementedProfileRegistryServiceHandler) ListProfiles(context.Context, *connect.Request[v1alpha1.ListProfilesRequest]) (*connect.Response[v1alpha1.ListProfilesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("granted.registry.aws.v1alpha1.ProfileRegistryService.ListProfiles is not implemented"))
}