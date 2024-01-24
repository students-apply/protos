// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: internship/internship.proto

package internship

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

type CreateInternshipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId   string   `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	City        string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	IsPaid      bool     `protobuf:"varint,5,opt,name=is_paid,json=isPaid,proto3" json:"is_paid,omitempty"`
	MaxInterns  int32    `protobuf:"varint,6,opt,name=max_interns,json=maxInterns,proto3" json:"max_interns,omitempty"`
	Keywords    []string `protobuf:"bytes,7,rep,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *CreateInternshipRequest) Reset() {
	*x = CreateInternshipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internship_internship_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInternshipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInternshipRequest) ProtoMessage() {}

func (x *CreateInternshipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internship_internship_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInternshipRequest.ProtoReflect.Descriptor instead.
func (*CreateInternshipRequest) Descriptor() ([]byte, []int) {
	return file_internship_internship_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInternshipRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CreateInternshipRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateInternshipRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateInternshipRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CreateInternshipRequest) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

func (x *CreateInternshipRequest) GetMaxInterns() int32 {
	if x != nil {
		return x.MaxInterns
	}
	return 0
}

func (x *CreateInternshipRequest) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

type CreateInternshipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateInternshipResponse) Reset() {
	*x = CreateInternshipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internship_internship_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInternshipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInternshipResponse) ProtoMessage() {}

func (x *CreateInternshipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internship_internship_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInternshipResponse.ProtoReflect.Descriptor instead.
func (*CreateInternshipResponse) Descriptor() ([]byte, []int) {
	return file_internship_internship_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInternshipResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetInternshipsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor     string `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CompanyId  string `protobuf:"bytes,3,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	IsPaid     bool   `protobuf:"varint,4,opt,name=is_paid,json=isPaid,proto3" json:"is_paid,omitempty"`
	IsActive   bool   `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	MaxInterns int32  `protobuf:"varint,6,opt,name=max_interns,json=maxInterns,proto3" json:"max_interns,omitempty"`
	Keyword    string `protobuf:"bytes,7,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *GetInternshipsRequest) Reset() {
	*x = GetInternshipsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internship_internship_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInternshipsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInternshipsRequest) ProtoMessage() {}

func (x *GetInternshipsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internship_internship_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInternshipsRequest.ProtoReflect.Descriptor instead.
func (*GetInternshipsRequest) Descriptor() ([]byte, []int) {
	return file_internship_internship_proto_rawDescGZIP(), []int{2}
}

func (x *GetInternshipsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *GetInternshipsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetInternshipsRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *GetInternshipsRequest) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

func (x *GetInternshipsRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *GetInternshipsRequest) GetMaxInterns() int32 {
	if x != nil {
		return x.MaxInterns
	}
	return 0
}

func (x *GetInternshipsRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type GetInternshipsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Internships []*Internship `protobuf:"bytes,1,rep,name=internships,proto3" json:"internships,omitempty"`
	Cursor      string        `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *GetInternshipsResponse) Reset() {
	*x = GetInternshipsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internship_internship_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInternshipsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInternshipsResponse) ProtoMessage() {}

func (x *GetInternshipsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internship_internship_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInternshipsResponse.ProtoReflect.Descriptor instead.
func (*GetInternshipsResponse) Descriptor() ([]byte, []int) {
	return file_internship_internship_proto_rawDescGZIP(), []int{3}
}

func (x *GetInternshipsResponse) GetInternships() []*Internship {
	if x != nil {
		return x.Internships
	}
	return nil
}

func (x *GetInternshipsResponse) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type Internship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CompanyId   string                 `protobuf:"bytes,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	City        string                 `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	IsPaid      bool                   `protobuf:"varint,6,opt,name=is_paid,json=isPaid,proto3" json:"is_paid,omitempty"`
	IsActive    bool                   `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	MaxInterns  int32                  `protobuf:"varint,8,opt,name=max_interns,json=maxInterns,proto3" json:"max_interns,omitempty"`
	Attachment  string                 `protobuf:"bytes,9,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Keywords    []string               `protobuf:"bytes,10,rep,name=keywords,proto3" json:"keywords,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Internship) Reset() {
	*x = Internship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internship_internship_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Internship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Internship) ProtoMessage() {}

func (x *Internship) ProtoReflect() protoreflect.Message {
	mi := &file_internship_internship_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Internship.ProtoReflect.Descriptor instead.
func (*Internship) Descriptor() ([]byte, []int) {
	return file_internship_internship_proto_rawDescGZIP(), []int{4}
}

func (x *Internship) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Internship) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Internship) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Internship) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *Internship) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Internship) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

func (x *Internship) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Internship) GetMaxInterns() int32 {
	if x != nil {
		return x.MaxInterns
	}
	return 0
}

func (x *Internship) GetAttachment() string {
	if x != nil {
		return x.Attachment
	}
	return ""
}

func (x *Internship) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *Internship) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type UpdateInternshipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId   string   `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Id          string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	City        string   `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	IsPaid      bool     `protobuf:"varint,6,opt,name=is_paid,json=isPaid,proto3" json:"is_paid,omitempty"`
	IsActive    bool     `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	MaxInterns  int32    `protobuf:"varint,8,opt,name=max_interns,json=maxInterns,proto3" json:"max_interns,omitempty"`
	Keywords    []string `protobuf:"bytes,9,rep,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *UpdateInternshipRequest) Reset() {
	*x = UpdateInternshipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internship_internship_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInternshipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInternshipRequest) ProtoMessage() {}

func (x *UpdateInternshipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internship_internship_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInternshipRequest.ProtoReflect.Descriptor instead.
func (*UpdateInternshipRequest) Descriptor() ([]byte, []int) {
	return file_internship_internship_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateInternshipRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *UpdateInternshipRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateInternshipRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateInternshipRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateInternshipRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UpdateInternshipRequest) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

func (x *UpdateInternshipRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *UpdateInternshipRequest) GetMaxInterns() int32 {
	if x != nil {
		return x.MaxInterns
	}
	return 0
}

func (x *UpdateInternshipRequest) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

type DeleteInternshipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteInternshipRequest) Reset() {
	*x = DeleteInternshipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internship_internship_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInternshipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInternshipRequest) ProtoMessage() {}

func (x *DeleteInternshipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internship_internship_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInternshipRequest.ProtoReflect.Descriptor instead.
func (*DeleteInternshipRequest) Descriptor() ([]byte, []int) {
	return file_internship_internship_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteInternshipRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *DeleteInternshipRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteInternshipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteInternshipResponse) Reset() {
	*x = DeleteInternshipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internship_internship_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInternshipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInternshipResponse) ProtoMessage() {}

func (x *DeleteInternshipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internship_internship_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInternshipResponse.ProtoReflect.Descriptor instead.
func (*DeleteInternshipResponse) Descriptor() ([]byte, []int) {
	return file_internship_internship_proto_rawDescGZIP(), []int{7}
}

var File_internship_internship_proto protoreflect.FileDescriptor

var file_internship_internship_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d,
	0x61, 0x78, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x2a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xd3, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0b,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x22, 0xd3, 0x02, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x73, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73,
	0x50, 0x61, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x85, 0x02, 0x0a, 0x17, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0x48, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfb, 0x02, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x21,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x5d, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2d, 0x61, 0x70, 0x70,
	0x6c, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internship_internship_proto_rawDescOnce sync.Once
	file_internship_internship_proto_rawDescData = file_internship_internship_proto_rawDesc
)

func file_internship_internship_proto_rawDescGZIP() []byte {
	file_internship_internship_proto_rawDescOnce.Do(func() {
		file_internship_internship_proto_rawDescData = protoimpl.X.CompressGZIP(file_internship_internship_proto_rawDescData)
	})
	return file_internship_internship_proto_rawDescData
}

var file_internship_internship_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internship_internship_proto_goTypes = []interface{}{
	(*CreateInternshipRequest)(nil),  // 0: internship.CreateInternshipRequest
	(*CreateInternshipResponse)(nil), // 1: internship.CreateInternshipResponse
	(*GetInternshipsRequest)(nil),    // 2: internship.GetInternshipsRequest
	(*GetInternshipsResponse)(nil),   // 3: internship.GetInternshipsResponse
	(*Internship)(nil),               // 4: internship.Internship
	(*UpdateInternshipRequest)(nil),  // 5: internship.UpdateInternshipRequest
	(*DeleteInternshipRequest)(nil),  // 6: internship.DeleteInternshipRequest
	(*DeleteInternshipResponse)(nil), // 7: internship.DeleteInternshipResponse
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
}
var file_internship_internship_proto_depIdxs = []int32{
	4, // 0: internship.GetInternshipsResponse.internships:type_name -> internship.Internship
	8, // 1: internship.Internship.created_at:type_name -> google.protobuf.Timestamp
	0, // 2: internship.InternshipService.CreateInternship:input_type -> internship.CreateInternshipRequest
	2, // 3: internship.InternshipService.GetInternships:input_type -> internship.GetInternshipsRequest
	5, // 4: internship.InternshipService.UpdateInternship:input_type -> internship.UpdateInternshipRequest
	6, // 5: internship.InternshipService.DeleteInternship:input_type -> internship.DeleteInternshipRequest
	1, // 6: internship.InternshipService.CreateInternship:output_type -> internship.CreateInternshipResponse
	3, // 7: internship.InternshipService.GetInternships:output_type -> internship.GetInternshipsResponse
	4, // 8: internship.InternshipService.UpdateInternship:output_type -> internship.Internship
	7, // 9: internship.InternshipService.DeleteInternship:output_type -> internship.DeleteInternshipResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internship_internship_proto_init() }
func file_internship_internship_proto_init() {
	if File_internship_internship_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internship_internship_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInternshipRequest); i {
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
		file_internship_internship_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInternshipResponse); i {
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
		file_internship_internship_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInternshipsRequest); i {
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
		file_internship_internship_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInternshipsResponse); i {
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
		file_internship_internship_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Internship); i {
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
		file_internship_internship_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInternshipRequest); i {
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
		file_internship_internship_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInternshipRequest); i {
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
		file_internship_internship_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInternshipResponse); i {
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
			RawDescriptor: file_internship_internship_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internship_internship_proto_goTypes,
		DependencyIndexes: file_internship_internship_proto_depIdxs,
		MessageInfos:      file_internship_internship_proto_msgTypes,
	}.Build()
	File_internship_internship_proto = out.File
	file_internship_internship_proto_rawDesc = nil
	file_internship_internship_proto_goTypes = nil
	file_internship_internship_proto_depIdxs = nil
}
