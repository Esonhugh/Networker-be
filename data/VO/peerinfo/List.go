package peerinfo

type PeerList struct {
	Peers []SimplePeer
}

type SimplePeer struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Asn      string `json:"asn"`
}
