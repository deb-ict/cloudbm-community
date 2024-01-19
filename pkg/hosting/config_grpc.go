package hosting

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	DEFAULT_GRPC_BIND string = "0.0.0.0"
	DEFAULT_GRPC_PORT int    = 81
)

type GrpcConfig struct {
	Bind string `yaml:"bind"`
	Port int    `yaml:"port"`
}

func (cfg *GrpcConfig) GetBindAddress() string {
	return fmt.Sprintf("%s:%d", cfg.Bind, cfg.Port)
}

func (cfg *GrpcConfig) LoadEnvironment() {
	grpc_bind, ok := os.LookupEnv("GRPC_BIND")
	if ok {
		logrus.Info("Override grpc bind from environment")
		cfg.Bind = grpc_bind
	}
	grpc_port, ok := os.LookupEnv("GRPC_PORT")
	if ok {
		port, err := strconv.Atoi(grpc_port)
		if err == nil {
			logrus.Info("Override grpc port from environment")
			cfg.Port = port
		}
	}
}

func (cfg *GrpcConfig) EnsureDefaults() {
	if cfg.Bind == "" {
		cfg.Bind = DEFAULT_GRPC_BIND
	}
	if cfg.Port <= 0 || cfg.Port > 65535 {
		cfg.Port = DEFAULT_GRPC_PORT
	}
}
