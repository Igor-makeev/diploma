// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/secret.proto

package proto

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
	Secret_CreateSecret_FullMethodName           = "/proto.Secret/CreateSecret"
	Secret_GetSecret_FullMethodName              = "/proto.Secret/GetSecret"
	Secret_DeleteSecret_FullMethodName           = "/proto.Secret/DeleteSecret"
	Secret_EditSecret_FullMethodName             = "/proto.Secret/EditSecret"
	Secret_GetListOfSecretsByType_FullMethodName = "/proto.Secret/GetListOfSecretsByType"
)

// SecretClient is the client API for Secret service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretClient interface {
	CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretResponse, error)
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error)
	DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*DeleteSecretResponse, error)
	EditSecret(ctx context.Context, in *EditSecretRequest, opts ...grpc.CallOption) (*EditSecretResponse, error)
	GetListOfSecretsByType(ctx context.Context, in *GetListOfSecretsByTypeRequest, opts ...grpc.CallOption) (*GetListOfSecretsByTypeResponse, error)
}

type secretClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretClient(cc grpc.ClientConnInterface) SecretClient {
	return &secretClient{cc}
}

func (c *secretClient) CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretResponse, error) {
	out := new(CreateSecretResponse)
	err := c.cc.Invoke(ctx, Secret_CreateSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error) {
	out := new(GetSecretResponse)
	err := c.cc.Invoke(ctx, Secret_GetSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*DeleteSecretResponse, error) {
	out := new(DeleteSecretResponse)
	err := c.cc.Invoke(ctx, Secret_DeleteSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) EditSecret(ctx context.Context, in *EditSecretRequest, opts ...grpc.CallOption) (*EditSecretResponse, error) {
	out := new(EditSecretResponse)
	err := c.cc.Invoke(ctx, Secret_EditSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) GetListOfSecretsByType(ctx context.Context, in *GetListOfSecretsByTypeRequest, opts ...grpc.CallOption) (*GetListOfSecretsByTypeResponse, error) {
	out := new(GetListOfSecretsByTypeResponse)
	err := c.cc.Invoke(ctx, Secret_GetListOfSecretsByType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServer is the server API for Secret service.
// All implementations must embed UnimplementedSecretServer
// for forward compatibility
type SecretServer interface {
	CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretResponse, error)
	GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error)
	DeleteSecret(context.Context, *DeleteSecretRequest) (*DeleteSecretResponse, error)
	EditSecret(context.Context, *EditSecretRequest) (*EditSecretResponse, error)
	GetListOfSecretsByType(context.Context, *GetListOfSecretsByTypeRequest) (*GetListOfSecretsByTypeResponse, error)
	mustEmbedUnimplementedSecretServer()
}

// UnimplementedSecretServer must be embedded to have forward compatible implementations.
type UnimplementedSecretServer struct {
}

func (UnimplementedSecretServer) CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedSecretServer) GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedSecretServer) DeleteSecret(context.Context, *DeleteSecretRequest) (*DeleteSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}
func (UnimplementedSecretServer) EditSecret(context.Context, *EditSecretRequest) (*EditSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditSecret not implemented")
}
func (UnimplementedSecretServer) GetListOfSecretsByType(context.Context, *GetListOfSecretsByTypeRequest) (*GetListOfSecretsByTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListOfSecretsByType not implemented")
}
func (UnimplementedSecretServer) mustEmbedUnimplementedSecretServer() {}

// UnsafeSecretServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretServer will
// result in compilation errors.
type UnsafeSecretServer interface {
	mustEmbedUnimplementedSecretServer()
}

func RegisterSecretServer(s grpc.ServiceRegistrar, srv SecretServer) {
	s.RegisterService(&Secret_ServiceDesc, srv)
}

func _Secret_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_CreateSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).CreateSecret(ctx, req.(*CreateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_GetSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_DeleteSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).DeleteSecret(ctx, req.(*DeleteSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_EditSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).EditSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_EditSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).EditSecret(ctx, req.(*EditSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_GetListOfSecretsByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListOfSecretsByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).GetListOfSecretsByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Secret_GetListOfSecretsByType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).GetListOfSecretsByType(ctx, req.(*GetListOfSecretsByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Secret_ServiceDesc is the grpc.ServiceDesc for Secret service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Secret_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Secret",
	HandlerType: (*SecretServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSecret",
			Handler:    _Secret_CreateSecret_Handler,
		},
		{
			MethodName: "GetSecret",
			Handler:    _Secret_GetSecret_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _Secret_DeleteSecret_Handler,
		},
		{
			MethodName: "EditSecret",
			Handler:    _Secret_EditSecret_Handler,
		},
		{
			MethodName: "GetListOfSecretsByType",
			Handler:    _Secret_GetListOfSecretsByType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/secret.proto",
}
