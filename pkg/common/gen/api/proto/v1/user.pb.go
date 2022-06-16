// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: api/proto/v1/user.proto

package pbv1

import (
	_ "github.com/alta/protopatch/patch/gopb"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserGroup int32

const (
	UserGroupGuest    UserGroup = 0
	UserGroupCustomer UserGroup = 1
)

// Enum value maps for UserGroup.
var (
	UserGroup_name = map[int32]string{
		0: "USER_GROUP_GUEST",
		1: "USER_GROUP_CUSTOMER",
	}
	UserGroup_value = map[string]int32{
		"USER_GROUP_GUEST":    0,
		"USER_GROUP_CUSTOMER": 1,
	}
)

func (x UserGroup) Enum() *UserGroup {
	p := new(UserGroup)
	*p = x
	return p
}

func (x UserGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_v1_user_proto_enumTypes[0].Descriptor()
}

func (UserGroup) Type() protoreflect.EnumType {
	return &file_api_proto_v1_user_proto_enumTypes[0]
}

func (x UserGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserGroup.Descriptor instead.
func (UserGroup) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_v1_user_proto_rawDescGZIP(), []int{0}
}

type UserStatus int32

const (
	UserStatusUnspecified UserStatus = 0
	UserStatusUnconfirmed UserStatus = 1
	UserStatusActive      UserStatus = 2
	UserStatusBlocked     UserStatus = 3
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0: "USER_STATUS_UNSPECIFIED",
		1: "USER_STATUS_UNCONFIRMED",
		2: "USER_STATUS_ACTIVE",
		3: "USER_STATUS_BLOCKED",
	}
	UserStatus_value = map[string]int32{
		"USER_STATUS_UNSPECIFIED": 0,
		"USER_STATUS_UNCONFIRMED": 1,
		"USER_STATUS_ACTIVE":      2,
		"USER_STATUS_BLOCKED":     3,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_v1_user_proto_enumTypes[1].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_api_proto_v1_user_proto_enumTypes[1]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_v1_user_proto_rawDescGZIP(), []int{1}
}

type PermissionObject int32

const (
	PermissionObjectUnspecified PermissionObject = 0
	PermissionObjectUser        PermissionObject = 1
	PermissionObjectProduct     PermissionObject = 2
	PermissionObjectOrder       PermissionObject = 3
	PermissionObjectUserProduct PermissionObject = 4
)

// Enum value maps for PermissionObject.
var (
	PermissionObject_name = map[int32]string{
		0: "PERMISSION_OBJECT_UNSPECIFIED",
		1: "PERMISSION_OBJECT_USER",
		2: "PERMISSION_OBJECT_PRODUCT",
		3: "PERMISSION_OBJECT_ORDER",
		4: "PERMISSION_OBJECT_USER_PRODUCT",
	}
	PermissionObject_value = map[string]int32{
		"PERMISSION_OBJECT_UNSPECIFIED":  0,
		"PERMISSION_OBJECT_USER":         1,
		"PERMISSION_OBJECT_PRODUCT":      2,
		"PERMISSION_OBJECT_ORDER":        3,
		"PERMISSION_OBJECT_USER_PRODUCT": 4,
	}
)

func (x PermissionObject) Enum() *PermissionObject {
	p := new(PermissionObject)
	*p = x
	return p
}

func (x PermissionObject) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionObject) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_v1_user_proto_enumTypes[2].Descriptor()
}

func (PermissionObject) Type() protoreflect.EnumType {
	return &file_api_proto_v1_user_proto_enumTypes[2]
}

func (x PermissionObject) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionObject.Descriptor instead.
func (PermissionObject) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_v1_user_proto_rawDescGZIP(), []int{2}
}

type PermissionOperation int32

const (
	PermissionOperationUnspecified PermissionOperation = 0
	PermissionOperationRead        PermissionOperation = 1
	PermissionOperationWrite       PermissionOperation = 2
)

// Enum value maps for PermissionOperation.
var (
	PermissionOperation_name = map[int32]string{
		0: "PERMISSION_OPERATION_UNSPECIFIED",
		1: "PERMISSION_OPERATION_READ",
		2: "PERMISSION_OPERATION_WRITE",
	}
	PermissionOperation_value = map[string]int32{
		"PERMISSION_OPERATION_UNSPECIFIED": 0,
		"PERMISSION_OPERATION_READ":        1,
		"PERMISSION_OPERATION_WRITE":       2,
	}
)

func (x PermissionOperation) Enum() *PermissionOperation {
	p := new(PermissionOperation)
	*p = x
	return p
}

func (x PermissionOperation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionOperation) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_v1_user_proto_enumTypes[3].Descriptor()
}

func (PermissionOperation) Type() protoreflect.EnumType {
	return &file_api_proto_v1_user_proto_enumTypes[3]
}

func (x PermissionOperation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionOperation.Descriptor instead.
func (PermissionOperation) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_v1_user_proto_rawDescGZIP(), []int{3}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Group       UserGroup              `protobuf:"varint,2,opt,name=group,proto3,enum=pb.v1.UserGroup" json:"group,omitempty"`
	Status      UserStatus             `protobuf:"varint,3,opt,name=status,proto3,enum=pb.v1.UserStatus" json:"status,omitempty"`
	PhoneNumber string                 `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Lastname    string                 `protobuf:"bytes,5,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Firstname   string                 `protobuf:"bytes,6,opt,name=firstname,proto3" json:"firstname,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *User) GetGroup() UserGroup {
	if x != nil {
		return x.Group
	}
	return UserGroupGuest
}

func (x *User) GetStatus() UserStatus {
	if x != nil {
		return x.Status
	}
	return UserStatusUnspecified
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Group UserGroup `protobuf:"varint,2,opt,name=group,proto3,enum=pb.v1.UserGroup" json:"group,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *Session) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Session) GetGroup() UserGroup {
	if x != nil {
		return x.Group
	}
	return UserGroupGuest
}

var File_api_proto_v1_user_proto protoreflect.FileDescriptor

var file_api_proto_v1_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x85, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x45, 0x0a, 0x07, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2a, 0x69, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2a, 0x0a,
	0x10, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x47, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x00, 0x1a, 0x14, 0xca, 0xb5, 0x03, 0x10, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52,
	0x10, 0x01, 0x1a, 0x17, 0xca, 0xb5, 0x03, 0x13, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2a, 0xe2, 0x01, 0x0a, 0x0a,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x17, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x1b, 0xca, 0xb5, 0x03, 0x17, 0x0a, 0x15, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x17, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10,
	0x01, 0x1a, 0x1b, 0xca, 0xb5, 0x03, 0x17, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x2e,
	0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x1a, 0x16, 0xca, 0xb5, 0x03, 0x12, 0x0a, 0x10, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x30,
	0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x4c,
	0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x03, 0x1a, 0x17, 0xca, 0xb5, 0x03, 0x13, 0x0a, 0x11, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x2a, 0xcf, 0x02, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x44, 0x0a, 0x1d, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x21, 0xca, 0xb5, 0x03, 0x1d, 0x0a, 0x1b,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x1a, 0x1a, 0xca, 0xb5, 0x03, 0x16, 0x0a, 0x14, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x19, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54,
	0x10, 0x02, 0x1a, 0x1d, 0xca, 0xb5, 0x03, 0x19, 0x0a, 0x17, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x38, 0x0a, 0x17, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x03, 0x1a, 0x1b,
	0xca, 0xb5, 0x03, 0x17, 0x0a, 0x15, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x1e, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x10, 0x04, 0x1a,
	0x21, 0xca, 0xb5, 0x03, 0x1d, 0x0a, 0x1b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2a, 0xdf, 0x01, 0x0a, 0x13, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x20, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x1a, 0x24, 0xca, 0xb5, 0x03, 0x20, 0x0a, 0x1e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x19, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52,
	0x45, 0x41, 0x44, 0x10, 0x01, 0x1a, 0x1d, 0xca, 0xb5, 0x03, 0x19, 0x0a, 0x17, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x61, 0x64, 0x12, 0x3e, 0x0a, 0x1a, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x10, 0x02, 0x1a, 0x1e, 0xca, 0xb5, 0x03, 0x1a, 0x0a, 0x18, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x79, 0x6f, 0x72, 0x74, 0x79, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_user_proto_rawDescOnce sync.Once
	file_api_proto_v1_user_proto_rawDescData = file_api_proto_v1_user_proto_rawDesc
)

func file_api_proto_v1_user_proto_rawDescGZIP() []byte {
	file_api_proto_v1_user_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_user_proto_rawDescData)
	})
	return file_api_proto_v1_user_proto_rawDescData
}

var file_api_proto_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_api_proto_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_v1_user_proto_goTypes = []interface{}{
	(UserGroup)(0),                // 0: pb.v1.UserGroup
	(UserStatus)(0),               // 1: pb.v1.UserStatus
	(PermissionObject)(0),         // 2: pb.v1.PermissionObject
	(PermissionOperation)(0),      // 3: pb.v1.PermissionOperation
	(*User)(nil),                  // 4: pb.v1.User
	(*Session)(nil),               // 5: pb.v1.Session
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_api_proto_v1_user_proto_depIdxs = []int32{
	0, // 0: pb.v1.User.group:type_name -> pb.v1.UserGroup
	1, // 1: pb.v1.User.status:type_name -> pb.v1.UserStatus
	6, // 2: pb.v1.User.created_at:type_name -> google.protobuf.Timestamp
	0, // 3: pb.v1.Session.group:type_name -> pb.v1.UserGroup
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_v1_user_proto_init() }
func file_api_proto_v1_user_proto_init() {
	if File_api_proto_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_v1_user_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_v1_user_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_user_proto_depIdxs,
		EnumInfos:         file_api_proto_v1_user_proto_enumTypes,
		MessageInfos:      file_api_proto_v1_user_proto_msgTypes,
	}.Build()
	File_api_proto_v1_user_proto = out.File
	file_api_proto_v1_user_proto_rawDesc = nil
	file_api_proto_v1_user_proto_goTypes = nil
	file_api_proto_v1_user_proto_depIdxs = nil
}
