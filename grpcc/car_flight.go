package grpcc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	carflightV1 "github.com/jianbo-zh/jypb/api/carflight/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ICarFlight interface {
	CreateFlight(context.Context, *carflightV1.CreateFlightRequest, ...filterc.Filter) (*carflightV1.Flight, error)
	StartFlight(context.Context, *carflightV1.StartFlightRequest, ...filterc.Filter) (*emptypb.Empty, error)
	EndFlight(context.Context, *carflightV1.EndFlightRequest, ...filterc.Filter) (*emptypb.Empty, error)
	TempParking(context.Context, *carflightV1.TempParkingRequest, ...filterc.Filter) (*emptypb.Empty, error)
	KeepDriving(context.Context, *carflightV1.KeepDrivingRequest, ...filterc.Filter) (*emptypb.Empty, error)
	ArrivalStop(context.Context, *carflightV1.ArrivalStopRequest, ...filterc.Filter) (*emptypb.Empty, error)
	GotoNextStop(context.Context, *carflightV1.GotoNextStopRequest, ...filterc.Filter) (*emptypb.Empty, error)
	GetFlight(context.Context, *carflightV1.GetFlightRequest, ...filterc.Filter) (*carflightV1.Flight, error)
	GetOptionalFlights(context.Context, *carflightV1.GetOptionalFlightsRequest, ...filterc.Filter) (*carflightV1.GetOptionalFlightsReply, error)
	GetFlightPath(context.Context, *carflightV1.GetFlightPathRequest, ...filterc.Filter) (*carflightV1.GetFlightPathReply, error)
	GetRoutePath(context.Context, *carflightV1.GetRoutePathRequest, ...filterc.Filter) (*carflightV1.GetRoutePathReply, error)
	UpdateFlightBooking(context.Context, *carflightV1.UpdateFlightBookingRequest, ...filterc.Filter) (*emptypb.Empty, error)
	UpdateFlightCoord(context.Context, *carflightV1.UpdateFlightCoordRequest, ...filterc.Filter) (*emptypb.Empty, error)

	// 调用zelos查询车辆详情
	GetYokeeCarDetail(context.Context, *carflightV1.GetYokeeCarDetailRequest, ...filterc.Filter) (*carflightV1.GetYokeeCarDetailReply, error)
}

type CarFlightGrpc struct {
	client IClient
}

func NewCarFlightGrpc(cli IClient) ICarFlight {
	return &CarFlightGrpc{
		client: cli,
	}
}

func (c *CarFlightGrpc) CreateFlight(ctx context.Context, req *carflightV1.CreateFlightRequest, filters ...filterc.Filter) (*carflightV1.Flight, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.CreateFlight(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreateFlight error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) StartFlight(ctx context.Context, req *carflightV1.StartFlightRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.StartFlight(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.StartFlight error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) EndFlight(ctx context.Context, req *carflightV1.EndFlightRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.EndFlight(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EndFlight error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) TempParking(ctx context.Context, req *carflightV1.TempParkingRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.TempParking(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.TempParking error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) KeepDriving(ctx context.Context, req *carflightV1.KeepDrivingRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.KeepDriving(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.KeepDriving error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) ArrivalStop(ctx context.Context, req *carflightV1.ArrivalStopRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.ArrivalStop(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.ArrivalStop error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) GotoNextStop(ctx context.Context, req *carflightV1.GotoNextStopRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.GotoNextStop(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GotoNextStop error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) GetFlight(ctx context.Context, req *carflightV1.GetFlightRequest, filters ...filterc.Filter) (*carflightV1.Flight, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.GetFlight(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFlight error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) GetOptionalFlights(ctx context.Context, req *carflightV1.GetOptionalFlightsRequest, filters ...filterc.Filter) (*carflightV1.GetOptionalFlightsReply, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.GetOptionalFlights(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetOptionalFlights error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) GetFlightPath(ctx context.Context, req *carflightV1.GetFlightPathRequest, filters ...filterc.Filter) (*carflightV1.GetFlightPathReply, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.GetFlightPath(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFlightPath error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) GetRoutePath(ctx context.Context, req *carflightV1.GetRoutePathRequest, filters ...filterc.Filter) (*carflightV1.GetRoutePathReply, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.GetRoutePath(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetRoutePath error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) UpdateFlightBooking(ctx context.Context, req *carflightV1.UpdateFlightBookingRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.UpdateFlightBooking(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UpdateFlightBooking error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) UpdateFlightCoord(ctx context.Context, req *carflightV1.UpdateFlightCoordRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.UpdateFlightCoord(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.UpdateFlightCoord error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) GetYokeeCarDetail(ctx context.Context, req *carflightV1.GetYokeeCarDetailRequest, filters ...filterc.Filter) (*carflightV1.GetYokeeCarDetailReply, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.GetYokeeCarDetail(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetYokeeCarDetail error: %w", errors.FromError(err))
	}
	return reply, nil
}
