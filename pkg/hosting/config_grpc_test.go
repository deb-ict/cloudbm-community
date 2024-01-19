package hosting

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrpcConfig_GetBindAddress(t *testing.T) {
	cfg := &GrpcConfig{
		Bind: "127.0.0.1",
		Port: 8080,
	}

	expectedBindAddress := "127.0.0.1:8080"
	actualBindAddress := cfg.GetBindAddress()

	assert.Equal(t, expectedBindAddress, actualBindAddress, "Bind address should match")
}

func TestGrpcConfig_LoadEnvironment(t *testing.T) {
	os.Setenv("GRPC_BIND", "127.0.0.1")
	os.Setenv("GRPC_PORT", "8080")

	cfg := &GrpcConfig{}
	cfg.LoadEnvironment()

	expectedBind := "127.0.0.1"
	expectedPort := 8080

	assert.Equal(t, expectedBind, cfg.Bind, "Bind should match environment variable")
	assert.Equal(t, expectedPort, cfg.Port, "Port should match environment variable")

	os.Unsetenv("GRPC_BIND")
	os.Unsetenv("GRPC_PORT")
}

func TestGrpcConfig_EnsureDefaults(t *testing.T) {
	cfg := &GrpcConfig{}
	cfg.EnsureDefaults()

	expectedBind := DEFAULT_GRPC_BIND
	expectedPort := DEFAULT_GRPC_PORT

	assert.Equal(t, expectedBind, cfg.Bind, "Bind should match default value")
	assert.Equal(t, expectedPort, cfg.Port, "Port should match default value")
}
