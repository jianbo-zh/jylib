package grpcc

import (
	"context"
	"fmt"

	"github.com/jianbo-zh/jylib/errc"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	mapV1 "github.com/jianbo-zh/jypb/api/parkmap/v1"
)

type IParkMap interface {
	CreateMapVersion(context.Context, *mapV1.CreateMapVersionRequest, ...filterc.Filter) (*mapV1.MapVersion, error)
	GetMapVersions(context.Context, *mapV1.GetMapVersionsRequest, ...filterc.Filter) (*mapV1.GetMapVersionsReply, error)
	GetMapVersion(context.Context, *mapV1.GetMapVersionRequest, ...filterc.Filter) (*mapV1.MapVersion, error)
	GetVersionMap(context.Context, *mapV1.GetVersionMapRequest, ...filterc.Filter) (*mapV1.GetVersionMapReply, error)
	GetVersionPois(context.Context, *mapV1.GetVersionPoisRequest, ...filterc.Filter) (*mapV1.GetVersionPoisReply, error)
}

type ParkMapGrpc struct {
	client IClient
}

func NewParkMapGrpc(cli IClient) IParkMap {
	return &ParkMapGrpc{
		client: cli,
	}
}

func (c *ParkMapGrpc) CreateMapVersion(ctx context.Context, req *mapV1.CreateMapVersionRequest, filters ...filterc.Filter) (*mapV1.MapVersion, error) {
	cli, err := c.client.ParkMapClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.ParkMapClient error: %w", err)
	}
	reply, err := cli.CreateMapVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreateMapVersion error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *ParkMapGrpc) GetMapVersions(ctx context.Context, req *mapV1.GetMapVersionsRequest, filters ...filterc.Filter) (*mapV1.GetMapVersionsReply, error) {
	cli, err := c.client.ParkMapClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.ParkMapClient error: %w", err)
	}
	reply, err := cli.GetMapVersions(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetMapVersions error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *ParkMapGrpc) GetMapVersion(ctx context.Context, req *mapV1.GetMapVersionRequest, filters ...filterc.Filter) (*mapV1.MapVersion, error) {
	cli, err := c.client.ParkMapClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.ParkMapClient error: %w", err)
	}
	reply, err := cli.GetMapVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetMapVersion error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *ParkMapGrpc) GetVersionMap(ctx context.Context, req *mapV1.GetVersionMapRequest, filters ...filterc.Filter) (*mapV1.GetVersionMapReply, error) {
	cli, err := c.client.ParkMapClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.ParkMapClient error: %w", err)
	}
	reply, err := cli.GetVersionMap(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetVersionMap error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *ParkMapGrpc) GetVersionPois(ctx context.Context, req *mapV1.GetVersionPoisRequest, filters ...filterc.Filter) (*mapV1.GetVersionPoisReply, error) {
	cli, err := c.client.ParkMapClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.ParkMapClient error: %w", err)
	}
	reply, err := cli.GetVersionPois(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetVersionPois error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}
