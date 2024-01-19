package hosting

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpConfig_GetBindAddress(t *testing.T) {
	cfg := &HttpConfig{
		Bind: "127.0.0.1",
		Port: 8080,
	}

	expectedBindAddress := "127.0.0.1:8080"
	actualBindAddress := cfg.GetBindAddress()

	assert.Equal(t, expectedBindAddress, actualBindAddress, "Bind address should match")
}

func TestHttpConfig_LoadEnvironment(t *testing.T) {
	os.Setenv("HTTP_BIND", "127.0.0.1")
	os.Setenv("HTTP_PORT", "8080")

	cfg := &HttpConfig{}
	cfg.LoadEnvironment()

	expectedBind := "127.0.0.1"
	expectedPort := 8080

	assert.Equal(t, expectedBind, cfg.Bind, "Bind should match environment variable")
	assert.Equal(t, expectedPort, cfg.Port, "Port should match environment variable")

	os.Unsetenv("HTTP_BIND")
	os.Unsetenv("HTTP_PORT")
}

func TestHttpConfig_EnsureDefaults(t *testing.T) {
	cfg := &HttpConfig{}
	cfg.EnsureDefaults()

	expectedBind := DEFAULT_HTTP_BIND
	expectedPort := DEFAULT_HTTP_PORT

	assert.Equal(t, expectedBind, cfg.Bind, "Bind should match default value")
	assert.Equal(t, expectedPort, cfg.Port, "Port should match default value")
}
