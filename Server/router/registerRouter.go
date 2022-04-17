package router

import (
	"Network-be/Server/handlers"
	"Network-be/Server/router/jwt"
)

func (g *GinList) RegisterRouter() {
	apis := g.MainWeb.Group("/api/v1")
	apis.GET("/ping", handlers.Ping)
	apis.GET("/config", handlers.GetConfig)
	peerinfo := apis.Group("/peerinfo")
	peerinfo.GET("/list", jwt.JWTAuthMiddleware(), handlers.GetPeerList)
	peerinfo.GET("/:id", jwt.JWTAuthMiddleware(), handlers.GetPeerInfo)
	auth := apis.Group("/auth")
	auth.POST("/login", jwt.AuthHandler)
	auth.POST("/register", jwt.RegisterHandler)
}
