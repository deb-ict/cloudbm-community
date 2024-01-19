package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironmentName(t *testing.T) {
	tests := []struct {
		environment string
		expected    Environment
	}{
		{"Production", Environment("Production")},
		{"Staging", Environment("Staging")},
		{"Development", Environment("Development")},
		{"", Environment("Production")},
	}

	for _, test := range tests {
		os.Setenv("CBME_ENVIRONMENT", test.environment)
		result := EnvironmentName()
		assert.Equal(t, test.expected, result, "Environment(%s) = %s, expected %s", test.environment, result, test.expected)
		os.Unsetenv("CBME_ENVIRONMENT")
	}
}

func TestEnvironment_String(t *testing.T) {
	env := Environment("production")
	expected := "production"
	result := env.String()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestEnvironment_ShortName(t *testing.T) {
	tests := []struct {
		environment string
		expected    string
	}{
		{"production", "prd"},
		{"staging", "stg"},
		{"development", "dev"},
		{"invalid", "prd"},
		{"", "prd"},
	}

	for _, test := range tests {
		env := Environment(test.environment)
		result := env.ShortName()
		assert.Equal(t, test.expected, result, "Environment(%s).ShortName() = %s, expected %s", test.environment, result, test.expected)
	}
}
