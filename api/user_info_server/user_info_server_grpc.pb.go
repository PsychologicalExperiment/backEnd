// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: user_info_server/user_info_server.proto

package user_info_server

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// register a user
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRsp, error)
	// login a user
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRsp, error)
	// 查询用户信息
	GetUserInfoBySearchKey(ctx context.Context, in *GetUserInfoBySearchKeyReq, opts ...grpc.CallOption) (*GetUserInfoBySearchKeyRsp, error)
	// 批量查询用户信息
	BatchGetUserInfos(ctx context.Context, in *BatchGetUserInfoReq, opts ...grpc.CallOption) (*BatchGetUserInfoRsp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRsp, error) {
	out := new(RegisterRsp)
	err := c.cc.Invoke(ctx, "/grpc.psychological_experiment.user_info_server.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRsp, error) {
	out := new(LoginRsp)
	err := c.cc.Invoke(ctx, "/grpc.psychological_experiment.user_info_server.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfoBySearchKey(ctx context.Context, in *GetUserInfoBySearchKeyReq, opts ...grpc.CallOption) (*GetUserInfoBySearchKeyRsp, error) {
	out := new(GetUserInfoBySearchKeyRsp)
	err := c.cc.Invoke(ctx, "/grpc.psychological_experiment.user_info_server.UserService/GetUserInfoBySearchKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BatchGetUserInfos(ctx context.Context, in *BatchGetUserInfoReq, opts ...grpc.CallOption) (*BatchGetUserInfoRsp, error) {
	out := new(BatchGetUserInfoRsp)
	err := c.cc.Invoke(ctx, "/grpc.psychological_experiment.user_info_server.UserService/BatchGetUserInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// register a user
	Register(context.Context, *RegisterReq) (*RegisterRsp, error)
	// login a user
	Login(context.Context, *LoginReq) (*LoginRsp, error)
	// 查询用户信息
	GetUserInfoBySearchKey(context.Context, *GetUserInfoBySearchKeyReq) (*GetUserInfoBySearchKeyRsp, error)
	// 批量查询用户信息
	BatchGetUserInfos(context.Context, *BatchGetUserInfoReq) (*BatchGetUserInfoRsp, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *RegisterReq) (*RegisterRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginReq) (*LoginRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfoBySearchKey(context.Context, *GetUserInfoBySearchKeyReq) (*GetUserInfoBySearchKeyRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoBySearchKey not implemented")
}
func (UnimplementedUserServiceServer) BatchGetUserInfos(context.Context, *BatchGetUserInfoReq) (*BatchGetUserInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetUserInfos not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.psychological_experiment.user_info_server.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.psychological_experiment.user_info_server.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfoBySearchKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoBySearchKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfoBySearchKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.psychological_experiment.user_info_server.UserService/GetUserInfoBySearchKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfoBySearchKey(ctx, req.(*GetUserInfoBySearchKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BatchGetUserInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BatchGetUserInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.psychological_experiment.user_info_server.UserService/BatchGetUserInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BatchGetUserInfos(ctx, req.(*BatchGetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.psychological_experiment.user_info_server.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "GetUserInfoBySearchKey",
			Handler:    _UserService_GetUserInfoBySearchKey_Handler,
		},
		{
			MethodName: "BatchGetUserInfos",
			Handler:    _UserService_BatchGetUserInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_info_server/user_info_server.proto",
}
