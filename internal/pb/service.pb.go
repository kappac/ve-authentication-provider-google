// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	error1 "github.com/kappac/ve-back-end-utils/pkg/proto/error"
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

type VEValidateTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VEValidateTokenRequest) Reset()         { *m = VEValidateTokenRequest{} }
func (m *VEValidateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*VEValidateTokenRequest) ProtoMessage()    {}
func (*VEValidateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *VEValidateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VEValidateTokenRequest.Unmarshal(m, b)
}
func (m *VEValidateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VEValidateTokenRequest.Marshal(b, m, deterministic)
}
func (m *VEValidateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VEValidateTokenRequest.Merge(m, src)
}
func (m *VEValidateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_VEValidateTokenRequest.Size(m)
}
func (m *VEValidateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VEValidateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VEValidateTokenRequest proto.InternalMessageInfo

func (m *VEValidateTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VEValidateTokenResponse struct {
	Info                 *VEProviderInfo         `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Request              *VEValidateTokenRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Error                *error1.VEError         `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *VEValidateTokenResponse) Reset()         { *m = VEValidateTokenResponse{} }
func (m *VEValidateTokenResponse) String() string { return proto.CompactTextString(m) }
func (*VEValidateTokenResponse) ProtoMessage()    {}
func (*VEValidateTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *VEValidateTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VEValidateTokenResponse.Unmarshal(m, b)
}
func (m *VEValidateTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VEValidateTokenResponse.Marshal(b, m, deterministic)
}
func (m *VEValidateTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VEValidateTokenResponse.Merge(m, src)
}
func (m *VEValidateTokenResponse) XXX_Size() int {
	return xxx_messageInfo_VEValidateTokenResponse.Size(m)
}
func (m *VEValidateTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VEValidateTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VEValidateTokenResponse proto.InternalMessageInfo

func (m *VEValidateTokenResponse) GetInfo() *VEProviderInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *VEValidateTokenResponse) GetRequest() *VEValidateTokenRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *VEValidateTokenResponse) GetError() *error1.VEError {
	if m != nil {
		return m.Error
	}
	return nil
}

type VEProviderInfo struct {
	FullName             string   `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	GivenName            string   `protobuf:"bytes,2,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName           string   `protobuf:"bytes,3,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	Picture              string   `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VEProviderInfo) Reset()         { *m = VEProviderInfo{} }
func (m *VEProviderInfo) String() string { return proto.CompactTextString(m) }
func (*VEProviderInfo) ProtoMessage()    {}
func (*VEProviderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *VEProviderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VEProviderInfo.Unmarshal(m, b)
}
func (m *VEProviderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VEProviderInfo.Marshal(b, m, deterministic)
}
func (m *VEProviderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VEProviderInfo.Merge(m, src)
}
func (m *VEProviderInfo) XXX_Size() int {
	return xxx_messageInfo_VEProviderInfo.Size(m)
}
func (m *VEProviderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VEProviderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VEProviderInfo proto.InternalMessageInfo

func (m *VEProviderInfo) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *VEProviderInfo) GetGivenName() string {
	if m != nil {
		return m.GivenName
	}
	return ""
}

func (m *VEProviderInfo) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

func (m *VEProviderInfo) GetPicture() string {
	if m != nil {
		return m.Picture
	}
	return ""
}

func (m *VEProviderInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*VEValidateTokenRequest)(nil), "pb.VEValidateTokenRequest")
	proto.RegisterType((*VEValidateTokenResponse)(nil), "pb.VEValidateTokenResponse")
	proto.RegisterType((*VEProviderInfo)(nil), "pb.VEProviderInfo")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x6b, 0xe3, 0x30,
	0x10, 0xc5, 0xc9, 0xbf, 0xcd, 0x7a, 0x42, 0x72, 0x10, 0xcb, 0xae, 0x49, 0x58, 0x76, 0x09, 0xa5,
	0xf4, 0x12, 0x1b, 0xd2, 0x7e, 0x83, 0x62, 0x68, 0x2f, 0xa5, 0xb8, 0xad, 0xaf, 0x45, 0xb6, 0xc7,
	0xae, 0xb0, 0x2d, 0xa9, 0xb2, 0x6c, 0xe8, 0x87, 0xe9, 0x77, 0x2d, 0x1e, 0x35, 0x87, 0xd0, 0xd0,
	0x93, 0x79, 0xf3, 0x7b, 0xcf, 0xd2, 0x3c, 0xc1, 0xb2, 0x45, 0xd3, 0x8b, 0x0c, 0x03, 0x6d, 0x94,
	0x55, 0x6c, 0xac, 0xd3, 0xf5, 0x75, 0x29, 0xec, 0x4b, 0x97, 0x06, 0x99, 0x6a, 0xc2, 0x8a, 0x6b,
	0xcd, 0xb3, 0xb0, 0xc7, 0x5d, 0xca, 0xb3, 0x6a, 0x87, 0x32, 0xdf, 0x75, 0x56, 0xd4, 0x6d, 0xa8,
	0xab, 0x32, 0xa4, 0x48, 0x88, 0xc6, 0x28, 0x13, 0x96, 0x28, 0xd1, 0x70, 0x8b, 0xb9, 0xfb, 0xd1,
	0x36, 0x80, 0xdf, 0x49, 0x94, 0xf0, 0x5a, 0xe4, 0xdc, 0xe2, 0xa3, 0xaa, 0x50, 0xc6, 0xf8, 0xda,
	0x61, 0x6b, 0xd9, 0x2f, 0x98, 0xd9, 0x41, 0xfb, 0xa3, 0xff, 0xa3, 0x0b, 0x2f, 0x76, 0x62, 0xfb,
	0x3e, 0x82, 0x3f, 0x5f, 0x02, 0xad, 0x56, 0xb2, 0x45, 0x76, 0x0e, 0x53, 0x21, 0x0b, 0x45, 0x81,
	0xc5, 0x9e, 0x05, 0x3a, 0x0d, 0x92, 0xe8, 0xde, 0xa8, 0x5e, 0xe4, 0x68, 0x6e, 0x65, 0xa1, 0x62,
	0xe2, 0xec, 0x0a, 0xe6, 0xc6, 0x1d, 0xe2, 0x8f, 0xc9, 0xba, 0x76, 0xd6, 0x53, 0xd7, 0x88, 0x0f,
	0x56, 0x76, 0x06, 0x33, 0x5a, 0xc1, 0x9f, 0x50, 0x66, 0x15, 0x90, 0x0a, 0x92, 0x28, 0x1a, 0xbe,
	0xb1, 0x83, 0xc3, 0xfd, 0x56, 0xc7, 0x87, 0xb2, 0x0d, 0x78, 0x45, 0x57, 0xd7, 0xcf, 0x92, 0x37,
	0xf8, 0xb9, 0xcc, 0xcf, 0x61, 0x70, 0xc7, 0x1b, 0x64, 0x7f, 0x01, 0x4a, 0xd1, 0xa3, 0x74, 0x74,
	0x4c, 0xd4, 0xa3, 0x09, 0xe1, 0x7f, 0xb0, 0x28, 0x78, 0x23, 0xea, 0x37, 0xc7, 0x27, 0xc4, 0xc1,
	0x8d, 0xc8, 0xe0, 0xc3, 0x5c, 0x8b, 0xcc, 0x76, 0x06, 0xfd, 0x29, 0xc1, 0x83, 0x1c, 0xfa, 0xc3,
	0x86, 0x8b, 0xda, 0x9f, 0xb9, 0xfe, 0x48, 0xec, 0x9f, 0xc0, 0x4b, 0xa2, 0x07, 0xf7, 0x96, 0xec,
	0x06, 0x96, 0x47, 0x3b, 0xb3, 0x6f, 0x8a, 0x58, 0x6f, 0x4e, 0x32, 0x57, 0x7d, 0xfa, 0x83, 0x5e,
	0xf3, 0xf2, 0x23, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x42, 0x48, 0x14, 0x27, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VEServiceClient is the client API for VEService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VEServiceClient interface {
	ValidateToken(ctx context.Context, in *VEValidateTokenRequest, opts ...grpc.CallOption) (*VEValidateTokenResponse, error)
}

type vEServiceClient struct {
	cc *grpc.ClientConn
}

func NewVEServiceClient(cc *grpc.ClientConn) VEServiceClient {
	return &vEServiceClient{cc}
}

func (c *vEServiceClient) ValidateToken(ctx context.Context, in *VEValidateTokenRequest, opts ...grpc.CallOption) (*VEValidateTokenResponse, error) {
	out := new(VEValidateTokenResponse)
	err := c.cc.Invoke(ctx, "/pb.VEService/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VEServiceServer is the server API for VEService service.
type VEServiceServer interface {
	ValidateToken(context.Context, *VEValidateTokenRequest) (*VEValidateTokenResponse, error)
}

func RegisterVEServiceServer(s *grpc.Server, srv VEServiceServer) {
	s.RegisterService(&_VEService_serviceDesc, srv)
}

func _VEService_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VEValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VEServiceServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VEService/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VEServiceServer).ValidateToken(ctx, req.(*VEValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VEService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VEService",
	HandlerType: (*VEServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateToken",
			Handler:    _VEService_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
