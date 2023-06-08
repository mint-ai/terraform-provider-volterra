// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/waf_signatures_changelog/public_waf_signature_changelog_customapi.proto

// Waf signature related Custom APIs
//
// x-displayName: "Waf Signatures Changelog"
// Waf Signatures Changelog custom APIs

package waf_signatures_changelog

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

// ReleasedSignaturesReq
//
// x-displayName: "Released Signatures Request"
// Request to get the list of released signatures
type ReleasedSignaturesReq struct {
	// namespace
	//
	// x-displayName: "Namespace"
	// x-example: "shared"
	// Fetch waf signatures changelog for the given namespace
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// vh_name
	//
	// x-displayName: "Virtual Host Name"
	// x-example: "blogging-app"
	// Virtual Host for current request
	VhName string `protobuf:"bytes,2,opt,name=vh_name,json=vhName,proto3" json:"vh_name,omitempty"`
}

func (m *ReleasedSignaturesReq) Reset()      { *m = ReleasedSignaturesReq{} }
func (*ReleasedSignaturesReq) ProtoMessage() {}
func (*ReleasedSignaturesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_af91ff46163493fa, []int{0}
}
func (m *ReleasedSignaturesReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReleasedSignaturesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReleasedSignaturesReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReleasedSignaturesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleasedSignaturesReq.Merge(m, src)
}
func (m *ReleasedSignaturesReq) XXX_Size() int {
	return m.Size()
}
func (m *ReleasedSignaturesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleasedSignaturesReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReleasedSignaturesReq proto.InternalMessageInfo

func (m *ReleasedSignaturesReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ReleasedSignaturesReq) GetVhName() string {
	if m != nil {
		return m.VhName
	}
	return ""
}

// ReleasedSignaturesRes
//
// x-displayName: "Released Signatures Response"
// Response to get the list of released signatures
type ReleasedSignaturesRsp struct {
	// released_signatures
	//
	// x-displayName: "Released Signatures"
	// Released Signatures
	ReleasedSignatures []*GlobalSpecType `protobuf:"bytes,1,rep,name=released_signatures,json=releasedSignatures,proto3" json:"released_signatures,omitempty"`
}

func (m *ReleasedSignaturesRsp) Reset()      { *m = ReleasedSignaturesRsp{} }
func (*ReleasedSignaturesRsp) ProtoMessage() {}
func (*ReleasedSignaturesRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_af91ff46163493fa, []int{1}
}
func (m *ReleasedSignaturesRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReleasedSignaturesRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReleasedSignaturesRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReleasedSignaturesRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleasedSignaturesRsp.Merge(m, src)
}
func (m *ReleasedSignaturesRsp) XXX_Size() int {
	return m.Size()
}
func (m *ReleasedSignaturesRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleasedSignaturesRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ReleasedSignaturesRsp proto.InternalMessageInfo

func (m *ReleasedSignaturesRsp) GetReleasedSignatures() []*GlobalSpecType {
	if m != nil {
		return m.ReleasedSignatures
	}
	return nil
}

func init() {
	proto.RegisterType((*ReleasedSignaturesReq)(nil), "ves.io.schema.waf_signatures_changelog.ReleasedSignaturesReq")
	golang_proto.RegisterType((*ReleasedSignaturesReq)(nil), "ves.io.schema.waf_signatures_changelog.ReleasedSignaturesReq")
	proto.RegisterType((*ReleasedSignaturesRsp)(nil), "ves.io.schema.waf_signatures_changelog.ReleasedSignaturesRsp")
	golang_proto.RegisterType((*ReleasedSignaturesRsp)(nil), "ves.io.schema.waf_signatures_changelog.ReleasedSignaturesRsp")
}

func init() {
	proto.RegisterFile("ves.io/schema/waf_signatures_changelog/public_waf_signature_changelog_customapi.proto", fileDescriptor_af91ff46163493fa)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/waf_signatures_changelog/public_waf_signature_changelog_customapi.proto", fileDescriptor_af91ff46163493fa)
}

var fileDescriptor_af91ff46163493fa = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x6b, 0xd4, 0x4e,
	0x14, 0xc7, 0x33, 0x29, 0xf4, 0x47, 0xf3, 0xbb, 0x48, 0x4a, 0x31, 0xac, 0x65, 0x28, 0x7b, 0x90,
	0x5e, 0x92, 0x81, 0x15, 0xbc, 0x88, 0x07, 0xed, 0xa1, 0x20, 0x52, 0x61, 0xab, 0x08, 0x22, 0x84,
	0x49, 0xf6, 0xed, 0x64, 0x34, 0xc9, 0x8c, 0x33, 0x93, 0x68, 0x29, 0x85, 0x52, 0xfc, 0x03, 0x04,
	0xc1, 0xbf, 0xc1, 0x3f, 0x41, 0xe8, 0xa5, 0x37, 0x7b, 0x92, 0x45, 0x2f, 0x3d, 0xba, 0xd9, 0x1e,
	0x3c, 0xf6, 0x4f, 0x10, 0xb3, 0xe9, 0xb6, 0x6b, 0x57, 0x5c, 0xf0, 0xf6, 0xde, 0x7c, 0xde, 0xfb,
	0xe6, 0xe5, 0xbd, 0xaf, 0xf3, 0xa4, 0x04, 0x1d, 0x70, 0x41, 0x74, 0x9c, 0x40, 0x46, 0xc9, 0x6b,
	0xda, 0x0f, 0x35, 0x67, 0x39, 0x35, 0x85, 0x02, 0x1d, 0xc6, 0x09, 0xcd, 0x19, 0xa4, 0x82, 0x11,
	0x59, 0x44, 0x29, 0x8f, 0xc3, 0x29, 0x7e, 0x81, 0xc3, 0xb8, 0xd0, 0x46, 0x64, 0x54, 0xf2, 0x40,
	0x2a, 0x61, 0x84, 0x7b, 0x73, 0x2c, 0x1b, 0x8c, 0x65, 0x83, 0x3f, 0xc9, 0xb6, 0x7c, 0xc6, 0x4d,
	0x52, 0x44, 0x41, 0x2c, 0x32, 0xc2, 0x04, 0x13, 0xa4, 0x6e, 0x8f, 0x8a, 0x7e, 0x9d, 0xd5, 0x49,
	0x1d, 0x8d, 0x65, 0x5b, 0xab, 0x4c, 0x08, 0x96, 0x02, 0xa1, 0x92, 0x13, 0x9a, 0xe7, 0xc2, 0x50,
	0xc3, 0x45, 0xae, 0x1b, 0x7a, 0x63, 0xfa, 0x5f, 0x84, 0xbc, 0x0c, 0xdb, 0xd3, 0xb0, 0x04, 0x0d,
	0x79, 0xf9, 0x5b, 0x4d, 0x67, 0xce, 0x65, 0x98, 0x1d, 0x09, 0x4d, 0x4f, 0x7b, 0xcb, 0x59, 0xe9,
	0x42, 0x0a, 0x54, 0x43, 0x6f, 0x7b, 0x52, 0xdb, 0x85, 0x57, 0xee, 0xaa, 0xb3, 0x94, 0xd3, 0x0c,
	0xb4, 0xa4, 0x31, 0x78, 0x68, 0x0d, 0xad, 0x2f, 0x75, 0x2f, 0x1e, 0xdc, 0xeb, 0xce, 0x7f, 0x65,
	0x12, 0xfe, 0xca, 0x3d, 0xbb, 0x66, 0x8b, 0x65, 0xb2, 0x45, 0x33, 0x68, 0xef, 0xa3, 0x99, 0x82,
	0x5a, 0xba, 0xcc, 0x59, 0x56, 0x0d, 0xb8, 0x34, 0x96, 0x87, 0xd6, 0x16, 0xd6, 0xff, 0xef, 0xdc,
	0x0e, 0xe6, 0xdb, 0x78, 0xb0, 0x99, 0x8a, 0x88, 0xa6, 0xdb, 0x12, 0xe2, 0xc7, 0x3b, 0x12, 0xba,
	0xae, 0xba, 0xf2, 0xad, 0xce, 0xa9, 0xed, 0xe0, 0xa7, 0xb4, 0x3f, 0x79, 0xd9, 0x38, 0x6f, 0xde,
	0xa8, 0xaf, 0x7c, 0x4f, 0x72, 0xf7, 0xad, 0xed, 0xac, 0x6c, 0x82, 0xb9, 0x3a, 0xa8, 0x7b, 0x77,
	0xde, 0x41, 0x66, 0x6e, 0xad, 0xf5, 0x2f, 0xed, 0x5a, 0xb6, 0xb3, 0xe3, 0x4f, 0x36, 0xaa, 0x3e,
	0x7b, 0xcb, 0x25, 0x68, 0x9f, 0x0b, 0x9f, 0x41, 0x0e, 0x8a, 0xa6, 0xbe, 0x02, 0xda, 0x3b, 0xf8,
	0x76, 0xfa, 0xde, 0x7e, 0xe8, 0x3e, 0x68, 0xbc, 0x4c, 0x26, 0xc7, 0xd0, 0x64, 0x77, 0x12, 0xef,
	0x91, 0x92, 0x2b, 0x53, 0xd0, 0x34, 0x4c, 0x84, 0x36, 0x9a, 0xec, 0x36, 0x57, 0xda, 0x23, 0x33,
	0x76, 0xdf, 0xba, 0x73, 0x74, 0x88, 0x16, 0xbe, 0x1e, 0x22, 0x7f, 0xce, 0xa1, 0x1f, 0x45, 0x2f,
	0x20, 0x36, 0x07, 0x5f, 0x3c, 0xfb, 0x1a, 0xba, 0xff, 0x01, 0x0d, 0x86, 0xd8, 0x3a, 0x19, 0x62,
	0xeb, 0x6c, 0x88, 0xd1, 0x7e, 0x85, 0xd1, 0xc7, 0x0a, 0xa3, 0xe3, 0x0a, 0xa3, 0x41, 0x85, 0xd1,
	0xf7, 0x0a, 0xa3, 0x1f, 0x15, 0xb6, 0xce, 0x2a, 0x8c, 0xde, 0x8d, 0xb0, 0x75, 0x34, 0xc2, 0x68,
	0x30, 0xc2, 0xd6, 0xc9, 0x08, 0x5b, 0xcf, 0x9e, 0x33, 0x21, 0x5f, 0xb2, 0xa0, 0x14, 0xa9, 0x01,
	0xa5, 0x68, 0x50, 0x68, 0x52, 0x07, 0x7d, 0xa1, 0x32, 0x5f, 0x2a, 0x51, 0xf2, 0x1e, 0x28, 0xff,
	0x1c, 0x13, 0x19, 0x31, 0x41, 0xe0, 0x8d, 0x69, 0x7c, 0xfd, 0x17, 0x7b, 0x47, 0x8b, 0xb5, 0xb3,
	0x6f, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x59, 0x8c, 0x0d, 0x1c, 0x04, 0x00, 0x00,
}

func (this *ReleasedSignaturesReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReleasedSignaturesReq)
	if !ok {
		that2, ok := that.(ReleasedSignaturesReq)
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
	if this.VhName != that1.VhName {
		return false
	}
	return true
}
func (this *ReleasedSignaturesRsp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ReleasedSignaturesRsp)
	if !ok {
		that2, ok := that.(ReleasedSignaturesRsp)
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
	if len(this.ReleasedSignatures) != len(that1.ReleasedSignatures) {
		return false
	}
	for i := range this.ReleasedSignatures {
		if !this.ReleasedSignatures[i].Equal(that1.ReleasedSignatures[i]) {
			return false
		}
	}
	return true
}
func (this *ReleasedSignaturesReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&waf_signatures_changelog.ReleasedSignaturesReq{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "VhName: "+fmt.Sprintf("%#v", this.VhName)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ReleasedSignaturesRsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&waf_signatures_changelog.ReleasedSignaturesRsp{")
	if this.ReleasedSignatures != nil {
		s = append(s, "ReleasedSignatures: "+fmt.Sprintf("%#v", this.ReleasedSignatures)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicWafSignatureChangelogCustomapi(v interface{}, typ string) string {
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

// WafSignatureChangelogCustomApiClient is the client API for WafSignatureChangelogCustomApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WafSignatureChangelogCustomApiClient interface {
	// ReleasedSignatures
	//
	// x-displayName: "Released Signatures"
	// API to get Released Signatures
	GetReleasedSignatures(ctx context.Context, in *ReleasedSignaturesReq, opts ...grpc.CallOption) (*ReleasedSignaturesRsp, error)
}

type wafSignatureChangelogCustomApiClient struct {
	cc *grpc.ClientConn
}

func NewWafSignatureChangelogCustomApiClient(cc *grpc.ClientConn) WafSignatureChangelogCustomApiClient {
	return &wafSignatureChangelogCustomApiClient{cc}
}

func (c *wafSignatureChangelogCustomApiClient) GetReleasedSignatures(ctx context.Context, in *ReleasedSignaturesReq, opts ...grpc.CallOption) (*ReleasedSignaturesRsp, error) {
	out := new(ReleasedSignaturesRsp)
	err := c.cc.Invoke(ctx, "/ves.io.schema.waf_signatures_changelog.WafSignatureChangelogCustomApi/GetReleasedSignatures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WafSignatureChangelogCustomApiServer is the server API for WafSignatureChangelogCustomApi service.
type WafSignatureChangelogCustomApiServer interface {
	// ReleasedSignatures
	//
	// x-displayName: "Released Signatures"
	// API to get Released Signatures
	GetReleasedSignatures(context.Context, *ReleasedSignaturesReq) (*ReleasedSignaturesRsp, error)
}

// UnimplementedWafSignatureChangelogCustomApiServer can be embedded to have forward compatible implementations.
type UnimplementedWafSignatureChangelogCustomApiServer struct {
}

func (*UnimplementedWafSignatureChangelogCustomApiServer) GetReleasedSignatures(ctx context.Context, req *ReleasedSignaturesReq) (*ReleasedSignaturesRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReleasedSignatures not implemented")
}

func RegisterWafSignatureChangelogCustomApiServer(s *grpc.Server, srv WafSignatureChangelogCustomApiServer) {
	s.RegisterService(&_WafSignatureChangelogCustomApi_serviceDesc, srv)
}

func _WafSignatureChangelogCustomApi_GetReleasedSignatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleasedSignaturesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WafSignatureChangelogCustomApiServer).GetReleasedSignatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.waf_signatures_changelog.WafSignatureChangelogCustomApi/GetReleasedSignatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WafSignatureChangelogCustomApiServer).GetReleasedSignatures(ctx, req.(*ReleasedSignaturesReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _WafSignatureChangelogCustomApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.waf_signatures_changelog.WafSignatureChangelogCustomApi",
	HandlerType: (*WafSignatureChangelogCustomApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReleasedSignatures",
			Handler:    _WafSignatureChangelogCustomApi_GetReleasedSignatures_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/waf_signatures_changelog/public_waf_signature_changelog_customapi.proto",
}

func (m *ReleasedSignaturesReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReleasedSignaturesReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReleasedSignaturesReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VhName) > 0 {
		i -= len(m.VhName)
		copy(dAtA[i:], m.VhName)
		i = encodeVarintPublicWafSignatureChangelogCustomapi(dAtA, i, uint64(len(m.VhName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicWafSignatureChangelogCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReleasedSignaturesRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReleasedSignaturesRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReleasedSignaturesRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReleasedSignatures) > 0 {
		for iNdEx := len(m.ReleasedSignatures) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReleasedSignatures[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPublicWafSignatureChangelogCustomapi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintPublicWafSignatureChangelogCustomapi(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublicWafSignatureChangelogCustomapi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReleasedSignaturesReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicWafSignatureChangelogCustomapi(uint64(l))
	}
	l = len(m.VhName)
	if l > 0 {
		n += 1 + l + sovPublicWafSignatureChangelogCustomapi(uint64(l))
	}
	return n
}

func (m *ReleasedSignaturesRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ReleasedSignatures) > 0 {
		for _, e := range m.ReleasedSignatures {
			l = e.Size()
			n += 1 + l + sovPublicWafSignatureChangelogCustomapi(uint64(l))
		}
	}
	return n
}

func sovPublicWafSignatureChangelogCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicWafSignatureChangelogCustomapi(x uint64) (n int) {
	return sovPublicWafSignatureChangelogCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ReleasedSignaturesReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReleasedSignaturesReq{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`VhName:` + fmt.Sprintf("%v", this.VhName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReleasedSignaturesRsp) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForReleasedSignatures := "[]*GlobalSpecType{"
	for _, f := range this.ReleasedSignatures {
		repeatedStringForReleasedSignatures += strings.Replace(fmt.Sprintf("%v", f), "GlobalSpecType", "GlobalSpecType", 1) + ","
	}
	repeatedStringForReleasedSignatures += "}"
	s := strings.Join([]string{`&ReleasedSignaturesRsp{`,
		`ReleasedSignatures:` + repeatedStringForReleasedSignatures + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicWafSignatureChangelogCustomapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ReleasedSignaturesReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicWafSignatureChangelogCustomapi
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
			return fmt.Errorf("proto: ReleasedSignaturesReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReleasedSignaturesReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicWafSignatureChangelogCustomapi
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
				return ErrInvalidLengthPublicWafSignatureChangelogCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicWafSignatureChangelogCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VhName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicWafSignatureChangelogCustomapi
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
				return ErrInvalidLengthPublicWafSignatureChangelogCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicWafSignatureChangelogCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VhName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicWafSignatureChangelogCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicWafSignatureChangelogCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicWafSignatureChangelogCustomapi
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
func (m *ReleasedSignaturesRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicWafSignatureChangelogCustomapi
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
			return fmt.Errorf("proto: ReleasedSignaturesRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReleasedSignaturesRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReleasedSignatures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicWafSignatureChangelogCustomapi
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
				return ErrInvalidLengthPublicWafSignatureChangelogCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublicWafSignatureChangelogCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReleasedSignatures = append(m.ReleasedSignatures, &GlobalSpecType{})
			if err := m.ReleasedSignatures[len(m.ReleasedSignatures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicWafSignatureChangelogCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicWafSignatureChangelogCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicWafSignatureChangelogCustomapi
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
func skipPublicWafSignatureChangelogCustomapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicWafSignatureChangelogCustomapi
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
					return 0, ErrIntOverflowPublicWafSignatureChangelogCustomapi
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
					return 0, ErrIntOverflowPublicWafSignatureChangelogCustomapi
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
				return 0, ErrInvalidLengthPublicWafSignatureChangelogCustomapi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublicWafSignatureChangelogCustomapi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublicWafSignatureChangelogCustomapi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublicWafSignatureChangelogCustomapi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicWafSignatureChangelogCustomapi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublicWafSignatureChangelogCustomapi = fmt.Errorf("proto: unexpected end of group")
)
