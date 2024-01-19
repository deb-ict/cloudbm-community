package hosting

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHostConfig(t *testing.T) {
	cfg := NewHostConfig()

	assert.NotNil(t, cfg)
	assert.NotNil(t, cfg.http)
	assert.NotNil(t, cfg.grpc)
	assert.NotNil(t, cfg.health)
}

func TestLoadEnvironment(t *testing.T) {
	os.Setenv("HTTP_BIND", "127.0.0.1")
	os.Setenv("HTTP_PORT", "8080")
	os.Setenv("GRPC_BIND", "127.0.0.1")
	os.Setenv("GRPC_PORT", "8081")
	os.Setenv("HEALTH_BIND", "127.0.0.1")
	os.Setenv("HEALTH_PORT", "8082")

	cfg := NewHostConfig()
	cfg.LoadEnvironment()

	expectedHttpBind := "127.0.0.1"
	expectedHttpPort := 8080
	expectedGrpcBind := "127.0.0.1"
	expectedGrpcPort := 8081
	expectedHealthBind := "127.0.0.1"
	expectedHealthPort := 8082

	assert.Equal(t, expectedHttpBind, cfg.GetHttpConfig().Bind, "Http bind should match environment variable")
	assert.Equal(t, expectedHttpPort, cfg.GetHttpConfig().Port, "Http port should match environment variable")
	assert.Equal(t, expectedGrpcBind, cfg.GetGrpcConfig().Bind, "Grpc bind should match environment variable")
	assert.Equal(t, expectedGrpcPort, cfg.GetGrpcConfig().Port, "Grpc port should match environment variable")
	assert.Equal(t, expectedHealthBind, cfg.GetHealthConfig().Bind, "Health bind should match environment variable")
	assert.Equal(t, expectedHealthPort, cfg.GetHealthConfig().Port, "Health port should match environment variable")

	os.Unsetenv("HTTP_BIND")
	os.Unsetenv("HTTP_PORT")
	os.Unsetenv("GRPC_BIND")
	os.Unsetenv("GRPC_PORT")
	os.Unsetenv("HEALTH_BIND")
	os.Unsetenv("HEALTH_PORT")
}

func TestEnsureDefaults(t *testing.T) {
	cfg := NewHostConfig()
	cfg.EnsureDefaults()

	expectedHttpBind := DEFAULT_HTTP_BIND
	expectedHttpPort := DEFAULT_HTTP_PORT
	expectedGrpcBind := DEFAULT_GRPC_BIND
	expectedGrpcPort := DEFAULT_GRPC_PORT
	expectedHealthBind := DEFAULT_HEALTH_BIND
	expectedHealthPort := DEFAULT_HEALTH_PORT

	assert.Equal(t, expectedHttpBind, cfg.GetHttpConfig().Bind, "Http bind should match default value")
	assert.Equal(t, expectedHttpPort, cfg.GetHttpConfig().Port, "Http port should match default value")
	assert.Equal(t, expectedGrpcBind, cfg.GetGrpcConfig().Bind, "Grpc bind should match default value")
	assert.Equal(t, expectedGrpcPort, cfg.GetGrpcConfig().Port, "Grpc port should match default value")
	assert.Equal(t, expectedHealthBind, cfg.GetHealthConfig().Bind, "Health bind should match default value")
	assert.Equal(t, expectedHealthPort, cfg.GetHealthConfig().Port, "Health port should match default value")
}

func TestGetHttpConfig(t *testing.T) {
	cfg := NewHostConfig()
	httpCfg := cfg.GetHttpConfig()

	assert.NotNil(t, httpCfg)
}

func TestGetGrpcConfig(t *testing.T) {
	cfg := NewHostConfig()
	grpcCfg := cfg.GetGrpcConfig()

	assert.NotNil(t, grpcCfg)
}

func TestGetHealthConfig(t *testing.T) {
	cfg := NewHostConfig()
	healthCfg := cfg.GetHealthConfig()

	assert.NotNil(t, healthCfg)
}
