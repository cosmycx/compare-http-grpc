// Code generated by protoc-gen-go. DO NOT EDIT.
// source: validate.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type RowData struct {
	Field                []string `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RowData) Reset()         { *m = RowData{} }
func (m *RowData) String() string { return proto.CompactTextString(m) }
func (*RowData) ProtoMessage()    {}
func (*RowData) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ce066df60f429f, []int{0}
}

func (m *RowData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RowData.Unmarshal(m, b)
}
func (m *RowData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RowData.Marshal(b, m, deterministic)
}
func (m *RowData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RowData.Merge(m, src)
}
func (m *RowData) XXX_Size() int {
	return xxx_messageInfo_RowData.Size(m)
}
func (m *RowData) XXX_DiscardUnknown() {
	xxx_messageInfo_RowData.DiscardUnknown(m)
}

var xxx_messageInfo_RowData proto.InternalMessageInfo

func (m *RowData) GetField() []string {
	if m != nil {
		return m.Field
	}
	return nil
}

type Row struct {
	RowNo                int64    `protobuf:"varint,1,opt,name=rowNo,proto3" json:"rowNo,omitempty"`
	RowData              *RowData `protobuf:"bytes,2,opt,name=rowData,proto3" json:"rowData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ce066df60f429f, []int{1}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetRowNo() int64 {
	if m != nil {
		return m.RowNo
	}
	return 0
}

func (m *Row) GetRowData() *RowData {
	if m != nil {
		return m.RowData
	}
	return nil
}

type ValidationPBRequest struct {
	ParticipID           string   `protobuf:"bytes,1,opt,name=participID,proto3" json:"participID,omitempty"`
	Datarows             []*Row   `protobuf:"bytes,2,rep,name=datarows,proto3" json:"datarows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationPBRequest) Reset()         { *m = ValidationPBRequest{} }
func (m *ValidationPBRequest) String() string { return proto.CompactTextString(m) }
func (*ValidationPBRequest) ProtoMessage()    {}
func (*ValidationPBRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ce066df60f429f, []int{2}
}

func (m *ValidationPBRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationPBRequest.Unmarshal(m, b)
}
func (m *ValidationPBRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationPBRequest.Marshal(b, m, deterministic)
}
func (m *ValidationPBRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationPBRequest.Merge(m, src)
}
func (m *ValidationPBRequest) XXX_Size() int {
	return xxx_messageInfo_ValidationPBRequest.Size(m)
}
func (m *ValidationPBRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationPBRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationPBRequest proto.InternalMessageInfo

func (m *ValidationPBRequest) GetParticipID() string {
	if m != nil {
		return m.ParticipID
	}
	return ""
}

func (m *ValidationPBRequest) GetDatarows() []*Row {
	if m != nil {
		return m.Datarows
	}
	return nil
}

type ValidationPBResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidationPBResponse) Reset()         { *m = ValidationPBResponse{} }
func (m *ValidationPBResponse) String() string { return proto.CompactTextString(m) }
func (*ValidationPBResponse) ProtoMessage()    {}
func (*ValidationPBResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18ce066df60f429f, []int{3}
}

func (m *ValidationPBResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationPBResponse.Unmarshal(m, b)
}
func (m *ValidationPBResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationPBResponse.Marshal(b, m, deterministic)
}
func (m *ValidationPBResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationPBResponse.Merge(m, src)
}
func (m *ValidationPBResponse) XXX_Size() int {
	return xxx_messageInfo_ValidationPBResponse.Size(m)
}
func (m *ValidationPBResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationPBResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationPBResponse proto.InternalMessageInfo

func (m *ValidationPBResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*RowData)(nil), "main.RowData")
	proto.RegisterType((*Row)(nil), "main.Row")
	proto.RegisterType((*ValidationPBRequest)(nil), "main.ValidationPBRequest")
	proto.RegisterType((*ValidationPBResponse)(nil), "main.ValidationPBResponse")
}

func init() { proto.RegisterFile("validate.proto", fileDescriptor_18ce066df60f429f) }

var fileDescriptor_18ce066df60f429f = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xbf, 0x6d, 0x3e, 0xd3, 0x66, 0x8a, 0x82, 0x6b, 0x91, 0xd8, 0x83, 0x86, 0x05, 0x71,
	0x4f, 0x39, 0xc4, 0x7f, 0x10, 0xe2, 0xc1, 0x8b, 0x94, 0x15, 0x3c, 0x14, 0x2f, 0x6b, 0x33, 0xc2,
	0x42, 0xcc, 0xc4, 0xdd, 0x6d, 0xf3, 0xf7, 0xa5, 0xd9, 0xb4, 0x28, 0xf4, 0xf8, 0xce, 0xbb, 0xfb,
	0x3c, 0xc3, 0xc0, 0xc5, 0x4e, 0x37, 0xa6, 0xd6, 0x1e, 0xf3, 0xce, 0x92, 0x27, 0xfe, 0xff, 0x4b,
	0x9b, 0x56, 0xdc, 0xc1, 0x54, 0x51, 0x5f, 0x69, 0xaf, 0xf9, 0x02, 0xce, 0x3e, 0x0d, 0x36, 0x75,
	0xca, 0xb2, 0x48, 0x26, 0x2a, 0x04, 0x51, 0x41, 0xa4, 0xa8, 0xdf, 0x97, 0x96, 0xfa, 0x17, 0x4a,
	0x59, 0xc6, 0x64, 0xa4, 0x42, 0xe0, 0x0f, 0x30, 0xb5, 0xe1, 0x77, 0x3a, 0xc9, 0x98, 0x9c, 0x17,
	0xe7, 0xf9, 0x9e, 0x9a, 0x8f, 0x48, 0x75, 0x68, 0xc5, 0x3b, 0x5c, 0xbd, 0x05, 0xbd, 0xa1, 0x76,
	0x55, 0x2a, 0xfc, 0xde, 0xa2, 0xf3, 0xfc, 0x16, 0xa0, 0xd3, 0xd6, 0x9b, 0x8d, 0xe9, 0x9e, 0xab,
	0x01, 0x9d, 0xa8, 0x5f, 0x13, 0x7e, 0x0f, 0xb3, 0x5a, 0x7b, 0x6d, 0xa9, 0x77, 0xe9, 0x24, 0x8b,
	0xe4, 0xbc, 0x48, 0x8e, 0x02, 0x75, 0xac, 0x44, 0x0e, 0x8b, 0xbf, 0x74, 0xd7, 0x51, 0xeb, 0x90,
	0x5f, 0x43, 0x6c, 0xd1, 0x6d, 0x1b, 0x3f, 0xa2, 0xc7, 0x54, 0xac, 0xe1, 0x72, 0x7c, 0x8f, 0xab,
	0xf2, 0x15, 0xed, 0xce, 0x6c, 0x90, 0x3f, 0xc1, 0xec, 0x30, 0xe4, 0x37, 0xc1, 0x72, 0x62, 0xe5,
	0xe5, 0xf2, 0x54, 0x15, 0x7c, 0xe2, 0x9f, 0x64, 0x65, 0xbc, 0x1e, 0x0e, 0xfb, 0x11, 0x0f, 0x57,
	0x7e, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x00, 0xd3, 0xf8, 0x77, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ValidatePBServiceClient is the client API for ValidatePBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ValidatePBServiceClient interface {
	Validate(ctx context.Context, opts ...grpc.CallOption) (ValidatePBService_ValidateClient, error)
}

type validatePBServiceClient struct {
	cc *grpc.ClientConn
}

func NewValidatePBServiceClient(cc *grpc.ClientConn) ValidatePBServiceClient {
	return &validatePBServiceClient{cc}
}

func (c *validatePBServiceClient) Validate(ctx context.Context, opts ...grpc.CallOption) (ValidatePBService_ValidateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ValidatePBService_serviceDesc.Streams[0], "/main.ValidatePBService/Validate", opts...)
	if err != nil {
		return nil, err
	}
	x := &validatePBServiceValidateClient{stream}
	return x, nil
}

type ValidatePBService_ValidateClient interface {
	Send(*ValidationPBRequest) error
	CloseAndRecv() (*ValidationPBResponse, error)
	grpc.ClientStream
}

type validatePBServiceValidateClient struct {
	grpc.ClientStream
}

func (x *validatePBServiceValidateClient) Send(m *ValidationPBRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *validatePBServiceValidateClient) CloseAndRecv() (*ValidationPBResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ValidationPBResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ValidatePBServiceServer is the server API for ValidatePBService service.
type ValidatePBServiceServer interface {
	Validate(ValidatePBService_ValidateServer) error
}

func RegisterValidatePBServiceServer(s *grpc.Server, srv ValidatePBServiceServer) {
	s.RegisterService(&_ValidatePBService_serviceDesc, srv)
}

func _ValidatePBService_Validate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ValidatePBServiceServer).Validate(&validatePBServiceValidateServer{stream})
}

type ValidatePBService_ValidateServer interface {
	SendAndClose(*ValidationPBResponse) error
	Recv() (*ValidationPBRequest, error)
	grpc.ServerStream
}

type validatePBServiceValidateServer struct {
	grpc.ServerStream
}

func (x *validatePBServiceValidateServer) SendAndClose(m *ValidationPBResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *validatePBServiceValidateServer) Recv() (*ValidationPBRequest, error) {
	m := new(ValidationPBRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ValidatePBService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.ValidatePBService",
	HandlerType: (*ValidatePBServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Validate",
			Handler:       _ValidatePBService_Validate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "validate.proto",
}
