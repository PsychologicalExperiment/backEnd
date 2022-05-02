// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/user_info_server/user_info_server.proto

package user_info_server

import (
	api_common "github.com/PsychologicalExperiment/backEnd/api/api_common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserType int32

const (
	// 主试类型
	UserType_USER_TYPE_RESEARCHER UserType = 0
	// 被试类型
	UserType_USER_TYPE_PARTICIPANT UserType = 1
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "USER_TYPE_RESEARCHER",
		1: "USER_TYPE_PARTICIPANT",
	}
	UserType_value = map[string]int32{
		"USER_TYPE_RESEARCHER":  0,
		"USER_TYPE_PARTICIPANT": 1,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_info_server_user_info_server_proto_enumTypes[0].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_api_user_info_server_user_info_server_proto_enumTypes[0]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_api_user_info_server_user_info_server_proto_rawDescGZIP(), []int{0}
}

type GenderType int32

const (
	GenderType_GENDER_TYPE_INVALID GenderType = 0
	// man
	GenderType_GENDER_TYPE_MAN GenderType = 1
	// woman
	GenderType_GENDER_TYPE_WOMAN GenderType = 2
)

// Enum value maps for GenderType.
var (
	GenderType_name = map[int32]string{
		0: "GENDER_TYPE_INVALID",
		1: "GENDER_TYPE_MAN",
		2: "GENDER_TYPE_WOMAN",
	}
	GenderType_value = map[string]int32{
		"GENDER_TYPE_INVALID": 0,
		"GENDER_TYPE_MAN":     1,
		"GENDER_TYPE_WOMAN":   2,
	}
)

func (x GenderType) Enum() *GenderType {
	p := new(GenderType)
	*p = x
	return p
}

func (x GenderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GenderType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_user_info_server_user_info_server_proto_enumTypes[1].Descriptor()
}

func (GenderType) Type() protoreflect.EnumType {
	return &file_api_user_info_server_user_info_server_proto_enumTypes[1]
}

func (x GenderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GenderType.Descriptor instead.
func (GenderType) EnumDescriptor() ([]byte, []int) {
	return file_api_user_info_server_user_info_server_proto_rawDescGZIP(), []int{1}
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//email
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	//phone number
	PhoneNumber string `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	//user name
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// gender
	Gender GenderType `protobuf:"varint,4,opt,name=gender,proto3,enum=grpc.psychological_experiment.user_info_server.GenderType" json:"gender,omitempty"`
	// user type
	UserType UserType `protobuf:"varint,5,opt,name=user_type,json=userType,proto3,enum=grpc.psychological_experiment.user_info_server.UserType" json:"user_type,omitempty"`
	// custom user form
	Extra string `protobuf:"bytes,6,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_info_server_user_info_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_info_server_user_info_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_api_user_info_server_user_info_server_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserInfo) GetGender() GenderType {
	if x != nil {
		return x.Gender
	}
	return GenderType_GENDER_TYPE_INVALID
}

func (x *UserInfo) GetUserType() UserType {
	if x != nil {
		return x.UserType
	}
	return UserType_USER_TYPE_RESEARCHER
}

func (x *UserInfo) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//user info
	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	//password
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_info_server_user_info_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_info_server_user_info_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_api_user_info_server_user_info_server_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReq) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// common rsp
	CommonRsp *api_common.CommonRsp `protobuf:"bytes,1,opt,name=common_rsp,json=commonRsp,proto3" json:"common_rsp,omitempty"`
	// uin
	Uin string `protobuf:"bytes,2,opt,name=uin,proto3" json:"uin,omitempty"`
}

func (x *RegisterRsp) Reset() {
	*x = RegisterRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_info_server_user_info_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRsp) ProtoMessage() {}

func (x *RegisterRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_info_server_user_info_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRsp.ProtoReflect.Descriptor instead.
func (*RegisterRsp) Descriptor() ([]byte, []int) {
	return file_api_user_info_server_user_info_server_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRsp) GetCommonRsp() *api_common.CommonRsp {
	if x != nil {
		return x.CommonRsp
	}
	return nil
}

func (x *RegisterRsp) GetUin() string {
	if x != nil {
		return x.Uin
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//email
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	//phone number
	PhoneNumber string `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	//password
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_info_server_user_info_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_info_server_user_info_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_api_user_info_server_user_info_server_proto_rawDescGZIP(), []int{3}
}

func (x *LoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// common rsp
	CommonRsp *api_common.CommonRsp `protobuf:"bytes,1,opt,name=common_rsp,json=commonRsp,proto3" json:"common_rsp,omitempty"`
	// user info
	UserInfo *UserInfo `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *LoginRsp) Reset() {
	*x = LoginRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_info_server_user_info_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRsp) ProtoMessage() {}

func (x *LoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_info_server_user_info_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRsp.ProtoReflect.Descriptor instead.
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return file_api_user_info_server_user_info_server_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRsp) GetCommonRsp() *api_common.CommonRsp {
	if x != nil {
		return x.CommonRsp
	}
	return nil
}

func (x *LoginRsp) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

var File_api_user_info_server_user_info_server_proto protoreflect.FileDescriptor

var file_api_user_info_server_user_info_server_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x73, 0x79, 0x63, 0x68, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x1f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1,
	0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x52, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x3a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x73, 0x79, 0x63, 0x68, 0x6f, 0x6c,
	0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x73, 0x79, 0x63, 0x68, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x55, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x73, 0x79,
	0x63, 0x68, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x73, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x73, 0x70, 0x12, 0x52, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x72,
	0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x73, 0x79, 0x63, 0x68, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x22, 0x5f, 0x0a, 0x08, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x52, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x72, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x73, 0x79, 0x63, 0x68, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73,
	0x70, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x55, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x73, 0x79, 0x63, 0x68, 0x6f, 0x6c, 0x6f, 0x67,
	0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2a, 0x3f, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x14, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x53,
	0x45, 0x41, 0x52, 0x43, 0x48, 0x45, 0x52, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x49, 0x43, 0x49, 0x50, 0x41,
	0x4e, 0x54, 0x10, 0x01, 0x2a, 0x51, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x47,
	0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x57, 0x4f, 0x4d, 0x41, 0x4e, 0x10, 0x02, 0x32, 0x95, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x73, 0x79, 0x63,
	0x68, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x3b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x73, 0x79, 0x63, 0x68, 0x6f, 0x6c,
	0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x7d, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x38, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x73, 0x79, 0x63, 0x68, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x38, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x73, 0x79, 0x63, 0x68,
	0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42,
	0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x73,
	0x79, 0x63, 0x68, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x45, 0x6e, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_info_server_user_info_server_proto_rawDescOnce sync.Once
	file_api_user_info_server_user_info_server_proto_rawDescData = file_api_user_info_server_user_info_server_proto_rawDesc
)

func file_api_user_info_server_user_info_server_proto_rawDescGZIP() []byte {
	file_api_user_info_server_user_info_server_proto_rawDescOnce.Do(func() {
		file_api_user_info_server_user_info_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_info_server_user_info_server_proto_rawDescData)
	})
	return file_api_user_info_server_user_info_server_proto_rawDescData
}

var file_api_user_info_server_user_info_server_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_user_info_server_user_info_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_user_info_server_user_info_server_proto_goTypes = []interface{}{
	(UserType)(0),                // 0: grpc.psychological_experiment.user_info_server.UserType
	(GenderType)(0),              // 1: grpc.psychological_experiment.user_info_server.GenderType
	(*UserInfo)(nil),             // 2: grpc.psychological_experiment.user_info_server.UserInfo
	(*RegisterReq)(nil),          // 3: grpc.psychological_experiment.user_info_server.RegisterReq
	(*RegisterRsp)(nil),          // 4: grpc.psychological_experiment.user_info_server.RegisterRsp
	(*LoginReq)(nil),             // 5: grpc.psychological_experiment.user_info_server.LoginReq
	(*LoginRsp)(nil),             // 6: grpc.psychological_experiment.user_info_server.LoginRsp
	(*api_common.CommonRsp)(nil), // 7: grpc.psychological_experiment.api_common.CommonRsp
}
var file_api_user_info_server_user_info_server_proto_depIdxs = []int32{
	1, // 0: grpc.psychological_experiment.user_info_server.UserInfo.gender:type_name -> grpc.psychological_experiment.user_info_server.GenderType
	0, // 1: grpc.psychological_experiment.user_info_server.UserInfo.user_type:type_name -> grpc.psychological_experiment.user_info_server.UserType
	2, // 2: grpc.psychological_experiment.user_info_server.RegisterReq.user_info:type_name -> grpc.psychological_experiment.user_info_server.UserInfo
	7, // 3: grpc.psychological_experiment.user_info_server.RegisterRsp.common_rsp:type_name -> grpc.psychological_experiment.api_common.CommonRsp
	7, // 4: grpc.psychological_experiment.user_info_server.LoginRsp.common_rsp:type_name -> grpc.psychological_experiment.api_common.CommonRsp
	2, // 5: grpc.psychological_experiment.user_info_server.LoginRsp.user_info:type_name -> grpc.psychological_experiment.user_info_server.UserInfo
	3, // 6: grpc.psychological_experiment.user_info_server.UserService.Register:input_type -> grpc.psychological_experiment.user_info_server.RegisterReq
	5, // 7: grpc.psychological_experiment.user_info_server.UserService.Login:input_type -> grpc.psychological_experiment.user_info_server.LoginReq
	4, // 8: grpc.psychological_experiment.user_info_server.UserService.Register:output_type -> grpc.psychological_experiment.user_info_server.RegisterRsp
	6, // 9: grpc.psychological_experiment.user_info_server.UserService.Login:output_type -> grpc.psychological_experiment.user_info_server.LoginRsp
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_user_info_server_user_info_server_proto_init() }
func file_api_user_info_server_user_info_server_proto_init() {
	if File_api_user_info_server_user_info_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_info_server_user_info_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_api_user_info_server_user_info_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_api_user_info_server_user_info_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRsp); i {
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
		file_api_user_info_server_user_info_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_api_user_info_server_user_info_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRsp); i {
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
			RawDescriptor: file_api_user_info_server_user_info_server_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_info_server_user_info_server_proto_goTypes,
		DependencyIndexes: file_api_user_info_server_user_info_server_proto_depIdxs,
		EnumInfos:         file_api_user_info_server_user_info_server_proto_enumTypes,
		MessageInfos:      file_api_user_info_server_user_info_server_proto_msgTypes,
	}.Build()
	File_api_user_info_server_user_info_server_proto = out.File
	file_api_user_info_server_user_info_server_proto_rawDesc = nil
	file_api_user_info_server_user_info_server_proto_goTypes = nil
	file_api_user_info_server_user_info_server_proto_depIdxs = nil
}
