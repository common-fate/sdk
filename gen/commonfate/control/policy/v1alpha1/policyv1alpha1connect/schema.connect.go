// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/policy/v1alpha1/schema.proto

package policyv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/control/policy/v1alpha1"
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
	// SchemaServiceName is the fully-qualified name of the SchemaService service.
	SchemaServiceName = "commonfate.control.policy.v1alpha1.SchemaService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SchemaServiceGetSchemaJSONStringProcedure is the fully-qualified name of the SchemaService's
	// GetSchemaJSONString RPC.
	SchemaServiceGetSchemaJSONStringProcedure = "/commonfate.control.policy.v1alpha1.SchemaService/GetSchemaJSONString"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	schemaServiceServiceDescriptor                   = v1alpha1.File_commonfate_control_policy_v1alpha1_schema_proto.Services().ByName("SchemaService")
	schemaServiceGetSchemaJSONStringMethodDescriptor = schemaServiceServiceDescriptor.Methods().ByName("GetSchemaJSONString")
)

// SchemaServiceClient is a client for the commonfate.control.policy.v1alpha1.SchemaService service.
type SchemaServiceClient interface {
	// Retrieves a copy of the Cedar schema in JSON format, as a string.
	GetSchemaJSONString(context.Context, *connect.Request[v1alpha1.GetSchemaJSONStringRequest]) (*connect.Response[v1alpha1.GetSchemaJSONStringResponse], error)
}

// NewSchemaServiceClient constructs a client for the
// commonfate.control.policy.v1alpha1.SchemaService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSchemaServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SchemaServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &schemaServiceClient{
		getSchemaJSONString: connect.NewClient[v1alpha1.GetSchemaJSONStringRequest, v1alpha1.GetSchemaJSONStringResponse](
			httpClient,
			baseURL+SchemaServiceGetSchemaJSONStringProcedure,
			connect.WithSchema(schemaServiceGetSchemaJSONStringMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// schemaServiceClient implements SchemaServiceClient.
type schemaServiceClient struct {
	getSchemaJSONString *connect.Client[v1alpha1.GetSchemaJSONStringRequest, v1alpha1.GetSchemaJSONStringResponse]
}

// GetSchemaJSONString calls commonfate.control.policy.v1alpha1.SchemaService.GetSchemaJSONString.
func (c *schemaServiceClient) GetSchemaJSONString(ctx context.Context, req *connect.Request[v1alpha1.GetSchemaJSONStringRequest]) (*connect.Response[v1alpha1.GetSchemaJSONStringResponse], error) {
	return c.getSchemaJSONString.CallUnary(ctx, req)
}

// SchemaServiceHandler is an implementation of the commonfate.control.policy.v1alpha1.SchemaService
// service.
type SchemaServiceHandler interface {
	// Retrieves a copy of the Cedar schema in JSON format, as a string.
	GetSchemaJSONString(context.Context, *connect.Request[v1alpha1.GetSchemaJSONStringRequest]) (*connect.Response[v1alpha1.GetSchemaJSONStringResponse], error)
}

// NewSchemaServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSchemaServiceHandler(svc SchemaServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	schemaServiceGetSchemaJSONStringHandler := connect.NewUnaryHandler(
		SchemaServiceGetSchemaJSONStringProcedure,
		svc.GetSchemaJSONString,
		connect.WithSchema(schemaServiceGetSchemaJSONStringMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.control.policy.v1alpha1.SchemaService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SchemaServiceGetSchemaJSONStringProcedure:
			schemaServiceGetSchemaJSONStringHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSchemaServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSchemaServiceHandler struct{}

func (UnimplementedSchemaServiceHandler) GetSchemaJSONString(context.Context, *connect.Request[v1alpha1.GetSchemaJSONStringRequest]) (*connect.Response[v1alpha1.GetSchemaJSONStringResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.policy.v1alpha1.SchemaService.GetSchemaJSONString is not implemented"))
}
