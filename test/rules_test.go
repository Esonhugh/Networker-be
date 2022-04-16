package test

import (
	"Network-be/Server/db"
	"Network-be/config"
	"Network-be/data/VO/getConfig"
	"Network-be/utils"
	"testing"
)

func TestDns(t *testing.T) {
	goodDomain := "eson.ninja"
	badDomain := "eson.ninja.co>m"
	domains(goodDomain)
	domains(badDomain)
}

func domains(dns string) {
	print("dns is ", dns)
	println(" result:", utils.IsGoodDoamin(dns))
}

func TestConfig(t *testing.T) {
	config.Init()
	a := getConfig.AdminConfig()
	println(a)
}

func TestConfig2(t *testing.T) {
	config.Init()
	a := config.GlobalConfig.GetString("db.mainDB.DB")
	MainDB, err := db.CreateDB(struct {
		Addr           string
		User           string
		Pass           string
		DB             string
		ConnectTimeout uint
	}{
		Addr:           config.GlobalConfig.GetString("db.mainDB.Addr"),
		User:           config.GlobalConfig.GetString("db.mainDB.User"),
		Pass:           config.GlobalConfig.GetString("db.mainDB.Pass"),
		DB:             config.GlobalConfig.GetString("db.mainDB.DB"),
		ConnectTimeout: 10})
	print(a, MainDB, err)
}
