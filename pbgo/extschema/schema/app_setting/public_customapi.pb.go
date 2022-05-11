// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/app_setting/public_customapi.proto

// app setting object
//
// x-displayName: "App Setting"
// App setting object
// app setting object refers to an app type

package app_setting

import (
	context "context"
	encoding_binary "encoding/binary"
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

// Suspicious User Status Request
//
// x-displayName: "Suspicious User Request"
// Request to fetch suspicious users
type SuspiciousUserStatusReq struct {
	// namespace
	//
	// x-displayName: "Suspicious User Status Request"
	// fetch suspicious users for a given namespace
	// x-example: "bloggin-app-namespace-1"
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// name
	//
	// x-displayName: "name"
	// x-example: "value"
	// fetch suspicious users based on a given app setting
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// query
	//
	// x-displayName: "Query"
	// x-example: "query={app_type="blogging_app"}"
	// query is used to specify the list of matchers
	// syntax for query := {[<matcher>]}
	// <matcher> := <field_name><operator>"<value>"
	// <field_name> := string
	//   One or more of these fields in the security event may be specified in the query.
	//     app_type - application type
	//     vh_name - name of the virtual host
	// <value> := string
	// <operator> := ["="|"!="]
	//   = : equal to
	//   != : not equal to
	// When more than one matcher is specified in the query, then security events matching ALL the matchers will be returned in the response.
	// Example: query={country="United States", city="California"} will return all security events originating from California, United States.
	//
	// Optional: If not specified, all the security events matching the given tenant and namespace will be returned in the response.
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// start time
	//
	// x-displayName: "Start Time"
	// x-example: "1570007981"
	// fetch suspicious users during timestamp >= start_time
	// format: unix_timestamp|rfc 3339
	//
	// Optional: If not specified, then the start_time will be evaluated to end_time-10m
	//           If end_time is not specified, then the start_time will be evaluated to <current time>-10m
	StartTime string `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// end time
	//
	// x-displayName: "End Time"
	// x-example: "1570007981"
	// fetch suspicious users during timestamp <= end_time
	// format: unix_timestamp|rfc 3339
	//
	// Optional: If not specified, then the end_time will be evaluated to start_time+10m
	//           If start_time is not specified, then the end_time will be evaluated to <current time>
	EndTime string `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// topn
	//
	// x-displayName: "TopN"
	// x-example: "None of int32 samples [0 1 10 42 100 1024 2048] satisfied rules map[ves.io.schema.rules.uint32.gte:1 ves.io.schema.rules.uint32.lte:100]"
	// x-example: 10
	// fetch top 10 suspicious users
	//
	// Number of top field values to be returned in the response.
	// Optional: If not specified, top 5 values will be returned in the response.
	Topn uint32 `protobuf:"varint,6,opt,name=topn,proto3" json:"topn,omitempty"`
}

func (m *SuspiciousUserStatusReq) Reset()      { *m = SuspiciousUserStatusReq{} }
func (*SuspiciousUserStatusReq) ProtoMessage() {}
func (*SuspiciousUserStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fc14533964f7f4a, []int{0}
}
func (m *SuspiciousUserStatusReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuspiciousUserStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuspiciousUserStatusReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuspiciousUserStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuspiciousUserStatusReq.Merge(m, src)
}
func (m *SuspiciousUserStatusReq) XXX_Size() int {
	return m.Size()
}
func (m *SuspiciousUserStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SuspiciousUserStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_SuspiciousUserStatusReq proto.InternalMessageInfo

func (m *SuspiciousUserStatusReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SuspiciousUserStatusReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SuspiciousUserStatusReq) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SuspiciousUserStatusReq) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *SuspiciousUserStatusReq) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *SuspiciousUserStatusReq) GetTopn() uint32 {
	if m != nil {
		return m.Topn
	}
	return 0
}

// Suspicious User Status Response
//
// x-displayName: "Suspicious User Status Response"
// Response message for SuspiciousUserStatusReq
type SuspiciousUserStatusRsp struct {
	// suspicious users
	//
	// x-displayName: "Suspicious Users"
	// List of suspicious users
	SuspiciousUsers []*SuspiciousUser `protobuf:"bytes,1,rep,name=suspicious_users,json=suspiciousUsers,proto3" json:"suspicious_users,omitempty"`
}

func (m *SuspiciousUserStatusRsp) Reset()      { *m = SuspiciousUserStatusRsp{} }
func (*SuspiciousUserStatusRsp) ProtoMessage() {}
func (*SuspiciousUserStatusRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fc14533964f7f4a, []int{1}
}
func (m *SuspiciousUserStatusRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuspiciousUserStatusRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuspiciousUserStatusRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuspiciousUserStatusRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuspiciousUserStatusRsp.Merge(m, src)
}
func (m *SuspiciousUserStatusRsp) XXX_Size() int {
	return m.Size()
}
func (m *SuspiciousUserStatusRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SuspiciousUserStatusRsp.DiscardUnknown(m)
}

var xxx_messageInfo_SuspiciousUserStatusRsp proto.InternalMessageInfo

func (m *SuspiciousUserStatusRsp) GetSuspiciousUsers() []*SuspiciousUser {
	if m != nil {
		return m.SuspiciousUsers
	}
	return nil
}

// Suspicious User Data
//
// x-displayName: "Suspicious User Data"
// Message containing suspicious user data
type SuspiciousUser struct {
	// user ID
	//
	// x-displayName: "UserID"
	// x-example: "value"
	// String representing the user (ex: Source IP)
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// logs
	//
	// x-displayName: "Logs"
	// x-example: "value"
	// list of security events that matched the query. Contains no more than 100 messages.
	Logs []string `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
	// suspicion score
	//
	// x-displayName: "Suspicion Score"
	// x-example: "0"
	// total number of security events that matched the query.
	SuspicionScore float64 `protobuf:"fixed64,3,opt,name=suspicion_score,json=suspicionScore,proto3" json:"suspicion_score,omitempty"`
}

func (m *SuspiciousUser) Reset()      { *m = SuspiciousUser{} }
func (*SuspiciousUser) ProtoMessage() {}
func (*SuspiciousUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fc14533964f7f4a, []int{2}
}
func (m *SuspiciousUser) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuspiciousUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuspiciousUser.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuspiciousUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuspiciousUser.Merge(m, src)
}
func (m *SuspiciousUser) XXX_Size() int {
	return m.Size()
}
func (m *SuspiciousUser) XXX_DiscardUnknown() {
	xxx_messageInfo_SuspiciousUser.DiscardUnknown(m)
}

var xxx_messageInfo_SuspiciousUser proto.InternalMessageInfo

func (m *SuspiciousUser) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SuspiciousUser) GetLogs() []string {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *SuspiciousUser) GetSuspicionScore() float64 {
	if m != nil {
		return m.SuspicionScore
	}
	return 0
}

func init() {
	proto.RegisterType((*SuspiciousUserStatusReq)(nil), "ves.io.schema.app_setting.SuspiciousUserStatusReq")
	golang_proto.RegisterType((*SuspiciousUserStatusReq)(nil), "ves.io.schema.app_setting.SuspiciousUserStatusReq")
	proto.RegisterType((*SuspiciousUserStatusRsp)(nil), "ves.io.schema.app_setting.SuspiciousUserStatusRsp")
	golang_proto.RegisterType((*SuspiciousUserStatusRsp)(nil), "ves.io.schema.app_setting.SuspiciousUserStatusRsp")
	proto.RegisterType((*SuspiciousUser)(nil), "ves.io.schema.app_setting.SuspiciousUser")
	golang_proto.RegisterType((*SuspiciousUser)(nil), "ves.io.schema.app_setting.SuspiciousUser")
}

func init() {
	proto.RegisterFile("ves.io/schema/app_setting/public_customapi.proto", fileDescriptor_5fc14533964f7f4a)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/app_setting/public_customapi.proto", fileDescriptor_5fc14533964f7f4a)
}

var fileDescriptor_5fc14533964f7f4a = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x4f, 0x13, 0x41,
	0x14, 0xc7, 0x3b, 0x6d, 0x29, 0x32, 0x44, 0x34, 0x13, 0x12, 0x96, 0x8a, 0x63, 0xd3, 0x8b, 0xd5,
	0xa4, 0x3b, 0x04, 0x3f, 0x81, 0x10, 0x0f, 0x9c, 0x30, 0x05, 0x2f, 0x5e, 0x9a, 0xe9, 0xee, 0xb0,
	0x8c, 0x76, 0xf7, 0x0d, 0x33, 0xb3, 0x0d, 0xc4, 0x90, 0x18, 0x6e, 0xde, 0x4c, 0xfc, 0x12, 0xc6,
	0x4f, 0x60, 0xc4, 0x03, 0x37, 0x39, 0x19, 0xa2, 0x17, 0x8e, 0xb2, 0xf5, 0xa0, 0x89, 0x07, 0x3e,
	0x82, 0xd9, 0x69, 0xad, 0x14, 0x25, 0xd1, 0xdb, 0xfb, 0xbf, 0xdf, 0xbe, 0x7f, 0x66, 0xde, 0xbc,
	0xb7, 0x78, 0xb1, 0x27, 0x8c, 0x2f, 0x81, 0x99, 0x60, 0x4b, 0xc4, 0x9c, 0x71, 0xa5, 0xda, 0x46,
	0x58, 0x2b, 0x93, 0x88, 0xa9, 0xb4, 0xd3, 0x95, 0x41, 0x3b, 0x48, 0x8d, 0x85, 0x98, 0x2b, 0xe9,
	0x2b, 0x0d, 0x16, 0xc8, 0xfc, 0xa0, 0xc2, 0x1f, 0x54, 0xf8, 0xe7, 0x2a, 0xaa, 0xcd, 0x48, 0xda,
	0xad, 0xb4, 0xe3, 0x07, 0x10, 0xb3, 0x08, 0x22, 0x60, 0xae, 0xa2, 0x93, 0x6e, 0x3a, 0xe5, 0x84,
	0x8b, 0x06, 0x4e, 0xd5, 0x85, 0x08, 0x20, 0xea, 0x0a, 0xc6, 0x95, 0x64, 0x3c, 0x49, 0xc0, 0x72,
	0x2b, 0x21, 0x31, 0x43, 0x7a, 0x63, 0xfc, 0x64, 0xa0, 0xce, 0xc3, 0x85, 0x71, 0xd8, 0xe3, 0x5d,
	0x19, 0x72, 0x2b, 0x86, 0xb4, 0x7e, 0x81, 0x0a, 0x23, 0x92, 0xde, 0xb8, 0x43, 0xfd, 0x3d, 0xc2,
	0x73, 0xeb, 0xa9, 0x51, 0x32, 0x90, 0x90, 0x9a, 0x47, 0x46, 0xe8, 0x75, 0xcb, 0x6d, 0x6a, 0x5a,
	0x62, 0x9b, 0x2c, 0xe0, 0xa9, 0x84, 0xc7, 0xc2, 0x28, 0x1e, 0x08, 0x0f, 0xd5, 0x50, 0x63, 0xaa,
	0xf5, 0x3b, 0x41, 0x08, 0x2e, 0xe7, 0xc2, 0x2b, 0x3a, 0xe0, 0x62, 0x32, 0x8b, 0x27, 0xb6, 0x53,
	0xa1, 0x77, 0xbd, 0x92, 0x4b, 0x0e, 0x04, 0xb9, 0x89, 0xb1, 0xb1, 0x5c, 0xdb, 0xb6, 0x95, 0xb1,
	0xf0, 0xca, 0x03, 0x23, 0x97, 0xd9, 0x90, 0xb1, 0x20, 0xf3, 0xf8, 0x8a, 0x48, 0xc2, 0x01, 0x9c,
	0x70, 0x70, 0x52, 0x24, 0xa1, 0x43, 0xb7, 0x70, 0xd9, 0x82, 0x4a, 0xbc, 0x4a, 0x0d, 0x35, 0xae,
	0x2e, 0x4f, 0xbf, 0xfb, 0x7e, 0x58, 0xaa, 0xdc, 0x2d, 0x7b, 0x61, 0x03, 0xb5, 0x1c, 0xa8, 0xc3,
	0x25, 0xa7, 0x37, 0x8a, 0x6c, 0xe0, 0xeb, 0x66, 0x84, 0xda, 0xa9, 0x11, 0xda, 0x78, 0xa8, 0x56,
	0x6a, 0x4c, 0x2f, 0xdd, 0xf1, 0x2f, 0x7d, 0x3b, 0x7f, 0xdc, 0xad, 0x75, 0xcd, 0x8c, 0x69, 0x53,
	0xdf, 0xc4, 0x33, 0xe3, 0x9f, 0x90, 0x39, 0x3c, 0x99, 0x9b, 0xb7, 0x65, 0x38, 0xec, 0x51, 0x25,
	0x97, 0xab, 0x61, 0xde, 0xa0, 0x2e, 0x44, 0xc6, 0x2b, 0xd6, 0x4a, 0x79, 0x83, 0xf2, 0x98, 0xdc,
	0xc6, 0x23, 0xc7, 0xa4, 0x6d, 0x02, 0xd0, 0xc2, 0xb5, 0x0a, 0xb5, 0x66, 0x46, 0xe9, 0xf5, 0x3c,
	0xbb, 0xf4, 0xa6, 0x88, 0xa7, 0x56, 0xdc, 0xc8, 0xdd, 0x7f, 0xb8, 0x4a, 0x7e, 0x20, 0x3c, 0xfb,
	0xb7, 0x7b, 0x92, 0xa5, 0x7f, 0xbe, 0xca, 0xe8, 0x59, 0xab, 0xff, 0x5d, 0x63, 0x54, 0xbd, 0x7b,
	0xf4, 0xb6, 0x88, 0xb2, 0x0f, 0x5e, 0x3e, 0xf5, 0x4d, 0x09, 0x4d, 0xa5, 0x61, 0x67, 0xb7, 0x19,
	0x43, 0x22, 0x2d, 0xe8, 0xa6, 0x16, 0x3c, 0xdc, 0xff, 0xfc, 0xf5, 0x55, 0xf1, 0x01, 0x59, 0x19,
	0xee, 0x0b, 0x1b, 0x0d, 0x8a, 0x61, 0xcf, 0x46, 0xf1, 0xde, 0xf9, 0xcd, 0x1a, 0x82, 0x3d, 0x76,
	0xf1, 0x99, 0xaa, 0xfe, 0xe1, 0x01, 0x2a, 0x7d, 0x3a, 0x40, 0xb5, 0xcb, 0x0f, 0xba, 0xd6, 0x79,
	0x22, 0x02, 0xbb, 0xff, 0xd1, 0x2b, 0x2e, 0xa2, 0xe5, 0x17, 0xe8, 0xf8, 0x94, 0x16, 0x4e, 0x4e,
	0x69, 0xe1, 0xec, 0x94, 0xa2, 0xe7, 0x19, 0x45, 0xaf, 0x33, 0x8a, 0x8e, 0x32, 0x8a, 0x8e, 0x33,
	0x8a, 0xbe, 0x64, 0x14, 0x7d, 0xcb, 0x68, 0xe1, 0x2c, 0xa3, 0xe8, 0x65, 0x9f, 0x16, 0x0e, 0xfb,
	0x14, 0x1d, 0xf7, 0x69, 0xe1, 0xa4, 0x4f, 0x0b, 0x8f, 0xd7, 0x22, 0x50, 0x4f, 0x23, 0xbf, 0x07,
	0x5d, 0x2b, 0xb4, 0xe6, 0x7e, 0x6a, 0x98, 0x0b, 0x36, 0x41, 0xc7, 0xf9, 0x2d, 0x7b, 0x32, 0x14,
	0xba, 0xf9, 0x0b, 0x33, 0xd5, 0x89, 0x80, 0x89, 0x1d, 0x3b, 0x5c, 0xa9, 0x3f, 0x7f, 0x17, 0x9d,
	0x8a, 0xdb, 0xab, 0x7b, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x8c, 0x86, 0xb1, 0x52, 0x04,
	0x00, 0x00,
}

func (this *SuspiciousUserStatusReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SuspiciousUserStatusReq)
	if !ok {
		that2, ok := that.(SuspiciousUserStatusReq)
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
	if this.Query != that1.Query {
		return false
	}
	if this.StartTime != that1.StartTime {
		return false
	}
	if this.EndTime != that1.EndTime {
		return false
	}
	if this.Topn != that1.Topn {
		return false
	}
	return true
}
func (this *SuspiciousUserStatusRsp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SuspiciousUserStatusRsp)
	if !ok {
		that2, ok := that.(SuspiciousUserStatusRsp)
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
	if len(this.SuspiciousUsers) != len(that1.SuspiciousUsers) {
		return false
	}
	for i := range this.SuspiciousUsers {
		if !this.SuspiciousUsers[i].Equal(that1.SuspiciousUsers[i]) {
			return false
		}
	}
	return true
}
func (this *SuspiciousUser) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SuspiciousUser)
	if !ok {
		that2, ok := that.(SuspiciousUser)
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
	if this.UserId != that1.UserId {
		return false
	}
	if len(this.Logs) != len(that1.Logs) {
		return false
	}
	for i := range this.Logs {
		if this.Logs[i] != that1.Logs[i] {
			return false
		}
	}
	if this.SuspicionScore != that1.SuspicionScore {
		return false
	}
	return true
}
func (this *SuspiciousUserStatusReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&app_setting.SuspiciousUserStatusReq{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Query: "+fmt.Sprintf("%#v", this.Query)+",\n")
	s = append(s, "StartTime: "+fmt.Sprintf("%#v", this.StartTime)+",\n")
	s = append(s, "EndTime: "+fmt.Sprintf("%#v", this.EndTime)+",\n")
	s = append(s, "Topn: "+fmt.Sprintf("%#v", this.Topn)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SuspiciousUserStatusRsp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&app_setting.SuspiciousUserStatusRsp{")
	if this.SuspiciousUsers != nil {
		s = append(s, "SuspiciousUsers: "+fmt.Sprintf("%#v", this.SuspiciousUsers)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SuspiciousUser) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&app_setting.SuspiciousUser{")
	s = append(s, "UserId: "+fmt.Sprintf("%#v", this.UserId)+",\n")
	s = append(s, "Logs: "+fmt.Sprintf("%#v", this.Logs)+",\n")
	s = append(s, "SuspicionScore: "+fmt.Sprintf("%#v", this.SuspicionScore)+",\n")
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
	// Suspicious User Status
	//
	// x-displayName: "Get Status of Suspicious users"
	// Get status of suspicious users
	SuspiciousUserStatus(ctx context.Context, in *SuspiciousUserStatusReq, opts ...grpc.CallOption) (*SuspiciousUserStatusRsp, error)
}

type customAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomAPIClient(cc *grpc.ClientConn) CustomAPIClient {
	return &customAPIClient{cc}
}

func (c *customAPIClient) SuspiciousUserStatus(ctx context.Context, in *SuspiciousUserStatusReq, opts ...grpc.CallOption) (*SuspiciousUserStatusRsp, error) {
	out := new(SuspiciousUserStatusRsp)
	err := c.cc.Invoke(ctx, "/ves.io.schema.app_setting.CustomAPI/SuspiciousUserStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAPIServer is the server API for CustomAPI service.
type CustomAPIServer interface {
	// Suspicious User Status
	//
	// x-displayName: "Get Status of Suspicious users"
	// Get status of suspicious users
	SuspiciousUserStatus(context.Context, *SuspiciousUserStatusReq) (*SuspiciousUserStatusRsp, error)
}

// UnimplementedCustomAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomAPIServer struct {
}

func (*UnimplementedCustomAPIServer) SuspiciousUserStatus(ctx context.Context, req *SuspiciousUserStatusReq) (*SuspiciousUserStatusRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuspiciousUserStatus not implemented")
}

func RegisterCustomAPIServer(s *grpc.Server, srv CustomAPIServer) {
	s.RegisterService(&_CustomAPI_serviceDesc, srv)
}

func _CustomAPI_SuspiciousUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuspiciousUserStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAPIServer).SuspiciousUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.app_setting.CustomAPI/SuspiciousUserStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAPIServer).SuspiciousUserStatus(ctx, req.(*SuspiciousUserStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.app_setting.CustomAPI",
	HandlerType: (*CustomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuspiciousUserStatus",
			Handler:    _CustomAPI_SuspiciousUserStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/app_setting/public_customapi.proto",
}

func (m *SuspiciousUserStatusReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuspiciousUserStatusReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuspiciousUserStatusReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Topn != 0 {
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(m.Topn))
		i--
		dAtA[i] = 0x30
	}
	if len(m.EndTime) > 0 {
		i -= len(m.EndTime)
		copy(dAtA[i:], m.EndTime)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.EndTime)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.StartTime) > 0 {
		i -= len(m.StartTime)
		copy(dAtA[i:], m.StartTime)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.StartTime)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0x1a
	}
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

func (m *SuspiciousUserStatusRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuspiciousUserStatusRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuspiciousUserStatusRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SuspiciousUsers) > 0 {
		for iNdEx := len(m.SuspiciousUsers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SuspiciousUsers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPublicCustomapi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SuspiciousUser) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuspiciousUser) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuspiciousUser) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SuspicionScore != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.SuspicionScore))))
		i--
		dAtA[i] = 0x19
	}
	if len(m.Logs) > 0 {
		for iNdEx := len(m.Logs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Logs[iNdEx])
			copy(dAtA[i:], m.Logs[iNdEx])
			i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.Logs[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.UserId) > 0 {
		i -= len(m.UserId)
		copy(dAtA[i:], m.UserId)
		i = encodeVarintPublicCustomapi(dAtA, i, uint64(len(m.UserId)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *SuspiciousUserStatusReq) Size() (n int) {
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
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.StartTime)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	l = len(m.EndTime)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	if m.Topn != 0 {
		n += 1 + sovPublicCustomapi(uint64(m.Topn))
	}
	return n
}

func (m *SuspiciousUserStatusRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SuspiciousUsers) > 0 {
		for _, e := range m.SuspiciousUsers {
			l = e.Size()
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
	return n
}

func (m *SuspiciousUser) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovPublicCustomapi(uint64(l))
	}
	if len(m.Logs) > 0 {
		for _, s := range m.Logs {
			l = len(s)
			n += 1 + l + sovPublicCustomapi(uint64(l))
		}
	}
	if m.SuspicionScore != 0 {
		n += 9
	}
	return n
}

func sovPublicCustomapi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublicCustomapi(x uint64) (n int) {
	return sovPublicCustomapi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SuspiciousUserStatusReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SuspiciousUserStatusReq{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Query:` + fmt.Sprintf("%v", this.Query) + `,`,
		`StartTime:` + fmt.Sprintf("%v", this.StartTime) + `,`,
		`EndTime:` + fmt.Sprintf("%v", this.EndTime) + `,`,
		`Topn:` + fmt.Sprintf("%v", this.Topn) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SuspiciousUserStatusRsp) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForSuspiciousUsers := "[]*SuspiciousUser{"
	for _, f := range this.SuspiciousUsers {
		repeatedStringForSuspiciousUsers += strings.Replace(f.String(), "SuspiciousUser", "SuspiciousUser", 1) + ","
	}
	repeatedStringForSuspiciousUsers += "}"
	s := strings.Join([]string{`&SuspiciousUserStatusRsp{`,
		`SuspiciousUsers:` + repeatedStringForSuspiciousUsers + `,`,
		`}`,
	}, "")
	return s
}
func (this *SuspiciousUser) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SuspiciousUser{`,
		`UserId:` + fmt.Sprintf("%v", this.UserId) + `,`,
		`Logs:` + fmt.Sprintf("%v", this.Logs) + `,`,
		`SuspicionScore:` + fmt.Sprintf("%v", this.SuspicionScore) + `,`,
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
func (m *SuspiciousUserStatusReq) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SuspiciousUserStatusReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuspiciousUserStatusReq: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
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
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
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
			m.StartTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
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
			m.EndTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topn", wireType)
			}
			m.Topn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Topn |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *SuspiciousUserStatusRsp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SuspiciousUserStatusRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuspiciousUserStatusRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuspiciousUsers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublicCustomapi
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
				return ErrInvalidLengthPublicCustomapi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublicCustomapi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SuspiciousUsers = append(m.SuspiciousUsers, &SuspiciousUser{})
			if err := m.SuspiciousUsers[len(m.SuspiciousUsers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *SuspiciousUser) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SuspiciousUser: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuspiciousUser: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
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
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
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
			m.Logs = append(m.Logs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuspicionScore", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.SuspicionScore = float64(math.Float64frombits(v))
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
