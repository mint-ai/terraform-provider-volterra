// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/signup/private_customapi.proto

package signup

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// TenantByUserCnameRequest
//
// x-displayName: "TenantByUserCnameRequest"
// Request for finding the tenant by parameters
type SignupFreemiumSSOTenantRequest struct {
	// User email
	//
	// x-displayName: "User name"
	// User email of the domain owner for specific tenant.
	UserEmail string `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
}

func (m *SignupFreemiumSSOTenantRequest) Reset()      { *m = SignupFreemiumSSOTenantRequest{} }
func (*SignupFreemiumSSOTenantRequest) ProtoMessage() {}
func (*SignupFreemiumSSOTenantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_860cf8f7e01427bf, []int{0}
}
func (m *SignupFreemiumSSOTenantRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignupFreemiumSSOTenantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignupFreemiumSSOTenantRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignupFreemiumSSOTenantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupFreemiumSSOTenantRequest.Merge(m, src)
}
func (m *SignupFreemiumSSOTenantRequest) XXX_Size() int {
	return m.Size()
}
func (m *SignupFreemiumSSOTenantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupFreemiumSSOTenantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignupFreemiumSSOTenantRequest proto.InternalMessageInfo

func (m *SignupFreemiumSSOTenantRequest) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

// TenantByUserCnameResponse
//
// x-displayName: "TenantByUserCnameResponse"
// Response of tenant information by finding it with parameters.
type SignupFreemiumSSOTenantResponse struct {
	// ResponseStatus
	//
	// x-displayName: "response_status"
	// ResponseStatus to identify the status of the response
	ResponseStatus *schema.ResponseMeta `protobuf:"bytes,1,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"`
	// tenant_id
	//
	// x-displayName: "TenantId"
	// x-example: "abc-def"
	// TenantId will denote the name of the tenant.
	TenantId string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (m *SignupFreemiumSSOTenantResponse) Reset()      { *m = SignupFreemiumSSOTenantResponse{} }
func (*SignupFreemiumSSOTenantResponse) ProtoMessage() {}
func (*SignupFreemiumSSOTenantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_860cf8f7e01427bf, []int{1}
}
func (m *SignupFreemiumSSOTenantResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignupFreemiumSSOTenantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignupFreemiumSSOTenantResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignupFreemiumSSOTenantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupFreemiumSSOTenantResponse.Merge(m, src)
}
func (m *SignupFreemiumSSOTenantResponse) XXX_Size() int {
	return m.Size()
}
func (m *SignupFreemiumSSOTenantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupFreemiumSSOTenantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignupFreemiumSSOTenantResponse proto.InternalMessageInfo

func (m *SignupFreemiumSSOTenantResponse) GetResponseStatus() *schema.ResponseMeta {
	if m != nil {
		return m.ResponseStatus
	}
	return nil
}

func (m *SignupFreemiumSSOTenantResponse) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func init() {
	proto.RegisterType((*SignupFreemiumSSOTenantRequest)(nil), "ves.io.schema.signup.SignupFreemiumSSOTenantRequest")
	golang_proto.RegisterType((*SignupFreemiumSSOTenantRequest)(nil), "ves.io.schema.signup.SignupFreemiumSSOTenantRequest")
	proto.RegisterType((*SignupFreemiumSSOTenantResponse)(nil), "ves.io.schema.signup.SignupFreemiumSSOTenantResponse")
	golang_proto.RegisterType((*SignupFreemiumSSOTenantResponse)(nil), "ves.io.schema.signup.SignupFreemiumSSOTenantResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/signup/private_customapi.proto", fileDescriptor_860cf8f7e01427bf)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/signup/private_customapi.proto", fileDescriptor_860cf8f7e01427bf)
}

var fileDescriptor_860cf8f7e01427bf = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6b, 0x13, 0x41,
	0x18, 0xde, 0x89, 0x45, 0xcd, 0x0a, 0xb5, 0x2c, 0x82, 0x31, 0x29, 0xa3, 0xe4, 0x24, 0x6a, 0x66,
	0xa0, 0x2a, 0x88, 0xe0, 0xc1, 0x56, 0x85, 0x0a, 0xd2, 0x92, 0x78, 0xf2, 0xb2, 0x4e, 0x92, 0x37,
	0xdb, 0xc1, 0xec, 0xbe, 0xe3, 0x7c, 0xa4, 0xf6, 0x56, 0xa4, 0x3f, 0x40, 0xf0, 0x4f, 0xf8, 0x1b,
	0xf4, 0x92, 0xa3, 0xe0, 0x25, 0xe0, 0xa5, 0x47, 0xb3, 0x11, 0xd1, 0x5b, 0x7f, 0x82, 0x64, 0x36,
	0x81, 0x46, 0x5a, 0xa1, 0xb7, 0x77, 0xe6, 0xf9, 0xe0, 0x79, 0x3f, 0xc2, 0x3b, 0x03, 0x30, 0x4c,
	0x22, 0x37, 0x9d, 0x1d, 0x48, 0x05, 0x37, 0x32, 0xc9, 0x9c, 0xe2, 0x4a, 0xcb, 0x81, 0xb0, 0x10,
	0x77, 0x9c, 0xb1, 0x98, 0x0a, 0x25, 0x99, 0xd2, 0x68, 0x31, 0xba, 0x52, 0xb0, 0x59, 0xc1, 0x66,
	0x05, 0xbb, 0xda, 0x48, 0xa4, 0xdd, 0x71, 0x6d, 0xd6, 0xc1, 0x94, 0x27, 0x98, 0x20, 0xf7, 0xe4,
	0xb6, 0xeb, 0xf9, 0x97, 0x7f, 0xf8, 0xaa, 0x30, 0xa9, 0xae, 0x26, 0x88, 0x49, 0x1f, 0xb8, 0x50,
	0x92, 0x8b, 0x2c, 0x43, 0x2b, 0xac, 0xc4, 0xcc, 0xcc, 0xd0, 0xda, 0x62, 0x20, 0x54, 0xc7, 0xc1,
	0xdb, 0x27, 0xa7, 0x75, 0xed, 0xbe, 0xec, 0xfc, 0x1b, 0xb6, 0x7a, 0x6d, 0x91, 0x6c, 0xf7, 0x14,
	0xcc, 0x7d, 0x56, 0x17, 0xa1, 0x81, 0xe8, 0xcb, 0xae, 0xb0, 0x50, 0xa0, 0xf5, 0xad, 0x90, 0xb6,
	0xbc, 0xf3, 0x33, 0x0d, 0x90, 0x4a, 0x97, 0xb6, 0x5a, 0x5b, 0x2f, 0x21, 0x13, 0x99, 0x6d, 0xc2,
	0x5b, 0x07, 0xc6, 0x46, 0x8d, 0x30, 0x74, 0x06, 0x74, 0x0c, 0xa9, 0x90, 0xfd, 0x0a, 0xb9, 0x41,
	0x6e, 0x96, 0xd7, 0x97, 0x3f, 0xff, 0x19, 0x9e, 0x2b, 0xeb, 0x0b, 0x2b, 0xa4, 0xb2, 0x5f, 0x7a,
	0x4d, 0x9a, 0xe5, 0x29, 0xe3, 0xe9, 0x94, 0x50, 0x3f, 0x20, 0xe1, 0xf5, 0x53, 0x1d, 0x8d, 0xc2,
	0xcc, 0x40, 0xf4, 0x24, 0xbc, 0xac, 0x67, 0x75, 0x6c, 0xac, 0xb0, 0xce, 0x78, 0xdf, 0x4b, 0x6b,
	0x35, 0xb6, 0x38, 0xf4, 0xb9, 0xe2, 0x05, 0x58, 0xd1, 0x5c, 0x9e, 0x6b, 0x5a, 0x5e, 0x12, 0xd5,
	0xc2, 0xb2, 0xf5, 0xbe, 0xb1, 0xec, 0x56, 0x4a, 0xd3, 0x5c, 0xcd, 0x8b, 0xc5, 0xc7, 0x66, 0x77,
	0xed, 0x17, 0x09, 0x57, 0x36, 0xfc, 0x90, 0xb6, 0x8b, 0xfd, 0x3e, 0xde, 0xde, 0x8c, 0xbe, 0x91,
	0xf0, 0xea, 0x29, 0xd9, 0xa2, 0x7b, 0xec, 0xa4, 0x7d, 0xb3, 0xff, 0x0f, 0xa7, 0x7a, 0xff, 0x8c,
	0xaa, 0xa2, 0x81, 0xfa, 0xc6, 0xfb, 0xef, 0x3f, 0x3f, 0x96, 0x1e, 0xd5, 0x1f, 0xf0, 0x99, 0x1c,
	0xf6, 0x76, 0x05, 0x97, 0x99, 0xd5, 0x68, 0x14, 0x74, 0x2c, 0xdf, 0xd5, 0xd2, 0x02, 0x2f, 0xba,
	0x99, 0xad, 0x3e, 0xee, 0xcd, 0xec, 0x62, 0x63, 0xf0, 0x21, 0xb9, 0x55, 0x5d, 0x1a, 0x7e, 0x21,
	0x4b, 0xeb, 0x07, 0x64, 0x34, 0xa6, 0xc1, 0xe1, 0x98, 0x06, 0x47, 0x63, 0x4a, 0xf6, 0x73, 0x4a,
	0x3e, 0xe5, 0x94, 0x7c, 0xcd, 0x29, 0x19, 0xe5, 0x94, 0xfc, 0xc8, 0x29, 0xf9, 0x9d, 0xd3, 0xe0,
	0x28, 0xa7, 0xe4, 0xc3, 0x84, 0x06, 0xc3, 0x09, 0x25, 0xa3, 0x09, 0x0d, 0x0e, 0x27, 0x34, 0x78,
	0xf5, 0x3c, 0x41, 0xf5, 0x26, 0x61, 0x03, 0xec, 0x5b, 0xd0, 0x5a, 0x30, 0x67, 0xb8, 0x2f, 0x7a,
	0xa8, 0xd3, 0x86, 0xd2, 0x38, 0x90, 0x5d, 0xd0, 0x8d, 0x39, 0xcc, 0x55, 0x3b, 0x41, 0x0e, 0xef,
	0xec, 0xfc, 0x20, 0x8f, 0xdf, 0x65, 0xfb, 0xbc, 0x3f, 0xa7, 0xbb, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xe9, 0x9f, 0x93, 0x18, 0x64, 0x03, 0x00, 0x00,
}

func (this *SignupFreemiumSSOTenantRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SignupFreemiumSSOTenantRequest)
	if !ok {
		that2, ok := that.(SignupFreemiumSSOTenantRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.UserEmail != that1.UserEmail {
		return false
	}
	return true
}
func (this *SignupFreemiumSSOTenantResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SignupFreemiumSSOTenantResponse)
	if !ok {
		that2, ok := that.(SignupFreemiumSSOTenantResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ResponseStatus.Equal(that1.ResponseStatus) {
		return false
	}
	if this.TenantId != that1.TenantId {
		return false
	}
	return true
}
func (this *SignupFreemiumSSOTenantRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&signup.SignupFreemiumSSOTenantRequest{")
	s = append(s, "UserEmail: "+fmt.Sprintf("%#v", this.UserEmail)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SignupFreemiumSSOTenantResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&signup.SignupFreemiumSSOTenantResponse{")
	if this.ResponseStatus != nil {
		s = append(s, "ResponseStatus: "+fmt.Sprintf("%#v", this.ResponseStatus)+",\n")
	}
	s = append(s, "TenantId: "+fmt.Sprintf("%#v", this.TenantId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPrivateCustomapi(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomPrivateAPIClient is the client API for CustomPrivateAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomPrivateAPIClient interface {
	// SignupFreemiumSSOTenant
	//
	// x-displayName: "Signup Freemium SSO Tenant"
	// SignupFreemiumSSOTenant signs up a new freemium sso tenant with given user details if the user
	// is found in IAM
	SignupFreemiumSSOTenant(ctx context.Context, in *SignupFreemiumSSOTenantRequest, opts ...grpc.CallOption) (*SignupFreemiumSSOTenantResponse, error)
}

type customPrivateAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomPrivateAPIClient(cc *grpc.ClientConn) CustomPrivateAPIClient {
	return &customPrivateAPIClient{cc}
}

func (c *customPrivateAPIClient) SignupFreemiumSSOTenant(ctx context.Context, in *SignupFreemiumSSOTenantRequest, opts ...grpc.CallOption) (*SignupFreemiumSSOTenantResponse, error) {
	out := new(SignupFreemiumSSOTenantResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.signup.CustomPrivateAPI/SignupFreemiumSSOTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomPrivateAPIServer is the server API for CustomPrivateAPI service.
type CustomPrivateAPIServer interface {
	// SignupFreemiumSSOTenant
	//
	// x-displayName: "Signup Freemium SSO Tenant"
	// SignupFreemiumSSOTenant signs up a new freemium sso tenant with given user details if the user
	// is found in IAM
	SignupFreemiumSSOTenant(context.Context, *SignupFreemiumSSOTenantRequest) (*SignupFreemiumSSOTenantResponse, error)
}

// UnimplementedCustomPrivateAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomPrivateAPIServer struct {
}

func (*UnimplementedCustomPrivateAPIServer) SignupFreemiumSSOTenant(ctx context.Context, req *SignupFreemiumSSOTenantRequest) (*SignupFreemiumSSOTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignupFreemiumSSOTenant not implemented")
}

func RegisterCustomPrivateAPIServer(s *grpc.Server, srv CustomPrivateAPIServer) {
	s.RegisterService(&_CustomPrivateAPI_serviceDesc, srv)
}

func _CustomPrivateAPI_SignupFreemiumSSOTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupFreemiumSSOTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomPrivateAPIServer).SignupFreemiumSSOTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.signup.CustomPrivateAPI/SignupFreemiumSSOTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomPrivateAPIServer).SignupFreemiumSSOTenant(ctx, req.(*SignupFreemiumSSOTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomPrivateAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.signup.CustomPrivateAPI",
	HandlerType: (*CustomPrivateAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignupFreemiumSSOTenant",
			Handler:    _CustomPrivateAPI_SignupFreemiumSSOTenant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/signup/private_customapi.proto",
}

func (m *SignupFreemiumSSOTenantRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignupFreemiumSSOTenantRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignupFreemiumSSOTenantRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UserEmail) > 0 {
		i -= len(m.UserEmail)
		copy(dAtA[i:], m.UserEmail)
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.UserEmail)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SignupFreemiumSSOTenantResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignupFreemiumSSOTenantResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignupFreemiumSSOTenantResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TenantId) > 0 {
		i -= len(m.TenantId)
		copy(dAtA[i:], m.TenantId)
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.TenantId)))
		i--
		dAtA[i] = 0x12
	}
	if m.ResponseStatus != nil {
		{
			size, err := m.ResponseStatus.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPrivateCustomapi(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPrivateCustomapi(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrivateCustomapi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignupFreemiumSSOTenantRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserEmail)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	return n
}

func (m *SignupFreemiumSSOTenantResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResponseStatus != nil {
		l = m.ResponseStatus.Size()
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	l = len(m.TenantId)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	return n
}

func sovPrivateCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrivateCustomapi(x uint64) (n int) {
	return sovPrivateCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SignupFreemiumSSOTenantRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SignupFreemiumSSOTenantRequest{`,
		`UserEmail:` + fmt.Sprintf("%v", this.UserEmail) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SignupFreemiumSSOTenantResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SignupFreemiumSSOTenantResponse{`,
		`ResponseStatus:` + strings.Replace(fmt.Sprintf("%v", this.ResponseStatus), "ResponseMeta", "schema.ResponseMeta", 1) + `,`,
		`TenantId:` + fmt.Sprintf("%v", this.TenantId) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPrivateCustomapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SignupFreemiumSSOTenantRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrivateCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignupFreemiumSSOTenantRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignupFreemiumSSOTenantRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserEmail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserEmail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrivateCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SignupFreemiumSSOTenantResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrivateCustomapi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignupFreemiumSSOTenantResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignupFreemiumSSOTenantResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResponseStatus == nil {
				m.ResponseStatus = &schema.ResponseMeta{}
			}
			if err := m.ResponseStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TenantId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrivateCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TenantId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrivateCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPrivateCustomapi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPrivateCustomapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrivateCustomapi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPrivateCustomapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPrivateCustomapi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPrivateCustomapi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrivateCustomapi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrivateCustomapi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrivateCustomapi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrivateCustomapi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrivateCustomapi = fmt.Errorf("proto: unexpected end of group")
)
