package hosting

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type HostConfig interface {
	Load(configPath string) error
	GetHttpConfig() HttpConfig
	GetGrpcConfig() GrpcConfig
}

type HttpConfig interface {
	GetBindAddress() string
}

type GrpcConfig interface {
	GetBindAddress() string
}

func newHostConfig() HostConfig {
	return &hostConfig{
		Http: httpConfig{
			Bind: "0.0.0.0",
			Port: "80",
		},
		Grpc: grpcConfig{
			Bind: "0.0.0.0",
			Port: "81",
		},
	}
}

type hostConfig struct {
	Http httpConfig `yaml:"http"`
	Grpc grpcConfig `yaml:"grpc"`
}

type httpConfig struct {
	Bind string `yaml:"bind"`
	Port string `yaml:"port"`
}

type grpcConfig struct {
	Bind string `yaml:"bind"`
	Port string `yaml:"port"`
}

func (cfg *hostConfig) Load(configPath string) error {
	err := cfg.loadYaml(configPath)
	cfg.loadEnvironment()
	return err
}

func (cfg *hostConfig) loadYaml(configPath string) error {
	logrus.Info("Loading host config")

	// Open the config file
	if _, err := os.Stat(configPath); err == nil {
		file, err := os.Open(configPath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Read the config file
		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(cfg)
		if err != nil {
			return err
		}
	} else if len(configPath) > 0 {
		logrus.Warnf("Failed to load config %s: %v", configPath, err)
	}

	return nil
}

func (cfg *hostConfig) loadEnvironment() {
	http_bind, ok := os.LookupEnv("HTTP_BIND")
	if ok && len(http_bind) > 0 {
		logrus.Info("Override http bind from environment")
		cfg.Http.Bind = http_bind
	}
	http_port, ok := os.LookupEnv("HTTP_PORT")
	if ok && len(http_port) > 0 {
		logrus.Info("Override http port from environment")
		cfg.Http.Port = http_port
	}
	grpc_bind, ok := os.LookupEnv("GRPC_BIND")
	if ok && len(grpc_bind) > 0 {
		logrus.Info("Override grpc bind from environment")
		cfg.Grpc.Bind = grpc_bind
	}
	grpc_port, ok := os.LookupEnv("GRPC_PORT")
	if ok && len(grpc_port) > 0 {
		logrus.Info("Override grpc port from environment")
		cfg.Grpc.Port = grpc_port
	}
}

func (cfg *hostConfig) GetHttpConfig() HttpConfig {
	return &cfg.Http
}

func (cfg *hostConfig) GetGrpcConfig() GrpcConfig {
	return &cfg.Grpc
}

func (cfg *httpConfig) GetBindAddress() string {
	return cfg.Bind + ":" + cfg.Port
}

func (cfg *grpcConfig) GetBindAddress() string {
	return cfg.Bind + ":" + cfg.Port
}
