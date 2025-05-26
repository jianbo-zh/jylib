package filterc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/selector"
)

// Metadata is metadata filter.
func Metadata(md map[string]string) selector.NodeFilter {
	return func(_ context.Context, nodes []selector.Node) []selector.Node {
		newNodes := make([]selector.Node, 0, len(nodes))
		for _, n := range nodes {
			fmt.Println(n.ServiceName(), n.Address())
			isMatch := true
			for k, v := range md {
				if n.Metadata()[k] != v {
					isMatch = false
					break
				}
			}
			if isMatch {
				newNodes = append(newNodes, n)
			}
		}
		return newNodes
	}
}

// Version is version filter.
func Version(version string) selector.NodeFilter {
	return func(_ context.Context, nodes []selector.Node) []selector.Node {
		newNodes := make([]selector.Node, 0, len(nodes))
		for _, n := range nodes {
			if n.Version() == version {
				newNodes = append(newNodes, n)
			}
		}
		return newNodes
	}
}
