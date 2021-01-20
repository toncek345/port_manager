package clientapi

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/toncek345/port_manager/internal/portdomain/proto"
	"google.golang.org/grpc"
)

type PortServiceMock struct {
	UpsertFn  func(context.Context, ...grpc.CallOption) (pb.PortDomain_UpsertClient, error)
	GetPortFn func(context.Context, *pb.GetPortRequest, ...grpc.CallOption) (*pb.Port, error)
}

func (m *PortServiceMock) Upsert(ctx context.Context, opts ...grpc.CallOption) (pb.PortDomain_UpsertClient, error) {
	return m.UpsertFn(ctx, opts...)
}

func (m *PortServiceMock) GetPort(ctx context.Context, in *pb.GetPortRequest, opts ...grpc.CallOption) (*pb.Port, error) {
	return m.GetPortFn(ctx, in, opts...)
}

type PortUpsertClientMock struct {
	SendFn         func(*pb.Port) error
	CloseAndRecvFn func() (*empty.Empty, error)
	grpc.ClientStream
}

func (m *PortUpsertClientMock) Send(port *pb.Port) error {
	return m.SendFn(port)
}

func (m *PortUpsertClientMock) CloseAndRecv() (*empty.Empty, error) {
	return m.CloseAndRecvFn()
}
