// Code generated by protoc-gen-gogo.
// source: slice.proto
// DO NOT EDIT!

/*
	Package events is a generated protocol buffer package.

	It is generated from these files:
		slice.proto

	It has these top-level messages:
		SessionSlice
		SessionChunk
*/
package events

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SessionSlice is a slice of submitted chunks
type SessionSlice struct {
	Namespace string          `protobuf:"bytes,1,opt,name=Namespace,json=namespace,proto3" json:"Namespace,omitempty"`
	SessionID string          `protobuf:"bytes,2,opt,name=SessionID,json=sessionID,proto3" json:"SessionID,omitempty"`
	Chunks    []*SessionChunk `protobuf:"bytes,3,rep,name=Chunks,json=chunks" json:"Chunks,omitempty"`
}

func (m *SessionSlice) Reset()                    { *m = SessionSlice{} }
func (m *SessionSlice) String() string            { return proto.CompactTextString(m) }
func (*SessionSlice) ProtoMessage()               {}
func (*SessionSlice) Descriptor() ([]byte, []int) { return fileDescriptorSlice, []int{0} }

func (m *SessionSlice) GetChunks() []*SessionChunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

// SessionChunk is a chunk to be posted in the context of the session
type SessionChunk struct {
	// Time is the occurence of this event
	Time int64 `protobuf:"varint,2,opt,name=Time,json=time,proto3" json:"Time,omitempty"`
	// Data is captured data
	Data []byte `protobuf:"bytes,3,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
}

func (m *SessionChunk) Reset()                    { *m = SessionChunk{} }
func (m *SessionChunk) String() string            { return proto.CompactTextString(m) }
func (*SessionChunk) ProtoMessage()               {}
func (*SessionChunk) Descriptor() ([]byte, []int) { return fileDescriptorSlice, []int{1} }

func init() {
	proto.RegisterType((*SessionSlice)(nil), "events.SessionSlice")
	proto.RegisterType((*SessionChunk)(nil), "events.SessionChunk")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for AuditLog service

type AuditLogClient interface {
	SubmitSessionSlice(ctx context.Context, opts ...grpc.CallOption) (AuditLog_SubmitSessionSliceClient, error)
}

type auditLogClient struct {
	cc *grpc.ClientConn
}

func NewAuditLogClient(cc *grpc.ClientConn) AuditLogClient {
	return &auditLogClient{cc}
}

func (c *auditLogClient) SubmitSessionSlice(ctx context.Context, opts ...grpc.CallOption) (AuditLog_SubmitSessionSliceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AuditLog_serviceDesc.Streams[0], c.cc, "/events.AuditLog/SubmitSessionSlice", opts...)
	if err != nil {
		return nil, err
	}
	x := &auditLogSubmitSessionSliceClient{stream}
	return x, nil
}

type AuditLog_SubmitSessionSliceClient interface {
	Send(*SessionSlice) error
	CloseAndRecv() (*google_protobuf1.Empty, error)
	grpc.ClientStream
}

type auditLogSubmitSessionSliceClient struct {
	grpc.ClientStream
}

func (x *auditLogSubmitSessionSliceClient) Send(m *SessionSlice) error {
	return x.ClientStream.SendMsg(m)
}

func (x *auditLogSubmitSessionSliceClient) CloseAndRecv() (*google_protobuf1.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(google_protobuf1.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AuditLog service

type AuditLogServer interface {
	SubmitSessionSlice(AuditLog_SubmitSessionSliceServer) error
}

func RegisterAuditLogServer(s *grpc.Server, srv AuditLogServer) {
	s.RegisterService(&_AuditLog_serviceDesc, srv)
}

func _AuditLog_SubmitSessionSlice_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AuditLogServer).SubmitSessionSlice(&auditLogSubmitSessionSliceServer{stream})
}

type AuditLog_SubmitSessionSliceServer interface {
	SendAndClose(*google_protobuf1.Empty) error
	Recv() (*SessionSlice, error)
	grpc.ServerStream
}

type auditLogSubmitSessionSliceServer struct {
	grpc.ServerStream
}

func (x *auditLogSubmitSessionSliceServer) SendAndClose(m *google_protobuf1.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *auditLogSubmitSessionSliceServer) Recv() (*SessionSlice, error) {
	m := new(SessionSlice)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AuditLog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.AuditLog",
	HandlerType: (*AuditLogServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubmitSessionSlice",
			Handler:       _AuditLog_SubmitSessionSlice_Handler,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptorSlice,
}

func (m *SessionSlice) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SessionSlice) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintSlice(data, i, uint64(len(m.Namespace)))
		i += copy(data[i:], m.Namespace)
	}
	if len(m.SessionID) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintSlice(data, i, uint64(len(m.SessionID)))
		i += copy(data[i:], m.SessionID)
	}
	if len(m.Chunks) > 0 {
		for _, msg := range m.Chunks {
			data[i] = 0x1a
			i++
			i = encodeVarintSlice(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SessionChunk) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SessionChunk) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Time != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintSlice(data, i, uint64(m.Time))
	}
	if len(m.Data) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintSlice(data, i, uint64(len(m.Data)))
		i += copy(data[i:], m.Data)
	}
	return i, nil
}

func encodeFixed64Slice(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Slice(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSlice(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *SessionSlice) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovSlice(uint64(l))
	}
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovSlice(uint64(l))
	}
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.Size()
			n += 1 + l + sovSlice(uint64(l))
		}
	}
	return n
}

func (m *SessionChunk) Size() (n int) {
	var l int
	_ = l
	if m.Time != 0 {
		n += 1 + sovSlice(uint64(m.Time))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovSlice(uint64(l))
	}
	return n
}

func sovSlice(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSlice(x uint64) (n int) {
	return sovSlice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SessionSlice) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SessionSlice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionSlice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks, &SessionChunk{})
			if err := m.Chunks[len(m.Chunks)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlice(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlice
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
func (m *SessionChunk) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SessionChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], data[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlice(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlice
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
func skipSlice(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSlice
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSlice
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSlice
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSlice(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSlice = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSlice   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("slice.proto", fileDescriptorSlice) }

var fileDescriptorSlice = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x17, 0x3b, 0x8a, 0xcb, 0x76, 0x90, 0x20, 0x52, 0xa6, 0x94, 0xb1, 0x53, 0x0f, 0x9a,
	0xc2, 0x04, 0xef, 0xea, 0x14, 0x04, 0xf1, 0xd0, 0x79, 0xf2, 0x22, 0x69, 0xf7, 0xd9, 0x05, 0xd7,
	0x24, 0x34, 0x5f, 0xd5, 0xfe, 0x13, 0x7f, 0x92, 0x47, 0x7f, 0x82, 0xcc, 0x3f, 0x22, 0x4d, 0x3b,
	0x41, 0x6f, 0xf9, 0xde, 0xe7, 0x4d, 0xde, 0x7c, 0x2f, 0x1d, 0xda, 0xb5, 0xcc, 0x80, 0x9b, 0x52,
	0xa3, 0x66, 0x3e, 0xbc, 0x80, 0x42, 0x3b, 0x7e, 0xc8, 0x25, 0xae, 0xaa, 0x94, 0x67, 0xba, 0x88,
	0xf3, 0xd2, 0x64, 0x27, 0x90, 0x69, 0x5b, 0x5b, 0x84, 0x6e, 0xcc, 0x05, 0xc2, 0xab, 0xa8, 0x63,
	0x5c, 0xc9, 0x72, 0xf9, 0x68, 0x44, 0x89, 0x75, 0x9c, 0x6b, 0x9d, 0xaf, 0x41, 0x18, 0x69, 0xbb,
	0x63, 0x2c, 0x8c, 0x8c, 0x85, 0x52, 0x1a, 0x05, 0x4a, 0xad, 0x6c, 0x9b, 0x31, 0x3e, 0xec, 0xa8,
	0x9b, 0xd2, 0xea, 0x29, 0x86, 0xc2, 0x60, 0xdd, 0xc2, 0xe9, 0x1b, 0x1d, 0x2d, 0xc0, 0x5a, 0xa9,
	0xd5, 0xa2, 0xf9, 0x16, 0x3b, 0xa2, 0x83, 0x3b, 0x51, 0x80, 0x35, 0x22, 0x83, 0x80, 0x4c, 0x48,
	0x34, 0x48, 0x06, 0x6a, 0x2b, 0x34, 0xb4, 0x73, 0xdf, 0xcc, 0x83, 0x9d, 0x96, 0xda, 0xad, 0xc0,
	0x8e, 0xa9, 0x7f, 0xb9, 0xaa, 0xd4, 0xb3, 0x0d, 0xbc, 0x89, 0x17, 0x0d, 0x67, 0xfb, 0xbc, 0xdd,
	0x8e, 0x77, 0x77, 0x1c, 0x4c, 0xfc, 0xcc, 0x79, 0xa6, 0x67, 0xbf, 0xc9, 0x4e, 0x67, 0x8c, 0xf6,
	0xef, 0x65, 0x01, 0xee, 0x59, 0x2f, 0xe9, 0xa3, 0x2c, 0xa0, 0xd1, 0xe6, 0x02, 0x45, 0xe0, 0x4d,
	0x48, 0x34, 0x4a, 0xfa, 0x4b, 0x81, 0x62, 0x96, 0xd0, 0xdd, 0xf3, 0x6a, 0x29, 0xf1, 0x56, 0xe7,
	0xec, 0x9a, 0xb2, 0x45, 0x95, 0x16, 0x12, 0xff, 0xec, 0xf0, 0x3f, 0xd7, 0xa9, 0xe3, 0x03, 0xde,
	0xf6, 0xc0, 0xb7, 0x3d, 0xf0, 0xab, 0xa6, 0x87, 0x69, 0x2f, 0x22, 0x17, 0x7b, 0x1f, 0x9b, 0x90,
	0x7c, 0x6e, 0x42, 0xf2, 0xb5, 0x09, 0xc9, 0xfb, 0x77, 0xd8, 0x4b, 0x7d, 0xe7, 0x3a, 0xfd, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x21, 0xc6, 0xeb, 0xd4, 0xae, 0x01, 0x00, 0x00,
}
