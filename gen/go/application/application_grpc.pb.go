// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: application/application.proto

package application

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
	UserApplicationService_CreateApplication_FullMethodName  = "/application.UserApplicationService/CreateApplication"
	UserApplicationService_GetApplicationByID_FullMethodName = "/application.UserApplicationService/GetApplicationByID"
	UserApplicationService_GetApplications_FullMethodName    = "/application.UserApplicationService/GetApplications"
	UserApplicationService_UpdateApplication_FullMethodName  = "/application.UserApplicationService/UpdateApplication"
	UserApplicationService_DeleteApplication_FullMethodName  = "/application.UserApplicationService/DeleteApplication"
)

// UserApplicationServiceClient is the client API for UserApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserApplicationServiceClient interface {
	CreateApplication(ctx context.Context, in *CreateUserApplicationRequest, opts ...grpc.CallOption) (*CreateUserApplicationResponse, error)
	GetApplicationByID(ctx context.Context, in *GetUserApplicationByIDRequest, opts ...grpc.CallOption) (*UserApplication, error)
	GetApplications(ctx context.Context, in *GetUserApplicationsRequest, opts ...grpc.CallOption) (*UserApplications, error)
	UpdateApplication(ctx context.Context, in *UpdateUserApplicationRequest, opts ...grpc.CallOption) (*UserApplication, error)
	DeleteApplication(ctx context.Context, in *DeleteUserApplicationRequest, opts ...grpc.CallOption) (*DeleteUserApplicationResponse, error)
}

type userApplicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserApplicationServiceClient(cc grpc.ClientConnInterface) UserApplicationServiceClient {
	return &userApplicationServiceClient{cc}
}

func (c *userApplicationServiceClient) CreateApplication(ctx context.Context, in *CreateUserApplicationRequest, opts ...grpc.CallOption) (*CreateUserApplicationResponse, error) {
	out := new(CreateUserApplicationResponse)
	err := c.cc.Invoke(ctx, UserApplicationService_CreateApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApplicationServiceClient) GetApplicationByID(ctx context.Context, in *GetUserApplicationByIDRequest, opts ...grpc.CallOption) (*UserApplication, error) {
	out := new(UserApplication)
	err := c.cc.Invoke(ctx, UserApplicationService_GetApplicationByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApplicationServiceClient) GetApplications(ctx context.Context, in *GetUserApplicationsRequest, opts ...grpc.CallOption) (*UserApplications, error) {
	out := new(UserApplications)
	err := c.cc.Invoke(ctx, UserApplicationService_GetApplications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApplicationServiceClient) UpdateApplication(ctx context.Context, in *UpdateUserApplicationRequest, opts ...grpc.CallOption) (*UserApplication, error) {
	out := new(UserApplication)
	err := c.cc.Invoke(ctx, UserApplicationService_UpdateApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userApplicationServiceClient) DeleteApplication(ctx context.Context, in *DeleteUserApplicationRequest, opts ...grpc.CallOption) (*DeleteUserApplicationResponse, error) {
	out := new(DeleteUserApplicationResponse)
	err := c.cc.Invoke(ctx, UserApplicationService_DeleteApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserApplicationServiceServer is the server API for UserApplicationService service.
// All implementations must embed UnimplementedUserApplicationServiceServer
// for forward compatibility
type UserApplicationServiceServer interface {
	CreateApplication(context.Context, *CreateUserApplicationRequest) (*CreateUserApplicationResponse, error)
	GetApplicationByID(context.Context, *GetUserApplicationByIDRequest) (*UserApplication, error)
	GetApplications(context.Context, *GetUserApplicationsRequest) (*UserApplications, error)
	UpdateApplication(context.Context, *UpdateUserApplicationRequest) (*UserApplication, error)
	DeleteApplication(context.Context, *DeleteUserApplicationRequest) (*DeleteUserApplicationResponse, error)
	mustEmbedUnimplementedUserApplicationServiceServer()
}

// UnimplementedUserApplicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserApplicationServiceServer struct {
}

func (UnimplementedUserApplicationServiceServer) CreateApplication(context.Context, *CreateUserApplicationRequest) (*CreateUserApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApplication not implemented")
}
func (UnimplementedUserApplicationServiceServer) GetApplicationByID(context.Context, *GetUserApplicationByIDRequest) (*UserApplication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplicationByID not implemented")
}
func (UnimplementedUserApplicationServiceServer) GetApplications(context.Context, *GetUserApplicationsRequest) (*UserApplications, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplications not implemented")
}
func (UnimplementedUserApplicationServiceServer) UpdateApplication(context.Context, *UpdateUserApplicationRequest) (*UserApplication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApplication not implemented")
}
func (UnimplementedUserApplicationServiceServer) DeleteApplication(context.Context, *DeleteUserApplicationRequest) (*DeleteUserApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApplication not implemented")
}
func (UnimplementedUserApplicationServiceServer) mustEmbedUnimplementedUserApplicationServiceServer() {
}

// UnsafeUserApplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserApplicationServiceServer will
// result in compilation errors.
type UnsafeUserApplicationServiceServer interface {
	mustEmbedUnimplementedUserApplicationServiceServer()
}

func RegisterUserApplicationServiceServer(s grpc.ServiceRegistrar, srv UserApplicationServiceServer) {
	s.RegisterService(&UserApplicationService_ServiceDesc, srv)
}

func _UserApplicationService_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApplicationServiceServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserApplicationService_CreateApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApplicationServiceServer).CreateApplication(ctx, req.(*CreateUserApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApplicationService_GetApplicationByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserApplicationByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApplicationServiceServer).GetApplicationByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserApplicationService_GetApplicationByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApplicationServiceServer).GetApplicationByID(ctx, req.(*GetUserApplicationByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApplicationService_GetApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApplicationServiceServer).GetApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserApplicationService_GetApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApplicationServiceServer).GetApplications(ctx, req.(*GetUserApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApplicationService_UpdateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApplicationServiceServer).UpdateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserApplicationService_UpdateApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApplicationServiceServer).UpdateApplication(ctx, req.(*UpdateUserApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserApplicationService_DeleteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserApplicationServiceServer).DeleteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserApplicationService_DeleteApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserApplicationServiceServer).DeleteApplication(ctx, req.(*DeleteUserApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserApplicationService_ServiceDesc is the grpc.ServiceDesc for UserApplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserApplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "application.UserApplicationService",
	HandlerType: (*UserApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApplication",
			Handler:    _UserApplicationService_CreateApplication_Handler,
		},
		{
			MethodName: "GetApplicationByID",
			Handler:    _UserApplicationService_GetApplicationByID_Handler,
		},
		{
			MethodName: "GetApplications",
			Handler:    _UserApplicationService_GetApplications_Handler,
		},
		{
			MethodName: "UpdateApplication",
			Handler:    _UserApplicationService_UpdateApplication_Handler,
		},
		{
			MethodName: "DeleteApplication",
			Handler:    _UserApplicationService_DeleteApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application/application.proto",
}

const (
	CompanyApplicationService_CreateApplication_FullMethodName  = "/application.CompanyApplicationService/CreateApplication"
	CompanyApplicationService_GetApplicationByID_FullMethodName = "/application.CompanyApplicationService/GetApplicationByID"
	CompanyApplicationService_GetApplications_FullMethodName    = "/application.CompanyApplicationService/GetApplications"
	CompanyApplicationService_UpdateApplication_FullMethodName  = "/application.CompanyApplicationService/UpdateApplication"
	CompanyApplicationService_DeleteApplication_FullMethodName  = "/application.CompanyApplicationService/DeleteApplication"
)

// CompanyApplicationServiceClient is the client API for CompanyApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyApplicationServiceClient interface {
	CreateApplication(ctx context.Context, in *CreateCompanyApplicationRequest, opts ...grpc.CallOption) (*CreateCompanyApplicationResponse, error)
	GetApplicationByID(ctx context.Context, in *GetCompanyApplicationByIDRequest, opts ...grpc.CallOption) (*CompanyApplication, error)
	GetApplications(ctx context.Context, in *GetCompanyApplicationsRequest, opts ...grpc.CallOption) (*CompanyApplications, error)
	UpdateApplication(ctx context.Context, in *UpdateCompanyApplicationRequest, opts ...grpc.CallOption) (*CompanyApplication, error)
	DeleteApplication(ctx context.Context, in *DeleteCompanyApplicationRequest, opts ...grpc.CallOption) (*DeleteCompanyApplicationResponse, error)
}

type companyApplicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyApplicationServiceClient(cc grpc.ClientConnInterface) CompanyApplicationServiceClient {
	return &companyApplicationServiceClient{cc}
}

func (c *companyApplicationServiceClient) CreateApplication(ctx context.Context, in *CreateCompanyApplicationRequest, opts ...grpc.CallOption) (*CreateCompanyApplicationResponse, error) {
	out := new(CreateCompanyApplicationResponse)
	err := c.cc.Invoke(ctx, CompanyApplicationService_CreateApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyApplicationServiceClient) GetApplicationByID(ctx context.Context, in *GetCompanyApplicationByIDRequest, opts ...grpc.CallOption) (*CompanyApplication, error) {
	out := new(CompanyApplication)
	err := c.cc.Invoke(ctx, CompanyApplicationService_GetApplicationByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyApplicationServiceClient) GetApplications(ctx context.Context, in *GetCompanyApplicationsRequest, opts ...grpc.CallOption) (*CompanyApplications, error) {
	out := new(CompanyApplications)
	err := c.cc.Invoke(ctx, CompanyApplicationService_GetApplications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyApplicationServiceClient) UpdateApplication(ctx context.Context, in *UpdateCompanyApplicationRequest, opts ...grpc.CallOption) (*CompanyApplication, error) {
	out := new(CompanyApplication)
	err := c.cc.Invoke(ctx, CompanyApplicationService_UpdateApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyApplicationServiceClient) DeleteApplication(ctx context.Context, in *DeleteCompanyApplicationRequest, opts ...grpc.CallOption) (*DeleteCompanyApplicationResponse, error) {
	out := new(DeleteCompanyApplicationResponse)
	err := c.cc.Invoke(ctx, CompanyApplicationService_DeleteApplication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyApplicationServiceServer is the server API for CompanyApplicationService service.
// All implementations must embed UnimplementedCompanyApplicationServiceServer
// for forward compatibility
type CompanyApplicationServiceServer interface {
	CreateApplication(context.Context, *CreateCompanyApplicationRequest) (*CreateCompanyApplicationResponse, error)
	GetApplicationByID(context.Context, *GetCompanyApplicationByIDRequest) (*CompanyApplication, error)
	GetApplications(context.Context, *GetCompanyApplicationsRequest) (*CompanyApplications, error)
	UpdateApplication(context.Context, *UpdateCompanyApplicationRequest) (*CompanyApplication, error)
	DeleteApplication(context.Context, *DeleteCompanyApplicationRequest) (*DeleteCompanyApplicationResponse, error)
	mustEmbedUnimplementedCompanyApplicationServiceServer()
}

// UnimplementedCompanyApplicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompanyApplicationServiceServer struct {
}

func (UnimplementedCompanyApplicationServiceServer) CreateApplication(context.Context, *CreateCompanyApplicationRequest) (*CreateCompanyApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApplication not implemented")
}
func (UnimplementedCompanyApplicationServiceServer) GetApplicationByID(context.Context, *GetCompanyApplicationByIDRequest) (*CompanyApplication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplicationByID not implemented")
}
func (UnimplementedCompanyApplicationServiceServer) GetApplications(context.Context, *GetCompanyApplicationsRequest) (*CompanyApplications, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplications not implemented")
}
func (UnimplementedCompanyApplicationServiceServer) UpdateApplication(context.Context, *UpdateCompanyApplicationRequest) (*CompanyApplication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApplication not implemented")
}
func (UnimplementedCompanyApplicationServiceServer) DeleteApplication(context.Context, *DeleteCompanyApplicationRequest) (*DeleteCompanyApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApplication not implemented")
}
func (UnimplementedCompanyApplicationServiceServer) mustEmbedUnimplementedCompanyApplicationServiceServer() {
}

// UnsafeCompanyApplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyApplicationServiceServer will
// result in compilation errors.
type UnsafeCompanyApplicationServiceServer interface {
	mustEmbedUnimplementedCompanyApplicationServiceServer()
}

func RegisterCompanyApplicationServiceServer(s grpc.ServiceRegistrar, srv CompanyApplicationServiceServer) {
	s.RegisterService(&CompanyApplicationService_ServiceDesc, srv)
}

func _CompanyApplicationService_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyApplicationServiceServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyApplicationService_CreateApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyApplicationServiceServer).CreateApplication(ctx, req.(*CreateCompanyApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyApplicationService_GetApplicationByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyApplicationByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyApplicationServiceServer).GetApplicationByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyApplicationService_GetApplicationByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyApplicationServiceServer).GetApplicationByID(ctx, req.(*GetCompanyApplicationByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyApplicationService_GetApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyApplicationServiceServer).GetApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyApplicationService_GetApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyApplicationServiceServer).GetApplications(ctx, req.(*GetCompanyApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyApplicationService_UpdateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyApplicationServiceServer).UpdateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyApplicationService_UpdateApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyApplicationServiceServer).UpdateApplication(ctx, req.(*UpdateCompanyApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyApplicationService_DeleteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyApplicationServiceServer).DeleteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyApplicationService_DeleteApplication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyApplicationServiceServer).DeleteApplication(ctx, req.(*DeleteCompanyApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyApplicationService_ServiceDesc is the grpc.ServiceDesc for CompanyApplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyApplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "application.CompanyApplicationService",
	HandlerType: (*CompanyApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApplication",
			Handler:    _CompanyApplicationService_CreateApplication_Handler,
		},
		{
			MethodName: "GetApplicationByID",
			Handler:    _CompanyApplicationService_GetApplicationByID_Handler,
		},
		{
			MethodName: "GetApplications",
			Handler:    _CompanyApplicationService_GetApplications_Handler,
		},
		{
			MethodName: "UpdateApplication",
			Handler:    _CompanyApplicationService_UpdateApplication_Handler,
		},
		{
			MethodName: "DeleteApplication",
			Handler:    _CompanyApplicationService_DeleteApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application/application.proto",
}
