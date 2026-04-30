package target

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("warpgate_target_group", func(r *config.Resource) {
		r.ShortGroup = "target"
	})
	p.AddResourceConfigurator("warpgate_target", func(r *config.Resource) {
		r.ShortGroup = "target"
		r.References["group_id"] = config.Reference{TerraformName: "warpgate_target_group"}
	})
	p.AddResourceConfigurator("warpgate_target_role", func(r *config.Resource) {
		r.ShortGroup = "target"
		r.Kind = "TargetRole"
		r.References["target_id"] = config.Reference{TerraformName: "warpgate_target"}
		r.References["role_id"] = config.Reference{TerraformName: "warpgate_role"}
	})
}
