// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/v1/rate/rate.proto

/*
Package rate is a generated protocol buffer package.

It is generated from these files:
	pb/v1/rate/rate.proto

It has these top-level messages:
	Request
	Result
	RatePlan
	RoomType
*/
package rate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Request struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
	InDate   string   `protobuf:"bytes,2,opt,name=inDate" json:"inDate,omitempty"`
	OutDate  string   `protobuf:"bytes,3,opt,name=outDate" json:"outDate,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func (m *Request) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *Request) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

type Result struct {
	RatePlans []*RatePlan `protobuf:"bytes,1,rep,name=ratePlans" json:"ratePlans,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetRatePlans() []*RatePlan {
	if m != nil {
		return m.RatePlans
	}
	return nil
}

type RatePlan struct {
	HotelId  string    `protobuf:"bytes,1,opt,name=hotelId" json:"hotelId,omitempty"`
	Code     string    `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	InDate   string    `protobuf:"bytes,3,opt,name=inDate" json:"inDate,omitempty"`
	OutDate  string    `protobuf:"bytes,4,opt,name=outDate" json:"outDate,omitempty"`
	RoomType *RoomType `protobuf:"bytes,5,opt,name=roomType" json:"roomType,omitempty"`
}

func (m *RatePlan) Reset()                    { *m = RatePlan{} }
func (m *RatePlan) String() string            { return proto.CompactTextString(m) }
func (*RatePlan) ProtoMessage()               {}
func (*RatePlan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RatePlan) GetHotelId() string {
	if m != nil {
		return m.HotelId
	}
	return ""
}

func (m *RatePlan) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RatePlan) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *RatePlan) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

func (m *RatePlan) GetRoomType() *RoomType {
	if m != nil {
		return m.RoomType
	}
	return nil
}

type RoomType struct {
	BookableRate       float64 `protobuf:"fixed64,1,opt,name=bookableRate" json:"bookableRate,omitempty"`
	TotalRate          float64 `protobuf:"fixed64,2,opt,name=totalRate" json:"totalRate,omitempty"`
	TotalRateInclusive float64 `protobuf:"fixed64,3,opt,name=totalRateInclusive" json:"totalRateInclusive,omitempty"`
	Code               string  `protobuf:"bytes,4,opt,name=code" json:"code,omitempty"`
	Currency           string  `protobuf:"bytes,5,opt,name=currency" json:"currency,omitempty"`
	RoomDescription    string  `protobuf:"bytes,6,opt,name=roomDescription" json:"roomDescription,omitempty"`
}

func (m *RoomType) Reset()                    { *m = RoomType{} }
func (m *RoomType) String() string            { return proto.CompactTextString(m) }
func (*RoomType) ProtoMessage()               {}
func (*RoomType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RoomType) GetBookableRate() float64 {
	if m != nil {
		return m.BookableRate
	}
	return 0
}

func (m *RoomType) GetTotalRate() float64 {
	if m != nil {
		return m.TotalRate
	}
	return 0
}

func (m *RoomType) GetTotalRateInclusive() float64 {
	if m != nil {
		return m.TotalRateInclusive
	}
	return 0
}

func (m *RoomType) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RoomType) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *RoomType) GetRoomDescription() string {
	if m != nil {
		return m.RoomDescription
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "rate.Request")
	proto.RegisterType((*Result)(nil), "rate.Result")
	proto.RegisterType((*RatePlan)(nil), "rate.RatePlan")
	proto.RegisterType((*RoomType)(nil), "rate.RoomType")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Rate service

type RateClient interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type rateClient struct {
	cc *grpc.ClientConn
}

func NewRateClient(cc *grpc.ClientConn) RateClient {
	return &rateClient{cc}
}

func (c *rateClient) GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/rate.Rate/GetRates", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rate service

type RateServer interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(context.Context, *Request) (*Result, error)
}

func RegisterRateServer(s *grpc.Server, srv RateServer) {
	s.RegisterService(&_Rate_serviceDesc, srv)
}

func _Rate_GetRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServer).GetRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rate.Rate/GetRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServer).GetRates(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rate.Rate",
	HandlerType: (*RateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRates",
			Handler:    _Rate_GetRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/v1/rate/rate.proto",
}

func init() { proto.RegisterFile("pb/v1/rate/rate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdf, 0x4a, 0xfb, 0x30,
	0x14, 0xc7, 0xc9, 0x6f, 0xfd, 0x75, 0xed, 0x71, 0x2a, 0x1c, 0x50, 0xca, 0xf0, 0x62, 0xf4, 0xc6,
	0x21, 0xb2, 0xe1, 0x04, 0x9f, 0x60, 0x20, 0xbb, 0x93, 0x83, 0xe0, 0x75, 0xdb, 0x05, 0x2c, 0xc6,
	0xa6, 0x26, 0xe9, 0x60, 0x2f, 0xe2, 0xa3, 0xf9, 0x3c, 0x92, 0x34, 0xed, 0x3a, 0xd1, 0x9b, 0x72,
	0x3e, 0xdf, 0x13, 0x4e, 0x3f, 0xf9, 0x03, 0x17, 0x75, 0xbe, 0xdc, 0xdd, 0x2d, 0x55, 0x66, 0xb8,
	0xfb, 0x2c, 0x6a, 0x25, 0x8d, 0xc4, 0xc0, 0xd6, 0xe9, 0x0b, 0x8c, 0x89, 0x7f, 0x34, 0x5c, 0x1b,
	0x9c, 0x42, 0xf4, 0x2a, 0x0d, 0x17, 0x9b, 0xad, 0x4e, 0xd8, 0x6c, 0x34, 0x8f, 0xa9, 0x67, 0xbc,
	0x84, 0xb0, 0xac, 0xd6, 0x99, 0xe1, 0xc9, 0xbf, 0x19, 0x9b, 0xc7, 0xe4, 0x09, 0x13, 0x18, 0xcb,
	0xc6, 0xb8, 0xc6, 0xc8, 0x35, 0x3a, 0x4c, 0x1f, 0x20, 0x24, 0xae, 0x1b, 0x61, 0xf0, 0x16, 0x62,
	0xfb, 0xab, 0x27, 0x91, 0x55, 0xed, 0xe0, 0x93, 0xd5, 0xd9, 0xc2, 0x89, 0x90, 0x8f, 0xe9, 0xb0,
	0x20, 0xfd, 0x64, 0x10, 0x75, 0xb9, 0x1d, 0xef, 0x15, 0x12, 0xd6, 0x8e, 0xf7, 0x88, 0x08, 0x41,
	0x21, 0xb7, 0x9d, 0x8e, 0xab, 0x07, 0x92, 0xa3, 0xbf, 0x24, 0x83, 0x23, 0x49, 0xbc, 0x81, 0x48,
	0x49, 0xf9, 0xfe, 0xbc, 0xaf, 0x79, 0xf2, 0x7f, 0xc6, 0x06, 0x66, 0x3e, 0xa5, 0xbe, 0x9f, 0x7e,
	0x59, 0x31, 0x0f, 0x98, 0xc2, 0x24, 0x97, 0xf2, 0x2d, 0xcb, 0x05, 0xb7, 0xb2, 0xce, 0x8e, 0xd1,
	0x51, 0x86, 0x57, 0x10, 0x1b, 0x69, 0x32, 0x41, 0xdd, 0xb1, 0x31, 0x3a, 0x04, 0xb8, 0x00, 0xec,
	0x61, 0x53, 0x15, 0xa2, 0xd1, 0xe5, 0xae, 0x15, 0x67, 0xf4, 0x4b, 0xa7, 0xdf, 0x70, 0x30, 0xd8,
	0xf0, 0x14, 0xa2, 0xa2, 0x51, 0x8a, 0x57, 0xc5, 0xde, 0xe9, 0xc7, 0xd4, 0x33, 0xce, 0xe1, 0xdc,
	0xaa, 0xaf, 0xb9, 0x2e, 0x54, 0x59, 0x9b, 0x52, 0x56, 0x49, 0xe8, 0x96, 0xfc, 0x8c, 0x57, 0x4b,
	0x08, 0x9c, 0xd1, 0x35, 0x44, 0x8f, 0xdc, 0xd8, 0x52, 0xe3, 0xa9, 0x3f, 0x86, 0xf6, 0x69, 0x4c,
	0x27, 0x1d, 0xda, 0x0b, 0xcd, 0x43, 0xf7, 0x80, 0xee, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x20,
	0xdb, 0x0d, 0x54, 0x59, 0x02, 0x00, 0x00,
}