// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comment.proto

package comment

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

type Filter struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type Comment struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Text                 string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Channel              string   `protobuf:"bytes,5,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_749aee09ea917828, []int{1}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Comment) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Comment) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Comment) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Comment) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func init() {
	proto.RegisterType((*Filter)(nil), "comment.Filter")
	proto.RegisterType((*Comment)(nil), "comment.Comment")
}

func init() { proto.RegisterFile("comment.proto", fileDescriptor_749aee09ea917828) }

var fileDescriptor_749aee09ea917828 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xcd, 0xca, 0x82, 0x40,
	0x14, 0x86, 0x3f, 0xbf, 0xfc, 0xc1, 0x03, 0xfd, 0x70, 0x68, 0x31, 0x44, 0x44, 0xb8, 0x6a, 0x25,
	0x91, 0x97, 0x10, 0xd5, 0xbe, 0xae, 0xc0, 0xe4, 0x40, 0x82, 0xa3, 0x36, 0x1e, 0x43, 0x17, 0xdd,
	0x7b, 0x38, 0x33, 0xd6, 0xee, 0xbc, 0xcf, 0xfb, 0xc0, 0xcc, 0x0b, 0xd3, 0xac, 0x92, 0x92, 0x4a,
	0x8e, 0x6b, 0x55, 0x71, 0x85, 0x81, 0x8d, 0xd1, 0x06, 0xfc, 0x73, 0x5e, 0x30, 0x29, 0x5c, 0x82,
	0xf7, 0x6c, 0x49, 0xf5, 0xc2, 0xd9, 0x3a, 0xbb, 0xf0, 0x6a, 0x42, 0xf4, 0x86, 0xe0, 0x68, 0x54,
	0x44, 0x70, 0xb9, 0xaf, 0xc9, 0xf6, 0xfa, 0x1e, 0x58, 0xdb, 0x90, 0x12, 0xff, 0x86, 0x0d, 0xb7,
	0xf6, 0xa8, 0x63, 0x31, 0xb1, 0x1e, 0x75, 0x8c, 0x6b, 0x08, 0x39, 0x97, 0xd4, 0x70, 0x2a, 0x6b,
	0xe1, 0xea, 0xe2, 0x07, 0x50, 0x40, 0x90, 0x3d, 0xd2, 0xb2, 0xa4, 0x42, 0x78, 0xba, 0x1b, 0xe3,
	0xe1, 0x04, 0x33, 0xfb, 0xfc, 0x8d, 0xd4, 0x2b, 0xcf, 0x08, 0x13, 0x80, 0x0b, 0xf1, 0xf8, 0xa7,
	0x79, 0x3c, 0xee, 0x32, 0x2b, 0x56, 0x8b, 0x2f, 0xb0, 0x4a, 0xf4, 0xb7, 0x77, 0xee, 0xbe, 0x5e,
	0x9d, 0x7c, 0x02, 0x00, 0x00, 0xff, 0xff, 0x85, 0x9e, 0xe8, 0xeb, 0x06, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentServiceClient interface {
	GetComment(ctx context.Context, in *Filter, opts ...grpc.CallOption) (CommentService_GetCommentClient, error)
}

type commentServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommentServiceClient(cc *grpc.ClientConn) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) GetComment(ctx context.Context, in *Filter, opts ...grpc.CallOption) (CommentService_GetCommentClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CommentService_serviceDesc.Streams[0], "/comment.CommentService/GetComment", opts...)
	if err != nil {
		return nil, err
	}
	x := &commentServiceGetCommentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CommentService_GetCommentClient interface {
	Recv() (*Comment, error)
	grpc.ClientStream
}

type commentServiceGetCommentClient struct {
	grpc.ClientStream
}

func (x *commentServiceGetCommentClient) Recv() (*Comment, error) {
	m := new(Comment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommentServiceServer is the server API for CommentService service.
type CommentServiceServer interface {
	GetComment(*Filter, CommentService_GetCommentServer) error
}

// UnimplementedCommentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (*UnimplementedCommentServiceServer) GetComment(req *Filter, srv CommentService_GetCommentServer) error {
	return status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}

func RegisterCommentServiceServer(s *grpc.Server, srv CommentServiceServer) {
	s.RegisterService(&_CommentService_serviceDesc, srv)
}

func _CommentService_GetComment_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Filter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommentServiceServer).GetComment(m, &commentServiceGetCommentServer{stream})
}

type CommentService_GetCommentServer interface {
	Send(*Comment) error
	grpc.ServerStream
}

type commentServiceGetCommentServer struct {
	grpc.ServerStream
}

func (x *commentServiceGetCommentServer) Send(m *Comment) error {
	return x.ServerStream.SendMsg(m)
}

var _CommentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetComment",
			Handler:       _CommentService_GetComment_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "comment.proto",
}
