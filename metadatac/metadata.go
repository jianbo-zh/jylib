package metadata

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/metadata"
	zzzV1 "github.com/jianbo-zh/jypb/api/zzz/v1"
)

const (
	Metadata_Origin   = "x-md-global-jyorigin"
	Metadata_Role     = "x-md-global-jyrole"
	Metadata_UserId   = "x-md-global-jyuserid"
	Metadata_ClientIp = "x-md-global-jyclientip"
)

// WithOperator
func WithOperator(ctx context.Context, operator *zzzV1.Operator) context.Context {
	md := metadata.New(map[string][]string{
		Metadata_Origin:   {operator.Origin.String()},
		Metadata_Role:     {operator.Role.String()},
		Metadata_UserId:   {strconv.Itoa(int(operator.UserId))},
		Metadata_ClientIp: {operator.ClientIp},
	})
	return metadata.MergeToClientContext(ctx, md)
}

// ExtractOperator
func ExtractOperator(ctx context.Context) *zzzV1.Operator {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return nil
	}

	userIdStr := md.Get(Metadata_UserId)
	if userIdStr == "" {
		userIdStr = "0"
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		userId = -1
	}

	return &zzzV1.Operator{
		Origin:   zzzV1.Operator_Origin(zzzV1.Operator_Origin_value[md.Get(Metadata_Origin)]),
		Role:     zzzV1.Operator_Role(zzzV1.Operator_Role_value[md.Get(Metadata_Role)]),
		UserId:   int32(userId),
		ClientIp: md.Get(Metadata_ClientIp),
	}
}
