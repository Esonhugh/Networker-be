package main

import (
	"Network-be/Server/db"
	"Network-be/config"
	"Network-be/data/PO"
	"Network-be/data/VO/peerinfo"
	"encoding/json"
)

func main() {
	var peerList []peerinfo.DetailPeer
	config.Init()
	db.InitDB()
	db.DBService.MainDB.Model(&PO.Config{}).Find(&peerList)
	for _, v := range peerList {
		json, _ := json.Marshal(v)
		println(string(json))
	}
}
