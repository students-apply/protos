// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.24.4
// source: view/view.proto

package view

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

type CreateViewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResumeId  string `protobuf:"bytes,1,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	CompanyId string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *CreateViewRequest) Reset() {
	*x = CreateViewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateViewRequest) ProtoMessage() {}

func (x *CreateViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_view_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateViewRequest.ProtoReflect.Descriptor instead.
func (*CreateViewRequest) Descriptor() ([]byte, []int) {
	return file_view_view_proto_rawDescGZIP(), []int{0}
}

func (x *CreateViewRequest) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

func (x *CreateViewRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

type CreateViewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ViewId string `protobuf:"bytes,1,opt,name=view_id,json=viewId,proto3" json:"view_id,omitempty"`
}

func (x *CreateViewResponse) Reset() {
	*x = CreateViewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_view_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateViewResponse) ProtoMessage() {}

func (x *CreateViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_view_view_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateViewResponse.ProtoReflect.Descriptor instead.
func (*CreateViewResponse) Descriptor() ([]byte, []int) {
	return file_view_view_proto_rawDescGZIP(), []int{1}
}

func (x *CreateViewResponse) GetViewId() string {
	if x != nil {
		return x.ViewId
	}
	return ""
}

type GetResumeViewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResumeId string `protobuf:"bytes,1,opt,name=resume_id,json=resumeId,proto3" json:"resume_id,omitempty"`
	LastDays *int32 `protobuf:"varint,2,opt,name=last_days,json=lastDays,proto3,oneof" json:"last_days,omitempty"`
}

func (x *GetResumeViewsRequest) Reset() {
	*x = GetResumeViewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_view_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResumeViewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResumeViewsRequest) ProtoMessage() {}

func (x *GetResumeViewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_view_view_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResumeViewsRequest.ProtoReflect.Descriptor instead.
func (*GetResumeViewsRequest) Descriptor() ([]byte, []int) {
	return file_view_view_proto_rawDescGZIP(), []int{2}
}

func (x *GetResumeViewsRequest) GetResumeId() string {
	if x != nil {
		return x.ResumeId
	}
	return ""
}

func (x *GetResumeViewsRequest) GetLastDays() int32 {
	if x != nil && x.LastDays != nil {
		return *x.LastDays
	}
	return 0
}

type GetResumeViewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Views []*View `protobuf:"bytes,1,rep,name=views,proto3" json:"views,omitempty"`
	Total int32   `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetResumeViewsResponse) Reset() {
	*x = GetResumeViewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_view_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResumeViewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResumeViewsResponse) ProtoMessage() {}

func (x *GetResumeViewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_view_view_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResumeViewsResponse.ProtoReflect.Descriptor instead.
func (*GetResumeViewsResponse) Descriptor() ([]byte, []int) {
	return file_view_view_proto_rawDescGZIP(), []int{3}
}

func (x *GetResumeViewsResponse) GetViews() []*View {
	if x != nil {
		return x.Views
	}
	return nil
}

func (x *GetResumeViewsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type View struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ViewId        string                 `protobuf:"bytes,1,opt,name=view_id,json=viewId,proto3" json:"view_id,omitempty"`
	CompanyId     string                 `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	CompanyName   string                 `protobuf:"bytes,3,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	CompanyAvatar string                 `protobuf:"bytes,4,opt,name=company_avatar,json=companyAvatar,proto3" json:"company_avatar,omitempty"`
	ViewedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=viewed_at,json=viewedAt,proto3" json:"viewed_at,omitempty"`
}

func (x *View) Reset() {
	*x = View{}
	if protoimpl.UnsafeEnabled {
		mi := &file_view_view_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *View) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*View) ProtoMessage() {}

func (x *View) ProtoReflect() protoreflect.Message {
	mi := &file_view_view_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use View.ProtoReflect.Descriptor instead.
func (*View) Descriptor() ([]byte, []int) {
	return file_view_view_proto_rawDescGZIP(), []int{4}
}

func (x *View) GetViewId() string {
	if x != nil {
		return x.ViewId
	}
	return ""
}

func (x *View) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *View) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *View) GetCompanyAvatar() string {
	if x != nil {
		return x.CompanyAvatar
	}
	return ""
}

func (x *View) GetViewedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ViewedAt
	}
	return nil
}

var File_view_view_proto protoreflect.FileDescriptor

var file_view_view_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x56, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x61, 0x79, 0x73, 0x88, 0x01, 0x01,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x22, 0x50,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x56, 0x69, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x56,
	0x69, 0x65, 0x77, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0xc1, 0x01, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x69, 0x65, 0x77,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x76,
	0x69, 0x65, 0x77, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x76, 0x69, 0x65, 0x77,
	0x65, 0x64, 0x41, 0x74, 0x32, 0x9b, 0x01, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x17, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x56, 0x69, 0x65, 0x77, 0x73, 0x12, 0x1b, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x56, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6d, 0x65, 0x56, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x76, 0x69, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_view_view_proto_rawDescOnce sync.Once
	file_view_view_proto_rawDescData = file_view_view_proto_rawDesc
)

func file_view_view_proto_rawDescGZIP() []byte {
	file_view_view_proto_rawDescOnce.Do(func() {
		file_view_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_view_view_proto_rawDescData)
	})
	return file_view_view_proto_rawDescData
}

var file_view_view_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_view_view_proto_goTypes = []interface{}{
	(*CreateViewRequest)(nil),      // 0: view.CreateViewRequest
	(*CreateViewResponse)(nil),     // 1: view.CreateViewResponse
	(*GetResumeViewsRequest)(nil),  // 2: view.GetResumeViewsRequest
	(*GetResumeViewsResponse)(nil), // 3: view.GetResumeViewsResponse
	(*View)(nil),                   // 4: view.View
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
}
var file_view_view_proto_depIdxs = []int32{
	4, // 0: view.GetResumeViewsResponse.views:type_name -> view.View
	5, // 1: view.View.viewed_at:type_name -> google.protobuf.Timestamp
	0, // 2: view.ViewService.CreateView:input_type -> view.CreateViewRequest
	2, // 3: view.ViewService.GetResumeViews:input_type -> view.GetResumeViewsRequest
	1, // 4: view.ViewService.CreateView:output_type -> view.CreateViewResponse
	3, // 5: view.ViewService.GetResumeViews:output_type -> view.GetResumeViewsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_view_view_proto_init() }
func file_view_view_proto_init() {
	if File_view_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_view_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateViewRequest); i {
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
		file_view_view_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateViewResponse); i {
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
		file_view_view_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResumeViewsRequest); i {
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
		file_view_view_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResumeViewsResponse); i {
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
		file_view_view_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*View); i {
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
	file_view_view_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_view_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_view_view_proto_goTypes,
		DependencyIndexes: file_view_view_proto_depIdxs,
		MessageInfos:      file_view_view_proto_msgTypes,
	}.Build()
	File_view_view_proto = out.File
	file_view_view_proto_rawDesc = nil
	file_view_view_proto_goTypes = nil
	file_view_view_proto_depIdxs = nil
}
