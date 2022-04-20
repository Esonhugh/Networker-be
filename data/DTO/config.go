package DTO

type Config struct {
	ID           int64
	Username     string
	Asn          string
	PublicAccess string
	WireguardKey string
	DN42_IPv4    string
	DN42_IPv6    string
}

func (c *Config) ToVO() {

}

/*
func (c *Config) ToDetailPeer() *peerinfo.DetailPeer {
	return &peerinfo.DetailPeer{
		ID:           c.ID,
		Username:     c.Username,
		Asn:          c.Asn,
		PublicAccess: c.PublicAccess,
		WireguardKey: c.WireguardKey,
		DN42_IPv4:    c.DN42_IPv4,
		DN42_IPv6:    c.DN42_IPv6,
	}
}

func (c *Config) ToSimplePeer() *peerinfo.SimplePeer {
	return &peerinfo.SimplePeer{
		ID:       c.ID,
		Username: c.Username,
		Asn:      c.Asn,
	}
}
*/
