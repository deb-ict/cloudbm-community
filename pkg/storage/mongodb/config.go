package mongodb

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config interface {
	GetUri() string
	GetDatabase() string
	GetMaxPageSize() int
	Load(configPath string) error
}

type config struct {
	MongoDb mongodbConfig `yaml:"mongodb"`
}

type mongodbConfig struct {
	Uri         string `yaml:"uri"`
	Database    string `yaml:"database"`
	MaxPageSize int    `yaml:"maxPageSize"`
}

func NewConfig() Config {
	return &config{
		MongoDb: mongodbConfig{
			Uri:         "mongodb://localhost:27017",
			Database:    "cloudbm",
			MaxPageSize: 150,
		},
	}
}

func (cfg *config) GetUri() string {
	return cfg.MongoDb.Uri
}

func (cfg *config) GetDatabase() string {
	return cfg.MongoDb.Database
}

func (cfg *config) GetMaxPageSize() int {
	return cfg.MongoDb.MaxPageSize
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
	mongo_uri, ok := os.LookupEnv("MONGO_URI")
	if ok && len(mongo_uri) > 0 {
		cfg.MongoDb.Uri = mongo_uri
	}
	mongo_db, ok := os.LookupEnv("MONGO_DB")
	if ok && len(mongo_db) > 0 {
		cfg.MongoDb.Database = mongo_db
	}
}
