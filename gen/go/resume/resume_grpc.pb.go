// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: resume/resume.proto

package resume

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
	ResumeService_CreateResume_FullMethodName   = "/resume.ResumeService/CreateResume"
	ResumeService_UpdateResume_FullMethodName   = "/resume.ResumeService/UpdateResume"
	ResumeService_GetResumeByID_FullMethodName  = "/resume.ResumeService/GetResumeByID"
	ResumeService_GetUserResumes_FullMethodName = "/resume.ResumeService/GetUserResumes"
	ResumeService_DeleteResume_FullMethodName   = "/resume.ResumeService/DeleteResume"
)

// ResumeServiceClient is the client API for ResumeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResumeServiceClient interface {
	CreateResume(ctx context.Context, in *CreateResumeRequest, opts ...grpc.CallOption) (*CreateResumeResponse, error)
	UpdateResume(ctx context.Context, in *UpdateResumeRequest, opts ...grpc.CallOption) (*Resume, error)
	GetResumeByID(ctx context.Context, in *GetResumeByIDRequest, opts ...grpc.CallOption) (*Resume, error)
	GetUserResumes(ctx context.Context, in *GetUserResumesRequest, opts ...grpc.CallOption) (*GetUserResumesResponse, error)
	DeleteResume(ctx context.Context, in *DeleteResumeRequest, opts ...grpc.CallOption) (*DeleteResumeResponse, error)
}

type resumeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResumeServiceClient(cc grpc.ClientConnInterface) ResumeServiceClient {
	return &resumeServiceClient{cc}
}

func (c *resumeServiceClient) CreateResume(ctx context.Context, in *CreateResumeRequest, opts ...grpc.CallOption) (*CreateResumeResponse, error) {
	out := new(CreateResumeResponse)
	err := c.cc.Invoke(ctx, ResumeService_CreateResume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) UpdateResume(ctx context.Context, in *UpdateResumeRequest, opts ...grpc.CallOption) (*Resume, error) {
	out := new(Resume)
	err := c.cc.Invoke(ctx, ResumeService_UpdateResume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) GetResumeByID(ctx context.Context, in *GetResumeByIDRequest, opts ...grpc.CallOption) (*Resume, error) {
	out := new(Resume)
	err := c.cc.Invoke(ctx, ResumeService_GetResumeByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) GetUserResumes(ctx context.Context, in *GetUserResumesRequest, opts ...grpc.CallOption) (*GetUserResumesResponse, error) {
	out := new(GetUserResumesResponse)
	err := c.cc.Invoke(ctx, ResumeService_GetUserResumes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resumeServiceClient) DeleteResume(ctx context.Context, in *DeleteResumeRequest, opts ...grpc.CallOption) (*DeleteResumeResponse, error) {
	out := new(DeleteResumeResponse)
	err := c.cc.Invoke(ctx, ResumeService_DeleteResume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResumeServiceServer is the server API for ResumeService service.
// All implementations must embed UnimplementedResumeServiceServer
// for forward compatibility
type ResumeServiceServer interface {
	CreateResume(context.Context, *CreateResumeRequest) (*CreateResumeResponse, error)
	UpdateResume(context.Context, *UpdateResumeRequest) (*Resume, error)
	GetResumeByID(context.Context, *GetResumeByIDRequest) (*Resume, error)
	GetUserResumes(context.Context, *GetUserResumesRequest) (*GetUserResumesResponse, error)
	DeleteResume(context.Context, *DeleteResumeRequest) (*DeleteResumeResponse, error)
	mustEmbedUnimplementedResumeServiceServer()
}

// UnimplementedResumeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResumeServiceServer struct {
}

func (UnimplementedResumeServiceServer) CreateResume(context.Context, *CreateResumeRequest) (*CreateResumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResume not implemented")
}
func (UnimplementedResumeServiceServer) UpdateResume(context.Context, *UpdateResumeRequest) (*Resume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResume not implemented")
}
func (UnimplementedResumeServiceServer) GetResumeByID(context.Context, *GetResumeByIDRequest) (*Resume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResumeByID not implemented")
}
func (UnimplementedResumeServiceServer) GetUserResumes(context.Context, *GetUserResumesRequest) (*GetUserResumesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserResumes not implemented")
}
func (UnimplementedResumeServiceServer) DeleteResume(context.Context, *DeleteResumeRequest) (*DeleteResumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResume not implemented")
}
func (UnimplementedResumeServiceServer) mustEmbedUnimplementedResumeServiceServer() {}

// UnsafeResumeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResumeServiceServer will
// result in compilation errors.
type UnsafeResumeServiceServer interface {
	mustEmbedUnimplementedResumeServiceServer()
}

func RegisterResumeServiceServer(s grpc.ServiceRegistrar, srv ResumeServiceServer) {
	s.RegisterService(&ResumeService_ServiceDesc, srv)
}

func _ResumeService_CreateResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).CreateResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_CreateResume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).CreateResume(ctx, req.(*CreateResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_UpdateResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).UpdateResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_UpdateResume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).UpdateResume(ctx, req.(*UpdateResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_GetResumeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResumeByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).GetResumeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_GetResumeByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).GetResumeByID(ctx, req.(*GetResumeByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_GetUserResumes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserResumesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).GetUserResumes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_GetUserResumes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).GetUserResumes(ctx, req.(*GetUserResumesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResumeService_DeleteResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).DeleteResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResumeService_DeleteResume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).DeleteResume(ctx, req.(*DeleteResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResumeService_ServiceDesc is the grpc.ServiceDesc for ResumeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResumeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resume.ResumeService",
	HandlerType: (*ResumeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResume",
			Handler:    _ResumeService_CreateResume_Handler,
		},
		{
			MethodName: "UpdateResume",
			Handler:    _ResumeService_UpdateResume_Handler,
		},
		{
			MethodName: "GetResumeByID",
			Handler:    _ResumeService_GetResumeByID_Handler,
		},
		{
			MethodName: "GetUserResumes",
			Handler:    _ResumeService_GetUserResumes_Handler,
		},
		{
			MethodName: "DeleteResume",
			Handler:    _ResumeService_DeleteResume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resume/resume.proto",
}
