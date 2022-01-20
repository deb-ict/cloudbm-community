package webhost

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config interface {
	GetHttpBind() string
	GetHttpPort() string
	IsApiEnabled() bool
	Load(configPath string) error
}

type config struct {
	http HttpConfig `yaml:"http"`
}

type HttpConfig struct {
	Bind string `yaml:"bind"`
	Port string `yaml:"port"`
	Api  bool   `yaml:"api"`
}

func NewConfig() Config {
	return &config{
		http: HttpConfig{
			Bind: "0.0.0.0",
			Port: "80",
			Api:  false,
		},
	}
}

func (cfg *config) GetHttpBind() string {
	return cfg.http.Bind
}

func (cfg *config) GetHttpPort() string {
	return cfg.http.Port
}

func (cfg *config) IsApiEnabled() bool {
	return cfg.http.Api
}

func (cfg *config) Load(configPath string) error {
	err := cfg.loadYaml(configPath)
	cfg.LoadEnvironment()
	return err
}

func (cfg *config) loadYaml(configPath string) error {
	// Open the config file
	/*
		file, err := os.Open(configPath)
		if err != nil {
			return err
		}
		defer file.Close()
	*/

	yamlData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	// Read the config file
	//decoder := yaml.NewDecoder(file)
	/*
		err = decoder.Decode(cfg)
		if err != nil {
			return err
		}
	*/

	var config config
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *config) LoadEnvironment() {
	http_bind, ok := os.LookupEnv("HTTP_BIND")
	if ok && len(http_bind) > 0 {
		cfg.http.Bind = http_bind
	}
	http_port, ok := os.LookupEnv("HTTP_PORT")
	if ok && len(http_port) > 0 {
		cfg.http.Port = http_port
	}
}
