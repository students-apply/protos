// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: experience/experience.proto

package experience

import (
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

type CreateExperienceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResumeId   string      `protobuf:"bytes,1,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	UserId     string      `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Experience *Experience `protobuf:"bytes,3,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *CreateExperienceRequest) Reset() {
	*x = CreateExperienceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experience_experience_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExperienceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExperienceRequest) ProtoMessage() {}

func (x *CreateExperienceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_experience_experience_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExperienceRequest.ProtoReflect.Descriptor instead.
func (*CreateExperienceRequest) Descriptor() ([]byte, []int) {
	return file_experience_experience_proto_rawDescGZIP(), []int{0}
}

func (x *CreateExperienceRequest) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

func (x *CreateExperienceRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateExperienceRequest) GetExperience() *Experience {
	if x != nil {
		return x.Experience
	}
	return nil
}

type CreateExperienceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExperienceId string `protobuf:"bytes,1,opt,name=experience_id,json=experienceId,proto3" json:"experience_id,omitempty"`
}

func (x *CreateExperienceResponse) Reset() {
	*x = CreateExperienceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experience_experience_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExperienceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExperienceResponse) ProtoMessage() {}

func (x *CreateExperienceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_experience_experience_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExperienceResponse.ProtoReflect.Descriptor instead.
func (*CreateExperienceResponse) Descriptor() ([]byte, []int) {
	return file_experience_experience_proto_rawDescGZIP(), []int{1}
}

func (x *CreateExperienceResponse) GetExperienceId() string {
	if x != nil {
		return x.ExperienceId
	}
	return ""
}

type UpdateExperienceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExperienceId string      `protobuf:"bytes,1,opt,name=experience_id,json=experienceId,proto3" json:"experience_id,omitempty"`
	UserId       string      `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Experience   *Experience `protobuf:"bytes,3,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *UpdateExperienceRequest) Reset() {
	*x = UpdateExperienceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experience_experience_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExperienceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExperienceRequest) ProtoMessage() {}

func (x *UpdateExperienceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_experience_experience_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExperienceRequest.ProtoReflect.Descriptor instead.
func (*UpdateExperienceRequest) Descriptor() ([]byte, []int) {
	return file_experience_experience_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateExperienceRequest) GetExperienceId() string {
	if x != nil {
		return x.ExperienceId
	}
	return ""
}

func (x *UpdateExperienceRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateExperienceRequest) GetExperience() *Experience {
	if x != nil {
		return x.Experience
	}
	return nil
}

type GetExperienceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExperienceId string `protobuf:"bytes,1,opt,name=experience_id,json=experienceId,proto3" json:"experience_id,omitempty"`
}

func (x *GetExperienceRequest) Reset() {
	*x = GetExperienceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experience_experience_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExperienceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExperienceRequest) ProtoMessage() {}

func (x *GetExperienceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_experience_experience_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExperienceRequest.ProtoReflect.Descriptor instead.
func (*GetExperienceRequest) Descriptor() ([]byte, []int) {
	return file_experience_experience_proto_rawDescGZIP(), []int{3}
}

func (x *GetExperienceRequest) GetExperienceId() string {
	if x != nil {
		return x.ExperienceId
	}
	return ""
}

type DeleteExperienceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExperienceId string `protobuf:"bytes,1,opt,name=experience_id,json=experienceId,proto3" json:"experience_id,omitempty"`
	UserId       string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteExperienceRequest) Reset() {
	*x = DeleteExperienceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experience_experience_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExperienceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExperienceRequest) ProtoMessage() {}

func (x *DeleteExperienceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_experience_experience_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExperienceRequest.ProtoReflect.Descriptor instead.
func (*DeleteExperienceRequest) Descriptor() ([]byte, []int) {
	return file_experience_experience_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteExperienceRequest) GetExperienceId() string {
	if x != nil {
		return x.ExperienceId
	}
	return ""
}

func (x *DeleteExperienceRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteExperienceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteExperienceResponse) Reset() {
	*x = DeleteExperienceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experience_experience_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExperienceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExperienceResponse) ProtoMessage() {}

func (x *DeleteExperienceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_experience_experience_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExperienceResponse.ProtoReflect.Descriptor instead.
func (*DeleteExperienceResponse) Descriptor() ([]byte, []int) {
	return file_experience_experience_proto_rawDescGZIP(), []int{5}
}

type Experience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResumeId    string                 `protobuf:"bytes,1,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	Position    string                 `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Company     string                 `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	StartDate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Description string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Experience) Reset() {
	*x = Experience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experience_experience_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Experience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Experience) ProtoMessage() {}

func (x *Experience) ProtoReflect() protoreflect.Message {
	mi := &file_experience_experience_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Experience.ProtoReflect.Descriptor instead.
func (*Experience) Descriptor() ([]byte, []int) {
	return file_experience_experience_proto_rawDescGZIP(), []int{6}
}

func (x *Experience) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

func (x *Experience) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Experience) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Experience) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Experience) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Experience) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_experience_experience_proto protoreflect.FileDescriptor

var file_experience_experience_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x22, 0x3f, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x36, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x3b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1a, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xf3, 0x01, 0x0a, 0x0a, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0xed, 0x02, 0x0a, 0x11, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x5d, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x3b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_experience_experience_proto_rawDescOnce sync.Once
	file_experience_experience_proto_rawDescData = file_experience_experience_proto_rawDesc
)

func file_experience_experience_proto_rawDescGZIP() []byte {
	file_experience_experience_proto_rawDescOnce.Do(func() {
		file_experience_experience_proto_rawDescData = protoimpl.X.CompressGZIP(file_experience_experience_proto_rawDescData)
	})
	return file_experience_experience_proto_rawDescData
}

var file_experience_experience_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_experience_experience_proto_goTypes = []interface{}{
	(*CreateExperienceRequest)(nil),  // 0: experience.CreateExperienceRequest
	(*CreateExperienceResponse)(nil), // 1: experience.CreateExperienceResponse
	(*UpdateExperienceRequest)(nil),  // 2: experience.UpdateExperienceRequest
	(*GetExperienceRequest)(nil),     // 3: experience.GetExperienceRequest
	(*DeleteExperienceRequest)(nil),  // 4: experience.DeleteExperienceRequest
	(*DeleteExperienceResponse)(nil), // 5: experience.DeleteExperienceResponse
	(*Experience)(nil),               // 6: experience.Experience
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
}
var file_experience_experience_proto_depIdxs = []int32{
	6, // 0: experience.CreateExperienceRequest.experience:type_name -> experience.Experience
	6, // 1: experience.UpdateExperienceRequest.experience:type_name -> experience.Experience
	7, // 2: experience.Experience.start_date:type_name -> google.protobuf.Timestamp
	7, // 3: experience.Experience.end_date:type_name -> google.protobuf.Timestamp
	0, // 4: experience.ExperienceService.CreateExperience:input_type -> experience.CreateExperienceRequest
	2, // 5: experience.ExperienceService.UpdateExperience:input_type -> experience.UpdateExperienceRequest
	3, // 6: experience.ExperienceService.GetExperience:input_type -> experience.GetExperienceRequest
	4, // 7: experience.ExperienceService.DeleteExperience:input_type -> experience.DeleteExperienceRequest
	1, // 8: experience.ExperienceService.CreateExperience:output_type -> experience.CreateExperienceResponse
	6, // 9: experience.ExperienceService.UpdateExperience:output_type -> experience.Experience
	6, // 10: experience.ExperienceService.GetExperience:output_type -> experience.Experience
	5, // 11: experience.ExperienceService.DeleteExperience:output_type -> experience.DeleteExperienceResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_experience_experience_proto_init() }
func file_experience_experience_proto_init() {
	if File_experience_experience_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_experience_experience_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExperienceRequest); i {
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
		file_experience_experience_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExperienceResponse); i {
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
		file_experience_experience_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExperienceRequest); i {
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
		file_experience_experience_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExperienceRequest); i {
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
		file_experience_experience_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteExperienceRequest); i {
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
		file_experience_experience_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteExperienceResponse); i {
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
		file_experience_experience_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Experience); i {
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
			RawDescriptor: file_experience_experience_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_experience_experience_proto_goTypes,
		DependencyIndexes: file_experience_experience_proto_depIdxs,
		MessageInfos:      file_experience_experience_proto_msgTypes,
	}.Build()
	File_experience_experience_proto = out.File
	file_experience_experience_proto_rawDesc = nil
	file_experience_experience_proto_goTypes = nil
	file_experience_experience_proto_depIdxs = nil
}