// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package compile

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MeterReportServiceClient is the client API for MeterReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeterReportServiceClient interface {
	// Meter data is reported in a certain period. The agent/SDK should report all
	// collected metrics in this period through one stream.
	Collect(ctx context.Context, opts ...grpc.CallOption) (MeterReportService_CollectClient, error)
}

type meterReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeterReportServiceClient(cc grpc.ClientConnInterface) MeterReportServiceClient {
	return &meterReportServiceClient{cc}
}

func (c *meterReportServiceClient) Collect(ctx context.Context, opts ...grpc.CallOption) (MeterReportService_CollectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MeterReportService_serviceDesc.Streams[0], "/skywalking.v3.MeterReportService/collect", opts...)
	if err != nil {
		return nil, err
	}
	x := &meterReportServiceCollectClient{stream}
	return x, nil
}

type MeterReportService_CollectClient interface {
	Send(*MeterData) error
	CloseAndRecv() (*Commands, error)
	grpc.ClientStream
}

type meterReportServiceCollectClient struct {
	grpc.ClientStream
}

func (x *meterReportServiceCollectClient) Send(m *MeterData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *meterReportServiceCollectClient) CloseAndRecv() (*Commands, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Commands)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MeterReportServiceServer is the server API for MeterReportService service.
// All implementations must embed UnimplementedMeterReportServiceServer
// for forward compatibility
type MeterReportServiceServer interface {
	// Meter data is reported in a certain period. The agent/SDK should report all
	// collected metrics in this period through one stream.
	Collect(MeterReportService_CollectServer) error
	mustEmbedUnimplementedMeterReportServiceServer()
}

// UnimplementedMeterReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMeterReportServiceServer struct{}

func (UnimplementedMeterReportServiceServer) Collect(MeterReportService_CollectServer) error {
	return status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedMeterReportServiceServer) mustEmbedUnimplementedMeterReportServiceServer() {}

// UnsafeMeterReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeterReportServiceServer will
// result in compilation errors.
type UnsafeMeterReportServiceServer interface {
	mustEmbedUnimplementedMeterReportServiceServer()
}

func RegisterMeterReportServiceServer(s grpc.ServiceRegistrar, srv MeterReportServiceServer) {
	s.RegisterService(&_MeterReportService_serviceDesc, srv)
}

func _MeterReportService_Collect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MeterReportServiceServer).Collect(&meterReportServiceCollectServer{stream})
}

type MeterReportService_CollectServer interface {
	SendAndClose(*Commands) error
	Recv() (*MeterData, error)
	grpc.ServerStream
}

type meterReportServiceCollectServer struct {
	grpc.ServerStream
}

func (x *meterReportServiceCollectServer) SendAndClose(m *Commands) error {
	return x.ServerStream.SendMsg(m)
}

func (x *meterReportServiceCollectServer) Recv() (*MeterData, error) {
	m := new(MeterData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MeterReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skywalking.v3.MeterReportService",
	HandlerType: (*MeterReportServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collect",
			Handler:       _MeterReportService_Collect_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "plugins/inputs/skywalking/v3/proto/language-agent/Meter.proto",
}