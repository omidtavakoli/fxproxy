package main

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
	"sidecar/internal/logger"
	"sidecar/internal/proxy"
	"sidecar/pkg/app"
)

type ServerConfig struct {
	User        string `yaml:"USER"`
	Pass        string `yaml:"PASS"`
	AuthEnabled bool   `yaml:"AUTH_ENABLED"`
	Port        int32  `yaml:"PORT"`
}

type MainConfig struct {
	Logger logger.Config `yaml:"LOGGER"`
	Proxy  proxy.Config  `yaml:"PROXY"`
	Server ServerConfig  `yaml:"SERVER"`
	App    app.Config    `yaml:"APP"`
}

// LoadConfig loads configs form provided yaml file or overrides it with env variables
func LoadConfig(filePath string) (*MainConfig, error) {
	cfg := MainConfig{}
	if filePath != "" {
		err := readFile(&cfg, filePath)
		if err != nil {
			return nil, err
		}
	}
	err := readEnv(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func readFile(cfg *MainConfig, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func readEnv(cfg *MainConfig) error {
	return envconfig.Process("", cfg)
}
