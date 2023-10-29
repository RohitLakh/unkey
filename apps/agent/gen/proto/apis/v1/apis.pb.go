// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/apis/v1/apis.proto

package apisv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type AuthType int32

const (
	AuthType_AUTH_TYPE_UNSPECIFIED AuthType = 0
	AuthType_AUTH_TYPE_KEY         AuthType = 1
	AuthType_AUTH_TYPE_JWT         AuthType = 2
)

// Enum value maps for AuthType.
var (
	AuthType_name = map[int32]string{
		0: "AUTH_TYPE_UNSPECIFIED",
		1: "AUTH_TYPE_KEY",
		2: "AUTH_TYPE_JWT",
	}
	AuthType_value = map[string]int32{
		"AUTH_TYPE_UNSPECIFIED": 0,
		"AUTH_TYPE_KEY":         1,
		"AUTH_TYPE_JWT":         2,
	}
)

func (x AuthType) Enum() *AuthType {
	p := new(AuthType)
	*p = x
	return p
}

func (x AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_apis_v1_apis_proto_enumTypes[0].Descriptor()
}

func (AuthType) Type() protoreflect.EnumType {
	return &file_proto_apis_v1_apis_proto_enumTypes[0]
}

func (x AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthType.Descriptor instead.
func (AuthType) EnumDescriptor() ([]byte, []int) {
	return file_proto_apis_v1_apis_proto_rawDescGZIP(), []int{0}
}

type CreateApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceId string `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateApiRequest) Reset() {
	*x = CreateApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apis_v1_apis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiRequest) ProtoMessage() {}

func (x *CreateApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apis_v1_apis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiRequest.ProtoReflect.Descriptor instead.
func (*CreateApiRequest) Descriptor() ([]byte, []int) {
	return file_proto_apis_v1_apis_proto_rawDescGZIP(), []int{0}
}

func (x *CreateApiRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *CreateApiRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateApiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiId string `protobuf:"bytes,1,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"`
}

func (x *CreateApiResponse) Reset() {
	*x = CreateApiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apis_v1_apis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiResponse) ProtoMessage() {}

func (x *CreateApiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apis_v1_apis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiResponse.ProtoReflect.Descriptor instead.
func (*CreateApiResponse) Descriptor() ([]byte, []int) {
	return file_proto_apis_v1_apis_proto_rawDescGZIP(), []int{1}
}

func (x *CreateApiResponse) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

type Api struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiId       string   `protobuf:"bytes,1,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"`
	WorkspaceId string   `protobuf:"bytes,2,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreateAt    int64    `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	DeletedAt   *int64   `protobuf:"varint,5,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
	IpWhitelist []string `protobuf:"bytes,6,rep,name=ip_whitelist,json=ipWhitelist,proto3" json:"ip_whitelist,omitempty"`
	AuthType    AuthType `protobuf:"varint,7,opt,name=auth_type,json=authType,proto3,enum=proto.apis.v1.AuthType" json:"auth_type,omitempty"`
	KeyAuthId   *string  `protobuf:"bytes,8,opt,name=key_auth_id,json=keyAuthId,proto3,oneof" json:"key_auth_id,omitempty"`
	JwtAuthId   *string  `protobuf:"bytes,9,opt,name=jwt_auth_id,json=jwtAuthId,proto3,oneof" json:"jwt_auth_id,omitempty"`
}

func (x *Api) Reset() {
	*x = Api{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apis_v1_apis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Api) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Api) ProtoMessage() {}

func (x *Api) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apis_v1_apis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Api.ProtoReflect.Descriptor instead.
func (*Api) Descriptor() ([]byte, []int) {
	return file_proto_apis_v1_apis_proto_rawDescGZIP(), []int{2}
}

func (x *Api) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

func (x *Api) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *Api) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Api) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Api) GetDeletedAt() int64 {
	if x != nil && x.DeletedAt != nil {
		return *x.DeletedAt
	}
	return 0
}

func (x *Api) GetIpWhitelist() []string {
	if x != nil {
		return x.IpWhitelist
	}
	return nil
}

func (x *Api) GetAuthType() AuthType {
	if x != nil {
		return x.AuthType
	}
	return AuthType_AUTH_TYPE_UNSPECIFIED
}

func (x *Api) GetKeyAuthId() string {
	if x != nil && x.KeyAuthId != nil {
		return *x.KeyAuthId
	}
	return ""
}

func (x *Api) GetJwtAuthId() string {
	if x != nil && x.JwtAuthId != nil {
		return *x.JwtAuthId
	}
	return ""
}

type DeleteApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiId string `protobuf:"bytes,1,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"`
}

func (x *DeleteApiRequest) Reset() {
	*x = DeleteApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apis_v1_apis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApiRequest) ProtoMessage() {}

func (x *DeleteApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apis_v1_apis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApiRequest.ProtoReflect.Descriptor instead.
func (*DeleteApiRequest) Descriptor() ([]byte, []int) {
	return file_proto_apis_v1_apis_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteApiRequest) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

type DeleteApiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteApiResponse) Reset() {
	*x = DeleteApiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apis_v1_apis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApiResponse) ProtoMessage() {}

func (x *DeleteApiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apis_v1_apis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApiResponse.ProtoReflect.Descriptor instead.
func (*DeleteApiResponse) Descriptor() ([]byte, []int) {
	return file_proto_apis_v1_apis_proto_rawDescGZIP(), []int{4}
}

type FindApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiId string `protobuf:"bytes,1,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"`
}

func (x *FindApiRequest) Reset() {
	*x = FindApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apis_v1_apis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindApiRequest) ProtoMessage() {}

func (x *FindApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apis_v1_apis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindApiRequest.ProtoReflect.Descriptor instead.
func (*FindApiRequest) Descriptor() ([]byte, []int) {
	return file_proto_apis_v1_apis_proto_rawDescGZIP(), []int{5}
}

func (x *FindApiRequest) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

type FindApiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api *Api `protobuf:"bytes,1,opt,name=api,proto3,oneof" json:"api,omitempty"`
}

func (x *FindApiResponse) Reset() {
	*x = FindApiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apis_v1_apis_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindApiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindApiResponse) ProtoMessage() {}

func (x *FindApiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apis_v1_apis_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindApiResponse.ProtoReflect.Descriptor instead.
func (*FindApiResponse) Descriptor() ([]byte, []int) {
	return file_proto_apis_v1_apis_proto_rawDescGZIP(), []int{6}
}

func (x *FindApiResponse) GetApi() *Api {
	if x != nil {
		return x.Api
	}
	return nil
}

type ListApisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceId string `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
}

func (x *ListApisRequest) Reset() {
	*x = ListApisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apis_v1_apis_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApisRequest) ProtoMessage() {}

func (x *ListApisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apis_v1_apis_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApisRequest.ProtoReflect.Descriptor instead.
func (*ListApisRequest) Descriptor() ([]byte, []int) {
	return file_proto_apis_v1_apis_proto_rawDescGZIP(), []int{7}
}

func (x *ListApisRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type ListApisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apis []*Api `protobuf:"bytes,1,rep,name=apis,proto3" json:"apis,omitempty"`
}

func (x *ListApisResponse) Reset() {
	*x = ListApisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_apis_v1_apis_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApisResponse) ProtoMessage() {}

func (x *ListApisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_apis_v1_apis_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApisResponse.ProtoReflect.Descriptor instead.
func (*ListApisResponse) Descriptor() ([]byte, []int) {
	return file_proto_apis_v1_apis_proto_rawDescGZIP(), []int{8}
}

func (x *ListApisResponse) GetApis() []*Api {
	if x != nil {
		return x.Apis
	}
	return nil
}

var File_proto_apis_v1_apis_proto protoreflect.FileDescriptor

var file_proto_apis_v1_apis_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x22,
	0xf8, 0x02, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x1e, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba,
	0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x70, 0x5f, 0x77,
	0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x69, 0x70, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x23, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x41, 0x75, 0x74,
	0x68, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0b, 0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x6a,
	0x77, 0x74, 0x41, 0x75, 0x74, 0x68, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6a,
	0x77, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x06, 0x61, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x22, 0x13,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x30, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05,
	0x61, 0x70, 0x69, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x48, 0x00, 0x52, 0x03, 0x61, 0x70, 0x69,
	0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x3d, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x70, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x61, 0x70, 0x69, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69,
	0x52, 0x04, 0x61, 0x70, 0x69, 0x73, 0x2a, 0x4b, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x57,
	0x54, 0x10, 0x02, 0x32, 0xcb, 0x02, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x12,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70,
	0x69, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x70,
	0x69, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x73, 0x12, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x75, 0x6e, 0x6b, 0x65, 0x79, 0x65, 0x64, 0x2f, 0x75, 0x6e, 0x6b, 0x65, 0x79, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x73, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_apis_v1_apis_proto_rawDescOnce sync.Once
	file_proto_apis_v1_apis_proto_rawDescData = file_proto_apis_v1_apis_proto_rawDesc
)

func file_proto_apis_v1_apis_proto_rawDescGZIP() []byte {
	file_proto_apis_v1_apis_proto_rawDescOnce.Do(func() {
		file_proto_apis_v1_apis_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_apis_v1_apis_proto_rawDescData)
	})
	return file_proto_apis_v1_apis_proto_rawDescData
}

var file_proto_apis_v1_apis_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_apis_v1_apis_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_apis_v1_apis_proto_goTypes = []interface{}{
	(AuthType)(0),             // 0: proto.apis.v1.AuthType
	(*CreateApiRequest)(nil),  // 1: proto.apis.v1.CreateApiRequest
	(*CreateApiResponse)(nil), // 2: proto.apis.v1.CreateApiResponse
	(*Api)(nil),               // 3: proto.apis.v1.Api
	(*DeleteApiRequest)(nil),  // 4: proto.apis.v1.DeleteApiRequest
	(*DeleteApiResponse)(nil), // 5: proto.apis.v1.DeleteApiResponse
	(*FindApiRequest)(nil),    // 6: proto.apis.v1.FindApiRequest
	(*FindApiResponse)(nil),   // 7: proto.apis.v1.FindApiResponse
	(*ListApisRequest)(nil),   // 8: proto.apis.v1.ListApisRequest
	(*ListApisResponse)(nil),  // 9: proto.apis.v1.ListApisResponse
}
var file_proto_apis_v1_apis_proto_depIdxs = []int32{
	0, // 0: proto.apis.v1.Api.auth_type:type_name -> proto.apis.v1.AuthType
	3, // 1: proto.apis.v1.FindApiResponse.api:type_name -> proto.apis.v1.Api
	3, // 2: proto.apis.v1.ListApisResponse.apis:type_name -> proto.apis.v1.Api
	1, // 3: proto.apis.v1.ApiService.CreateApi:input_type -> proto.apis.v1.CreateApiRequest
	4, // 4: proto.apis.v1.ApiService.DeleteApi:input_type -> proto.apis.v1.DeleteApiRequest
	6, // 5: proto.apis.v1.ApiService.FindApi:input_type -> proto.apis.v1.FindApiRequest
	8, // 6: proto.apis.v1.ApiService.ListApis:input_type -> proto.apis.v1.ListApisRequest
	2, // 7: proto.apis.v1.ApiService.CreateApi:output_type -> proto.apis.v1.CreateApiResponse
	5, // 8: proto.apis.v1.ApiService.DeleteApi:output_type -> proto.apis.v1.DeleteApiResponse
	7, // 9: proto.apis.v1.ApiService.FindApi:output_type -> proto.apis.v1.FindApiResponse
	9, // 10: proto.apis.v1.ApiService.ListApis:output_type -> proto.apis.v1.ListApisResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_apis_v1_apis_proto_init() }
func file_proto_apis_v1_apis_proto_init() {
	if File_proto_apis_v1_apis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_apis_v1_apis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiRequest); i {
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
		file_proto_apis_v1_apis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiResponse); i {
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
		file_proto_apis_v1_apis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Api); i {
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
		file_proto_apis_v1_apis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApiRequest); i {
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
		file_proto_apis_v1_apis_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApiResponse); i {
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
		file_proto_apis_v1_apis_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindApiRequest); i {
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
		file_proto_apis_v1_apis_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindApiResponse); i {
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
		file_proto_apis_v1_apis_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApisRequest); i {
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
		file_proto_apis_v1_apis_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApisResponse); i {
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
	file_proto_apis_v1_apis_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_proto_apis_v1_apis_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_apis_v1_apis_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_apis_v1_apis_proto_goTypes,
		DependencyIndexes: file_proto_apis_v1_apis_proto_depIdxs,
		EnumInfos:         file_proto_apis_v1_apis_proto_enumTypes,
		MessageInfos:      file_proto_apis_v1_apis_proto_msgTypes,
	}.Build()
	File_proto_apis_v1_apis_proto = out.File
	file_proto_apis_v1_apis_proto_rawDesc = nil
	file_proto_apis_v1_apis_proto_goTypes = nil
	file_proto_apis_v1_apis_proto_depIdxs = nil
}
