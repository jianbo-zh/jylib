package grpcc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	proxyV1 "github.com/jianbo-zh/jypb/api/carproxy/v1"
)

type CarControlGrpc struct {
	client IClient
}

func NewCarControlGrpc(cli IClient) ICarProxy {
	return &CarControlGrpc{
		client: cli,
	}
}

func (c *CarControlGrpc) UserSOSLockCar(ctx context.Context, req *proxyV1.UserSOSLockCarRequest, filters ...filterc.Filter) (*proxyV1.UserSOSLockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserSOSLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserSOSLockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) UserLockCar(ctx context.Context, req *proxyV1.UserLockCarRequest, filters ...filterc.Filter) (*proxyV1.UserLockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserLockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) UserUnlockCar(ctx context.Context, req *proxyV1.UserUnlockCarRequest, filters ...filterc.Filter) (*proxyV1.UserUnlockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserUnlockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserUnlockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) UserReturnCar(ctx context.Context, req *proxyV1.UserReturnCarRequest, filters ...filterc.Filter) (*proxyV1.UserReturnCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserReturnCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) UserAutoDriving(ctx context.Context, req *proxyV1.UserAutoDrivingRequest, filters ...filterc.Filter) (*proxyV1.UserAutoDrivingReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserAutoDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserAutoDriving error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) UserManualDriving(ctx context.Context, req *proxyV1.UserManualDrivingRequest, filters ...filterc.Filter) (*proxyV1.UserManualDrivingReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserManualDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserManualDriving error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsEnterMaintainMode(ctx context.Context, req *proxyV1.DevopsEnterMaintainModeRequest, filters ...filterc.Filter) (*proxyV1.DevopsEnterMaintainModeReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsEnterMaintainMode(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsEnterMaintainMode error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsExitMaintainMode(ctx context.Context, req *proxyV1.DevopsExitMaintainModeRequest, filters ...filterc.Filter) (*proxyV1.DevopsExitMaintainModeReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsExitMaintainMode(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsExitMaintainMode error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsUnlockCar(ctx context.Context, req *proxyV1.DevopsUnlockCarRequest, filters ...filterc.Filter) (*proxyV1.DevopsUnlockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsUnlockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsUnlockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsLockCar(ctx context.Context, req *proxyV1.DevopsLockCarRequest, filters ...filterc.Filter) (*proxyV1.DevopsLockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsLockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsReturnCar(ctx context.Context, req *proxyV1.DevopsReturnCarRequest, filters ...filterc.Filter) (*proxyV1.DevopsReturnCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsReturnCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsAutoDriving(ctx context.Context, req *proxyV1.DevopsAutoDrivingRequest, filters ...filterc.Filter) (*proxyV1.DevopsAutoDrivingReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsAutoDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsAutoDriving error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsManualDriving(ctx context.Context, req *proxyV1.DevopsManualDrivingRequest, filters ...filterc.Filter) (*proxyV1.DevopsManualDrivingReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsManualDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsManualDriving error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminReturnCar(ctx context.Context, req *proxyV1.AdminReturnCarRequest, filters ...filterc.Filter) (*proxyV1.AdminReturnCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminReturnCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminUpdateMapVersion(ctx context.Context, req *proxyV1.AdminUpdateMapVersionRequest, filters ...filterc.Filter) (*proxyV1.AdminUpdateMapVersionReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminUpdateMapVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminUpdateMapVersion error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminUpdateStaticConfigs(ctx context.Context, req *proxyV1.AdminUpdateStaticConfigsRequest, filters ...filterc.Filter) (*proxyV1.AdminUpdateStaticConfigsReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminUpdateStaticConfigs(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminUpdateStaticConfigs error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminQueryLogBagFiles(ctx context.Context, req *proxyV1.AdminQueryLogBagFilesRequest, filters ...filterc.Filter) (*proxyV1.AdminQueryLogBagFilesReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminQueryLogBagFiles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminQueryLogBagFiles error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminPackUploadFiles(ctx context.Context, req *proxyV1.AdminUploadLogBagFilesRequest, filters ...filterc.Filter) (*proxyV1.AdminUploadLogBagFilesReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminPackUploadFiles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminPackUploadFiles error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminRemoteOperationLogin(ctx context.Context, req *proxyV1.AdminRemoteOperationLoginRequest, filters ...filterc.Filter) (*proxyV1.AdminRemoteOperationLoginReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminRemoteOperationLogin(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminRemoteOperationLogin error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminRemoteOperationLogout(ctx context.Context, req *proxyV1.AdminRemoteOperationLogoutRequest, filters ...filterc.Filter) (*proxyV1.AdminRemoteOperationLogoutReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminRemoteOperationLogout(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminRemoteOperationLogout error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) TaskSystemReturnCar(ctx context.Context, req *proxyV1.TaskSystemReturnCarRequest, filters ...filterc.Filter) (*proxyV1.TaskSystemReturnCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.TaskSystemReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.TaskSystemReturnCar error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DispatchPlanPath(ctx context.Context, req *proxyV1.DispatchPlanPathRequest, filters ...filterc.Filter) (*proxyV1.DispatchPlanPathReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DispatchPlanPath(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchPlanPath error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DispatchStart(ctx context.Context, req *proxyV1.DispatchStartRequest, filters ...filterc.Filter) (*proxyV1.DispatchStartReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DispatchStart(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchStart error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DispatchRestart(ctx context.Context, req *proxyV1.DispatchRestartRequest, filters ...filterc.Filter) (*proxyV1.DispatchRestartReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DispatchRestart(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchRestart error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DispatchLockCar(ctx context.Context, req *proxyV1.DispatchLockCarRequest, filters ...filterc.Filter) (*proxyV1.DispatchLockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DispatchLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchLockCar error: %w", errors.FromError(err))
	}
	return reply, nil
}
