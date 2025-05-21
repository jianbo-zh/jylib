package grpcc

import (
	"context"
	"fmt"

	fsV1 "github.com/jianbo-zh/jyapi/api/filestorage/v1"
)

type IFileStorage interface {
	SaveFile(context.Context, *fsV1.SaveFileRequest) (*fsV1.File, error)
	GetFiles(context.Context, *fsV1.GetFilesRequest) (*fsV1.GetFilesReply, error)
	GetFile(context.Context, *fsV1.GetFileRequest) (*fsV1.File, error)
	GetFileData(context.Context, *fsV1.GetFileDataRequest) (*fsV1.GetFileDataReply, error)
}

type FileStorageGrpc struct {
	client IClient
}

func NewFileStorageGrpc(cli IClient) IFileStorage {
	return &FileStorageGrpc{
		client: cli,
	}
}

func (c *FileStorageGrpc) SaveFile(ctx context.Context, req *fsV1.SaveFileRequest) (*fsV1.File, error) {
	cli, err := c.client.FileStorageClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.client.FileStorageClient error: %w", err)
	}
	reply, err := cli.SaveFile(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.SaveFile error: %w", err)
	}
	return reply, nil
}

func (c *FileStorageGrpc) GetFiles(ctx context.Context, req *fsV1.GetFilesRequest) (*fsV1.GetFilesReply, error) {
	cli, err := c.client.FileStorageClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.client.FileStorageClient error: %w", err)
	}
	reply, err := cli.GetFiles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFiles error: %w", err)
	}
	return reply, nil
}

func (c *FileStorageGrpc) GetFile(ctx context.Context, req *fsV1.GetFileRequest) (*fsV1.File, error) {
	cli, err := c.client.FileStorageClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.client.FileStorageClient error: %w", err)
	}
	reply, err := cli.GetFile(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFile error: %w", err)
	}
	return reply, nil
}

func (c *FileStorageGrpc) GetFileData(ctx context.Context, req *fsV1.GetFileDataRequest) (*fsV1.GetFileDataReply, error) {
	cli, err := c.client.FileStorageClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.client.FileStorageClient error: %w", err)
	}
	reply, err := cli.GetFileData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFileData error: %w", err)
	}
	return reply, nil
}
