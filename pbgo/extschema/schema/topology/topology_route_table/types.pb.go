// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/topology/topology_route_table/types.proto

package topology_route_table

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	topology "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/topology"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Topology Route Table
//
// x-displayName: "Route Table"
//
//	A node represents route table in the topology graph.
type GlobalSpecType struct {
	// Topology Metadata
	//
	// x-displayName: "Metadata"
	// A common metadata for topology.
	TopologyMetadata *topology.MetaType `protobuf:"bytes,1,opt,name=topology_metadata,json=topologyMetadata,proto3" json:"topology_metadata,omitempty"`
	// Topology Spec
	//
	// x-displayName: "Spec"
	// A canonical spec for this topology node.
	TopologySpec *topology.RouteTableType `protobuf:"bytes,2,opt,name=topology_spec,json=topologySpec,proto3" json:"topology_spec,omitempty"`
}

func (m *GlobalSpecType) Reset()      { *m = GlobalSpecType{} }
func (*GlobalSpecType) ProtoMessage() {}
func (*GlobalSpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_200c625eefa92c57, []int{0}
}
func (m *GlobalSpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GlobalSpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *GlobalSpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalSpecType.Merge(m, src)
}
func (m *GlobalSpecType) XXX_Size() int {
	return m.Size()
}
func (m *GlobalSpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalSpecType.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalSpecType proto.InternalMessageInfo

func (m *GlobalSpecType) GetTopologyMetadata() *topology.MetaType {
	if m != nil {
		return m.TopologyMetadata
	}
	return nil
}

func (m *GlobalSpecType) GetTopologySpec() *topology.RouteTableType {
	if m != nil {
		return m.TopologySpec
	}
	return nil
}

func init() {
	proto.RegisterType((*GlobalSpecType)(nil), "ves.io.schema.topology.topology_route_table.GlobalSpecType")
}

func init() {
	proto.RegisterFile("ves.io/schema/topology/topology_route_table/types.proto", fileDescriptor_200c625eefa92c57)
}

var fileDescriptor_200c625eefa92c57 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4e, 0x02, 0x41,
	0x10, 0xc6, 0x6f, 0x2d, 0x2c, 0xce, 0x3f, 0x51, 0x2a, 0xa4, 0x98, 0x10, 0x0a, 0x63, 0x62, 0xd8,
	0x4d, 0xb4, 0xb0, 0xb7, 0xb1, 0x30, 0x34, 0x48, 0x65, 0x43, 0xf6, 0x8e, 0x61, 0x21, 0xde, 0x65,
	0x36, 0x7b, 0x7b, 0x44, 0x3a, 0x1f, 0xc1, 0xde, 0x17, 0xb0, 0xf0, 0x41, 0x2c, 0x29, 0x29, 0x65,
	0x69, 0x2c, 0x79, 0x04, 0x73, 0x0b, 0x07, 0x89, 0xf1, 0x12, 0xbb, 0x99, 0x9d, 0xf9, 0x7e, 0xfb,
	0xed, 0xb7, 0xe1, 0xcd, 0x04, 0x33, 0x3e, 0x26, 0x91, 0xc5, 0x23, 0x4c, 0xa5, 0xb0, 0xa4, 0x29,
	0x21, 0x35, 0xdd, 0x16, 0x7d, 0x43, 0xb9, 0xc5, 0xbe, 0x95, 0x51, 0x82, 0xc2, 0x4e, 0x35, 0x66,
	0x5c, 0x1b, 0xb2, 0x54, 0xbb, 0x5c, 0x0b, 0xf9, 0x5a, 0xc8, 0xcb, 0x7d, 0xfe, 0x97, 0xb0, 0xd1,
	0x56, 0x63, 0x3b, 0xca, 0x23, 0x1e, 0x53, 0x2a, 0x14, 0x29, 0x12, 0x9e, 0x11, 0xe5, 0x43, 0xdf,
	0xf9, 0xc6, 0x57, 0x6b, 0x76, 0xa3, 0x55, 0x65, 0x6a, 0x77, 0x7f, 0xe3, 0xec, 0xd7, 0xce, 0x6e,
	0xd4, 0xfa, 0x60, 0xe1, 0xf1, 0x5d, 0x42, 0x91, 0x4c, 0x1e, 0x34, 0xc6, 0xbd, 0xa9, 0xc6, 0x5a,
	0x27, 0x3c, 0xdd, 0x1a, 0x4b, 0xd1, 0xca, 0x81, 0xb4, 0xb2, 0xce, 0x9a, 0xec, 0xe2, 0xe0, 0xaa,
	0xc9, 0x2b, 0x5e, 0xd2, 0x41, 0x2b, 0x0b, 0x71, 0xf7, 0xa4, 0x3c, 0xea, 0x6c, 0x94, 0xb5, 0xfb,
	0xf0, 0x68, 0x8b, 0xcb, 0x34, 0xc6, 0xf5, 0x3d, 0x8f, 0x3a, 0xaf, 0x42, 0x75, 0x8b, 0x2c, 0x7a,
	0x45, 0x14, 0x1e, 0x78, 0x58, 0x0e, 0x0a, 0x7f, 0xb7, 0x6f, 0x6c, 0xb6, 0x80, 0x60, 0xbe, 0x80,
	0x60, 0xb5, 0x00, 0xf6, 0xe2, 0x80, 0xbd, 0x3b, 0x60, 0x9f, 0x0e, 0xd8, 0xcc, 0x01, 0x9b, 0x3b,
	0x60, 0x5f, 0x0e, 0xd8, 0xb7, 0x83, 0x60, 0xe5, 0x80, 0xbd, 0x2e, 0x21, 0x98, 0x2d, 0x21, 0x98,
	0x2f, 0x21, 0x78, 0x94, 0x8a, 0xf4, 0x93, 0xe2, 0x13, 0x4a, 0x2c, 0x1a, 0x23, 0x79, 0x9e, 0x09,
	0x5f, 0x0c, 0xc9, 0xa4, 0x6d, 0x6d, 0x68, 0x32, 0x1e, 0xa0, 0x69, 0x97, 0x63, 0xa1, 0x23, 0x45,
	0x02, 0x9f, 0xed, 0x26, 0xb3, 0xff, 0xfc, 0x79, 0xb4, 0xef, 0x33, 0xbd, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xc6, 0x39, 0x36, 0x70, 0x29, 0x02, 0x00, 0x00,
}

func (this *GlobalSpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GlobalSpecType)
	if !ok {
		that2, ok := that.(GlobalSpecType)
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
	if !this.TopologyMetadata.Equal(that1.TopologyMetadata) {
		return false
	}
	if !this.TopologySpec.Equal(that1.TopologySpec) {
		return false
	}
	return true
}
func (this *GlobalSpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&topology_route_table.GlobalSpecType{")
	if this.TopologyMetadata != nil {
		s = append(s, "TopologyMetadata: "+fmt.Sprintf("%#v", this.TopologyMetadata)+",\n")
	}
	if this.TopologySpec != nil {
		s = append(s, "TopologySpec: "+fmt.Sprintf("%#v", this.TopologySpec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GlobalSpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalSpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalSpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TopologySpec != nil {
		{
			size, err := m.TopologySpec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.TopologyMetadata != nil {
		{
			size, err := m.TopologyMetadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GlobalSpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TopologyMetadata != nil {
		l = m.TopologyMetadata.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.TopologySpec != nil {
		l = m.TopologySpec.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GlobalSpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GlobalSpecType{`,
		`TopologyMetadata:` + strings.Replace(fmt.Sprintf("%v", this.TopologyMetadata), "MetaType", "topology.MetaType", 1) + `,`,
		`TopologySpec:` + strings.Replace(fmt.Sprintf("%v", this.TopologySpec), "RouteTableType", "topology.RouteTableType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GlobalSpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GlobalSpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GlobalSpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopologyMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TopologyMetadata == nil {
				m.TopologyMetadata = &topology.MetaType{}
			}
			if err := m.TopologyMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopologySpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TopologySpec == nil {
				m.TopologySpec = &topology.RouteTableType{}
			}
			if err := m.TopologySpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)