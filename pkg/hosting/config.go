package hosting

import (
	"context"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type HostConfig struct {
	http   HttpConfig   `yaml:"http"`
	grpc   GrpcConfig   `yaml:"grpc"`
	health HealthConfig `yaml:"health"`
}

func NewHostConfig() *HostConfig {
	return &HostConfig{
		http: HttpConfig{
			Bind: "0.0.0.0",
			Port: 80,
		},
		grpc: GrpcConfig{
			Bind: "0.0.0.0",
			Port: 81,
		},
		health: HealthConfig{
			Bind: "0.0.0.0",
			Port: 8080,
		},
	}
}

func (cfg *HostConfig) Load(configPath string) error {
	err := cfg.loadYaml(configPath)
	cfg.LoadEnvironment()
	cfg.EnsureDefaults()
	return err
}

func (cfg *HostConfig) loadYaml(configPath string) error {
	slog.InfoContext(context.Background(), "Loading host config")

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
			slog.ErrorContext(context.Background(), "Failed to parse config",
				slog.String("configPath", configPath),
				slog.Any("error", err),
			)
			return err
		}
	} else if len(configPath) > 0 {
		slog.WarnContext(context.Background(), "Failed to load config",
			slog.String("configPath", configPath),
			slog.Any("error", err),
		)
	}

	return nil
}

func (cfg *HostConfig) LoadEnvironment() {
	cfg.http.LoadEnvironment()
	cfg.grpc.LoadEnvironment()
	cfg.health.LoadEnvironment()
}

func (cfg *HostConfig) EnsureDefaults() {
	cfg.http.EnsureDefaults()
	cfg.grpc.EnsureDefaults()
	cfg.health.EnsureDefaults()
}

func (cfg *HostConfig) GetHttpConfig() *HttpConfig {
	return &cfg.http
}

func (cfg *HostConfig) GetGrpcConfig() *GrpcConfig {
	return &cfg.grpc
}

func (cfg *HostConfig) GetHealthConfig() *HealthConfig {
	return &cfg.health
}
