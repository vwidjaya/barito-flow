// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/timber.proto

package timber

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Timber struct {
	Context              *TimberContext      `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	Content              map[string]*any.Any `protobuf:"bytes,2,rep,name=content,proto3" json:"content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Timber) Reset()         { *m = Timber{} }
func (m *Timber) String() string { return proto.CompactTextString(m) }
func (*Timber) ProtoMessage()    {}
func (*Timber) Descriptor() ([]byte, []int) {
	return fileDescriptor_timber_04b8121c0f798b35, []int{0}
}
func (m *Timber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timber.Unmarshal(m, b)
}
func (m *Timber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timber.Marshal(b, m, deterministic)
}
func (dst *Timber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timber.Merge(dst, src)
}
func (m *Timber) XXX_Size() int {
	return xxx_messageInfo_Timber.Size(m)
}
func (m *Timber) XXX_DiscardUnknown() {
	xxx_messageInfo_Timber.DiscardUnknown(m)
}

var xxx_messageInfo_Timber proto.InternalMessageInfo

func (m *Timber) GetContext() *TimberContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Timber) GetContent() map[string]*any.Any {
	if m != nil {
		return m.Content
	}
	return nil
}

type TimberCollection struct {
	Context              *TimberContext `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	Items                []*Timber      `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TimberCollection) Reset()         { *m = TimberCollection{} }
func (m *TimberCollection) String() string { return proto.CompactTextString(m) }
func (*TimberCollection) ProtoMessage()    {}
func (*TimberCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_timber_04b8121c0f798b35, []int{1}
}
func (m *TimberCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimberCollection.Unmarshal(m, b)
}
func (m *TimberCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimberCollection.Marshal(b, m, deterministic)
}
func (dst *TimberCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimberCollection.Merge(dst, src)
}
func (m *TimberCollection) XXX_Size() int {
	return xxx_messageInfo_TimberCollection.Size(m)
}
func (m *TimberCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_TimberCollection.DiscardUnknown(m)
}

var xxx_messageInfo_TimberCollection proto.InternalMessageInfo

func (m *TimberCollection) GetContext() *TimberContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *TimberCollection) GetItems() []*Timber {
	if m != nil {
		return m.Items
	}
	return nil
}

type TimberContext struct {
	KafkaTopic             string   `protobuf:"bytes,1,opt,name=kafka_topic,json=kafkaTopic,proto3" json:"kafka_topic,omitempty"`
	KafkaPartition         int32    `protobuf:"varint,2,opt,name=kafka_partition,json=kafkaPartition,proto3" json:"kafka_partition,omitempty"`
	KafkaReplicationFactor int32    `protobuf:"varint,3,opt,name=kafka_replication_factor,json=kafkaReplicationFactor,proto3" json:"kafka_replication_factor,omitempty"`
	EsIndexPrefix          string   `protobuf:"bytes,4,opt,name=es_index_prefix,json=esIndexPrefix,proto3" json:"es_index_prefix,omitempty"`
	EsDocumentType         string   `protobuf:"bytes,5,opt,name=es_document_type,json=esDocumentType,proto3" json:"es_document_type,omitempty"`
	AppMaxTps              int32    `protobuf:"varint,6,opt,name=app_max_tps,json=appMaxTps,proto3" json:"app_max_tps,omitempty"`
	AppSecret              string   `protobuf:"bytes,7,opt,name=app_secret,json=appSecret,proto3" json:"app_secret,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *TimberContext) Reset()         { *m = TimberContext{} }
func (m *TimberContext) String() string { return proto.CompactTextString(m) }
func (*TimberContext) ProtoMessage()    {}
func (*TimberContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_timber_04b8121c0f798b35, []int{2}
}
func (m *TimberContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimberContext.Unmarshal(m, b)
}
func (m *TimberContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimberContext.Marshal(b, m, deterministic)
}
func (dst *TimberContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimberContext.Merge(dst, src)
}
func (m *TimberContext) XXX_Size() int {
	return xxx_messageInfo_TimberContext.Size(m)
}
func (m *TimberContext) XXX_DiscardUnknown() {
	xxx_messageInfo_TimberContext.DiscardUnknown(m)
}

var xxx_messageInfo_TimberContext proto.InternalMessageInfo

func (m *TimberContext) GetKafkaTopic() string {
	if m != nil {
		return m.KafkaTopic
	}
	return ""
}

func (m *TimberContext) GetKafkaPartition() int32 {
	if m != nil {
		return m.KafkaPartition
	}
	return 0
}

func (m *TimberContext) GetKafkaReplicationFactor() int32 {
	if m != nil {
		return m.KafkaReplicationFactor
	}
	return 0
}

func (m *TimberContext) GetEsIndexPrefix() string {
	if m != nil {
		return m.EsIndexPrefix
	}
	return ""
}

func (m *TimberContext) GetEsDocumentType() string {
	if m != nil {
		return m.EsDocumentType
	}
	return ""
}

func (m *TimberContext) GetAppMaxTps() int32 {
	if m != nil {
		return m.AppMaxTps
	}
	return 0
}

func (m *TimberContext) GetAppSecret() string {
	if m != nil {
		return m.AppSecret
	}
	return ""
}

type ProduceResponse struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceResponse) Reset()         { *m = ProduceResponse{} }
func (m *ProduceResponse) String() string { return proto.CompactTextString(m) }
func (*ProduceResponse) ProtoMessage()    {}
func (*ProduceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_timber_04b8121c0f798b35, []int{3}
}
func (m *ProduceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceResponse.Unmarshal(m, b)
}
func (m *ProduceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceResponse.Marshal(b, m, deterministic)
}
func (dst *ProduceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceResponse.Merge(dst, src)
}
func (m *ProduceResponse) XXX_Size() int {
	return xxx_messageInfo_ProduceResponse.Size(m)
}
func (m *ProduceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceResponse proto.InternalMessageInfo

func (m *ProduceResponse) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*Timber)(nil), "timber.Timber")
	proto.RegisterMapType((map[string]*any.Any)(nil), "timber.Timber.ContentEntry")
	proto.RegisterType((*TimberCollection)(nil), "timber.TimberCollection")
	proto.RegisterType((*TimberContext)(nil), "timber.TimberContext")
	proto.RegisterType((*ProduceResponse)(nil), "timber.ProduceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BaritoProducerClient is the client API for BaritoProducer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BaritoProducerClient interface {
	Produce(ctx context.Context, in *Timber, opts ...grpc.CallOption) (*ProduceResponse, error)
	ProduceBatch(ctx context.Context, in *TimberCollection, opts ...grpc.CallOption) (*ProduceResponse, error)
}

type baritoProducerClient struct {
	cc *grpc.ClientConn
}

func NewBaritoProducerClient(cc *grpc.ClientConn) BaritoProducerClient {
	return &baritoProducerClient{cc}
}

func (c *baritoProducerClient) Produce(ctx context.Context, in *Timber, opts ...grpc.CallOption) (*ProduceResponse, error) {
	out := new(ProduceResponse)
	err := c.cc.Invoke(ctx, "/timber.BaritoProducer/Produce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baritoProducerClient) ProduceBatch(ctx context.Context, in *TimberCollection, opts ...grpc.CallOption) (*ProduceResponse, error) {
	out := new(ProduceResponse)
	err := c.cc.Invoke(ctx, "/timber.BaritoProducer/ProduceBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaritoProducerServer is the server API for BaritoProducer service.
type BaritoProducerServer interface {
	Produce(context.Context, *Timber) (*ProduceResponse, error)
	ProduceBatch(context.Context, *TimberCollection) (*ProduceResponse, error)
}

func RegisterBaritoProducerServer(s *grpc.Server, srv BaritoProducerServer) {
	s.RegisterService(&_BaritoProducer_serviceDesc, srv)
}

func _BaritoProducer_Produce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaritoProducerServer).Produce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timber.BaritoProducer/Produce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaritoProducerServer).Produce(ctx, req.(*Timber))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaritoProducer_ProduceBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimberCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaritoProducerServer).ProduceBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timber.BaritoProducer/ProduceBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaritoProducerServer).ProduceBatch(ctx, req.(*TimberCollection))
	}
	return interceptor(ctx, in, info, handler)
}

var _BaritoProducer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "timber.BaritoProducer",
	HandlerType: (*BaritoProducerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Produce",
			Handler:    _BaritoProducer_Produce_Handler,
		},
		{
			MethodName: "ProduceBatch",
			Handler:    _BaritoProducer_ProduceBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/timber.proto",
}

func init() { proto.RegisterFile("proto/timber.proto", fileDescriptor_timber_04b8121c0f798b35) }

var fileDescriptor_timber_04b8121c0f798b35 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0x49, 0xd7, 0xb4, 0xec, 0xe9, 0x6e, 0x5b, 0x86, 0x55, 0xc7, 0x8a, 0xba, 0x14, 0x71,
	0x8b, 0x17, 0x29, 0x54, 0x84, 0xc5, 0xbb, 0xdd, 0x55, 0xc1, 0x0b, 0xa1, 0xc4, 0xde, 0x87, 0x69,
	0x7a, 0xba, 0x0e, 0x4d, 0x67, 0x86, 0x99, 0xa9, 0x24, 0x4f, 0xe0, 0x8b, 0xf8, 0x24, 0x3e, 0x99,
	0xe4, 0x4c, 0x82, 0xdb, 0x82, 0x17, 0xde, 0xe5, 0xfc, 0xff, 0x77, 0xce, 0xfc, 0x27, 0x33, 0xc0,
	0x8c, 0xd5, 0x5e, 0xcf, 0xbc, 0xdc, 0xad, 0xd0, 0x26, 0x54, 0xb0, 0x6e, 0xa8, 0xc6, 0xcf, 0xee,
	0xb5, 0xbe, 0x2f, 0x70, 0x46, 0xea, 0x6a, 0xbf, 0x99, 0x09, 0x55, 0x05, 0x64, 0xf2, 0x3b, 0x82,
	0xee, 0x92, 0x28, 0x36, 0x83, 0x5e, 0xae, 0x95, 0xc7, 0xd2, 0xf3, 0xe8, 0x32, 0x9a, 0xf6, 0xe7,
	0x8f, 0x93, 0x66, 0x5a, 0x00, 0xee, 0x82, 0x99, 0xb6, 0x14, 0x7b, 0xdf, 0x34, 0x28, 0xcf, 0x3b,
	0x97, 0x27, 0xd3, 0xfe, 0xfc, 0xf9, 0x61, 0x43, 0x72, 0x17, 0xdc, 0x4f, 0xca, 0xdb, 0x2a, 0x6d,
	0xd9, 0xf1, 0x02, 0xce, 0x1e, 0x1a, 0x6c, 0x04, 0x27, 0x5b, 0xac, 0xe8, 0xcc, 0xd3, 0xb4, 0xfe,
	0x64, 0x6f, 0x21, 0xfe, 0x21, 0x8a, 0x3d, 0xf2, 0x0e, 0xe5, 0xb8, 0x48, 0x42, 0xfe, 0xa4, 0xcd,
	0x9f, 0xdc, 0xa8, 0x2a, 0x0d, 0xc8, 0x87, 0xce, 0x75, 0x34, 0x91, 0x30, 0x6a, 0x23, 0x16, 0x05,
	0xe6, 0x5e, 0x6a, 0xf5, 0xff, 0xdb, 0xbc, 0x86, 0x58, 0x7a, 0xdc, 0xb9, 0x66, 0x97, 0xc1, 0x21,
	0x9e, 0x06, 0x73, 0xf2, 0xab, 0x03, 0xe7, 0x07, 0x03, 0xd8, 0x2b, 0xe8, 0x6f, 0xc5, 0x66, 0x2b,
	0x32, 0xaf, 0x8d, 0xcc, 0x9b, 0x35, 0x80, 0xa4, 0x65, 0xad, 0xb0, 0x2b, 0x18, 0x06, 0xc0, 0x08,
	0xeb, 0x65, 0x1d, 0x8e, 0xf6, 0x8a, 0xd3, 0x01, 0xc9, 0x8b, 0x56, 0x65, 0xd7, 0xc0, 0x03, 0x68,
	0xd1, 0x14, 0x32, 0x17, 0xb5, 0x98, 0x6d, 0x44, 0xee, 0xb5, 0xe5, 0x27, 0xd4, 0xf1, 0x84, 0xfc,
	0xf4, 0xaf, 0xfd, 0x99, 0x5c, 0xf6, 0x06, 0x86, 0xe8, 0x32, 0xa9, 0xd6, 0x58, 0x66, 0xc6, 0xe2,
	0x46, 0x96, 0xfc, 0x11, 0xe5, 0x38, 0x47, 0xf7, 0xa5, 0x56, 0x17, 0x24, 0xb2, 0x29, 0x8c, 0xd0,
	0x65, 0x6b, 0x9d, 0xef, 0x77, 0xa8, 0x7c, 0xe6, 0x2b, 0x83, 0x3c, 0x26, 0x70, 0x80, 0xee, 0x63,
	0x23, 0x2f, 0x2b, 0x83, 0xec, 0x25, 0xf4, 0x85, 0x31, 0xd9, 0x4e, 0x94, 0x99, 0x37, 0x8e, 0x77,
	0xe9, 0xf8, 0x53, 0x61, 0xcc, 0x57, 0x51, 0x2e, 0x8d, 0x63, 0x2f, 0x00, 0x6a, 0xdf, 0x61, 0x6e,
	0xd1, 0xf3, 0x1e, 0xcd, 0xa8, 0xed, 0x6f, 0x24, 0x4c, 0xae, 0x60, 0xb8, 0xb0, 0x7a, 0xbd, 0xcf,
	0x31, 0x45, 0x67, 0xb4, 0x72, 0xc8, 0x2e, 0x20, 0x7e, 0xf8, 0x87, 0x42, 0x31, 0xff, 0x19, 0xc1,
	0xe0, 0x56, 0x58, 0xe9, 0x75, 0xc3, 0x5b, 0x36, 0x87, 0x5e, 0xf3, 0xcd, 0x8e, 0x2e, 0x61, 0xfc,
	0xb4, 0xad, 0x8f, 0x87, 0xdf, 0xc0, 0x59, 0x23, 0xdd, 0x0a, 0x9f, 0x7f, 0x67, 0xfc, 0xf8, 0xb2,
	0xdb, 0x77, 0xf1, 0xcf, 0x11, 0xab, 0x2e, 0xbd, 0xae, 0x77, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x6a, 0x8e, 0xc4, 0x1f, 0x49, 0x03, 0x00, 0x00,
}