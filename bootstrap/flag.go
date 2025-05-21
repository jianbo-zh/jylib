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
	flag.StringVar(&f.ConfigPath, "conf", "./configs/config.yaml", "config path, eg: -conf bootstrap.yaml")
	flag.StringVar(&f.Env, "env", "dev", "runtime environment, eg: -env dev")
	flag.StringVar(&f.ConfigType, "ctype", "", "config server host, eg: -ctype consul")
	flag.StringVar(&f.ConfigHost, "chost", "", "config server host, eg: -chost 127.0.0.1:8500")
	flag.StringVar(&f.ConfigKey, "ckey", "", "config key path, eg: -ckey /config")
}
