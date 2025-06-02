package bootstrap

import (
	"fmt"
	"time"

	"github.com/jianbo-zh/jylib/grpcc/metadatac"
)

var Service *ServiceInfo

type ServiceInfo struct {
	Id       string
	Name     string
	Version  string
	Metadata map[string]string
}

func RegisterServiceInfo(name string, version string, metadata map[string]string) {
	id := fmt.Sprintf("%d", time.Now().UnixMilli())

	if metadata == nil {
		metadata = make(map[string]string)
	}
	metadata[metadatac.Key_InstanceId] = id

	Service = &ServiceInfo{
		Id:       id,
		Name:     name,
		Version:  version,
		Metadata: metadata,
	}
}

func (s *ServiceInfo) GetInstanceId() string {
	return s.Name + "." + s.Id
}

func (s *ServiceInfo) SetMataData(k, v string) {
	if s.Metadata == nil {
		s.Metadata = make(map[string]string)
	}
	s.Metadata[k] = v
}
