// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/user/private_customapi.proto

// User Custom Private API
//
// x-displayName: "User custom private API"
// Custom private APIs for user management

package user

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
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

// LastLoginUpdateRequest
//
// x-displayName: "Last Login Update Request"
// Request to update user login timestamp.
type LastLoginUpdateRequest struct {
	// user
	//
	// x-displayName: "User"
	// x-example: "user@company1.com"
	// User ID of the user. typically email id
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// tenant
	//
	// x-displayName: "Tenant"
	// x-example: "company1-as432s"
	// Tenant ID of the tenant user belongs to.
	Tenant string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// last_login_timestamp
	//
	// x-displayName: "Last Login Timestamp"
	// Last successful login timestamp of the user .
	LastLoginTimestamp *types.Timestamp `protobuf:"bytes,3,opt,name=last_login_timestamp,json=lastLoginTimestamp,proto3" json:"last_login_timestamp,omitempty"`
}

func (m *LastLoginUpdateRequest) Reset()      { *m = LastLoginUpdateRequest{} }
func (*LastLoginUpdateRequest) ProtoMessage() {}
func (*LastLoginUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce107763f3250c6, []int{0}
}
func (m *LastLoginUpdateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastLoginUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastLoginUpdateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastLoginUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastLoginUpdateRequest.Merge(m, src)
}
func (m *LastLoginUpdateRequest) XXX_Size() int {
	return m.Size()
}
func (m *LastLoginUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LastLoginUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LastLoginUpdateRequest proto.InternalMessageInfo

func (m *LastLoginUpdateRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *LastLoginUpdateRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *LastLoginUpdateRequest) GetLastLoginTimestamp() *types.Timestamp {
	if m != nil {
		return m.LastLoginTimestamp
	}
	return nil
}

// LastLoginUpdateResponse
//
// x-displayName: "Last Login Update Response"
type LastLoginUpdateResponse struct {
}

func (m *LastLoginUpdateResponse) Reset()      { *m = LastLoginUpdateResponse{} }
func (*LastLoginUpdateResponse) ProtoMessage() {}
func (*LastLoginUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce107763f3250c6, []int{1}
}
func (m *LastLoginUpdateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastLoginUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastLoginUpdateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastLoginUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastLoginUpdateResponse.Merge(m, src)
}
func (m *LastLoginUpdateResponse) XXX_Size() int {
	return m.Size()
}
func (m *LastLoginUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LastLoginUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LastLoginUpdateResponse proto.InternalMessageInfo

// PrivateCascadeDeleteRequest
//
// x-displayName: "Delete the User and Associated Namespace Roles"
// PrivateCascadeDeleteRequest is the request to delete the user along with the associated namespace role objects.
type PrivateCascadeDeleteRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-example: "value"
	// Value of namespace is always "system"
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Tenant name
	//
	// x-displayName: "Tenant name"
	// x-required
	// User deletion will be executed within this tenant.
	TenantName string `protobuf:"bytes,2,opt,name=tenant_name,json=tenantName,proto3" json:"tenant_name,omitempty"`
	// email of the user
	//
	// x-displayName: "Email"
	// x-example: "value"
	// x-required
	// email of the user requesting for
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (m *PrivateCascadeDeleteRequest) Reset()      { *m = PrivateCascadeDeleteRequest{} }
func (*PrivateCascadeDeleteRequest) ProtoMessage() {}
func (*PrivateCascadeDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce107763f3250c6, []int{2}
}
func (m *PrivateCascadeDeleteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrivateCascadeDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrivateCascadeDeleteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrivateCascadeDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateCascadeDeleteRequest.Merge(m, src)
}
func (m *PrivateCascadeDeleteRequest) XXX_Size() int {
	return m.Size()
}
func (m *PrivateCascadeDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateCascadeDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateCascadeDeleteRequest proto.InternalMessageInfo

func (m *PrivateCascadeDeleteRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PrivateCascadeDeleteRequest) GetTenantName() string {
	if m != nil {
		return m.TenantName
	}
	return ""
}

func (m *PrivateCascadeDeleteRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*LastLoginUpdateRequest)(nil), "ves.io.schema.user.LastLoginUpdateRequest")
	golang_proto.RegisterType((*LastLoginUpdateRequest)(nil), "ves.io.schema.user.LastLoginUpdateRequest")
	proto.RegisterType((*LastLoginUpdateResponse)(nil), "ves.io.schema.user.LastLoginUpdateResponse")
	golang_proto.RegisterType((*LastLoginUpdateResponse)(nil), "ves.io.schema.user.LastLoginUpdateResponse")
	proto.RegisterType((*PrivateCascadeDeleteRequest)(nil), "ves.io.schema.user.PrivateCascadeDeleteRequest")
	golang_proto.RegisterType((*PrivateCascadeDeleteRequest)(nil), "ves.io.schema.user.PrivateCascadeDeleteRequest")
}

func init() {
	proto.RegisterFile("ves.io/schema/user/private_customapi.proto", fileDescriptor_7ce107763f3250c6)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/user/private_customapi.proto", fileDescriptor_7ce107763f3250c6)
}

var fileDescriptor_7ce107763f3250c6 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x4f, 0x14, 0x3d,
	0x18, 0xde, 0x02, 0x1f, 0x09, 0x25, 0x5f, 0x34, 0x0d, 0x41, 0x18, 0x48, 0xc1, 0x3d, 0xc1, 0x1a,
	0xa6, 0x09, 0xde, 0xb8, 0x29, 0x1e, 0x34, 0x12, 0x25, 0x1b, 0x35, 0xc6, 0xc4, 0x6c, 0xba, 0xb3,
	0x2f, 0x43, 0x75, 0x66, 0x5a, 0xa7, 0x9d, 0x55, 0x62, 0x4c, 0x0c, 0xbf, 0xc0, 0xe8, 0x9f, 0xf0,
	0x3f, 0x70, 0x21, 0xf1, 0xe2, 0xc5, 0x48, 0xf4, 0xc2, 0x51, 0x66, 0x3d, 0x78, 0xe4, 0x27, 0x98,
	0xed, 0x74, 0x87, 0x5d, 0x98, 0xa8, 0xb7, 0xf6, 0x7d, 0x9f, 0x3e, 0x7d, 0x9e, 0xa7, 0x7d, 0x71,
	0xa3, 0x0b, 0xda, 0x17, 0x92, 0xe9, 0x60, 0x17, 0x62, 0xce, 0x32, 0x0d, 0x29, 0x53, 0xa9, 0xe8,
	0x72, 0x03, 0xad, 0x20, 0xd3, 0x46, 0xc6, 0x5c, 0x09, 0x5f, 0xa5, 0xd2, 0x48, 0x42, 0x0a, 0xac,
	0x5f, 0x60, 0xfd, 0x3e, 0xd6, 0x5b, 0x0b, 0x85, 0xd9, 0xcd, 0xda, 0x7e, 0x20, 0x63, 0x16, 0xca,
	0x50, 0x32, 0x0b, 0x6d, 0x67, 0x3b, 0x76, 0x67, 0x37, 0x76, 0x55, 0x50, 0x78, 0x8b, 0xa1, 0x94,
	0x61, 0x04, 0x8c, 0x2b, 0xc1, 0x78, 0x92, 0x48, 0xc3, 0x8d, 0x90, 0x89, 0x76, 0xdd, 0x25, 0xd7,
	0x2d, 0x39, 0x8c, 0x88, 0x41, 0x1b, 0x1e, 0x2b, 0x07, 0x58, 0x18, 0x55, 0x2b, 0xd5, 0xf0, 0xe9,
	0xf9, 0xd1, 0xa6, 0xd9, 0x53, 0x50, 0x12, 0x57, 0xb8, 0x94, 0xed, 0x67, 0x10, 0x18, 0x07, 0x58,
	0xad, 0x8a, 0x21, 0x6b, 0x47, 0x22, 0x38, 0x9f, 0x82, 0x47, 0x2b, 0xa0, 0x43, 0x77, 0xd5, 0xdf,
	0x23, 0x3c, 0xbb, 0xc5, 0xb5, 0xd9, 0x92, 0xa1, 0x48, 0x1e, 0xaa, 0x0e, 0x37, 0xd0, 0x84, 0x17,
	0x19, 0x68, 0x43, 0x08, 0x9e, 0xe8, 0xc3, 0xe7, 0xd0, 0x32, 0x5a, 0x99, 0x6a, 0xda, 0x35, 0x99,
	0xc5, 0x93, 0x06, 0x12, 0x9e, 0x98, 0xb9, 0x31, 0x5b, 0x75, 0x3b, 0xb2, 0x85, 0x67, 0x22, 0xae,
	0x4d, 0x2b, 0xea, 0xd3, 0xb4, 0xca, 0x20, 0xe6, 0xc6, 0x97, 0xd1, 0xca, 0xf4, 0xba, 0xe7, 0x17,
	0x51, 0xf9, 0x83, 0xa8, 0xfc, 0x07, 0x03, 0x44, 0x93, 0x44, 0x83, 0xdb, 0xcb, 0x5a, 0x7d, 0x1e,
	0x5f, 0xb9, 0xa0, 0x49, 0x2b, 0x99, 0x68, 0xa8, 0xa7, 0x78, 0x61, 0xbb, 0x78, 0xf0, 0x4d, 0xae,
	0x03, 0xde, 0x81, 0x5b, 0x10, 0xc1, 0x99, 0xe6, 0x45, 0x3c, 0x95, 0xf0, 0x18, 0xb4, 0xe2, 0x01,
	0x38, 0xe1, 0x67, 0x05, 0xb2, 0x84, 0xa7, 0x0b, 0xbd, 0xad, 0x7e, 0xcd, 0x59, 0xc0, 0x45, 0xe9,
	0x1e, 0x8f, 0x81, 0xcc, 0xe0, 0xff, 0x20, 0xe6, 0x22, 0xb2, 0xba, 0xa7, 0x9a, 0xc5, 0x66, 0xfd,
	0xcb, 0x38, 0xbe, 0xbc, 0x69, 0x73, 0x75, 0x57, 0xdf, 0xd8, 0xbe, 0x43, 0xbe, 0x22, 0x7c, 0xa9,
	0xd0, 0x56, 0x4a, 0x25, 0x0d, 0xff, 0xe2, 0x9f, 0xf3, 0xab, 0xd3, 0xf5, 0xae, 0xfd, 0x13, 0xd6,
	0xb9, 0x0e, 0xf6, 0xbf, 0xff, 0xfc, 0x30, 0xf6, 0xb4, 0xfe, 0x98, 0x8d, 0x1c, 0x62, 0x22, 0x31,
	0xa9, 0xd4, 0x0a, 0x02, 0xc3, 0x5e, 0xa6, 0xc2, 0xc0, 0x60, 0x18, 0x58, 0xf1, 0x0d, 0x58, 0x69,
	0x5d, 0x33, 0xbd, 0xa7, 0x0d, 0xc4, 0xf6, 0x03, 0x68, 0x96, 0x59, 0xfa, 0xd6, 0xd9, 0x93, 0x6d,
	0xa0, 0x06, 0xf9, 0x84, 0xf0, 0xff, 0x23, 0xa1, 0x12, 0x56, 0xa5, 0xf1, 0x0f, 0xf1, 0x7b, 0xab,
	0x55, 0x07, 0xce, 0x21, 0x9d, 0xa5, 0x47, 0xd6, 0xd2, 0x76, 0xfd, 0xee, 0xdf, 0x2c, 0x0d, 0x79,
	0x78, 0x5d, 0xae, 0xdf, 0x38, 0x23, 0x41, 0xc1, 0xdd, 0xea, 0x58, 0xf2, 0x0d, 0xd4, 0xf0, 0xae,
	0x1e, 0x1e, 0xa0, 0x89, 0x6f, 0x07, 0x68, 0xbe, 0x42, 0xc9, 0x7d, 0x3b, 0x44, 0x37, 0xf7, 0xd1,
	0xd1, 0x09, 0xad, 0x1d, 0x9f, 0xd0, 0xda, 0xe9, 0x09, 0x45, 0x6f, 0x73, 0x8a, 0x3e, 0xe6, 0x14,
	0x7d, 0xce, 0x29, 0x3a, 0xca, 0x29, 0xfa, 0x91, 0x53, 0xf4, 0x2b, 0xa7, 0xb5, 0xd3, 0x9c, 0xa2,
	0x77, 0x3d, 0x5a, 0x3b, 0xec, 0x51, 0x74, 0xd4, 0xa3, 0xb5, 0xe3, 0x1e, 0xad, 0x3d, 0xb9, 0x1d,
	0x4a, 0xf5, 0x3c, 0xf4, 0xbb, 0x32, 0x32, 0x90, 0xa6, 0x7d, 0x5a, 0x66, 0x17, 0x3b, 0x32, 0x8d,
	0xd7, 0x54, 0x2a, 0xbb, 0xa2, 0x03, 0xe9, 0xda, 0xa0, 0xcd, 0x54, 0x3b, 0x94, 0x0c, 0x5e, 0x19,
	0xe7, 0x6d, 0x68, 0x08, 0xdb, 0x93, 0x76, 0x16, 0xae, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x29,
	0xc1, 0x17, 0x20, 0xd3, 0x04, 0x00, 0x00,
}

func (this *LastLoginUpdateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LastLoginUpdateRequest)
	if !ok {
		that2, ok := that.(LastLoginUpdateRequest)
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
	if this.User != that1.User {
		return false
	}
	if this.Tenant != that1.Tenant {
		return false
	}
	if !this.LastLoginTimestamp.Equal(that1.LastLoginTimestamp) {
		return false
	}
	return true
}
func (this *LastLoginUpdateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LastLoginUpdateResponse)
	if !ok {
		that2, ok := that.(LastLoginUpdateResponse)
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
	return true
}
func (this *PrivateCascadeDeleteRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PrivateCascadeDeleteRequest)
	if !ok {
		that2, ok := that.(PrivateCascadeDeleteRequest)
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
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.TenantName != that1.TenantName {
		return false
	}
	if this.Email != that1.Email {
		return false
	}
	return true
}
func (this *LastLoginUpdateRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&user.LastLoginUpdateRequest{")
	s = append(s, "User: "+fmt.Sprintf("%#v", this.User)+",\n")
	s = append(s, "Tenant: "+fmt.Sprintf("%#v", this.Tenant)+",\n")
	if this.LastLoginTimestamp != nil {
		s = append(s, "LastLoginTimestamp: "+fmt.Sprintf("%#v", this.LastLoginTimestamp)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *LastLoginUpdateResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&user.LastLoginUpdateResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PrivateCascadeDeleteRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&user.PrivateCascadeDeleteRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "TenantName: "+fmt.Sprintf("%#v", this.TenantName)+",\n")
	s = append(s, "Email: "+fmt.Sprintf("%#v", this.Email)+",\n")
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
	// Update Last Login
	//
	// x-displayName: "Update Last Login"
	// API to update last login timestamp of user
	UpdateLastLogin(ctx context.Context, in *LastLoginUpdateRequest, opts ...grpc.CallOption) (*LastLoginUpdateResponse, error)
	// CascadeDelete
	//
	// x-displayName: "Delete User and Related Objects"
	// CascadeDelete deletes the user and associated namespace roles for this user.
	// Use this only if the user and its referenced objects need to be wiped out altogether.
	// Note: users will always be in the system namespace.
	CascadeDelete(ctx context.Context, in *PrivateCascadeDeleteRequest, opts ...grpc.CallOption) (*CascadeDeleteResponse, error)
}

type customPrivateAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomPrivateAPIClient(cc *grpc.ClientConn) CustomPrivateAPIClient {
	return &customPrivateAPIClient{cc}
}

func (c *customPrivateAPIClient) UpdateLastLogin(ctx context.Context, in *LastLoginUpdateRequest, opts ...grpc.CallOption) (*LastLoginUpdateResponse, error) {
	out := new(LastLoginUpdateResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.user.CustomPrivateAPI/UpdateLastLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customPrivateAPIClient) CascadeDelete(ctx context.Context, in *PrivateCascadeDeleteRequest, opts ...grpc.CallOption) (*CascadeDeleteResponse, error) {
	out := new(CascadeDeleteResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.user.CustomPrivateAPI/CascadeDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomPrivateAPIServer is the server API for CustomPrivateAPI service.
type CustomPrivateAPIServer interface {
	// Update Last Login
	//
	// x-displayName: "Update Last Login"
	// API to update last login timestamp of user
	UpdateLastLogin(context.Context, *LastLoginUpdateRequest) (*LastLoginUpdateResponse, error)
	// CascadeDelete
	//
	// x-displayName: "Delete User and Related Objects"
	// CascadeDelete deletes the user and associated namespace roles for this user.
	// Use this only if the user and its referenced objects need to be wiped out altogether.
	// Note: users will always be in the system namespace.
	CascadeDelete(context.Context, *PrivateCascadeDeleteRequest) (*CascadeDeleteResponse, error)
}

// UnimplementedCustomPrivateAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomPrivateAPIServer struct {
}

func (*UnimplementedCustomPrivateAPIServer) UpdateLastLogin(ctx context.Context, req *LastLoginUpdateRequest) (*LastLoginUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLastLogin not implemented")
}
func (*UnimplementedCustomPrivateAPIServer) CascadeDelete(ctx context.Context, req *PrivateCascadeDeleteRequest) (*CascadeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CascadeDelete not implemented")
}

func RegisterCustomPrivateAPIServer(s *grpc.Server, srv CustomPrivateAPIServer) {
	s.RegisterService(&_CustomPrivateAPI_serviceDesc, srv)
}

func _CustomPrivateAPI_UpdateLastLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastLoginUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomPrivateAPIServer).UpdateLastLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.user.CustomPrivateAPI/UpdateLastLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomPrivateAPIServer).UpdateLastLogin(ctx, req.(*LastLoginUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomPrivateAPI_CascadeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivateCascadeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomPrivateAPIServer).CascadeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.user.CustomPrivateAPI/CascadeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomPrivateAPIServer).CascadeDelete(ctx, req.(*PrivateCascadeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomPrivateAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.user.CustomPrivateAPI",
	HandlerType: (*CustomPrivateAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateLastLogin",
			Handler:    _CustomPrivateAPI_UpdateLastLogin_Handler,
		},
		{
			MethodName: "CascadeDelete",
			Handler:    _CustomPrivateAPI_CascadeDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/user/private_customapi.proto",
}

func (m *LastLoginUpdateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastLoginUpdateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastLoginUpdateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastLoginTimestamp != nil {
		{
			size, err := m.LastLoginTimestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPrivateCustomapi(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Tenant) > 0 {
		i -= len(m.Tenant)
		copy(dAtA[i:], m.Tenant)
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.Tenant)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LastLoginUpdateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastLoginUpdateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastLoginUpdateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PrivateCascadeDeleteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrivateCascadeDeleteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrivateCascadeDeleteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TenantName) > 0 {
		i -= len(m.TenantName)
		copy(dAtA[i:], m.TenantName)
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.TenantName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPrivateCustomapi(dAtA, i, uint64(len(m.Namespace)))
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
func (m *LastLoginUpdateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	l = len(m.Tenant)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	if m.LastLoginTimestamp != nil {
		l = m.LastLoginTimestamp.Size()
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	return n
}

func (m *LastLoginUpdateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PrivateCascadeDeleteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	l = len(m.TenantName)
	if l > 0 {
		n += 1 + l + sovPrivateCustomapi(uint64(l))
	}
	l = len(m.Email)
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
func (this *LastLoginUpdateRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LastLoginUpdateRequest{`,
		`User:` + fmt.Sprintf("%v", this.User) + `,`,
		`Tenant:` + fmt.Sprintf("%v", this.Tenant) + `,`,
		`LastLoginTimestamp:` + strings.Replace(fmt.Sprintf("%v", this.LastLoginTimestamp), "Timestamp", "types.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LastLoginUpdateResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LastLoginUpdateResponse{`,
		`}`,
	}, "")
	return s
}
func (this *PrivateCascadeDeleteRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PrivateCascadeDeleteRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`TenantName:` + fmt.Sprintf("%v", this.TenantName) + `,`,
		`Email:` + fmt.Sprintf("%v", this.Email) + `,`,
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
func (m *LastLoginUpdateRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LastLoginUpdateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastLoginUpdateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
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
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenant", wireType)
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
			m.Tenant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastLoginTimestamp", wireType)
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
			if m.LastLoginTimestamp == nil {
				m.LastLoginTimestamp = &types.Timestamp{}
			}
			if err := m.LastLoginTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *LastLoginUpdateResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LastLoginUpdateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastLoginUpdateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *PrivateCascadeDeleteRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PrivateCascadeDeleteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrivateCascadeDeleteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
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
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TenantName", wireType)
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
			m.TenantName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
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
			m.Email = string(dAtA[iNdEx:postIndex])
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
