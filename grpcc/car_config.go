package grpcc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/selector"
	configV1 "github.com/jianbo-zh/jypb/api/carconfig/v1"
)

type ICarConfig interface {
	CreateBaseConfig(context.Context, *configV1.CreateBaseConfigRequest, ...selector.NodeFilter) (*configV1.CarConfig, error)
	CreateScenicAreaConfig(context.Context, *configV1.CreateScenicAreaConfigRequest, ...selector.NodeFilter) (*configV1.CarConfig, error)
	CreateHdMapConfig(context.Context, *configV1.CreateHdMapConfigRequest, ...selector.NodeFilter) (*configV1.CarConfig, error)
	CreatePoiConfig(context.Context, *configV1.CreatePoiConfigRequest, ...selector.NodeFilter) (*configV1.CarConfig, error)
	CreateCarUiConfig(context.Context, *configV1.CreateCarUiConfigRequest, ...selector.NodeFilter) (*configV1.CarConfig, error)
	GetCarConfig(context.Context, *configV1.GetCarConfigRequest, ...selector.NodeFilter) (*configV1.CarConfig, error)
	GetCarConfigs(context.Context, *configV1.GetCarConfigsRequest, ...selector.NodeFilter) (*configV1.GetCarConfigsReply, error)
	GetCarConfigList(context.Context, *configV1.GetCarConfigListRequest, ...selector.NodeFilter) (*configV1.GetCarConfigListReply, error)
	DeleteCarConfig(context.Context, *configV1.DeleteCarConfigRequest, ...selector.NodeFilter) (*configV1.DeleteCarConfigReply, error)
	SendConfigToCar(context.Context, *configV1.SendConfigToCarRequest, ...selector.NodeFilter) (*configV1.SendConfigToCarReply, error)
	CreateCarConfigPack(context.Context, *configV1.CreateCarConfigPackRequest, ...selector.NodeFilter) (*configV1.CarConfigPack, error)
	GetCarConfigPack(context.Context, *configV1.GetCarConfigPackRequest, ...selector.NodeFilter) (*configV1.CarConfigPack, error)
	GetCarConfigPackList(context.Context, *configV1.GetCarConfigPackListRequest, ...selector.NodeFilter) (*configV1.GetCarConfigPackListReply, error)
	DeleteCarConfigPack(context.Context, *configV1.DeleteCarConfigPackRequest, ...selector.NodeFilter) (*configV1.DeleteCarConfigPackReply, error)
	SendConfigPackToCar(context.Context, *configV1.SendConfigPackToCarRequest, ...selector.NodeFilter) (*configV1.SendConfigPackToCarReply, error)
	GetDownloadProcess(context.Context, *configV1.GetDownloadProcessRequest, ...selector.NodeFilter) (*configV1.GetDownloadProcessReply, error)
	UpdateConfigStatus(context.Context, *configV1.UpdateConfigStatusRequest, ...selector.NodeFilter) (*configV1.UpdateConfigStatusReply, error)
}

type CarConfigGrpc struct {
	client IClient
}

func NewCarConfigGrpc(cli IClient) ICarConfig {
	return &CarConfigGrpc{
		client: cli,
	}
}

func (c *CarConfigGrpc) CreateBaseConfig(ctx context.Context, req *configV1.CreateBaseConfigRequest, filters ...selector.NodeFilter) (*configV1.CarConfig, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreateBaseConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreateBaseConfigerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) CreateScenicAreaConfig(ctx context.Context, req *configV1.CreateScenicAreaConfigRequest, filters ...selector.NodeFilter) (*configV1.CarConfig, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreateScenicAreaConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreateScenicAreaConfigerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) CreateHdMapConfig(ctx context.Context, req *configV1.CreateHdMapConfigRequest, filters ...selector.NodeFilter) (*configV1.CarConfig, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreateHdMapConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreateHdMapConfigerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) CreatePoiConfig(ctx context.Context, req *configV1.CreatePoiConfigRequest, filters ...selector.NodeFilter) (*configV1.CarConfig, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreatePoiConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreatePoiConfigerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) CreateCarUiConfig(ctx context.Context, req *configV1.CreateCarUiConfigRequest, filters ...selector.NodeFilter) (*configV1.CarConfig, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreateCarUiConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreateCarUiConfigerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetCarConfig(ctx context.Context, req *configV1.GetCarConfigRequest, filters ...selector.NodeFilter) (*configV1.CarConfig, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetCarConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetCarConfigerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetCarConfigs(ctx context.Context, req *configV1.GetCarConfigsRequest, filters ...selector.NodeFilter) (*configV1.GetCarConfigsReply, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetCarConfigs(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetCarConfigserror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetCarConfigList(ctx context.Context, req *configV1.GetCarConfigListRequest, filters ...selector.NodeFilter) (*configV1.GetCarConfigListReply, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetCarConfigList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetCarConfigListerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) DeleteCarConfig(ctx context.Context, req *configV1.DeleteCarConfigRequest, filters ...selector.NodeFilter) (*configV1.DeleteCarConfigReply, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.DeleteCarConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DeleteCarConfigerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) SendConfigToCar(ctx context.Context, req *configV1.SendConfigToCarRequest, filters ...selector.NodeFilter) (*configV1.SendConfigToCarReply, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.SendConfigToCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.SendConfigToCarerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) CreateCarConfigPack(ctx context.Context, req *configV1.CreateCarConfigPackRequest, filters ...selector.NodeFilter) (*configV1.CarConfigPack, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreateCarConfigPack(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreateCarConfigPackerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetCarConfigPack(ctx context.Context, req *configV1.GetCarConfigPackRequest, filters ...selector.NodeFilter) (*configV1.CarConfigPack, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetCarConfigPack(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetCarConfigPackerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetCarConfigPackList(ctx context.Context, req *configV1.GetCarConfigPackListRequest, filters ...selector.NodeFilter) (*configV1.GetCarConfigPackListReply, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetCarConfigPackList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetCarConfigPackListerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) DeleteCarConfigPack(ctx context.Context, req *configV1.DeleteCarConfigPackRequest, filters ...selector.NodeFilter) (*configV1.DeleteCarConfigPackReply, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.DeleteCarConfigPack(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DeleteCarConfigPackerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) SendConfigPackToCar(ctx context.Context, req *configV1.SendConfigPackToCarRequest, filters ...selector.NodeFilter) (*configV1.SendConfigPackToCarReply, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.SendConfigPackToCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.SendConfigPackToCarerror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetDownloadProcess(ctx context.Context, req *configV1.GetDownloadProcessRequest, filters ...selector.NodeFilter) (*configV1.GetDownloadProcessReply, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetDownloadProcess(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetDownloadProcesserror: %w", errors.FromError(err))
	}

	return reply, nil
}

func (c *CarConfigGrpc) UpdateConfigStatus(ctx context.Context, req *configV1.UpdateConfigStatusRequest, filters ...selector.NodeFilter) (*configV1.UpdateConfigStatusReply, error) {
	cli, err := c.client.CarConfigClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.UpdateConfigStatus(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UpdateConfigStatuserror: %w", errors.FromError(err))
	}

	return reply, nil
}
