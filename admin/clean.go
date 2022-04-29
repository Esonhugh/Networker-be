package main

import (
	"Network-be/Server/db"
	"Network-be/config"
	"Network-be/data/PO"
)

func main() {
	config.Init()
	db.InitDB()
	if err := db.DBService.MainDB.Unscoped().Where("verify = ?", false).Delete(&PO.Auth{}).Error; err != nil {
		println(err)
	} else {
		println("Cleaned.")
	}
}
