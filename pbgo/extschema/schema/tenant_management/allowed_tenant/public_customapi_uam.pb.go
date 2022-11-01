// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/tenant_management/allowed_tenant/public_customapi_uam.proto

// Allowed Tenant
//
// x-displayName: "Allowed Tenant"
// Additional public APIs for allowed_tenant config object.

package allowed_tenant

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
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
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

// GetSupportTenantAccessReq
//
// x-displayName: "Get Support Tenant Access Request"
// Request to get access control configurations for a support tenant.
type GetSupportTenantAccessReq struct {
}

func (m *GetSupportTenantAccessReq) Reset()      { *m = GetSupportTenantAccessReq{} }
func (*GetSupportTenantAccessReq) ProtoMessage() {}
func (*GetSupportTenantAccessReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9792f429656e3beb, []int{0}
}
func (m *GetSupportTenantAccessReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSupportTenantAccessReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSupportTenantAccessReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSupportTenantAccessReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSupportTenantAccessReq.Merge(m, src)
}
func (m *GetSupportTenantAccessReq) XXX_Size() int {
	return m.Size()
}
func (m *GetSupportTenantAccessReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSupportTenantAccessReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetSupportTenantAccessReq proto.InternalMessageInfo

// GetSupportTenantAccessResp
//
// x-displayName: "Get Support Tenant Access Response"
// Response to get access control configurations for a support tenant.
type GetSupportTenantAccessResp struct {
	// access_config
	//
	// x-displayName: "Access Config"
	// Allowed access configuration details for the tenant.
	AccessConfig *AllowedAccessConfig `protobuf:"bytes,1,opt,name=access_config,json=accessConfig,proto3" json:"access_config,omitempty"`
}

func (m *GetSupportTenantAccessResp) Reset()      { *m = GetSupportTenantAccessResp{} }
func (*GetSupportTenantAccessResp) ProtoMessage() {}
func (*GetSupportTenantAccessResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9792f429656e3beb, []int{1}
}
func (m *GetSupportTenantAccessResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSupportTenantAccessResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSupportTenantAccessResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSupportTenantAccessResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSupportTenantAccessResp.Merge(m, src)
}
func (m *GetSupportTenantAccessResp) XXX_Size() int {
	return m.Size()
}
func (m *GetSupportTenantAccessResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSupportTenantAccessResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetSupportTenantAccessResp proto.InternalMessageInfo

func (m *GetSupportTenantAccessResp) GetAccessConfig() *AllowedAccessConfig {
	if m != nil {
		return m.AccessConfig
	}
	return nil
}

// UpdateSupportTenantAccessReq
//
// x-displayName: "Support Tenant Access Update Request"
// Request to update access control configurations for a support tenant.
type UpdateSupportTenantAccessReq struct {
	// access_config
	//
	// x-displayName: "Access Config"
	// Allowed access configuration details for the tenant.
	AccessConfig *AllowedAccessConfig `protobuf:"bytes,2,opt,name=access_config,json=accessConfig,proto3" json:"access_config,omitempty"`
}

func (m *UpdateSupportTenantAccessReq) Reset()      { *m = UpdateSupportTenantAccessReq{} }
func (*UpdateSupportTenantAccessReq) ProtoMessage() {}
func (*UpdateSupportTenantAccessReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9792f429656e3beb, []int{2}
}
func (m *UpdateSupportTenantAccessReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateSupportTenantAccessReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateSupportTenantAccessReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateSupportTenantAccessReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSupportTenantAccessReq.Merge(m, src)
}
func (m *UpdateSupportTenantAccessReq) XXX_Size() int {
	return m.Size()
}
func (m *UpdateSupportTenantAccessReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSupportTenantAccessReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSupportTenantAccessReq proto.InternalMessageInfo

func (m *UpdateSupportTenantAccessReq) GetAccessConfig() *AllowedAccessConfig {
	if m != nil {
		return m.AccessConfig
	}
	return nil
}

// UpdateSupportTenantAccessResp
//
// x-displayName: "Support Tenant Access Update Response"
// Response to update access control configurations for a support tenant.
type UpdateSupportTenantAccessResp struct {
	// access_config
	//
	// x-displayName: "Access Config"
	// Allowed access configuration details for the tenant.
	AccessConfig *AllowedAccessConfig `protobuf:"bytes,1,opt,name=access_config,json=accessConfig,proto3" json:"access_config,omitempty"`
}

func (m *UpdateSupportTenantAccessResp) Reset()      { *m = UpdateSupportTenantAccessResp{} }
func (*UpdateSupportTenantAccessResp) ProtoMessage() {}
func (*UpdateSupportTenantAccessResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9792f429656e3beb, []int{3}
}
func (m *UpdateSupportTenantAccessResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateSupportTenantAccessResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateSupportTenantAccessResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateSupportTenantAccessResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSupportTenantAccessResp.Merge(m, src)
}
func (m *UpdateSupportTenantAccessResp) XXX_Size() int {
	return m.Size()
}
func (m *UpdateSupportTenantAccessResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSupportTenantAccessResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSupportTenantAccessResp proto.InternalMessageInfo

func (m *UpdateSupportTenantAccessResp) GetAccessConfig() *AllowedAccessConfig {
	if m != nil {
		return m.AccessConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*GetSupportTenantAccessReq)(nil), "ves.io.schema.tenant_management.allowed_tenant.GetSupportTenantAccessReq")
	golang_proto.RegisterType((*GetSupportTenantAccessReq)(nil), "ves.io.schema.tenant_management.allowed_tenant.GetSupportTenantAccessReq")
	proto.RegisterType((*GetSupportTenantAccessResp)(nil), "ves.io.schema.tenant_management.allowed_tenant.GetSupportTenantAccessResp")
	golang_proto.RegisterType((*GetSupportTenantAccessResp)(nil), "ves.io.schema.tenant_management.allowed_tenant.GetSupportTenantAccessResp")
	proto.RegisterType((*UpdateSupportTenantAccessReq)(nil), "ves.io.schema.tenant_management.allowed_tenant.UpdateSupportTenantAccessReq")
	golang_proto.RegisterType((*UpdateSupportTenantAccessReq)(nil), "ves.io.schema.tenant_management.allowed_tenant.UpdateSupportTenantAccessReq")
	proto.RegisterType((*UpdateSupportTenantAccessResp)(nil), "ves.io.schema.tenant_management.allowed_tenant.UpdateSupportTenantAccessResp")
	golang_proto.RegisterType((*UpdateSupportTenantAccessResp)(nil), "ves.io.schema.tenant_management.allowed_tenant.UpdateSupportTenantAccessResp")
}

func init() {
	proto.RegisterFile("ves.io/schema/tenant_management/allowed_tenant/public_customapi_uam.proto", fileDescriptor_9792f429656e3beb)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/tenant_management/allowed_tenant/public_customapi_uam.proto", fileDescriptor_9792f429656e3beb)
}

var fileDescriptor_9792f429656e3beb = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xbf, 0x6f, 0xd4, 0x3e,
	0x18, 0xc6, 0xcf, 0xf7, 0xed, 0x17, 0x41, 0x00, 0x09, 0x65, 0x40, 0xd7, 0xb4, 0x98, 0x12, 0x16,
	0x84, 0x48, 0x2c, 0x8a, 0x58, 0xba, 0xa0, 0x6b, 0x91, 0x50, 0x2b, 0x10, 0xa8, 0xc0, 0xc2, 0x12,
	0x39, 0xc9, 0x7b, 0x39, 0xd3, 0xc4, 0x36, 0xb1, 0x73, 0x6d, 0x37, 0xd4, 0x01, 0xa9, 0x1b, 0x12,
	0x8c, 0x88, 0x19, 0xfe, 0x86, 0x2e, 0xdd, 0xca, 0x84, 0x2a, 0x58, 0x3a, 0xd2, 0x1c, 0x03, 0x63,
	0xff, 0x04, 0xd4, 0xe4, 0x2a, 0x5d, 0xd2, 0x1f, 0xe2, 0x44, 0xc5, 0x96, 0x57, 0x8f, 0xfd, 0xf8,
	0xe3, 0x27, 0x7e, 0x5f, 0x63, 0xbe, 0x07, 0xca, 0x65, 0x82, 0xa8, 0xa0, 0x0b, 0x09, 0x25, 0x1a,
	0x38, 0xe5, 0xda, 0x4b, 0x28, 0xa7, 0x11, 0x24, 0xc0, 0x35, 0xa1, 0x71, 0x2c, 0x96, 0x21, 0xf4,
	0x4a, 0x85, 0xc8, 0xcc, 0x8f, 0x59, 0xe0, 0x05, 0x99, 0xd2, 0x22, 0xa1, 0x92, 0x79, 0x19, 0x4d,
	0x5c, 0x99, 0x0a, 0x2d, 0x4c, 0xb7, 0xb4, 0x72, 0x4b, 0x2b, 0xf7, 0x90, 0x95, 0x5b, 0xb5, 0xb2,
	0x9c, 0x88, 0xe9, 0x6e, 0xe6, 0xbb, 0x81, 0x48, 0x48, 0x24, 0x22, 0x41, 0x0a, 0x1b, 0x3f, 0xeb,
	0x14, 0x55, 0x51, 0x14, 0x5f, 0xa5, 0xbd, 0x35, 0x19, 0x09, 0x11, 0xc5, 0x40, 0xa8, 0x64, 0x84,
	0x72, 0x2e, 0x34, 0xd5, 0x4c, 0x70, 0x35, 0x50, 0x27, 0xaa, 0xf7, 0x10, 0x72, 0x58, 0x9c, 0x19,
	0xf1, 0x92, 0x7a, 0x55, 0xc2, 0xc1, 0xde, 0xf1, 0xda, 0xde, 0x21, 0x69, 0xb2, 0x2a, 0xf5, 0x68,
	0xcc, 0x42, 0xaa, 0x61, 0xa0, 0xda, 0x35, 0x15, 0x14, 0xf0, 0x5e, 0x0d, 0x6c, 0xaa, 0xb6, 0x86,
	0xc1, 0xb2, 0x57, 0x5d, 0x71, 0xf5, 0xf0, 0x0a, 0x35, 0x0c, 0x61, 0x4f, 0x18, 0xe3, 0x0f, 0x40,
	0x3f, 0xcd, 0xa4, 0x14, 0xa9, 0x7e, 0x56, 0xf0, 0xb7, 0x83, 0x00, 0x94, 0x5a, 0x84, 0x57, 0xf6,
	0x1b, 0x64, 0x58, 0xc7, 0xa9, 0x4a, 0x9a, 0x5d, 0xe3, 0x22, 0x2d, 0x2a, 0x2f, 0x10, 0xbc, 0xc3,
	0xa2, 0x16, 0x9a, 0x42, 0x37, 0xce, 0x4f, 0xcf, 0x8d, 0xf8, 0x27, 0xdd, 0x76, 0x59, 0x96, 0xce,
	0x73, 0x85, 0xd5, 0xe2, 0x05, 0x3a, 0x54, 0xd9, 0xef, 0x91, 0x31, 0xf9, 0x5c, 0xee, 0xa7, 0x73,
	0x34, 0xa9, 0xc9, 0xeb, 0x28, 0xcd, 0x53, 0x43, 0x99, 0x1d, 0xdb, 0xdb, 0x40, 0xa8, 0x0a, 0xb4,
	0x30, 0x76, 0x16, 0x5d, 0x6a, 0xda, 0xeb, 0xc8, 0xb8, 0x72, 0x02, 0xd6, 0xbf, 0x8c, 0x68, 0x7a,
	0xfd, 0x7f, 0xe3, 0xdc, 0x5c, 0xd1, 0x56, 0xed, 0x27, 0xf3, 0xe6, 0xe7, 0xa6, 0x31, 0x7e, 0x2c,
	0x99, 0xf9, 0x70, 0xd4, 0xe3, 0x4f, 0xca, 0xde, 0x7a, 0x74, 0x8a, 0x6e, 0x4a, 0xda, 0x22, 0xdf,
	0x6a, 0x5d, 0xeb, 0xdc, 0x5d, 0x09, 0x1c, 0x9f, 0x2a, 0xb8, 0xb5, 0xff, 0xc6, 0x1d, 0x26, 0x9c,
	0x8c, 0x26, 0x0e, 0x0d, 0x13, 0xc6, 0x9d, 0xe5, 0x94, 0x69, 0x58, 0xfb, 0xfe, 0xf3, 0x5d, 0xf3,
	0xbe, 0x7d, 0x6f, 0x30, 0x51, 0x08, 0xa7, 0x09, 0x28, 0x49, 0x03, 0x50, 0x44, 0xad, 0x2a, 0x0d,
	0x49, 0xad, 0x29, 0x15, 0x51, 0xe5, 0x79, 0xce, 0xa0, 0x49, 0xcb, 0xf0, 0x66, 0xd0, 0x4d, 0xf3,
	0x43, 0xd3, 0xb8, 0x7c, 0xf4, 0x2b, 0x37, 0xe7, 0x47, 0xbd, 0xda, 0xb1, 0xbd, 0x64, 0x2d, 0x9c,
	0x96, 0x95, 0x92, 0xf6, 0x52, 0xbe, 0xd5, 0x9a, 0x3a, 0x29, 0xa2, 0x14, 0x68, 0x58, 0x24, 0xd4,
	0x36, 0xff, 0x36, 0x21, 0xeb, 0xf6, 0xe6, 0x06, 0xfa, 0xef, 0xdb, 0x06, 0xba, 0x5e, 0xe5, 0xaf,
	0xd1, 0x3e, 0xf6, 0x5f, 0x42, 0xa0, 0xd7, 0xbe, 0xb6, 0x9a, 0x5d, 0x34, 0xfb, 0x11, 0x6d, 0xef,
	0xe2, 0xc6, 0xce, 0x2e, 0x6e, 0xec, 0xed, 0x62, 0xf4, 0x3a, 0xc7, 0xe8, 0x53, 0x8e, 0xd1, 0x97,
	0x1c, 0xa3, 0xed, 0x1c, 0xa3, 0x1f, 0x39, 0x46, 0xbf, 0x72, 0xdc, 0xd8, 0xcb, 0x31, 0x7a, 0xdb,
	0xc7, 0x8d, 0xcd, 0x3e, 0x46, 0xdb, 0x7d, 0xdc, 0xd8, 0xe9, 0xe3, 0xc6, 0x8b, 0x30, 0x12, 0x72,
	0x29, 0x72, 0x7b, 0x22, 0xd6, 0x90, 0xa6, 0xd4, 0xcd, 0x14, 0x29, 0x3e, 0x3a, 0x22, 0x4d, 0x1c,
	0x99, 0x8a, 0x1e, 0x0b, 0x21, 0x75, 0x0e, 0x64, 0x22, 0xfd, 0x48, 0x10, 0x58, 0xd1, 0x83, 0xf9,
	0xf6, 0x87, 0x13, 0xda, 0x3f, 0x53, 0x0c, 0xbf, 0x3b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x40,
	0xfc, 0x22, 0x9c, 0xbf, 0x06, 0x00, 0x00,
}

func (this *GetSupportTenantAccessReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSupportTenantAccessReq)
	if !ok {
		that2, ok := that.(GetSupportTenantAccessReq)
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
func (this *GetSupportTenantAccessResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSupportTenantAccessResp)
	if !ok {
		that2, ok := that.(GetSupportTenantAccessResp)
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
	if !this.AccessConfig.Equal(that1.AccessConfig) {
		return false
	}
	return true
}
func (this *UpdateSupportTenantAccessReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateSupportTenantAccessReq)
	if !ok {
		that2, ok := that.(UpdateSupportTenantAccessReq)
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
	if !this.AccessConfig.Equal(that1.AccessConfig) {
		return false
	}
	return true
}
func (this *UpdateSupportTenantAccessResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateSupportTenantAccessResp)
	if !ok {
		that2, ok := that.(UpdateSupportTenantAccessResp)
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
	if !this.AccessConfig.Equal(that1.AccessConfig) {
		return false
	}
	return true
}
func (this *GetSupportTenantAccessReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&allowed_tenant.GetSupportTenantAccessReq{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSupportTenantAccessResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&allowed_tenant.GetSupportTenantAccessResp{")
	if this.AccessConfig != nil {
		s = append(s, "AccessConfig: "+fmt.Sprintf("%#v", this.AccessConfig)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UpdateSupportTenantAccessReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&allowed_tenant.UpdateSupportTenantAccessReq{")
	if this.AccessConfig != nil {
		s = append(s, "AccessConfig: "+fmt.Sprintf("%#v", this.AccessConfig)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UpdateSupportTenantAccessResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&allowed_tenant.UpdateSupportTenantAccessResp{")
	if this.AccessConfig != nil {
		s = append(s, "AccessConfig: "+fmt.Sprintf("%#v", this.AccessConfig)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicCustomapiUam(v interface{}, typ string) string {
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

// CustomAPIClient is the client API for CustomAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomAPIClient interface {
	// UpdateSupportTenantAccess
	//
	// x-displayName: "Update Support Tenant Access"
	// This RPC can be used to manage user access for all flavors of support tenants currently
	// supported by the platform. Use read-only, read-write with specific namespaces or
	// admin can specify custom groups to control access by the support tenant user.
	// Name is well-known identifier for a specific support related tenant.
	UpdateSupportTenantAccess(ctx context.Context, in *UpdateSupportTenantAccessReq, opts ...grpc.CallOption) (*UpdateSupportTenantAccessResp, error)
	// GetSupportTenantAccess
	//
	// x-displayName: "Get Support Tenant Access"
	// Get current access details for the support tenant.
	GetSupportTenantAccess(ctx context.Context, in *GetSupportTenantAccessReq, opts ...grpc.CallOption) (*GetSupportTenantAccessResp, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) UpdateSupportTenantAccess(ctx context.Context, in *UpdateSupportTenantAccessReq, opts ...grpc.CallOption) (*UpdateSupportTenantAccessResp, error) {
	out := new(UpdateSupportTenantAccessResp)
	err := c.cc.Invoke(ctx, "/ves.io.schema.tenant_management.allowed_tenant.CustomAPI/UpdateSupportTenantAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customAPIClient) GetSupportTenantAccess(ctx context.Context, in *GetSupportTenantAccessReq, opts ...grpc.CallOption) (*GetSupportTenantAccessResp, error) {
	out := new(GetSupportTenantAccessResp)
	err := c.cc.Invoke(ctx, "/ves.io.schema.tenant_management.allowed_tenant.CustomAPI/GetSupportTenantAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// UpdateSupportTenantAccess
	//
	// x-displayName: "Update Support Tenant Access"
	// This RPC can be used to manage user access for all flavors of support tenants currently
	// supported by the platform. Use read-only, read-write with specific namespaces or
	// admin can specify custom groups to control access by the support tenant user.
	// Name is well-known identifier for a specific support related tenant.
	UpdateSupportTenantAccess(context.Context, *UpdateSupportTenantAccessReq) (*UpdateSupportTenantAccessResp, error)
	// GetSupportTenantAccess
	//
	// x-displayName: "Get Support Tenant Access"
	// Get current access details for the support tenant.
	GetSupportTenantAccess(context.Context, *GetSupportTenantAccessReq) (*GetSupportTenantAccessResp, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) UpdateSupportTenantAccess(ctx context.Context, req *UpdateSupportTenantAccessReq) (*UpdateSupportTenantAccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSupportTenantAccess not implemented")
}
func (*UnimplementedCustomAPIServer) GetSupportTenantAccess(ctx context.Context, req *GetSupportTenantAccessReq) (*GetSupportTenantAccessResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportTenantAccess not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_UpdateSupportTenantAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSupportTenantAccessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).UpdateSupportTenantAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.tenant_management.allowed_tenant.CustomAPI/UpdateSupportTenantAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).UpdateSupportTenantAccess(ctx, req.(*UpdateSupportTenantAccessReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomAPI_GetSupportTenantAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupportTenantAccessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).GetSupportTenantAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.tenant_management.allowed_tenant.CustomAPI/GetSupportTenantAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).GetSupportTenantAccess(ctx, req.(*GetSupportTenantAccessReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.tenant_management.allowed_tenant.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateSupportTenantAccess",
			Handler:    _CustomAPI_UpdateSupportTenantAccess_Handler,
		},
		{
			MethodName: "GetSupportTenantAccess",
			Handler:    _CustomAPI_GetSupportTenantAccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/tenant_management/allowed_tenant/public_customapi_uam.proto",
}

func (m *GetSupportTenantAccessReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSupportTenantAccessReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSupportTenantAccessReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetSupportTenantAccessResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSupportTenantAccessResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSupportTenantAccessResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AccessConfig != nil {
		{
			size, err := m.AccessConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPublicCustomapiUam(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateSupportTenantAccessReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateSupportTenantAccessReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateSupportTenantAccessReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AccessConfig != nil {
		{
			size, err := m.AccessConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPublicCustomapiUam(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *UpdateSupportTenantAccessResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateSupportTenantAccessResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateSupportTenantAccessResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AccessConfig != nil {
		{
			size, err := m.AccessConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPublicCustomapiUam(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPublicCustomapiUam(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublicCustomapiUam(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetSupportTenantAccessReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetSupportTenantAccessResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccessConfig != nil {
		l = m.AccessConfig.Size()
		n += 1 + l + sovPublicCustomapiUam(uint64(l))
	}
	return n
}

func (m *UpdateSupportTenantAccessReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccessConfig != nil {
		l = m.AccessConfig.Size()
		n += 1 + l + sovPublicCustomapiUam(uint64(l))
	}
	return n
}

func (m *UpdateSupportTenantAccessResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccessConfig != nil {
		l = m.AccessConfig.Size()
		n += 1 + l + sovPublicCustomapiUam(uint64(l))
	}
	return n
}

func sovPublicCustomapiUam(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapiUam(x uint64) (n int) {
	return sovPublicCustomapiUam(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetSupportTenantAccessReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSupportTenantAccessReq{`,
		`}`,
	}, "")
	return s
}
func (this *GetSupportTenantAccessResp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSupportTenantAccessResp{`,
		`AccessConfig:` + strings.Replace(fmt.Sprintf("%v", this.AccessConfig), "AllowedAccessConfig", "AllowedAccessConfig", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UpdateSupportTenantAccessReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpdateSupportTenantAccessReq{`,
		`AccessConfig:` + strings.Replace(fmt.Sprintf("%v", this.AccessConfig), "AllowedAccessConfig", "AllowedAccessConfig", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *UpdateSupportTenantAccessResp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UpdateSupportTenantAccessResp{`,
		`AccessConfig:` + strings.Replace(fmt.Sprintf("%v", this.AccessConfig), "AllowedAccessConfig", "AllowedAccessConfig", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicCustomapiUam(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetSupportTenantAccessReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapiUam
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
			return fmt.Errorf("proto: GetSupportTenantAccessReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSupportTenantAccessReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapiUam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapiUam
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapiUam
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
func (m *GetSupportTenantAccessResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapiUam
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
			return fmt.Errorf("proto: GetSupportTenantAccessResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSupportTenantAccessResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapiUam
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
				return ErrInvalidLengthPublicCustomapiUam
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapiUam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AccessConfig == nil {
				m.AccessConfig = &AllowedAccessConfig{}
			}
			if err := m.AccessConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapiUam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapiUam
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapiUam
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
func (m *UpdateSupportTenantAccessReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapiUam
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
			return fmt.Errorf("proto: UpdateSupportTenantAccessReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateSupportTenantAccessReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapiUam
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
				return ErrInvalidLengthPublicCustomapiUam
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapiUam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AccessConfig == nil {
				m.AccessConfig = &AllowedAccessConfig{}
			}
			if err := m.AccessConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapiUam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapiUam
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapiUam
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
func (m *UpdateSupportTenantAccessResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapiUam
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
			return fmt.Errorf("proto: UpdateSupportTenantAccessResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateSupportTenantAccessResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapiUam
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
				return ErrInvalidLengthPublicCustomapiUam
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapiUam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AccessConfig == nil {
				m.AccessConfig = &AllowedAccessConfig{}
			}
			if err := m.AccessConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapiUam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapiUam
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapiUam
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
func skipPublicCustomapiUam(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicCustomapiUam
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
					return 0, ErrIntOverflowPublicCustomapiUam
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
					return 0, ErrIntOverflowPublicCustomapiUam
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
				return 0, ErrInvalidLengthPublicCustomapiUam
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublicCustomapiUam
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublicCustomapiUam
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublicCustomapiUam        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomapiUam          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublicCustomapiUam = fmt.Errorf("proto: unexpected end of group")
)
