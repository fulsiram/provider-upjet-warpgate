package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"warpgate_role":                  config.IdentifierFromProvider,
	"warpgate_user":                  config.IdentifierFromProvider,
	"warpgate_target":                config.IdentifierFromProvider,
	"warpgate_target_group":          config.IdentifierFromProvider,
	"warpgate_ticket":                config.IdentifierFromProvider,
	"warpgate_parameters":            config.IdentifierFromProvider,
	"warpgate_user_role":             config.IdentifierFromProvider,
	"warpgate_target_role":           config.IdentifierFromProvider,
	"warpgate_password_credential":   config.IdentifierFromProvider,
	"warpgate_public_key_credential": config.IdentifierFromProvider,
	"warpgate_user_sso_credential":   config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
