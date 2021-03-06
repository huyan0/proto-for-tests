// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opentelemetry/proto/logs/v1/logs.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v11 "github.com/open-telemetry/opentelemetry-proto/gen/go/common/v1"
	v1 "github.com/open-telemetry/opentelemetry-proto/gen/go/resource/v1"
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

// Possible values for LogRecord.SeverityNumber.
type SeverityNumber int32

const (
	SeverityNumber_UNDEFINED_SEVERITY_NUMBER SeverityNumber = 0
	SeverityNumber_TRACE                     SeverityNumber = 1
	SeverityNumber_TRACE2                    SeverityNumber = 2
	SeverityNumber_TRACE3                    SeverityNumber = 3
	SeverityNumber_TRACE4                    SeverityNumber = 4
	SeverityNumber_DEBUG                     SeverityNumber = 5
	SeverityNumber_DEBUG2                    SeverityNumber = 6
	SeverityNumber_DEBUG3                    SeverityNumber = 7
	SeverityNumber_DEBUG4                    SeverityNumber = 8
	SeverityNumber_INFO                      SeverityNumber = 9
	SeverityNumber_INFO2                     SeverityNumber = 10
	SeverityNumber_INFO3                     SeverityNumber = 11
	SeverityNumber_INFO4                     SeverityNumber = 12
	SeverityNumber_WARN                      SeverityNumber = 13
	SeverityNumber_WARN2                     SeverityNumber = 14
	SeverityNumber_WARN3                     SeverityNumber = 15
	SeverityNumber_WARN4                     SeverityNumber = 16
	SeverityNumber_ERROR                     SeverityNumber = 17
	SeverityNumber_ERROR2                    SeverityNumber = 18
	SeverityNumber_ERROR3                    SeverityNumber = 19
	SeverityNumber_ERROR4                    SeverityNumber = 20
	SeverityNumber_FATAL                     SeverityNumber = 21
	SeverityNumber_FATAL2                    SeverityNumber = 22
	SeverityNumber_FATAL3                    SeverityNumber = 23
	SeverityNumber_FATAL4                    SeverityNumber = 24
)

var SeverityNumber_name = map[int32]string{
	0:  "UNDEFINED_SEVERITY_NUMBER",
	1:  "TRACE",
	2:  "TRACE2",
	3:  "TRACE3",
	4:  "TRACE4",
	5:  "DEBUG",
	6:  "DEBUG2",
	7:  "DEBUG3",
	8:  "DEBUG4",
	9:  "INFO",
	10: "INFO2",
	11: "INFO3",
	12: "INFO4",
	13: "WARN",
	14: "WARN2",
	15: "WARN3",
	16: "WARN4",
	17: "ERROR",
	18: "ERROR2",
	19: "ERROR3",
	20: "ERROR4",
	21: "FATAL",
	22: "FATAL2",
	23: "FATAL3",
	24: "FATAL4",
}

var SeverityNumber_value = map[string]int32{
	"UNDEFINED_SEVERITY_NUMBER": 0,
	"TRACE":                     1,
	"TRACE2":                    2,
	"TRACE3":                    3,
	"TRACE4":                    4,
	"DEBUG":                     5,
	"DEBUG2":                    6,
	"DEBUG3":                    7,
	"DEBUG4":                    8,
	"INFO":                      9,
	"INFO2":                     10,
	"INFO3":                     11,
	"INFO4":                     12,
	"WARN":                      13,
	"WARN2":                     14,
	"WARN3":                     15,
	"WARN4":                     16,
	"ERROR":                     17,
	"ERROR2":                    18,
	"ERROR3":                    19,
	"ERROR4":                    20,
	"FATAL":                     21,
	"FATAL2":                    22,
	"FATAL3":                    23,
	"FATAL4":                    24,
}

func (x SeverityNumber) String() string {
	return proto.EnumName(SeverityNumber_name, int32(x))
}

func (SeverityNumber) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{0}
}

// Masks for LogRecord.flags field.
type LogRecordFlags int32

const (
	LogRecordFlags_UNDEFINED_LOG_RECORD_FLAG LogRecordFlags = 0
	LogRecordFlags_TRACE_FLAGS_MASK          LogRecordFlags = 255
)

var LogRecordFlags_name = map[int32]string{
	0:   "UNDEFINED_LOG_RECORD_FLAG",
	255: "TRACE_FLAGS_MASK",
}

var LogRecordFlags_value = map[string]int32{
	"UNDEFINED_LOG_RECORD_FLAG": 0,
	"TRACE_FLAGS_MASK":          255,
}

func (x LogRecordFlags) String() string {
	return proto.EnumName(LogRecordFlags_name, int32(x))
}

func (LogRecordFlags) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{1}
}

// A collection of InstrumentationLibraryLogs from a Resource.
type ResourceLogs struct {
	// The resource for the logs in this message.
	// If this field is not set then no resource info is known.
	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// A list of InstrumentationLibraryLogs that originate from a resource.
	InstrumentationLibraryLogs []*InstrumentationLibraryLogs `protobuf:"bytes,2,rep,name=instrumentation_library_logs,json=instrumentationLibraryLogs,proto3" json:"instrumentation_library_logs,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                      `json:"-"`
	XXX_unrecognized           []byte                        `json:"-"`
	XXX_sizecache              int32                         `json:"-"`
}

func (m *ResourceLogs) Reset()         { *m = ResourceLogs{} }
func (m *ResourceLogs) String() string { return proto.CompactTextString(m) }
func (*ResourceLogs) ProtoMessage()    {}
func (*ResourceLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{0}
}

func (m *ResourceLogs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceLogs.Unmarshal(m, b)
}
func (m *ResourceLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceLogs.Marshal(b, m, deterministic)
}
func (m *ResourceLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceLogs.Merge(m, src)
}
func (m *ResourceLogs) XXX_Size() int {
	return xxx_messageInfo_ResourceLogs.Size(m)
}
func (m *ResourceLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceLogs.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceLogs proto.InternalMessageInfo

func (m *ResourceLogs) GetResource() *v1.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ResourceLogs) GetInstrumentationLibraryLogs() []*InstrumentationLibraryLogs {
	if m != nil {
		return m.InstrumentationLibraryLogs
	}
	return nil
}

// A collection of Logs produced by an InstrumentationLibrary.
type InstrumentationLibraryLogs struct {
	// The instrumentation library information for the logs in this message.
	// If this field is not set then no library info is known.
	InstrumentationLibrary *v11.InstrumentationLibrary `protobuf:"bytes,1,opt,name=instrumentation_library,json=instrumentationLibrary,proto3" json:"instrumentation_library,omitempty"`
	// A list of log records.
	Logs                 []*LogRecord `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *InstrumentationLibraryLogs) Reset()         { *m = InstrumentationLibraryLogs{} }
func (m *InstrumentationLibraryLogs) String() string { return proto.CompactTextString(m) }
func (*InstrumentationLibraryLogs) ProtoMessage()    {}
func (*InstrumentationLibraryLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{1}
}

func (m *InstrumentationLibraryLogs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstrumentationLibraryLogs.Unmarshal(m, b)
}
func (m *InstrumentationLibraryLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstrumentationLibraryLogs.Marshal(b, m, deterministic)
}
func (m *InstrumentationLibraryLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentationLibraryLogs.Merge(m, src)
}
func (m *InstrumentationLibraryLogs) XXX_Size() int {
	return xxx_messageInfo_InstrumentationLibraryLogs.Size(m)
}
func (m *InstrumentationLibraryLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentationLibraryLogs.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentationLibraryLogs proto.InternalMessageInfo

func (m *InstrumentationLibraryLogs) GetInstrumentationLibrary() *v11.InstrumentationLibrary {
	if m != nil {
		return m.InstrumentationLibrary
	}
	return nil
}

func (m *InstrumentationLibraryLogs) GetLogs() []*LogRecord {
	if m != nil {
		return m.Logs
	}
	return nil
}

// A log record according to OpenTelemetry Log Data Model:
// https://github.com/open-telemetry/oteps/blob/master/text/logs/0097-log-data-model.md
type LogRecord struct {
	// time_unix_nano is the time when the event occurred.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	// Value of 0 indicates unknown or missing timestamp.
	TimeUnixNano uint64 `protobuf:"fixed64,1,opt,name=time_unix_nano,json=timeUnixNano,proto3" json:"time_unix_nano,omitempty"`
	// Numerical value of the severity, normalized to values described in Log Data Model.
	// [Optional].
	SeverityNumber SeverityNumber `protobuf:"varint,2,opt,name=severity_number,json=severityNumber,proto3,enum=opentelemetry.proto.logs.v1.SeverityNumber" json:"severity_number,omitempty"`
	// The severity text (also known as log level). The original string representation as
	// it is known at the source. [Optional].
	SeverityText string `protobuf:"bytes,3,opt,name=severity_text,json=severityText,proto3" json:"severity_text,omitempty"`
	// Short event identifier that does not contain varying parts. Name describes
	// what happened (e.g. "ProcessStarted"). Recommended to be no longer than 50
	// characters. Not guaranteed to be unique in any way. [Optional].
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// A value containing the body of the log record. Can be for example a human-readable
	// string message (including multi-line) describing the event in a free form or it can
	// be a structured data composed of arrays and maps of other values. [Optional].
	Body *v11.AnyValue `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// Additional attributes that describe the specific event occurrence. [Optional].
	Attributes             []*v11.KeyValue `protobuf:"bytes,6,rep,name=attributes,proto3" json:"attributes,omitempty"`
	DroppedAttributesCount uint32          `protobuf:"varint,7,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	// Flags, a bit field. 8 least significant bits are the trace flags as
	// defined in W3C Trace Context specification. 24 most significant bits are reserved
	// and must be set to 0. Readers must not assume that 24 most significant bits
	// will be zero and must correctly mask the bits when reading 8-bit trace flag (use
	// flags & TRACE_FLAGS_MASK). [Optional].
	Flags uint32 `protobuf:"fixed32,8,opt,name=flags,proto3" json:"flags,omitempty"`
	// A unique identifier for a trace. All logs from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
	// is considered invalid. Can be set for logs that are part of request processing
	// and have an assigned trace id. [Optional].
	TraceId []byte `protobuf:"bytes,9,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes is considered
	// invalid. Can be set for logs that are part of a particular processing span.
	// If span_id is present trace_id SHOULD be also present. [Optional].
	SpanId               []byte   `protobuf:"bytes,10,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRecord) Reset()         { *m = LogRecord{} }
func (m *LogRecord) String() string { return proto.CompactTextString(m) }
func (*LogRecord) ProtoMessage()    {}
func (*LogRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1c030a3ec7e961e, []int{2}
}

func (m *LogRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRecord.Unmarshal(m, b)
}
func (m *LogRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRecord.Marshal(b, m, deterministic)
}
func (m *LogRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRecord.Merge(m, src)
}
func (m *LogRecord) XXX_Size() int {
	return xxx_messageInfo_LogRecord.Size(m)
}
func (m *LogRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LogRecord proto.InternalMessageInfo

func (m *LogRecord) GetTimeUnixNano() uint64 {
	if m != nil {
		return m.TimeUnixNano
	}
	return 0
}

func (m *LogRecord) GetSeverityNumber() SeverityNumber {
	if m != nil {
		return m.SeverityNumber
	}
	return SeverityNumber_UNDEFINED_SEVERITY_NUMBER
}

func (m *LogRecord) GetSeverityText() string {
	if m != nil {
		return m.SeverityText
	}
	return ""
}

func (m *LogRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogRecord) GetBody() *v11.AnyValue {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *LogRecord) GetAttributes() []*v11.KeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *LogRecord) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

func (m *LogRecord) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *LogRecord) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *LogRecord) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func init() {
	proto.RegisterEnum("opentelemetry.proto.logs.v1.SeverityNumber", SeverityNumber_name, SeverityNumber_value)
	proto.RegisterEnum("opentelemetry.proto.logs.v1.LogRecordFlags", LogRecordFlags_name, LogRecordFlags_value)
	proto.RegisterType((*ResourceLogs)(nil), "opentelemetry.proto.logs.v1.ResourceLogs")
	proto.RegisterType((*InstrumentationLibraryLogs)(nil), "opentelemetry.proto.logs.v1.InstrumentationLibraryLogs")
	proto.RegisterType((*LogRecord)(nil), "opentelemetry.proto.logs.v1.LogRecord")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/logs/v1/logs.proto", fileDescriptor_d1c030a3ec7e961e)
}

var fileDescriptor_d1c030a3ec7e961e = []byte{
	// 733 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x5d, 0x6f, 0xda, 0x3c,
	0x14, 0xc7, 0x1b, 0xde, 0x71, 0x29, 0xf5, 0xe3, 0xa7, 0x2f, 0x29, 0x7b, 0x11, 0xea, 0xa6, 0x8e,
	0x75, 0x2a, 0xa8, 0x81, 0x69, 0xd3, 0xb6, 0x9b, 0x50, 0x02, 0x42, 0xa5, 0x50, 0x19, 0xe8, 0xb4,
	0xdd, 0x44, 0x01, 0x3c, 0x16, 0x0d, 0x6c, 0x94, 0x38, 0x08, 0x3e, 0xde, 0xee, 0xa6, 0x5d, 0xed,
	0x1b, 0x6d, 0xb2, 0x21, 0x69, 0x91, 0x80, 0x5e, 0xf5, 0xe7, 0xff, 0x39, 0xff, 0x93, 0x73, 0x4e,
	0xb1, 0xc1, 0x19, 0x9b, 0x10, 0xca, 0xc9, 0x88, 0x8c, 0x09, 0x77, 0xe6, 0x85, 0x89, 0xc3, 0x38,
	0x2b, 0x8c, 0xd8, 0xd0, 0x2d, 0x4c, 0x2f, 0xe5, 0xdf, 0xbc, 0x94, 0xd0, 0x93, 0x95, 0xbc, 0x85,
	0x98, 0x97, 0xf1, 0xe9, 0x65, 0xe6, 0x7c, 0x5d, 0x91, 0x3e, 0x1b, 0x8f, 0x19, 0x15, 0x65, 0x16,
	0xb4, 0xf0, 0x64, 0xf2, 0xeb, 0x72, 0x1d, 0xe2, 0x32, 0xcf, 0xe9, 0x13, 0x91, 0xed, 0xf3, 0x22,
	0xff, 0xf4, 0x8f, 0x02, 0x52, 0x78, 0x29, 0x35, 0xd8, 0xd0, 0x45, 0x06, 0x48, 0xf8, 0x29, 0xaa,
	0x92, 0x55, 0x72, 0xbb, 0xda, 0xeb, 0xfc, 0xba, 0xe6, 0x82, 0x3a, 0xd3, 0xcb, 0xbc, 0x5f, 0x00,
	0x07, 0x56, 0x34, 0x07, 0x4f, 0x6d, 0xea, 0x72, 0xc7, 0x1b, 0x13, 0xca, 0x2d, 0x6e, 0x33, 0x6a,
	0x8e, 0xec, 0x9e, 0x63, 0x39, 0x73, 0x53, 0x8c, 0xa5, 0x86, 0xb2, 0xe1, 0xdc, 0xae, 0xf6, 0x2e,
	0xbf, 0x65, 0xee, 0x7c, 0x7d, 0xb5, 0x40, 0x63, 0xe1, 0x17, 0x5d, 0xe2, 0x8c, 0xbd, 0x31, 0x76,
	0xfa, 0x4b, 0x01, 0x99, 0xcd, 0x56, 0x44, 0xc1, 0xf1, 0x86, 0xce, 0x96, 0xf3, 0xbe, 0x5d, 0xdb,
	0xd4, 0x72, 0xcb, 0x1b, 0xdb, 0xc2, 0x47, 0xeb, 0x5b, 0x42, 0x1f, 0x40, 0xe4, 0xc1, 0xc4, 0x67,
	0x5b, 0x27, 0x6e, 0xb0, 0x21, 0x26, 0x7d, 0xe6, 0x0c, 0xb0, 0xf4, 0x9c, 0xfe, 0x0e, 0x83, 0x64,
	0xa0, 0xa1, 0x97, 0x20, 0xcd, 0xed, 0x31, 0x31, 0x3d, 0x6a, 0xcf, 0x4c, 0x6a, 0x51, 0x26, 0x1b,
	0x8e, 0xe1, 0x94, 0x50, 0xbb, 0xd4, 0x9e, 0x35, 0x2d, 0xca, 0x50, 0x07, 0xec, 0xbb, 0x64, 0x4a,
	0x1c, 0x9b, 0xcf, 0x4d, 0xea, 0x8d, 0x7b, 0xc4, 0x51, 0x43, 0x59, 0x25, 0x97, 0xd6, 0xde, 0x6c,
	0xfd, 0x74, 0x7b, 0xe9, 0x69, 0x4a, 0x0b, 0x4e, 0xbb, 0x2b, 0x67, 0xf4, 0x02, 0xec, 0x05, 0x55,
	0x39, 0x99, 0x71, 0x35, 0x9c, 0x55, 0x72, 0x49, 0x9c, 0xf2, 0xc5, 0x0e, 0x99, 0x71, 0x84, 0x40,
	0x84, 0x5a, 0x63, 0xa2, 0x46, 0x64, 0x4c, 0x32, 0xfa, 0x08, 0x22, 0x3d, 0x36, 0x98, 0xab, 0x51,
	0xb9, 0xdb, 0x57, 0x8f, 0xec, 0x56, 0xa7, 0xf3, 0x3b, 0x6b, 0xe4, 0x11, 0x2c, 0x4d, 0xa8, 0x06,
	0x80, 0xc5, 0xb9, 0x63, 0xf7, 0x3c, 0x4e, 0x5c, 0x35, 0x26, 0x37, 0xf8, 0x58, 0x89, 0x6b, 0xb2,
	0x2c, 0xf1, 0xc0, 0x8a, 0xde, 0x03, 0x75, 0xe0, 0xb0, 0xc9, 0x84, 0x0c, 0xcc, 0x7b, 0xd5, 0xec,
	0x33, 0x8f, 0x72, 0x35, 0x9e, 0x55, 0x72, 0x7b, 0xf8, 0x68, 0x19, 0xd7, 0x83, 0xf0, 0x95, 0x88,
	0xa2, 0x03, 0x10, 0xfd, 0x36, 0xb2, 0x86, 0xae, 0x9a, 0xc8, 0x2a, 0xb9, 0x38, 0x5e, 0x1c, 0xd0,
	0x09, 0x48, 0x70, 0xc7, 0xea, 0x13, 0xd3, 0x1e, 0xa8, 0xc9, 0xac, 0x92, 0x4b, 0xe1, 0xb8, 0x3c,
	0xd7, 0x07, 0xe8, 0x18, 0xc4, 0xdd, 0x89, 0x45, 0x45, 0x04, 0xc8, 0x48, 0x4c, 0x1c, 0xeb, 0x83,
	0xf3, 0x9f, 0x21, 0x90, 0x5e, 0xdd, 0x32, 0x7a, 0x06, 0x4e, 0xba, 0xcd, 0x8a, 0x51, 0xad, 0x37,
	0x8d, 0x8a, 0xd9, 0x36, 0xee, 0x0c, 0x5c, 0xef, 0x7c, 0x31, 0x9b, 0xdd, 0x9b, 0xb2, 0x81, 0xe1,
	0x0e, 0x4a, 0x82, 0x68, 0x07, 0xeb, 0x57, 0x06, 0x54, 0x10, 0x00, 0x31, 0x89, 0x1a, 0x0c, 0x05,
	0x5c, 0x84, 0xe1, 0x80, 0x4b, 0x30, 0x22, 0xd2, 0x2b, 0x46, 0xb9, 0x5b, 0x83, 0x51, 0x21, 0x4b,
	0xd4, 0x60, 0x2c, 0xe0, 0x22, 0x8c, 0x07, 0x5c, 0x82, 0x09, 0x94, 0x00, 0x91, 0x7a, 0xb3, 0xda,
	0x82, 0x49, 0x61, 0x14, 0xa4, 0x41, 0xe0, 0x63, 0x11, 0xee, 0xfa, 0x58, 0x82, 0x29, 0x91, 0xfa,
	0x59, 0xc7, 0x4d, 0xb8, 0x27, 0x44, 0x41, 0x1a, 0x4c, 0xfb, 0x58, 0x84, 0xfb, 0x3e, 0x96, 0x20,
	0x14, 0x68, 0x60, 0xdc, 0xc2, 0xf0, 0x3f, 0xf1, 0x31, 0x89, 0x1a, 0x44, 0x01, 0x17, 0xe1, 0xff,
	0x01, 0x97, 0xe0, 0x81, 0x48, 0xaf, 0xea, 0x1d, 0xbd, 0x01, 0x0f, 0x85, 0x2c, 0x51, 0x83, 0x47,
	0x01, 0x17, 0xe1, 0x71, 0xc0, 0x25, 0xa8, 0x9e, 0x57, 0x41, 0x3a, 0xb8, 0x0f, 0x55, 0xf9, 0x9f,
	0x58, 0x59, 0x61, 0xa3, 0x55, 0x33, 0xb1, 0x71, 0xd5, 0xc2, 0x15, 0xb3, 0xda, 0xd0, 0x6b, 0x70,
	0x07, 0x1d, 0x02, 0x28, 0xf7, 0x23, 0xcf, 0x6d, 0xf3, 0x46, 0x6f, 0x5f, 0xc3, 0xbf, 0x4a, 0xf9,
	0x07, 0x78, 0x6e, 0xb3, 0x6d, 0xf7, 0xa1, 0x2c, 0xee, 0x9d, 0x7b, 0x2b, 0xa4, 0x5b, 0xe5, 0xeb,
	0xa7, 0xa1, 0xcd, 0xbf, 0x7b, 0x3d, 0xf1, 0x2b, 0x2b, 0x08, 0xd3, 0xc5, 0xfd, 0x0b, 0xbb, 0x52,
	0xe3, 0x62, 0xf1, 0xde, 0x0e, 0x09, 0x2d, 0x0c, 0x83, 0x77, 0xbe, 0x17, 0x93, 0x6a, 0xf1, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x91, 0x45, 0xa9, 0x0d, 0x06, 0x00, 0x00,
}
