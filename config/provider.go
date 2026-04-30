package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	parametersCluster "github.com/fulsiram/provider-upjet-warpgate/config/cluster/parameters"
	roleCluster "github.com/fulsiram/provider-upjet-warpgate/config/cluster/role"
	targetCluster "github.com/fulsiram/provider-upjet-warpgate/config/cluster/target"
	ticketCluster "github.com/fulsiram/provider-upjet-warpgate/config/cluster/ticket"
	userCluster "github.com/fulsiram/provider-upjet-warpgate/config/cluster/user"
	parametersNamespaced "github.com/fulsiram/provider-upjet-warpgate/config/namespaced/parameters"
	roleNamespaced "github.com/fulsiram/provider-upjet-warpgate/config/namespaced/role"
	targetNamespaced "github.com/fulsiram/provider-upjet-warpgate/config/namespaced/target"
	ticketNamespaced "github.com/fulsiram/provider-upjet-warpgate/config/namespaced/ticket"
	userNamespaced "github.com/fulsiram/provider-upjet-warpgate/config/namespaced/user"
)

const (
	resourcePrefix = "warpgate"
	modulePath     = "github.com/fulsiram/provider-upjet-warpgate"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("warpgate.salami.network"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		roleCluster.Configure,
		userCluster.Configure,
		targetCluster.Configure,
		ticketCluster.Configure,
		parametersCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("warpgate.m.salami.network"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		roleNamespaced.Configure,
		userNamespaced.Configure,
		targetNamespaced.Configure,
		ticketNamespaced.Configure,
		parametersNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
