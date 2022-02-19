package main

import (
	"avalon-fs/src/fileSystemService"
	"bufio"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"os"
)

type FileSystemServiceWorker struct {
	fileSystemService.UnimplementedFileSystemServer
}

func (f FileSystemServiceWorker) FSList(ctx context.Context, param *fileSystemService.Param) (*fileSystemService.Result, error) {
	path := param.GetSource()

	var FileList []*fileSystemService.FileInfo

	files, _ := ioutil.ReadDir(path)

	for i := range files {
		FileList = append(FileList, &fileSystemService.FileInfo{
			Index: int32(i),
			IsDir: files[i].IsDir(),
			Name:  files[i].Name(),
		})
	}

	return &fileSystemService.Result{
		Code:     0,
		Msg:      `OK`,
		FileInfo: FileList,
	}, nil
}

func (f FileSystemServiceWorker) FSCreate(server fileSystemService.FileSystem_FSCreateServer) error {
	currentIndex, stash := 0, map[int32]*[]byte{}
	var file *os.File
	for {
		param, err := server.Recv()
		if err == io.EOF {
			file.Close()

			return server.SendAndClose(&fileSystemService.Result{
				Code: 0,
				Msg:  `OK`,
			})
		}

		if err != nil {
			return err
		}

		//TODO:Improve algorithm
		if file == nil {
			file, err = os.OpenFile(param.GetSource(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
			if err != nil {
				return status.Error(codes.Internal, err.Error())
			}
		}

		if param.GetData().Index >= int32(currentIndex) {
			stash[param.GetData().Index] = &param.GetData().Data
		}

		w := bufio.NewWriter(file)
		for _, ok := stash[int32(currentIndex)]; ok; {
			_, err = w.Write(*stash[int32(currentIndex)])
			delete(stash, int32(currentIndex))
			if err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			currentIndex++
			_, ok = stash[int32(currentIndex)]
		}

		err = w.Flush()
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
}

func (f FileSystemServiceWorker) FSDelete(ctx context.Context, param *fileSystemService.Param) (*fileSystemService.Result, error) {
	fileName := param.GetSource()

	if fileName == `` {
		return nil, status.Error(codes.InvalidArgument, `Invalid Filename`)
	}

	err := os.RemoveAll(fileName)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
	}, nil
}

func (f FileSystemServiceWorker) FSMove(ctx context.Context, param *fileSystemService.Param) (*fileSystemService.Result, error) {
	src, dst := param.GetSource(), param.GetDestination()
	if src == `` || dst == `` {
		return nil, status.Error(codes.InvalidArgument, `Invalid Filename`)
	}

	err := os.Rename(src, dst)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &fileSystemService.Result{
		Code: 0,
		Msg:  `OK`,
	}, nil
}
