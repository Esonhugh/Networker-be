package getConfig

import (
	"Network-be/config"
	"Network-be/utils"
	"log"
)

type GetConfig struct {
	Username     string `json:"username"`
	Asn          string `json:"asn"`
	PublicAccess string `json:"public_access"`
	WireguardKey string `json:"wireguard_key"`
	DN42_IPv4    string `json:"dn42_ipv4"`
	DN42_IPv6    string `json:"dn42_ipv6"`
}

// ReadAdmin Config and Return
func AdminConfig() *GetConfig {
	conf := &GetConfig{
		Username:     config.GlobalConfig.GetString("admin.config.username"),
		Asn:          config.GlobalConfig.GetString("admin.config.asn"),
		PublicAccess: config.GlobalConfig.GetString("admin.config.public_access"),
		WireguardKey: config.GlobalConfig.GetString("admin.config.wireguard_key"),
		DN42_IPv4:    config.GlobalConfig.GetString("admin.config.dn42_ipv4"),
		DN42_IPv6:    config.GlobalConfig.GetString("admin.config.dn42_ipv6"),
	}
	if !utils.IsGoodIP(conf.DN42_IPv4) || !utils.IsGoodIP(conf.DN42_IPv6) {
		log.Panic("DN42_IPv4 or DN42_IPv6 is not a valid IP ", "\n",
			"DN42_IPv4 ", conf.DN42_IPv4, "\n",
			"DN42_IPv6 ", conf.DN42_IPv6, "\n")
		return nil
	}
	if !utils.IsGoodIP(conf.PublicAccess) && !utils.IsGoodDoamin(conf.PublicAccess) {
		log.Panic("PublicAccess is not a valid IP or Domain ", "\n",
			"PublicAccess ", conf.PublicAccess, "\n")
		return nil
	}
	return conf
}
