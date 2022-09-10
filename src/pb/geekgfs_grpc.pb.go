// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: geekgfs.proto

package pb

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

// MasterServerToClientClient is the client API for MasterServerToClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterServerToClientClient interface {
	ListFiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	CreateFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	AppendFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	CreateChunk(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	ReadFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	WriteFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	DeleteFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type masterServerToClientClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterServerToClientClient(cc grpc.ClientConnInterface) MasterServerToClientClient {
	return &masterServerToClientClient{cc}
}

func (c *masterServerToClientClient) ListFiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.MasterServerToClient/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServerToClientClient) CreateFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.MasterServerToClient/CreateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServerToClientClient) AppendFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.MasterServerToClient/AppendFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServerToClientClient) CreateChunk(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.MasterServerToClient/CreateChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServerToClientClient) ReadFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.MasterServerToClient/ReadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServerToClientClient) WriteFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.MasterServerToClient/WriteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServerToClientClient) DeleteFile(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.MasterServerToClient/DeleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServerToClientServer is the server API for MasterServerToClient service.
// All implementations must embed UnimplementedMasterServerToClientServer
// for forward compatibility
type MasterServerToClientServer interface {
	ListFiles(context.Context, *Request) (*Reply, error)
	CreateFile(context.Context, *Request) (*Reply, error)
	AppendFile(context.Context, *Request) (*Reply, error)
	CreateChunk(context.Context, *Request) (*Reply, error)
	ReadFile(context.Context, *Request) (*Reply, error)
	WriteFile(context.Context, *Request) (*Reply, error)
	DeleteFile(context.Context, *Request) (*Reply, error)
	mustEmbedUnimplementedMasterServerToClientServer()
}

// UnimplementedMasterServerToClientServer must be embedded to have forward compatible implementations.
type UnimplementedMasterServerToClientServer struct {
}

func (UnimplementedMasterServerToClientServer) ListFiles(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedMasterServerToClientServer) CreateFile(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedMasterServerToClientServer) AppendFile(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendFile not implemented")
}
func (UnimplementedMasterServerToClientServer) CreateChunk(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChunk not implemented")
}
func (UnimplementedMasterServerToClientServer) ReadFile(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadFile not implemented")
}
func (UnimplementedMasterServerToClientServer) WriteFile(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteFile not implemented")
}
func (UnimplementedMasterServerToClientServer) DeleteFile(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedMasterServerToClientServer) mustEmbedUnimplementedMasterServerToClientServer() {}

// UnsafeMasterServerToClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServerToClientServer will
// result in compilation errors.
type UnsafeMasterServerToClientServer interface {
	mustEmbedUnimplementedMasterServerToClientServer()
}

func RegisterMasterServerToClientServer(s grpc.ServiceRegistrar, srv MasterServerToClientServer) {
	s.RegisterService(&MasterServerToClient_ServiceDesc, srv)
}

func _MasterServerToClient_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServerToClientServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterServerToClient/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServerToClientServer).ListFiles(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterServerToClient_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServerToClientServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterServerToClient/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServerToClientServer).CreateFile(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterServerToClient_AppendFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServerToClientServer).AppendFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterServerToClient/AppendFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServerToClientServer).AppendFile(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterServerToClient_CreateChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServerToClientServer).CreateChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterServerToClient/CreateChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServerToClientServer).CreateChunk(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterServerToClient_ReadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServerToClientServer).ReadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterServerToClient/ReadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServerToClientServer).ReadFile(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterServerToClient_WriteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServerToClientServer).WriteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterServerToClient/WriteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServerToClientServer).WriteFile(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterServerToClient_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServerToClientServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MasterServerToClient/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServerToClientServer).DeleteFile(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterServerToClient_ServiceDesc is the grpc.ServiceDesc for MasterServerToClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterServerToClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MasterServerToClient",
	HandlerType: (*MasterServerToClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _MasterServerToClient_ListFiles_Handler,
		},
		{
			MethodName: "CreateFile",
			Handler:    _MasterServerToClient_CreateFile_Handler,
		},
		{
			MethodName: "AppendFile",
			Handler:    _MasterServerToClient_AppendFile_Handler,
		},
		{
			MethodName: "CreateChunk",
			Handler:    _MasterServerToClient_CreateChunk_Handler,
		},
		{
			MethodName: "ReadFile",
			Handler:    _MasterServerToClient_ReadFile_Handler,
		},
		{
			MethodName: "WriteFile",
			Handler:    _MasterServerToClient_WriteFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _MasterServerToClient_DeleteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geekgfs.proto",
}

// ChunkServerToClientClient is the client API for ChunkServerToClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChunkServerToClientClient interface {
	Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	GetChunkSpace(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	Write(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	Append(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	Read(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type chunkServerToClientClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkServerToClientClient(cc grpc.ClientConnInterface) ChunkServerToClientClient {
	return &chunkServerToClientClient{cc}
}

func (c *chunkServerToClientClient) Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.ChunkServerToClient/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServerToClientClient) GetChunkSpace(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.ChunkServerToClient/GetChunkSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServerToClientClient) Write(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.ChunkServerToClient/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServerToClientClient) Append(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.ChunkServerToClient/Append", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServerToClientClient) Read(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pb.ChunkServerToClient/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkServerToClientServer is the server API for ChunkServerToClient service.
// All implementations must embed UnimplementedChunkServerToClientServer
// for forward compatibility
type ChunkServerToClientServer interface {
	Create(context.Context, *Request) (*Reply, error)
	GetChunkSpace(context.Context, *Request) (*Reply, error)
	Write(context.Context, *Request) (*Reply, error)
	Append(context.Context, *Request) (*Reply, error)
	Read(context.Context, *Request) (*Reply, error)
	mustEmbedUnimplementedChunkServerToClientServer()
}

// UnimplementedChunkServerToClientServer must be embedded to have forward compatible implementations.
type UnimplementedChunkServerToClientServer struct {
}

func (UnimplementedChunkServerToClientServer) Create(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedChunkServerToClientServer) GetChunkSpace(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChunkSpace not implemented")
}
func (UnimplementedChunkServerToClientServer) Write(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedChunkServerToClientServer) Append(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Append not implemented")
}
func (UnimplementedChunkServerToClientServer) Read(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedChunkServerToClientServer) mustEmbedUnimplementedChunkServerToClientServer() {}

// UnsafeChunkServerToClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChunkServerToClientServer will
// result in compilation errors.
type UnsafeChunkServerToClientServer interface {
	mustEmbedUnimplementedChunkServerToClientServer()
}

func RegisterChunkServerToClientServer(s grpc.ServiceRegistrar, srv ChunkServerToClientServer) {
	s.RegisterService(&ChunkServerToClient_ServiceDesc, srv)
}

func _ChunkServerToClient_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerToClientServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChunkServerToClient/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerToClientServer).Create(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkServerToClient_GetChunkSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerToClientServer).GetChunkSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChunkServerToClient/GetChunkSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerToClientServer).GetChunkSpace(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkServerToClient_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerToClientServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChunkServerToClient/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerToClientServer).Write(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkServerToClient_Append_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerToClientServer).Append(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChunkServerToClient/Append",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerToClientServer).Append(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkServerToClient_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServerToClientServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChunkServerToClient/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServerToClientServer).Read(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ChunkServerToClient_ServiceDesc is the grpc.ServiceDesc for ChunkServerToClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChunkServerToClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ChunkServerToClient",
	HandlerType: (*ChunkServerToClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ChunkServerToClient_Create_Handler,
		},
		{
			MethodName: "GetChunkSpace",
			Handler:    _ChunkServerToClient_GetChunkSpace_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _ChunkServerToClient_Write_Handler,
		},
		{
			MethodName: "Append",
			Handler:    _ChunkServerToClient_Append_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ChunkServerToClient_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geekgfs.proto",
}