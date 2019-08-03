// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//
type User struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

//
type ListUserRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserRequest) Reset()         { *m = ListUserRequest{} }
func (m *ListUserRequest) String() string { return proto.CompactTextString(m) }
func (*ListUserRequest) ProtoMessage()    {}
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *ListUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserRequest.Unmarshal(m, b)
}
func (m *ListUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserRequest.Marshal(b, m, deterministic)
}
func (m *ListUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserRequest.Merge(m, src)
}
func (m *ListUserRequest) XXX_Size() int {
	return xxx_messageInfo_ListUserRequest.Size(m)
}
func (m *ListUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserRequest proto.InternalMessageInfo

type ListUserResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserResponse) Reset()         { *m = ListUserResponse{} }
func (m *ListUserResponse) String() string { return proto.CompactTextString(m) }
func (*ListUserResponse) ProtoMessage()    {}
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *ListUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserResponse.Unmarshal(m, b)
}
func (m *ListUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserResponse.Marshal(b, m, deterministic)
}
func (m *ListUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserResponse.Merge(m, src)
}
func (m *ListUserResponse) XXX_Size() int {
	return xxx_messageInfo_ListUserResponse.Size(m)
}
func (m *ListUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserResponse proto.InternalMessageInfo

func (m *ListUserResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ListUserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

//
type GetUserRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

//
type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateUserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

//
type UpdateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateUserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type UpdateUserResponse struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(m, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

func (m *UpdateUserResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateUserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

//
type DeleteUserRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type DeleteUserResponse struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (m *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(m, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

func (m *DeleteUserResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *DeleteUserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "v1.User")
	proto.RegisterType((*ListUserRequest)(nil), "v1.ListUserRequest")
	proto.RegisterType((*ListUserResponse)(nil), "v1.ListUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "v1.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "v1.GetUserResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "v1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "v1.CreateUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "v1.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "v1.UpdateUserResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "v1.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "v1.DeleteUserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0x38, 0x69, 0xcd, 0x54, 0x10, 0x3c, 0x6d, 0xa3, 0x95, 0xa9, 0xaa, 0xb0, 0x27,
	0xab, 0x48, 0xb1, 0x5c, 0x6e, 0x5c, 0x41, 0x41, 0x95, 0x72, 0x32, 0xca, 0x09, 0x09, 0xc9, 0x24,
	0xa3, 0x62, 0x29, 0xf5, 0x1a, 0xaf, 0x63, 0x0e, 0x88, 0x4b, 0x5f, 0x81, 0x47, 0xe2, 0x11, 0x78,
	0x05, 0x1e, 0x04, 0xed, 0xae, 0x5d, 0xdb, 0xb5, 0x8c, 0x44, 0x6e, 0x99, 0xd9, 0x99, 0xef, 0x9f,
	0xdd, 0xf9, 0x63, 0x80, 0xbd, 0xa4, 0x7c, 0x91, 0xe5, 0xa2, 0x10, 0x38, 0x2a, 0x43, 0xef, 0xe2,
	0x56, 0x88, 0xdb, 0x1d, 0x05, 0x71, 0x96, 0x04, 0x71, 0x9a, 0x8a, 0x22, 0x2e, 0x12, 0x91, 0x4a,
	0x53, 0xc1, 0xef, 0x2d, 0x18, 0xaf, 0x25, 0xe5, 0x38, 0x83, 0x23, 0xd5, 0x78, 0xb3, 0x65, 0xd6,
	0xdc, 0xf2, 0xed, 0xa8, 0x8a, 0xd0, 0x03, 0x47, 0xfd, 0x4a, 0xe3, 0x3b, 0x62, 0xa3, 0xb9, 0xe5,
	0x3f, 0x89, 0x1e, 0x62, 0x3c, 0x83, 0x09, 0xdd, 0xc5, 0xc9, 0x8e, 0xd9, 0xfa, 0xc0, 0x04, 0xaa,
	0x23, 0x8b, 0xa5, 0xfc, 0x26, 0xf2, 0x2d, 0x1b, 0x9b, 0x8e, 0x3a, 0x56, 0x1d, 0xd9, 0x17, 0x91,
	0x12, 0x9b, 0x98, 0x0e, 0x1d, 0x70, 0x17, 0xa6, 0xab, 0x44, 0x16, 0x6a, 0x8e, 0x88, 0xbe, 0xee,
	0x49, 0x16, 0x7c, 0x05, 0xcf, 0x9b, 0x94, 0xcc, 0x44, 0x2a, 0x09, 0x2f, 0x61, 0xa2, 0xa4, 0x25,
	0xb3, 0xe6, 0xb6, 0x7f, 0x72, 0xed, 0x2c, 0xca, 0x70, 0xa1, 0x0b, 0x4c, 0x1a, 0x19, 0x1c, 0xcb,
	0xfd, 0x66, 0x43, 0x52, 0xea, 0x49, 0x9d, 0xa8, 0x0e, 0xb9, 0x0f, 0xcf, 0xde, 0x53, 0x9b, 0x3f,
	0x74, 0x5d, 0x1e, 0xc0, 0xf4, 0xa1, 0xb2, 0x92, 0xbd, 0x80, 0xb1, 0x3a, 0xd4, 0x85, 0x6d, 0x55,
	0x9d, 0xe5, 0x21, 0xb8, 0x6f, 0x73, 0x8a, 0x0b, 0x6a, 0xd3, 0xff, 0xdd, 0xb2, 0x04, 0x6c, 0xb7,
	0x54, 0x32, 0x43, 0x0b, 0x18, 0xbe, 0xd5, 0x0d, 0xb8, 0xeb, 0x6c, 0xfb, 0x3f, 0xd2, 0x2d, 0x91,
	0x51, 0xe7, 0xda, 0x4b, 0xc0, 0x36, 0xea, 0xe0, 0x91, 0x5e, 0x81, 0xfb, 0x8e, 0x76, 0xd4, 0x1d,
	0x69, 0xe8, 0xad, 0x97, 0x80, 0xed, 0xe2, 0x43, 0x45, 0xaf, 0x7f, 0xd9, 0x70, 0xa2, 0x10, 0x1f,
	0x28, 0x2f, 0x93, 0x0d, 0xe1, 0x0a, 0x9c, 0xda, 0x3b, 0x78, 0xaa, 0x1e, 0xe0, 0x91, 0xb9, 0xbc,
	0xb3, 0x6e, 0xd2, 0x08, 0xf3, 0xf3, 0xfb, 0xdf, 0x7f, 0x7e, 0x8e, 0xa6, 0xf8, 0x54, 0xff, 0x55,
	0xca, 0x30, 0x30, 0xae, 0x8a, 0xe0, 0xb8, 0x72, 0x04, 0xa2, 0xea, 0xeb, 0x1a, 0xc9, 0x3b, 0xed,
	0xe4, 0x2a, 0xd4, 0xa5, 0x46, 0x31, 0x9c, 0x75, 0x50, 0xc1, 0x77, 0x73, 0x95, 0x1f, 0xb8, 0x06,
	0x68, 0x1c, 0x80, 0xe7, 0x0a, 0xd1, 0x33, 0x91, 0x37, 0x7b, 0x9c, 0xae, 0xe0, 0x4c, 0xc3, 0x91,
	0x77, 0xe7, 0x7c, 0x63, 0x5d, 0xe1, 0x27, 0x80, 0x66, 0x8b, 0x06, 0xdb, 0x33, 0x88, 0xc1, 0xf6,
	0x97, 0xcd, 0x5f, 0x6a, 0xec, 0x0b, 0x6f, 0x60, 0x66, 0xc5, 0xff, 0x08, 0xd0, 0x2c, 0xcc, 0xf0,
	0x7b, 0xdb, 0x36, 0xfc, 0xfe, 0x5e, 0xeb, 0x37, 0xb9, 0x1a, 0xe0, 0x7f, 0x3e, 0xd2, 0x1f, 0xa4,
	0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xa5, 0x11, 0x98, 0xc0, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) ListUser(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServiceServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _UserService_ListUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
