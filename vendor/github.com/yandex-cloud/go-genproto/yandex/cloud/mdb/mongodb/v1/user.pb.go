// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mongodb/v1/user.proto

package mongodb // import "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A MongoDB User resource. For more information, see the
// [Developer's Guide](/docs/managed-mongodb/concepts).
type User struct {
	// Name of the MongoDB user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the MongoDB cluster the user belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Set of permissions granted to the user.
	Permissions          []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ce6db2103014c466, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *User) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Permission struct {
	// Name of the database that the permission grants access to.
	DatabaseName string `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	// MongoDB roles for the [database_name] database that the permission grants.
	Roles                []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ce6db2103014c466, []int{1}
}
func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (dst *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(dst, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

func (m *Permission) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type UserSpec struct {
	// Name of the MongoDB user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Password of the MongoDB user.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Set of permissions to grant to the user.
	Permissions          []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserSpec) Reset()         { *m = UserSpec{} }
func (m *UserSpec) String() string { return proto.CompactTextString(m) }
func (*UserSpec) ProtoMessage()    {}
func (*UserSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_ce6db2103014c466, []int{2}
}
func (m *UserSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSpec.Unmarshal(m, b)
}
func (m *UserSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSpec.Marshal(b, m, deterministic)
}
func (dst *UserSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSpec.Merge(dst, src)
}
func (m *UserSpec) XXX_Size() int {
	return xxx_messageInfo_UserSpec.Size(m)
}
func (m *UserSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UserSpec proto.InternalMessageInfo

func (m *UserSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserSpec) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserSpec) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "yandex.cloud.mdb.mongodb.v1.User")
	proto.RegisterType((*Permission)(nil), "yandex.cloud.mdb.mongodb.v1.Permission")
	proto.RegisterType((*UserSpec)(nil), "yandex.cloud.mdb.mongodb.v1.UserSpec")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mongodb/v1/user.proto", fileDescriptor_user_ce6db2103014c466)
}

var fileDescriptor_user_ce6db2103014c466 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xab, 0x4c, 0xcc, 0x4b,
	0x49, 0xad, 0xd0, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0xcf, 0x4d, 0x49, 0xd2, 0xcf, 0xcd, 0xcf,
	0x4b, 0xcf, 0x4f, 0x49, 0xd2, 0x2f, 0x33, 0xd4, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x92, 0x86, 0xa8, 0xd3, 0x03, 0xab, 0xd3, 0xcb, 0x4d, 0x49, 0xd2, 0x83, 0xaa,
	0xd3, 0x2b, 0x33, 0x94, 0x92, 0x45, 0x31, 0xa4, 0x2c, 0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0x33,
	0x3f, 0x0f, 0xa2, 0x57, 0xa9, 0x85, 0x91, 0x8b, 0x25, 0xb4, 0x38, 0xb5, 0x48, 0x48, 0x88, 0x8b,
	0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x92, 0xe5,
	0xe2, 0x4a, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0x8a, 0xcf, 0x4c, 0x91, 0x60, 0x02, 0xcb, 0x70,
	0x42, 0x45, 0x3c, 0x53, 0x84, 0x3c, 0xb9, 0xb8, 0x0b, 0x52, 0x8b, 0x72, 0x33, 0x8b, 0x8b, 0x33,
	0xf3, 0xf3, 0x8a, 0x25, 0x98, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xd4, 0xf5, 0xf0, 0xb8, 0x46, 0x2f,
	0x00, 0xae, 0x3e, 0x08, 0x59, 0xaf, 0x92, 0x3b, 0x17, 0x17, 0x42, 0x4a, 0x48, 0x99, 0x8b, 0x37,
	0x25, 0xb1, 0x24, 0x31, 0x29, 0xb1, 0x38, 0x35, 0x1e, 0xc9, 0x51, 0x3c, 0x30, 0x41, 0x3f, 0x90,
	0xe3, 0x44, 0xb8, 0x58, 0x8b, 0xf2, 0x73, 0x52, 0x8b, 0x25, 0x98, 0x14, 0x98, 0x35, 0x38, 0x83,
	0x20, 0x1c, 0xa5, 0xcd, 0x8c, 0x5c, 0x1c, 0x20, 0xff, 0x04, 0x17, 0xa4, 0x26, 0x0b, 0x19, 0x22,
	0xfb, 0xc9, 0x49, 0xf6, 0xc5, 0x71, 0x43, 0xc6, 0x4f, 0xc7, 0x0d, 0x79, 0xa3, 0x13, 0x75, 0xab,
	0x1c, 0x75, 0xa3, 0x0c, 0x74, 0x2d, 0xe3, 0x63, 0xb5, 0xba, 0x4e, 0x18, 0xb2, 0xd8, 0xd8, 0x9a,
	0x19, 0x43, 0xbd, 0xac, 0xc9, 0xc5, 0x51, 0x90, 0x58, 0x5c, 0x5c, 0x9e, 0x5f, 0x04, 0xf5, 0xb0,
	0x13, 0x2f, 0x48, 0x5b, 0xd7, 0x09, 0x43, 0x56, 0x0b, 0x5d, 0x43, 0x23, 0x8b, 0x20, 0xb8, 0x34,
	0x15, 0xbd, 0xef, 0xe4, 0x19, 0xe5, 0x9e, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f,
	0xab, 0x0f, 0x31, 0x41, 0x17, 0x12, 0x63, 0xe9, 0xf9, 0xba, 0xe9, 0xa9, 0x79, 0xe0, 0xc8, 0xd2,
	0xc7, 0x93, 0x1e, 0xac, 0xa1, 0xcc, 0x24, 0x36, 0xb0, 0x52, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb7, 0x04, 0x46, 0xbf, 0x3d, 0x02, 0x00, 0x00,
}