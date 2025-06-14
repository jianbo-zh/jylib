package grpcc

import (
	"context"
	"fmt"

	"github.com/jianbo-zh/jylib/errc"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	dispatchV1 "github.com/jianbo-zh/jypb/api/cardispatch/v1"
)

type ICarDispatch interface {
	CheckCarSchedulability(context.Context, *dispatchV1.CheckCarSchedulabilityRequest, ...filterc.Filter) (*dispatchV1.CheckCarSchedulabilityReply, error)
	CheckCarReachability(context.Context, *dispatchV1.CheckCarReachabilityRequest, ...filterc.Filter) (*dispatchV1.CheckCarReachabilityReply, error)
	StartScheTask(context.Context, *dispatchV1.StartScheTaskRequest, ...filterc.Filter) (*dispatchV1.ScheTask, error)
	PauseScheTask(context.Context, *dispatchV1.PauseScheTaskRequest, ...filterc.Filter) (*dispatchV1.PauseScheTaskReply, error)
	RestartScheTask(context.Context, *dispatchV1.RestartScheTaskRequest, ...filterc.Filter) (*dispatchV1.RestartScheTaskReply, error)
	GetScheTask(context.Context, *dispatchV1.GetScheTaskRequest, ...filterc.Filter) (*dispatchV1.ScheTask, error)
	GetScheTaskEvents(context.Context, *dispatchV1.GetScheTaskEventsRequest, ...filterc.Filter) (*dispatchV1.GetScheTaskEventsReply, error)
	CancelScheTask(context.Context, *dispatchV1.CancelScheTaskRequest, ...filterc.Filter) (*dispatchV1.CancelScheTaskReply, error)
	FinishScheTask(context.Context, *dispatchV1.FinishScheTaskRequest, ...filterc.Filter) (*dispatchV1.FinishScheTaskReply, error)
}

type CarDispatchGrpc struct {
	client IClient
}

func NewCarDispatchGrpc(cli IClient) ICarDispatch {
	return &CarDispatchGrpc{
		client: cli,
	}
}

func (c *CarDispatchGrpc) CheckCarSchedulability(ctx context.Context, req *dispatchV1.CheckCarSchedulabilityRequest, filters ...filterc.Filter) (*dispatchV1.CheckCarSchedulabilityReply, error) {
	cli, err := c.client.CarDispatchClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarDispatchClient error: %w", err)
	}
	reply, err := cli.CheckCarSchedulability(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CheckCarSchedulability error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarDispatchGrpc) CheckCarReachability(ctx context.Context, req *dispatchV1.CheckCarReachabilityRequest, filters ...filterc.Filter) (*dispatchV1.CheckCarReachabilityReply, error) {
	cli, err := c.client.CarDispatchClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarDispatchClient error: %w", err)
	}
	reply, err := cli.CheckCarReachability(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CheckCarReachability error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarDispatchGrpc) StartScheTask(ctx context.Context, req *dispatchV1.StartScheTaskRequest, filters ...filterc.Filter) (*dispatchV1.ScheTask, error) {
	cli, err := c.client.CarDispatchClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarDispatchClient error: %w", err)
	}
	reply, err := cli.StartScheTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.StartScheTask error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarDispatchGrpc) PauseScheTask(ctx context.Context, req *dispatchV1.PauseScheTaskRequest, filters ...filterc.Filter) (*dispatchV1.PauseScheTaskReply, error) {
	cli, err := c.client.CarDispatchClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarDispatchClient error: %w", err)
	}
	reply, err := cli.PauseScheTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.PauseScheTask error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarDispatchGrpc) RestartScheTask(ctx context.Context, req *dispatchV1.RestartScheTaskRequest, filters ...filterc.Filter) (*dispatchV1.RestartScheTaskReply, error) {
	cli, err := c.client.CarDispatchClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarDispatchClient error: %w", err)
	}
	reply, err := cli.RestartScheTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.RestartScheTask error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarDispatchGrpc) GetScheTask(ctx context.Context, req *dispatchV1.GetScheTaskRequest, filters ...filterc.Filter) (*dispatchV1.ScheTask, error) {
	cli, err := c.client.CarDispatchClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarDispatchClient error: %w", err)
	}
	reply, err := cli.GetScheTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetScheTask error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarDispatchGrpc) GetScheTaskEvents(ctx context.Context, req *dispatchV1.GetScheTaskEventsRequest, filters ...filterc.Filter) (*dispatchV1.GetScheTaskEventsReply, error) {
	cli, err := c.client.CarDispatchClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarDispatchClient error: %w", err)
	}
	reply, err := cli.GetScheTaskEvents(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetScheTaskEvents error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarDispatchGrpc) CancelScheTask(ctx context.Context, req *dispatchV1.CancelScheTaskRequest, filters ...filterc.Filter) (*dispatchV1.CancelScheTaskReply, error) {
	cli, err := c.client.CarDispatchClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarDispatchClient error: %w", err)
	}
	reply, err := cli.CancelScheTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CancelScheTask error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}

func (c *CarDispatchGrpc) FinishScheTask(ctx context.Context, req *dispatchV1.FinishScheTaskRequest, filters ...filterc.Filter) (*dispatchV1.FinishScheTaskReply, error) {
	cli, err := c.client.CarDispatchClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarDispatchClient error: %w", err)
	}
	reply, err := cli.FinishScheTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.FinishScheTask error: %w", errc.FromGrpcServiceError(err))
	}
	return reply, nil
}
