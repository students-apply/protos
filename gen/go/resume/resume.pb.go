// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: resume/resume.proto

package resume

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

type CreateResumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Resume *Resume `protobuf:"bytes,2,opt,name=resume,proto3" json:"resume,omitempty"`
}

func (x *CreateResumeRequest) Reset() {
	*x = CreateResumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResumeRequest) ProtoMessage() {}

func (x *CreateResumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResumeRequest.ProtoReflect.Descriptor instead.
func (*CreateResumeRequest) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{0}
}

func (x *CreateResumeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateResumeRequest) GetResume() *Resume {
	if x != nil {
		return x.Resume
	}
	return nil
}

type CreateResumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResumeId string `protobuf:"bytes,1,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
}

func (x *CreateResumeResponse) Reset() {
	*x = CreateResumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResumeResponse) ProtoMessage() {}

func (x *CreateResumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResumeResponse.ProtoReflect.Descriptor instead.
func (*CreateResumeResponse) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResumeResponse) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

type UpdateResumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResumeId string  `protobuf:"bytes,2,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	Resume   *Resume `protobuf:"bytes,3,opt,name=resume,proto3" json:"resume,omitempty"`
}

func (x *UpdateResumeRequest) Reset() {
	*x = UpdateResumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResumeRequest) ProtoMessage() {}

func (x *UpdateResumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResumeRequest.ProtoReflect.Descriptor instead.
func (*UpdateResumeRequest) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateResumeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateResumeRequest) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

func (x *UpdateResumeRequest) GetResume() *Resume {
	if x != nil {
		return x.Resume
	}
	return nil
}

type UpdateResumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resume *Resume `protobuf:"bytes,1,opt,name=resume,proto3" json:"resume,omitempty"`
}

func (x *UpdateResumeResponse) Reset() {
	*x = UpdateResumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResumeResponse) ProtoMessage() {}

func (x *UpdateResumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResumeResponse.ProtoReflect.Descriptor instead.
func (*UpdateResumeResponse) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateResumeResponse) GetResume() *Resume {
	if x != nil {
		return x.Resume
	}
	return nil
}

type GetResumeByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResumeId string `protobuf:"bytes,2,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
}

func (x *GetResumeByIDRequest) Reset() {
	*x = GetResumeByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResumeByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResumeByIDRequest) ProtoMessage() {}

func (x *GetResumeByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResumeByIDRequest.ProtoReflect.Descriptor instead.
func (*GetResumeByIDRequest) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{4}
}

func (x *GetResumeByIDRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetResumeByIDRequest) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

type GetResumeByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resume *Resume `protobuf:"bytes,4,opt,name=resume,proto3,oneof" json:"resume,omitempty"`
}

func (x *GetResumeByIDResponse) Reset() {
	*x = GetResumeByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResumeByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResumeByIDResponse) ProtoMessage() {}

func (x *GetResumeByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResumeByIDResponse.ProtoReflect.Descriptor instead.
func (*GetResumeByIDResponse) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{5}
}

func (x *GetResumeByIDResponse) GetResume() *Resume {
	if x != nil {
		return x.Resume
	}
	return nil
}

type GetUserResumesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserResumesRequest) Reset() {
	*x = GetUserResumesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResumesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResumesRequest) ProtoMessage() {}

func (x *GetUserResumesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResumesRequest.ProtoReflect.Descriptor instead.
func (*GetUserResumesRequest) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserResumesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserResumesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resumes []*Resume `protobuf:"bytes,1,rep,name=resumes,proto3" json:"resumes,omitempty"`
}

func (x *GetUserResumesResponse) Reset() {
	*x = GetUserResumesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResumesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResumesResponse) ProtoMessage() {}

func (x *GetUserResumesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResumesResponse.ProtoReflect.Descriptor instead.
func (*GetUserResumesResponse) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserResumesResponse) GetResumes() []*Resume {
	if x != nil {
		return x.Resumes
	}
	return nil
}

type DeleteResumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResumeId string `protobuf:"bytes,1,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
}

func (x *DeleteResumeRequest) Reset() {
	*x = DeleteResumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResumeRequest) ProtoMessage() {}

func (x *DeleteResumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResumeRequest.ProtoReflect.Descriptor instead.
func (*DeleteResumeRequest) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteResumeRequest) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

type DeleteResumeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResumeResponse) Reset() {
	*x = DeleteResumeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResumeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResumeResponse) ProtoMessage() {}

func (x *DeleteResumeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResumeResponse.ProtoReflect.Descriptor instead.
func (*DeleteResumeResponse) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{9}
}

type Resume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string                 `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   string                 `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	MiddleName *string                `protobuf:"bytes,3,opt,name=middle_name,json=middleName,proto3,oneof" json:"middle_name,omitempty"`
	Birthday   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Email      string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone      *string                `protobuf:"bytes,6,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	CityId     int32                  `protobuf:"varint,7,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
}

func (x *Resume) Reset() {
	*x = Resume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resume) ProtoMessage() {}

func (x *Resume) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resume.ProtoReflect.Descriptor instead.
func (*Resume) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{10}
}

func (x *Resume) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Resume) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Resume) GetMiddleName() string {
	if x != nil && x.MiddleName != nil {
		return *x.MiddleName
	}
	return ""
}

func (x *Resume) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *Resume) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Resume) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *Resume) GetCityId() int32 {
	if x != nil {
		return x.CityId
	}
	return 0
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResumeId    string `protobuf:"bytes,1,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	Chunk       []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_resume_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_resume_resume_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_resume_resume_proto_rawDescGZIP(), []int{11}
}

func (x *Image) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

func (x *Image) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *Image) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_resume_resume_proto protoreflect.FileDescriptor

var file_resume_resume_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x73, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x22, 0x3e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x22, 0x4c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x4f,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x22,
	0x30, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x42, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x86, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x36,
	0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x71, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x8f, 0x03,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x49, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12,
	0x1b, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_resume_resume_proto_rawDescOnce sync.Once
	file_resume_resume_proto_rawDescData = file_resume_resume_proto_rawDesc
)

func file_resume_resume_proto_rawDescGZIP() []byte {
	file_resume_resume_proto_rawDescOnce.Do(func() {
		file_resume_resume_proto_rawDescData = protoimpl.X.CompressGZIP(file_resume_resume_proto_rawDescData)
	})
	return file_resume_resume_proto_rawDescData
}

var file_resume_resume_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_resume_resume_proto_goTypes = []interface{}{
	(*CreateResumeRequest)(nil),    // 0: resume.CreateResumeRequest
	(*CreateResumeResponse)(nil),   // 1: resume.CreateResumeResponse
	(*UpdateResumeRequest)(nil),    // 2: resume.UpdateResumeRequest
	(*UpdateResumeResponse)(nil),   // 3: resume.UpdateResumeResponse
	(*GetResumeByIDRequest)(nil),   // 4: resume.GetResumeByIDRequest
	(*GetResumeByIDResponse)(nil),  // 5: resume.GetResumeByIDResponse
	(*GetUserResumesRequest)(nil),  // 6: resume.GetUserResumesRequest
	(*GetUserResumesResponse)(nil), // 7: resume.GetUserResumesResponse
	(*DeleteResumeRequest)(nil),    // 8: resume.DeleteResumeRequest
	(*DeleteResumeResponse)(nil),   // 9: resume.DeleteResumeResponse
	(*Resume)(nil),                 // 10: resume.Resume
	(*Image)(nil),                  // 11: resume.Image
	(*timestamppb.Timestamp)(nil),  // 12: google.protobuf.Timestamp
}
var file_resume_resume_proto_depIdxs = []int32{
	10, // 0: resume.CreateResumeRequest.resume:type_name -> resume.Resume
	10, // 1: resume.UpdateResumeRequest.resume:type_name -> resume.Resume
	10, // 2: resume.UpdateResumeResponse.resume:type_name -> resume.Resume
	10, // 3: resume.GetResumeByIDResponse.resume:type_name -> resume.Resume
	10, // 4: resume.GetUserResumesResponse.resumes:type_name -> resume.Resume
	12, // 5: resume.Resume.birthday:type_name -> google.protobuf.Timestamp
	0,  // 6: resume.ResumeService.CreateResume:input_type -> resume.CreateResumeRequest
	2,  // 7: resume.ResumeService.UpdateResume:input_type -> resume.UpdateResumeRequest
	4,  // 8: resume.ResumeService.GetResumeByID:input_type -> resume.GetResumeByIDRequest
	6,  // 9: resume.ResumeService.GetUserResumes:input_type -> resume.GetUserResumesRequest
	8,  // 10: resume.ResumeService.DeleteResume:input_type -> resume.DeleteResumeRequest
	1,  // 11: resume.ResumeService.CreateResume:output_type -> resume.CreateResumeResponse
	3,  // 12: resume.ResumeService.UpdateResume:output_type -> resume.UpdateResumeResponse
	5,  // 13: resume.ResumeService.GetResumeByID:output_type -> resume.GetResumeByIDResponse
	7,  // 14: resume.ResumeService.GetUserResumes:output_type -> resume.GetUserResumesResponse
	9,  // 15: resume.ResumeService.DeleteResume:output_type -> resume.DeleteResumeResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_resume_resume_proto_init() }
func file_resume_resume_proto_init() {
	if File_resume_resume_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resume_resume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResumeRequest); i {
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
		file_resume_resume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResumeResponse); i {
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
		file_resume_resume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResumeRequest); i {
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
		file_resume_resume_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResumeResponse); i {
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
		file_resume_resume_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResumeByIDRequest); i {
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
		file_resume_resume_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResumeByIDResponse); i {
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
		file_resume_resume_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResumesRequest); i {
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
		file_resume_resume_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResumesResponse); i {
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
		file_resume_resume_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResumeRequest); i {
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
		file_resume_resume_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResumeResponse); i {
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
		file_resume_resume_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resume); i {
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
		file_resume_resume_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
	file_resume_resume_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_resume_resume_proto_msgTypes[10].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resume_resume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resume_resume_proto_goTypes,
		DependencyIndexes: file_resume_resume_proto_depIdxs,
		MessageInfos:      file_resume_resume_proto_msgTypes,
	}.Build()
	File_resume_resume_proto = out.File
	file_resume_resume_proto_rawDesc = nil
	file_resume_resume_proto_goTypes = nil
	file_resume_resume_proto_depIdxs = nil
}
