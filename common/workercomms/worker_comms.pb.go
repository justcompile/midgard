// Code generated by protoc-gen-go. DO NOT EDIT.
// source: worker_comms.proto

package workercomms

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Worker struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Worker) Reset()         { *m = Worker{} }
func (m *Worker) String() string { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()    {}
func (*Worker) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f27c1b99d76830, []int{0}
}

func (m *Worker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Worker.Unmarshal(m, b)
}
func (m *Worker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Worker.Marshal(b, m, deterministic)
}
func (m *Worker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Worker.Merge(m, src)
}
func (m *Worker) XXX_Size() int {
	return xxx_messageInfo_Worker.Size(m)
}
func (m *Worker) XXX_DiscardUnknown() {
	xxx_messageInfo_Worker.DiscardUnknown(m)
}

var xxx_messageInfo_Worker proto.InternalMessageInfo

func (m *Worker) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Connection struct {
	State                int32    `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f27c1b99d76830, []int{1}
}

func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (m *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(m, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func init() {
	proto.RegisterType((*Worker)(nil), "workercomms.Worker")
	proto.RegisterType((*Connection)(nil), "workercomms.Connection")
}

func init() { proto.RegisterFile("worker_comms.proto", fileDescriptor_e4f27c1b99d76830) }

var fileDescriptor_e4f27c1b99d76830 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xcf, 0x2f, 0xca,
	0x4e, 0x2d, 0x8a, 0x4f, 0xce, 0xcf, 0xcd, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x86, 0x88, 0x81, 0x85, 0x94, 0x64, 0xb8, 0xd8, 0xc2, 0xc1, 0x5c, 0x21, 0x21, 0x2e, 0x96, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x89, 0x8b, 0xcb,
	0x39, 0x3f, 0x2f, 0x2f, 0x35, 0xb9, 0x24, 0x33, 0x3f, 0x4f, 0x48, 0x84, 0x8b, 0xb5, 0xb8, 0x24,
	0xb1, 0x04, 0xa2, 0x84, 0x35, 0x08, 0xc2, 0x31, 0xea, 0x60, 0xe4, 0xe2, 0xf5, 0xcd, 0x4c, 0x49,
	0x4f, 0x2c, 0x4a, 0x81, 0x9a, 0x64, 0xc9, 0xc5, 0x0e, 0xd5, 0x25, 0x24, 0xac, 0x87, 0x64, 0x99,
	0x1e, 0x44, 0x5e, 0x4a, 0x1c, 0x45, 0x10, 0x61, 0x81, 0x12, 0x83, 0x90, 0x0d, 0x17, 0x97, 0x4b,
	0x66, 0x71, 0x32, 0x79, 0xba, 0x93, 0xd8, 0xc0, 0x1e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x68, 0x47, 0x49, 0x87, 0xf6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MidgardWorkerClient is the client API for MidgardWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MidgardWorkerClient interface {
	Connect(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*Connection, error)
	Disconnect(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*Connection, error)
}

type midgardWorkerClient struct {
	cc *grpc.ClientConn
}

func NewMidgardWorkerClient(cc *grpc.ClientConn) MidgardWorkerClient {
	return &midgardWorkerClient{cc}
}

func (c *midgardWorkerClient) Connect(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*Connection, error) {
	out := new(Connection)
	err := c.cc.Invoke(ctx, "/workercomms.MidgardWorker/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midgardWorkerClient) Disconnect(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*Connection, error) {
	out := new(Connection)
	err := c.cc.Invoke(ctx, "/workercomms.MidgardWorker/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MidgardWorkerServer is the server API for MidgardWorker service.
type MidgardWorkerServer interface {
	Connect(context.Context, *Worker) (*Connection, error)
	Disconnect(context.Context, *Worker) (*Connection, error)
}

// UnimplementedMidgardWorkerServer can be embedded to have forward compatible implementations.
type UnimplementedMidgardWorkerServer struct {
}

func (*UnimplementedMidgardWorkerServer) Connect(ctx context.Context, req *Worker) (*Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedMidgardWorkerServer) Disconnect(ctx context.Context, req *Worker) (*Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}

func RegisterMidgardWorkerServer(s *grpc.Server, srv MidgardWorkerServer) {
	s.RegisterService(&_MidgardWorker_serviceDesc, srv)
}

func _MidgardWorker_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Worker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardWorkerServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workercomms.MidgardWorker/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardWorkerServer).Connect(ctx, req.(*Worker))
	}
	return interceptor(ctx, in, info, handler)
}

func _MidgardWorker_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Worker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardWorkerServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workercomms.MidgardWorker/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardWorkerServer).Disconnect(ctx, req.(*Worker))
	}
	return interceptor(ctx, in, info, handler)
}

var _MidgardWorker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "workercomms.MidgardWorker",
	HandlerType: (*MidgardWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _MidgardWorker_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _MidgardWorker_Disconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "worker_comms.proto",
}
