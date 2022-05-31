// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/edible/recipe/service.proto

package erpb

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

// EdibleRecipeServiceClient is the client API for EdibleRecipeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EdibleRecipeServiceClient interface {
	// Unary
	Store(ctx context.Context, in *RecipeStoreRequest, opts ...grpc.CallOption) (*RecipeStoreResponse, error)
}

type edibleRecipeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEdibleRecipeServiceClient(cc grpc.ClientConnInterface) EdibleRecipeServiceClient {
	return &edibleRecipeServiceClient{cc}
}

func (c *edibleRecipeServiceClient) Store(ctx context.Context, in *RecipeStoreRequest, opts ...grpc.CallOption) (*RecipeStoreResponse, error) {
	out := new(RecipeStoreResponse)
	err := c.cc.Invoke(ctx, "/edible.recipe.service.EdibleRecipeService/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EdibleRecipeServiceServer is the server API for EdibleRecipeService service.
// All implementations should embed UnimplementedEdibleRecipeServiceServer
// for forward compatibility
type EdibleRecipeServiceServer interface {
	// Unary
	Store(context.Context, *RecipeStoreRequest) (*RecipeStoreResponse, error)
}

// UnimplementedEdibleRecipeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEdibleRecipeServiceServer struct {
}

func (UnimplementedEdibleRecipeServiceServer) Store(context.Context, *RecipeStoreRequest) (*RecipeStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}

// UnsafeEdibleRecipeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EdibleRecipeServiceServer will
// result in compilation errors.
type UnsafeEdibleRecipeServiceServer interface {
	mustEmbedUnimplementedEdibleRecipeServiceServer()
}

func RegisterEdibleRecipeServiceServer(s grpc.ServiceRegistrar, srv EdibleRecipeServiceServer) {
	s.RegisterService(&EdibleRecipeService_ServiceDesc, srv)
}

func _EdibleRecipeService_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecipeStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdibleRecipeServiceServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edible.recipe.service.EdibleRecipeService/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdibleRecipeServiceServer).Store(ctx, req.(*RecipeStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EdibleRecipeService_ServiceDesc is the grpc.ServiceDesc for EdibleRecipeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EdibleRecipeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "edible.recipe.service.EdibleRecipeService",
	HandlerType: (*EdibleRecipeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _EdibleRecipeService_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/edible/recipe/service.proto",
}