package main

import (
	"os"

	"github.com/deb-ict/cloudbm-community/pkg/hosting"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var (
	config *applicationConfig
)

type applicationConfig struct {
	http hosting.HttpConfig `yaml:"http"`
}

func LoadConfig(configPath string) (*applicationConfig, error) {
	cfg := &applicationConfig{
		http: hosting.HttpConfig{},
	}

	err := cfg.loadYaml(configPath)
	cfg.loadEnvironment()
	cfg.ensureDefaults()
	return cfg, err
}

func (cfg *applicationConfig) loadYaml(configPath string) error {
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
			logrus.Errorf("Failed to parse config %s: %v", configPath, err)
			return err
		}
	} else if len(configPath) > 0 {
		logrus.Warnf("Failed to load config %s: %v", configPath, err)
	}

	return nil
}

func (cfg *applicationConfig) loadEnvironment() {
	cfg.http.LoadEnvironment()
}

func (cfg *applicationConfig) ensureDefaults() {
	cfg.http.EnsureDefaults()
}

func (cfg *applicationConfig) GetHttpConfig() *hosting.HttpConfig {
	return &cfg.http
}
