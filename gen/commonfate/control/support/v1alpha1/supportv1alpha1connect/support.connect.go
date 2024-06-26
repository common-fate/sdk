// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/support/v1alpha1/support.proto

package supportv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/control/support/v1alpha1"
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
	// SupportServiceName is the fully-qualified name of the SupportService service.
	SupportServiceName = "commonfate.control.support.v1alpha1.SupportService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SupportServiceContactProcedure is the fully-qualified name of the SupportService's Contact RPC.
	SupportServiceContactProcedure = "/commonfate.control.support.v1alpha1.SupportService/Contact"
	// SupportServiceCreateAttachmentProcedure is the fully-qualified name of the SupportService's
	// CreateAttachment RPC.
	SupportServiceCreateAttachmentProcedure = "/commonfate.control.support.v1alpha1.SupportService/CreateAttachment"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	supportServiceServiceDescriptor                = v1alpha1.File_commonfate_control_support_v1alpha1_support_proto.Services().ByName("SupportService")
	supportServiceContactMethodDescriptor          = supportServiceServiceDescriptor.Methods().ByName("Contact")
	supportServiceCreateAttachmentMethodDescriptor = supportServiceServiceDescriptor.Methods().ByName("CreateAttachment")
)

// SupportServiceClient is a client for the commonfate.control.support.v1alpha1.SupportService
// service.
type SupportServiceClient interface {
	// Contact Common Fate support.
	Contact(context.Context, *connect.Request[v1alpha1.ContactRequest]) (*connect.Response[v1alpha1.ContactResponse], error)
	// Create an attachment to add to a support ticket.
	CreateAttachment(context.Context, *connect.Request[v1alpha1.CreateAttachmentRequest]) (*connect.Response[v1alpha1.CreateAttachmentResponse], error)
}

// NewSupportServiceClient constructs a client for the
// commonfate.control.support.v1alpha1.SupportService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSupportServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SupportServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &supportServiceClient{
		contact: connect.NewClient[v1alpha1.ContactRequest, v1alpha1.ContactResponse](
			httpClient,
			baseURL+SupportServiceContactProcedure,
			connect.WithSchema(supportServiceContactMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createAttachment: connect.NewClient[v1alpha1.CreateAttachmentRequest, v1alpha1.CreateAttachmentResponse](
			httpClient,
			baseURL+SupportServiceCreateAttachmentProcedure,
			connect.WithSchema(supportServiceCreateAttachmentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// supportServiceClient implements SupportServiceClient.
type supportServiceClient struct {
	contact          *connect.Client[v1alpha1.ContactRequest, v1alpha1.ContactResponse]
	createAttachment *connect.Client[v1alpha1.CreateAttachmentRequest, v1alpha1.CreateAttachmentResponse]
}

// Contact calls commonfate.control.support.v1alpha1.SupportService.Contact.
func (c *supportServiceClient) Contact(ctx context.Context, req *connect.Request[v1alpha1.ContactRequest]) (*connect.Response[v1alpha1.ContactResponse], error) {
	return c.contact.CallUnary(ctx, req)
}

// CreateAttachment calls commonfate.control.support.v1alpha1.SupportService.CreateAttachment.
func (c *supportServiceClient) CreateAttachment(ctx context.Context, req *connect.Request[v1alpha1.CreateAttachmentRequest]) (*connect.Response[v1alpha1.CreateAttachmentResponse], error) {
	return c.createAttachment.CallUnary(ctx, req)
}

// SupportServiceHandler is an implementation of the
// commonfate.control.support.v1alpha1.SupportService service.
type SupportServiceHandler interface {
	// Contact Common Fate support.
	Contact(context.Context, *connect.Request[v1alpha1.ContactRequest]) (*connect.Response[v1alpha1.ContactResponse], error)
	// Create an attachment to add to a support ticket.
	CreateAttachment(context.Context, *connect.Request[v1alpha1.CreateAttachmentRequest]) (*connect.Response[v1alpha1.CreateAttachmentResponse], error)
}

// NewSupportServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSupportServiceHandler(svc SupportServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	supportServiceContactHandler := connect.NewUnaryHandler(
		SupportServiceContactProcedure,
		svc.Contact,
		connect.WithSchema(supportServiceContactMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	supportServiceCreateAttachmentHandler := connect.NewUnaryHandler(
		SupportServiceCreateAttachmentProcedure,
		svc.CreateAttachment,
		connect.WithSchema(supportServiceCreateAttachmentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.control.support.v1alpha1.SupportService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SupportServiceContactProcedure:
			supportServiceContactHandler.ServeHTTP(w, r)
		case SupportServiceCreateAttachmentProcedure:
			supportServiceCreateAttachmentHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSupportServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSupportServiceHandler struct{}

func (UnimplementedSupportServiceHandler) Contact(context.Context, *connect.Request[v1alpha1.ContactRequest]) (*connect.Response[v1alpha1.ContactResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.support.v1alpha1.SupportService.Contact is not implemented"))
}

func (UnimplementedSupportServiceHandler) CreateAttachment(context.Context, *connect.Request[v1alpha1.CreateAttachmentRequest]) (*connect.Response[v1alpha1.CreateAttachmentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.support.v1alpha1.SupportService.CreateAttachment is not implemented"))
}
