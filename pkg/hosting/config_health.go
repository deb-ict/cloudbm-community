package hosting

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	DEFAULT_HEALTH_BIND string = "0.0.0.0"
	DEFAULT_HEALTH_PORT int    = 8080
)

type HealthConfig struct {
	Bind string `yaml:"bind"`
	Port int    `yaml:"port"`
}

func (cfg *HealthConfig) GetBindAddress() string {
	return fmt.Sprintf("%s:%d", cfg.Bind, cfg.Port)
}

func (cfg *HealthConfig) LoadEnvironment() {
	http_bind, ok := os.LookupEnv("HEALTH_BIND")
	if ok {
		logrus.Info("Override health bind from environment")
		cfg.Bind = http_bind
	}
	http_port, ok := os.LookupEnv("HEALTH_PORT")
	if ok {
		port, err := strconv.Atoi(http_port)
		if err == nil {
			logrus.Info("Override health port from environment")
			cfg.Port = port
		}
	}
}

func (cfg *HealthConfig) EnsureDefaults() {
	if cfg.Bind == "" {
		cfg.Bind = DEFAULT_HEALTH_BIND
	}
	if cfg.Port <= 0 || cfg.Port > 65535 {
		cfg.Port = DEFAULT_HEALTH_PORT
	}
}
