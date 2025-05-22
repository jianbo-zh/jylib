package grpcc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	mapV1 "github.com/jianbo-zh/jypb/api/parkmap/v1"
)

type IParkMap interface {
	CreateMapVersion(context.Context, *mapV1.CreateMapVersionRequest) (*mapV1.MapVersion, error)
	GetMapVersions(context.Context, *mapV1.GetMapVersionsRequest) (*mapV1.GetMapVersionsReply, error)
	GetMapVersion(context.Context, *mapV1.GetMapVersionRequest) (*mapV1.MapVersion, error)
	GetVersionMap(context.Context, *mapV1.GetVersionMapRequest) (*mapV1.GetVersionMapReply, error)
	GetVersionPois(context.Context, *mapV1.GetVersionPoisRequest) (*mapV1.GetVersionPoisReply, error)
}

type ParkMapGrpc struct {
	client IClient
}

func NewParkMapGrpc(cli IClient) IParkMap {
	return &ParkMapGrpc{
		client: cli,
	}
}

func (c *ParkMapGrpc) CreateMapVersion(ctx context.Context, req *mapV1.CreateMapVersionRequest) (*mapV1.MapVersion, error) {
	cli, err := c.client.ParkMapClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.ParkMapClient error: %w", err)
	}
	reply, err := cli.CreateMapVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreateMapVersion error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *ParkMapGrpc) GetMapVersions(ctx context.Context, req *mapV1.GetMapVersionsRequest) (*mapV1.GetMapVersionsReply, error) {
	cli, err := c.client.ParkMapClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.ParkMapClient error: %w", err)
	}
	reply, err := cli.GetMapVersions(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetMapVersions error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *ParkMapGrpc) GetMapVersion(ctx context.Context, req *mapV1.GetMapVersionRequest) (*mapV1.MapVersion, error) {
	cli, err := c.client.ParkMapClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.ParkMapClient error: %w", err)
	}
	reply, err := cli.GetMapVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetMapVersion error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *ParkMapGrpc) GetVersionMap(ctx context.Context, req *mapV1.GetVersionMapRequest) (*mapV1.GetVersionMapReply, error) {
	cli, err := c.client.ParkMapClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.ParkMapClient error: %w", err)
	}
	reply, err := cli.GetVersionMap(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetVersionMap error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *ParkMapGrpc) GetVersionPois(ctx context.Context, req *mapV1.GetVersionPoisRequest) (*mapV1.GetVersionPoisReply, error) {
	cli, err := c.client.ParkMapClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.ParkMapClient error: %w", err)
	}
	reply, err := cli.GetVersionPois(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetVersionPois error: %w", errors.FromError(err))
	}
	return reply, nil
}
