package hosting

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	DEFAULT_HTTP_BIND string = "0.0.0.0"
	DEFAULT_HTTP_PORT int    = 80
)

type HttpConfig struct {
	Bind string `yaml:"bind"`
	Port int    `yaml:"port"`
}

func (cfg *HttpConfig) GetBaseUri() string {
	return fmt.Sprintf("http://%s", cfg.GetBindAddress())
}

func (cfg *HttpConfig) GetBindAddress() string {
	return fmt.Sprintf("%s:%d", cfg.Bind, cfg.Port)
}

func (cfg *HttpConfig) LoadEnvironment() {
	http_bind, ok := os.LookupEnv("HTTP_BIND")
	if ok {
		logrus.Info("Override http bind from environment")
		cfg.Bind = http_bind
	}
	http_port, ok := os.LookupEnv("HTTP_PORT")
	if ok {
		port, err := strconv.Atoi(http_port)
		if err == nil {
			logrus.Info("Override http port from environment")
			cfg.Port = port
		}
	}
}

func (cfg *HttpConfig) EnsureDefaults() {
	if cfg.Bind == "" {
		cfg.Bind = DEFAULT_HTTP_BIND
	}
	if cfg.Port <= 0 || cfg.Port > 65535 {
		cfg.Port = DEFAULT_HTTP_PORT
	}
}
