// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pixelfs/v1/block.proto

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
	// BlockServiceName is the fully-qualified name of the BlockService service.
	BlockServiceName = "pixelfs.v1.BlockService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BlockServiceGetBlockDurationProcedure is the fully-qualified name of the BlockService's
	// GetBlockDuration RPC.
	BlockServiceGetBlockDurationProcedure = "/pixelfs.v1.BlockService/GetBlockDuration"
	// BlockServiceSetBlockDurationProcedure is the fully-qualified name of the BlockService's
	// SetBlockDuration RPC.
	BlockServiceSetBlockDurationProcedure = "/pixelfs.v1.BlockService/SetBlockDuration"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	blockServiceServiceDescriptor                = v1.File_pixelfs_v1_block_proto.Services().ByName("BlockService")
	blockServiceGetBlockDurationMethodDescriptor = blockServiceServiceDescriptor.Methods().ByName("GetBlockDuration")
	blockServiceSetBlockDurationMethodDescriptor = blockServiceServiceDescriptor.Methods().ByName("SetBlockDuration")
)

// BlockServiceClient is a client for the pixelfs.v1.BlockService service.
type BlockServiceClient interface {
	GetBlockDuration(context.Context, *connect.Request[v1.GetBlockDurationRequest]) (*connect.Response[v1.GetBlockDurationResponse], error)
	SetBlockDuration(context.Context, *connect.Request[v1.SetBlockDurationRequest]) (*connect.Response[v1.SetBlockDurationResponse], error)
}

// NewBlockServiceClient constructs a client for the pixelfs.v1.BlockService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBlockServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BlockServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &blockServiceClient{
		getBlockDuration: connect.NewClient[v1.GetBlockDurationRequest, v1.GetBlockDurationResponse](
			httpClient,
			baseURL+BlockServiceGetBlockDurationProcedure,
			connect.WithSchema(blockServiceGetBlockDurationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setBlockDuration: connect.NewClient[v1.SetBlockDurationRequest, v1.SetBlockDurationResponse](
			httpClient,
			baseURL+BlockServiceSetBlockDurationProcedure,
			connect.WithSchema(blockServiceSetBlockDurationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// blockServiceClient implements BlockServiceClient.
type blockServiceClient struct {
	getBlockDuration *connect.Client[v1.GetBlockDurationRequest, v1.GetBlockDurationResponse]
	setBlockDuration *connect.Client[v1.SetBlockDurationRequest, v1.SetBlockDurationResponse]
}

// GetBlockDuration calls pixelfs.v1.BlockService.GetBlockDuration.
func (c *blockServiceClient) GetBlockDuration(ctx context.Context, req *connect.Request[v1.GetBlockDurationRequest]) (*connect.Response[v1.GetBlockDurationResponse], error) {
	return c.getBlockDuration.CallUnary(ctx, req)
}

// SetBlockDuration calls pixelfs.v1.BlockService.SetBlockDuration.
func (c *blockServiceClient) SetBlockDuration(ctx context.Context, req *connect.Request[v1.SetBlockDurationRequest]) (*connect.Response[v1.SetBlockDurationResponse], error) {
	return c.setBlockDuration.CallUnary(ctx, req)
}

// BlockServiceHandler is an implementation of the pixelfs.v1.BlockService service.
type BlockServiceHandler interface {
	GetBlockDuration(context.Context, *connect.Request[v1.GetBlockDurationRequest]) (*connect.Response[v1.GetBlockDurationResponse], error)
	SetBlockDuration(context.Context, *connect.Request[v1.SetBlockDurationRequest]) (*connect.Response[v1.SetBlockDurationResponse], error)
}

// NewBlockServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBlockServiceHandler(svc BlockServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	blockServiceGetBlockDurationHandler := connect.NewUnaryHandler(
		BlockServiceGetBlockDurationProcedure,
		svc.GetBlockDuration,
		connect.WithSchema(blockServiceGetBlockDurationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	blockServiceSetBlockDurationHandler := connect.NewUnaryHandler(
		BlockServiceSetBlockDurationProcedure,
		svc.SetBlockDuration,
		connect.WithSchema(blockServiceSetBlockDurationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/pixelfs.v1.BlockService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BlockServiceGetBlockDurationProcedure:
			blockServiceGetBlockDurationHandler.ServeHTTP(w, r)
		case BlockServiceSetBlockDurationProcedure:
			blockServiceSetBlockDurationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBlockServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBlockServiceHandler struct{}

func (UnimplementedBlockServiceHandler) GetBlockDuration(context.Context, *connect.Request[v1.GetBlockDurationRequest]) (*connect.Response[v1.GetBlockDurationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pixelfs.v1.BlockService.GetBlockDuration is not implemented"))
}

func (UnimplementedBlockServiceHandler) SetBlockDuration(context.Context, *connect.Request[v1.SetBlockDurationRequest]) (*connect.Response[v1.SetBlockDurationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pixelfs.v1.BlockService.SetBlockDuration is not implemented"))
}
