// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: skill/skill.proto

package skill

import (
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

type CreateSkillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResumeId string `protobuf:"bytes,1,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	Skill    *Skill `protobuf:"bytes,2,opt,name=skill,proto3" json:"skill,omitempty"`
}

func (x *CreateSkillRequest) Reset() {
	*x = CreateSkillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_skill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSkillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSkillRequest) ProtoMessage() {}

func (x *CreateSkillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skill_skill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSkillRequest.ProtoReflect.Descriptor instead.
func (*CreateSkillRequest) Descriptor() ([]byte, []int) {
	return file_skill_skill_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSkillRequest) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

func (x *CreateSkillRequest) GetSkill() *Skill {
	if x != nil {
		return x.Skill
	}
	return nil
}

type CreateSkillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId string `protobuf:"bytes,1,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
}

func (x *CreateSkillResponse) Reset() {
	*x = CreateSkillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_skill_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSkillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSkillResponse) ProtoMessage() {}

func (x *CreateSkillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_skill_skill_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSkillResponse.ProtoReflect.Descriptor instead.
func (*CreateSkillResponse) Descriptor() ([]byte, []int) {
	return file_skill_skill_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSkillResponse) GetSkillId() string {
	if x != nil {
		return x.SkillId
	}
	return ""
}

type UpdateSkillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId string `protobuf:"bytes,1,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Skill   *Skill `protobuf:"bytes,3,opt,name=skill,proto3" json:"skill,omitempty"`
}

func (x *UpdateSkillRequest) Reset() {
	*x = UpdateSkillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_skill_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSkillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSkillRequest) ProtoMessage() {}

func (x *UpdateSkillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skill_skill_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSkillRequest.ProtoReflect.Descriptor instead.
func (*UpdateSkillRequest) Descriptor() ([]byte, []int) {
	return file_skill_skill_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSkillRequest) GetSkillId() string {
	if x != nil {
		return x.SkillId
	}
	return ""
}

func (x *UpdateSkillRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateSkillRequest) GetSkill() *Skill {
	if x != nil {
		return x.Skill
	}
	return nil
}

type GetSkillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId string `protobuf:"bytes,1,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
}

func (x *GetSkillRequest) Reset() {
	*x = GetSkillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_skill_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSkillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSkillRequest) ProtoMessage() {}

func (x *GetSkillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skill_skill_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSkillRequest.ProtoReflect.Descriptor instead.
func (*GetSkillRequest) Descriptor() ([]byte, []int) {
	return file_skill_skill_proto_rawDescGZIP(), []int{3}
}

func (x *GetSkillRequest) GetSkillId() string {
	if x != nil {
		return x.SkillId
	}
	return ""
}

type DeleteSkillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId string `protobuf:"bytes,1,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteSkillRequest) Reset() {
	*x = DeleteSkillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_skill_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSkillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSkillRequest) ProtoMessage() {}

func (x *DeleteSkillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skill_skill_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSkillRequest.ProtoReflect.Descriptor instead.
func (*DeleteSkillRequest) Descriptor() ([]byte, []int) {
	return file_skill_skill_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSkillRequest) GetSkillId() string {
	if x != nil {
		return x.SkillId
	}
	return ""
}

func (x *DeleteSkillRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteSkillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSkillResponse) Reset() {
	*x = DeleteSkillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_skill_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSkillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSkillResponse) ProtoMessage() {}

func (x *DeleteSkillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_skill_skill_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSkillResponse.ProtoReflect.Descriptor instead.
func (*DeleteSkillResponse) Descriptor() ([]byte, []int) {
	return file_skill_skill_proto_rawDescGZIP(), []int{5}
}

type Skill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResumeId  string  `protobuf:"bytes,1,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	SkillName string  `protobuf:"bytes,2,opt,name=skill_name,json=skillName,proto3" json:"skill_name,omitempty"`
	Level     float32 `protobuf:"fixed32,3,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *Skill) Reset() {
	*x = Skill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_skill_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_skill_skill_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_skill_skill_proto_rawDescGZIP(), []int{6}
}

func (x *Skill) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

func (x *Skill) GetSkillName() string {
	if x != nil {
		return x.SkillName
	}
	return ""
}

func (x *Skill) GetLevel() float32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_skill_skill_proto protoreflect.FileDescriptor

var file_skill_skill_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x22, 0x55, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x22, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x22,
	0x48, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x59, 0x0a, 0x05, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x32, 0x84, 0x02, 0x0a, 0x0c,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x12, 0x19, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x30, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x44, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_skill_skill_proto_rawDescOnce sync.Once
	file_skill_skill_proto_rawDescData = file_skill_skill_proto_rawDesc
)

func file_skill_skill_proto_rawDescGZIP() []byte {
	file_skill_skill_proto_rawDescOnce.Do(func() {
		file_skill_skill_proto_rawDescData = protoimpl.X.CompressGZIP(file_skill_skill_proto_rawDescData)
	})
	return file_skill_skill_proto_rawDescData
}

var file_skill_skill_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_skill_skill_proto_goTypes = []interface{}{
	(*CreateSkillRequest)(nil),  // 0: skill.CreateSkillRequest
	(*CreateSkillResponse)(nil), // 1: skill.CreateSkillResponse
	(*UpdateSkillRequest)(nil),  // 2: skill.UpdateSkillRequest
	(*GetSkillRequest)(nil),     // 3: skill.GetSkillRequest
	(*DeleteSkillRequest)(nil),  // 4: skill.DeleteSkillRequest
	(*DeleteSkillResponse)(nil), // 5: skill.DeleteSkillResponse
	(*Skill)(nil),               // 6: skill.Skill
}
var file_skill_skill_proto_depIdxs = []int32{
	6, // 0: skill.CreateSkillRequest.skill:type_name -> skill.Skill
	6, // 1: skill.UpdateSkillRequest.skill:type_name -> skill.Skill
	0, // 2: skill.SkillService.CreateSkill:input_type -> skill.CreateSkillRequest
	2, // 3: skill.SkillService.UpdateSkill:input_type -> skill.UpdateSkillRequest
	3, // 4: skill.SkillService.GetSkill:input_type -> skill.GetSkillRequest
	4, // 5: skill.SkillService.DeleteSkill:input_type -> skill.DeleteSkillRequest
	1, // 6: skill.SkillService.CreateSkill:output_type -> skill.CreateSkillResponse
	6, // 7: skill.SkillService.UpdateSkill:output_type -> skill.Skill
	6, // 8: skill.SkillService.GetSkill:output_type -> skill.Skill
	5, // 9: skill.SkillService.DeleteSkill:output_type -> skill.DeleteSkillResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_skill_skill_proto_init() }
func file_skill_skill_proto_init() {
	if File_skill_skill_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skill_skill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSkillRequest); i {
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
		file_skill_skill_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSkillResponse); i {
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
		file_skill_skill_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSkillRequest); i {
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
		file_skill_skill_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSkillRequest); i {
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
		file_skill_skill_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSkillRequest); i {
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
		file_skill_skill_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSkillResponse); i {
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
		file_skill_skill_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skill); i {
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
			RawDescriptor: file_skill_skill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_skill_skill_proto_goTypes,
		DependencyIndexes: file_skill_skill_proto_depIdxs,
		MessageInfos:      file_skill_skill_proto_msgTypes,
	}.Build()
	File_skill_skill_proto = out.File
	file_skill_skill_proto_rawDesc = nil
	file_skill_skill_proto_goTypes = nil
	file_skill_skill_proto_depIdxs = nil
}