package main

import (
	"avalon-fs/src/fileSystemService"
	"bytes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (f FileSystemServiceWorker) RCloneCopy(param *fileSystemService.Param, server fileSystemService.FileSystem_RCloneCopyServer) error {
	src, dst := param.GetSource(), param.GetDestination()

	if src == `` || dst == `` {
		return status.Error(codes.InvalidArgument, `Invalid Filename`)
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 1,
		Msg:  `Received`,
	}); err != nil {
		return err
	}

	var stdout, stderr bytes.Buffer
	s := append([]string{`copy`, src, dst}, getArgs(param.GetArgs()...)...)
	err := RunCommand(&stdout, &stderr, `rclone`, s...)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if len(stderr.Bytes()) != 0 {
		return status.Error(codes.Internal, stderr.String())
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
		Data: stdout.Bytes(),
	}); err != nil {
		return err
	}

	return nil
}

func (f FileSystemServiceWorker) RCloneMove(param *fileSystemService.Param, server fileSystemService.FileSystem_RCloneMoveServer) error {
	src, dst := param.GetSource(), param.GetDestination()

	if src == `` || dst == `` {
		return status.Error(codes.InvalidArgument, `Invalid Filename`)
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 1,
		Msg:  `Received`,
	}); err != nil {
		return err
	}

	var stdout, stderr bytes.Buffer
	s := append([]string{`move`, src, dst}, getArgs(param.GetArgs()...)...)
	err := RunCommand(&stdout, &stderr, `rclone`, s...)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if len(stderr.Bytes()) != 0 {
		return status.Error(codes.Internal, stderr.String())
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
		Data: stdout.Bytes(),
	}); err != nil {
		return err
	}

	return nil
}

func (f FileSystemServiceWorker) RCloneLink(param *fileSystemService.Param, server fileSystemService.FileSystem_RCloneLinkServer) error {
	src := param.GetSource()

	if src == `` {
		return status.Error(codes.InvalidArgument, `Invalid Filename`)
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 1,
		Msg:  `Received`,
	}); err != nil {
		return err
	}

	var stdout, stderr bytes.Buffer
	s := append([]string{`link`, src}, getArgs(param.GetArgs()...)...)
	err := RunCommand(&stdout, &stderr, `rclone`, s...)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if len(stderr.Bytes()) != 0 {
		return status.Error(codes.Internal, stderr.String())
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
		Data: stdout.Bytes(),
	}); err != nil {
		return err
	}

	return nil
}

func (f FileSystemServiceWorker) RCloneList(param *fileSystemService.Param, server fileSystemService.FileSystem_RCloneListServer) error {
	src := param.GetSource()

	if src == `` {
		return status.Error(codes.InvalidArgument, `Invalid Filename`)
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 1,
		Msg:  `Received`,
	}); err != nil {
		return err
	}

	var stdout, stderr bytes.Buffer
	s := append([]string{`ls`, src}, getArgs(param.GetArgs()...)...)
	err := RunCommand(&stdout, &stderr, `rclone`, s...)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if len(stderr.Bytes()) != 0 {
		return status.Error(codes.Internal, stderr.String())
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
		Data: stdout.Bytes(),
	}); err != nil {
		return err
	}

	return nil
}

func (f FileSystemServiceWorker) RCloneListFormat(param *fileSystemService.Param, server fileSystemService.FileSystem_RCloneListFormatServer) error {
	src := param.GetSource()

	if src == `` {
		return status.Error(codes.InvalidArgument, `Invalid Filename`)
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 1,
		Msg:  `Received`,
	}); err != nil {
		return err
	}

	var stdout, stderr bytes.Buffer
	s := append([]string{`lsf`, src}, getArgs(param.GetArgs()...)...)
	err := RunCommand(&stdout, &stderr, `rclone`, s...)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if len(stderr.Bytes()) != 0 {
		return status.Error(codes.Internal, stderr.String())
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
		Data: stdout.Bytes(),
	}); err != nil {
		return err
	}

	return nil
}

func (f FileSystemServiceWorker) RCloneListDirectory(param *fileSystemService.Param, server fileSystemService.FileSystem_RCloneListDirectoryServer) error {
	src := param.GetSource()

	if src == `` {
		return status.Error(codes.InvalidArgument, `Invalid Filename`)
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 1,
		Msg:  `Received`,
	}); err != nil {
		return err
	}

	var stdout, stderr bytes.Buffer
	s := append([]string{`lsd`, src}, getArgs(param.GetArgs()...)...)
	err := RunCommand(&stdout, &stderr, `rclone`, s...)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if len(stderr.Bytes()) != 0 {
		return status.Error(codes.Internal, stderr.String())
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
		Data: stdout.Bytes(),
	}); err != nil {
		return err
	}

	return nil
}

func (f FileSystemServiceWorker) RCloneListJson(param *fileSystemService.Param, server fileSystemService.FileSystem_RCloneListJsonServer) error {
	src := param.GetSource()

	if src == `` {
		return status.Error(codes.InvalidArgument, `Invalid Filename`)
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 1,
		Msg:  `Received`,
	}); err != nil {
		return err
	}

	var stdout, stderr bytes.Buffer
	s := append([]string{`lsjson`, src}, getArgs(param.GetArgs()...)...)
	err := RunCommand(&stdout, &stderr, `rclone`, s...)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if len(stderr.Bytes()) != 0 {
		return status.Error(codes.Internal, stderr.String())
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
		Data: stdout.Bytes(),
	}); err != nil {
		return err
	}

	return nil
}

func (f FileSystemServiceWorker) RCloneListRemotes(param *fileSystemService.Param, server fileSystemService.FileSystem_RCloneListRemotesServer) error {
	if err := server.Send(&fileSystemService.Result{
		Code: 1,
		Msg:  `Received`,
	}); err != nil {
		return err
	}

	var stdout, stderr bytes.Buffer
	s := append([]string{`listremotes`}, getArgs(param.GetArgs()...)...)
	err := RunCommand(&stdout, &stderr, `rclone`, s...)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if len(stderr.Bytes()) != 0 {
		return status.Error(codes.Internal, stderr.String())
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
		Data: stdout.Bytes(),
	}); err != nil {
		return err
	}

	return nil
}

func (f FileSystemServiceWorker) RCloneMkdir(param *fileSystemService.Param, server fileSystemService.FileSystem_RCloneMkdirServer) error {
	src := param.GetSource()

	if src == `` {
		return status.Error(codes.InvalidArgument, `Invalid Filename`)
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 1,
		Msg:  `Received`,
	}); err != nil {
		return err
	}

	var stdout, stderr bytes.Buffer
	s := append([]string{`mkdir`, src}, getArgs(param.GetArgs()...)...)
	err := RunCommand(&stdout, &stderr, `rclone`, s...)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if len(stderr.Bytes()) != 0 {
		return status.Error(codes.Internal, stderr.String())
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
		Data: stdout.Bytes(),
	}); err != nil {
		return err
	}

	return nil
}

func (f FileSystemServiceWorker) RCloneAbout(param *fileSystemService.Param, server fileSystemService.FileSystem_RCloneAboutServer) error {
	src := param.GetSource()

	if src == `` {
		return status.Error(codes.InvalidArgument, `Invalid Filename`)
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 1,
		Msg:  `Received`,
	}); err != nil {
		return err
	}

	var stdout, stderr bytes.Buffer
	s := append([]string{`about`, src}, getArgs(param.GetArgs()...)...)
	err := RunCommand(&stdout, &stderr, `rclone`, s...)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if len(stderr.Bytes()) != 0 {
		return status.Error(codes.Internal, stderr.String())
	}

	if err := server.Send(&fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
		Data: stdout.Bytes(),
	}); err != nil {
		return err
	}

	return nil
}

func getArgs(param ...fileSystemService.RcloneArguments) []string {
	s := []string{}

	for i := range param {
		switch param[i] {
		case fileSystemService.RcloneArguments_CREATE_EMPTY_SRC_DIRS:
			s = append(s, `--create-empty-src-dirs`)
		case fileSystemService.RcloneArguments_DELETE_EMPTY_SRC_DIRS:
			s = append(s, `--delete-empty-src-dirs`)
		case fileSystemService.RcloneArguments_NO_TRAVERSE:
			s = append(s, `--no-traverse`)
		case fileSystemService.RcloneArguments_ABOUT_FORMAT_JSON:
			s = append(s, `--json`)
		default:
			continue
		}
	}

	return s
}
