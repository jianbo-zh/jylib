package grpcc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	configV1 "github.com/jianbo-zh/jypb/api/carconfig/v1"
)

type ICarConfig interface {
	CreateBaseConfig(context.Context, *configV1.CreateBaseConfigRequest, ...filterc.Filter) (*configV1.CarConfig, error)
	CreateScenicAreaConfig(context.Context, *configV1.CreateScenicAreaConfigRequest, ...filterc.Filter) (*configV1.CarConfig, error)
	CreateHdMapConfig(context.Context, *configV1.CreateHdMapConfigRequest, ...filterc.Filter) (*configV1.CarConfig, error)
	CreatePoiConfig(context.Context, *configV1.CreatePoiConfigRequest, ...filterc.Filter) (*configV1.CarConfig, error)
	CreateCarUiConfig(context.Context, *configV1.CreateCarUiConfigRequest, ...filterc.Filter) (*configV1.CarConfig, error)
	GetCarConfig(context.Context, *configV1.GetCarConfigRequest, ...filterc.Filter) (*configV1.CarConfig, error)
	GetCarConfigs(context.Context, *configV1.GetCarConfigsRequest, ...filterc.Filter) (*configV1.GetCarConfigsReply, error)
	GetCarConfigList(context.Context, *configV1.GetCarConfigListRequest, ...filterc.Filter) (*configV1.GetCarConfigListReply, error)
	DeleteCarConfig(context.Context, *configV1.DeleteCarConfigRequest, ...filterc.Filter) (*configV1.DeleteCarConfigReply, error)
	SendConfigToCar(context.Context, *configV1.SendConfigToCarRequest, ...filterc.Filter) (*configV1.SendConfigToCarReply, error)
	CreateCarConfigPack(context.Context, *configV1.CreateCarConfigPackRequest, ...filterc.Filter) (*configV1.CarConfigPack, error)
	GetCarConfigPack(context.Context, *configV1.GetCarConfigPackRequest, ...filterc.Filter) (*configV1.CarConfigPack, error)
	GetCarConfigPackList(context.Context, *configV1.GetCarConfigPackListRequest, ...filterc.Filter) (*configV1.GetCarConfigPackListReply, error)
	DeleteCarConfigPack(context.Context, *configV1.DeleteCarConfigPackRequest, ...filterc.Filter) (*configV1.DeleteCarConfigPackReply, error)
	SendConfigPackToCar(context.Context, *configV1.SendConfigPackToCarRequest, ...filterc.Filter) (*configV1.SendConfigPackToCarReply, error)
	GetDownloadProcess(context.Context, *configV1.GetDownloadProcessRequest, ...filterc.Filter) (*configV1.GetDownloadProcessReply, error)
	UpdateConfigStatus(context.Context, *configV1.UpdateConfigStatusRequest, ...filterc.Filter) (*configV1.UpdateConfigStatusReply, error)
}

type CarConfigGrpc struct {
	client IClient
}

func NewCarConfigGrpc(cli IClient) ICarConfig {
	return &CarConfigGrpc{
		client: cli,
	}
}

func (c *CarConfigGrpc) CreateBaseConfig(ctx context.Context, req *configV1.CreateBaseConfigRequest, filters ...filterc.Filter) (*configV1.CarConfig, error) {
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

func (c *CarConfigGrpc) CreateScenicAreaConfig(ctx context.Context, req *configV1.CreateScenicAreaConfigRequest, filters ...filterc.Filter) (*configV1.CarConfig, error) {
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

func (c *CarConfigGrpc) CreateHdMapConfig(ctx context.Context, req *configV1.CreateHdMapConfigRequest, filters ...filterc.Filter) (*configV1.CarConfig, error) {
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

func (c *CarConfigGrpc) CreatePoiConfig(ctx context.Context, req *configV1.CreatePoiConfigRequest, filters ...filterc.Filter) (*configV1.CarConfig, error) {
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

func (c *CarConfigGrpc) CreateCarUiConfig(ctx context.Context, req *configV1.CreateCarUiConfigRequest, filters ...filterc.Filter) (*configV1.CarConfig, error) {
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

func (c *CarConfigGrpc) GetCarConfig(ctx context.Context, req *configV1.GetCarConfigRequest, filters ...filterc.Filter) (*configV1.CarConfig, error) {
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

func (c *CarConfigGrpc) GetCarConfigs(ctx context.Context, req *configV1.GetCarConfigsRequest, filters ...filterc.Filter) (*configV1.GetCarConfigsReply, error) {
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

func (c *CarConfigGrpc) GetCarConfigList(ctx context.Context, req *configV1.GetCarConfigListRequest, filters ...filterc.Filter) (*configV1.GetCarConfigListReply, error) {
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

func (c *CarConfigGrpc) DeleteCarConfig(ctx context.Context, req *configV1.DeleteCarConfigRequest, filters ...filterc.Filter) (*configV1.DeleteCarConfigReply, error) {
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

func (c *CarConfigGrpc) SendConfigToCar(ctx context.Context, req *configV1.SendConfigToCarRequest, filters ...filterc.Filter) (*configV1.SendConfigToCarReply, error) {
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

func (c *CarConfigGrpc) CreateCarConfigPack(ctx context.Context, req *configV1.CreateCarConfigPackRequest, filters ...filterc.Filter) (*configV1.CarConfigPack, error) {
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

func (c *CarConfigGrpc) GetCarConfigPack(ctx context.Context, req *configV1.GetCarConfigPackRequest, filters ...filterc.Filter) (*configV1.CarConfigPack, error) {
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

func (c *CarConfigGrpc) GetCarConfigPackList(ctx context.Context, req *configV1.GetCarConfigPackListRequest, filters ...filterc.Filter) (*configV1.GetCarConfigPackListReply, error) {
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

func (c *CarConfigGrpc) DeleteCarConfigPack(ctx context.Context, req *configV1.DeleteCarConfigPackRequest, filters ...filterc.Filter) (*configV1.DeleteCarConfigPackReply, error) {
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

func (c *CarConfigGrpc) SendConfigPackToCar(ctx context.Context, req *configV1.SendConfigPackToCarRequest, filters ...filterc.Filter) (*configV1.SendConfigPackToCarReply, error) {
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

func (c *CarConfigGrpc) GetDownloadProcess(ctx context.Context, req *configV1.GetDownloadProcessRequest, filters ...filterc.Filter) (*configV1.GetDownloadProcessReply, error) {
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

func (c *CarConfigGrpc) UpdateConfigStatus(ctx context.Context, req *configV1.UpdateConfigStatusRequest, filters ...filterc.Filter) (*configV1.UpdateConfigStatusReply, error) {
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
