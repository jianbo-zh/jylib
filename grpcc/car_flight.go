package grpcc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/selector"
	carflightV1 "github.com/jianbo-zh/jypb/api/carflight/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ICarFlight interface {
	CreateFlight(context.Context, *carflightV1.CreateFlightRequest, ...selector.NodeFilter) (*carflightV1.Flight, error)
	StartFlight(context.Context, *carflightV1.StartFlightRequest, ...selector.NodeFilter) (*emptypb.Empty, error)
	EndFlight(context.Context, *carflightV1.EndFlightRequest, ...selector.NodeFilter) (*emptypb.Empty, error)
	TempParking(context.Context, *carflightV1.TempParkingRequest, ...selector.NodeFilter) (*emptypb.Empty, error)
	KeepDriving(context.Context, *carflightV1.KeepDrivingRequest, ...selector.NodeFilter) (*emptypb.Empty, error)
	ArrivalStop(context.Context, *carflightV1.ArrivalStopRequest, ...selector.NodeFilter) (*emptypb.Empty, error)
	GetFlight(context.Context, *carflightV1.GetFlightRequest, ...selector.NodeFilter) (*carflightV1.Flight, error)
	GetOptionalFlights(context.Context, *carflightV1.GetOptionalFlightsRequest, ...selector.NodeFilter) (*carflightV1.GetOptionalFlightsReply, error)
	GetFlightTrace(context.Context, *carflightV1.GetFlightTraceRequest, ...selector.NodeFilter) (*carflightV1.GetFlightTraceReply, error)
	GetRoutePath(context.Context, *carflightV1.GetRoutePathRequest, ...selector.NodeFilter) (*carflightV1.GetRoutePathReply, error)
	UpdateFlightBooking(context.Context, *carflightV1.UpdateFlightBookingRequest, ...selector.NodeFilter) (*emptypb.Empty, error)
}

type CarFlightGrpc struct {
	client IClient
}

func NewCarFlightGrpc(cli IClient) ICarFlight {
	return &CarFlightGrpc{
		client: cli,
	}
}

func (c *CarFlightGrpc) CreateFlight(ctx context.Context, req *carflightV1.CreateFlightRequest, filters ...selector.NodeFilter) (*carflightV1.Flight, error) {
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

func (c *CarFlightGrpc) StartFlight(ctx context.Context, req *carflightV1.StartFlightRequest, filters ...selector.NodeFilter) (*emptypb.Empty, error) {
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

func (c *CarFlightGrpc) EndFlight(ctx context.Context, req *carflightV1.EndFlightRequest, filters ...selector.NodeFilter) (*emptypb.Empty, error) {
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

func (c *CarFlightGrpc) TempParking(ctx context.Context, req *carflightV1.TempParkingRequest, filters ...selector.NodeFilter) (*emptypb.Empty, error) {
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

func (c *CarFlightGrpc) KeepDriving(ctx context.Context, req *carflightV1.KeepDrivingRequest, filters ...selector.NodeFilter) (*emptypb.Empty, error) {
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

func (c *CarFlightGrpc) ArrivalStop(ctx context.Context, req *carflightV1.ArrivalStopRequest, filters ...selector.NodeFilter) (*emptypb.Empty, error) {
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

func (c *CarFlightGrpc) GetFlight(ctx context.Context, req *carflightV1.GetFlightRequest, filters ...selector.NodeFilter) (*carflightV1.Flight, error) {
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

func (c *CarFlightGrpc) GetOptionalFlights(ctx context.Context, req *carflightV1.GetOptionalFlightsRequest, filters ...selector.NodeFilter) (*carflightV1.GetOptionalFlightsReply, error) {
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

func (c *CarFlightGrpc) GetFlightTrace(ctx context.Context, req *carflightV1.GetFlightTraceRequest, filters ...selector.NodeFilter) (*carflightV1.GetFlightTraceReply, error) {
	cli, err := c.client.CarFlightClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarFlightClient error: %w", err)
	}
	reply, err := cli.GetFlightTrace(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFlightTrace error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarFlightGrpc) GetRoutePath(ctx context.Context, req *carflightV1.GetRoutePathRequest, filters ...selector.NodeFilter) (*carflightV1.GetRoutePathReply, error) {
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

func (c *CarFlightGrpc) UpdateFlightBooking(ctx context.Context, req *carflightV1.UpdateFlightBookingRequest, filters ...selector.NodeFilter) (*emptypb.Empty, error) {
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
