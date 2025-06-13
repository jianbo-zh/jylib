package grpcc

import (
	"context"
	"fmt"

	"github.com/jianbo-zh/jylib/errc"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	zzzV1 "github.com/jianbo-zh/jypb/api/zzz/v1"
)

type CarControlGrpc struct {
	client IClient
}

func NewCarControlGrpc(cli IClient) ICarProxy {
	return &CarControlGrpc{
		client: cli,
	}
}

func (c *CarControlGrpc) UserSOSLockCar(ctx context.Context, req *zzzV1.UserSOSLockCarRequest, filters ...filterc.Filter) (*zzzV1.UserSOSLockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserSOSLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserSOSLockCar error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) UserLockCar(ctx context.Context, req *zzzV1.UserLockCarRequest, filters ...filterc.Filter) (*zzzV1.UserLockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserLockCar error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) UserUnlockCar(ctx context.Context, req *zzzV1.UserUnlockCarRequest, filters ...filterc.Filter) (*zzzV1.UserUnlockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserUnlockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserUnlockCar error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) UserReturnCar(ctx context.Context, req *zzzV1.UserReturnCarRequest, filters ...filterc.Filter) (*zzzV1.UserReturnCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserReturnCar error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) UserAutoDriving(ctx context.Context, req *zzzV1.UserAutoDrivingRequest, filters ...filterc.Filter) (*zzzV1.UserAutoDrivingReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserAutoDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserAutoDriving error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) UserManualDriving(ctx context.Context, req *zzzV1.UserManualDrivingRequest, filters ...filterc.Filter) (*zzzV1.UserManualDrivingReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.UserManualDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserManualDriving error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsEnterMaintainMode(ctx context.Context, req *zzzV1.DevopsEnterMaintainModeRequest, filters ...filterc.Filter) (*zzzV1.DevopsEnterMaintainModeReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsEnterMaintainMode(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsEnterMaintainMode error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsExitMaintainMode(ctx context.Context, req *zzzV1.DevopsExitMaintainModeRequest, filters ...filterc.Filter) (*zzzV1.DevopsExitMaintainModeReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsExitMaintainMode(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsExitMaintainMode error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsUnlockCar(ctx context.Context, req *zzzV1.DevopsUnlockCarRequest, filters ...filterc.Filter) (*zzzV1.DevopsUnlockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsUnlockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsUnlockCar error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsLockCar(ctx context.Context, req *zzzV1.DevopsLockCarRequest, filters ...filterc.Filter) (*zzzV1.DevopsLockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsLockCar error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsReturnCar(ctx context.Context, req *zzzV1.DevopsReturnCarRequest, filters ...filterc.Filter) (*zzzV1.DevopsReturnCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsReturnCar error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsAutoDriving(ctx context.Context, req *zzzV1.DevopsAutoDrivingRequest, filters ...filterc.Filter) (*zzzV1.DevopsAutoDrivingReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsAutoDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsAutoDriving error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DevopsManualDriving(ctx context.Context, req *zzzV1.DevopsManualDrivingRequest, filters ...filterc.Filter) (*zzzV1.DevopsManualDrivingReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DevopsManualDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsManualDriving error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminReturnCar(ctx context.Context, req *zzzV1.AdminReturnCarRequest, filters ...filterc.Filter) (*zzzV1.AdminReturnCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminReturnCar error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminUpdateMapVersion(ctx context.Context, req *zzzV1.AdminUpdateMapVersionRequest, filters ...filterc.Filter) (*zzzV1.AdminUpdateMapVersionReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminUpdateMapVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminUpdateMapVersion error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminUpdateStaticConfigs(ctx context.Context, req *zzzV1.AdminUpdateStaticConfigsRequest, filters ...filterc.Filter) (*zzzV1.AdminUpdateStaticConfigsReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminUpdateStaticConfigs(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminUpdateStaticConfigs error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminQueryLogBagFiles(ctx context.Context, req *zzzV1.AdminQueryLogBagFilesRequest, filters ...filterc.Filter) (*zzzV1.AdminQueryLogBagFilesReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminQueryLogBagFiles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminQueryLogBagFiles error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminPackUploadFiles(ctx context.Context, req *zzzV1.AdminUploadLogBagFilesRequest, filters ...filterc.Filter) (*zzzV1.AdminUploadLogBagFilesReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminPackUploadFiles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminPackUploadFiles error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminRemoteOperationLogin(ctx context.Context, req *zzzV1.AdminRemoteOperationLoginRequest, filters ...filterc.Filter) (*zzzV1.AdminRemoteOperationLoginReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminRemoteOperationLogin(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminRemoteOperationLogin error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) AdminRemoteOperationLogout(ctx context.Context, req *zzzV1.AdminRemoteOperationLogoutRequest, filters ...filterc.Filter) (*zzzV1.AdminRemoteOperationLogoutReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.AdminRemoteOperationLogout(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminRemoteOperationLogout error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) TaskSystemReturnCar(ctx context.Context, req *zzzV1.TaskSystemReturnCarRequest, filters ...filterc.Filter) (*zzzV1.TaskSystemReturnCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.TaskSystemReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.TaskSystemReturnCar error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) TaskSystemUpdateFlightSeat(ctx context.Context, req *zzzV1.TaskSystemUpdateFlightSeatRequest, filters ...filterc.Filter) (*zzzV1.TaskSystemUpdateFlightSeatReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.TaskSystemUpdateFlightSeat(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.TaskSystemUpdateFlightSeat error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DispatchPlanPath(ctx context.Context, req *zzzV1.DispatchPlanPathRequest, filters ...filterc.Filter) (*zzzV1.DispatchPlanPathReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DispatchPlanPath(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchPlanPath error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DispatchStart(ctx context.Context, req *zzzV1.DispatchStartRequest, filters ...filterc.Filter) (*zzzV1.DispatchStartReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DispatchStart(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchStart error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DispatchRestart(ctx context.Context, req *zzzV1.DispatchRestartRequest, filters ...filterc.Filter) (*zzzV1.DispatchRestartReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DispatchRestart(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchRestart error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarControlGrpc) DispatchLockCar(ctx context.Context, req *zzzV1.DispatchLockCarRequest, filters ...filterc.Filter) (*zzzV1.DispatchLockCarReply, error) {
	cli, err := c.client.CarControlClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarControlClient error: %w", err)
	}
	reply, err := cli.DispatchLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchLockCar error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}
