package grpcc

import (
	"context"
	"fmt"

	configV1 "github.com/jianbo-zh/jyapi/api/carconfig/v1"
)

type ICarConfig interface {
	CreateBaseConfig(context.Context, *configV1.CreateBaseConfigRequest) (*configV1.CreateBaseConfigReply, error)
	CreateScenicAreaConfig(context.Context, *configV1.CreateScenicAreaConfigRequest) (*configV1.CreateScenicAreaConfigReply, error)
	CreateHdMapConfig(context.Context, *configV1.CreateHdMapConfigRequest) (*configV1.CreateHdMapConfigReply, error)
	CreatePoiConfig(context.Context, *configV1.CreatePoiConfigRequest) (*configV1.CreatePoiConfigReply, error)
	CreateCarUiConfig(context.Context, *configV1.CreateCarUiConfigRequest) (*configV1.CreateCarUiConfigReply, error)
	GetCarConfig(context.Context, *configV1.GetCarConfigRequest) (*configV1.GetCarConfigReply, error)
	GetCarConfigs(context.Context, *configV1.GetCarConfigsRequest) (*configV1.GetCarConfigsReply, error)
	GetCarConfigList(context.Context, *configV1.GetCarConfigListRequest) (*configV1.GetCarConfigListReply, error)
	DeleteCarConfig(context.Context, *configV1.DeleteCarConfigRequest) (*configV1.DeleteCarConfigReply, error)
	SendConfigToCar(context.Context, *configV1.SendConfigToCarRequest) (*configV1.SendConfigToCarReply, error)
	CreateCarConfigPack(context.Context, *configV1.CreateCarConfigPackRequest) (*configV1.CreateCarConfigPackReply, error)
	GetCarConfigPack(context.Context, *configV1.GetCarConfigPackRequest) (*configV1.GetCarConfigPackReply, error)
	GetCarConfigPackList(context.Context, *configV1.GetCarConfigPackListRequest) (*configV1.GetCarConfigPackListReply, error)
	DeleteCarConfigPack(context.Context, *configV1.DeleteCarConfigPackRequest) (*configV1.DeleteCarConfigPackReply, error)
	SendConfigPackToCar(context.Context, *configV1.SendConfigPackToCarRequest) (*configV1.SendConfigPackToCarReply, error)
	GetDownloadProcess(context.Context, *configV1.GetDownloadProcessRequest) (*configV1.GetDownloadProcessReply, error)
	UpdateConfigStatus(context.Context, *configV1.UpdateConfigStatusRequest) (*configV1.UpdateConfigStatusReply, error)
}

type CarConfigGrpc struct {
	client IClient
}

func NewCarConfigGrpc(cli IClient) ICarConfig {
	return &CarConfigGrpc{
		client: cli,
	}
}

func (c *CarConfigGrpc) CreateBaseConfig(ctx context.Context, req *configV1.CreateBaseConfigRequest) (*configV1.CreateBaseConfigReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreateBaseConfig(ctx, &configV1.CreateBaseConfigRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.CreateBaseConfig error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) CreateScenicAreaConfig(ctx context.Context, req *configV1.CreateScenicAreaConfigRequest) (*configV1.CreateScenicAreaConfigReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreateScenicAreaConfig(ctx, &configV1.CreateScenicAreaConfigRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.CreateScenicAreaConfig error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) CreateHdMapConfig(ctx context.Context, req *configV1.CreateHdMapConfigRequest) (*configV1.CreateHdMapConfigReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreateHdMapConfig(ctx, &configV1.CreateHdMapConfigRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.CreateHdMapConfig error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) CreatePoiConfig(ctx context.Context, req *configV1.CreatePoiConfigRequest) (*configV1.CreatePoiConfigReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreatePoiConfig(ctx, &configV1.CreatePoiConfigRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.CreatePoiConfig error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) CreateCarUiConfig(ctx context.Context, req *configV1.CreateCarUiConfigRequest) (*configV1.CreateCarUiConfigReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreateCarUiConfig(ctx, &configV1.CreateCarUiConfigRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.CreateCarUiConfig error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetCarConfig(ctx context.Context, req *configV1.GetCarConfigRequest) (*configV1.GetCarConfigReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetCarConfig(ctx, &configV1.GetCarConfigRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.GetCarConfig error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetCarConfigs(ctx context.Context, req *configV1.GetCarConfigsRequest) (*configV1.GetCarConfigsReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetCarConfigs(ctx, &configV1.GetCarConfigsRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.GetCarConfigs error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetCarConfigList(ctx context.Context, req *configV1.GetCarConfigListRequest) (*configV1.GetCarConfigListReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetCarConfigList(ctx, &configV1.GetCarConfigListRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.GetCarConfigList error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) DeleteCarConfig(ctx context.Context, req *configV1.DeleteCarConfigRequest) (*configV1.DeleteCarConfigReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.DeleteCarConfig(ctx, &configV1.DeleteCarConfigRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.DeleteCarConfig error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) SendConfigToCar(ctx context.Context, req *configV1.SendConfigToCarRequest) (*configV1.SendConfigToCarReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.SendConfigToCar(ctx, &configV1.SendConfigToCarRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.SendConfigToCar error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) CreateCarConfigPack(ctx context.Context, req *configV1.CreateCarConfigPackRequest) (*configV1.CreateCarConfigPackReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.CreateCarConfigPack(ctx, &configV1.CreateCarConfigPackRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.CreateCarConfigPack error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetCarConfigPack(ctx context.Context, req *configV1.GetCarConfigPackRequest) (*configV1.GetCarConfigPackReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetCarConfigPack(ctx, &configV1.GetCarConfigPackRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.GetCarConfigPack error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetCarConfigPackList(ctx context.Context, req *configV1.GetCarConfigPackListRequest) (*configV1.GetCarConfigPackListReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetCarConfigPackList(ctx, &configV1.GetCarConfigPackListRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.GetCarConfigPackList error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) DeleteCarConfigPack(ctx context.Context, req *configV1.DeleteCarConfigPackRequest) (*configV1.DeleteCarConfigPackReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.DeleteCarConfigPack(ctx, &configV1.DeleteCarConfigPackRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.DeleteCarConfigPack error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) SendConfigPackToCar(ctx context.Context, req *configV1.SendConfigPackToCarRequest) (*configV1.SendConfigPackToCarReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.SendConfigPackToCar(ctx, &configV1.SendConfigPackToCarRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.SendConfigPackToCar error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) GetDownloadProcess(ctx context.Context, req *configV1.GetDownloadProcessRequest) (*configV1.GetDownloadProcessReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.GetDownloadProcess(ctx, &configV1.GetDownloadProcessRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.GetDownloadProcess error: %w", err)
	}

	return reply, nil
}

func (c *CarConfigGrpc) UpdateConfigStatus(ctx context.Context, req *configV1.UpdateConfigStatusRequest) (*configV1.UpdateConfigStatusReply, error) {
	cli, err := c.client.CarConfigClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("client.CarConfigClient error: %w", err)
	}
	reply, err := cli.UpdateConfigStatus(ctx, &configV1.UpdateConfigStatusRequest{})
	if err != nil {
		return nil, fmt.Errorf("cli.UpdateConfigStatus error: %w", err)
	}

	return reply, nil
}
