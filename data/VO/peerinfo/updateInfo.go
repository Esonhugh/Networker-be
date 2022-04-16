package peerinfo

import "Network-be/data/PO"

type UpdateInfo struct {
	Username     string `json:"username"`
	Asn          string `json:"asn"`
	PublicAccess string `json:"public_access"`
	WireguradKey string `json:"wireguard_key"`
	DN42_IPv4    string `json:"dn42_ipv4"`
	DN42_IPv6    string `json:"dn42_ipv6"`
}

func (u *UpdateInfo) ToPO() *PO.Config {
	return &PO.Config{
		Username:     u.Username,
		Asn:          u.Asn,
		PublicAccess: u.PublicAccess,
		WireguardKey: u.WireguradKey,
		DN42_IPv4:    u.DN42_IPv4,
		DN42_IPv6:    u.DN42_IPv6,
	}
}
