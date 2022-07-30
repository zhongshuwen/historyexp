// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: dfuse/eosio/tokenmeta/v1/tokenmeta.proto

package pbtokenmeta

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

// TokenMetaClient is the client API for TokenMeta service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenMetaClient interface {
	GetTokens(ctx context.Context, in *GetTokensRequest, opts ...grpc.CallOption) (*TokensResponse, error)
	GetAccountBalances(ctx context.Context, in *GetAccountBalancesRequest, opts ...grpc.CallOption) (*AccountBalancesResponse, error)
	GetTokenBalances(ctx context.Context, in *GetTokenBalancesRequest, opts ...grpc.CallOption) (*TokenBalancesResponse, error)
}

type tokenMetaClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenMetaClient(cc grpc.ClientConnInterface) TokenMetaClient {
	return &tokenMetaClient{cc}
}

func (c *tokenMetaClient) GetTokens(ctx context.Context, in *GetTokensRequest, opts ...grpc.CallOption) (*TokensResponse, error) {
	out := new(TokensResponse)
	err := c.cc.Invoke(ctx, "/dfuse.eosio.tokenmeta.v1.TokenMeta/GetTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenMetaClient) GetAccountBalances(ctx context.Context, in *GetAccountBalancesRequest, opts ...grpc.CallOption) (*AccountBalancesResponse, error) {
	out := new(AccountBalancesResponse)
	err := c.cc.Invoke(ctx, "/dfuse.eosio.tokenmeta.v1.TokenMeta/GetAccountBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenMetaClient) GetTokenBalances(ctx context.Context, in *GetTokenBalancesRequest, opts ...grpc.CallOption) (*TokenBalancesResponse, error) {
	out := new(TokenBalancesResponse)
	err := c.cc.Invoke(ctx, "/dfuse.eosio.tokenmeta.v1.TokenMeta/GetTokenBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenMetaServer is the server API for TokenMeta service.
// All implementations must embed UnimplementedTokenMetaServer
// for forward compatibility
type TokenMetaServer interface {
	GetTokens(context.Context, *GetTokensRequest) (*TokensResponse, error)
	GetAccountBalances(context.Context, *GetAccountBalancesRequest) (*AccountBalancesResponse, error)
	GetTokenBalances(context.Context, *GetTokenBalancesRequest) (*TokenBalancesResponse, error)
	mustEmbedUnimplementedTokenMetaServer()
}

// UnimplementedTokenMetaServer must be embedded to have forward compatible implementations.
type UnimplementedTokenMetaServer struct {
}

func (UnimplementedTokenMetaServer) GetTokens(context.Context, *GetTokensRequest) (*TokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokens not implemented")
}
func (UnimplementedTokenMetaServer) GetAccountBalances(context.Context, *GetAccountBalancesRequest) (*AccountBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountBalances not implemented")
}
func (UnimplementedTokenMetaServer) GetTokenBalances(context.Context, *GetTokenBalancesRequest) (*TokenBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenBalances not implemented")
}
func (UnimplementedTokenMetaServer) mustEmbedUnimplementedTokenMetaServer() {}

// UnsafeTokenMetaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenMetaServer will
// result in compilation errors.
type UnsafeTokenMetaServer interface {
	mustEmbedUnimplementedTokenMetaServer()
}

func RegisterTokenMetaServer(s grpc.ServiceRegistrar, srv TokenMetaServer) {
	s.RegisterService(&TokenMeta_ServiceDesc, srv)
}

func _TokenMeta_GetTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenMetaServer).GetTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfuse.eosio.tokenmeta.v1.TokenMeta/GetTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenMetaServer).GetTokens(ctx, req.(*GetTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenMeta_GetAccountBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenMetaServer).GetAccountBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfuse.eosio.tokenmeta.v1.TokenMeta/GetAccountBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenMetaServer).GetAccountBalances(ctx, req.(*GetAccountBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenMeta_GetTokenBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenMetaServer).GetTokenBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfuse.eosio.tokenmeta.v1.TokenMeta/GetTokenBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenMetaServer).GetTokenBalances(ctx, req.(*GetTokenBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenMeta_ServiceDesc is the grpc.ServiceDesc for TokenMeta service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenMeta_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.eosio.tokenmeta.v1.TokenMeta",
	HandlerType: (*TokenMetaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTokens",
			Handler:    _TokenMeta_GetTokens_Handler,
		},
		{
			MethodName: "GetAccountBalances",
			Handler:    _TokenMeta_GetAccountBalances_Handler,
		},
		{
			MethodName: "GetTokenBalances",
			Handler:    _TokenMeta_GetTokenBalances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dfuse/eosio/tokenmeta/v1/tokenmeta.proto",
}
