package grpcc

import (
	"context"
	"fmt"

	proxyV1 "github.com/jianbo-zh/jypb/api/carproxy/v1"
)

type ICarProxy interface {
	UserSOSLockCar(context.Context, *proxyV1.UserSOSLockCarRequest) (*proxyV1.UserSOSLockCarReply, error)
	UserLockCar(context.Context, *proxyV1.UserLockCarRequest) (*proxyV1.UserLockCarReply, error)
	UserUnlockCar(context.Context, *proxyV1.UserUnlockCarRequest) (*proxyV1.UserUnlockCarReply, error)
	UserReturnCar(context.Context, *proxyV1.UserReturnCarRequest) (*proxyV1.UserReturnCarReply, error)
	UserAutoDriving(context.Context, *proxyV1.UserAutoDrivingRequest) (*proxyV1.UserAutoDrivingReply, error)
	UserManualDriving(context.Context, *proxyV1.UserManualDrivingRequest) (*proxyV1.UserManualDrivingReply, error)
	DevopsEnterMaintainMode(context.Context, *proxyV1.DevopsEnterMaintainModeRequest) (*proxyV1.DevopsEnterMaintainModeReply, error)
	DevopsExitMaintainMode(context.Context, *proxyV1.DevopsExitMaintainModeRequest) (*proxyV1.DevopsExitMaintainModeReply, error)
	DevopsUnlockCar(context.Context, *proxyV1.DevopsUnlockCarRequest) (*proxyV1.DevopsUnlockCarReply, error)
	DevopsLockCar(context.Context, *proxyV1.DevopsLockCarRequest) (*proxyV1.DevopsLockCarReply, error)
	DevopsReturnCar(context.Context, *proxyV1.DevopsReturnCarRequest) (*proxyV1.DevopsReturnCarReply, error)
	DevopsAutoDriving(context.Context, *proxyV1.DevopsAutoDrivingRequest) (*proxyV1.DevopsAutoDrivingReply, error)
	DevopsManualDriving(context.Context, *proxyV1.DevopsManualDrivingRequest) (*proxyV1.DevopsManualDrivingReply, error)
	AdminReturnCar(context.Context, *proxyV1.AdminReturnCarRequest) (*proxyV1.AdminReturnCarReply, error)
	AdminUpdateMapVersion(context.Context, *proxyV1.AdminUpdateMapVersionRequest) (*proxyV1.AdminUpdateMapVersionReply, error)
	AdminUpdateStaticConfigs(context.Context, *proxyV1.AdminUpdateStaticConfigsRequest) (*proxyV1.AdminUpdateStaticConfigsReply, error)
	AdminQueryLogBagFiles(context.Context, *proxyV1.AdminQueryLogBagFilesRequest) (*proxyV1.AdminQueryLogBagFilesReply, error)
	AdminPackUploadFiles(context.Context, *proxyV1.AdminUploadLogBagFilesRequest) (*proxyV1.AdminUploadLogBagFilesReply, error)
	AdminRemoteOperationLogin(context.Context, *proxyV1.AdminRemoteOperationLoginRequest) (*proxyV1.AdminRemoteOperationLoginReply, error)
	AdminRemoteOperationLogout(context.Context, *proxyV1.AdminRemoteOperationLogoutRequest) (*proxyV1.AdminRemoteOperationLogoutReply, error)
	TaskSystemReturnCar(context.Context, *proxyV1.TaskSystemReturnCarRequest) (*proxyV1.TaskSystemReturnCarReply, error)
	DispatchPlanPath(context.Context, *proxyV1.DispatchPlanPathRequest) (*proxyV1.DispatchPlanPathReply, error)
	DispatchStart(context.Context, *proxyV1.DispatchStartRequest) (*proxyV1.DispatchStartReply, error)
	DispatchRestart(context.Context, *proxyV1.DispatchRestartRequest) (*proxyV1.DispatchRestartReply, error)
	DispatchLockCar(context.Context, *proxyV1.DispatchLockCarRequest) (*proxyV1.DispatchLockCarReply, error)
}

type CarProxyGrpc struct {
	client IClient
}

func NewCarProxyGrpc(cli IClient) ICarProxy {
	return &CarProxyGrpc{
		client: cli,
	}
}

func (c *CarProxyGrpc) UserSOSLockCar(ctx context.Context, req *proxyV1.UserSOSLockCarRequest) (*proxyV1.UserSOSLockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserSOSLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserSOSLockCar error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) UserLockCar(ctx context.Context, req *proxyV1.UserLockCarRequest) (*proxyV1.UserLockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserLockCar error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) UserUnlockCar(ctx context.Context, req *proxyV1.UserUnlockCarRequest) (*proxyV1.UserUnlockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserUnlockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserUnlockCar error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) UserReturnCar(ctx context.Context, req *proxyV1.UserReturnCarRequest) (*proxyV1.UserReturnCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserReturnCar error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) UserAutoDriving(ctx context.Context, req *proxyV1.UserAutoDrivingRequest) (*proxyV1.UserAutoDrivingReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserAutoDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserAutoDriving error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) UserManualDriving(ctx context.Context, req *proxyV1.UserManualDrivingRequest) (*proxyV1.UserManualDrivingReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.UserManualDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UserManualDriving error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsEnterMaintainMode(ctx context.Context, req *proxyV1.DevopsEnterMaintainModeRequest) (*proxyV1.DevopsEnterMaintainModeReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsEnterMaintainMode(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsEnterMaintainMode error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsExitMaintainMode(ctx context.Context, req *proxyV1.DevopsExitMaintainModeRequest) (*proxyV1.DevopsExitMaintainModeReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsExitMaintainMode(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsExitMaintainMode error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsUnlockCar(ctx context.Context, req *proxyV1.DevopsUnlockCarRequest) (*proxyV1.DevopsUnlockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsUnlockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsUnlockCar error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsLockCar(ctx context.Context, req *proxyV1.DevopsLockCarRequest) (*proxyV1.DevopsLockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsLockCar error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsReturnCar(ctx context.Context, req *proxyV1.DevopsReturnCarRequest) (*proxyV1.DevopsReturnCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsReturnCar error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsAutoDriving(ctx context.Context, req *proxyV1.DevopsAutoDrivingRequest) (*proxyV1.DevopsAutoDrivingReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsAutoDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsAutoDriving error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) DevopsManualDriving(ctx context.Context, req *proxyV1.DevopsManualDrivingRequest) (*proxyV1.DevopsManualDrivingReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DevopsManualDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DevopsManualDriving error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminReturnCar(ctx context.Context, req *proxyV1.AdminReturnCarRequest) (*proxyV1.AdminReturnCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminReturnCar error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminUpdateMapVersion(ctx context.Context, req *proxyV1.AdminUpdateMapVersionRequest) (*proxyV1.AdminUpdateMapVersionReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminUpdateMapVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminUpdateMapVersion error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminUpdateStaticConfigs(ctx context.Context, req *proxyV1.AdminUpdateStaticConfigsRequest) (*proxyV1.AdminUpdateStaticConfigsReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminUpdateStaticConfigs(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminUpdateStaticConfigs error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminQueryLogBagFiles(ctx context.Context, req *proxyV1.AdminQueryLogBagFilesRequest) (*proxyV1.AdminQueryLogBagFilesReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminQueryLogBagFiles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminQueryLogBagFiles error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminPackUploadFiles(ctx context.Context, req *proxyV1.AdminUploadLogBagFilesRequest) (*proxyV1.AdminUploadLogBagFilesReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminPackUploadFiles(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminPackUploadFiles error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminRemoteOperationLogin(ctx context.Context, req *proxyV1.AdminRemoteOperationLoginRequest) (*proxyV1.AdminRemoteOperationLoginReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminRemoteOperationLogin(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminRemoteOperationLogin error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) AdminRemoteOperationLogout(ctx context.Context, req *proxyV1.AdminRemoteOperationLogoutRequest) (*proxyV1.AdminRemoteOperationLogoutReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.AdminRemoteOperationLogout(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AdminRemoteOperationLogout error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) TaskSystemReturnCar(ctx context.Context, req *proxyV1.TaskSystemReturnCarRequest) (*proxyV1.TaskSystemReturnCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.TaskSystemReturnCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.TaskSystemReturnCar error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) DispatchPlanPath(ctx context.Context, req *proxyV1.DispatchPlanPathRequest) (*proxyV1.DispatchPlanPathReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DispatchPlanPath(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchPlanPath error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) DispatchStart(ctx context.Context, req *proxyV1.DispatchStartRequest) (*proxyV1.DispatchStartReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DispatchStart(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchStart error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) DispatchRestart(ctx context.Context, req *proxyV1.DispatchRestartRequest) (*proxyV1.DispatchRestartReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DispatchRestart(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchRestart error: %w", err)
	}
	return reply, nil
}

func (c *CarProxyGrpc) DispatchLockCar(ctx context.Context, req *proxyV1.DispatchLockCarRequest) (*proxyV1.DispatchLockCarReply, error) {
	cli, err := c.client.CarProxyClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarProxyClient error: %w", err)
	}
	reply, err := cli.DispatchLockCar(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.DispatchLockCar error: %w", err)
	}
	return reply, nil
}
