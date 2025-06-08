package grpcc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	orderV1 "github.com/jianbo-zh/jypb/api/carorder/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ICarOrder interface {
	GetOrder(context.Context, *orderV1.GetOrderRequest, ...filterc.Filter) (*orderV1.Order, error)
	GetGoingOrder(context.Context, *orderV1.GetGoingOrderRequest, ...filterc.Filter) (*orderV1.GetGoingOrderReply, error)
	GetOrderList(context.Context, *orderV1.GetOrderListRequest, ...filterc.Filter) (*orderV1.GetOrderListReply, error)
	CreateOrder(context.Context, *orderV1.CreateOrderRequest, ...filterc.Filter) (*orderV1.Order, error)
	CancelOrder(context.Context, *orderV1.CancelOrderRequest, ...filterc.Filter) (*orderV1.CancelOrderReply, error)
	StartOrder(context.Context, *orderV1.StartOrderRequest, ...filterc.Filter) (*orderV1.StartOrderReply, error)
	FinishOrder(context.Context, *orderV1.FinishOrderRequest, ...filterc.Filter) (*orderV1.FinishOrderReply, error)
	CommentOrder(context.Context, *orderV1.CommentOrderRequest, ...filterc.Filter) (*orderV1.CommentOrderReply, error)
	CheckOrderTimeoutAbort(context.Context, *orderV1.CheckOrderTimeoutAbortRequest, ...filterc.Filter) (*orderV1.CheckOrderTimeoutAbortReply, error)
	CheckOrderExpiredClose(context.Context, *orderV1.CheckOrderExpiredCloseRequest, ...filterc.Filter) (*orderV1.CheckOrderExpiredCloseReply, error)
	CheckOrderTimeoutCancel(context.Context, *orderV1.CheckOrderTimeoutCancelRequest, ...filterc.Filter) (*orderV1.CheckOrderTimeoutCancelReply, error)
	PrepayOrder(context.Context, *orderV1.PrepayOrderRequest, ...filterc.Filter) (*orderV1.PrepayOrderReply, error)
	CancelPayment(context.Context, *orderV1.CancelPaymentRequest, ...filterc.Filter) (*orderV1.CancelPaymentReply, error)
	LaunchOrderRefund(context.Context, *orderV1.LaunchOrderRefundRequest, ...filterc.Filter) (*orderV1.LaunchOrderRefundReply, error)
	ExecuteOrderRefund(context.Context, *orderV1.ExecuteOrderRefundRequest, ...filterc.Filter) (*orderV1.ExecuteOrderRefundReply, error)
	GetOrderRefunds(context.Context, *orderV1.GetOrderRefundsRequest, ...filterc.Filter) (*orderV1.GetOrderRefundsReply, error)
	GetOrderBilling(context.Context, *orderV1.GetOrderBillingRequest, ...filterc.Filter) (*orderV1.OrderBilling, error)
	ExecuteOrderSharing(context.Context, *orderV1.ExecuteOrderSharingRequest, ...filterc.Filter) (*orderV1.ExecuteOrderSharingReply, error)
	CheckOrderSharingResult(context.Context, *orderV1.CheckOrderSharingResultRequest, ...filterc.Filter) (*orderV1.CheckOrderSharingResultReply, error)
	EmitSOSEvent(context.Context, *orderV1.EmitSOSEventRequest, ...filterc.Filter) (*orderV1.EmitSOSEventReply, error)
	CancelSOSEvent(context.Context, *orderV1.CancelSOSEventRequest, ...filterc.Filter) (*orderV1.CancelSOSEventReply, error)
	EstimateFlightOrderAmount(context.Context, *orderV1.EstimateFlightOrderAmountRequest, ...filterc.Filter) (*orderV1.EstimateFlightOrderAmountReply, error)
	CancelUnuseFlightOrders(context.Context, *orderV1.CancelUnuseFlightOrdersRequest, ...filterc.Filter) (*emptypb.Empty, error)
	CancelInuseFlightOrders(context.Context, *orderV1.CancelInuseFlightOrdersRequest, ...filterc.Filter) (*emptypb.Empty, error)
	FinishInuseFlightOrders(context.Context, *orderV1.FinishInuseFlightOrdersRequest, ...filterc.Filter) (*emptypb.Empty, error)
	GetUserCoupons(context.Context, *orderV1.GetUserCouponsRequest, ...filterc.Filter) (*orderV1.GetUserCouponsReply, error)
	GetFlightOrder(context.Context, *orderV1.GetFlightOrderRequest, ...filterc.Filter) (*orderV1.FlightOrder, error)
	CreateFlightOrder(context.Context, *orderV1.CreateFlightOrderRequest, ...filterc.Filter) (*orderV1.FlightOrder, error)
	GetFlightOrderList(context.Context, *orderV1.GetFlightOrderListRequest, ...filterc.Filter) (*orderV1.GetFlightOrderListReply, error)
	GetFlightOrderStops(context.Context, *orderV1.GetFlightOrderStopsRequest, ...filterc.Filter) (*orderV1.GetFlightOrderStopsReply, error)

	GetOrderAppeal(context.Context, *orderV1.GetOrderAppealRequest, ...filterc.Filter) (*orderV1.OrderAppeal, error)
	CreateOrderAppeal(context.Context, *orderV1.CreateOrderAppealRequest, ...filterc.Filter) (*orderV1.OrderAppeal, error)
	CancelOrderAppeal(context.Context, *orderV1.CancelOrderAppealRequest, ...filterc.Filter) (*emptypb.Empty, error)
	GetOrderAppealList(context.Context, *orderV1.GetOrderAppealListRequest, ...filterc.Filter) (*orderV1.GetOrderAppealListReply, error)
}

type CarOrderGrpc struct {
	client IClient
}

func NewCarOrderGrpc(cli IClient) ICarOrder {
	return &CarOrderGrpc{
		client: cli,
	}
}

func (c *CarOrderGrpc) GetOrder(ctx context.Context, req *orderV1.GetOrderRequest, filters ...filterc.Filter) (*orderV1.Order, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.GetOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetOrder error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetGoingOrder(ctx context.Context, req *orderV1.GetGoingOrderRequest, filters ...filterc.Filter) (*orderV1.GetGoingOrderReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.GetGoingOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetGoingOrder error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetOrderList(ctx context.Context, req *orderV1.GetOrderListRequest, filters ...filterc.Filter) (*orderV1.GetOrderListReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.GetOrderList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetOrderList error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) CreateOrder(ctx context.Context, req *orderV1.CreateOrderRequest, filters ...filterc.Filter) (*orderV1.Order, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CreateOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreateOrder error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) CancelOrder(ctx context.Context, req *orderV1.CancelOrderRequest, filters ...filterc.Filter) (*orderV1.CancelOrderReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CancelOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CancelOrder error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) StartOrder(ctx context.Context, req *orderV1.StartOrderRequest, filters ...filterc.Filter) (*orderV1.StartOrderReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.StartOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.StartOrder error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) FinishOrder(ctx context.Context, req *orderV1.FinishOrderRequest, filters ...filterc.Filter) (*orderV1.FinishOrderReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.FinishOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.FinishOrder error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) CommentOrder(ctx context.Context, req *orderV1.CommentOrderRequest, filters ...filterc.Filter) (*orderV1.CommentOrderReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CommentOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CommentOrder error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) CheckOrderTimeoutAbort(ctx context.Context, req *orderV1.CheckOrderTimeoutAbortRequest, filters ...filterc.Filter) (*orderV1.CheckOrderTimeoutAbortReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CheckOrderTimeoutAbort(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CheckOrderTimeoutAbort error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) CheckOrderExpiredClose(ctx context.Context, req *orderV1.CheckOrderExpiredCloseRequest, filters ...filterc.Filter) (*orderV1.CheckOrderExpiredCloseReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CheckOrderExpiredClose(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CheckOrderExpiredClose error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) CheckOrderTimeoutCancel(ctx context.Context, req *orderV1.CheckOrderTimeoutCancelRequest, filters ...filterc.Filter) (*orderV1.CheckOrderTimeoutCancelReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CheckOrderTimeoutCancel(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CheckOrderTimeoutCancel error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) PrepayOrder(ctx context.Context, req *orderV1.PrepayOrderRequest, filters ...filterc.Filter) (*orderV1.PrepayOrderReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.PrepayOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.PrepayOrder error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) CancelPayment(ctx context.Context, req *orderV1.CancelPaymentRequest, filters ...filterc.Filter) (*orderV1.CancelPaymentReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CancelPayment(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CancelPayment error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) LaunchOrderRefund(ctx context.Context, req *orderV1.LaunchOrderRefundRequest, filters ...filterc.Filter) (*orderV1.LaunchOrderRefundReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.LaunchOrderRefund(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.LaunchOrderRefund error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) ExecuteOrderRefund(ctx context.Context, req *orderV1.ExecuteOrderRefundRequest, filters ...filterc.Filter) (*orderV1.ExecuteOrderRefundReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.ExecuteOrderRefund(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.ExecuteOrderRefund error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetOrderRefunds(ctx context.Context, req *orderV1.GetOrderRefundsRequest, filters ...filterc.Filter) (*orderV1.GetOrderRefundsReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.GetOrderRefunds(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetOrderRefunds error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetOrderBilling(ctx context.Context, req *orderV1.GetOrderBillingRequest, filters ...filterc.Filter) (*orderV1.OrderBilling, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.GetOrderBilling(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetOrderBilling error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) ExecuteOrderSharing(ctx context.Context, req *orderV1.ExecuteOrderSharingRequest, filters ...filterc.Filter) (*orderV1.ExecuteOrderSharingReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.ExecuteOrderSharing(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.ExecuteOrderSharing error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) CheckOrderSharingResult(ctx context.Context, req *orderV1.CheckOrderSharingResultRequest, filters ...filterc.Filter) (*orderV1.CheckOrderSharingResultReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CheckOrderSharingResult(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CheckOrderSharingResult error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) EmitSOSEvent(ctx context.Context, req *orderV1.EmitSOSEventRequest, filters ...filterc.Filter) (*orderV1.EmitSOSEventReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.EmitSOSEvent(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EmitSOSEvent error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) CancelSOSEvent(ctx context.Context, req *orderV1.CancelSOSEventRequest, filters ...filterc.Filter) (*orderV1.CancelSOSEventReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CancelSOSEvent(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CancelSOSEvent error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) EstimateFlightOrderAmount(ctx context.Context, req *orderV1.EstimateFlightOrderAmountRequest, filters ...filterc.Filter) (*orderV1.EstimateFlightOrderAmountReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.EstimateFlightOrderAmount(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.EstimateFlightOrderAmount error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) CancelUnuseFlightOrders(ctx context.Context, req *orderV1.CancelUnuseFlightOrdersRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CancelUnuseFlightOrders(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CancelUnuseFlightOrders error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) CancelInuseFlightOrders(ctx context.Context, req *orderV1.CancelInuseFlightOrdersRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CancelInuseFlightOrders(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CancelInuseFlightOrders error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) FinishInuseFlightOrders(ctx context.Context, req *orderV1.FinishInuseFlightOrdersRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.FinishInuseFlightOrders(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.FinishInuseFlightOrders error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetUserCoupons(ctx context.Context, req *orderV1.GetUserCouponsRequest, filters ...filterc.Filter) (*orderV1.GetUserCouponsReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.GetUserCoupons(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetUserCoupons error: %w", err)
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetFlightOrder(ctx context.Context, req *orderV1.GetFlightOrderRequest, filters ...filterc.Filter) (*orderV1.FlightOrder, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.GetFlightOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFlightOrder error: %w", err)
	}
	return reply, nil
}

func (c *CarOrderGrpc) CreateFlightOrder(ctx context.Context, req *orderV1.CreateFlightOrderRequest, filters ...filterc.Filter) (*orderV1.FlightOrder, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CreateFlightOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreateFlightOrder error: %w", err)
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetFlightOrderList(ctx context.Context, req *orderV1.GetFlightOrderListRequest, filters ...filterc.Filter) (*orderV1.GetFlightOrderListReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.GetFlightOrderList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFlightOrderList error: %w", err)
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetFlightOrderStops(ctx context.Context, req *orderV1.GetFlightOrderStopsRequest, filters ...filterc.Filter) (*orderV1.GetFlightOrderStopsReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.GetFlightOrderStops(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetFlightOrderStops error: %w", err)
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetOrderAppeal(ctx context.Context, req *orderV1.GetOrderAppealRequest, filters ...filterc.Filter) (*orderV1.OrderAppeal, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.GetOrderAppeal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetOrderAppeal error: %w", err)
	}
	return reply, nil
}

func (c *CarOrderGrpc) CreateOrderAppeal(ctx context.Context, req *orderV1.CreateOrderAppealRequest, filters ...filterc.Filter) (*orderV1.OrderAppeal, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CreateOrderAppeal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CreateOrderAppeal error: %w", err)
	}
	return reply, nil
}

func (c *CarOrderGrpc) CancelOrderAppeal(ctx context.Context, req *orderV1.CancelOrderAppealRequest, filters ...filterc.Filter) (*emptypb.Empty, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CancelOrderAppeal(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CancelOrderAppeal error: %w", err)
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetOrderAppealList(ctx context.Context, req *orderV1.GetOrderAppealListRequest, filters ...filterc.Filter) (*orderV1.GetOrderAppealListReply, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.GetOrderAppealList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.GetOrderAppealList error: %w", err)
	}
	return reply, nil
}
