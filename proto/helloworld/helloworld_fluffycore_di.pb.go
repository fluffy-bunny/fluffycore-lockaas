// Code generated by protoc-gen-go-fluffycore-di. DO NOT EDIT.
// Code generated grpcGateway

package helloworld

import (
	context "context"
	fluffy_dozm_di "github.com/fluffy-bunny/fluffy-dozm-di"
	endpoint "github.com/fluffy-bunny/fluffycore/contracts/endpoint"
	dicontext "github.com/fluffy-bunny/fluffycore/middleware/dicontext"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc "google.golang.org/grpc"
)

// IFluffyCoreGreeterServer defines the grpc server
type IFluffyCoreGreeterServer interface {
	GreeterServer
}

type UnimplementedFluffyCoreGreeterServerEndpointRegistration struct {
}

func (UnimplementedFluffyCoreGreeterServerEndpointRegistration) RegisterFluffyCoreHandler(gwmux *runtime.ServeMux, conn *grpc.ClientConn) {
}

// GreeterFluffyCoreServer defines the grpc server truct
type GreeterFluffyCoreServer struct {
	UnimplementedGreeterServer
	UnimplementedFluffyCoreGreeterServerEndpointRegistration
}

// RegisterFluffyCoreGRPCService the server with grpc
func (srv *GreeterFluffyCoreServer) RegisterFluffyCoreGRPCService(s *grpc.Server) {
	RegisterGreeterServer(s, srv)
}

// AddGreeterServerWithExternalRegistration adds the fluffycore aware grpc server and external registration service.  Mainly used for grpc-gateway
func AddGreeterServerWithExternalRegistration(cb fluffy_dozm_di.ContainerBuilder, ctor any, register func() endpoint.IEndpointRegistration) {
	fluffy_dozm_di.AddSingleton[endpoint.IEndpointRegistration](cb, register)
	fluffy_dozm_di.AddScoped[IFluffyCoreGreeterServer](cb, ctor)
}

// AddGreeterServer adds the fluffycore aware grpc server
func AddGreeterServer(cb fluffy_dozm_di.ContainerBuilder, ctor any) {
	AddGreeterServerWithExternalRegistration(cb, ctor, func() endpoint.IEndpointRegistration {
		return &GreeterFluffyCoreServer{}
	})
}

// SayHello...
func (s *GreeterFluffyCoreServer) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	requestContainer := dicontext.GetRequestContainer(ctx)
	downstreamService := fluffy_dozm_di.Get[IFluffyCoreGreeterServer](requestContainer)
	return downstreamService.SayHello(ctx, request)
}

// IFluffyCoreMyStreamServiceServer defines the grpc server
type IFluffyCoreMyStreamServiceServer interface {
	MyStreamServiceServer
}

type UnimplementedFluffyCoreMyStreamServiceServerEndpointRegistration struct {
}

func (UnimplementedFluffyCoreMyStreamServiceServerEndpointRegistration) RegisterFluffyCoreHandler(gwmux *runtime.ServeMux, conn *grpc.ClientConn) {
}

// MyStreamServiceFluffyCoreServer defines the grpc server truct
type MyStreamServiceFluffyCoreServer struct {
	UnimplementedMyStreamServiceServer
	UnimplementedFluffyCoreMyStreamServiceServerEndpointRegistration
}

// RegisterFluffyCoreGRPCService the server with grpc
func (srv *MyStreamServiceFluffyCoreServer) RegisterFluffyCoreGRPCService(s *grpc.Server) {
	RegisterMyStreamServiceServer(s, srv)
}

// AddMyStreamServiceServerWithExternalRegistration adds the fluffycore aware grpc server and external registration service.  Mainly used for grpc-gateway
func AddMyStreamServiceServerWithExternalRegistration(cb fluffy_dozm_di.ContainerBuilder, ctor any, register func() endpoint.IEndpointRegistration) {
	fluffy_dozm_di.AddSingleton[endpoint.IEndpointRegistration](cb, register)
	fluffy_dozm_di.AddScoped[IFluffyCoreMyStreamServiceServer](cb, ctor)
}

// AddMyStreamServiceServer adds the fluffycore aware grpc server
func AddMyStreamServiceServer(cb fluffy_dozm_di.ContainerBuilder, ctor any) {
	AddMyStreamServiceServerWithExternalRegistration(cb, ctor, func() endpoint.IEndpointRegistration {
		return &MyStreamServiceFluffyCoreServer{}
	})
}

// RequestPoints...
func (s *MyStreamServiceFluffyCoreServer) RequestPoints(request *PointsRequest, stream MyStreamService_RequestPointsServer) error {
	ctx := stream.Context()
	requestContainer := dicontext.GetRequestContainer(ctx)
	downstreamService := fluffy_dozm_di.Get[IFluffyCoreMyStreamServiceServer](requestContainer)
	return downstreamService.RequestPoints(request, stream)
}

// StreamPoints...
func (s *MyStreamServiceFluffyCoreServer) StreamPoints(stream MyStreamService_StreamPointsServer) error {
	ctx := stream.Context()
	requestContainer := dicontext.GetRequestContainer(ctx)
	downstreamService := fluffy_dozm_di.Get[IFluffyCoreMyStreamServiceServer](requestContainer)
	return downstreamService.StreamPoints(stream)
}