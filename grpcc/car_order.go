package grpcc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/selector"
	orderV1 "github.com/jianbo-zh/jypb/api/carorder/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ICarOrder interface {
	GetOrder(context.Context, *orderV1.GetOrderRequest, ...selector.NodeFilter) (*orderV1.Order, error)
	GetGoingOrder(context.Context, *orderV1.GetGoingOrderRequest, ...selector.NodeFilter) (*orderV1.GetGoingOrderReply, error)
	GetOrderList(context.Context, *orderV1.GetOrderListRequest, ...selector.NodeFilter) (*orderV1.GetOrderListReply, error)
	CreateOrder(context.Context, *orderV1.CreateOrderRequest, ...selector.NodeFilter) (*orderV1.Order, error)
	CancelOrder(context.Context, *orderV1.CancelOrderRequest, ...selector.NodeFilter) (*orderV1.CancelOrderReply, error)
	StartOrder(context.Context, *orderV1.StartOrderRequest, ...selector.NodeFilter) (*orderV1.StartOrderReply, error)
	FinishOrder(context.Context, *orderV1.FinishOrderRequest, ...selector.NodeFilter) (*orderV1.FinishOrderReply, error)
	CommentOrder(context.Context, *orderV1.CommentOrderRequest, ...selector.NodeFilter) (*orderV1.CommentOrderReply, error)
	CheckOrderTimeoutAbort(context.Context, *orderV1.CheckOrderTimeoutAbortRequest, ...selector.NodeFilter) (*orderV1.CheckOrderTimeoutAbortReply, error)
	CheckOrderExpiredClose(context.Context, *orderV1.CheckOrderExpiredCloseRequest, ...selector.NodeFilter) (*orderV1.CheckOrderExpiredCloseReply, error)
	CheckOrderTimeoutCancel(context.Context, *orderV1.CheckOrderTimeoutCancelRequest, ...selector.NodeFilter) (*orderV1.CheckOrderTimeoutCancelReply, error)
	PrepayOrder(context.Context, *orderV1.PrepayOrderRequest, ...selector.NodeFilter) (*orderV1.PrepayOrderReply, error)
	CancelPayment(context.Context, *orderV1.CancelPaymentRequest, ...selector.NodeFilter) (*orderV1.CancelPaymentReply, error)
	LaunchOrderRefund(context.Context, *orderV1.LaunchOrderRefundRequest, ...selector.NodeFilter) (*orderV1.LaunchOrderRefundReply, error)
	ExecuteOrderRefund(context.Context, *orderV1.ExecuteOrderRefundRequest, ...selector.NodeFilter) (*orderV1.ExecuteOrderRefundReply, error)
	GetOrderRefunds(context.Context, *orderV1.GetOrderRefundsRequest, ...selector.NodeFilter) (*orderV1.GetOrderRefundsReply, error)
	GetOrderBilling(context.Context, *orderV1.GetOrderBillingRequest, ...selector.NodeFilter) (*orderV1.OrderBilling, error)
	ExecuteOrderSharing(context.Context, *orderV1.ExecuteOrderSharingRequest, ...selector.NodeFilter) (*orderV1.ExecuteOrderSharingReply, error)
	CheckOrderSharingResult(context.Context, *orderV1.CheckOrderSharingResultRequest, ...selector.NodeFilter) (*orderV1.CheckOrderSharingResultReply, error)
	EmitSOSEvent(context.Context, *orderV1.EmitSOSEventRequest, ...selector.NodeFilter) (*orderV1.EmitSOSEventReply, error)
	CancelSOSEvent(context.Context, *orderV1.CancelSOSEventRequest, ...selector.NodeFilter) (*orderV1.CancelSOSEventReply, error)
	CancelFlightOrders(context.Context, *orderV1.CancelFlightOrdersRequest, ...selector.NodeFilter) (*emptypb.Empty, error)
	FinishFlightOrders(context.Context, *orderV1.FinishFlightOrdersRequest, ...selector.NodeFilter) (*emptypb.Empty, error)
	GetUserCoupons(context.Context, *orderV1.GetUserCouponsRequest, ...selector.NodeFilter) (*orderV1.GetUserCouponsReply, error)
	GetFlightOrder(context.Context, *orderV1.GetFlightOrderRequest, ...selector.NodeFilter) (*orderV1.FlightOrder, error)
	CreateFlightOrder(context.Context, *orderV1.CreateFlightOrderRequest, ...selector.NodeFilter) (*orderV1.FlightOrder, error)
	StartFlightOrder(context.Context, *orderV1.StartFlightOrderRequest, ...selector.NodeFilter) (*orderV1.FlightOrder, error)
	GetFlightOrderList(context.Context, *orderV1.GetFlightOrderListRequest, ...selector.NodeFilter) (*orderV1.GetFlightOrderListReply, error)
	GetFlightOrderStops(context.Context, *orderV1.GetFlightOrderStopsRequest, ...selector.NodeFilter) (*orderV1.GetFlightOrderStopsReply, error)

	GetOrderAppeal(context.Context, *orderV1.GetOrderAppealRequest, ...selector.NodeFilter) (*orderV1.OrderAppeal, error)
	CreateOrderAppeal(context.Context, *orderV1.CreateOrderAppealRequest, ...selector.NodeFilter) (*orderV1.OrderAppeal, error)
	CancelOrderAppeal(context.Context, *orderV1.CancelOrderAppealRequest, ...selector.NodeFilter) (*emptypb.Empty, error)
	GetOrderAppealList(context.Context, *orderV1.GetOrderAppealListRequest, ...selector.NodeFilter) (*orderV1.GetOrderAppealListReply, error)
}

type CarOrderGrpc struct {
	client IClient
}

func NewCarOrderGrpc(cli IClient) ICarOrder {
	return &CarOrderGrpc{
		client: cli,
	}
}

func (c *CarOrderGrpc) GetOrder(ctx context.Context, req *orderV1.GetOrderRequest, filters ...selector.NodeFilter) (*orderV1.Order, error) {
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

func (c *CarOrderGrpc) GetGoingOrder(ctx context.Context, req *orderV1.GetGoingOrderRequest, filters ...selector.NodeFilter) (*orderV1.GetGoingOrderReply, error) {
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

func (c *CarOrderGrpc) GetOrderList(ctx context.Context, req *orderV1.GetOrderListRequest, filters ...selector.NodeFilter) (*orderV1.GetOrderListReply, error) {
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

func (c *CarOrderGrpc) CreateOrder(ctx context.Context, req *orderV1.CreateOrderRequest, filters ...selector.NodeFilter) (*orderV1.Order, error) {
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

func (c *CarOrderGrpc) CancelOrder(ctx context.Context, req *orderV1.CancelOrderRequest, filters ...selector.NodeFilter) (*orderV1.CancelOrderReply, error) {
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

func (c *CarOrderGrpc) StartOrder(ctx context.Context, req *orderV1.StartOrderRequest, filters ...selector.NodeFilter) (*orderV1.StartOrderReply, error) {
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

func (c *CarOrderGrpc) FinishOrder(ctx context.Context, req *orderV1.FinishOrderRequest, filters ...selector.NodeFilter) (*orderV1.FinishOrderReply, error) {
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

func (c *CarOrderGrpc) CommentOrder(ctx context.Context, req *orderV1.CommentOrderRequest, filters ...selector.NodeFilter) (*orderV1.CommentOrderReply, error) {
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

func (c *CarOrderGrpc) CheckOrderTimeoutAbort(ctx context.Context, req *orderV1.CheckOrderTimeoutAbortRequest, filters ...selector.NodeFilter) (*orderV1.CheckOrderTimeoutAbortReply, error) {
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

func (c *CarOrderGrpc) CheckOrderExpiredClose(ctx context.Context, req *orderV1.CheckOrderExpiredCloseRequest, filters ...selector.NodeFilter) (*orderV1.CheckOrderExpiredCloseReply, error) {
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

func (c *CarOrderGrpc) CheckOrderTimeoutCancel(ctx context.Context, req *orderV1.CheckOrderTimeoutCancelRequest, filters ...selector.NodeFilter) (*orderV1.CheckOrderTimeoutCancelReply, error) {
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

func (c *CarOrderGrpc) PrepayOrder(ctx context.Context, req *orderV1.PrepayOrderRequest, filters ...selector.NodeFilter) (*orderV1.PrepayOrderReply, error) {
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

func (c *CarOrderGrpc) CancelPayment(ctx context.Context, req *orderV1.CancelPaymentRequest, filters ...selector.NodeFilter) (*orderV1.CancelPaymentReply, error) {
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

func (c *CarOrderGrpc) LaunchOrderRefund(ctx context.Context, req *orderV1.LaunchOrderRefundRequest, filters ...selector.NodeFilter) (*orderV1.LaunchOrderRefundReply, error) {
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

func (c *CarOrderGrpc) ExecuteOrderRefund(ctx context.Context, req *orderV1.ExecuteOrderRefundRequest, filters ...selector.NodeFilter) (*orderV1.ExecuteOrderRefundReply, error) {
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

func (c *CarOrderGrpc) GetOrderRefunds(ctx context.Context, req *orderV1.GetOrderRefundsRequest, filters ...selector.NodeFilter) (*orderV1.GetOrderRefundsReply, error) {
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

func (c *CarOrderGrpc) GetOrderBilling(ctx context.Context, req *orderV1.GetOrderBillingRequest, filters ...selector.NodeFilter) (*orderV1.OrderBilling, error) {
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

func (c *CarOrderGrpc) ExecuteOrderSharing(ctx context.Context, req *orderV1.ExecuteOrderSharingRequest, filters ...selector.NodeFilter) (*orderV1.ExecuteOrderSharingReply, error) {
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

func (c *CarOrderGrpc) CheckOrderSharingResult(ctx context.Context, req *orderV1.CheckOrderSharingResultRequest, filters ...selector.NodeFilter) (*orderV1.CheckOrderSharingResultReply, error) {
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

func (c *CarOrderGrpc) EmitSOSEvent(ctx context.Context, req *orderV1.EmitSOSEventRequest, filters ...selector.NodeFilter) (*orderV1.EmitSOSEventReply, error) {
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

func (c *CarOrderGrpc) CancelSOSEvent(ctx context.Context, req *orderV1.CancelSOSEventRequest, filters ...selector.NodeFilter) (*orderV1.CancelSOSEventReply, error) {
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

func (c *CarOrderGrpc) CancelFlightOrders(ctx context.Context, req *orderV1.CancelFlightOrdersRequest, filters ...selector.NodeFilter) (*emptypb.Empty, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.CancelFlightOrders(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CancelFlightOrder error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) FinishFlightOrders(ctx context.Context, req *orderV1.FinishFlightOrdersRequest, filters ...selector.NodeFilter) (*emptypb.Empty, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.FinishFlightOrders(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.FinishFlightOrder error: %w", errors.FromError(err))
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetUserCoupons(ctx context.Context, req *orderV1.GetUserCouponsRequest, filters ...selector.NodeFilter) (*orderV1.GetUserCouponsReply, error) {
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

func (c *CarOrderGrpc) GetFlightOrder(ctx context.Context, req *orderV1.GetFlightOrderRequest, filters ...selector.NodeFilter) (*orderV1.FlightOrder, error) {
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

func (c *CarOrderGrpc) CreateFlightOrder(ctx context.Context, req *orderV1.CreateFlightOrderRequest, filters ...selector.NodeFilter) (*orderV1.FlightOrder, error) {
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

func (c *CarOrderGrpc) StartFlightOrder(ctx context.Context, req *orderV1.StartFlightOrderRequest, filters ...selector.NodeFilter) (*orderV1.FlightOrder, error) {
	cli, err := c.client.CarOrderClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.CarOrderClient error: %w", err)
	}
	reply, err := cli.StartFlightOrder(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.StartFlightOrder error: %w", err)
	}
	return reply, nil
}

func (c *CarOrderGrpc) GetFlightOrderList(ctx context.Context, req *orderV1.GetFlightOrderListRequest, filters ...selector.NodeFilter) (*orderV1.GetFlightOrderListReply, error) {
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

func (c *CarOrderGrpc) GetFlightOrderStops(ctx context.Context, req *orderV1.GetFlightOrderStopsRequest, filters ...selector.NodeFilter) (*orderV1.GetFlightOrderStopsReply, error) {
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

func (c *CarOrderGrpc) GetOrderAppeal(ctx context.Context, req *orderV1.GetOrderAppealRequest, filters ...selector.NodeFilter) (*orderV1.OrderAppeal, error) {
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

func (c *CarOrderGrpc) CreateOrderAppeal(ctx context.Context, req *orderV1.CreateOrderAppealRequest, filters ...selector.NodeFilter) (*orderV1.OrderAppeal, error) {
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

func (c *CarOrderGrpc) CancelOrderAppeal(ctx context.Context, req *orderV1.CancelOrderAppealRequest, filters ...selector.NodeFilter) (*emptypb.Empty, error) {
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

func (c *CarOrderGrpc) GetOrderAppealList(ctx context.Context, req *orderV1.GetOrderAppealListRequest, filters ...selector.NodeFilter) (*orderV1.GetOrderAppealListReply, error) {
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
