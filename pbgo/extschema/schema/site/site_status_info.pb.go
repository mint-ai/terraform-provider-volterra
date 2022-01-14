// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/site/site_status_info.proto

// Site Status Info
//
// x-displayName: "Site"
// APIs to get status of a site or individual nodes in a multi-node site.
// Some of the status metrics for a site are listed below:
// - Cpu, Memory and Disk usage
// - In/Out throughput
// - Number of active flows and flow setup rate
// - Number active system and vK8s pods
// - Throughput to Regional Edges (REs)

package site

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Site Status Metrics Request
//
// x-displayName: "Site Status Metrics Request"
// Request to get status metrics for a site
type SiteStatusMetricsRequest struct {
	// Namespace
	//
	// x-displayName: "Namespace"
	// x-required
	// x-example: "system"
	// Only system namespace is valid for this request
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Site
	//
	// x-displayName: "Site"
	// x-required
	// x-example: "ce01"
	// Name of the site
	Site string `protobuf:"bytes,2,opt,name=site,proto3" json:"site,omitempty"`
	// Field Selector
	//
	// x-displayName: "Field Selector"
	// x-required
	// List of fields to be returned in the response
	FieldSelector []SiteStatusMetricsField `protobuf:"varint,3,rep,packed,name=field_selector,json=fieldSelector,proto3,enum=ves.io.schema.site.SiteStatusMetricsField" json:"field_selector,omitempty"`
	// Start time
	//
	// x-displayName: "Start Time"
	// x-example: "1570007981"
	// start time of metric collection from which data will be considered.
	// Format: unix_timestamp|rfc 3339
	//
	// Optional: If not specified, then the start_time will be evaluated to end_time-10m
	//           If end_time is not specified, then the start_time will be evaluated to <current time>-10m
	StartTime string `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End time
	//
	// x-displayName: "End Time"
	// x-example: "1570007981"
	// end time of metric collection from which data will be considered.
	// Format: unix_timestamp|rfc 3339
	//
	// Optional: If not specified, then the end_time will be evaluated to start_time+10m
	//           If start_time is not specified, then the end_time will be evaluated to <current time>
	EndTime string `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Step
	//
	// x-displayName: "Step"
	// x-example: "15m"
	// step is the resolution width, which determines the number of the data points [x-axis (time)] to be returned in the response.
	// The timestamps in the response will be t1=start_time, t2=t1+step, ... tn=tn-1+step, where tn <= end_time.
	// Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days
	//
	// Optional: If not specified, then step size is evaluated to <end_time - start_time>
	Step string `protobuf:"bytes,6,opt,name=step,proto3" json:"step,omitempty"`
}

func (m *SiteStatusMetricsRequest) Reset()      { *m = SiteStatusMetricsRequest{} }
func (*SiteStatusMetricsRequest) ProtoMessage() {}
func (*SiteStatusMetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc66d63407bdffe4, []int{0}
}
func (m *SiteStatusMetricsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SiteStatusMetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SiteStatusMetricsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SiteStatusMetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SiteStatusMetricsRequest.Merge(m, src)
}
func (m *SiteStatusMetricsRequest) XXX_Size() int {
	return m.Size()
}
func (m *SiteStatusMetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SiteStatusMetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SiteStatusMetricsRequest proto.InternalMessageInfo

func (m *SiteStatusMetricsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SiteStatusMetricsRequest) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

func (m *SiteStatusMetricsRequest) GetFieldSelector() []SiteStatusMetricsField {
	if m != nil {
		return m.FieldSelector
	}
	return nil
}

func (m *SiteStatusMetricsRequest) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *SiteStatusMetricsRequest) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *SiteStatusMetricsRequest) GetStep() string {
	if m != nil {
		return m.Step
	}
	return ""
}

// Site Status Metrics Response
//
// x-displayName: "Site Status Metrics Response"
// Response for the Site Status Metrics Request
type SiteStatusMetricsResponse struct {
	// Data
	//
	// x-displayName: "Data"
	// Data contains time-series data for the status fields in the request
	Data []*SiteStatusMetricsData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *SiteStatusMetricsResponse) Reset()      { *m = SiteStatusMetricsResponse{} }
func (*SiteStatusMetricsResponse) ProtoMessage() {}
func (*SiteStatusMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc66d63407bdffe4, []int{1}
}
func (m *SiteStatusMetricsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SiteStatusMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SiteStatusMetricsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SiteStatusMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SiteStatusMetricsResponse.Merge(m, src)
}
func (m *SiteStatusMetricsResponse) XXX_Size() int {
	return m.Size()
}
func (m *SiteStatusMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SiteStatusMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SiteStatusMetricsResponse proto.InternalMessageInfo

func (m *SiteStatusMetricsResponse) GetData() []*SiteStatusMetricsData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SiteStatusMetricsRequest)(nil), "ves.io.schema.site.SiteStatusMetricsRequest")
	golang_proto.RegisterType((*SiteStatusMetricsRequest)(nil), "ves.io.schema.site.SiteStatusMetricsRequest")
	proto.RegisterType((*SiteStatusMetricsResponse)(nil), "ves.io.schema.site.SiteStatusMetricsResponse")
	golang_proto.RegisterType((*SiteStatusMetricsResponse)(nil), "ves.io.schema.site.SiteStatusMetricsResponse")
}

func init() {
	proto.RegisterFile("ves.io/schema/site/site_status_info.proto", fileDescriptor_bc66d63407bdffe4)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/site/site_status_info.proto", fileDescriptor_bc66d63407bdffe4)
}

var fileDescriptor_bc66d63407bdffe4 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6a, 0x14, 0x4d,
	0x14, 0x86, 0xa7, 0x66, 0xf2, 0xe5, 0x33, 0x25, 0x06, 0x6c, 0x37, 0x3d, 0x63, 0x2c, 0xc2, 0x6c,
	0x4c, 0x82, 0xdd, 0x05, 0x71, 0xa5, 0xa0, 0xe0, 0x0f, 0xa2, 0x0b, 0x51, 0x27, 0xae, 0xb2, 0x19,
	0xaa, 0x7b, 0xce, 0x74, 0x4a, 0xa7, 0xbb, 0xca, 0xaa, 0xd3, 0x83, 0x12, 0x02, 0x92, 0x2b, 0x10,
	0xc4, 0x7b, 0xf0, 0x1e, 0xb2, 0xc9, 0x4e, 0x57, 0x12, 0xd4, 0x45, 0x96, 0xa6, 0xc7, 0x85, 0x1b,
	0x21, 0x97, 0x20, 0x5d, 0x3d, 0x93, 0x64, 0x48, 0x16, 0xd9, 0x14, 0xa7, 0xea, 0x79, 0xcf, 0xa9,
	0xb7, 0x7e, 0x0e, 0x5d, 0x1e, 0x82, 0x0d, 0xa5, 0xe2, 0x36, 0xde, 0x80, 0x54, 0x70, 0x2b, 0x11,
	0xdc, 0xd0, 0xb5, 0x28, 0x30, 0xb7, 0x5d, 0x99, 0xf5, 0x55, 0xa8, 0x8d, 0x42, 0xe5, 0x79, 0x95,
	0x34, 0xac, 0xa4, 0x61, 0xa9, 0x6a, 0x05, 0x89, 0xc4, 0x8d, 0x3c, 0x0a, 0x63, 0x95, 0xf2, 0x44,
	0x25, 0x8a, 0x3b, 0x69, 0x94, 0xf7, 0xdd, 0xcc, 0x4d, 0x5c, 0x54, 0x95, 0x68, 0x2d, 0x24, 0x4a,
	0x25, 0x03, 0xe0, 0x42, 0x4b, 0x2e, 0xb2, 0x4c, 0xa1, 0x40, 0xa9, 0x32, 0x3b, 0xa6, 0xad, 0x69,
	0x2f, 0x60, 0x8c, 0x32, 0x13, 0x76, 0x75, 0x9a, 0x29, 0x7d, 0x32, 0x91, 0x9d, 0x71, 0x08, 0x7c,
	0xa7, 0x61, 0xc2, 0x9b, 0xd3, 0xfc, 0x24, 0x6a, 0x4f, 0xa3, 0x21, 0x58, 0xc8, 0x86, 0xd3, 0xe5,
	0xdb, 0x7f, 0x09, 0xf5, 0xd7, 0x24, 0xc2, 0x9a, 0xbb, 0x92, 0xa7, 0x80, 0x46, 0xc6, 0xb6, 0x03,
	0x6f, 0x72, 0xb0, 0xe8, 0x2d, 0xd0, 0xb9, 0x4c, 0xa4, 0x60, 0xb5, 0x88, 0xc1, 0x27, 0x8b, 0x64,
	0x69, 0xae, 0x73, 0xbc, 0xe0, 0x79, 0x74, 0xa6, 0x74, 0xe3, 0xd7, 0x1d, 0x70, 0xb1, 0xf7, 0x82,
	0xce, 0xf7, 0x25, 0x0c, 0x7a, 0x5d, 0x0b, 0x03, 0x88, 0x51, 0x19, 0xbf, 0xb1, 0xd8, 0x58, 0x9a,
	0x5f, 0x5d, 0x09, 0x4f, 0x5f, 0x70, 0x78, 0x6a, 0xdf, 0x47, 0x65, 0x6a, 0xe7, 0x92, 0xab, 0xb0,
	0x36, 0x2e, 0xe0, 0x5d, 0xa3, 0xd4, 0xa2, 0x30, 0xd8, 0x45, 0x99, 0x82, 0x3f, 0x53, 0xb9, 0x70,
	0x2b, 0x2f, 0x65, 0x0a, 0x5e, 0x93, 0x5e, 0x80, 0xac, 0x57, 0xc1, 0xff, 0x1c, 0xfc, 0x1f, 0xb2,
	0x9e, 0x43, 0xa5, 0x41, 0x04, 0xed, 0xcf, 0x8e, 0x0d, 0x22, 0xe8, 0xf6, 0x3a, 0x6d, 0x9e, 0x71,
	0x5c, 0xab, 0x55, 0x66, 0xc1, 0xbb, 0x43, 0x67, 0x7a, 0x02, 0x85, 0x4f, 0x16, 0x1b, 0x4b, 0x17,
	0x57, 0x97, 0xcf, 0xe5, 0xf9, 0xa1, 0x40, 0xd1, 0x71, 0x69, 0xab, 0x9f, 0xea, 0xf4, 0xca, 0x83,
	0xdc, 0xa2, 0x4a, 0x8f, 0x55, 0xf7, 0x9e, 0x3f, 0xf1, 0x7e, 0x12, 0x7a, 0xf9, 0x54, 0x9e, 0x77,
	0xe3, 0x5c, 0xe5, 0xc7, 0x4f, 0xd1, 0x0a, 0xce, 0xa9, 0xae, 0x4e, 0xd2, 0x8e, 0x8a, 0x2f, 0x7e,
	0xf9, 0xe5, 0x02, 0xa9, 0x02, 0x99, 0xf5, 0x8d, 0xb0, 0x68, 0xf2, 0x18, 0x73, 0x03, 0x81, 0x01,
	0xd1, 0xdb, 0xfe, 0xf1, 0xfb, 0x63, 0xfd, 0x6e, 0xfb, 0x16, 0xd7, 0x79, 0x34, 0x90, 0x31, 0x3f,
	0x7a, 0x54, 0xcb, 0x37, 0x8f, 0xe2, 0xad, 0xea, 0xb3, 0x6d, 0x96, 0xe3, 0x16, 0xaf, 0x7a, 0x86,
	0xa7, 0xd5, 0x46, 0xb7, 0xc9, 0x4a, 0xeb, 0xfa, 0xee, 0x0e, 0x69, 0x7c, 0xdf, 0x21, 0xcd, 0x33,
	0x9c, 0x3d, 0x8b, 0x5e, 0x41, 0x8c, 0xdb, 0xdf, 0xfc, 0xba, 0x4f, 0xee, 0x6f, 0x93, 0xbd, 0x03,
	0x56, 0xdb, 0x3f, 0x60, 0xb5, 0xc3, 0x03, 0x46, 0xde, 0x17, 0x8c, 0x7c, 0x2e, 0x18, 0xf9, 0x5a,
	0x30, 0xb2, 0x57, 0x30, 0xf2, 0xab, 0x60, 0xe4, 0x4f, 0xc1, 0x6a, 0x87, 0x05, 0x23, 0x1f, 0x46,
	0xac, 0xb6, 0x3b, 0x62, 0x64, 0x6f, 0xc4, 0x6a, 0xfb, 0x23, 0x56, 0x5b, 0x7f, 0x9c, 0x28, 0xfd,
	0x3a, 0x09, 0x87, 0x6a, 0x80, 0x60, 0x8c, 0x08, 0x73, 0xcb, 0x5d, 0xd0, 0x57, 0x26, 0x0d, 0xb4,
	0x51, 0x43, 0xd9, 0x03, 0x13, 0x4c, 0x30, 0xd7, 0x51, 0xa2, 0x38, 0xbc, 0xc5, 0x49, 0xb3, 0x1c,
	0xf7, 0x4c, 0x34, 0xeb, 0xfe, 0xfb, 0xcd, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31, 0x5c, 0xcb,
	0x81, 0x15, 0x04, 0x00, 0x00,
}

func (this *SiteStatusMetricsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SiteStatusMetricsRequest)
	if !ok {
		that2, ok := that.(SiteStatusMetricsRequest)
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
	if this.Site != that1.Site {
		return false
	}
	if len(this.FieldSelector) != len(that1.FieldSelector) {
		return false
	}
	for i := range this.FieldSelector {
		if this.FieldSelector[i] != that1.FieldSelector[i] {
			return false
		}
	}
	if this.StartTime != that1.StartTime {
		return false
	}
	if this.EndTime != that1.EndTime {
		return false
	}
	if this.Step != that1.Step {
		return false
	}
	return true
}
func (this *SiteStatusMetricsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SiteStatusMetricsResponse)
	if !ok {
		that2, ok := that.(SiteStatusMetricsResponse)
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
	if len(this.Data) != len(that1.Data) {
		return false
	}
	for i := range this.Data {
		if !this.Data[i].Equal(that1.Data[i]) {
			return false
		}
	}
	return true
}
func (this *SiteStatusMetricsRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&site.SiteStatusMetricsRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	s = append(s, "Site: "+fmt.Sprintf("%#v", this.Site)+",\n")
	s = append(s, "FieldSelector: "+fmt.Sprintf("%#v", this.FieldSelector)+",\n")
	s = append(s, "StartTime: "+fmt.Sprintf("%#v", this.StartTime)+",\n")
	s = append(s, "EndTime: "+fmt.Sprintf("%#v", this.EndTime)+",\n")
	s = append(s, "Step: "+fmt.Sprintf("%#v", this.Step)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SiteStatusMetricsResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&site.SiteStatusMetricsResponse{")
	if this.Data != nil {
		s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSiteStatusInfo(v interface{}, typ string) string {
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

// CustomSiteStatusAPIClient is the client API for CustomSiteStatusAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomSiteStatusAPIClient interface {
	// Site Status Metrics
	//
	// x-displayName: "Site Status Metrics"
	// Get status metrics for a site
	SiteStatusMetrics(ctx context.Context, in *SiteStatusMetricsRequest, opts ...grpc.CallOption) (*SiteStatusMetricsResponse, error)
}

type customSiteStatusAPIClient struct {
	cc *grpc.ClientConn
}

func NewCustomSiteStatusAPIClient(cc *grpc.ClientConn) CustomSiteStatusAPIClient {
	return &customSiteStatusAPIClient{cc}
}

func (c *customSiteStatusAPIClient) SiteStatusMetrics(ctx context.Context, in *SiteStatusMetricsRequest, opts ...grpc.CallOption) (*SiteStatusMetricsResponse, error) {
	out := new(SiteStatusMetricsResponse)
	err := c.cc.Invoke(ctx, "/ves.io.schema.site.CustomSiteStatusAPI/SiteStatusMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomSiteStatusAPIServer is the server API for CustomSiteStatusAPI service.
type CustomSiteStatusAPIServer interface {
	// Site Status Metrics
	//
	// x-displayName: "Site Status Metrics"
	// Get status metrics for a site
	SiteStatusMetrics(context.Context, *SiteStatusMetricsRequest) (*SiteStatusMetricsResponse, error)
}

// UnimplementedCustomSiteStatusAPIServer can be embedded to have forward compatible implementations.
type UnimplementedCustomSiteStatusAPIServer struct {
}

func (*UnimplementedCustomSiteStatusAPIServer) SiteStatusMetrics(ctx context.Context, req *SiteStatusMetricsRequest) (*SiteStatusMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SiteStatusMetrics not implemented")
}

func RegisterCustomSiteStatusAPIServer(s *grpc.Server, srv CustomSiteStatusAPIServer) {
	s.RegisterService(&_CustomSiteStatusAPI_serviceDesc, srv)
}

func _CustomSiteStatusAPI_SiteStatusMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiteStatusMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomSiteStatusAPIServer).SiteStatusMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ves.io.schema.site.CustomSiteStatusAPI/SiteStatusMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomSiteStatusAPIServer).SiteStatusMetrics(ctx, req.(*SiteStatusMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomSiteStatusAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ves.io.schema.site.CustomSiteStatusAPI",
	HandlerType: (*CustomSiteStatusAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SiteStatusMetrics",
			Handler:    _CustomSiteStatusAPI_SiteStatusMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ves.io/schema/site/site_status_info.proto",
}

func (m *SiteStatusMetricsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SiteStatusMetricsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SiteStatusMetricsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Step) > 0 {
		i -= len(m.Step)
		copy(dAtA[i:], m.Step)
		i = encodeVarintSiteStatusInfo(dAtA, i, uint64(len(m.Step)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.EndTime) > 0 {
		i -= len(m.EndTime)
		copy(dAtA[i:], m.EndTime)
		i = encodeVarintSiteStatusInfo(dAtA, i, uint64(len(m.EndTime)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.StartTime) > 0 {
		i -= len(m.StartTime)
		copy(dAtA[i:], m.StartTime)
		i = encodeVarintSiteStatusInfo(dAtA, i, uint64(len(m.StartTime)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FieldSelector) > 0 {
		dAtA2 := make([]byte, len(m.FieldSelector)*10)
		var j1 int
		for _, num := range m.FieldSelector {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintSiteStatusInfo(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Site) > 0 {
		i -= len(m.Site)
		copy(dAtA[i:], m.Site)
		i = encodeVarintSiteStatusInfo(dAtA, i, uint64(len(m.Site)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintSiteStatusInfo(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SiteStatusMetricsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SiteStatusMetricsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SiteStatusMetricsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSiteStatusInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintSiteStatusInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovSiteStatusInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SiteStatusMetricsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovSiteStatusInfo(uint64(l))
	}
	l = len(m.Site)
	if l > 0 {
		n += 1 + l + sovSiteStatusInfo(uint64(l))
	}
	if len(m.FieldSelector) > 0 {
		l = 0
		for _, e := range m.FieldSelector {
			l += sovSiteStatusInfo(uint64(e))
		}
		n += 1 + sovSiteStatusInfo(uint64(l)) + l
	}
	l = len(m.StartTime)
	if l > 0 {
		n += 1 + l + sovSiteStatusInfo(uint64(l))
	}
	l = len(m.EndTime)
	if l > 0 {
		n += 1 + l + sovSiteStatusInfo(uint64(l))
	}
	l = len(m.Step)
	if l > 0 {
		n += 1 + l + sovSiteStatusInfo(uint64(l))
	}
	return n
}

func (m *SiteStatusMetricsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovSiteStatusInfo(uint64(l))
		}
	}
	return n
}

func sovSiteStatusInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSiteStatusInfo(x uint64) (n int) {
	return sovSiteStatusInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SiteStatusMetricsRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SiteStatusMetricsRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Site:` + fmt.Sprintf("%v", this.Site) + `,`,
		`FieldSelector:` + fmt.Sprintf("%v", this.FieldSelector) + `,`,
		`StartTime:` + fmt.Sprintf("%v", this.StartTime) + `,`,
		`EndTime:` + fmt.Sprintf("%v", this.EndTime) + `,`,
		`Step:` + fmt.Sprintf("%v", this.Step) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SiteStatusMetricsResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForData := "[]*SiteStatusMetricsData{"
	for _, f := range this.Data {
		repeatedStringForData += strings.Replace(fmt.Sprintf("%v", f), "SiteStatusMetricsData", "SiteStatusMetricsData", 1) + ","
	}
	repeatedStringForData += "}"
	s := strings.Join([]string{`&SiteStatusMetricsResponse{`,
		`Data:` + repeatedStringForData + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSiteStatusInfo(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SiteStatusMetricsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSiteStatusInfo
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
			return fmt.Errorf("proto: SiteStatusMetricsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SiteStatusMetricsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSiteStatusInfo
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
				return ErrInvalidLengthSiteStatusInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSiteStatusInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Site", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSiteStatusInfo
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
				return ErrInvalidLengthSiteStatusInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSiteStatusInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Site = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v SiteStatusMetricsField
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSiteStatusInfo
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= SiteStatusMetricsField(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.FieldSelector = append(m.FieldSelector, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSiteStatusInfo
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthSiteStatusInfo
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthSiteStatusInfo
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.FieldSelector) == 0 {
					m.FieldSelector = make([]SiteStatusMetricsField, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v SiteStatusMetricsField
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSiteStatusInfo
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= SiteStatusMetricsField(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.FieldSelector = append(m.FieldSelector, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldSelector", wireType)
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSiteStatusInfo
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
				return ErrInvalidLengthSiteStatusInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSiteStatusInfo
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
					return ErrIntOverflowSiteStatusInfo
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
				return ErrInvalidLengthSiteStatusInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSiteStatusInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSiteStatusInfo
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
				return ErrInvalidLengthSiteStatusInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSiteStatusInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Step = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSiteStatusInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSiteStatusInfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSiteStatusInfo
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
func (m *SiteStatusMetricsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSiteStatusInfo
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
			return fmt.Errorf("proto: SiteStatusMetricsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SiteStatusMetricsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSiteStatusInfo
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
				return ErrInvalidLengthSiteStatusInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSiteStatusInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &SiteStatusMetricsData{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSiteStatusInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSiteStatusInfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSiteStatusInfo
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
func skipSiteStatusInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSiteStatusInfo
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
					return 0, ErrIntOverflowSiteStatusInfo
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
					return 0, ErrIntOverflowSiteStatusInfo
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
				return 0, ErrInvalidLengthSiteStatusInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSiteStatusInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSiteStatusInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSiteStatusInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSiteStatusInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSiteStatusInfo = fmt.Errorf("proto: unexpected end of group")
)
