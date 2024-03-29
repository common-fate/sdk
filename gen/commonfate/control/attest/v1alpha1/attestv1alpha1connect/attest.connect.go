// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/attest/v1alpha1/attest.proto

package attestv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/control/attest/v1alpha1"
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
	// AttestServiceName is the fully-qualified name of the AttestService service.
	AttestServiceName = "commonfate.control.attest.v1alpha1.AttestService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AttestServiceRegisterDeviceProcedure is the fully-qualified name of the AttestService's
	// RegisterDevice RPC.
	AttestServiceRegisterDeviceProcedure = "/commonfate.control.attest.v1alpha1.AttestService/RegisterDevice"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	attestServiceServiceDescriptor              = v1alpha1.File_commonfate_control_attest_v1alpha1_attest_proto.Services().ByName("AttestService")
	attestServiceRegisterDeviceMethodDescriptor = attestServiceServiceDescriptor.Methods().ByName("RegisterDevice")
)

// AttestServiceClient is a client for the commonfate.control.attest.v1alpha1.AttestService service.
type AttestServiceClient interface {
	RegisterDevice(context.Context, *connect.Request[v1alpha1.RegisterDeviceRequest]) (*connect.Response[v1alpha1.RegisterDeviceResponse], error)
}

// NewAttestServiceClient constructs a client for the
// commonfate.control.attest.v1alpha1.AttestService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAttestServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AttestServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &attestServiceClient{
		registerDevice: connect.NewClient[v1alpha1.RegisterDeviceRequest, v1alpha1.RegisterDeviceResponse](
			httpClient,
			baseURL+AttestServiceRegisterDeviceProcedure,
			connect.WithSchema(attestServiceRegisterDeviceMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// attestServiceClient implements AttestServiceClient.
type attestServiceClient struct {
	registerDevice *connect.Client[v1alpha1.RegisterDeviceRequest, v1alpha1.RegisterDeviceResponse]
}

// RegisterDevice calls commonfate.control.attest.v1alpha1.AttestService.RegisterDevice.
func (c *attestServiceClient) RegisterDevice(ctx context.Context, req *connect.Request[v1alpha1.RegisterDeviceRequest]) (*connect.Response[v1alpha1.RegisterDeviceResponse], error) {
	return c.registerDevice.CallUnary(ctx, req)
}

// AttestServiceHandler is an implementation of the commonfate.control.attest.v1alpha1.AttestService
// service.
type AttestServiceHandler interface {
	RegisterDevice(context.Context, *connect.Request[v1alpha1.RegisterDeviceRequest]) (*connect.Response[v1alpha1.RegisterDeviceResponse], error)
}

// NewAttestServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAttestServiceHandler(svc AttestServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	attestServiceRegisterDeviceHandler := connect.NewUnaryHandler(
		AttestServiceRegisterDeviceProcedure,
		svc.RegisterDevice,
		connect.WithSchema(attestServiceRegisterDeviceMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.control.attest.v1alpha1.AttestService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AttestServiceRegisterDeviceProcedure:
			attestServiceRegisterDeviceHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAttestServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAttestServiceHandler struct{}

func (UnimplementedAttestServiceHandler) RegisterDevice(context.Context, *connect.Request[v1alpha1.RegisterDeviceRequest]) (*connect.Response[v1alpha1.RegisterDeviceResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.attest.v1alpha1.AttestService.RegisterDevice is not implemented"))
}
