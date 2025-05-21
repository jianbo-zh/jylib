package grpcc

import (
	"context"
	"fmt"

	msgV1 "github.com/jianbo-zh/jypb/api/message/v1"
)

type IMessage interface {
	AppPush(context.Context, *msgV1.AppPushRequest) (*msgV1.AppPushReply, error)
}

type MessageGrpc struct {
	client IClient
}

func NewMessageGrpc(cli IClient) IMessage {
	return &MessageGrpc{
		client: cli,
	}
}

func (c *MessageGrpc) AppPush(ctx context.Context, req *msgV1.AppPushRequest) (*msgV1.AppPushReply, error) {
	cli, err := c.client.MessageClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.MessageClient error: %w", err)
	}
	reply, err := cli.AppPush(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.AppPush error: %w", err)
	}
	return reply, nil
}
