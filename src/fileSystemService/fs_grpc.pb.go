// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package fileSystemService

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

// FileSystemClient is the client API for FileSystem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileSystemClient interface {
	FSCreate(ctx context.Context, opts ...grpc.CallOption) (FileSystem_FSCreateClient, error)
	FSDelete(ctx context.Context, in *Param, opts ...grpc.CallOption) (*Result, error)
	FSMove(ctx context.Context, in *Param, opts ...grpc.CallOption) (*Result, error)
	FSList(ctx context.Context, in *Param, opts ...grpc.CallOption) (*Result, error)
	RCloneCopy(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneCopyClient, error)
	RCloneMove(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneMoveClient, error)
	RCloneLink(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneLinkClient, error)
	RCloneList(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneListClient, error)
	RCloneListFormat(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneListFormatClient, error)
	RCloneListDirectory(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneListDirectoryClient, error)
	RCloneListRemotes(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneListRemotesClient, error)
	RCloneListJson(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneListJsonClient, error)
	RCloneMkdir(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneMkdirClient, error)
	RCloneAbout(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneAboutClient, error)
}

type fileSystemClient struct {
	cc grpc.ClientConnInterface
}

func NewFileSystemClient(cc grpc.ClientConnInterface) FileSystemClient {
	return &fileSystemClient{cc}
}

func (c *fileSystemClient) FSCreate(ctx context.Context, opts ...grpc.CallOption) (FileSystem_FSCreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSystem_ServiceDesc.Streams[0], "/fileSystemService.FileSystem/FSCreate", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSystemFSCreateClient{stream}
	return x, nil
}

type FileSystem_FSCreateClient interface {
	Send(*Param) error
	CloseAndRecv() (*Result, error)
	grpc.ClientStream
}

type fileSystemFSCreateClient struct {
	grpc.ClientStream
}

func (x *fileSystemFSCreateClient) Send(m *Param) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileSystemFSCreateClient) CloseAndRecv() (*Result, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSystemClient) FSDelete(ctx context.Context, in *Param, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/fileSystemService.FileSystem/FSDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemClient) FSMove(ctx context.Context, in *Param, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/fileSystemService.FileSystem/FSMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemClient) FSList(ctx context.Context, in *Param, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/fileSystemService.FileSystem/FSList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemClient) RCloneCopy(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneCopyClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSystem_ServiceDesc.Streams[1], "/fileSystemService.FileSystem/RCloneCopy", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSystemRCloneCopyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSystem_RCloneCopyClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type fileSystemRCloneCopyClient struct {
	grpc.ClientStream
}

func (x *fileSystemRCloneCopyClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSystemClient) RCloneMove(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneMoveClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSystem_ServiceDesc.Streams[2], "/fileSystemService.FileSystem/RCloneMove", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSystemRCloneMoveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSystem_RCloneMoveClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type fileSystemRCloneMoveClient struct {
	grpc.ClientStream
}

func (x *fileSystemRCloneMoveClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSystemClient) RCloneLink(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneLinkClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSystem_ServiceDesc.Streams[3], "/fileSystemService.FileSystem/RCloneLink", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSystemRCloneLinkClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSystem_RCloneLinkClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type fileSystemRCloneLinkClient struct {
	grpc.ClientStream
}

func (x *fileSystemRCloneLinkClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSystemClient) RCloneList(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneListClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSystem_ServiceDesc.Streams[4], "/fileSystemService.FileSystem/RCloneList", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSystemRCloneListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSystem_RCloneListClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type fileSystemRCloneListClient struct {
	grpc.ClientStream
}

func (x *fileSystemRCloneListClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSystemClient) RCloneListFormat(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneListFormatClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSystem_ServiceDesc.Streams[5], "/fileSystemService.FileSystem/RCloneListFormat", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSystemRCloneListFormatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSystem_RCloneListFormatClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type fileSystemRCloneListFormatClient struct {
	grpc.ClientStream
}

func (x *fileSystemRCloneListFormatClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSystemClient) RCloneListDirectory(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneListDirectoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSystem_ServiceDesc.Streams[6], "/fileSystemService.FileSystem/RCloneListDirectory", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSystemRCloneListDirectoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSystem_RCloneListDirectoryClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type fileSystemRCloneListDirectoryClient struct {
	grpc.ClientStream
}

func (x *fileSystemRCloneListDirectoryClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSystemClient) RCloneListRemotes(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneListRemotesClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSystem_ServiceDesc.Streams[7], "/fileSystemService.FileSystem/RCloneListRemotes", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSystemRCloneListRemotesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSystem_RCloneListRemotesClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type fileSystemRCloneListRemotesClient struct {
	grpc.ClientStream
}

func (x *fileSystemRCloneListRemotesClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSystemClient) RCloneListJson(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneListJsonClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSystem_ServiceDesc.Streams[8], "/fileSystemService.FileSystem/RCloneListJson", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSystemRCloneListJsonClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSystem_RCloneListJsonClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type fileSystemRCloneListJsonClient struct {
	grpc.ClientStream
}

func (x *fileSystemRCloneListJsonClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSystemClient) RCloneMkdir(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneMkdirClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSystem_ServiceDesc.Streams[9], "/fileSystemService.FileSystem/RCloneMkdir", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSystemRCloneMkdirClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSystem_RCloneMkdirClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type fileSystemRCloneMkdirClient struct {
	grpc.ClientStream
}

func (x *fileSystemRCloneMkdirClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSystemClient) RCloneAbout(ctx context.Context, in *Param, opts ...grpc.CallOption) (FileSystem_RCloneAboutClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSystem_ServiceDesc.Streams[10], "/fileSystemService.FileSystem/RCloneAbout", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSystemRCloneAboutClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSystem_RCloneAboutClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type fileSystemRCloneAboutClient struct {
	grpc.ClientStream
}

func (x *fileSystemRCloneAboutClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileSystemServer is the server API for FileSystem service.
// All implementations must embed UnimplementedFileSystemServer
// for forward compatibility
type FileSystemServer interface {
	FSCreate(FileSystem_FSCreateServer) error
	FSDelete(context.Context, *Param) (*Result, error)
	FSMove(context.Context, *Param) (*Result, error)
	FSList(context.Context, *Param) (*Result, error)
	RCloneCopy(*Param, FileSystem_RCloneCopyServer) error
	RCloneMove(*Param, FileSystem_RCloneMoveServer) error
	RCloneLink(*Param, FileSystem_RCloneLinkServer) error
	RCloneList(*Param, FileSystem_RCloneListServer) error
	RCloneListFormat(*Param, FileSystem_RCloneListFormatServer) error
	RCloneListDirectory(*Param, FileSystem_RCloneListDirectoryServer) error
	RCloneListRemotes(*Param, FileSystem_RCloneListRemotesServer) error
	RCloneListJson(*Param, FileSystem_RCloneListJsonServer) error
	RCloneMkdir(*Param, FileSystem_RCloneMkdirServer) error
	RCloneAbout(*Param, FileSystem_RCloneAboutServer) error
	mustEmbedUnimplementedFileSystemServer()
}

// UnimplementedFileSystemServer must be embedded to have forward compatible implementations.
type UnimplementedFileSystemServer struct {
}

func (UnimplementedFileSystemServer) FSCreate(FileSystem_FSCreateServer) error {
	return status.Errorf(codes.Unimplemented, "method FSCreate not implemented")
}
func (UnimplementedFileSystemServer) FSDelete(context.Context, *Param) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FSDelete not implemented")
}
func (UnimplementedFileSystemServer) FSMove(context.Context, *Param) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FSMove not implemented")
}
func (UnimplementedFileSystemServer) FSList(context.Context, *Param) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FSList not implemented")
}
func (UnimplementedFileSystemServer) RCloneCopy(*Param, FileSystem_RCloneCopyServer) error {
	return status.Errorf(codes.Unimplemented, "method RCloneCopy not implemented")
}
func (UnimplementedFileSystemServer) RCloneMove(*Param, FileSystem_RCloneMoveServer) error {
	return status.Errorf(codes.Unimplemented, "method RCloneMove not implemented")
}
func (UnimplementedFileSystemServer) RCloneLink(*Param, FileSystem_RCloneLinkServer) error {
	return status.Errorf(codes.Unimplemented, "method RCloneLink not implemented")
}
func (UnimplementedFileSystemServer) RCloneList(*Param, FileSystem_RCloneListServer) error {
	return status.Errorf(codes.Unimplemented, "method RCloneList not implemented")
}
func (UnimplementedFileSystemServer) RCloneListFormat(*Param, FileSystem_RCloneListFormatServer) error {
	return status.Errorf(codes.Unimplemented, "method RCloneListFormat not implemented")
}
func (UnimplementedFileSystemServer) RCloneListDirectory(*Param, FileSystem_RCloneListDirectoryServer) error {
	return status.Errorf(codes.Unimplemented, "method RCloneListDirectory not implemented")
}
func (UnimplementedFileSystemServer) RCloneListRemotes(*Param, FileSystem_RCloneListRemotesServer) error {
	return status.Errorf(codes.Unimplemented, "method RCloneListRemotes not implemented")
}
func (UnimplementedFileSystemServer) RCloneListJson(*Param, FileSystem_RCloneListJsonServer) error {
	return status.Errorf(codes.Unimplemented, "method RCloneListJson not implemented")
}
func (UnimplementedFileSystemServer) RCloneMkdir(*Param, FileSystem_RCloneMkdirServer) error {
	return status.Errorf(codes.Unimplemented, "method RCloneMkdir not implemented")
}
func (UnimplementedFileSystemServer) RCloneAbout(*Param, FileSystem_RCloneAboutServer) error {
	return status.Errorf(codes.Unimplemented, "method RCloneAbout not implemented")
}
func (UnimplementedFileSystemServer) mustEmbedUnimplementedFileSystemServer() {}

// UnsafeFileSystemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileSystemServer will
// result in compilation errors.
type UnsafeFileSystemServer interface {
	mustEmbedUnimplementedFileSystemServer()
}

func RegisterFileSystemServer(s grpc.ServiceRegistrar, srv FileSystemServer) {
	s.RegisterService(&FileSystem_ServiceDesc, srv)
}

func _FileSystem_FSCreate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileSystemServer).FSCreate(&fileSystemFSCreateServer{stream})
}

type FileSystem_FSCreateServer interface {
	SendAndClose(*Result) error
	Recv() (*Param, error)
	grpc.ServerStream
}

type fileSystemFSCreateServer struct {
	grpc.ServerStream
}

func (x *fileSystemFSCreateServer) SendAndClose(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileSystemFSCreateServer) Recv() (*Param, error) {
	m := new(Param)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileSystem_FSDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Param)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemServer).FSDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileSystemService.FileSystem/FSDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemServer).FSDelete(ctx, req.(*Param))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystem_FSMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Param)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemServer).FSMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileSystemService.FileSystem/FSMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemServer).FSMove(ctx, req.(*Param))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystem_FSList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Param)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemServer).FSList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileSystemService.FileSystem/FSList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemServer).FSList(ctx, req.(*Param))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystem_RCloneCopy_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Param)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSystemServer).RCloneCopy(m, &fileSystemRCloneCopyServer{stream})
}

type FileSystem_RCloneCopyServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type fileSystemRCloneCopyServer struct {
	grpc.ServerStream
}

func (x *fileSystemRCloneCopyServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSystem_RCloneMove_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Param)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSystemServer).RCloneMove(m, &fileSystemRCloneMoveServer{stream})
}

type FileSystem_RCloneMoveServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type fileSystemRCloneMoveServer struct {
	grpc.ServerStream
}

func (x *fileSystemRCloneMoveServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSystem_RCloneLink_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Param)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSystemServer).RCloneLink(m, &fileSystemRCloneLinkServer{stream})
}

type FileSystem_RCloneLinkServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type fileSystemRCloneLinkServer struct {
	grpc.ServerStream
}

func (x *fileSystemRCloneLinkServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSystem_RCloneList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Param)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSystemServer).RCloneList(m, &fileSystemRCloneListServer{stream})
}

type FileSystem_RCloneListServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type fileSystemRCloneListServer struct {
	grpc.ServerStream
}

func (x *fileSystemRCloneListServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSystem_RCloneListFormat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Param)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSystemServer).RCloneListFormat(m, &fileSystemRCloneListFormatServer{stream})
}

type FileSystem_RCloneListFormatServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type fileSystemRCloneListFormatServer struct {
	grpc.ServerStream
}

func (x *fileSystemRCloneListFormatServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSystem_RCloneListDirectory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Param)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSystemServer).RCloneListDirectory(m, &fileSystemRCloneListDirectoryServer{stream})
}

type FileSystem_RCloneListDirectoryServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type fileSystemRCloneListDirectoryServer struct {
	grpc.ServerStream
}

func (x *fileSystemRCloneListDirectoryServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSystem_RCloneListRemotes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Param)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSystemServer).RCloneListRemotes(m, &fileSystemRCloneListRemotesServer{stream})
}

type FileSystem_RCloneListRemotesServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type fileSystemRCloneListRemotesServer struct {
	grpc.ServerStream
}

func (x *fileSystemRCloneListRemotesServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSystem_RCloneListJson_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Param)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSystemServer).RCloneListJson(m, &fileSystemRCloneListJsonServer{stream})
}

type FileSystem_RCloneListJsonServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type fileSystemRCloneListJsonServer struct {
	grpc.ServerStream
}

func (x *fileSystemRCloneListJsonServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSystem_RCloneMkdir_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Param)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSystemServer).RCloneMkdir(m, &fileSystemRCloneMkdirServer{stream})
}

type FileSystem_RCloneMkdirServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type fileSystemRCloneMkdirServer struct {
	grpc.ServerStream
}

func (x *fileSystemRCloneMkdirServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSystem_RCloneAbout_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Param)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSystemServer).RCloneAbout(m, &fileSystemRCloneAboutServer{stream})
}

type FileSystem_RCloneAboutServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type fileSystemRCloneAboutServer struct {
	grpc.ServerStream
}

func (x *fileSystemRCloneAboutServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

// FileSystem_ServiceDesc is the grpc.ServiceDesc for FileSystem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileSystem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fileSystemService.FileSystem",
	HandlerType: (*FileSystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FSDelete",
			Handler:    _FileSystem_FSDelete_Handler,
		},
		{
			MethodName: "FSMove",
			Handler:    _FileSystem_FSMove_Handler,
		},
		{
			MethodName: "FSList",
			Handler:    _FileSystem_FSList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FSCreate",
			Handler:       _FileSystem_FSCreate_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RCloneCopy",
			Handler:       _FileSystem_RCloneCopy_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RCloneMove",
			Handler:       _FileSystem_RCloneMove_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RCloneLink",
			Handler:       _FileSystem_RCloneLink_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RCloneList",
			Handler:       _FileSystem_RCloneList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RCloneListFormat",
			Handler:       _FileSystem_RCloneListFormat_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RCloneListDirectory",
			Handler:       _FileSystem_RCloneListDirectory_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RCloneListRemotes",
			Handler:       _FileSystem_RCloneListRemotes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RCloneListJson",
			Handler:       _FileSystem_RCloneListJson_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RCloneMkdir",
			Handler:       _FileSystem_RCloneMkdir_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RCloneAbout",
			Handler:       _FileSystem_RCloneAbout_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fs.proto",
}
