// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/stakeibc/validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Validator struct {
	Name                        string                                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address                     string                                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Weight                      uint64                                 `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	Delegation                  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=delegation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"delegation"`
	SlashQueryProgressTracker   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=slash_query_progress_tracker,json=slashQueryProgressTracker,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"slash_query_progress_tracker"`
	SlashQueryCheckpoint        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,12,opt,name=slash_query_checkpoint,json=slashQueryCheckpoint,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"slash_query_checkpoint"`
	SharesToTokensRate          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=shares_to_tokens_rate,json=sharesToTokensRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"shares_to_tokens_rate"`
	DelegationChangesInProgress int64                                  `protobuf:"varint,11,opt,name=delegation_changes_in_progress,json=delegationChangesInProgress,proto3" json:"delegation_changes_in_progress,omitempty"`
	SlashQueryInProgress        bool                                   `protobuf:"varint,13,opt,name=slash_query_in_progress,json=slashQueryInProgress,proto3" json:"slash_query_in_progress,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d2f32e16bd6ab8f, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Validator) GetWeight() uint64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Validator) GetDelegationChangesInProgress() int64 {
	if m != nil {
		return m.DelegationChangesInProgress
	}
	return 0
}

func (m *Validator) GetSlashQueryInProgress() bool {
	if m != nil {
		return m.SlashQueryInProgress
	}
	return false
}

func init() {
	proto.RegisterType((*Validator)(nil), "stride.stakeibc.Validator")
}

func init() { proto.RegisterFile("stride/stakeibc/validator.proto", fileDescriptor_5d2f32e16bd6ab8f) }

var fileDescriptor_5d2f32e16bd6ab8f = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x1b, 0x96, 0x75, 0xa9, 0x01, 0x51, 0x59, 0x65, 0x64, 0x03, 0xa5, 0x15, 0x07, 0xd4,
	0x4b, 0x13, 0xa9, 0x88, 0x1b, 0x17, 0xd6, 0x5d, 0x56, 0x4d, 0x08, 0xb2, 0x8a, 0x03, 0x97, 0xc8,
	0x75, 0x3e, 0x25, 0x56, 0x5a, 0xbb, 0xd8, 0xde, 0x60, 0x6f, 0xc1, 0x8d, 0x17, 0xd9, 0x43, 0xec,
	0x38, 0xed, 0x84, 0x38, 0x4c, 0xa8, 0x7d, 0x11, 0x14, 0x3b, 0x5d, 0x72, 0x65, 0xa7, 0xd8, 0x9f,
	0xff, 0xfe, 0xfd, 0xf3, 0xfd, 0x93, 0x0f, 0xf5, 0x95, 0x96, 0x2c, 0x85, 0x48, 0x69, 0x52, 0x00,
	0x9b, 0xd3, 0xe8, 0x82, 0x2c, 0x58, 0x4a, 0xb4, 0x90, 0xe1, 0x4a, 0x0a, 0x2d, 0xf0, 0x33, 0x2b,
	0x08, 0xb7, 0x82, 0xc3, 0x03, 0x2a, 0xd4, 0x52, 0xa8, 0xc4, 0x1c, 0x47, 0x76, 0x63, 0xb5, 0x87,
	0xbd, 0x4c, 0x64, 0xc2, 0xd6, 0xcb, 0x95, 0xad, 0xbe, 0xfe, 0xb5, 0x8b, 0x3a, 0x5f, 0xb6, 0x54,
	0x8c, 0x91, 0xcb, 0xc9, 0x12, 0x7c, 0x67, 0xe0, 0x0c, 0x3b, 0xb1, 0x59, 0xe3, 0x31, 0xda, 0x23,
	0x69, 0x2a, 0x41, 0x29, 0xff, 0x51, 0x59, 0x3e, 0xf2, 0x6f, 0xaf, 0x46, 0xbd, 0x0a, 0xfd, 0xc1,
	0x9e, 0x9c, 0x69, 0xc9, 0x78, 0x16, 0x6f, 0x85, 0x78, 0x1f, 0xb5, 0xbf, 0x03, 0xcb, 0x72, 0xed,
	0xb7, 0x07, 0xce, 0xd0, 0x8d, 0xab, 0x1d, 0xfe, 0x88, 0x50, 0x0a, 0x0b, 0xc8, 0x88, 0x66, 0x82,
	0xfb, 0xbb, 0x06, 0x17, 0x5e, 0xdf, 0xf5, 0x5b, 0x7f, 0xee, 0xfa, 0x6f, 0x32, 0xa6, 0xf3, 0xf3,
	0x79, 0x48, 0xc5, 0xb2, 0x7a, 0xf1, 0xea, 0x31, 0x52, 0x69, 0x11, 0xe9, 0xcb, 0x15, 0xa8, 0xf0,
	0x84, 0xeb, 0xb8, 0x41, 0xc0, 0x02, 0xbd, 0x52, 0x0b, 0xa2, 0xf2, 0xe4, 0xdb, 0x39, 0xc8, 0xcb,
	0xb2, 0xeb, 0xac, 0xf4, 0x4f, 0xb4, 0x24, 0xb4, 0x00, 0xe9, 0x77, 0x1e, 0xe4, 0x70, 0x60, 0x98,
	0x9f, 0x4b, 0xe4, 0xa7, 0x8a, 0x38, 0xb3, 0x40, 0x9c, 0xa2, 0xfd, 0xa6, 0x21, 0xcd, 0x81, 0x16,
	0x2b, 0xc1, 0xb8, 0xf6, 0x9f, 0x3c, 0xc8, 0xaa, 0x57, 0x5b, 0x4d, 0xee, 0x59, 0x58, 0xa0, 0xe7,
	0x2a, 0x27, 0x12, 0x54, 0xa2, 0x45, 0xa2, 0x45, 0x01, 0x5c, 0x25, 0x92, 0x68, 0xf0, 0x91, 0x31,
	0x79, 0xff, 0x1f, 0x26, 0xc7, 0x40, 0x6f, 0xaf, 0x46, 0xa8, 0xfa, 0x5c, 0xc7, 0x40, 0x63, 0x6c,
	0xd1, 0x33, 0x31, 0x33, 0xe0, 0x98, 0x68, 0xc0, 0x13, 0x14, 0xd4, 0xa9, 0x26, 0x34, 0x27, 0x3c,
	0x03, 0x95, 0x30, 0x7e, 0x9f, 0xa8, 0xff, 0x78, 0xe0, 0x0c, 0x77, 0xe2, 0x97, 0xb5, 0x6a, 0x62,
	0x45, 0x27, 0x7c, 0x1b, 0x11, 0x7e, 0x87, 0x5e, 0x34, 0xb3, 0x69, 0xde, 0x7e, 0x3a, 0x70, 0x86,
	0x5e, 0xb3, 0xd9, 0xfa, 0xda, 0xd4, 0xf5, 0x76, 0xba, 0xee, 0xd4, 0xf5, 0xdc, 0xee, 0xee, 0xd4,
	0xf5, 0xf6, 0xba, 0xde, 0xd4, 0xf5, 0xbc, 0x6e, 0xe7, 0xe8, 0xf4, 0x7a, 0x1d, 0x38, 0x37, 0xeb,
	0xc0, 0xf9, 0xbb, 0x0e, 0x9c, 0x9f, 0x9b, 0xa0, 0x75, 0xb3, 0x09, 0x5a, 0xbf, 0x37, 0x41, 0xeb,
	0xeb, 0xb8, 0xd1, 0xf7, 0x99, 0x19, 0x80, 0xd1, 0x29, 0x99, 0xab, 0xa8, 0x9a, 0x96, 0x8b, 0xf1,
	0x38, 0xfa, 0x51, 0xcf, 0x8c, 0xc9, 0x61, 0xde, 0x36, 0xbf, 0xfb, 0xdb, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x84, 0x3d, 0xc4, 0x6d, 0x53, 0x03, 0x00, 0x00,
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SlashQueryInProgress {
		i--
		if m.SlashQueryInProgress {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	{
		size := m.SlashQueryCheckpoint.Size()
		i -= size
		if _, err := m.SlashQueryCheckpoint.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if m.DelegationChangesInProgress != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.DelegationChangesInProgress))
		i--
		dAtA[i] = 0x58
	}
	{
		size := m.SharesToTokensRate.Size()
		i -= size
		if _, err := m.SharesToTokensRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.SlashQueryProgressTracker.Size()
		i -= size
		if _, err := m.SlashQueryProgressTracker.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.Weight != 0 {
		i = encodeVarintValidator(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Delegation.Size()
		i -= size
		if _, err := m.Delegation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = m.Delegation.Size()
	n += 1 + l + sovValidator(uint64(l))
	if m.Weight != 0 {
		n += 1 + sovValidator(uint64(m.Weight))
	}
	l = m.SlashQueryProgressTracker.Size()
	n += 1 + l + sovValidator(uint64(l))
	l = m.SharesToTokensRate.Size()
	n += 1 + l + sovValidator(uint64(l))
	if m.DelegationChangesInProgress != 0 {
		n += 1 + sovValidator(uint64(m.DelegationChangesInProgress))
	}
	l = m.SlashQueryCheckpoint.Size()
	n += 1 + l + sovValidator(uint64(l))
	if m.SlashQueryInProgress {
		n += 2
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Delegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashQueryProgressTracker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashQueryProgressTracker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesToTokensRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesToTokensRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationChangesInProgress", wireType)
			}
			m.DelegationChangesInProgress = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelegationChangesInProgress |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashQueryCheckpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashQueryCheckpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashQueryInProgress", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SlashQueryInProgress = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
