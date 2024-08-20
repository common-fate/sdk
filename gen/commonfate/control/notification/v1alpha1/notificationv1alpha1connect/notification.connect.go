// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/notification/v1alpha1/notification.proto

package notificationv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/control/notification/v1alpha1"
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
	// UserNotificationSettingsServiceName is the fully-qualified name of the
	// UserNotificationSettingsService service.
	UserNotificationSettingsServiceName = "commonfate.control.notification.v1alpha1.UserNotificationSettingsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserNotificationSettingsServiceGetUserUserNotificationSettingsProcedure is the fully-qualified
	// name of the UserNotificationSettingsService's GetUserUserNotificationSettings RPC.
	UserNotificationSettingsServiceGetUserUserNotificationSettingsProcedure = "/commonfate.control.notification.v1alpha1.UserNotificationSettingsService/GetUserUserNotificationSettings"
	// UserNotificationSettingsServiceUpdateUserNotificationSettingsProcedure is the fully-qualified
	// name of the UserNotificationSettingsService's UpdateUserNotificationSettings RPC.
	UserNotificationSettingsServiceUpdateUserNotificationSettingsProcedure = "/commonfate.control.notification.v1alpha1.UserNotificationSettingsService/UpdateUserNotificationSettings"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userNotificationSettingsServiceServiceDescriptor                               = v1alpha1.File_commonfate_control_notification_v1alpha1_notification_proto.Services().ByName("UserNotificationSettingsService")
	userNotificationSettingsServiceGetUserUserNotificationSettingsMethodDescriptor = userNotificationSettingsServiceServiceDescriptor.Methods().ByName("GetUserUserNotificationSettings")
	userNotificationSettingsServiceUpdateUserNotificationSettingsMethodDescriptor  = userNotificationSettingsServiceServiceDescriptor.Methods().ByName("UpdateUserNotificationSettings")
)

// UserNotificationSettingsServiceClient is a client for the
// commonfate.control.notification.v1alpha1.UserNotificationSettingsService service.
type UserNotificationSettingsServiceClient interface {
	// Returns a list of enabled user_notification_settings.
	GetUserUserNotificationSettings(context.Context, *connect.Request[v1alpha1.GetUserUserNotificationSettingsRequest]) (*connect.Response[v1alpha1.GetUserUserNotificationSettingsResponse], error)
	UpdateUserNotificationSettings(context.Context, *connect.Request[v1alpha1.UpdateUserNotificationSettingsRequest]) (*connect.Response[v1alpha1.UpdateUserNotificationSettingsResponse], error)
}

// NewUserNotificationSettingsServiceClient constructs a client for the
// commonfate.control.notification.v1alpha1.UserNotificationSettingsService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserNotificationSettingsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserNotificationSettingsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userNotificationSettingsServiceClient{
		getUserUserNotificationSettings: connect.NewClient[v1alpha1.GetUserUserNotificationSettingsRequest, v1alpha1.GetUserUserNotificationSettingsResponse](
			httpClient,
			baseURL+UserNotificationSettingsServiceGetUserUserNotificationSettingsProcedure,
			connect.WithSchema(userNotificationSettingsServiceGetUserUserNotificationSettingsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateUserNotificationSettings: connect.NewClient[v1alpha1.UpdateUserNotificationSettingsRequest, v1alpha1.UpdateUserNotificationSettingsResponse](
			httpClient,
			baseURL+UserNotificationSettingsServiceUpdateUserNotificationSettingsProcedure,
			connect.WithSchema(userNotificationSettingsServiceUpdateUserNotificationSettingsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userNotificationSettingsServiceClient implements UserNotificationSettingsServiceClient.
type userNotificationSettingsServiceClient struct {
	getUserUserNotificationSettings *connect.Client[v1alpha1.GetUserUserNotificationSettingsRequest, v1alpha1.GetUserUserNotificationSettingsResponse]
	updateUserNotificationSettings  *connect.Client[v1alpha1.UpdateUserNotificationSettingsRequest, v1alpha1.UpdateUserNotificationSettingsResponse]
}

// GetUserUserNotificationSettings calls
// commonfate.control.notification.v1alpha1.UserNotificationSettingsService.GetUserUserNotificationSettings.
func (c *userNotificationSettingsServiceClient) GetUserUserNotificationSettings(ctx context.Context, req *connect.Request[v1alpha1.GetUserUserNotificationSettingsRequest]) (*connect.Response[v1alpha1.GetUserUserNotificationSettingsResponse], error) {
	return c.getUserUserNotificationSettings.CallUnary(ctx, req)
}

// UpdateUserNotificationSettings calls
// commonfate.control.notification.v1alpha1.UserNotificationSettingsService.UpdateUserNotificationSettings.
func (c *userNotificationSettingsServiceClient) UpdateUserNotificationSettings(ctx context.Context, req *connect.Request[v1alpha1.UpdateUserNotificationSettingsRequest]) (*connect.Response[v1alpha1.UpdateUserNotificationSettingsResponse], error) {
	return c.updateUserNotificationSettings.CallUnary(ctx, req)
}

// UserNotificationSettingsServiceHandler is an implementation of the
// commonfate.control.notification.v1alpha1.UserNotificationSettingsService service.
type UserNotificationSettingsServiceHandler interface {
	// Returns a list of enabled user_notification_settings.
	GetUserUserNotificationSettings(context.Context, *connect.Request[v1alpha1.GetUserUserNotificationSettingsRequest]) (*connect.Response[v1alpha1.GetUserUserNotificationSettingsResponse], error)
	UpdateUserNotificationSettings(context.Context, *connect.Request[v1alpha1.UpdateUserNotificationSettingsRequest]) (*connect.Response[v1alpha1.UpdateUserNotificationSettingsResponse], error)
}

// NewUserNotificationSettingsServiceHandler builds an HTTP handler from the service implementation.
// It returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserNotificationSettingsServiceHandler(svc UserNotificationSettingsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userNotificationSettingsServiceGetUserUserNotificationSettingsHandler := connect.NewUnaryHandler(
		UserNotificationSettingsServiceGetUserUserNotificationSettingsProcedure,
		svc.GetUserUserNotificationSettings,
		connect.WithSchema(userNotificationSettingsServiceGetUserUserNotificationSettingsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userNotificationSettingsServiceUpdateUserNotificationSettingsHandler := connect.NewUnaryHandler(
		UserNotificationSettingsServiceUpdateUserNotificationSettingsProcedure,
		svc.UpdateUserNotificationSettings,
		connect.WithSchema(userNotificationSettingsServiceUpdateUserNotificationSettingsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.control.notification.v1alpha1.UserNotificationSettingsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserNotificationSettingsServiceGetUserUserNotificationSettingsProcedure:
			userNotificationSettingsServiceGetUserUserNotificationSettingsHandler.ServeHTTP(w, r)
		case UserNotificationSettingsServiceUpdateUserNotificationSettingsProcedure:
			userNotificationSettingsServiceUpdateUserNotificationSettingsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserNotificationSettingsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserNotificationSettingsServiceHandler struct{}

func (UnimplementedUserNotificationSettingsServiceHandler) GetUserUserNotificationSettings(context.Context, *connect.Request[v1alpha1.GetUserUserNotificationSettingsRequest]) (*connect.Response[v1alpha1.GetUserUserNotificationSettingsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.notification.v1alpha1.UserNotificationSettingsService.GetUserUserNotificationSettings is not implemented"))
}

func (UnimplementedUserNotificationSettingsServiceHandler) UpdateUserNotificationSettings(context.Context, *connect.Request[v1alpha1.UpdateUserNotificationSettingsRequest]) (*connect.Response[v1alpha1.UpdateUserNotificationSettingsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.notification.v1alpha1.UserNotificationSettingsService.UpdateUserNotificationSettings is not implemented"))
}
