package grpcc

import (
	"context"
	"fmt"

	authV1 "github.com/jianbo-zh/jypb/api/carauth/v1"
)

type ICarAuth interface {
	CheckAuth(context.Context, *authV1.CheckAuthRequest) (*authV1.CheckAuthReply, error)
}

type CarAuthGrpc struct {
	client IClient
}

func NewCarAuthGrpc(cli IClient) ICarAuth {
	return &CarAuthGrpc{
		client: cli,
	}
}

func (c *CarAuthGrpc) CheckAuth(ctx context.Context, req *authV1.CheckAuthRequest) (*authV1.CheckAuthReply, error) {
	cli, err := c.client.CarAuthClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("c.CarAuthClient error: %w", err)
	}
	reply, err := cli.CheckAuth(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cli.CheckAuth error: %w", err)
	}
	return reply, nil
}
