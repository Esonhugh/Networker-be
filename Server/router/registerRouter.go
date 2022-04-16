package router

import "Network-be/Server/handlers"

func (g *GinList) RegisterRouter() {
	apis := g.MainWeb.Group("/api/v1")
	apis.GET("/ping", handlers.Ping)
	apis.GET("/config", handlers.GetConfig)
	peerinfo := apis.Group("/peerinfo")
	peerinfo.GET("/list", handlers.GetPeerList)
	peerinfo.GET("/:id", handlers.GetPeerInfo)
}
