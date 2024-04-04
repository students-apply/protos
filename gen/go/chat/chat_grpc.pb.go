// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: chat/chat.proto

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserChatService_CreateChat_FullMethodName        = "/chat.UserChatService/CreateChat"
	UserChatService_GetChats_FullMethodName          = "/chat.UserChatService/GetChats"
	UserChatService_DeleteChat_FullMethodName        = "/chat.UserChatService/DeleteChat"
	UserChatService_CreateUserMessage_FullMethodName = "/chat.UserChatService/CreateUserMessage"
	UserChatService_UpdateUserMessage_FullMethodName = "/chat.UserChatService/UpdateUserMessage"
	UserChatService_DeleteUserMessage_FullMethodName = "/chat.UserChatService/DeleteUserMessage"
	UserChatService_ViewMessage_FullMethodName       = "/chat.UserChatService/ViewMessage"
)

// UserChatServiceClient is the client API for UserChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserChatServiceClient interface {
	CreateChat(ctx context.Context, in *CreateUserChatRequest, opts ...grpc.CallOption) (*CreateUserChatResponse, error)
	GetChats(ctx context.Context, in *GetUserChatsRequest, opts ...grpc.CallOption) (*Chats, error)
	DeleteChat(ctx context.Context, in *DeleteUserChatRequest, opts ...grpc.CallOption) (*DeleteUserChatResponse, error)
	CreateUserMessage(ctx context.Context, in *CreateUserMessageRequest, opts ...grpc.CallOption) (*CreateUserMessageResponse, error)
	UpdateUserMessage(ctx context.Context, in *UpdateUserMessageRequest, opts ...grpc.CallOption) (*UpdateUserMessageResponse, error)
	DeleteUserMessage(ctx context.Context, in *DeleteUserMessageRequest, opts ...grpc.CallOption) (*DeleteUserMessageResponse, error)
	ViewMessage(ctx context.Context, in *ViewUserMessageRequest, opts ...grpc.CallOption) (*ViewUserMessageResponse, error)
}

type userChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserChatServiceClient(cc grpc.ClientConnInterface) UserChatServiceClient {
	return &userChatServiceClient{cc}
}

func (c *userChatServiceClient) CreateChat(ctx context.Context, in *CreateUserChatRequest, opts ...grpc.CallOption) (*CreateUserChatResponse, error) {
	out := new(CreateUserChatResponse)
	err := c.cc.Invoke(ctx, UserChatService_CreateChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChatServiceClient) GetChats(ctx context.Context, in *GetUserChatsRequest, opts ...grpc.CallOption) (*Chats, error) {
	out := new(Chats)
	err := c.cc.Invoke(ctx, UserChatService_GetChats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChatServiceClient) DeleteChat(ctx context.Context, in *DeleteUserChatRequest, opts ...grpc.CallOption) (*DeleteUserChatResponse, error) {
	out := new(DeleteUserChatResponse)
	err := c.cc.Invoke(ctx, UserChatService_DeleteChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChatServiceClient) CreateUserMessage(ctx context.Context, in *CreateUserMessageRequest, opts ...grpc.CallOption) (*CreateUserMessageResponse, error) {
	out := new(CreateUserMessageResponse)
	err := c.cc.Invoke(ctx, UserChatService_CreateUserMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChatServiceClient) UpdateUserMessage(ctx context.Context, in *UpdateUserMessageRequest, opts ...grpc.CallOption) (*UpdateUserMessageResponse, error) {
	out := new(UpdateUserMessageResponse)
	err := c.cc.Invoke(ctx, UserChatService_UpdateUserMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChatServiceClient) DeleteUserMessage(ctx context.Context, in *DeleteUserMessageRequest, opts ...grpc.CallOption) (*DeleteUserMessageResponse, error) {
	out := new(DeleteUserMessageResponse)
	err := c.cc.Invoke(ctx, UserChatService_DeleteUserMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userChatServiceClient) ViewMessage(ctx context.Context, in *ViewUserMessageRequest, opts ...grpc.CallOption) (*ViewUserMessageResponse, error) {
	out := new(ViewUserMessageResponse)
	err := c.cc.Invoke(ctx, UserChatService_ViewMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserChatServiceServer is the server API for UserChatService service.
// All implementations must embed UnimplementedUserChatServiceServer
// for forward compatibility
type UserChatServiceServer interface {
	CreateChat(context.Context, *CreateUserChatRequest) (*CreateUserChatResponse, error)
	GetChats(context.Context, *GetUserChatsRequest) (*Chats, error)
	DeleteChat(context.Context, *DeleteUserChatRequest) (*DeleteUserChatResponse, error)
	CreateUserMessage(context.Context, *CreateUserMessageRequest) (*CreateUserMessageResponse, error)
	UpdateUserMessage(context.Context, *UpdateUserMessageRequest) (*UpdateUserMessageResponse, error)
	DeleteUserMessage(context.Context, *DeleteUserMessageRequest) (*DeleteUserMessageResponse, error)
	ViewMessage(context.Context, *ViewUserMessageRequest) (*ViewUserMessageResponse, error)
	mustEmbedUnimplementedUserChatServiceServer()
}

// UnimplementedUserChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserChatServiceServer struct {
}

func (UnimplementedUserChatServiceServer) CreateChat(context.Context, *CreateUserChatRequest) (*CreateUserChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedUserChatServiceServer) GetChats(context.Context, *GetUserChatsRequest) (*Chats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (UnimplementedUserChatServiceServer) DeleteChat(context.Context, *DeleteUserChatRequest) (*DeleteUserChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (UnimplementedUserChatServiceServer) CreateUserMessage(context.Context, *CreateUserMessageRequest) (*CreateUserMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserMessage not implemented")
}
func (UnimplementedUserChatServiceServer) UpdateUserMessage(context.Context, *UpdateUserMessageRequest) (*UpdateUserMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMessage not implemented")
}
func (UnimplementedUserChatServiceServer) DeleteUserMessage(context.Context, *DeleteUserMessageRequest) (*DeleteUserMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserMessage not implemented")
}
func (UnimplementedUserChatServiceServer) ViewMessage(context.Context, *ViewUserMessageRequest) (*ViewUserMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewMessage not implemented")
}
func (UnimplementedUserChatServiceServer) mustEmbedUnimplementedUserChatServiceServer() {}

// UnsafeUserChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserChatServiceServer will
// result in compilation errors.
type UnsafeUserChatServiceServer interface {
	mustEmbedUnimplementedUserChatServiceServer()
}

func RegisterUserChatServiceServer(s grpc.ServiceRegistrar, srv UserChatServiceServer) {
	s.RegisterService(&UserChatService_ServiceDesc, srv)
}

func _UserChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChatService_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChatServiceServer).CreateChat(ctx, req.(*CreateUserChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChatService_GetChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChatServiceServer).GetChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChatService_GetChats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChatServiceServer).GetChats(ctx, req.(*GetUserChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChatService_DeleteChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChatServiceServer).DeleteChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChatService_DeleteChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChatServiceServer).DeleteChat(ctx, req.(*DeleteUserChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChatService_CreateUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChatServiceServer).CreateUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChatService_CreateUserMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChatServiceServer).CreateUserMessage(ctx, req.(*CreateUserMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChatService_UpdateUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChatServiceServer).UpdateUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChatService_UpdateUserMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChatServiceServer).UpdateUserMessage(ctx, req.(*UpdateUserMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChatService_DeleteUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChatServiceServer).DeleteUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChatService_DeleteUserMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChatServiceServer).DeleteUserMessage(ctx, req.(*DeleteUserMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserChatService_ViewMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewUserMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserChatServiceServer).ViewMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserChatService_ViewMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserChatServiceServer).ViewMessage(ctx, req.(*ViewUserMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserChatService_ServiceDesc is the grpc.ServiceDesc for UserChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.UserChatService",
	HandlerType: (*UserChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _UserChatService_CreateChat_Handler,
		},
		{
			MethodName: "GetChats",
			Handler:    _UserChatService_GetChats_Handler,
		},
		{
			MethodName: "DeleteChat",
			Handler:    _UserChatService_DeleteChat_Handler,
		},
		{
			MethodName: "CreateUserMessage",
			Handler:    _UserChatService_CreateUserMessage_Handler,
		},
		{
			MethodName: "UpdateUserMessage",
			Handler:    _UserChatService_UpdateUserMessage_Handler,
		},
		{
			MethodName: "DeleteUserMessage",
			Handler:    _UserChatService_DeleteUserMessage_Handler,
		},
		{
			MethodName: "ViewMessage",
			Handler:    _UserChatService_ViewMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}

const (
	CompanyChatService_CreateChat_FullMethodName        = "/chat.CompanyChatService/CreateChat"
	CompanyChatService_GetChats_FullMethodName          = "/chat.CompanyChatService/GetChats"
	CompanyChatService_DeleteChat_FullMethodName        = "/chat.CompanyChatService/DeleteChat"
	CompanyChatService_CreateUserMessage_FullMethodName = "/chat.CompanyChatService/CreateUserMessage"
	CompanyChatService_UpdateUserMessage_FullMethodName = "/chat.CompanyChatService/UpdateUserMessage"
	CompanyChatService_DeleteUserMessage_FullMethodName = "/chat.CompanyChatService/DeleteUserMessage"
	CompanyChatService_ViewMessage_FullMethodName       = "/chat.CompanyChatService/ViewMessage"
)

// CompanyChatServiceClient is the client API for CompanyChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyChatServiceClient interface {
	CreateChat(ctx context.Context, in *CreateCompanyChatRequest, opts ...grpc.CallOption) (*CreateCompanyChatResponse, error)
	GetChats(ctx context.Context, in *GetCompanyChatsRequest, opts ...grpc.CallOption) (*Chats, error)
	DeleteChat(ctx context.Context, in *DeleteCompanyChatRequest, opts ...grpc.CallOption) (*DeleteCompanyChatResponse, error)
	CreateUserMessage(ctx context.Context, in *CreateCompanyMessageRequest, opts ...grpc.CallOption) (*CreateCompanyMessageResponse, error)
	UpdateUserMessage(ctx context.Context, in *CreateCompanyMessageRequest, opts ...grpc.CallOption) (*CreateCompanyMessageResponse, error)
	DeleteUserMessage(ctx context.Context, in *CreateCompanyMessageRequest, opts ...grpc.CallOption) (*CreateCompanyMessageResponse, error)
	ViewMessage(ctx context.Context, in *ViewCompanyMessageRequest, opts ...grpc.CallOption) (*ViewCompanyMessageResponse, error)
}

type companyChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyChatServiceClient(cc grpc.ClientConnInterface) CompanyChatServiceClient {
	return &companyChatServiceClient{cc}
}

func (c *companyChatServiceClient) CreateChat(ctx context.Context, in *CreateCompanyChatRequest, opts ...grpc.CallOption) (*CreateCompanyChatResponse, error) {
	out := new(CreateCompanyChatResponse)
	err := c.cc.Invoke(ctx, CompanyChatService_CreateChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyChatServiceClient) GetChats(ctx context.Context, in *GetCompanyChatsRequest, opts ...grpc.CallOption) (*Chats, error) {
	out := new(Chats)
	err := c.cc.Invoke(ctx, CompanyChatService_GetChats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyChatServiceClient) DeleteChat(ctx context.Context, in *DeleteCompanyChatRequest, opts ...grpc.CallOption) (*DeleteCompanyChatResponse, error) {
	out := new(DeleteCompanyChatResponse)
	err := c.cc.Invoke(ctx, CompanyChatService_DeleteChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyChatServiceClient) CreateUserMessage(ctx context.Context, in *CreateCompanyMessageRequest, opts ...grpc.CallOption) (*CreateCompanyMessageResponse, error) {
	out := new(CreateCompanyMessageResponse)
	err := c.cc.Invoke(ctx, CompanyChatService_CreateUserMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyChatServiceClient) UpdateUserMessage(ctx context.Context, in *CreateCompanyMessageRequest, opts ...grpc.CallOption) (*CreateCompanyMessageResponse, error) {
	out := new(CreateCompanyMessageResponse)
	err := c.cc.Invoke(ctx, CompanyChatService_UpdateUserMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyChatServiceClient) DeleteUserMessage(ctx context.Context, in *CreateCompanyMessageRequest, opts ...grpc.CallOption) (*CreateCompanyMessageResponse, error) {
	out := new(CreateCompanyMessageResponse)
	err := c.cc.Invoke(ctx, CompanyChatService_DeleteUserMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyChatServiceClient) ViewMessage(ctx context.Context, in *ViewCompanyMessageRequest, opts ...grpc.CallOption) (*ViewCompanyMessageResponse, error) {
	out := new(ViewCompanyMessageResponse)
	err := c.cc.Invoke(ctx, CompanyChatService_ViewMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyChatServiceServer is the server API for CompanyChatService service.
// All implementations must embed UnimplementedCompanyChatServiceServer
// for forward compatibility
type CompanyChatServiceServer interface {
	CreateChat(context.Context, *CreateCompanyChatRequest) (*CreateCompanyChatResponse, error)
	GetChats(context.Context, *GetCompanyChatsRequest) (*Chats, error)
	DeleteChat(context.Context, *DeleteCompanyChatRequest) (*DeleteCompanyChatResponse, error)
	CreateUserMessage(context.Context, *CreateCompanyMessageRequest) (*CreateCompanyMessageResponse, error)
	UpdateUserMessage(context.Context, *CreateCompanyMessageRequest) (*CreateCompanyMessageResponse, error)
	DeleteUserMessage(context.Context, *CreateCompanyMessageRequest) (*CreateCompanyMessageResponse, error)
	ViewMessage(context.Context, *ViewCompanyMessageRequest) (*ViewCompanyMessageResponse, error)
	mustEmbedUnimplementedCompanyChatServiceServer()
}

// UnimplementedCompanyChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompanyChatServiceServer struct {
}

func (UnimplementedCompanyChatServiceServer) CreateChat(context.Context, *CreateCompanyChatRequest) (*CreateCompanyChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedCompanyChatServiceServer) GetChats(context.Context, *GetCompanyChatsRequest) (*Chats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (UnimplementedCompanyChatServiceServer) DeleteChat(context.Context, *DeleteCompanyChatRequest) (*DeleteCompanyChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (UnimplementedCompanyChatServiceServer) CreateUserMessage(context.Context, *CreateCompanyMessageRequest) (*CreateCompanyMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserMessage not implemented")
}
func (UnimplementedCompanyChatServiceServer) UpdateUserMessage(context.Context, *CreateCompanyMessageRequest) (*CreateCompanyMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMessage not implemented")
}
func (UnimplementedCompanyChatServiceServer) DeleteUserMessage(context.Context, *CreateCompanyMessageRequest) (*CreateCompanyMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserMessage not implemented")
}
func (UnimplementedCompanyChatServiceServer) ViewMessage(context.Context, *ViewCompanyMessageRequest) (*ViewCompanyMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewMessage not implemented")
}
func (UnimplementedCompanyChatServiceServer) mustEmbedUnimplementedCompanyChatServiceServer() {}

// UnsafeCompanyChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyChatServiceServer will
// result in compilation errors.
type UnsafeCompanyChatServiceServer interface {
	mustEmbedUnimplementedCompanyChatServiceServer()
}

func RegisterCompanyChatServiceServer(s grpc.ServiceRegistrar, srv CompanyChatServiceServer) {
	s.RegisterService(&CompanyChatService_ServiceDesc, srv)
}

func _CompanyChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyChatService_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyChatServiceServer).CreateChat(ctx, req.(*CreateCompanyChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyChatService_GetChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyChatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyChatServiceServer).GetChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyChatService_GetChats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyChatServiceServer).GetChats(ctx, req.(*GetCompanyChatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyChatService_DeleteChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyChatServiceServer).DeleteChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyChatService_DeleteChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyChatServiceServer).DeleteChat(ctx, req.(*DeleteCompanyChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyChatService_CreateUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyChatServiceServer).CreateUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyChatService_CreateUserMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyChatServiceServer).CreateUserMessage(ctx, req.(*CreateCompanyMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyChatService_UpdateUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyChatServiceServer).UpdateUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyChatService_UpdateUserMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyChatServiceServer).UpdateUserMessage(ctx, req.(*CreateCompanyMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyChatService_DeleteUserMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyChatServiceServer).DeleteUserMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyChatService_DeleteUserMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyChatServiceServer).DeleteUserMessage(ctx, req.(*CreateCompanyMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyChatService_ViewMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewCompanyMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyChatServiceServer).ViewMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyChatService_ViewMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyChatServiceServer).ViewMessage(ctx, req.(*ViewCompanyMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyChatService_ServiceDesc is the grpc.ServiceDesc for CompanyChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.CompanyChatService",
	HandlerType: (*CompanyChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _CompanyChatService_CreateChat_Handler,
		},
		{
			MethodName: "GetChats",
			Handler:    _CompanyChatService_GetChats_Handler,
		},
		{
			MethodName: "DeleteChat",
			Handler:    _CompanyChatService_DeleteChat_Handler,
		},
		{
			MethodName: "CreateUserMessage",
			Handler:    _CompanyChatService_CreateUserMessage_Handler,
		},
		{
			MethodName: "UpdateUserMessage",
			Handler:    _CompanyChatService_UpdateUserMessage_Handler,
		},
		{
			MethodName: "DeleteUserMessage",
			Handler:    _CompanyChatService_DeleteUserMessage_Handler,
		},
		{
			MethodName: "ViewMessage",
			Handler:    _CompanyChatService_ViewMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}