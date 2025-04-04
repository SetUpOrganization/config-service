// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc2
// source: api.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Config_Parse_FullMethodName            = "/config.Config/Parse"
	Config_CreateConfig_FullMethodName     = "/config.Config/CreateConfig"
	Config_UpdateConfig_FullMethodName     = "/config.Config/UpdateConfig"
	Config_DeleteConfig_FullMethodName     = "/config.Config/DeleteConfig"
	Config_GetConfig_FullMethodName        = "/config.Config/GetConfig"
	Config_GetConfigsBuUser_FullMethodName = "/config.Config/GetConfigsBuUser"
)

// ConfigClient is the client API for Config service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigClient interface {
	Parse(ctx context.Context, in *ParseReq, opts ...grpc.CallOption) (*ParseRes, error)
	CreateConfig(ctx context.Context, in *CreateConfigReq, opts ...grpc.CallOption) (*CreateConfigRes, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigReq, opts ...grpc.CallOption) (*UpdateConfigRes, error)
	DeleteConfig(ctx context.Context, in *DeleteConfigReq, opts ...grpc.CallOption) (*DeleteConfigRes, error)
	GetConfig(ctx context.Context, in *GetConfigReq, opts ...grpc.CallOption) (*CreateConfigRes, error)
	GetConfigsBuUser(ctx context.Context, in *GetConfigsBuUserReq, opts ...grpc.CallOption) (*GetConfigsBuUserRes, error)
}

type configClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigClient(cc grpc.ClientConnInterface) ConfigClient {
	return &configClient{cc}
}

func (c *configClient) Parse(ctx context.Context, in *ParseReq, opts ...grpc.CallOption) (*ParseRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ParseRes)
	err := c.cc.Invoke(ctx, Config_Parse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) CreateConfig(ctx context.Context, in *CreateConfigReq, opts ...grpc.CallOption) (*CreateConfigRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateConfigRes)
	err := c.cc.Invoke(ctx, Config_CreateConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) UpdateConfig(ctx context.Context, in *UpdateConfigReq, opts ...grpc.CallOption) (*UpdateConfigRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateConfigRes)
	err := c.cc.Invoke(ctx, Config_UpdateConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) DeleteConfig(ctx context.Context, in *DeleteConfigReq, opts ...grpc.CallOption) (*DeleteConfigRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteConfigRes)
	err := c.cc.Invoke(ctx, Config_DeleteConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) GetConfig(ctx context.Context, in *GetConfigReq, opts ...grpc.CallOption) (*CreateConfigRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateConfigRes)
	err := c.cc.Invoke(ctx, Config_GetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) GetConfigsBuUser(ctx context.Context, in *GetConfigsBuUserReq, opts ...grpc.CallOption) (*GetConfigsBuUserRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConfigsBuUserRes)
	err := c.cc.Invoke(ctx, Config_GetConfigsBuUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServer is the server API for Config service.
// All implementations must embed UnimplementedConfigServer
// for forward compatibility.
type ConfigServer interface {
	Parse(context.Context, *ParseReq) (*ParseRes, error)
	CreateConfig(context.Context, *CreateConfigReq) (*CreateConfigRes, error)
	UpdateConfig(context.Context, *UpdateConfigReq) (*UpdateConfigRes, error)
	DeleteConfig(context.Context, *DeleteConfigReq) (*DeleteConfigRes, error)
	GetConfig(context.Context, *GetConfigReq) (*CreateConfigRes, error)
	GetConfigsBuUser(context.Context, *GetConfigsBuUserReq) (*GetConfigsBuUserRes, error)
	mustEmbedUnimplementedConfigServer()
}

// UnimplementedConfigServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConfigServer struct{}

func (UnimplementedConfigServer) Parse(context.Context, *ParseReq) (*ParseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parse not implemented")
}
func (UnimplementedConfigServer) CreateConfig(context.Context, *CreateConfigReq) (*CreateConfigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConfig not implemented")
}
func (UnimplementedConfigServer) UpdateConfig(context.Context, *UpdateConfigReq) (*UpdateConfigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedConfigServer) DeleteConfig(context.Context, *DeleteConfigReq) (*DeleteConfigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfig not implemented")
}
func (UnimplementedConfigServer) GetConfig(context.Context, *GetConfigReq) (*CreateConfigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedConfigServer) GetConfigsBuUser(context.Context, *GetConfigsBuUserReq) (*GetConfigsBuUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigsBuUser not implemented")
}
func (UnimplementedConfigServer) mustEmbedUnimplementedConfigServer() {}
func (UnimplementedConfigServer) testEmbeddedByValue()                {}

// UnsafeConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServer will
// result in compilation errors.
type UnsafeConfigServer interface {
	mustEmbedUnimplementedConfigServer()
}

func RegisterConfigServer(s grpc.ServiceRegistrar, srv ConfigServer) {
	// If the following call pancis, it indicates UnimplementedConfigServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Config_ServiceDesc, srv)
}

func _Config_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_Parse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).Parse(ctx, req.(*ParseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_CreateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).CreateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_CreateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).CreateConfig(ctx, req.(*CreateConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_UpdateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).UpdateConfig(ctx, req.(*UpdateConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_DeleteConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).DeleteConfig(ctx, req.(*DeleteConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).GetConfig(ctx, req.(*GetConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Config_GetConfigsBuUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigsBuUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServer).GetConfigsBuUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Config_GetConfigsBuUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServer).GetConfigsBuUser(ctx, req.(*GetConfigsBuUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Config_ServiceDesc is the grpc.ServiceDesc for Config service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Config_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.Config",
	HandlerType: (*ConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _Config_Parse_Handler,
		},
		{
			MethodName: "CreateConfig",
			Handler:    _Config_CreateConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _Config_UpdateConfig_Handler,
		},
		{
			MethodName: "DeleteConfig",
			Handler:    _Config_DeleteConfig_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _Config_GetConfig_Handler,
		},
		{
			MethodName: "GetConfigsBuUser",
			Handler:    _Config_GetConfigsBuUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

const (
	Users_GetUserAvatar_FullMethodName = "/config.Users/GetUserAvatar"
)

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	GetUserAvatar(ctx context.Context, in *GetUsersAvatarRequest, opts ...grpc.CallOption) (*GetUsersAvatarResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) GetUserAvatar(ctx context.Context, in *GetUsersAvatarRequest, opts ...grpc.CallOption) (*GetUsersAvatarResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUsersAvatarResponse)
	err := c.cc.Invoke(ctx, Users_GetUserAvatar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility.
type UsersServer interface {
	GetUserAvatar(context.Context, *GetUsersAvatarRequest) (*GetUsersAvatarResponse, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsersServer struct{}

func (UnimplementedUsersServer) GetUserAvatar(context.Context, *GetUsersAvatarRequest) (*GetUsersAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAvatar not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}
func (UnimplementedUsersServer) testEmbeddedByValue()               {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	// If the following call pancis, it indicates UnimplementedUsersServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_GetUserAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetUserAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserAvatar(ctx, req.(*GetUsersAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserAvatar",
			Handler:    _Users_GetUserAvatar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
