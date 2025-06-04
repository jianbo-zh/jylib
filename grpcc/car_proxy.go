package grpcc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	"github.com/jianbo-zh/jypb/api/zzz/v1"
)

type ICarProxy interface {
	UserSOSLockCar(context.Context, *zzzV1.UserSOSLockCarRequest, ...filterc.Filter) (*zzzV1.UserSOSLockCarReply, error)
	UserLockCar(context.Context, *zzzV1.UserLockCarRequest, ...filterc.Filter) (*zzzV1.UserLockCarReply, error)
	UserUnlockCar(context.Context, *zzzV1.UserUnlockCarRequest, ...filterc.Filter) (*zzzV1.UserUnlockCarReply, error)
	UserReturnCar(context.Context, *zzzV1.UserReturnCarRequest, ...filterc.Filter) (*zzzV1.UserReturnCarReply, error)
	UserAutoDriving(context.Context, *zzzV1.UserAutoDrivingRequest, ...filterc.Filter) (*zzzV1.UserAutoDrivingReply, error)
	UserManualDriving(context.Context, *zzzV1.UserManualDrivingRequest, ...filterc.Filter) (*zzzV1.UserManualDrivingReply, error)

	DevopsEnterMaintainMode(context.Context, *zzzV1.DevopsEnterMaintainModeRequest, ...filterc.Filter) (*zzzV1.DevopsEnterMaintainModeReply, error)
	DevopsExitMaintainMode(context.Context, *zzzV1.DevopsExitMaintainModeRequest, ...filterc.Filter) (*zzzV1.DevopsExitMaintainModeReply, error)
	DevopsUnlockCar(context.Context, *zzzV1.DevopsUnlockCarRequest, ...filterc.Filter) (*zzzV1.DevopsUnlockCarReply, error)
	DevopsLockCar(context.Context, *zzzV1.DevopsLockCarRequest, ...filterc.Filter) (*zzzV1.DevopsLockCarReply, error)
	DevopsReturnCar(context.Context, *zzzV1.DevopsReturnCarRequest, ...filterc.Filter) (*zzzV1.DevopsReturnCarReply, error)
	DevopsAutoDriving(context.Context, *zzzV1.DevopsAutoDrivingRequest, ...filterc.Filter) (*zzzV1.DevopsAutoDrivingReply, error)
	DevopsManualDriving(context.Context, *zzzV1.DevopsManualDrivingRequest, ...filterc.Filter) (*zzzV1.DevopsManualDrivingReply, error)

	AdminReturnCar(context.Context, *zzzV1.AdminReturnCarRequest, ...filterc.Filter) (*zzzV1.AdminReturnCarReply, error)
	AdminUpdateMapVersion(context.Context, *zzzV1.AdminUpdateMapVersionRequest, ...filterc.Filter) (*zzzV1.AdminUpdateMapVersionReply, error)
	AdminUpdateStaticConfigs(context.Context, *zzzV1.AdminUpdateStaticConfigsRequest, ...filterc.Filter) (*zzzV1.AdminUpdateStaticConfigsReply, error)
	AdminQueryLogBagFiles(context.Context, *zzzV1.AdminQueryLogBagFilesRequest, ...filterc.Filter) (*zzzV1.AdminQueryLogBagFilesReply, error)
	AdminPackUploadFiles(context.Context, *zzzV1.AdminUploadLogBagFilesRequest, ...filterc.Filter) (*zzzV1.AdminUploadLogBagFilesReply, error)
	AdminRemoteOperationLogin(context.Context, *zzzV1.AdminRemoteOperationLoginRequest, ...filterc.Filter) (*zzzV1.AdminRemoteOperationLoginReply, error)
	AdminRemoteOperationLogout(context.Context, *zzzV1.AdminRemoteOperationLogoutRequest, ...filterc.Filter) (*zzzV1.AdminRemoteOperationLogoutReply, error)

	TaskSystemReturnCar(context.Context, *zzzV1.TaskSystemReturnCarRequest, ...filterc.Filter) (*zzzV1.TaskSystemReturnCarReply, error)

	DispatchPlanPath(context.Context, *zzzV1.DispatchPlanPathRequest, ...filterc.Filter) (*zzzV1.DispatchPlanPathReply, error)
	DispatchStart(context.Context, *zzzV1.DispatchStartRequest, ...filterc.Filter) (*zzzV1.DispatchStartReply, error)
	DispatchRestart(context.Context, *zzzV1.DispatchRestartRequest, ...filterc.Filter) (*zzzV1.DispatchRestartReply, error)
	DispatchLockCar(context.Context, *zzzV1.DispatchLockCarRequest, ...filterc.Filter) (*zzzV1.DispatchLockCarReply, error)
}

type CarProxyGrpc struct {
	client IClient
}

func NewCarProxyGrpc(cli IClient) ICarProxy {
	return &CarProxyGrpc{
		client: cli,
	}
}

func (c *CarProxyGrpc) UserSOSLockCar(ctx context.Context, req *zzzV1.UserSOSLockCarRequest, filters ...filterc.Filter) (*zzzV1.UserSOSLockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserSOSLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserSOSLockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) UserLockCar(ctx context.Context, req *zzzV1.UserLockCarRequest, filters ...filterc.Filter) (*zzzV1.UserLockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserLockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) UserUnlockCar(ctx context.Context, req *zzzV1.UserUnlockCarRequest, filters ...filterc.Filter) (*zzzV1.UserUnlockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserUnlockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserUnlockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) UserReturnCar(ctx context.Context, req *zzzV1.UserReturnCarRequest, filters ...filterc.Filter) (*zzzV1.UserReturnCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserReturnCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) UserAutoDriving(ctx context.Context, req *zzzV1.UserAutoDrivingRequest, filters ...filterc.Filter) (*zzzV1.UserAutoDrivingReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserAutoDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserAutoDriving error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) UserManualDriving(ctx context.Context, req *zzzV1.UserManualDrivingRequest, filters ...filterc.Filter) (*zzzV1.UserManualDrivingReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserManualDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserManualDriving error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsEnterMaintainMode(ctx context.Context, req *zzzV1.DevopsEnterMaintainModeRequest, filters ...filterc.Filter) (*zzzV1.DevopsEnterMaintainModeReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsEnterMaintainMode(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsEnterMaintainMode error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsExitMaintainMode(ctx context.Context, req *zzzV1.DevopsExitMaintainModeRequest, filters ...filterc.Filter) (*zzzV1.DevopsExitMaintainModeReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsExitMaintainMode(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsExitMaintainMode error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsUnlockCar(ctx context.Context, req *zzzV1.DevopsUnlockCarRequest, filters ...filterc.Filter) (*zzzV1.DevopsUnlockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsUnlockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsUnlockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsLockCar(ctx context.Context, req *zzzV1.DevopsLockCarRequest, filters ...filterc.Filter) (*zzzV1.DevopsLockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsLockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsReturnCar(ctx context.Context, req *zzzV1.DevopsReturnCarRequest, filters ...filterc.Filter) (*zzzV1.DevopsReturnCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsReturnCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsAutoDriving(ctx context.Context, req *zzzV1.DevopsAutoDrivingRequest, filters ...filterc.Filter) (*zzzV1.DevopsAutoDrivingReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsAutoDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsAutoDriving error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsManualDriving(ctx context.Context, req *zzzV1.DevopsManualDrivingRequest, filters ...filterc.Filter) (*zzzV1.DevopsManualDrivingReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsManualDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsManualDriving error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminReturnCar(ctx context.Context, req *zzzV1.AdminReturnCarRequest, filters ...filterc.Filter) (*zzzV1.AdminReturnCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminReturnCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminUpdateMapVersion(ctx context.Context, req *zzzV1.AdminUpdateMapVersionRequest, filters ...filterc.Filter) (*zzzV1.AdminUpdateMapVersionReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminUpdateMapVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminUpdateMapVersion error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminUpdateStaticConfigs(ctx context.Context, req *zzzV1.AdminUpdateStaticConfigsRequest, filters ...filterc.Filter) (*zzzV1.AdminUpdateStaticConfigsReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminUpdateStaticConfigs(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminUpdateStaticConfigs error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminQueryLogBagFiles(ctx context.Context, req *zzzV1.AdminQueryLogBagFilesRequest, filters ...filterc.Filter) (*zzzV1.AdminQueryLogBagFilesReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminQueryLogBagFiles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminQueryLogBagFiles error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminPackUploadFiles(ctx context.Context, req *zzzV1.AdminUploadLogBagFilesRequest, filters ...filterc.Filter) (*zzzV1.AdminUploadLogBagFilesReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminPackUploadFiles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminPackUploadFiles error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminRemoteOperationLogin(ctx context.Context, req *zzzV1.AdminRemoteOperationLoginRequest, filters ...filterc.Filter) (*zzzV1.AdminRemoteOperationLoginReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminRemoteOperationLogin(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminRemoteOperationLogin error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminRemoteOperationLogout(ctx context.Context, req *zzzV1.AdminRemoteOperationLogoutRequest, filters ...filterc.Filter) (*zzzV1.AdminRemoteOperationLogoutReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminRemoteOperationLogout(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminRemoteOperationLogout error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) TaskSystemReturnCar(ctx context.Context, req *zzzV1.TaskSystemReturnCarRequest, filters ...filterc.Filter) (*zzzV1.TaskSystemReturnCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.TaskSystemReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.TaskSystemReturnCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) DispatchPlanPath(ctx context.Context, req *zzzV1.DispatchPlanPathRequest, filters ...filterc.Filter) (*zzzV1.DispatchPlanPathReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DispatchPlanPath(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchPlanPath error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) DispatchStart(ctx context.Context, req *zzzV1.DispatchStartRequest, filters ...filterc.Filter) (*zzzV1.DispatchStartReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DispatchStart(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchStart error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) DispatchRestart(ctx context.Context, req *zzzV1.DispatchRestartRequest, filters ...filterc.Filter) (*zzzV1.DispatchRestartReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DispatchRestart(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchRestart error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarProxyGrpc) DispatchLockCar(ctx context.Context, req *zzzV1.DispatchLockCarRequest, filters ...filterc.Filter) (*zzzV1.DispatchLockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DispatchLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchLockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}
