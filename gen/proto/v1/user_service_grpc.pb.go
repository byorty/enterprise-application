// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/v1/user_service.proto

package pbv1

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
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	GetByUUID(ctx context.Context, in *GetByUserUUIDRequest, opts ...grpc.CallOption) (*User, error)
	GetUserProducts(ctx context.Context, in *GetByUserUUIDRequest, opts ...grpc.CallOption) (*UserProductsResponse, error)
	PutProduct(ctx context.Context, in *PutProductRequest, opts ...grpc.CallOption) (*UserProductsResponse, error)
	ChangeProduct(ctx context.Context, in *ChangeProductRequest, opts ...grpc.CallOption) (*UserProductsResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/pb.v1.UserService/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/pb.v1.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByUUID(ctx context.Context, in *GetByUserUUIDRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.v1.UserService/GetByUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserProducts(ctx context.Context, in *GetByUserUUIDRequest, opts ...grpc.CallOption) (*UserProductsResponse, error) {
	out := new(UserProductsResponse)
	err := c.cc.Invoke(ctx, "/pb.v1.UserService/GetUserProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PutProduct(ctx context.Context, in *PutProductRequest, opts ...grpc.CallOption) (*UserProductsResponse, error) {
	out := new(UserProductsResponse)
	err := c.cc.Invoke(ctx, "/pb.v1.UserService/PutProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeProduct(ctx context.Context, in *ChangeProductRequest, opts ...grpc.CallOption) (*UserProductsResponse, error) {
	out := new(UserProductsResponse)
	err := c.cc.Invoke(ctx, "/pb.v1.UserService/ChangeProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Authorize(context.Context, *AuthorizeRequest) (*TokenResponse, error)
	Register(context.Context, *RegisterRequest) (*TokenResponse, error)
	GetByUUID(context.Context, *GetByUserUUIDRequest) (*User, error)
	GetUserProducts(context.Context, *GetByUserUUIDRequest) (*UserProductsResponse, error)
	PutProduct(context.Context, *PutProductRequest) (*UserProductsResponse, error)
	ChangeProduct(context.Context, *ChangeProductRequest) (*UserProductsResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Authorize(context.Context, *AuthorizeRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedUserServiceServer) Register(context.Context, *RegisterRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) GetByUUID(context.Context, *GetByUserUUIDRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUUID not implemented")
}
func (UnimplementedUserServiceServer) GetUserProducts(context.Context, *GetByUserUUIDRequest) (*UserProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProducts not implemented")
}
func (UnimplementedUserServiceServer) PutProduct(context.Context, *PutProductRequest) (*UserProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutProduct not implemented")
}
func (UnimplementedUserServiceServer) ChangeProduct(context.Context, *ChangeProductRequest) (*UserProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeProduct not implemented")
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

func _UserService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.v1.UserService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.v1.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetByUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUserUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.v1.UserService/GetByUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByUUID(ctx, req.(*GetByUserUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUserUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.v1.UserService/GetUserProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserProducts(ctx, req.(*GetByUserUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PutProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PutProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.v1.UserService/PutProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PutProduct(ctx, req.(*PutProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.v1.UserService/ChangeProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeProduct(ctx, req.(*ChangeProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _UserService_Authorize_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "GetByUUID",
			Handler:    _UserService_GetByUUID_Handler,
		},
		{
			MethodName: "GetUserProducts",
			Handler:    _UserService_GetUserProducts_Handler,
		},
		{
			MethodName: "PutProduct",
			Handler:    _UserService_PutProduct_Handler,
		},
		{
			MethodName: "ChangeProduct",
			Handler:    _UserService_ChangeProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/user_service.proto",
}
