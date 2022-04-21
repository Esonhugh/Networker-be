package peerinfo

import "Network-be/data/PO"

type UpdateInfo struct {
	Username     string `json:"username" binding:"required,min=1,max=20"`
	Asn          string `json:"asn"`
	PublicAccess string `json:"public_access" binding:"required,fqdn|ipv4|ipv6"`
	WireguradKey string `json:"wireguard_key" binding:"required,base64"`
	DN42_IPv4    string `json:"dn42_ipv4" binding:"required,ipv4,omitempty"`
	DN42_IPv6    string `json:"dn42_ipv6" binding:"required,ipv6"`
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
