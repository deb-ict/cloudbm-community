package webhost

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Config interface {
	GetHttpBind() string
	GetHttpPort() string
	IsApiEnabled() bool
	Load(configPath string) error
}

type config struct {
	Http httpConfig `yaml:"http"`
}

type httpConfig struct {
	Bind string `yaml:"bind"`
	Port string `yaml:"port"`
	Api  bool   `yaml:"api"`
}

func NewConfig() Config {
	return &config{
		Http: httpConfig{
			Bind: "0.0.0.0",
			Port: "80",
			Api:  false,
		},
	}
}

func (cfg *config) GetHttpBind() string {
	return cfg.Http.Bind
}

func (cfg *config) GetHttpPort() string {
	return cfg.Http.Port
}

func (cfg *config) IsApiEnabled() bool {
	return cfg.Http.Api
}

func (cfg *config) Load(configPath string) error {
	err := cfg.loadYaml(configPath)
	cfg.LoadEnvironment()
	return err
}

func (cfg *config) loadYaml(configPath string) error {
	// Open the config file
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

	return nil
}

func (cfg *config) LoadEnvironment() {
	http_bind, ok := os.LookupEnv("HTTP_BIND")
	if ok && len(http_bind) > 0 {
		cfg.Http.Bind = http_bind
	}
	http_port, ok := os.LookupEnv("HTTP_PORT")
	if ok && len(http_port) > 0 {
		cfg.Http.Port = http_port
	}
	api_enabled, ok := os.LookupEnv("API_ENABLED")
	if ok {
		api_enable_value, err := strconv.ParseBool(api_enabled)
		if err == nil {
			cfg.Http.Api = api_enable_value
		}
	}
}
