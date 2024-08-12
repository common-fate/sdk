// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: granted/ingest/aws/v1alpha1/ingest.proto

package awsv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/common-fate/sdk/gen/granted/ingest/aws/v1alpha1"
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
	// IngestServiceName is the fully-qualified name of the IngestService service.
	IngestServiceName = "granted.ingest.aws.v1alpha1.IngestService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IngestServiceBatchWriteEventsProcedure is the fully-qualified name of the IngestService's
	// BatchWriteEvents RPC.
	IngestServiceBatchWriteEventsProcedure = "/granted.ingest.aws.v1alpha1.IngestService/BatchWriteEvents"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	ingestServiceServiceDescriptor                = v1alpha1.File_granted_ingest_aws_v1alpha1_ingest_proto.Services().ByName("IngestService")
	ingestServiceBatchWriteEventsMethodDescriptor = ingestServiceServiceDescriptor.Methods().ByName("BatchWriteEvents")
)

// IngestServiceClient is a client for the granted.ingest.aws.v1alpha1.IngestService service.
type IngestServiceClient interface {
	// Writes a batch of IAM observability events.
	BatchWriteEvents(context.Context, *connect.Request[v1alpha1.BatchWriteEventsRequest]) (*connect.Response[v1alpha1.BatchWriteEventsResponse], error)
}

// NewIngestServiceClient constructs a client for the granted.ingest.aws.v1alpha1.IngestService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIngestServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) IngestServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &ingestServiceClient{
		batchWriteEvents: connect.NewClient[v1alpha1.BatchWriteEventsRequest, v1alpha1.BatchWriteEventsResponse](
			httpClient,
			baseURL+IngestServiceBatchWriteEventsProcedure,
			connect.WithSchema(ingestServiceBatchWriteEventsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// ingestServiceClient implements IngestServiceClient.
type ingestServiceClient struct {
	batchWriteEvents *connect.Client[v1alpha1.BatchWriteEventsRequest, v1alpha1.BatchWriteEventsResponse]
}

// BatchWriteEvents calls granted.ingest.aws.v1alpha1.IngestService.BatchWriteEvents.
func (c *ingestServiceClient) BatchWriteEvents(ctx context.Context, req *connect.Request[v1alpha1.BatchWriteEventsRequest]) (*connect.Response[v1alpha1.BatchWriteEventsResponse], error) {
	return c.batchWriteEvents.CallUnary(ctx, req)
}

// IngestServiceHandler is an implementation of the granted.ingest.aws.v1alpha1.IngestService
// service.
type IngestServiceHandler interface {
	// Writes a batch of IAM observability events.
	BatchWriteEvents(context.Context, *connect.Request[v1alpha1.BatchWriteEventsRequest]) (*connect.Response[v1alpha1.BatchWriteEventsResponse], error)
}

// NewIngestServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIngestServiceHandler(svc IngestServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	ingestServiceBatchWriteEventsHandler := connect.NewUnaryHandler(
		IngestServiceBatchWriteEventsProcedure,
		svc.BatchWriteEvents,
		connect.WithSchema(ingestServiceBatchWriteEventsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/granted.ingest.aws.v1alpha1.IngestService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IngestServiceBatchWriteEventsProcedure:
			ingestServiceBatchWriteEventsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIngestServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIngestServiceHandler struct{}

func (UnimplementedIngestServiceHandler) BatchWriteEvents(context.Context, *connect.Request[v1alpha1.BatchWriteEventsRequest]) (*connect.Response[v1alpha1.BatchWriteEventsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("granted.ingest.aws.v1alpha1.IngestService.BatchWriteEvents is not implemented"))
}
