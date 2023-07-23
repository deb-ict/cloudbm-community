package core

import (
	"os"
	"strings"
)

type Environment string

func EnvironmentName() Environment {
	env, ok := os.LookupEnv("CBME_ENVIRONMENT")
	if !ok || env == "" {
		env = "Production"
	}
	return Environment(env)
}

func (env Environment) String() string {
	return string(env)
}

func (env Environment) ShortName() string {
	name := strings.ToLower(env.String())
	switch name {
	case "production":
		return "prd"
	case "staging":
		return "stg"
	default:
		return "dev"
	}
}
