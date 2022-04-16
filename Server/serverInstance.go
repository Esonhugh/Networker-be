package Server

import (
	"Network-be/Server/db"
	"Network-be/Server/router"
	"Network-be/data/PO"
)

var DB *db.DBList
var Gin *router.GinList

func Run() {
	DB = db.InitDB()
	Gin = router.InitGin()

	DB.MainDB.AutoMigrate(&PO.Auth{}, &PO.Config{})

	Gin.RegisterRouter()
	Gin.MainWeb.Run(":8080")
}
