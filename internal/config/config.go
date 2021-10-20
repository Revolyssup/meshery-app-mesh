package config

import (
	"path"
	"strings"

	"github.com/layer5io/meshery-adapter-library/common"
	"github.com/layer5io/meshery-adapter-library/config"
	configprovider "github.com/layer5io/meshery-adapter-library/config/provider"
	"github.com/layer5io/meshery-adapter-library/status"

	//	"github.com/layer5io/meshery-adapter-library"
	"github.com/layer5io/meshkit/utils"
	smp "github.com/layer5io/service-mesh-performance/spec"
)

const (
	LabelNamespace = "label-namespace"

	// OAM Metadata constants
	OAMAdapterNameMetadataKey       = "adapter.meshery.io/name"
	OAMComponentCategoryMetadataKey = "ui.meshery.io/category"

	// supported addons
	PrometheusAddon = "prometheus-addon"
	GrafanaAddon    = "grafana-addon"
	JaegerAddon     = "jaeger-addon"
	SpireAgent      = "spire-agent"
	SpireServer     = "spire-server"
	AppMeshGateway  = "appmesh-gateway"
)

var (
	AppMeshOperation = strings.ToLower(smp.ServiceMesh_APP_MESH.Enum().String())

	ServerVersion  = status.None
	ServerGitSHA   = status.None
	configRootPath = path.Join(utils.GetHome(), ".meshery")

	Config = configprovider.Options{
		ServerConfig:   ServerConfig,
		MeshSpec:       MeshSpec,
		ProviderConfig: ProviderConfig,
		Operations:     Operations,
	}

	// ServerConfig is the configuration for the gRPC server
	ServerConfig = map[string]string{
		"name":     smp.ServiceMesh_APP_MESH.Enum().String(),
		"type":     "adapter",
		"port":     "10005",
		"traceurl": status.None,
	}

	// MeshSpec is the spec for the service mesh associated with this adapter
	MeshSpec = map[string]string{
		"name":    smp.ServiceMesh_APP_MESH.Enum().String(),
		"status":  status.NotInstalled,
		"version": status.None,
	}

	// ProviderConfig is the config for the configuration provider
	ProviderConfig = map[string]string{
		configprovider.FilePath: configRootPath,
		configprovider.FileType: "yaml",
		configprovider.FileName: "app-mesh",
	}

	// KubeConfig - Controlling the kubeconfig lifecycle with viper
	KubeConfig = map[string]string{
		configprovider.FilePath: configRootPath,
		configprovider.FileType: "yaml",
		configprovider.FileName: "kubeconfig",
	}

	// Operations represents the set of valid operations that are available
	// to the adapter
	Operations = getOperations(common.Operations)
)

// New creates a new config instance
func New(provider string) (config.Handler, error) {
	// Config provider
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(Config)
	case configprovider.InMemKey:
		return configprovider.NewInMem(Config)
	}

	return nil, ErrEmptyConfig

}

func NewKubeconfigBuilder(provider string) (config.Handler, error) {
	opts := configprovider.Options{}
	opts.ProviderConfig = KubeConfig

	// Config provider
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(opts)
	case configprovider.InMemKey:
		return configprovider.NewInMem(opts)
	}
	return nil, ErrEmptyConfig
}

// RootPath returns the config root path for the adapter
func RootPath() string {
	return configRootPath
}
