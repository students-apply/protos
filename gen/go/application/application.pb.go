// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: application/application.proto

package application

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

type CreateApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResumeId     string `protobuf:"bytes,2,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	CoverLetter  string `protobuf:"bytes,3,opt,name=cover_letter,json=coverLetter,proto3" json:"cover_letter,omitempty"`
	InternshipId string `protobuf:"bytes,4,opt,name=internship_id,json=internshipId,proto3" json:"internship_id,omitempty"`
}

func (x *CreateApplicationRequest) Reset() {
	*x = CreateApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApplicationRequest) ProtoMessage() {}

func (x *CreateApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApplicationRequest.ProtoReflect.Descriptor instead.
func (*CreateApplicationRequest) Descriptor() ([]byte, []int) {
	return file_application_application_proto_rawDescGZIP(), []int{0}
}

func (x *CreateApplicationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateApplicationRequest) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

func (x *CreateApplicationRequest) GetCoverLetter() string {
	if x != nil {
		return x.CoverLetter
	}
	return ""
}

func (x *CreateApplicationRequest) GetInternshipId() string {
	if x != nil {
		return x.InternshipId
	}
	return ""
}

type CreateApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateApplicationResponse) Reset() {
	*x = CreateApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApplicationResponse) ProtoMessage() {}

func (x *CreateApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_application_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApplicationResponse.ProtoReflect.Descriptor instead.
func (*CreateApplicationResponse) Descriptor() ([]byte, []int) {
	return file_application_application_proto_rawDescGZIP(), []int{1}
}

func (x *CreateApplicationResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyId string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Status    string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateApplicationRequest) Reset() {
	*x = UpdateApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApplicationRequest) ProtoMessage() {}

func (x *UpdateApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApplicationRequest.ProtoReflect.Descriptor instead.
func (*UpdateApplicationRequest) Descriptor() ([]byte, []int) {
	return file_application_application_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateApplicationRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateApplicationRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *UpdateApplicationRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId       string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResumeId     string                 `protobuf:"bytes,3,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	InternshipId string                 `protobuf:"bytes,4,opt,name=internship_id,json=internshipId,proto3" json:"internship_id,omitempty"`
	Status       string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_application_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_application_application_proto_rawDescGZIP(), []int{3}
}

func (x *Application) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Application) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Application) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

func (x *Application) GetInternshipId() string {
	if x != nil {
		return x.InternshipId
	}
	return ""
}

func (x *Application) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Application) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Application) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type GetApplicationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetApplicationsRequest) Reset() {
	*x = GetApplicationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_application_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApplicationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApplicationsRequest) ProtoMessage() {}

func (x *GetApplicationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_application_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApplicationsRequest.ProtoReflect.Descriptor instead.
func (*GetApplicationsRequest) Descriptor() ([]byte, []int) {
	return file_application_application_proto_rawDescGZIP(), []int{4}
}

func (x *GetApplicationsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Applications struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications []*Application `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty"`
}

func (x *Applications) Reset() {
	*x = Applications{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_application_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Applications) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Applications) ProtoMessage() {}

func (x *Applications) ProtoReflect() protoreflect.Message {
	mi := &file_application_application_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Applications.ProtoReflect.Descriptor instead.
func (*Applications) Descriptor() ([]byte, []int) {
	return file_application_application_proto_rawDescGZIP(), []int{5}
}

func (x *Applications) GetApplications() []*Application {
	if x != nil {
		return x.Applications
	}
	return nil
}

type DeleteApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteApplicationRequest) Reset() {
	*x = DeleteApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_application_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApplicationRequest) ProtoMessage() {}

func (x *DeleteApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_application_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApplicationRequest.ProtoReflect.Descriptor instead.
func (*DeleteApplicationRequest) Descriptor() ([]byte, []int) {
	return file_application_application_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteApplicationRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteApplicationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteApplicationResponse) Reset() {
	*x = DeleteApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_application_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApplicationResponse) ProtoMessage() {}

func (x *DeleteApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_application_application_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApplicationResponse.ProtoReflect.Descriptor instead.
func (*DeleteApplicationResponse) Descriptor() ([]byte, []int) {
	return file_application_application_proto_rawDescGZIP(), []int{7}
}

var File_application_application_proto protoreflect.FileDescriptor

var file_application_application_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x4c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x86, 0x02, 0x0a, 0x0b, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x31, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x43, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x85, 0x03, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x51, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x54, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x2d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x3b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_application_application_proto_rawDescOnce sync.Once
	file_application_application_proto_rawDescData = file_application_application_proto_rawDesc
)

func file_application_application_proto_rawDescGZIP() []byte {
	file_application_application_proto_rawDescOnce.Do(func() {
		file_application_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_application_application_proto_rawDescData)
	})
	return file_application_application_proto_rawDescData
}

var file_application_application_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_application_application_proto_goTypes = []interface{}{
	(*CreateApplicationRequest)(nil),  // 0: application.CreateApplicationRequest
	(*CreateApplicationResponse)(nil), // 1: application.CreateApplicationResponse
	(*UpdateApplicationRequest)(nil),  // 2: application.UpdateApplicationRequest
	(*Application)(nil),               // 3: application.Application
	(*GetApplicationsRequest)(nil),    // 4: application.GetApplicationsRequest
	(*Applications)(nil),              // 5: application.Applications
	(*DeleteApplicationRequest)(nil),  // 6: application.DeleteApplicationRequest
	(*DeleteApplicationResponse)(nil), // 7: application.DeleteApplicationResponse
	(*timestamppb.Timestamp)(nil),     // 8: google.protobuf.Timestamp
}
var file_application_application_proto_depIdxs = []int32{
	8, // 0: application.Application.created_at:type_name -> google.protobuf.Timestamp
	8, // 1: application.Application.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: application.Applications.applications:type_name -> application.Application
	0, // 3: application.ApplicationService.CreateApplication:input_type -> application.CreateApplicationRequest
	4, // 4: application.ApplicationService.GetApplications:input_type -> application.GetApplicationsRequest
	2, // 5: application.ApplicationService.UpdateApplication:input_type -> application.UpdateApplicationRequest
	6, // 6: application.ApplicationService.DeleteApplication:input_type -> application.DeleteApplicationRequest
	1, // 7: application.ApplicationService.CreateApplication:output_type -> application.CreateApplicationResponse
	5, // 8: application.ApplicationService.GetApplications:output_type -> application.Applications
	3, // 9: application.ApplicationService.UpdateApplication:output_type -> application.Application
	7, // 10: application.ApplicationService.DeleteApplication:output_type -> application.DeleteApplicationResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_application_application_proto_init() }
func file_application_application_proto_init() {
	if File_application_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_application_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApplicationRequest); i {
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
		file_application_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApplicationResponse); i {
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
		file_application_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateApplicationRequest); i {
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
		file_application_application_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_application_application_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApplicationsRequest); i {
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
		file_application_application_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Applications); i {
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
		file_application_application_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApplicationRequest); i {
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
		file_application_application_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApplicationResponse); i {
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
			RawDescriptor: file_application_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_application_application_proto_goTypes,
		DependencyIndexes: file_application_application_proto_depIdxs,
		MessageInfos:      file_application_application_proto_msgTypes,
	}.Build()
	File_application_application_proto = out.File
	file_application_application_proto_rawDesc = nil
	file_application_application_proto_goTypes = nil
	file_application_application_proto_depIdxs = nil
}
