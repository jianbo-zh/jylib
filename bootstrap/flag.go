package bootstrap

import "flag"

type CommandFlags struct {
	Env        string
	ConfigPath string
	ConfigType string
	ConfigHost string
	ConfigKey  string
}

func NewCommandFlags() *CommandFlags {
	return &CommandFlags{
		Env:        "",
		ConfigPath: "",
		ConfigType: "",
		ConfigHost: "",
		ConfigKey:  "",
	}
}

func (f *CommandFlags) Init() {
	flag.StringVar(&f.ConfigPath, "conf", "./config.yaml", "config path, eg: -conf bootstrap.yaml")
	flag.StringVar(&f.Env, "env", "dev", "runtime environment, eg: -env dev")
	flag.StringVar(&f.ConfigType, "ctype", "consul", "config server host, eg: -ctype consul")
	flag.StringVar(&f.ConfigHost, "chost", "127.0.0.1:8500", "config server host, eg: -chost 127.0.0.1:8500")
	flag.StringVar(&f.ConfigKey, "ckey", "/config", "config key path, eg: -ckey /config")
}
