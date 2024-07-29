package acmedns
/*
	THIS MODULE IS GENERATED AUTOMATICALLY
	DO NOT EDIT THIS FILE
*/
import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/providers/dns/alidns"
	"github.com/go-acme/lego/v4/providers/dns/allinkl"
	"github.com/go-acme/lego/v4/providers/dns/arvancloud"
	"github.com/go-acme/lego/v4/providers/dns/auroradns"
	"github.com/go-acme/lego/v4/providers/dns/autodns"
	"github.com/go-acme/lego/v4/providers/dns/azure"
	"github.com/go-acme/lego/v4/providers/dns/azuredns"
	"github.com/go-acme/lego/v4/providers/dns/bindman"
	"github.com/go-acme/lego/v4/providers/dns/bluecat"
	"github.com/go-acme/lego/v4/providers/dns/brandit"
	"github.com/go-acme/lego/v4/providers/dns/bunny"
	"github.com/go-acme/lego/v4/providers/dns/checkdomain"
	"github.com/go-acme/lego/v4/providers/dns/civo"
	"github.com/go-acme/lego/v4/providers/dns/clouddns"
	"github.com/go-acme/lego/v4/providers/dns/cloudflare"
	"github.com/go-acme/lego/v4/providers/dns/cloudns"
	"github.com/go-acme/lego/v4/providers/dns/cloudru"
	"github.com/go-acme/lego/v4/providers/dns/cloudxns"
	"github.com/go-acme/lego/v4/providers/dns/conoha"
	"github.com/go-acme/lego/v4/providers/dns/constellix"
	"github.com/go-acme/lego/v4/providers/dns/cpanel"
	"github.com/go-acme/lego/v4/providers/dns/derak"
	"github.com/go-acme/lego/v4/providers/dns/desec"
	"github.com/go-acme/lego/v4/providers/dns/digitalocean"
	"github.com/go-acme/lego/v4/providers/dns/dnshomede"
	"github.com/go-acme/lego/v4/providers/dns/dnsimple"
	"github.com/go-acme/lego/v4/providers/dns/dnsmadeeasy"
	"github.com/go-acme/lego/v4/providers/dns/dnspod"
	"github.com/go-acme/lego/v4/providers/dns/dode"
	"github.com/go-acme/lego/v4/providers/dns/domeneshop"
	"github.com/go-acme/lego/v4/providers/dns/dreamhost"
	"github.com/go-acme/lego/v4/providers/dns/duckdns"
	"github.com/go-acme/lego/v4/providers/dns/dyn"
	"github.com/go-acme/lego/v4/providers/dns/dynu"
	"github.com/go-acme/lego/v4/providers/dns/easydns"
	"github.com/go-acme/lego/v4/providers/dns/efficientip"
	"github.com/go-acme/lego/v4/providers/dns/epik"
	"github.com/go-acme/lego/v4/providers/dns/exoscale"
	"github.com/go-acme/lego/v4/providers/dns/freemyip"
	"github.com/go-acme/lego/v4/providers/dns/gandi"
	"github.com/go-acme/lego/v4/providers/dns/gandiv5"
	"github.com/go-acme/lego/v4/providers/dns/gcore"
	"github.com/go-acme/lego/v4/providers/dns/glesys"
	"github.com/go-acme/lego/v4/providers/dns/godaddy"
	"github.com/go-acme/lego/v4/providers/dns/googledomains"
	"github.com/go-acme/lego/v4/providers/dns/hetzner"
	"github.com/go-acme/lego/v4/providers/dns/hostingde"
	"github.com/go-acme/lego/v4/providers/dns/hosttech"
	"github.com/go-acme/lego/v4/providers/dns/httpnet"
	"github.com/go-acme/lego/v4/providers/dns/hyperone"
	"github.com/go-acme/lego/v4/providers/dns/ibmcloud"
	"github.com/go-acme/lego/v4/providers/dns/iij"
	"github.com/go-acme/lego/v4/providers/dns/iijdpf"
	"github.com/go-acme/lego/v4/providers/dns/infoblox"
	"github.com/go-acme/lego/v4/providers/dns/infomaniak"
	"github.com/go-acme/lego/v4/providers/dns/internetbs"
	"github.com/go-acme/lego/v4/providers/dns/inwx"
	"github.com/go-acme/lego/v4/providers/dns/ionos"
	"github.com/go-acme/lego/v4/providers/dns/ipv64"
	"github.com/go-acme/lego/v4/providers/dns/iwantmyname"
	"github.com/go-acme/lego/v4/providers/dns/joker"
	"github.com/go-acme/lego/v4/providers/dns/liara"
	"github.com/go-acme/lego/v4/providers/dns/lightsail"
	"github.com/go-acme/lego/v4/providers/dns/linode"
	"github.com/go-acme/lego/v4/providers/dns/liquidweb"
	"github.com/go-acme/lego/v4/providers/dns/loopia"
	"github.com/go-acme/lego/v4/providers/dns/luadns"
	"github.com/go-acme/lego/v4/providers/dns/mailinabox"
	"github.com/go-acme/lego/v4/providers/dns/metaname"
	"github.com/go-acme/lego/v4/providers/dns/mydnsjp"
	"github.com/go-acme/lego/v4/providers/dns/namecheap"
	"github.com/go-acme/lego/v4/providers/dns/namedotcom"
	"github.com/go-acme/lego/v4/providers/dns/namesilo"
	"github.com/go-acme/lego/v4/providers/dns/nearlyfreespeech"
	"github.com/go-acme/lego/v4/providers/dns/netcup"
	"github.com/go-acme/lego/v4/providers/dns/netlify"
	"github.com/go-acme/lego/v4/providers/dns/nicmanager"
	"github.com/go-acme/lego/v4/providers/dns/nifcloud"
	"github.com/go-acme/lego/v4/providers/dns/njalla"
	"github.com/go-acme/lego/v4/providers/dns/nodion"
	"github.com/go-acme/lego/v4/providers/dns/ns1"
	"github.com/go-acme/lego/v4/providers/dns/otc"
	"github.com/go-acme/lego/v4/providers/dns/ovh"
	"github.com/go-acme/lego/v4/providers/dns/pdns"
	"github.com/go-acme/lego/v4/providers/dns/plesk"
	"github.com/go-acme/lego/v4/providers/dns/porkbun"
	"github.com/go-acme/lego/v4/providers/dns/rackspace"
	"github.com/go-acme/lego/v4/providers/dns/rcodezero"
	"github.com/go-acme/lego/v4/providers/dns/regru"
	"github.com/go-acme/lego/v4/providers/dns/rfc2136"
	"github.com/go-acme/lego/v4/providers/dns/rimuhosting"
	"github.com/go-acme/lego/v4/providers/dns/route53"
	"github.com/go-acme/lego/v4/providers/dns/safedns"
	"github.com/go-acme/lego/v4/providers/dns/sakuracloud"
	"github.com/go-acme/lego/v4/providers/dns/scaleway"
	"github.com/go-acme/lego/v4/providers/dns/selectel"
	"github.com/go-acme/lego/v4/providers/dns/servercow"
	"github.com/go-acme/lego/v4/providers/dns/shellrent"
	"github.com/go-acme/lego/v4/providers/dns/simply"
	"github.com/go-acme/lego/v4/providers/dns/sonic"
	"github.com/go-acme/lego/v4/providers/dns/stackpath"
	"github.com/go-acme/lego/v4/providers/dns/tencentcloud"
	"github.com/go-acme/lego/v4/providers/dns/transip"
	"github.com/go-acme/lego/v4/providers/dns/ultradns"
	"github.com/go-acme/lego/v4/providers/dns/variomedia"
	"github.com/go-acme/lego/v4/providers/dns/vegadns"
	"github.com/go-acme/lego/v4/providers/dns/vercel"
	"github.com/go-acme/lego/v4/providers/dns/versio"
	"github.com/go-acme/lego/v4/providers/dns/vinyldns"
	"github.com/go-acme/lego/v4/providers/dns/vkcloud"
	"github.com/go-acme/lego/v4/providers/dns/vscale"
	"github.com/go-acme/lego/v4/providers/dns/vultr"
	"github.com/go-acme/lego/v4/providers/dns/webnames"
	"github.com/go-acme/lego/v4/providers/dns/websupport"
	"github.com/go-acme/lego/v4/providers/dns/wedos"
	"github.com/go-acme/lego/v4/providers/dns/yandex"
	"github.com/go-acme/lego/v4/providers/dns/yandex360"
	"github.com/go-acme/lego/v4/providers/dns/yandexcloud"
	"github.com/go-acme/lego/v4/providers/dns/zoneee"
	"github.com/go-acme/lego/v4/providers/dns/zonomi"

)

//name is the DNS provider name, e.g. cloudflare or gandi
//JSON (js) must be in key-value string that match ConfigableFields Title in providers.json, e.g. {"Username":"far","Password":"boo"}
func GetDNSProviderByJsonConfig(name string, js string)(challenge.Provider, error){
	switch name {

	case "alidns":
		cfg := alidns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return alidns.NewDNSProviderConfig(cfg)
	case "allinkl":
		cfg := allinkl.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return allinkl.NewDNSProviderConfig(cfg)
	case "arvancloud":
		cfg := arvancloud.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return arvancloud.NewDNSProviderConfig(cfg)
	case "auroradns":
		cfg := auroradns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return auroradns.NewDNSProviderConfig(cfg)
	case "autodns":
		cfg := autodns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return autodns.NewDNSProviderConfig(cfg)
	case "azure":
		cfg := azure.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return azure.NewDNSProviderConfig(cfg)
	case "azuredns":
		cfg := azuredns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return azuredns.NewDNSProviderConfig(cfg)
	case "bindman":
		cfg := bindman.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return bindman.NewDNSProviderConfig(cfg)
	case "bluecat":
		cfg := bluecat.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return bluecat.NewDNSProviderConfig(cfg)
	case "brandit":
		cfg := brandit.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return brandit.NewDNSProviderConfig(cfg)
	case "bunny":
		cfg := bunny.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return bunny.NewDNSProviderConfig(cfg)
	case "checkdomain":
		cfg := checkdomain.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return checkdomain.NewDNSProviderConfig(cfg)
	case "civo":
		cfg := civo.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return civo.NewDNSProviderConfig(cfg)
	case "clouddns":
		cfg := clouddns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return clouddns.NewDNSProviderConfig(cfg)
	case "cloudflare":
		cfg := cloudflare.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return cloudflare.NewDNSProviderConfig(cfg)
	case "cloudns":
		cfg := cloudns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return cloudns.NewDNSProviderConfig(cfg)
	case "cloudru":
		cfg := cloudru.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return cloudru.NewDNSProviderConfig(cfg)
	case "cloudxns":
		cfg := cloudxns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return cloudxns.NewDNSProviderConfig(cfg)
	case "conoha":
		cfg := conoha.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return conoha.NewDNSProviderConfig(cfg)
	case "constellix":
		cfg := constellix.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return constellix.NewDNSProviderConfig(cfg)
	case "cpanel":
		cfg := cpanel.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return cpanel.NewDNSProviderConfig(cfg)
	case "derak":
		cfg := derak.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return derak.NewDNSProviderConfig(cfg)
	case "desec":
		cfg := desec.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return desec.NewDNSProviderConfig(cfg)
	case "digitalocean":
		cfg := digitalocean.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return digitalocean.NewDNSProviderConfig(cfg)
	case "dnshomede":
		cfg := dnshomede.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return dnshomede.NewDNSProviderConfig(cfg)
	case "dnsimple":
		cfg := dnsimple.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return dnsimple.NewDNSProviderConfig(cfg)
	case "dnsmadeeasy":
		cfg := dnsmadeeasy.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return dnsmadeeasy.NewDNSProviderConfig(cfg)
	case "dnspod":
		cfg := dnspod.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return dnspod.NewDNSProviderConfig(cfg)
	case "dode":
		cfg := dode.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return dode.NewDNSProviderConfig(cfg)
	case "domeneshop":
		cfg := domeneshop.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return domeneshop.NewDNSProviderConfig(cfg)
	case "dreamhost":
		cfg := dreamhost.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return dreamhost.NewDNSProviderConfig(cfg)
	case "duckdns":
		cfg := duckdns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return duckdns.NewDNSProviderConfig(cfg)
	case "dyn":
		cfg := dyn.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return dyn.NewDNSProviderConfig(cfg)
	case "dynu":
		cfg := dynu.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return dynu.NewDNSProviderConfig(cfg)
	case "easydns":
		cfg := easydns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return easydns.NewDNSProviderConfig(cfg)
	case "efficientip":
		cfg := efficientip.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return efficientip.NewDNSProviderConfig(cfg)
	case "epik":
		cfg := epik.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return epik.NewDNSProviderConfig(cfg)
	case "exoscale":
		cfg := exoscale.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return exoscale.NewDNSProviderConfig(cfg)
	case "freemyip":
		cfg := freemyip.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return freemyip.NewDNSProviderConfig(cfg)
	case "gandi":
		cfg := gandi.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return gandi.NewDNSProviderConfig(cfg)
	case "gandiv5":
		cfg := gandiv5.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return gandiv5.NewDNSProviderConfig(cfg)
	case "gcore":
		cfg := gcore.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return gcore.NewDNSProviderConfig(cfg)
	case "glesys":
		cfg := glesys.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return glesys.NewDNSProviderConfig(cfg)
	case "godaddy":
		cfg := godaddy.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return godaddy.NewDNSProviderConfig(cfg)
	case "googledomains":
		cfg := googledomains.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return googledomains.NewDNSProviderConfig(cfg)
	case "hetzner":
		cfg := hetzner.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return hetzner.NewDNSProviderConfig(cfg)
	case "hostingde":
		cfg := hostingde.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return hostingde.NewDNSProviderConfig(cfg)
	case "hosttech":
		cfg := hosttech.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return hosttech.NewDNSProviderConfig(cfg)
	case "httpnet":
		cfg := httpnet.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return httpnet.NewDNSProviderConfig(cfg)
	case "hyperone":
		cfg := hyperone.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return hyperone.NewDNSProviderConfig(cfg)
	case "ibmcloud":
		cfg := ibmcloud.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return ibmcloud.NewDNSProviderConfig(cfg)
	case "iij":
		cfg := iij.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return iij.NewDNSProviderConfig(cfg)
	case "iijdpf":
		cfg := iijdpf.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return iijdpf.NewDNSProviderConfig(cfg)
	case "infoblox":
		cfg := infoblox.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return infoblox.NewDNSProviderConfig(cfg)
	case "infomaniak":
		cfg := infomaniak.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return infomaniak.NewDNSProviderConfig(cfg)
	case "internetbs":
		cfg := internetbs.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return internetbs.NewDNSProviderConfig(cfg)
	case "inwx":
		cfg := inwx.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return inwx.NewDNSProviderConfig(cfg)
	case "ionos":
		cfg := ionos.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return ionos.NewDNSProviderConfig(cfg)
	case "ipv64":
		cfg := ipv64.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return ipv64.NewDNSProviderConfig(cfg)
	case "iwantmyname":
		cfg := iwantmyname.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return iwantmyname.NewDNSProviderConfig(cfg)
	case "joker":
		cfg := joker.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return joker.NewDNSProviderConfig(cfg)
	case "liara":
		cfg := liara.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return liara.NewDNSProviderConfig(cfg)
	case "lightsail":
		cfg := lightsail.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return lightsail.NewDNSProviderConfig(cfg)
	case "linode":
		cfg := linode.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return linode.NewDNSProviderConfig(cfg)
	case "liquidweb":
		cfg := liquidweb.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return liquidweb.NewDNSProviderConfig(cfg)
	case "loopia":
		cfg := loopia.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return loopia.NewDNSProviderConfig(cfg)
	case "luadns":
		cfg := luadns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return luadns.NewDNSProviderConfig(cfg)
	case "mailinabox":
		cfg := mailinabox.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return mailinabox.NewDNSProviderConfig(cfg)
	case "metaname":
		cfg := metaname.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return metaname.NewDNSProviderConfig(cfg)
	case "mydnsjp":
		cfg := mydnsjp.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return mydnsjp.NewDNSProviderConfig(cfg)
	case "namecheap":
		cfg := namecheap.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return namecheap.NewDNSProviderConfig(cfg)
	case "namedotcom":
		cfg := namedotcom.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return namedotcom.NewDNSProviderConfig(cfg)
	case "namesilo":
		cfg := namesilo.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return namesilo.NewDNSProviderConfig(cfg)
	case "nearlyfreespeech":
		cfg := nearlyfreespeech.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return nearlyfreespeech.NewDNSProviderConfig(cfg)
		case "netcup":
			cfg := netcup.NewDefaultConfig()
			err := json.Unmarshal([]byte(js), &cfg)
			if err != nil {
				return nil, err
			}
			cfg.PropagationTimeout = 20*time.Minute
			return netcup.NewDNSProviderConfig(cfg)
	case "netlify":
		cfg := netlify.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return netlify.NewDNSProviderConfig(cfg)
	case "nicmanager":
		cfg := nicmanager.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return nicmanager.NewDNSProviderConfig(cfg)
	case "nifcloud":
		cfg := nifcloud.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return nifcloud.NewDNSProviderConfig(cfg)
	case "njalla":
		cfg := njalla.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return njalla.NewDNSProviderConfig(cfg)
	case "nodion":
		cfg := nodion.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return nodion.NewDNSProviderConfig(cfg)
	case "ns1":
		cfg := ns1.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return ns1.NewDNSProviderConfig(cfg)
	case "otc":
		cfg := otc.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return otc.NewDNSProviderConfig(cfg)
	case "ovh":
		cfg := ovh.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return ovh.NewDNSProviderConfig(cfg)
	case "pdns":
		cfg := pdns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return pdns.NewDNSProviderConfig(cfg)
	case "plesk":
		cfg := plesk.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return plesk.NewDNSProviderConfig(cfg)
	case "porkbun":
		cfg := porkbun.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return porkbun.NewDNSProviderConfig(cfg)
	case "rackspace":
		cfg := rackspace.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return rackspace.NewDNSProviderConfig(cfg)
	case "rcodezero":
		cfg := rcodezero.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return rcodezero.NewDNSProviderConfig(cfg)
	case "regru":
		cfg := regru.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return regru.NewDNSProviderConfig(cfg)
	case "rfc2136":
		cfg := rfc2136.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return rfc2136.NewDNSProviderConfig(cfg)
	case "rimuhosting":
		cfg := rimuhosting.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return rimuhosting.NewDNSProviderConfig(cfg)
	case "route53":
		cfg := route53.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return route53.NewDNSProviderConfig(cfg)
	case "safedns":
		cfg := safedns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return safedns.NewDNSProviderConfig(cfg)
	case "sakuracloud":
		cfg := sakuracloud.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return sakuracloud.NewDNSProviderConfig(cfg)
	case "scaleway":
		cfg := scaleway.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return scaleway.NewDNSProviderConfig(cfg)
	case "selectel":
		cfg := selectel.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return selectel.NewDNSProviderConfig(cfg)
	case "servercow":
		cfg := servercow.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return servercow.NewDNSProviderConfig(cfg)
	case "shellrent":
		cfg := shellrent.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return shellrent.NewDNSProviderConfig(cfg)
	case "simply":
		cfg := simply.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return simply.NewDNSProviderConfig(cfg)
	case "sonic":
		cfg := sonic.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return sonic.NewDNSProviderConfig(cfg)
	case "stackpath":
		cfg := stackpath.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return stackpath.NewDNSProviderConfig(cfg)
	case "tencentcloud":
		cfg := tencentcloud.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return tencentcloud.NewDNSProviderConfig(cfg)
	case "transip":
		cfg := transip.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return transip.NewDNSProviderConfig(cfg)
	case "ultradns":
		cfg := ultradns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return ultradns.NewDNSProviderConfig(cfg)
	case "variomedia":
		cfg := variomedia.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return variomedia.NewDNSProviderConfig(cfg)
	case "vegadns":
		cfg := vegadns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return vegadns.NewDNSProviderConfig(cfg)
	case "vercel":
		cfg := vercel.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return vercel.NewDNSProviderConfig(cfg)
	case "versio":
		cfg := versio.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return versio.NewDNSProviderConfig(cfg)
	case "vinyldns":
		cfg := vinyldns.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return vinyldns.NewDNSProviderConfig(cfg)
	case "vkcloud":
		cfg := vkcloud.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return vkcloud.NewDNSProviderConfig(cfg)
	case "vscale":
		cfg := vscale.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return vscale.NewDNSProviderConfig(cfg)
	case "vultr":
		cfg := vultr.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return vultr.NewDNSProviderConfig(cfg)
	case "webnames":
		cfg := webnames.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return webnames.NewDNSProviderConfig(cfg)
	case "websupport":
		cfg := websupport.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return websupport.NewDNSProviderConfig(cfg)
	case "wedos":
		cfg := wedos.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return wedos.NewDNSProviderConfig(cfg)
	case "yandex":
		cfg := yandex.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return yandex.NewDNSProviderConfig(cfg)
	case "yandex360":
		cfg := yandex360.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return yandex360.NewDNSProviderConfig(cfg)
	case "yandexcloud":
		cfg := yandexcloud.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return yandexcloud.NewDNSProviderConfig(cfg)
	case "zoneee":
		cfg := zoneee.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return zoneee.NewDNSProviderConfig(cfg)
	case "zonomi":
		cfg := zonomi.NewDefaultConfig()
		err := json.Unmarshal([]byte(js), &cfg)
		if err != nil {
			return nil, err
		}
		cfg.PropagationTimeout = 5*time.Minute
		return zonomi.NewDNSProviderConfig(cfg)
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}
