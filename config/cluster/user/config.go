package user

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("warpgate_user", func(r *config.Resource) {
		r.ShortGroup = "user"
	})
	p.AddResourceConfigurator("warpgate_user_role", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.Kind = "UserRole"
		r.References["user_id"] = config.Reference{TerraformName: "warpgate_user"}
		r.References["role_id"] = config.Reference{TerraformName: "warpgate_role"}
	})
	p.AddResourceConfigurator("warpgate_password_credential", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.Kind = "PasswordCredential"
		r.References["user_id"] = config.Reference{TerraformName: "warpgate_user"}
	})
	p.AddResourceConfigurator("warpgate_public_key_credential", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.References["user_id"] = config.Reference{TerraformName: "warpgate_user"}
	})
	p.AddResourceConfigurator("warpgate_user_sso_credential", func(r *config.Resource) {
		r.ShortGroup = "user"
		r.References["user_id"] = config.Reference{TerraformName: "warpgate_user"}
	})
}
