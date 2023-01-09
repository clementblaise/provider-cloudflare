/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	account "github.com/cdloh/provider-cloudflare/internal/controller/account/account"
	apitoken "github.com/cdloh/provider-cloudflare/internal/controller/account/apitoken"
	member "github.com/cdloh/provider-cloudflare/internal/controller/account/member"
	apishield "github.com/cdloh/provider-cloudflare/internal/controller/apishield/apishield"
	argo "github.com/cdloh/provider-cloudflare/internal/controller/argo/argo"
	tunnel "github.com/cdloh/provider-cloudflare/internal/controller/argo/tunnel"
	tunnelconfig "github.com/cdloh/provider-cloudflare/internal/controller/argo/tunnelconfig"
	tunnelroute "github.com/cdloh/provider-cloudflare/internal/controller/argo/tunnelroute"
	tunnelvirtualnetwork "github.com/cdloh/provider-cloudflare/internal/controller/argo/tunnelvirtualnetwork"
	authenticatedoriginspulls "github.com/cdloh/provider-cloudflare/internal/controller/authenticatedoriginpulls/authenticatedoriginspulls"
	certificate "github.com/cdloh/provider-cloudflare/internal/controller/authenticatedoriginpulls/certificate"
	ipprefix "github.com/cdloh/provider-cloudflare/internal/controller/byoip/ipprefix"
	pack "github.com/cdloh/provider-cloudflare/internal/controller/certificate/pack"
	pages "github.com/cdloh/provider-cloudflare/internal/controller/custom/pages"
	ssl "github.com/cdloh/provider-cloudflare/internal/controller/custom/ssl"
	fallbackorigin "github.com/cdloh/provider-cloudflare/internal/controller/customhostname/fallbackorigin"
	hostname "github.com/cdloh/provider-cloudflare/internal/controller/customhostname/hostname"
	profile "github.com/cdloh/provider-cloudflare/internal/controller/dlp/profile"
	record "github.com/cdloh/provider-cloudflare/internal/controller/dns/record"
	address "github.com/cdloh/provider-cloudflare/internal/controller/emailrouting/address"
	catchall "github.com/cdloh/provider-cloudflare/internal/controller/emailrouting/catchall"
	rule "github.com/cdloh/provider-cloudflare/internal/controller/emailrouting/rule"
	settings "github.com/cdloh/provider-cloudflare/internal/controller/emailrouting/settings"
	loadbalancer "github.com/cdloh/provider-cloudflare/internal/controller/loadbalancer/loadbalancer"
	monitor "github.com/cdloh/provider-cloudflare/internal/controller/loadbalancer/monitor"
	pool "github.com/cdloh/provider-cloudflare/internal/controller/loadbalancer/pool"
	job "github.com/cdloh/provider-cloudflare/internal/controller/logpush/job"
	ownershipchallenge "github.com/cdloh/provider-cloudflare/internal/controller/logpush/ownershipchallenge"
	rulepage "github.com/cdloh/provider-cloudflare/internal/controller/page/rule"
	domain "github.com/cdloh/provider-cloudflare/internal/controller/pages/domain"
	project "github.com/cdloh/provider-cloudflare/internal/controller/pages/project"
	providerconfig "github.com/cdloh/provider-cloudflare/internal/controller/providerconfig"
	accountteams "github.com/cdloh/provider-cloudflare/internal/controller/teams/account"
	list "github.com/cdloh/provider-cloudflare/internal/controller/teams/list"
	location "github.com/cdloh/provider-cloudflare/internal/controller/teams/location"
	proxyendpoint "github.com/cdloh/provider-cloudflare/internal/controller/teams/proxyendpoint"
	ruleteams "github.com/cdloh/provider-cloudflare/internal/controller/teams/rule"
	group "github.com/cdloh/provider-cloudflare/internal/controller/waf/group"
	override "github.com/cdloh/provider-cloudflare/internal/controller/waf/override"
	rulewaf "github.com/cdloh/provider-cloudflare/internal/controller/waf/rule"
	wafpackage "github.com/cdloh/provider-cloudflare/internal/controller/waf/wafpackage"
	event "github.com/cdloh/provider-cloudflare/internal/controller/waitingroom/event"
	room "github.com/cdloh/provider-cloudflare/internal/controller/waitingroom/room"
	rules "github.com/cdloh/provider-cloudflare/internal/controller/waitingroom/rules"
	devicepolicycertificates "github.com/cdloh/provider-cloudflare/internal/controller/warp/devicepolicycertificates"
	devicepostureintegration "github.com/cdloh/provider-cloudflare/internal/controller/warp/devicepostureintegration"
	deviceposturerule "github.com/cdloh/provider-cloudflare/internal/controller/warp/deviceposturerule"
	devicesettingspolicy "github.com/cdloh/provider-cloudflare/internal/controller/warp/devicesettingspolicy"
	splittunnel "github.com/cdloh/provider-cloudflare/internal/controller/warp/splittunnel"
	crontrigger "github.com/cdloh/provider-cloudflare/internal/controller/worker/crontrigger"
	kv "github.com/cdloh/provider-cloudflare/internal/controller/worker/kv"
	kvnamespace "github.com/cdloh/provider-cloudflare/internal/controller/worker/kvnamespace"
	route "github.com/cdloh/provider-cloudflare/internal/controller/worker/route"
	script "github.com/cdloh/provider-cloudflare/internal/controller/worker/script"
	dnssec "github.com/cdloh/provider-cloudflare/internal/controller/zone/dnssec"
	settingsoverride "github.com/cdloh/provider-cloudflare/internal/controller/zone/settingsoverride"
	zone "github.com/cdloh/provider-cloudflare/internal/controller/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		apitoken.Setup,
		member.Setup,
		apishield.Setup,
		argo.Setup,
		tunnel.Setup,
		tunnelconfig.Setup,
		tunnelroute.Setup,
		tunnelvirtualnetwork.Setup,
		authenticatedoriginspulls.Setup,
		certificate.Setup,
		ipprefix.Setup,
		pack.Setup,
		pages.Setup,
		ssl.Setup,
		fallbackorigin.Setup,
		hostname.Setup,
		profile.Setup,
		record.Setup,
		address.Setup,
		catchall.Setup,
		rule.Setup,
		settings.Setup,
		loadbalancer.Setup,
		monitor.Setup,
		pool.Setup,
		job.Setup,
		ownershipchallenge.Setup,
		rulepage.Setup,
		domain.Setup,
		project.Setup,
		providerconfig.Setup,
		accountteams.Setup,
		list.Setup,
		location.Setup,
		proxyendpoint.Setup,
		ruleteams.Setup,
		group.Setup,
		override.Setup,
		rulewaf.Setup,
		wafpackage.Setup,
		event.Setup,
		room.Setup,
		rules.Setup,
		devicepolicycertificates.Setup,
		devicepostureintegration.Setup,
		deviceposturerule.Setup,
		devicesettingspolicy.Setup,
		splittunnel.Setup,
		crontrigger.Setup,
		kv.Setup,
		kvnamespace.Setup,
		route.Setup,
		script.Setup,
		dnssec.Setup,
		settingsoverride.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
