// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/v1/user.proto

package pbv1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/type/datetime"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type UserGroup int32

const (
	UserGroupGuest    UserGroup = 0
	UserGroupCustomer UserGroup = 1
)

var UserGroup_name = map[int32]string{
	0: "UserGroupGuest",
	1: "UserGroupCustomer",
}

var UserGroup_value = map[string]int32{
	"UserGroupGuest":    0,
	"UserGroupCustomer": 1,
}

func (UserGroup) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b1212c297c1885f, []int{0}
}

type UserStatus int32

const (
	UserStatusUnverified  UserStatus = 0
	UserStatusUnconfirmed UserStatus = 1
	UserStatusActive      UserStatus = 2
	UserStatusBlocked     UserStatus = 3
)

var UserStatus_name = map[int32]string{
	0: "UserStatusUnverified",
	1: "UserStatusUnconfirmed",
	2: "UserStatusActive",
	3: "UserStatusBlocked",
}

var UserStatus_value = map[string]int32{
	"UserStatusUnverified":  0,
	"UserStatusUnconfirmed": 1,
	"UserStatusActive":      2,
	"UserStatusBlocked":     3,
}

func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1b1212c297c1885f, []int{1}
}

type User struct {
	Uuid      string     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Group     UserGroup  `protobuf:"varint,2,opt,name=group,proto3,enum=pb.v1.UserGroup" json:"group,omitempty"`
	Status    UserStatus `protobuf:"varint,3,opt,name=status,proto3,enum=pb.v1.UserStatus" json:"status,omitempty"`
	Email     string     `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Lastname  string     `protobuf:"bytes,5,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Firstname string     `protobuf:"bytes,6,opt,name=firstname,proto3" json:"firstname,omitempty"`
	CreatedAt *time.Time `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty"`
}

func (m *User) Reset()      { *m = User{} }
func (*User) ProtoMessage() {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b1212c297c1885f, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *User) GetGroup() UserGroup {
	if m != nil {
		return m.Group
	}
	return UserGroupGuest
}

func (m *User) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatusUnverified
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *User) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *User) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.v1.UserGroup", UserGroup_name, UserGroup_value)
	golang_proto.RegisterEnum("pb.v1.UserGroup", UserGroup_name, UserGroup_value)
	proto.RegisterEnum("pb.v1.UserStatus", UserStatus_name, UserStatus_value)
	golang_proto.RegisterEnum("pb.v1.UserStatus", UserStatus_name, UserStatus_value)
	proto.RegisterType((*User)(nil), "pb.v1.User")
	golang_proto.RegisterType((*User)(nil), "pb.v1.User")
}

func init() { proto.RegisterFile("proto/v1/user.proto", fileDescriptor_1b1212c297c1885f) }
func init() { golang_proto.RegisterFile("proto/v1/user.proto", fileDescriptor_1b1212c297c1885f) }

var fileDescriptor_1b1212c297c1885f = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x31, 0x6f, 0xd3, 0x50,
	0x14, 0x85, 0xfd, 0xda, 0x24, 0x90, 0x8b, 0x54, 0xb9, 0x8f, 0x44, 0x32, 0x11, 0x7a, 0x44, 0x0c,
	0x28, 0x54, 0xaa, 0xad, 0x14, 0x89, 0x81, 0xad, 0x01, 0xa9, 0x7b, 0xa0, 0x0b, 0x0b, 0x7a, 0xb6,
	0x6f, 0xcc, 0x13, 0xb6, 0x9f, 0xf5, 0x7c, 0x6d, 0x29, 0x1b, 0x3f, 0xa1, 0x23, 0x3f, 0x81, 0x9f,
	0xc1, 0xc8, 0x98, 0xb1, 0x1b, 0xc4, 0x5e, 0x18, 0x3b, 0xc2, 0x86, 0x62, 0xa7, 0x49, 0xb6, 0x7b,
	0xce, 0x77, 0xa4, 0xa3, 0x23, 0x5d, 0x78, 0x9c, 0x19, 0x4d, 0xda, 0x2b, 0xa7, 0x5e, 0x91, 0xa3,
	0x71, 0x1b, 0xc5, 0xbb, 0x99, 0xef, 0x96, 0xd3, 0xd1, 0x20, 0xd2, 0x91, 0x6e, 0xf9, 0xe6, 0x6a,
	0xe1, 0x68, 0x14, 0x69, 0x1d, 0xc5, 0xe8, 0xd1, 0x32, 0x43, 0x2f, 0x94, 0x84, 0xa4, 0x12, 0x6c,
	0xd9, 0xf3, 0x7f, 0x0c, 0x3a, 0xd7, 0x39, 0x1a, 0xce, 0xa1, 0x53, 0x14, 0x2a, 0x74, 0xd8, 0x98,
	0x4d, 0xfa, 0xf3, 0xe6, 0xe6, 0x2f, 0xa0, 0x1b, 0x19, 0x5d, 0x64, 0xce, 0xd1, 0x98, 0x4d, 0x4e,
	0x2e, 0x6c, 0xb7, 0x69, 0x71, 0x37, 0xf9, 0xab, 0x8d, 0x3f, 0x6f, 0x31, 0x7f, 0x09, 0xbd, 0x9c,
	0x24, 0x15, 0xb9, 0x73, 0xdc, 0x04, 0x4f, 0x0f, 0x82, 0xef, 0x1b, 0x30, 0xdf, 0x06, 0xf8, 0x00,
	0xba, 0x98, 0x48, 0x15, 0x3b, 0x9d, 0xa6, 0xa7, 0x15, 0x7c, 0x04, 0x0f, 0x63, 0x99, 0x53, 0x2a,
	0x13, 0x74, 0xba, 0x0d, 0xd8, 0x69, 0xfe, 0x14, 0xfa, 0x0b, 0x65, 0xb6, 0xb0, 0xd7, 0xc0, 0xbd,
	0xc1, 0xdf, 0x00, 0x04, 0x06, 0x25, 0x61, 0xf8, 0x49, 0x92, 0xf3, 0x60, 0xcc, 0x26, 0x8f, 0x2e,
	0x86, 0x6e, 0x3b, 0xd8, 0xdd, 0x0c, 0x76, 0xdf, 0x49, 0xc2, 0x0f, 0x2a, 0xc1, 0x59, 0xe7, 0xe6,
	0xd7, 0x33, 0x36, 0xef, 0x6f, 0xe3, 0x97, 0x74, 0xf6, 0x1a, 0xfa, 0xbb, 0x29, 0x9c, 0xc3, 0xc9,
	0x4e, 0x5c, 0x15, 0x98, 0x93, 0x6d, 0xf1, 0x21, 0x9c, 0xee, 0xbc, 0xb7, 0x45, 0x4e, 0x3a, 0x41,
	0x63, 0xb3, 0xb3, 0x14, 0x60, 0xbf, 0x8c, 0x3b, 0x30, 0xd8, 0xab, 0xeb, 0xb4, 0x44, 0xa3, 0x16,
	0x0a, 0x43, 0xdb, 0xe2, 0x4f, 0x60, 0x78, 0x48, 0x02, 0x9d, 0x2e, 0x94, 0x49, 0x30, 0xb4, 0x19,
	0x1f, 0x80, 0xbd, 0x47, 0x97, 0x01, 0xa9, 0x12, 0xed, 0xa3, 0xfb, 0xbe, 0xd6, 0x9d, 0xc5, 0x3a,
	0xf8, 0x82, 0xa1, 0x7d, 0x3c, 0xd3, 0xab, 0xb5, 0xb0, 0x6e, 0xd7, 0xc2, 0xba, 0x5b, 0x0b, 0xf6,
	0xb5, 0x12, 0xec, 0x7b, 0x25, 0xd8, 0xcf, 0x4a, 0xb0, 0x55, 0x25, 0xd8, 0xef, 0x4a, 0xb0, 0x3f,
	0x95, 0xb0, 0xee, 0x2a, 0xc1, 0xfe, 0x56, 0x82, 0xdd, 0xd4, 0xc2, 0xfa, 0x56, 0x0b, 0xf6, 0xa3,
	0x16, 0x6c, 0x55, 0x0b, 0xeb, 0xb6, 0x16, 0xd6, 0xc7, 0xf3, 0x48, 0xd1, 0xe7, 0xc2, 0x77, 0x03,
	0x9d, 0x78, 0xfe, 0x52, 0x1b, 0x5a, 0x7a, 0x98, 0x12, 0x9a, 0xcc, 0xa8, 0x1c, 0xcf, 0x65, 0x96,
	0xc5, 0x2a, 0x90, 0xa4, 0x74, 0xea, 0x65, 0x7e, 0x39, 0xf5, 0x7b, 0xcd, 0x6f, 0xbc, 0xfa, 0x1f,
	0x00, 0x00, 0xff, 0xff, 0xf0, 0x7d, 0xda, 0x7b, 0x6b, 0x02, 0x00, 0x00,
}

func (x UserGroup) String() string {
	s, ok := UserGroup_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x UserStatus) String() string {
	s, ok := UserStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *User) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*User)
	if !ok {
		that2, ok := that.(User)
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
	if this.Uuid != that1.Uuid {
		return false
	}
	if this.Group != that1.Group {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if this.Email != that1.Email {
		return false
	}
	if this.Lastname != that1.Lastname {
		return false
	}
	if this.Firstname != that1.Firstname {
		return false
	}
	if that1.CreatedAt == nil {
		if this.CreatedAt != nil {
			return false
		}
	} else if !this.CreatedAt.Equal(*that1.CreatedAt) {
		return false
	}
	return true
}
func (this *User) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&pbv1.User{")
	s = append(s, "Uuid: "+fmt.Sprintf("%#v", this.Uuid)+",\n")
	s = append(s, "Group: "+fmt.Sprintf("%#v", this.Group)+",\n")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Email: "+fmt.Sprintf("%#v", this.Email)+",\n")
	s = append(s, "Lastname: "+fmt.Sprintf("%#v", this.Lastname)+",\n")
	s = append(s, "Firstname: "+fmt.Sprintf("%#v", this.Firstname)+",\n")
	s = append(s, "CreatedAt: "+fmt.Sprintf("%#v", this.CreatedAt)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringUser(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintUser(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Firstname) > 0 {
		i -= len(m.Firstname)
		copy(dAtA[i:], m.Firstname)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Firstname)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Lastname) > 0 {
		i -= len(m.Lastname)
		copy(dAtA[i:], m.Lastname)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Lastname)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Email) > 0 {
		i -= len(m.Email)
		copy(dAtA[i:], m.Email)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Email)))
		i--
		dAtA[i] = 0x22
	}
	if m.Status != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if m.Group != 0 {
		i = encodeVarintUser(dAtA, i, uint64(m.Group))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Uuid) > 0 {
		i -= len(m.Uuid)
		copy(dAtA[i:], m.Uuid)
		i = encodeVarintUser(dAtA, i, uint64(len(m.Uuid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.Group != 0 {
		n += 1 + sovUser(uint64(m.Group))
	}
	if m.Status != 0 {
		n += 1 + sovUser(uint64(m.Status))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Lastname)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Firstname)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.CreatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func sovUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *User) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&User{`,
		`Uuid:` + fmt.Sprintf("%v", this.Uuid) + `,`,
		`Group:` + fmt.Sprintf("%v", this.Group) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Email:` + fmt.Sprintf("%v", this.Email) + `,`,
		`Lastname:` + fmt.Sprintf("%v", this.Lastname) + `,`,
		`Firstname:` + fmt.Sprintf("%v", this.Firstname) + `,`,
		`CreatedAt:` + strings.Replace(fmt.Sprintf("%v", this.CreatedAt), "DateTime", "datetime.DateTime", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringUser(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			m.Group = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Group |= UserGroup(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= UserStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lastname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lastname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Firstname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Firstname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUser = fmt.Errorf("proto: unexpected end of group")
)
