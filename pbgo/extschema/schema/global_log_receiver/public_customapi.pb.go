// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/global_log_receiver/public_customapi.proto

// Global Log Receiver
//
// x-displayName: "Global Log Receiver"
// Custom APIs for Global log receiver object

package global_log_receiver

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

// Test Global Log Receiver Request
//
// x-displayName: "Test Global Log Receiver Request"
// Request to send test log
type TestGlobalLogReceiverRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-required
	// x-example: "system"
	// Namespace in which the Global log receiver is configured
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name
	//
	// x-displayName: "Name"
	// x-example: "slack1"
	// Name of the Global log receiver
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *TestGlobalLogReceiverRequest) Reset()      { *m = TestGlobalLogReceiverRequest{} }
func (*TestGlobalLogReceiverRequest) ProtoMessage() {}
func (*TestGlobalLogReceiverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c4bdaebe7e79181, []int{0}
}
func (m *TestGlobalLogReceiverRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestGlobalLogReceiverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestGlobalLogReceiverRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestGlobalLogReceiverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestGlobalLogReceiverRequest.Merge(m, src)
}
func (m *TestGlobalLogReceiverRequest) XXX_Size() int {
	return m.Size()
}
func (m *TestGlobalLogReceiverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestGlobalLogReceiverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestGlobalLogReceiverRequest proto.InternalMessageInfo

func (m *TestGlobalLogReceiverRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *TestGlobalLogReceiverRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Test Global Log Receiver Response
//
// x-displayName: "Test Global Log Receiver Response"
// Response for the Global Log Receiver test request; empty because the only returned
// information is error message.
type TestGlobalLogReceiverResponse struct {
}

func (m *TestGlobalLogReceiverResponse) Reset()      { *m = TestGlobalLogReceiverResponse{} }
func (*TestGlobalLogReceiverResponse) ProtoMessage() {}
func (*TestGlobalLogReceiverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c4bdaebe7e79181, []int{1}
}
func (m *TestGlobalLogReceiverResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestGlobalLogReceiverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestGlobalLogReceiverResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestGlobalLogReceiverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestGlobalLogReceiverResponse.Merge(m, src)
}
func (m *TestGlobalLogReceiverResponse) XXX_Size() int {
	return m.Size()
}
func (m *TestGlobalLogReceiverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestGlobalLogReceiverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestGlobalLogReceiverResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TestGlobalLogReceiverRequest)(nil), "ves.io.schema.global_log_receiver.TestGlobalLogReceiverRequest")
	golang_proto.RegisterType((*TestGlobalLogReceiverRequest)(nil), "ves.io.schema.global_log_receiver.TestGlobalLogReceiverRequest")
	proto.RegisterType((*TestGlobalLogReceiverResponse)(nil), "ves.io.schema.global_log_receiver.TestGlobalLogReceiverResponse")
	golang_proto.RegisterType((*TestGlobalLogReceiverResponse)(nil), "ves.io.schema.global_log_receiver.TestGlobalLogReceiverResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/global_log_receiver/public_customapi.proto", fileDescriptor_0c4bdaebe7e79181)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/global_log_receiver/public_customapi.proto", fileDescriptor_0c4bdaebe7e79181)
}

var fileDescriptor_0c4bdaebe7e79181 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6a, 0x14, 0x4d,
	0x14, 0xc5, 0xbb, 0x26, 0x1f, 0x1f, 0x4c, 0x2f, 0x1b, 0x85, 0x71, 0x1c, 0x4b, 0xed, 0x55, 0x10,
	0xba, 0x0a, 0x74, 0xa1, 0xb8, 0x89, 0xff, 0x40, 0x04, 0xc1, 0x30, 0x08, 0x82, 0x9b, 0xa1, 0xba,
	0x73, 0x53, 0x29, 0xed, 0xee, 0x5b, 0x56, 0x55, 0xb7, 0x8a, 0x04, 0x24, 0x4f, 0xa0, 0xf8, 0x12,
	0xbe, 0x43, 0x36, 0xd9, 0x88, 0xae, 0x64, 0xd0, 0x4d, 0x96, 0x4e, 0x8f, 0x0b, 0x97, 0xf3, 0x04,
	0x22, 0xd6, 0x74, 0xa2, 0x23, 0x13, 0x03, 0xee, 0xee, 0xed, 0x5f, 0xdd, 0xc3, 0xa9, 0xdb, 0xa7,
	0xc2, 0x2b, 0x35, 0x58, 0xa6, 0x90, 0xdb, 0x6c, 0x0b, 0x0a, 0xc1, 0x65, 0x8e, 0xa9, 0xc8, 0x47,
	0x39, 0xca, 0x91, 0x81, 0x0c, 0x54, 0x0d, 0x86, 0xeb, 0x2a, 0xcd, 0x55, 0x36, 0xca, 0x2a, 0xeb,
	0xb0, 0x10, 0x5a, 0x31, 0x6d, 0xd0, 0x61, 0x74, 0x7e, 0x3e, 0xc9, 0xe6, 0x93, 0x6c, 0xc9, 0x64,
	0x3f, 0x91, 0xca, 0x6d, 0x55, 0x29, 0xcb, 0xb0, 0xe0, 0x12, 0x25, 0x72, 0x3f, 0x99, 0x56, 0x9b,
	0xbe, 0xf3, 0x8d, 0xaf, 0xe6, 0x8a, 0xfd, 0x81, 0x44, 0x94, 0x39, 0x70, 0xa1, 0x15, 0x17, 0x65,
	0x89, 0x4e, 0x38, 0x85, 0xa5, 0x6d, 0xe9, 0xe9, 0x45, 0xa7, 0xa8, 0x7f, 0x87, 0xa7, 0x16, 0xa1,
	0x7b, 0xae, 0xe1, 0x00, 0x0d, 0x16, 0x51, 0x2d, 0x72, 0xb5, 0x21, 0x1c, 0xb4, 0x34, 0xfe, 0x83,
	0x82, 0x85, 0xb2, 0x5e, 0x14, 0x8f, 0xd7, 0xc3, 0xc1, 0x7d, 0xb0, 0xee, 0xb6, 0xbf, 0xe1, 0x5d,
	0x94, 0xc3, 0xf6, 0x7e, 0x43, 0x78, 0x52, 0x81, 0x75, 0xd1, 0x20, 0xec, 0x96, 0xa2, 0x00, 0xab,
	0x45, 0x06, 0x3d, 0x72, 0x8e, 0xac, 0x76, 0x87, 0xbf, 0x3e, 0x44, 0x51, 0xf8, 0xdf, 0xcf, 0xa6,
	0xd7, 0xf1, 0xc0, 0xd7, 0xf1, 0xd9, 0xf0, 0xcc, 0x11, 0x8a, 0x56, 0x63, 0x69, 0xe1, 0xe2, 0xbb,
	0x4e, 0xd8, 0xbd, 0xe9, 0x17, 0x7e, 0x7d, 0xfd, 0x4e, 0xf4, 0x9d, 0x84, 0x27, 0x97, 0x9e, 0x8f,
	0xd6, 0xd8, 0xb1, 0x7f, 0x81, 0xfd, 0xcd, 0x7b, 0xff, 0xda, 0xbf, 0x0b, 0xcc, 0xad, 0xc6, 0x69,
	0xf3, 0xbe, 0x77, 0xa2, 0x06, 0x9b, 0x28, 0x4c, 0x24, 0x94, 0x60, 0x44, 0x9e, 0x3c, 0x35, 0xca,
	0xc1, 0xce, 0xe7, 0xaf, 0x6f, 0x3a, 0xb7, 0xe2, 0xb5, 0x36, 0x40, 0xfc, 0x70, 0x25, 0x96, 0xbf,
	0x38, 0xac, 0xb7, 0x97, 0x45, 0xae, 0x3d, 0xb0, 0xcd, 0x1d, 0x58, 0x77, 0x95, 0x5c, 0xe8, 0x5f,
	0xde, 0xdb, 0x25, 0x2b, 0x9f, 0x76, 0xc9, 0xea, 0xf1, 0x66, 0xef, 0xa5, 0x8f, 0x20, 0x73, 0x3b,
	0x1f, 0x7b, 0x2b, 0x33, 0x42, 0x6e, 0xbc, 0x26, 0xe3, 0x09, 0x0d, 0xf6, 0x27, 0x34, 0x98, 0x4d,
	0x28, 0x79, 0xd9, 0x50, 0xf2, 0xb6, 0xa1, 0xe4, 0x43, 0x43, 0xc9, 0xb8, 0xa1, 0xe4, 0x4b, 0x43,
	0xc9, 0xb7, 0x86, 0x06, 0xb3, 0x86, 0x92, 0x57, 0x53, 0x1a, 0xec, 0x4d, 0x29, 0x19, 0x4f, 0x69,
	0xb0, 0x3f, 0xa5, 0xc1, 0xc3, 0x07, 0x12, 0xf5, 0x63, 0xc9, 0x6a, 0xcc, 0x1d, 0x18, 0x23, 0x58,
	0x65, 0xb9, 0x2f, 0x36, 0xd1, 0x14, 0x89, 0x36, 0x58, 0xab, 0x0d, 0x30, 0xc9, 0x01, 0xe6, 0x3a,
	0x95, 0xc8, 0xe1, 0x99, 0x6b, 0x83, 0x74, 0xf4, 0x7b, 0x4a, 0xff, 0xf7, 0xa9, 0xba, 0xf4, 0x23,
	0x00, 0x00, 0xff, 0xff, 0xde, 0x46, 0x9e, 0x29, 0x7b, 0x03, 0x00, 0x00,
}

func (this *TestGlobalLogReceiverRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TestGlobalLogReceiverRequest)
	if !ok {
		that2, ok := that.(TestGlobalLogReceiverRequest)
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
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *TestGlobalLogReceiverResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TestGlobalLogReceiverResponse)
	if !ok {
		that2, ok := that.(TestGlobalLogReceiverResponse)
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
func (this *TestGlobalLogReceiverRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&global_log_receiver.TestGlobalLogReceiverRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TestGlobalLogReceiverResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&global_log_receiver.TestGlobalLogReceiverResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPublicCustomapi(v interface{}, typ string) string {
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
	// Test Global Log Receiver
	//
	// x-displayName: "Test Global Log Receiver"
	// API to send test log
	TestGlobalLogReceiver(ctx context.Context, in *TestGlobalLogReceiverRequest, opts ...grpc.CallOption) (*TestGlobalLogReceiverResponse, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) TestGlobalLogReceiver(ctx context.Context, in *TestGlobalLogReceiverRequest, opts ...grpc.CallOption) (*TestGlobalLogReceiverResponse, error) {
	out := new(TestGlobalLogReceiverResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.global_log_receiver.CustomAPI/TestGlobalLogReceiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// Test Global Log Receiver
	//
	// x-displayName: "Test Global Log Receiver"
	// API to send test log
	TestGlobalLogReceiver(context.Context, *TestGlobalLogReceiverRequest) (*TestGlobalLogReceiverResponse, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) TestGlobalLogReceiver(ctx context.Context, req *TestGlobalLogReceiverRequest) (*TestGlobalLogReceiverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestGlobalLogReceiver not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_TestGlobalLogReceiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestGlobalLogReceiverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).TestGlobalLogReceiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.global_log_receiver.CustomAPI/TestGlobalLogReceiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).TestGlobalLogReceiver(ctx, req.(*TestGlobalLogReceiverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.global_log_receiver.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestGlobalLogReceiver",
			Handler:    _CustomAPI_TestGlobalLogReceiver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/global_log_receiver/public_customapi.proto",
}

func (m *TestGlobalLogReceiverRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestGlobalLogReceiverRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestGlobalLogReceiverRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestGlobalLogReceiverResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestGlobalLogReceiverResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestGlobalLogReceiverResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintPublicCustomapi(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublicCustomapi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TestGlobalLogReceiverRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	return n
}

func (m *TestGlobalLogReceiverResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TestGlobalLogReceiverRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TestGlobalLogReceiverRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TestGlobalLogReceiverResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TestGlobalLogReceiverResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringPublicCustomapi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TestGlobalLogReceiverRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
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
			return fmt.Errorf("proto: TestGlobalLogReceiverRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestGlobalLogReceiverRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func (m *TestGlobalLogReceiverResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublicCustomapi
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
			return fmt.Errorf("proto: TestGlobalLogReceiverResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestGlobalLogReceiverResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPublicCustomapi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPublicCustomapi
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
func skipPublicCustomapi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublicCustomapi
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
					return 0, ErrIntOverflowPublicCustomapi
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
					return 0, ErrIntOverflowPublicCustomapi
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
				return 0, ErrInvalidLengthPublicCustomapi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublicCustomapi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublicCustomapi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublicCustomapi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublicCustomapi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublicCustomapi = fmt.Errorf("proto: unexpected end of group")
)
