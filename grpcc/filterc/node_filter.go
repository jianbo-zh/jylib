package filterc

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/selector"
)

func WithMetadata0(md map[string]string) selector.NodeFilter {
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
