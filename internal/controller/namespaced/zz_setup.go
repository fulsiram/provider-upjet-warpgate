// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	parameters "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/parameters/parameters"
	providerconfig "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/providerconfig"
	role "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/role/role"
	group "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/target/group"
	target "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/target/target"
	targetrole "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/target/targetrole"
	ticket "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/ticket/ticket"
	keycredential "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/user/keycredential"
	passwordcredential "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/user/passwordcredential"
	ssocredential "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/user/ssocredential"
	user "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/user/user"
	userrole "github.com/fulsiram/provider-upjet-warpgate/internal/controller/namespaced/user/userrole"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		parameters.Setup,
		providerconfig.Setup,
		role.Setup,
		group.Setup,
		target.Setup,
		targetrole.Setup,
		ticket.Setup,
		keycredential.Setup,
		passwordcredential.Setup,
		ssocredential.Setup,
		user.Setup,
		userrole.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		parameters.SetupGated,
		providerconfig.SetupGated,
		role.SetupGated,
		group.SetupGated,
		target.SetupGated,
		targetrole.SetupGated,
		ticket.SetupGated,
		keycredential.SetupGated,
		passwordcredential.SetupGated,
		ssocredential.SetupGated,
		user.SetupGated,
		userrole.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
