// Code generated by protoc-gen-go. DO NOT EDIT.
// source: index.proto

package user

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

type UserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserResponse struct {
	Id                   int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Username             string      `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Avatarurl            string      `protobuf:"bytes,4,opt,name=avatarurl,proto3" json:"avatarurl,omitempty"`
	Location             string      `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Statistics           *Statistics `protobuf:"bytes,6,opt,name=statistics,proto3" json:"statistics,omitempty"`
	ListURLs             []string    `protobuf:"bytes,7,rep,name=listURLs,proto3" json:"listURLs,omitempty"`
	NodeId               string      `protobuf:"bytes,8,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	CreatedAt            int64       `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Bio                  string      `protobuf:"bytes,10,opt,name=bio,proto3" json:"bio,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserResponse) GetAvatarurl() string {
	if m != nil {
		return m.Avatarurl
	}
	return ""
}

func (m *UserResponse) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *UserResponse) GetStatistics() *Statistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *UserResponse) GetListURLs() []string {
	if m != nil {
		return m.ListURLs
	}
	return nil
}

func (m *UserResponse) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *UserResponse) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *UserResponse) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

type Statistics struct {
	Followers            int64    `protobuf:"varint,1,opt,name=followers,proto3" json:"followers,omitempty"`
	Following            int64    `protobuf:"varint,2,opt,name=following,proto3" json:"following,omitempty"`
	Repos                int64    `protobuf:"varint,3,opt,name=repos,proto3" json:"repos,omitempty"`
	Gists                int64    `protobuf:"varint,4,opt,name=gists,proto3" json:"gists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Statistics) Reset()         { *m = Statistics{} }
func (m *Statistics) String() string { return proto.CompactTextString(m) }
func (*Statistics) ProtoMessage()    {}
func (*Statistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_f750e0f7889345b5, []int{2}
}

func (m *Statistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statistics.Unmarshal(m, b)
}
func (m *Statistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statistics.Marshal(b, m, deterministic)
}
func (m *Statistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statistics.Merge(m, src)
}
func (m *Statistics) XXX_Size() int {
	return xxx_messageInfo_Statistics.Size(m)
}
func (m *Statistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Statistics.DiscardUnknown(m)
}

var xxx_messageInfo_Statistics proto.InternalMessageInfo

func (m *Statistics) GetFollowers() int64 {
	if m != nil {
		return m.Followers
	}
	return 0
}

func (m *Statistics) GetFollowing() int64 {
	if m != nil {
		return m.Following
	}
	return 0
}

func (m *Statistics) GetRepos() int64 {
	if m != nil {
		return m.Repos
	}
	return 0
}

func (m *Statistics) GetGists() int64 {
	if m != nil {
		return m.Gists
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "UserRequest")
	proto.RegisterType((*UserResponse)(nil), "UserResponse")
	proto.RegisterType((*Statistics)(nil), "Statistics")
}

func init() { proto.RegisterFile("index.proto", fileDescriptor_f750e0f7889345b5) }

var fileDescriptor_f750e0f7889345b5 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xfd, 0x9a, 0xb4, 0x69, 0x33, 0xe9, 0xf7, 0xf1, 0xb1, 0x08, 0x2e, 0x45, 0xa1, 0xe4, 0x54,
	0x11, 0x72, 0xa8, 0x07, 0xcf, 0x7a, 0x11, 0xc1, 0x53, 0x4a, 0x2f, 0x5e, 0xca, 0x36, 0x19, 0xcb,
	0x42, 0xcc, 0xd6, 0x9d, 0x69, 0xf5, 0xbf, 0xf8, 0x67, 0x25, 0x93, 0xda, 0xb4, 0xb7, 0x7d, 0x6f,
	0xf6, 0xbd, 0x65, 0xde, 0x5b, 0x48, 0x6c, 0x5d, 0xe2, 0x57, 0xb6, 0xf5, 0x8e, 0x5d, 0x7a, 0x03,
	0xc9, 0x92, 0xd0, 0xe7, 0xf8, 0xb1, 0x43, 0x62, 0x35, 0x81, 0xd1, 0x8e, 0xd0, 0xd7, 0xe6, 0x1d,
	0x75, 0x6f, 0xda, 0x9b, 0xc5, 0xf9, 0x11, 0xa7, 0xdf, 0x01, 0x8c, 0xdb, 0xbb, 0xb4, 0x75, 0x35,
	0xa1, 0xfa, 0x07, 0x81, 0x2d, 0xe5, 0x5a, 0x98, 0x07, 0xb6, 0x54, 0x0a, 0xfa, 0x22, 0x0c, 0x44,
	0x28, 0xe7, 0x33, 0xc3, 0xf0, 0xdc, 0x50, 0x5d, 0x41, 0x6c, 0xf6, 0x86, 0x8d, 0xdf, 0xf9, 0x4a,
	0xf7, 0x65, 0xd8, 0x11, 0x8d, 0xb2, 0x72, 0x85, 0x61, 0xeb, 0x6a, 0x3d, 0x68, 0x95, 0xbf, 0x58,
	0xdd, 0x02, 0x10, 0x1b, 0xb6, 0xc4, 0xb6, 0x20, 0x1d, 0x4d, 0x7b, 0xb3, 0x64, 0x9e, 0x64, 0x8b,
	0x23, 0x95, 0x9f, 0x8c, 0xc5, 0xc8, 0x12, 0x2f, 0xf3, 0x17, 0xd2, 0xc3, 0x69, 0x28, 0x46, 0x07,
	0xac, 0x2e, 0x61, 0x58, 0xbb, 0x12, 0x57, 0xb6, 0xd4, 0x23, 0x79, 0x23, 0x6a, 0xe0, 0x73, 0xa9,
	0xae, 0x01, 0x0a, 0x8f, 0x86, 0xb1, 0x5c, 0x19, 0xd6, 0xb1, 0xec, 0x18, 0x1f, 0x98, 0x07, 0x56,
	0xff, 0x21, 0x5c, 0x5b, 0xa7, 0x41, 0x34, 0xcd, 0x31, 0x65, 0x80, 0xee, 0xfd, 0x66, 0xb5, 0x37,
	0x57, 0x55, 0xee, 0x13, 0x3d, 0x1d, 0x12, 0xea, 0x88, 0x6e, 0x6a, 0xeb, 0x8d, 0xa4, 0x75, 0x9c,
	0xda, 0x7a, 0xa3, 0x2e, 0x60, 0xe0, 0x71, 0xeb, 0x48, 0xf2, 0x0a, 0xf3, 0x16, 0x34, 0xec, 0xc6,
	0x12, 0x93, 0x04, 0x15, 0xe6, 0x2d, 0x98, 0xdf, 0xb7, 0xf5, 0x2d, 0xd0, 0xef, 0x6d, 0x81, 0x6a,
	0x06, 0xc3, 0x27, 0xe4, 0x86, 0x51, 0xe3, 0xec, 0xa4, 0xd7, 0xc9, 0xdf, 0xec, 0xb4, 0xb9, 0xf4,
	0xcf, 0x63, 0xf4, 0xda, 0x6f, 0x7a, 0x58, 0x47, 0xf2, 0x0d, 0xee, 0x7e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x8a, 0x99, 0xcc, 0x15, 0x15, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUser(context.Context, *UserRequest) (*UserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index.proto",
}
