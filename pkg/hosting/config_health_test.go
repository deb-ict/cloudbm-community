package hosting

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthConfig_GetBindAddress(t *testing.T) {
	cfg := &HealthConfig{
		Bind: "127.0.0.1",
		Port: 8080,
	}

	expectedBindAddress := "127.0.0.1:8080"
	actualBindAddress := cfg.GetBindAddress()

	assert.Equal(t, expectedBindAddress, actualBindAddress, "Bind address should match")
}

func TestHealthConfig_LoadEnvironment(t *testing.T) {
	os.Setenv("HEALTH_BIND", "127.0.0.1")
	os.Setenv("HEALTH_PORT", "8080")

	cfg := &HealthConfig{}
	cfg.LoadEnvironment()

	expectedBind := "127.0.0.1"
	expectedPort := 8080

	assert.Equal(t, expectedBind, cfg.Bind, "Bind should match environment variable")
	assert.Equal(t, expectedPort, cfg.Port, "Port should match environment variable")

	os.Unsetenv("HEALTH_BIND")
	os.Unsetenv("HEALTH_PORT")
}

func TestHealthConfig_EnsureDefaults(t *testing.T) {
	cfg := &HealthConfig{}
	cfg.EnsureDefaults()

	expectedBind := DEFAULT_HEALTH_BIND
	expectedPort := DEFAULT_HEALTH_PORT

	assert.Equal(t, expectedBind, cfg.Bind, "Bind should match default value")
	assert.Equal(t, expectedPort, cfg.Port, "Port should match default value")
}
