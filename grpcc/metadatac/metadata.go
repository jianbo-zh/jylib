package metadatac

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/metadata"
	zzzV1 "github.com/jianbo-zh/jypb/api/zzz/v1"
)

// 上下文元数据常量
const (
	Ctx_Origin   = "x-md-global-jyorigin"
	Ctx_Role     = "x-md-global-jyrole"
	Ctx_UserId   = "x-md-global-jyuserid"
	Ctx_ClientIp = "x-md-global-jyclientip"
)

// 其他元数据常量
const (
	Key_InstanceId            = "instance-id"
	Key_ApolloVersion         = "apollo-version"
	Key_CommunicationProtocol = "communication-protocol"
)

// WithOperator
func WithOperator(ctx context.Context, operator *zzzV1.Operator) context.Context {
	md := metadata.New(map[string][]string{
		Ctx_Origin:   {operator.Origin.String()},
		Ctx_Role:     {operator.Role.String()},
		Ctx_UserId:   {strconv.Itoa(int(operator.UserId))},
		Ctx_ClientIp: {operator.ClientIp},
	})
	return metadata.MergeToClientContext(ctx, md)
}

// ExtractOperator
func ExtractOperator(ctx context.Context) *zzzV1.Operator {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return nil
	}

	userIdStr := md.Get(Ctx_UserId)
	if userIdStr == "" {
		userIdStr = "0"
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		userId = -1
	}

	return &zzzV1.Operator{
		Origin:   zzzV1.Operator_Origin(zzzV1.Operator_Origin_value[md.Get(Ctx_Origin)]),
		Role:     zzzV1.Operator_Role(zzzV1.Operator_Role_value[md.Get(Ctx_Role)]),
		UserId:   int32(userId),
		ClientIp: md.Get(Ctx_ClientIp),
	}
}
