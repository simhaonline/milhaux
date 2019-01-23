// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend.proto

package backend

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type MessageSubmitRequest struct {
	To                   *string  `protobuf:"bytes,1,req,name=To" json:"To,omitempty"`
	From                 *string  `protobuf:"bytes,2,req,name=From" json:"From,omitempty"`
	Data                 *string  `protobuf:"bytes,3,req,name=Data" json:"Data,omitempty"`
	QueueId              *string  `protobuf:"bytes,4,req,name=QueueId" json:"QueueId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageSubmitRequest) Reset()         { *m = MessageSubmitRequest{} }
func (m *MessageSubmitRequest) String() string { return proto.CompactTextString(m) }
func (*MessageSubmitRequest) ProtoMessage()    {}
func (*MessageSubmitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ab9ba5b8d8b2ba5, []int{0}
}

func (m *MessageSubmitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageSubmitRequest.Unmarshal(m, b)
}
func (m *MessageSubmitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageSubmitRequest.Marshal(b, m, deterministic)
}
func (m *MessageSubmitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageSubmitRequest.Merge(m, src)
}
func (m *MessageSubmitRequest) XXX_Size() int {
	return xxx_messageInfo_MessageSubmitRequest.Size(m)
}
func (m *MessageSubmitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageSubmitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageSubmitRequest proto.InternalMessageInfo

func (m *MessageSubmitRequest) GetTo() string {
	if m != nil && m.To != nil {
		return *m.To
	}
	return ""
}

func (m *MessageSubmitRequest) GetFrom() string {
	if m != nil && m.From != nil {
		return *m.From
	}
	return ""
}

func (m *MessageSubmitRequest) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func (m *MessageSubmitRequest) GetQueueId() string {
	if m != nil && m.QueueId != nil {
		return *m.QueueId
	}
	return ""
}

type MessageSubmitResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageSubmitResponse) Reset()         { *m = MessageSubmitResponse{} }
func (m *MessageSubmitResponse) String() string { return proto.CompactTextString(m) }
func (*MessageSubmitResponse) ProtoMessage()    {}
func (*MessageSubmitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ab9ba5b8d8b2ba5, []int{1}
}

func (m *MessageSubmitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageSubmitResponse.Unmarshal(m, b)
}
func (m *MessageSubmitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageSubmitResponse.Marshal(b, m, deterministic)
}
func (m *MessageSubmitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageSubmitResponse.Merge(m, src)
}
func (m *MessageSubmitResponse) XXX_Size() int {
	return xxx_messageInfo_MessageSubmitResponse.Size(m)
}
func (m *MessageSubmitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageSubmitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageSubmitResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MessageSubmitRequest)(nil), "MessageSubmitRequest")
	proto.RegisterType((*MessageSubmitResponse)(nil), "MessageSubmitResponse")
}

func init() { proto.RegisterFile("backend.proto", fileDescriptor_5ab9ba5b8d8b2ba5) }

var fileDescriptor_5ab9ba5b8d8b2ba5 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4a, 0x4c, 0xce,
	0x4e, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xca, 0xe0, 0x12, 0xf1, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0x0d, 0x2e, 0x4d, 0xca, 0xcd, 0x2c, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0xe2, 0xe3, 0x62, 0x0a, 0xc9, 0x97, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0c, 0x62, 0x0a,
	0xc9, 0x17, 0x12, 0xe2, 0x62, 0x71, 0x2b, 0xca, 0xcf, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x20,
	0x31, 0x97, 0xc4, 0x92, 0x44, 0x09, 0x66, 0x88, 0x18, 0x88, 0x2d, 0x24, 0xc1, 0xc5, 0x1e, 0x58,
	0x9a, 0x5a, 0x9a, 0xea, 0x99, 0x22, 0xc1, 0x02, 0x16, 0x86, 0x71, 0x95, 0xc4, 0xb9, 0x44, 0xd1,
	0x6c, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x35, 0x0a, 0xe0, 0xe2, 0x81, 0x49, 0x94, 0xe4, 0x17,
	0xa5, 0x0a, 0x39, 0x70, 0xf1, 0x42, 0x54, 0x40, 0x45, 0x85, 0x44, 0xf5, 0xb0, 0x39, 0x51, 0x4a,
	0x4c, 0x0f, 0xab, 0x79, 0x4a, 0x0c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0x2b, 0x67, 0xb7,
	0xe4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageStoreClient is the client API for MessageStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageStoreClient interface {
	SubmitMessage(ctx context.Context, in *MessageSubmitRequest, opts ...grpc.CallOption) (*MessageSubmitResponse, error)
}

type messageStoreClient struct {
	cc *grpc.ClientConn
}

func NewMessageStoreClient(cc *grpc.ClientConn) MessageStoreClient {
	return &messageStoreClient{cc}
}

func (c *messageStoreClient) SubmitMessage(ctx context.Context, in *MessageSubmitRequest, opts ...grpc.CallOption) (*MessageSubmitResponse, error) {
	out := new(MessageSubmitResponse)
	err := c.cc.Invoke(ctx, "/MessageStore/SubmitMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageStoreServer is the server API for MessageStore service.
type MessageStoreServer interface {
	SubmitMessage(context.Context, *MessageSubmitRequest) (*MessageSubmitResponse, error)
}

func RegisterMessageStoreServer(s *grpc.Server, srv MessageStoreServer) {
	s.RegisterService(&_MessageStore_serviceDesc, srv)
}

func _MessageStore_SubmitMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageSubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageStoreServer).SubmitMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageStore/SubmitMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageStoreServer).SubmitMessage(ctx, req.(*MessageSubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MessageStore",
	HandlerType: (*MessageStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitMessage",
			Handler:    _MessageStore_SubmitMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend.proto",
}
