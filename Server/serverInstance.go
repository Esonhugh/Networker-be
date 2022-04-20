package Server

import (
	"Network-be/Server/db"
	"Network-be/Server/router"
	"Network-be/config"
	"Network-be/data/PO"
)

func Run() {

	DB := db.InitDB()
	Gin := router.InitGin()

	DB.MainDB.AutoMigrate(&PO.Auth{}, &PO.Config{})

	Gin.RegisterRouter()
	Gin.MainWeb.Run(":" + config.GlobalConfig.GetString("server.port"))
}
