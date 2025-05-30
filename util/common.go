package util

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPbTimestamp0(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func ToPbTimestamp1(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func ToPtr[T any](abc T) *T {
	return &abc
}

func ToPtrOrNil[T any](ok bool, abc T) *T {
	if ok {
		return &abc
	}
	return nil
}

func FromPtr[T any](abc *T) T {
	var val T
	if abc == nil {
		return val
	}
	return *abc
}

func IfThen[T any](cond bool, val1 T, val2 T) T {
	if cond {
		return val1
	}
	return val2
}

func InSlice[T comparable](item T, items []T) bool {
	if items == nil {
		return false
	}

	for _, val := range items {
		if item == val {
			return true
		}
	}
	return false
}

func IntSlice[T int32 | uint32 | int64](abc []T) []int {
	var ints []int
	for _, val := range abc {
		ints = append(ints, int(val))
	}
	return ints
}

func IntTSlice[T int32 | uint32 | int64 | uint64](abc []int) []T {
	var ints []T
	for _, val := range abc {
		ints = append(ints, T(val))
	}
	return ints
}

func ClientIP(ctx context.Context) (string, error) {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return "", fmt.Errorf("transport.FromServerContext fail")
	}

	return tr.RequestHeader().Get("X-Real-IP"), nil
}

func PrototextMarshal(m proto.Message) ([]byte, error) {
	return prototext.MarshalOptions{Multiline: true}.Marshal(m)
}

func PrototextUnmarshal(b []byte, m proto.Message) error {
	return prototext.Unmarshal(b, m)
}

func IgnoreError[T any](a T, _ error) T {
	return a
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
