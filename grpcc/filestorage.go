package grpcc

import (
	"context"
	"fmt"

	"github.com/jianbo-zh/jylib/errc"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	fsV1 "github.com/jianbo-zh/jypb/api/filestorage/v1"
)

type IFileStorage interface {
	SaveFile(context.Context, *fsV1.SaveFileRequest, ...filterc.Filter) (*fsV1.File, error)
	GetFiles(context.Context, *fsV1.GetFilesRequest, ...filterc.Filter) (*fsV1.GetFilesReply, error)
	GetFile(context.Context, *fsV1.GetFileRequest, ...filterc.Filter) (*fsV1.File, error)
	GetFileData(context.Context, *fsV1.GetFileDataRequest, ...filterc.Filter) (*fsV1.GetFileDataReply, error)
}

type FileStorageGrpc struct {
	client IClient
}

func NewFileStorageGrpc(cli IClient) IFileStorage {
	return &FileStorageGrpc{
		client: cli,
	}
}

func (c *FileStorageGrpc) SaveFile(ctx context.Context, req *fsV1.SaveFileRequest, filters ...filterc.Filter) (*fsV1.File, error) {
	cli, err := c.client.FileStorageClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.client.FileStorageClient error: %w", err)
	}
	reply, err := cli.SaveFile(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.SaveFile error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *FileStorageGrpc) GetFiles(ctx context.Context, req *fsV1.GetFilesRequest, filters ...filterc.Filter) (*fsV1.GetFilesReply, error) {
	cli, err := c.client.FileStorageClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.client.FileStorageClient error: %w", err)
	}
	reply, err := cli.GetFiles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFiles error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *FileStorageGrpc) GetFile(ctx context.Context, req *fsV1.GetFileRequest, filters ...filterc.Filter) (*fsV1.File, error) {
	cli, err := c.client.FileStorageClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.client.FileStorageClient error: %w", err)
	}
	reply, err := cli.GetFile(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFile error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *FileStorageGrpc) GetFileData(ctx context.Context, req *fsV1.GetFileDataRequest, filters ...filterc.Filter) (*fsV1.GetFileDataReply, error) {
	cli, err := c.client.FileStorageClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.client.FileStorageClient error: %w", err)
	}
	reply, err := cli.GetFileData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFileData error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}
