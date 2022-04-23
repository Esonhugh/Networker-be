package handlers

import (
	"Network-be/config"
	"github.com/gin-gonic/gin"
)

func SetCorsPolicy(c *gin.Context) {
	method := c.Request.Method
	origin := c.Request.Header.Get("Origin")
	if origin != "" {
		c.Header("Access-Control-Allow-Origin",
			config.GlobalConfig.GetString("server.cors.origin"))
		c.Header("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods",
			config.GlobalConfig.GetString("server.cors.methods"))
		c.Header("Access-Control-Allow-Credentials",
			config.GlobalConfig.GetString("server.cors.allowCredentials"))
	}
	if method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
