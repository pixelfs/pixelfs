// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pixelfs/v1/system.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/pixelfs/pixelfs/gen/pixelfs/v1"
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
	// SystemServiceName is the fully-qualified name of the SystemService service.
	SystemServiceName = "pixelfs.v1.SystemService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SystemServiceStartupProcedure is the fully-qualified name of the SystemService's Startup RPC.
	SystemServiceStartupProcedure = "/pixelfs.v1.SystemService/Startup"
	// SystemServiceShutdownProcedure is the fully-qualified name of the SystemService's Shutdown RPC.
	SystemServiceShutdownProcedure = "/pixelfs.v1.SystemService/Shutdown"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	systemServiceServiceDescriptor        = v1.File_pixelfs_v1_system_proto.Services().ByName("SystemService")
	systemServiceStartupMethodDescriptor  = systemServiceServiceDescriptor.Methods().ByName("Startup")
	systemServiceShutdownMethodDescriptor = systemServiceServiceDescriptor.Methods().ByName("Shutdown")
)

// SystemServiceClient is a client for the pixelfs.v1.SystemService service.
type SystemServiceClient interface {
	Startup(context.Context, *connect.Request[v1.SystemStartupRequest]) (*connect.Response[v1.SystemStartupResponse], error)
	Shutdown(context.Context, *connect.Request[v1.SystemShutdownRequest]) (*connect.Response[v1.SystemShutdownResponse], error)
}

// NewSystemServiceClient constructs a client for the pixelfs.v1.SystemService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSystemServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SystemServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &systemServiceClient{
		startup: connect.NewClient[v1.SystemStartupRequest, v1.SystemStartupResponse](
			httpClient,
			baseURL+SystemServiceStartupProcedure,
			connect.WithSchema(systemServiceStartupMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		shutdown: connect.NewClient[v1.SystemShutdownRequest, v1.SystemShutdownResponse](
			httpClient,
			baseURL+SystemServiceShutdownProcedure,
			connect.WithSchema(systemServiceShutdownMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// systemServiceClient implements SystemServiceClient.
type systemServiceClient struct {
	startup  *connect.Client[v1.SystemStartupRequest, v1.SystemStartupResponse]
	shutdown *connect.Client[v1.SystemShutdownRequest, v1.SystemShutdownResponse]
}

// Startup calls pixelfs.v1.SystemService.Startup.
func (c *systemServiceClient) Startup(ctx context.Context, req *connect.Request[v1.SystemStartupRequest]) (*connect.Response[v1.SystemStartupResponse], error) {
	return c.startup.CallUnary(ctx, req)
}

// Shutdown calls pixelfs.v1.SystemService.Shutdown.
func (c *systemServiceClient) Shutdown(ctx context.Context, req *connect.Request[v1.SystemShutdownRequest]) (*connect.Response[v1.SystemShutdownResponse], error) {
	return c.shutdown.CallUnary(ctx, req)
}

// SystemServiceHandler is an implementation of the pixelfs.v1.SystemService service.
type SystemServiceHandler interface {
	Startup(context.Context, *connect.Request[v1.SystemStartupRequest]) (*connect.Response[v1.SystemStartupResponse], error)
	Shutdown(context.Context, *connect.Request[v1.SystemShutdownRequest]) (*connect.Response[v1.SystemShutdownResponse], error)
}

// NewSystemServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSystemServiceHandler(svc SystemServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	systemServiceStartupHandler := connect.NewUnaryHandler(
		SystemServiceStartupProcedure,
		svc.Startup,
		connect.WithSchema(systemServiceStartupMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	systemServiceShutdownHandler := connect.NewUnaryHandler(
		SystemServiceShutdownProcedure,
		svc.Shutdown,
		connect.WithSchema(systemServiceShutdownMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/pixelfs.v1.SystemService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SystemServiceStartupProcedure:
			systemServiceStartupHandler.ServeHTTP(w, r)
		case SystemServiceShutdownProcedure:
			systemServiceShutdownHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSystemServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSystemServiceHandler struct{}

func (UnimplementedSystemServiceHandler) Startup(context.Context, *connect.Request[v1.SystemStartupRequest]) (*connect.Response[v1.SystemStartupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pixelfs.v1.SystemService.Startup is not implemented"))
}

func (UnimplementedSystemServiceHandler) Shutdown(context.Context, *connect.Request[v1.SystemShutdownRequest]) (*connect.Response[v1.SystemShutdownResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pixelfs.v1.SystemService.Shutdown is not implemented"))
}
