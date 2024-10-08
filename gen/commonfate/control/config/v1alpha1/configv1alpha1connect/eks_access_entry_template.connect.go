// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: commonfate/control/config/v1alpha1/eks_access_entry_template.proto

package configv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/control/config/v1alpha1"
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
	// EKSAccessEntryTemplateServiceName is the fully-qualified name of the
	// EKSAccessEntryTemplateService service.
	EKSAccessEntryTemplateServiceName = "commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EKSAccessEntryTemplateServiceCreateEKSAccessEntryTemplateProcedure is the fully-qualified name of
	// the EKSAccessEntryTemplateService's CreateEKSAccessEntryTemplate RPC.
	EKSAccessEntryTemplateServiceCreateEKSAccessEntryTemplateProcedure = "/commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService/CreateEKSAccessEntryTemplate"
	// EKSAccessEntryTemplateServiceGetEKSAccessEntryTemplateProcedure is the fully-qualified name of
	// the EKSAccessEntryTemplateService's GetEKSAccessEntryTemplate RPC.
	EKSAccessEntryTemplateServiceGetEKSAccessEntryTemplateProcedure = "/commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService/GetEKSAccessEntryTemplate"
	// EKSAccessEntryTemplateServiceUpdateEKSAccessEntryTemplateProcedure is the fully-qualified name of
	// the EKSAccessEntryTemplateService's UpdateEKSAccessEntryTemplate RPC.
	EKSAccessEntryTemplateServiceUpdateEKSAccessEntryTemplateProcedure = "/commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService/UpdateEKSAccessEntryTemplate"
	// EKSAccessEntryTemplateServiceDeleteEKSAccessEntryTemplateProcedure is the fully-qualified name of
	// the EKSAccessEntryTemplateService's DeleteEKSAccessEntryTemplate RPC.
	EKSAccessEntryTemplateServiceDeleteEKSAccessEntryTemplateProcedure = "/commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService/DeleteEKSAccessEntryTemplate"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	eKSAccessEntryTemplateServiceServiceDescriptor                            = v1alpha1.File_commonfate_control_config_v1alpha1_eks_access_entry_template_proto.Services().ByName("EKSAccessEntryTemplateService")
	eKSAccessEntryTemplateServiceCreateEKSAccessEntryTemplateMethodDescriptor = eKSAccessEntryTemplateServiceServiceDescriptor.Methods().ByName("CreateEKSAccessEntryTemplate")
	eKSAccessEntryTemplateServiceGetEKSAccessEntryTemplateMethodDescriptor    = eKSAccessEntryTemplateServiceServiceDescriptor.Methods().ByName("GetEKSAccessEntryTemplate")
	eKSAccessEntryTemplateServiceUpdateEKSAccessEntryTemplateMethodDescriptor = eKSAccessEntryTemplateServiceServiceDescriptor.Methods().ByName("UpdateEKSAccessEntryTemplate")
	eKSAccessEntryTemplateServiceDeleteEKSAccessEntryTemplateMethodDescriptor = eKSAccessEntryTemplateServiceServiceDescriptor.Methods().ByName("DeleteEKSAccessEntryTemplate")
)

// EKSAccessEntryTemplateServiceClient is a client for the
// commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService service.
type EKSAccessEntryTemplateServiceClient interface {
	CreateEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.CreateEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.CreateEKSAccessEntryTemplateResponse], error)
	GetEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.GetEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.GetEKSAccessEntryTemplateResponse], error)
	UpdateEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.UpdateEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.UpdateEKSAccessEntryTemplateResponse], error)
	DeleteEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.DeleteEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.DeleteEKSAccessEntryTemplateResponse], error)
}

// NewEKSAccessEntryTemplateServiceClient constructs a client for the
// commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEKSAccessEntryTemplateServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) EKSAccessEntryTemplateServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &eKSAccessEntryTemplateServiceClient{
		createEKSAccessEntryTemplate: connect.NewClient[v1alpha1.CreateEKSAccessEntryTemplateRequest, v1alpha1.CreateEKSAccessEntryTemplateResponse](
			httpClient,
			baseURL+EKSAccessEntryTemplateServiceCreateEKSAccessEntryTemplateProcedure,
			connect.WithSchema(eKSAccessEntryTemplateServiceCreateEKSAccessEntryTemplateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getEKSAccessEntryTemplate: connect.NewClient[v1alpha1.GetEKSAccessEntryTemplateRequest, v1alpha1.GetEKSAccessEntryTemplateResponse](
			httpClient,
			baseURL+EKSAccessEntryTemplateServiceGetEKSAccessEntryTemplateProcedure,
			connect.WithSchema(eKSAccessEntryTemplateServiceGetEKSAccessEntryTemplateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateEKSAccessEntryTemplate: connect.NewClient[v1alpha1.UpdateEKSAccessEntryTemplateRequest, v1alpha1.UpdateEKSAccessEntryTemplateResponse](
			httpClient,
			baseURL+EKSAccessEntryTemplateServiceUpdateEKSAccessEntryTemplateProcedure,
			connect.WithSchema(eKSAccessEntryTemplateServiceUpdateEKSAccessEntryTemplateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteEKSAccessEntryTemplate: connect.NewClient[v1alpha1.DeleteEKSAccessEntryTemplateRequest, v1alpha1.DeleteEKSAccessEntryTemplateResponse](
			httpClient,
			baseURL+EKSAccessEntryTemplateServiceDeleteEKSAccessEntryTemplateProcedure,
			connect.WithSchema(eKSAccessEntryTemplateServiceDeleteEKSAccessEntryTemplateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// eKSAccessEntryTemplateServiceClient implements EKSAccessEntryTemplateServiceClient.
type eKSAccessEntryTemplateServiceClient struct {
	createEKSAccessEntryTemplate *connect.Client[v1alpha1.CreateEKSAccessEntryTemplateRequest, v1alpha1.CreateEKSAccessEntryTemplateResponse]
	getEKSAccessEntryTemplate    *connect.Client[v1alpha1.GetEKSAccessEntryTemplateRequest, v1alpha1.GetEKSAccessEntryTemplateResponse]
	updateEKSAccessEntryTemplate *connect.Client[v1alpha1.UpdateEKSAccessEntryTemplateRequest, v1alpha1.UpdateEKSAccessEntryTemplateResponse]
	deleteEKSAccessEntryTemplate *connect.Client[v1alpha1.DeleteEKSAccessEntryTemplateRequest, v1alpha1.DeleteEKSAccessEntryTemplateResponse]
}

// CreateEKSAccessEntryTemplate calls
// commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService.CreateEKSAccessEntryTemplate.
func (c *eKSAccessEntryTemplateServiceClient) CreateEKSAccessEntryTemplate(ctx context.Context, req *connect.Request[v1alpha1.CreateEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.CreateEKSAccessEntryTemplateResponse], error) {
	return c.createEKSAccessEntryTemplate.CallUnary(ctx, req)
}

// GetEKSAccessEntryTemplate calls
// commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService.GetEKSAccessEntryTemplate.
func (c *eKSAccessEntryTemplateServiceClient) GetEKSAccessEntryTemplate(ctx context.Context, req *connect.Request[v1alpha1.GetEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.GetEKSAccessEntryTemplateResponse], error) {
	return c.getEKSAccessEntryTemplate.CallUnary(ctx, req)
}

// UpdateEKSAccessEntryTemplate calls
// commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService.UpdateEKSAccessEntryTemplate.
func (c *eKSAccessEntryTemplateServiceClient) UpdateEKSAccessEntryTemplate(ctx context.Context, req *connect.Request[v1alpha1.UpdateEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.UpdateEKSAccessEntryTemplateResponse], error) {
	return c.updateEKSAccessEntryTemplate.CallUnary(ctx, req)
}

// DeleteEKSAccessEntryTemplate calls
// commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService.DeleteEKSAccessEntryTemplate.
func (c *eKSAccessEntryTemplateServiceClient) DeleteEKSAccessEntryTemplate(ctx context.Context, req *connect.Request[v1alpha1.DeleteEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.DeleteEKSAccessEntryTemplateResponse], error) {
	return c.deleteEKSAccessEntryTemplate.CallUnary(ctx, req)
}

// EKSAccessEntryTemplateServiceHandler is an implementation of the
// commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService service.
type EKSAccessEntryTemplateServiceHandler interface {
	CreateEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.CreateEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.CreateEKSAccessEntryTemplateResponse], error)
	GetEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.GetEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.GetEKSAccessEntryTemplateResponse], error)
	UpdateEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.UpdateEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.UpdateEKSAccessEntryTemplateResponse], error)
	DeleteEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.DeleteEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.DeleteEKSAccessEntryTemplateResponse], error)
}

// NewEKSAccessEntryTemplateServiceHandler builds an HTTP handler from the service implementation.
// It returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEKSAccessEntryTemplateServiceHandler(svc EKSAccessEntryTemplateServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	eKSAccessEntryTemplateServiceCreateEKSAccessEntryTemplateHandler := connect.NewUnaryHandler(
		EKSAccessEntryTemplateServiceCreateEKSAccessEntryTemplateProcedure,
		svc.CreateEKSAccessEntryTemplate,
		connect.WithSchema(eKSAccessEntryTemplateServiceCreateEKSAccessEntryTemplateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	eKSAccessEntryTemplateServiceGetEKSAccessEntryTemplateHandler := connect.NewUnaryHandler(
		EKSAccessEntryTemplateServiceGetEKSAccessEntryTemplateProcedure,
		svc.GetEKSAccessEntryTemplate,
		connect.WithSchema(eKSAccessEntryTemplateServiceGetEKSAccessEntryTemplateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	eKSAccessEntryTemplateServiceUpdateEKSAccessEntryTemplateHandler := connect.NewUnaryHandler(
		EKSAccessEntryTemplateServiceUpdateEKSAccessEntryTemplateProcedure,
		svc.UpdateEKSAccessEntryTemplate,
		connect.WithSchema(eKSAccessEntryTemplateServiceUpdateEKSAccessEntryTemplateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	eKSAccessEntryTemplateServiceDeleteEKSAccessEntryTemplateHandler := connect.NewUnaryHandler(
		EKSAccessEntryTemplateServiceDeleteEKSAccessEntryTemplateProcedure,
		svc.DeleteEKSAccessEntryTemplate,
		connect.WithSchema(eKSAccessEntryTemplateServiceDeleteEKSAccessEntryTemplateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case EKSAccessEntryTemplateServiceCreateEKSAccessEntryTemplateProcedure:
			eKSAccessEntryTemplateServiceCreateEKSAccessEntryTemplateHandler.ServeHTTP(w, r)
		case EKSAccessEntryTemplateServiceGetEKSAccessEntryTemplateProcedure:
			eKSAccessEntryTemplateServiceGetEKSAccessEntryTemplateHandler.ServeHTTP(w, r)
		case EKSAccessEntryTemplateServiceUpdateEKSAccessEntryTemplateProcedure:
			eKSAccessEntryTemplateServiceUpdateEKSAccessEntryTemplateHandler.ServeHTTP(w, r)
		case EKSAccessEntryTemplateServiceDeleteEKSAccessEntryTemplateProcedure:
			eKSAccessEntryTemplateServiceDeleteEKSAccessEntryTemplateHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedEKSAccessEntryTemplateServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEKSAccessEntryTemplateServiceHandler struct{}

func (UnimplementedEKSAccessEntryTemplateServiceHandler) CreateEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.CreateEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.CreateEKSAccessEntryTemplateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService.CreateEKSAccessEntryTemplate is not implemented"))
}

func (UnimplementedEKSAccessEntryTemplateServiceHandler) GetEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.GetEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.GetEKSAccessEntryTemplateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService.GetEKSAccessEntryTemplate is not implemented"))
}

func (UnimplementedEKSAccessEntryTemplateServiceHandler) UpdateEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.UpdateEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.UpdateEKSAccessEntryTemplateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService.UpdateEKSAccessEntryTemplate is not implemented"))
}

func (UnimplementedEKSAccessEntryTemplateServiceHandler) DeleteEKSAccessEntryTemplate(context.Context, *connect.Request[v1alpha1.DeleteEKSAccessEntryTemplateRequest]) (*connect.Response[v1alpha1.DeleteEKSAccessEntryTemplateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("commonfate.control.config.v1alpha1.EKSAccessEntryTemplateService.DeleteEKSAccessEntryTemplate is not implemented"))
}
