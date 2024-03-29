// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.8.0
// source: users.proto

package userspb

import (
	context "context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Ip       string `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{0}
}

func (x *SignupRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SignupRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SignupRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignupRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type SignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SignupResponse) Reset() {
	*x = SignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResponse) ProtoMessage() {}

func (x *SignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResponse.ProtoReflect.Descriptor instead.
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{1}
}

func (x *SignupResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	// Types that are assignable to Identifier:
	//	*LookupRequest_Email
	//	*LookupRequest_Phone
	Identifier isLookupRequest_Identifier `protobuf_oneof:"identifier"`
}

func (x *LookupRequest) Reset() {
	*x = LookupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupRequest) ProtoMessage() {}

func (x *LookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupRequest.ProtoReflect.Descriptor instead.
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{2}
}

func (x *LookupRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LookupRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (m *LookupRequest) GetIdentifier() isLookupRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *LookupRequest) GetEmail() string {
	if x, ok := x.GetIdentifier().(*LookupRequest_Email); ok {
		return x.Email
	}
	return ""
}

func (x *LookupRequest) GetPhone() string {
	if x, ok := x.GetIdentifier().(*LookupRequest_Phone); ok {
		return x.Phone
	}
	return ""
}

type isLookupRequest_Identifier interface {
	isLookupRequest_Identifier()
}

type LookupRequest_Email struct {
	Email string `protobuf:"bytes,3,opt,name=email,proto3,oneof"`
}

type LookupRequest_Phone struct {
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3,oneof"`
}

func (*LookupRequest_Email) isLookupRequest_Identifier() {}

func (*LookupRequest_Phone) isLookupRequest_Identifier() {}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          int32                `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Email         string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified bool                 `protobuf:"varint,4,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	Phone         string               `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	PhoneVerified bool                 `protobuf:"varint,6,opt,name=phone_verified,json=phoneVerified,proto3" json:"phone_verified,omitempty"`
	Password      string               `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	SignupIp      string               `protobuf:"bytes,8,opt,name=signup_ip,json=signupIp,proto3" json:"signup_ip,omitempty"`
	CreatedAt     *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[3]
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
	return file_users_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetPhoneVerified() bool {
	if x != nil {
		return x.PhoneVerified
	}
	return false
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetSignupIp() string {
	if x != nil {
		return x.SignupIp
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type RequestPasswordResetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserType  int32  `protobuf:"varint,1,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Ip        string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	UserAgent string `protobuf:"bytes,5,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
}

func (x *RequestPasswordResetRequest) Reset() {
	*x = RequestPasswordResetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPasswordResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPasswordResetRequest) ProtoMessage() {}

func (x *RequestPasswordResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPasswordResetRequest.ProtoReflect.Descriptor instead.
func (*RequestPasswordResetRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{4}
}

func (x *RequestPasswordResetRequest) GetUserType() int32 {
	if x != nil {
		return x.UserType
	}
	return 0
}

func (x *RequestPasswordResetRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RequestPasswordResetRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RequestPasswordResetRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *RequestPasswordResetRequest) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

type RequestPasswordResetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrantId     string               `protobuf:"bytes,1,opt,name=grant_id,json=grantId,proto3" json:"grant_id,omitempty"`
	NextOtpSend *timestamp.Timestamp `protobuf:"bytes,2,opt,name=next_otp_send,json=nextOtpSend,proto3" json:"next_otp_send,omitempty"`
}

func (x *RequestPasswordResetResponse) Reset() {
	*x = RequestPasswordResetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPasswordResetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPasswordResetResponse) ProtoMessage() {}

func (x *RequestPasswordResetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPasswordResetResponse.ProtoReflect.Descriptor instead.
func (*RequestPasswordResetResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{5}
}

func (x *RequestPasswordResetResponse) GetGrantId() string {
	if x != nil {
		return x.GrantId
	}
	return ""
}

func (x *RequestPasswordResetResponse) GetNextOtpSend() *timestamp.Timestamp {
	if x != nil {
		return x.NextOtpSend
	}
	return nil
}

type ResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrantId  string `protobuf:"bytes,1,opt,name=grant_id,json=grantId,proto3" json:"grant_id,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Ip       string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Otp      string `protobuf:"bytes,4,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *ResetPasswordRequest) Reset() {
	*x = ResetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordRequest) ProtoMessage() {}

func (x *ResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{6}
}

func (x *ResetPasswordRequest) GetGrantId() string {
	if x != nil {
		return x.GrantId
	}
	return ""
}

func (x *ResetPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ResetPasswordRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ResetPasswordRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type ResetPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetPasswordResponse) Reset() {
	*x = ResetPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordResponse) ProtoMessage() {}

func (x *ResetPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordResponse.ProtoReflect.Descriptor instead.
func (*ResetPasswordResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{7}
}

var File_users_proto protoreflect.FileDescriptor

var file_users_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x22, 0x20, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x71, 0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x98, 0x02, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x5f, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x49, 0x70, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x1b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x79, 0x0a, 0x1c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x6f, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74,
	0x4f, 0x74, 0x70, 0x53, 0x65, 0x6e, 0x64, 0x22, 0x6f, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xa6, 0x03, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x46, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x46, 0x6f, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x00, 0x12, 0x65, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x6d, 0x61,
	0x6e, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_proto_rawDescOnce sync.Once
	file_users_proto_rawDescData = file_users_proto_rawDesc
)

func file_users_proto_rawDescGZIP() []byte {
	file_users_proto_rawDescOnce.Do(func() {
		file_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_proto_rawDescData)
	})
	return file_users_proto_rawDescData
}

var file_users_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_users_proto_goTypes = []interface{}{
	(*SignupRequest)(nil),                // 0: userspb.SignupRequest
	(*SignupResponse)(nil),               // 1: userspb.SignupResponse
	(*LookupRequest)(nil),                // 2: userspb.LookupRequest
	(*User)(nil),                         // 3: userspb.User
	(*RequestPasswordResetRequest)(nil),  // 4: userspb.RequestPasswordResetRequest
	(*RequestPasswordResetResponse)(nil), // 5: userspb.RequestPasswordResetResponse
	(*ResetPasswordRequest)(nil),         // 6: userspb.ResetPasswordRequest
	(*ResetPasswordResponse)(nil),        // 7: userspb.ResetPasswordResponse
	(*timestamp.Timestamp)(nil),          // 8: google.protobuf.Timestamp
}
var file_users_proto_depIdxs = []int32{
	8, // 0: userspb.User.created_at:type_name -> google.protobuf.Timestamp
	8, // 1: userspb.RequestPasswordResetResponse.next_otp_send:type_name -> google.protobuf.Timestamp
	0, // 2: userspb.Users.Signup:input_type -> userspb.SignupRequest
	2, // 3: userspb.Users.Lookup:input_type -> userspb.LookupRequest
	2, // 4: userspb.Users.LookupForEmail:input_type -> userspb.LookupRequest
	2, // 5: userspb.Users.LookupForPhone:input_type -> userspb.LookupRequest
	4, // 6: userspb.Users.RequestPasswordReset:input_type -> userspb.RequestPasswordResetRequest
	6, // 7: userspb.Users.ResetPassword:input_type -> userspb.ResetPasswordRequest
	1, // 8: userspb.Users.Signup:output_type -> userspb.SignupResponse
	3, // 9: userspb.Users.Lookup:output_type -> userspb.User
	3, // 10: userspb.Users.LookupForEmail:output_type -> userspb.User
	3, // 11: userspb.Users.LookupForPhone:output_type -> userspb.User
	5, // 12: userspb.Users.RequestPasswordReset:output_type -> userspb.RequestPasswordResetResponse
	7, // 13: userspb.Users.ResetPassword:output_type -> userspb.ResetPasswordResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_users_proto_init() }
func file_users_proto_init() {
	if File_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupRequest); i {
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
		file_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupResponse); i {
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
		file_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupRequest); i {
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
		file_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPasswordResetRequest); i {
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
		file_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPasswordResetResponse); i {
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
		file_users_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordRequest); i {
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
		file_users_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordResponse); i {
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
	file_users_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*LookupRequest_Email)(nil),
		(*LookupRequest_Phone)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_proto_goTypes,
		DependencyIndexes: file_users_proto_depIdxs,
		MessageInfos:      file_users_proto_msgTypes,
	}.Build()
	File_users_proto = out.File
	file_users_proto_rawDesc = nil
	file_users_proto_goTypes = nil
	file_users_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*User, error)
	LookupForEmail(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*User, error)
	LookupForPhone(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*User, error)
	RequestPasswordReset(ctx context.Context, in *RequestPasswordResetRequest, opts ...grpc.CallOption) (*RequestPasswordResetResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/userspb.Users/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/userspb.Users/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) LookupForEmail(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/userspb.Users/LookupForEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) LookupForPhone(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/userspb.Users/LookupForPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) RequestPasswordReset(ctx context.Context, in *RequestPasswordResetRequest, opts ...grpc.CallOption) (*RequestPasswordResetResponse, error) {
	out := new(RequestPasswordResetResponse)
	err := c.cc.Invoke(ctx, "/userspb.Users/RequestPasswordReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/userspb.Users/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	Lookup(context.Context, *LookupRequest) (*User, error)
	LookupForEmail(context.Context, *LookupRequest) (*User, error)
	LookupForPhone(context.Context, *LookupRequest) (*User, error)
	RequestPasswordReset(context.Context, *RequestPasswordResetRequest) (*RequestPasswordResetResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (*UnimplementedUsersServer) Lookup(context.Context, *LookupRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}
func (*UnimplementedUsersServer) LookupForEmail(context.Context, *LookupRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupForEmail not implemented")
}
func (*UnimplementedUsersServer) LookupForPhone(context.Context, *LookupRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupForPhone not implemented")
}
func (*UnimplementedUsersServer) RequestPasswordReset(context.Context, *RequestPasswordResetRequest) (*RequestPasswordResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPasswordReset not implemented")
}
func (*UnimplementedUsersServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Users/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Users/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_LookupForEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).LookupForEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Users/LookupForEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).LookupForEmail(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_LookupForPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).LookupForPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Users/LookupForPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).LookupForPhone(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_RequestPasswordReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPasswordResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).RequestPasswordReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Users/RequestPasswordReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).RequestPasswordReset(ctx, req.(*RequestPasswordResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userspb.Users/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userspb.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _Users_Signup_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _Users_Lookup_Handler,
		},
		{
			MethodName: "LookupForEmail",
			Handler:    _Users_LookupForEmail_Handler,
		},
		{
			MethodName: "LookupForPhone",
			Handler:    _Users_LookupForPhone_Handler,
		},
		{
			MethodName: "RequestPasswordReset",
			Handler:    _Users_RequestPasswordReset_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _Users_ResetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
