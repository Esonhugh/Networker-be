package router

import (
	"Network-be/Server/handlers"
	"Network-be/Server/router/jwt"
)

func (g *GinList) RegisterRouter() {
	// Global middleware Allow Cors Policy.
	g.MainWeb.Use(handlers.SetCorsPolicy)
	// apis Main api
	apis := g.MainWeb.Group("/api/v1")
	{
		apis.GET("/ping", handlers.Ping)
		apis.GET("/config", handlers.GetConfig)

		// auth sub-api path
		auth := apis.Group("/auth")
		{
			auth.POST("/login", handlers.AuthHandler)
			// auth.Any("/logout", handlers.LogoutHandler)
			auth.POST("/register", handlers.RegisterHandler)
			auth.GET("/verify/:ticket", handlers.VerifyHandler)
		}

		// PeerInfo sub-api path
		PeerInfo := apis.Group("/peerinfo")
		PeerInfo.Use(jwt.JWTAuthMiddleware)
		{
			PeerInfo.GET("/list", handlers.GetPeerList)
			PeerInfo.GET("/me", handlers.GetMyInfo)
			PeerInfo.GET("/:id", handlers.GetPeerInfo)
			PeerInfo.POST("/me", handlers.UpdatePeerInfo)
		}
	}
}
