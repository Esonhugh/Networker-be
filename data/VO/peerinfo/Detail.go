package peerinfo

type DetailPeer struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Asn          string `json:"asn"`
	PublicAccess string `json:"public_access"`
	WireguardKey string `json:"wireguard_key"`
	DN42_IPv4    string `json:"dn42_ipv4"`
	DN42_IPv6    string `json:"dn42_ipv6"`
}
