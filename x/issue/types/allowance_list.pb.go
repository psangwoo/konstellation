// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue/allowance_list.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type AllowanceList struct {
	Allowances []*Allowance `protobuf:"bytes,1,rep,name=allowances,proto3" json:"allowances,omitempty"`
	Id         uint64       `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *AllowanceList) Reset()         { *m = AllowanceList{} }
func (m *AllowanceList) String() string { return proto.CompactTextString(m) }
func (*AllowanceList) ProtoMessage()    {}
func (*AllowanceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e4061c655792577, []int{0}
}
func (m *AllowanceList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowanceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowanceList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowanceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowanceList.Merge(m, src)
}
func (m *AllowanceList) XXX_Size() int {
	return m.Size()
}
func (m *AllowanceList) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowanceList.DiscardUnknown(m)
}

var xxx_messageInfo_AllowanceList proto.InternalMessageInfo

func (m *AllowanceList) GetAllowances() []*Allowance {
	if m != nil {
		return m.Allowances
	}
	return nil
}

func (m *AllowanceList) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*AllowanceList)(nil), "konstellation.issue.AllowanceList")
}

func init() { proto.RegisterFile("issue/allowance_list.proto", fileDescriptor_8e4061c655792577) }

var fileDescriptor_8e4061c655792577 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x2c, 0x2e, 0x2e,
	0x4d, 0xd5, 0x4f, 0xcc, 0xc9, 0xc9, 0x2f, 0x4f, 0xcc, 0x4b, 0x4e, 0x8d, 0xcf, 0xc9, 0x2c, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xce, 0xce, 0xcf, 0x2b, 0x2e, 0x49, 0xcd, 0xc9,
	0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0xab, 0x94, 0x12, 0x45, 0xd3, 0x00, 0x51, 0xab, 0x14,
	0xcf, 0xc5, 0xeb, 0x08, 0x13, 0xf2, 0xc9, 0x2c, 0x2e, 0x11, 0xb2, 0xe3, 0xe2, 0x82, 0xab, 0x29,
	0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd3, 0xc3, 0x62, 0xa2, 0x1e, 0x5c, 0x5f, 0x10,
	0x92, 0x0e, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6,
	0xcc, 0x14, 0x27, 0xf5, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xe2, 0xad,
	0xd0, 0x87, 0xb8, 0xa9, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x20, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x4c, 0x1e, 0xd1, 0x4e, 0xda, 0x00, 0x00, 0x00,
}

func (m *AllowanceList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowanceList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowanceList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintAllowanceList(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Allowances) > 0 {
		for iNdEx := len(m.Allowances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Allowances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAllowanceList(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAllowanceList(dAtA []byte, offset int, v uint64) int {
	offset -= sovAllowanceList(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AllowanceList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Allowances) > 0 {
		for _, e := range m.Allowances {
			l = e.Size()
			n += 1 + l + sovAllowanceList(uint64(l))
		}
	}
	if m.Id != 0 {
		n += 1 + sovAllowanceList(uint64(m.Id))
	}
	return n
}

func sovAllowanceList(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAllowanceList(x uint64) (n int) {
	return sovAllowanceList(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AllowanceList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllowanceList
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
			return fmt.Errorf("proto: AllowanceList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowanceList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllowanceList
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
				return ErrInvalidLengthAllowanceList
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAllowanceList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Allowances = append(m.Allowances, &Allowance{})
			if err := m.Allowances[len(m.Allowances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllowanceList
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAllowanceList(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAllowanceList
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
func skipAllowanceList(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAllowanceList
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
					return 0, ErrIntOverflowAllowanceList
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
					return 0, ErrIntOverflowAllowanceList
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
				return 0, ErrInvalidLengthAllowanceList
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAllowanceList
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAllowanceList
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAllowanceList        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAllowanceList          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAllowanceList = fmt.Errorf("proto: unexpected end of group")
)
