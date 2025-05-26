package bootstrap

import (
	"github.com/google/uuid"
	"github.com/jianbo-zh/jylib/grpcc/metadatac"
)

var Service *ServiceInfo

type ServiceInfo struct {
	Id       string
	Name     string
	Version  string
	Metadata map[string]string
}

func RegisterServiceInfo(id string, name string, version string, metadata map[string]string) {
	if id == "" {
		id = uuid.NewString()
	}

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
	return s.Id + "." + s.Name
}

func (s *ServiceInfo) SetMataData(k, v string) {
	if s.Metadata == nil {
		s.Metadata = make(map[string]string)
	}
	s.Metadata[k] = v
}
