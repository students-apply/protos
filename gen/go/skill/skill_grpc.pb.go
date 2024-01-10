// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: skill/skill.proto

package skill

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
	SkillService_CreateSkill_FullMethodName = "/skill.SkillService/CreateSkill"
	SkillService_UpdateSkill_FullMethodName = "/skill.SkillService/UpdateSkill"
	SkillService_GetSkill_FullMethodName    = "/skill.SkillService/GetSkill"
	SkillService_DeleteSkill_FullMethodName = "/skill.SkillService/DeleteSkill"
)

// SkillServiceClient is the client API for SkillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SkillServiceClient interface {
	CreateSkill(ctx context.Context, in *CreateSkillRequest, opts ...grpc.CallOption) (*CreateSkillResponse, error)
	UpdateSkill(ctx context.Context, in *UpdateSkillRequest, opts ...grpc.CallOption) (*Skill, error)
	GetSkill(ctx context.Context, in *GetSkillRequest, opts ...grpc.CallOption) (*Skill, error)
	DeleteSkill(ctx context.Context, in *DeleteSkillRequest, opts ...grpc.CallOption) (*DeleteSkillResponse, error)
}

type skillServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSkillServiceClient(cc grpc.ClientConnInterface) SkillServiceClient {
	return &skillServiceClient{cc}
}

func (c *skillServiceClient) CreateSkill(ctx context.Context, in *CreateSkillRequest, opts ...grpc.CallOption) (*CreateSkillResponse, error) {
	out := new(CreateSkillResponse)
	err := c.cc.Invoke(ctx, SkillService_CreateSkill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillServiceClient) UpdateSkill(ctx context.Context, in *UpdateSkillRequest, opts ...grpc.CallOption) (*Skill, error) {
	out := new(Skill)
	err := c.cc.Invoke(ctx, SkillService_UpdateSkill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillServiceClient) GetSkill(ctx context.Context, in *GetSkillRequest, opts ...grpc.CallOption) (*Skill, error) {
	out := new(Skill)
	err := c.cc.Invoke(ctx, SkillService_GetSkill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillServiceClient) DeleteSkill(ctx context.Context, in *DeleteSkillRequest, opts ...grpc.CallOption) (*DeleteSkillResponse, error) {
	out := new(DeleteSkillResponse)
	err := c.cc.Invoke(ctx, SkillService_DeleteSkill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkillServiceServer is the server API for SkillService service.
// All implementations must embed UnimplementedSkillServiceServer
// for forward compatibility
type SkillServiceServer interface {
	CreateSkill(context.Context, *CreateSkillRequest) (*CreateSkillResponse, error)
	UpdateSkill(context.Context, *UpdateSkillRequest) (*Skill, error)
	GetSkill(context.Context, *GetSkillRequest) (*Skill, error)
	DeleteSkill(context.Context, *DeleteSkillRequest) (*DeleteSkillResponse, error)
	mustEmbedUnimplementedSkillServiceServer()
}

// UnimplementedSkillServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSkillServiceServer struct {
}

func (UnimplementedSkillServiceServer) CreateSkill(context.Context, *CreateSkillRequest) (*CreateSkillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSkill not implemented")
}
func (UnimplementedSkillServiceServer) UpdateSkill(context.Context, *UpdateSkillRequest) (*Skill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSkill not implemented")
}
func (UnimplementedSkillServiceServer) GetSkill(context.Context, *GetSkillRequest) (*Skill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkill not implemented")
}
func (UnimplementedSkillServiceServer) DeleteSkill(context.Context, *DeleteSkillRequest) (*DeleteSkillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSkill not implemented")
}
func (UnimplementedSkillServiceServer) mustEmbedUnimplementedSkillServiceServer() {}

// UnsafeSkillServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SkillServiceServer will
// result in compilation errors.
type UnsafeSkillServiceServer interface {
	mustEmbedUnimplementedSkillServiceServer()
}

func RegisterSkillServiceServer(s grpc.ServiceRegistrar, srv SkillServiceServer) {
	s.RegisterService(&SkillService_ServiceDesc, srv)
}

func _SkillService_CreateSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServiceServer).CreateSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillService_CreateSkill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServiceServer).CreateSkill(ctx, req.(*CreateSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillService_UpdateSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServiceServer).UpdateSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillService_UpdateSkill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServiceServer).UpdateSkill(ctx, req.(*UpdateSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillService_GetSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServiceServer).GetSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillService_GetSkill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServiceServer).GetSkill(ctx, req.(*GetSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillService_DeleteSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServiceServer).DeleteSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkillService_DeleteSkill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServiceServer).DeleteSkill(ctx, req.(*DeleteSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SkillService_ServiceDesc is the grpc.ServiceDesc for SkillService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SkillService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skill.SkillService",
	HandlerType: (*SkillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSkill",
			Handler:    _SkillService_CreateSkill_Handler,
		},
		{
			MethodName: "UpdateSkill",
			Handler:    _SkillService_UpdateSkill_Handler,
		},
		{
			MethodName: "GetSkill",
			Handler:    _SkillService_GetSkill_Handler,
		},
		{
			MethodName: "DeleteSkill",
			Handler:    _SkillService_DeleteSkill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "skill/skill.proto",
}
