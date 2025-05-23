package bootstrap

import (
	"flag"
	"log"

	"github.com/go-kratos/kratos/v2/config"
	configV1 "github.com/jianbo-zh/jypb/config/v1"
)

var Flags *CommandFlags

func LoadConfig() *configV1.Bootstrap {
	Flags = NewCommandFlags()
	Flags.Init()
	flag.Parse()

	var sources []config.Source
	if Flags.ConfigPath != "" {
		sources = append(sources, NewFileConfigSource(Flags.ConfigPath))
	}

	if Flags.ConfigType != "" && Flags.ConfigHost != "" && Flags.ConfigKey != "" {
		if source := NewRemoteConfigSource(Flags.ConfigType, Flags.ConfigHost, Flags.ConfigKey); source != nil {
			sources = append(sources, source)
		}
	}

	conf := config.New(
		config.WithSource(sources...),
	)

	var bc configV1.Bootstrap
	if err := conf.Load(); err != nil {
		log.Fatal(err)
	}
	if err := conf.Scan(&bc); err != nil {
		panic(err)
	}
	return &bc
}
