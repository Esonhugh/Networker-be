package PO

import (
	"Network-be/data/DTO"
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	ID           int64 `gorm:"column:id;primary_key;increase"`
	Username     string
	Asn          string
	PublicAccess string
	WireguardKey string
	DN42_IPv4    string
	DN42_IPv6    string
}

func (*Config) TableName() string {
	return "config"
}

func (c *Config) ToDto() *DTO.Config {
	return &DTO.Config{
		ID:           c.ID,
		Username:     c.Username,
		Asn:          c.Asn,
		PublicAccess: c.PublicAccess,
		WireguardKey: c.WireguardKey,
		DN42_IPv4:    c.DN42_IPv4,
		DN42_IPv6:    c.DN42_IPv6,
	}
}
