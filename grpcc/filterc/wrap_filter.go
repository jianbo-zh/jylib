package filterc

import (
	"context"
	"crypto/md5"
	"fmt"
	"jiuyouzx/pkg/util"

	"github.com/go-kratos/kratos/v2/selector"
)

type Filter interface {
	Filter() selector.NodeFilter
	Uuid() string
}

func WithMetadata(md map[string]string) Filter {
	return util.ToPtr(Metadata(md))
}

func WithVersion(version string) Filter {
	return util.ToPtr(Version(version))
}

type Metadata map[string]string

func (f *Metadata) Filter() selector.NodeFilter {
	return func(_ context.Context, nodes []selector.Node) []selector.Node {
		newNodes := make([]selector.Node, 0, len(nodes))
		for _, n := range nodes {
			fmt.Println(n.ServiceName(), n.Address())
			isMatch := true
			for k, v := range *f {
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

func (f *Metadata) Uuid() string {
	var str string
	for k, v := range *f {
		str += k + v
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

type Version string

func (f *Version) Filter() selector.NodeFilter {
	return func(_ context.Context, nodes []selector.Node) []selector.Node {
		newNodes := make([]selector.Node, 0, len(nodes))
		for _, n := range nodes {
			if n.Version() == string(*f) {
				newNodes = append(newNodes, n)
			}
		}
		return newNodes
	}
}

func (f *Version) Uuid() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(*f)))
}
