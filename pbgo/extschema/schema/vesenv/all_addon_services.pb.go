// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/vesenv/all_addon_services.proto

package vesenv

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
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

// AddonServiceName
//
// x-displayName: "Addon Service Name"
// Full set of addon services pre-approved names to be used as unique identifier/key.
// Every addon service is required to associate to one of the key in this list.
// Eywa to use this during addon service config object creation for metdata.name validation.
type AddonServiceName int32

const (
	// volterra_base
	//
	// x-displayName: "Volterra Base"
	// Full set of current volterra console services including native and non-native which are enabled by default.
	// for supporting legacy services before addon service migration purpose only.
	AS_VOLTERRA_BASE AddonServiceName = 0
	// shape_device_id
	//
	// x-displayName: "Shape Device ID"
	AS_SHAPE_DEVICE_ID AddonServiceName = 51
	// shape_client_side_defense
	//
	// x-displayName: "Shape Client Side Defense"
	AS_SHAPE_CLIENT_SIDE_DEFENSE AddonServiceName = 52
	// shape_safe_analyst
	//
	// x-displayName: "Shape Safe Analyst"
	AS_SHAPE_SAFE_ANALYST AddonServiceName = 53
)

var AddonServiceName_name = map[int32]string{
	0:  "AS_VOLTERRA_BASE",
	51: "AS_SHAPE_DEVICE_ID",
	52: "AS_SHAPE_CLIENT_SIDE_DEFENSE",
	53: "AS_SHAPE_SAFE_ANALYST",
}

var AddonServiceName_value = map[string]int32{
	"AS_VOLTERRA_BASE":             0,
	"AS_SHAPE_DEVICE_ID":           51,
	"AS_SHAPE_CLIENT_SIDE_DEFENSE": 52,
	"AS_SHAPE_SAFE_ANALYST":        53,
}

func (AddonServiceName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24b1043c1548ccb0, []int{0}
}

// NavigationTileName
//
// x-displayName: "Navigation Tile Name"
// Full set of Navigation Tile Names pre-approved names to be used as unique identifier/key.
// Eywa to use this during navigation tile config object creation for metdata.name validation.
type NavigationTileName int32

const (
	// cloud_and_edge_sites
	//
	// x-displayName: "Cloud And Edge Sites"
	//
	NT_CLOUD_AND_EDGE_SITES NavigationTileName = 0
	// shape_device_id
	//
	// x-displayName: "Shape Device ID"
	NT_SHAPE_DEVICE_ID NavigationTileName = 21
	// shape_client_side_defense
	//
	// x-displayName: "Shape Client Side Defense"
	NT_SHAPE_CLIENT_SIDE_DEFENSE NavigationTileName = 22
	// shape_safe_analyst
	//
	// x-displayName: "Shape Safe Analyst"
	NT_SHAPE_SAFE_ANALYST NavigationTileName = 23
)

var NavigationTileName_name = map[int32]string{
	0:  "NT_CLOUD_AND_EDGE_SITES",
	21: "NT_SHAPE_DEVICE_ID",
	22: "NT_SHAPE_CLIENT_SIDE_DEFENSE",
	23: "NT_SHAPE_SAFE_ANALYST",
}

var NavigationTileName_value = map[string]int32{
	"NT_CLOUD_AND_EDGE_SITES":      0,
	"NT_SHAPE_DEVICE_ID":           21,
	"NT_SHAPE_CLIENT_SIDE_DEFENSE": 22,
	"NT_SHAPE_SAFE_ANALYST":        23,
}

func (NavigationTileName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24b1043c1548ccb0, []int{1}
}

func init() {
	proto.RegisterEnum("ves.io.schema.vesenv.AddonServiceName", AddonServiceName_name, AddonServiceName_value)
	proto.RegisterEnum("ves.io.schema.vesenv.NavigationTileName", NavigationTileName_name, NavigationTileName_value)
}

func init() {
	proto.RegisterFile("ves.io/schema/vesenv/all_addon_services.proto", fileDescriptor_24b1043c1548ccb0)
}

var fileDescriptor_24b1043c1548ccb0 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbd, 0x6a, 0xeb, 0x30,
	0x00, 0x85, 0xad, 0xe5, 0x0e, 0x9a, 0x8c, 0xc9, 0x0f, 0xf7, 0x07, 0x71, 0xe9, 0x18, 0x88, 0x3d,
	0xa4, 0x7d, 0x00, 0x25, 0x56, 0x5a, 0x17, 0xe3, 0x94, 0xc8, 0x0d, 0xb4, 0x8b, 0x70, 0x12, 0xd5,
	0x35, 0x75, 0x22, 0x63, 0xbb, 0xa2, 0x4b, 0x21, 0x53, 0xe7, 0x3e, 0x46, 0x1f, 0xa5, 0x63, 0xc6,
	0x8c, 0x8d, 0xb2, 0x74, 0xcc, 0x23, 0x14, 0xc7, 0xa9, 0x21, 0x25, 0x9b, 0x38, 0xfa, 0x74, 0xf4,
	0xc1, 0x81, 0x6d, 0xc9, 0x33, 0x33, 0x12, 0x56, 0x36, 0xb9, 0xe7, 0xb3, 0xc0, 0x92, 0x3c, 0xe3,
	0x73, 0x69, 0x05, 0x71, 0xcc, 0x82, 0xe9, 0x54, 0xcc, 0x59, 0xc6, 0x53, 0x19, 0x4d, 0x78, 0x66,
	0x26, 0xa9, 0xc8, 0x85, 0x51, 0x2b, 0x71, 0xb3, 0xc4, 0xcd, 0x12, 0xff, 0x73, 0x72, 0xb4, 0x44,
	0x24, 0x79, 0x24, 0xe6, 0xfb, 0x97, 0xad, 0x67, 0xa8, 0xe3, 0xa2, 0x91, 0x96, 0x85, 0x5e, 0x30,
	0xe3, 0x46, 0x0d, 0xea, 0x98, 0xb2, 0xd1, 0xc0, 0xf5, 0xc9, 0x70, 0x88, 0x59, 0x17, 0x53, 0xa2,
	0x6b, 0x46, 0x03, 0x1a, 0x98, 0x32, 0x7a, 0x81, 0xaf, 0x08, 0xb3, 0xc9, 0xc8, 0xe9, 0x11, 0xe6,
	0xd8, 0x7a, 0xc7, 0xf8, 0x0f, 0xff, 0x55, 0x79, 0xcf, 0x75, 0x88, 0xe7, 0x33, 0xea, 0xd8, 0x05,
	0xd3, 0x27, 0x1e, 0x25, 0xfa, 0xa9, 0xf1, 0x1b, 0xd6, 0x2b, 0x82, 0xe2, 0x3e, 0x61, 0xd8, 0xc3,
	0xee, 0x0d, 0xf5, 0xf5, 0xb3, 0xd6, 0x0b, 0x80, 0x86, 0x17, 0xc8, 0x28, 0x0c, 0x0a, 0x29, 0x3f,
	0x8a, 0x4b, 0x83, 0xbf, 0xb0, 0xe9, 0xf9, 0xac, 0xe7, 0x0e, 0xae, 0x6d, 0x86, 0x3d, 0x9b, 0x11,
	0xfb, 0x9c, 0x30, 0xea, 0xf8, 0x84, 0x96, 0x22, 0xc5, 0x1f, 0x3f, 0x44, 0xea, 0x85, 0x48, 0x95,
	0x1f, 0x13, 0x69, 0x14, 0x22, 0x15, 0x71, 0x20, 0xd2, 0xec, 0x2e, 0xc0, 0x72, 0x8d, 0xb4, 0xd5,
	0x1a, 0x69, 0xdb, 0x35, 0x02, 0x0b, 0x85, 0xc0, 0x9b, 0x42, 0xe0, 0x5d, 0x21, 0xb0, 0x54, 0x08,
	0x7c, 0x28, 0x04, 0x3e, 0x15, 0xd2, 0xb6, 0x0a, 0x81, 0xd7, 0x0d, 0xd2, 0x96, 0x1b, 0xa4, 0xad,
	0x36, 0x48, 0xbb, 0xbd, 0x0c, 0x45, 0xf2, 0x10, 0x9a, 0x52, 0xc4, 0x39, 0x4f, 0xd3, 0xc0, 0x7c,
	0xcc, 0xac, 0xdd, 0xe1, 0x4e, 0xa4, 0xb3, 0x76, 0x92, 0x0a, 0x19, 0x4d, 0x79, 0xda, 0xfe, 0xbe,
	0xb6, 0x92, 0x71, 0x28, 0x2c, 0xfe, 0x94, 0xef, 0x17, 0x39, 0x18, 0x66, 0xfc, 0x6b, 0xb7, 0x48,
	0xe7, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x26, 0xd5, 0x29, 0x98, 0xfc, 0x01, 0x00, 0x00,
}

func (x AddonServiceName) String() string {
	s, ok := AddonServiceName_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x NavigationTileName) String() string {
	s, ok := NavigationTileName_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}