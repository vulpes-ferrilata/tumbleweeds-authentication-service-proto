// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.3
// source: protobuf/authentication.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	models "github.com/vulpes-ferrilata/authentication-service-proto/pb/models"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationClient interface {
	GetTokenByClaimID(ctx context.Context, in *models.GetTokenByClaimIDRequest, opts ...grpc.CallOption) (*models.Token, error)
	GetClaimByAccessToken(ctx context.Context, in *models.GetClaimByAccessTokenRequest, opts ...grpc.CallOption) (*models.Claim, error)
	GetTokenByRefreshToken(ctx context.Context, in *models.GetTokenByRefreshTokenRequest, opts ...grpc.CallOption) (*models.Token, error)
	CreateUserCredential(ctx context.Context, in *models.CreateUserCredentialRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteUserCredential(ctx context.Context, in *models.DeleteUserCredentialRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Login(ctx context.Context, in *models.LoginRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RevokeToken(ctx context.Context, in *models.RevokeTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) GetTokenByClaimID(ctx context.Context, in *models.GetTokenByClaimIDRequest, opts ...grpc.CallOption) (*models.Token, error) {
	out := new(models.Token)
	err := c.cc.Invoke(ctx, "/pb.Authentication/GetTokenByClaimID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) GetClaimByAccessToken(ctx context.Context, in *models.GetClaimByAccessTokenRequest, opts ...grpc.CallOption) (*models.Claim, error) {
	out := new(models.Claim)
	err := c.cc.Invoke(ctx, "/pb.Authentication/GetClaimByAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) GetTokenByRefreshToken(ctx context.Context, in *models.GetTokenByRefreshTokenRequest, opts ...grpc.CallOption) (*models.Token, error) {
	out := new(models.Token)
	err := c.cc.Invoke(ctx, "/pb.Authentication/GetTokenByRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) CreateUserCredential(ctx context.Context, in *models.CreateUserCredentialRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Authentication/CreateUserCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) DeleteUserCredential(ctx context.Context, in *models.DeleteUserCredentialRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Authentication/DeleteUserCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Login(ctx context.Context, in *models.LoginRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Authentication/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) RevokeToken(ctx context.Context, in *models.RevokeTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.Authentication/RevokeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
// All implementations must embed UnimplementedAuthenticationServer
// for forward compatibility
type AuthenticationServer interface {
	GetTokenByClaimID(context.Context, *models.GetTokenByClaimIDRequest) (*models.Token, error)
	GetClaimByAccessToken(context.Context, *models.GetClaimByAccessTokenRequest) (*models.Claim, error)
	GetTokenByRefreshToken(context.Context, *models.GetTokenByRefreshTokenRequest) (*models.Token, error)
	CreateUserCredential(context.Context, *models.CreateUserCredentialRequest) (*empty.Empty, error)
	DeleteUserCredential(context.Context, *models.DeleteUserCredentialRequest) (*empty.Empty, error)
	Login(context.Context, *models.LoginRequest) (*empty.Empty, error)
	RevokeToken(context.Context, *models.RevokeTokenRequest) (*empty.Empty, error)
	mustEmbedUnimplementedAuthenticationServer()
}

// UnimplementedAuthenticationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (UnimplementedAuthenticationServer) GetTokenByClaimID(context.Context, *models.GetTokenByClaimIDRequest) (*models.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenByClaimID not implemented")
}
func (UnimplementedAuthenticationServer) GetClaimByAccessToken(context.Context, *models.GetClaimByAccessTokenRequest) (*models.Claim, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClaimByAccessToken not implemented")
}
func (UnimplementedAuthenticationServer) GetTokenByRefreshToken(context.Context, *models.GetTokenByRefreshTokenRequest) (*models.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenByRefreshToken not implemented")
}
func (UnimplementedAuthenticationServer) CreateUserCredential(context.Context, *models.CreateUserCredentialRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserCredential not implemented")
}
func (UnimplementedAuthenticationServer) DeleteUserCredential(context.Context, *models.DeleteUserCredentialRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserCredential not implemented")
}
func (UnimplementedAuthenticationServer) Login(context.Context, *models.LoginRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthenticationServer) RevokeToken(context.Context, *models.RevokeTokenRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeToken not implemented")
}
func (UnimplementedAuthenticationServer) mustEmbedUnimplementedAuthenticationServer() {}

// UnsafeAuthenticationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServer will
// result in compilation errors.
type UnsafeAuthenticationServer interface {
	mustEmbedUnimplementedAuthenticationServer()
}

func RegisterAuthenticationServer(s grpc.ServiceRegistrar, srv AuthenticationServer) {
	s.RegisterService(&Authentication_ServiceDesc, srv)
}

func _Authentication_GetTokenByClaimID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.GetTokenByClaimIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GetTokenByClaimID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authentication/GetTokenByClaimID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GetTokenByClaimID(ctx, req.(*models.GetTokenByClaimIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_GetClaimByAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.GetClaimByAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GetClaimByAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authentication/GetClaimByAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GetClaimByAccessToken(ctx, req.(*models.GetClaimByAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_GetTokenByRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.GetTokenByRefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GetTokenByRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authentication/GetTokenByRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GetTokenByRefreshToken(ctx, req.(*models.GetTokenByRefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_CreateUserCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.CreateUserCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).CreateUserCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authentication/CreateUserCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).CreateUserCredential(ctx, req.(*models.CreateUserCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_DeleteUserCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.DeleteUserCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).DeleteUserCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authentication/DeleteUserCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).DeleteUserCredential(ctx, req.(*models.DeleteUserCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Login(ctx, req.(*models.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_RevokeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.RevokeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).RevokeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Authentication/RevokeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).RevokeToken(ctx, req.(*models.RevokeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authentication_ServiceDesc is the grpc.ServiceDesc for Authentication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authentication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTokenByClaimID",
			Handler:    _Authentication_GetTokenByClaimID_Handler,
		},
		{
			MethodName: "GetClaimByAccessToken",
			Handler:    _Authentication_GetClaimByAccessToken_Handler,
		},
		{
			MethodName: "GetTokenByRefreshToken",
			Handler:    _Authentication_GetTokenByRefreshToken_Handler,
		},
		{
			MethodName: "CreateUserCredential",
			Handler:    _Authentication_CreateUserCredential_Handler,
		},
		{
			MethodName: "DeleteUserCredential",
			Handler:    _Authentication_DeleteUserCredential_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Authentication_Login_Handler,
		},
		{
			MethodName: "RevokeToken",
			Handler:    _Authentication_RevokeToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/authentication.proto",
}
