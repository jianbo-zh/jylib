package grpcc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jianbo-zh/jylib/grpcc/filterc"
	msgV1 "github.com/jianbo-zh/jypb/api/message/v1"
)

type IMessage interface {
	AppPush(context.Context, *msgV1.AppPushRequest, ...filterc.Filter) (*msgV1.AppPushReply, error)
}

type MessageGrpc struct {
	client IClient
}

func NewMessageGrpc(cli IClient) IMessage {
	return &MessageGrpc{
		client: cli,
	}
}

func (c *MessageGrpc) AppPush(ctx context.Context, req *msgV1.AppPushRequest, filters ...filterc.Filter) (*msgV1.AppPushReply, error) {
	cli, err := c.client.MessageClient(ctx, filters...)
	if err != nil {
		return nil, fmt.Errorf("c.MessageClient error: %w", err)
	}
	reply, err := cli.AppPush(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AppPush error: %w", errors.FromError(err))
	}
	return reply, nil
}
