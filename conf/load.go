package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 从配置文件中加载
func LoadConfigFromToml(path string) error {
	cfg := newDefaultConfig()
	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		return err
	}

	setGlobalConfig(cfg)

	return nil
}

// 从环境变量中加载
func LoadConfigFromEnv() error {
	cfg := newDefaultConfig()
	err := env.Parse(&cfg)
	if err != nil {
		return err
	}

	setGlobalConfig(cfg)

	return nil
}
