// Code generated by protoc-gen-gogo.
// source: protocol.proto
// DO NOT EDIT!

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		protocol.proto

	It has these top-level messages:
		PubMsg
		Ask
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "gogoproto"

import strings "strings"
import reflect "reflect"

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

// How messages are delivered to the STAN cluster
type PubMsg struct {
	ClientID string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	Guid     string `protobuf:"bytes,2,opt,name=guid,proto3" json:"guid,omitempty"`
	Subject  string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Reply    string `protobuf:"bytes,4,opt,name=reply,proto3" json:"reply,omitempty"`
	Data     []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Sha256   []byte `protobuf:"bytes,10,opt,name=sha256,proto3" json:"sha256,omitempty"`
}

func (m *PubMsg) Reset()                    { *m = PubMsg{} }
func (m *PubMsg) String() string            { return proto.CompactTextString(m) }
func (*PubMsg) ProtoMessage()               {}
func (*PubMsg) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{0} }

type Ask struct {
	A string `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32  `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (m *Ask) Reset()                    { *m = Ask{} }
func (m *Ask) String() string            { return proto.CompactTextString(m) }
func (*Ask) ProtoMessage()               {}
func (*Ask) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{1} }

func init() {
	proto.RegisterType((*PubMsg)(nil), "pb.PubMsg")
	proto.RegisterType((*Ask)(nil), "pb.Ask")
}
func (this *PubMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&pb.PubMsg{")
	s = append(s, "ClientID: "+fmt.Sprintf("%#v", this.ClientID)+",\n")
	s = append(s, "Guid: "+fmt.Sprintf("%#v", this.Guid)+",\n")
	s = append(s, "Subject: "+fmt.Sprintf("%#v", this.Subject)+",\n")
	s = append(s, "Reply: "+fmt.Sprintf("%#v", this.Reply)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "Sha256: "+fmt.Sprintf("%#v", this.Sha256)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Ask) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.Ask{")
	s = append(s, "A: "+fmt.Sprintf("%#v", this.A)+",\n")
	s = append(s, "B: "+fmt.Sprintf("%#v", this.B)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringProtocol(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *PubMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ClientID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.ClientID)))
		i += copy(dAtA[i:], m.ClientID)
	}
	if len(m.Guid) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Guid)))
		i += copy(dAtA[i:], m.Guid)
	}
	if len(m.Subject) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Subject)))
		i += copy(dAtA[i:], m.Subject)
	}
	if len(m.Reply) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Reply)))
		i += copy(dAtA[i:], m.Reply)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if len(m.Sha256) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Sha256)))
		i += copy(dAtA[i:], m.Sha256)
	}
	return i, nil
}

func (m *Ask) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ask) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.A) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.A)))
		i += copy(dAtA[i:], m.A)
	}
	if m.B != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.B))
	}
	return i, nil
}

func encodeFixed64Protocol(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Protocol(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintProtocol(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PubMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Guid)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Reply)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.Sha256)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	return n
}

func (m *Ask) Size() (n int) {
	var l int
	_ = l
	l = len(m.A)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.B != 0 {
		n += 1 + sovProtocol(uint64(m.B))
	}
	return n
}

func sovProtocol(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtocol(x uint64) (n int) {
	return sovProtocol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PubMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PubMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Guid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Guid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reply = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sha256", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sha256 = append(m.Sha256[:0], dAtA[iNdEx:postIndex]...)
			if m.Sha256 == nil {
				m.Sha256 = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
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
func (m *Ask) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ask: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ask: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			m.B = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.B |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
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
func skipProtocol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtocol
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
					return 0, ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowProtocol
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProtocol
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtocol
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipProtocol(dAtA[start:])
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
	ErrInvalidLengthProtocol = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtocol   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protocol.proto", fileDescriptorProtocol) }

var fileDescriptorProtocol = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x03, 0x33, 0x84, 0x98, 0x0a, 0x92, 0xa4, 0x44, 0xd2, 0xf3, 0xd3,
	0xf3, 0xc1, 0x5c, 0x7d, 0x10, 0x0b, 0x22, 0xa3, 0x34, 0x85, 0x91, 0x8b, 0x2d, 0xa0, 0x34, 0xc9,
	0xb7, 0x38, 0x5d, 0x48, 0x8a, 0x8b, 0x23, 0x39, 0x27, 0x33, 0x35, 0xaf, 0xc4, 0xd3, 0x45, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xce, 0x17, 0x12, 0xe2, 0x62, 0x49, 0x2f, 0xcd, 0x4c, 0x91,
	0x60, 0x02, 0x8b, 0x83, 0xd9, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa5, 0x49, 0x59, 0xa9, 0xc9, 0x25,
	0x12, 0xcc, 0x60, 0x61, 0x18, 0x57, 0x48, 0x84, 0x8b, 0xb5, 0x28, 0xb5, 0x20, 0xa7, 0x52, 0x82,
	0x05, 0x2c, 0x0e, 0xe1, 0x80, 0xcc, 0x48, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x55, 0x60, 0xd4, 0xe0,
	0x09, 0x02, 0xb3, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x33, 0x12, 0x8d, 0x4c, 0xcd, 0x24, 0xb8, 0xc0,
	0xa2, 0x50, 0x9e, 0x92, 0x22, 0x17, 0xb3, 0x63, 0x71, 0xb6, 0x10, 0x0f, 0x17, 0x63, 0x22, 0xd4,
	0x2d, 0x8c, 0x89, 0x20, 0x5e, 0x12, 0xd8, 0x05, 0xac, 0x41, 0x8c, 0x49, 0x4e, 0x12, 0x27, 0x1e,
	0xca, 0x31, 0x7c, 0x78, 0x28, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0xf6, 0x9a, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x16, 0x42, 0x80, 0xaa, 0x06, 0x01, 0x00, 0x00,
}
