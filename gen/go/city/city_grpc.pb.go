// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: city/city.proto

package city

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
	CityService_GetCities_FullMethodName = "/city.CityService/GetCities"
	CityService_GetCity_FullMethodName   = "/city.CityService/GetCity"
)

// CityServiceClient is the client API for CityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CityServiceClient interface {
	GetCities(ctx context.Context, in *GetCitiesRequest, opts ...grpc.CallOption) (*GetCitiesResponse, error)
	GetCity(ctx context.Context, in *GetCityRequest, opts ...grpc.CallOption) (*GetCityResponse, error)
}

type cityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCityServiceClient(cc grpc.ClientConnInterface) CityServiceClient {
	return &cityServiceClient{cc}
}

func (c *cityServiceClient) GetCities(ctx context.Context, in *GetCitiesRequest, opts ...grpc.CallOption) (*GetCitiesResponse, error) {
	out := new(GetCitiesResponse)
	err := c.cc.Invoke(ctx, CityService_GetCities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) GetCity(ctx context.Context, in *GetCityRequest, opts ...grpc.CallOption) (*GetCityResponse, error) {
	out := new(GetCityResponse)
	err := c.cc.Invoke(ctx, CityService_GetCity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CityServiceServer is the server API for CityService service.
// All implementations must embed UnimplementedCityServiceServer
// for forward compatibility
type CityServiceServer interface {
	GetCities(context.Context, *GetCitiesRequest) (*GetCitiesResponse, error)
	GetCity(context.Context, *GetCityRequest) (*GetCityResponse, error)
	mustEmbedUnimplementedCityServiceServer()
}

// UnimplementedCityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCityServiceServer struct {
}

func (UnimplementedCityServiceServer) GetCities(context.Context, *GetCitiesRequest) (*GetCitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCities not implemented")
}
func (UnimplementedCityServiceServer) GetCity(context.Context, *GetCityRequest) (*GetCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCity not implemented")
}
func (UnimplementedCityServiceServer) mustEmbedUnimplementedCityServiceServer() {}

// UnsafeCityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CityServiceServer will
// result in compilation errors.
type UnsafeCityServiceServer interface {
	mustEmbedUnimplementedCityServiceServer()
}

func RegisterCityServiceServer(s grpc.ServiceRegistrar, srv CityServiceServer) {
	s.RegisterService(&CityService_ServiceDesc, srv)
}

func _CityService_GetCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).GetCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CityService_GetCities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).GetCities(ctx, req.(*GetCitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_GetCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).GetCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CityService_GetCity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).GetCity(ctx, req.(*GetCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CityService_ServiceDesc is the grpc.ServiceDesc for CityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "city.CityService",
	HandlerType: (*CityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCities",
			Handler:    _CityService_GetCities_Handler,
		},
		{
			MethodName: "GetCity",
			Handler:    _CityService_GetCity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "city/city.proto",
}
