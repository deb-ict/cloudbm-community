package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/deb-ict/cloudbm-community/pkg/hosting"
	"gopkg.in/yaml.v3"

	auth_svc "github.com/deb-ict/cloudbm-community/pkg/module/auth/service"
	contact_svc "github.com/deb-ict/cloudbm-community/pkg/module/contact/service"
	gallery_svc "github.com/deb-ict/cloudbm-community/pkg/module/gallery/service"
	product_svc "github.com/deb-ict/cloudbm-community/pkg/module/product/service"
	session_svc "github.com/deb-ict/cloudbm-community/pkg/module/session/service"
)

type config struct {
	Http           hosting.HttpConfig         `yaml:"http"`
	AuthService    auth_svc.ServiceOptions    `yaml:"auth_service"`
	ContactService contact_svc.ServiceOptions `yaml:"contact_service"`
	GalleryService gallery_svc.ServiceOptions `yaml:"gallery_service"`
	ProductService product_svc.ServiceOptions `yaml:"product_service"`
	SessionService session_svc.ServiceOptions `yaml:"session_service"`
}

func LoadConfig(configPath string) (*config, error) {
	cfg := &config{
		Http:           hosting.HttpConfig{},
		AuthService:    auth_svc.ServiceOptions{},
		ContactService: contact_svc.ServiceOptions{},
		GalleryService: gallery_svc.ServiceOptions{},
		ProductService: product_svc.ServiceOptions{},
		SessionService: session_svc.ServiceOptions{},
	}
	err := cfg.loadYaml(configPath)
	cfg.loadEnvironment()
	cfg.ensureDefaults()
	return cfg, err
}

func (cfg *config) loadYaml(configPath string) error {
	logger := slog.With(slog.String("configPath", configPath))
	logger.InfoContext(context.Background(), "Loading config")

	// Open the config file
	if _, err := os.Stat(configPath); err == nil {
		file, err := os.Open(configPath)
		if err != nil {
			logger.ErrorContext(context.Background(), "Failed to open config",
				slog.Any("error", err),
			)
			return err
		}
		defer file.Close()

		// Read the config file
		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(cfg)
		if err != nil {
			logger.ErrorContext(context.Background(), "Failed to parse config",
				slog.Any("error", err),
			)
			return err
		}
	} else if len(configPath) > 0 {
		logger.WarnContext(context.Background(), "Failed to load config",
			slog.Any("error", err),
		)
	}

	return nil
}

func (cfg *config) loadEnvironment() {
	cfg.Http.LoadEnvironment()
}

func (cfg *config) ensureDefaults() {
	cfg.Http.EnsureDefaults()
	cfg.AuthService.EnsureDefaults()
	cfg.ContactService.EnsureDefaults()
	cfg.GalleryService.EnsureDefaults()
	cfg.ProductService.EnsureDefaults()
	cfg.SessionService.EnsureDefaults()
}
