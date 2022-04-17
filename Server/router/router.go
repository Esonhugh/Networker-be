package router

import (
	"Network-be/Server/router/jwt"
	"github.com/gin-gonic/gin"
)

type GinList struct {
	MainWeb *gin.Engine
	JWTAuth *jwt.JWT
}

var GinServer *GinList

func init() {
	GinServer = InitGin()
}

func CreateGin() *gin.Engine {
	return gin.Default()
}

func InitGin() *GinList {
	GinServer := &GinList{
		MainWeb: CreateGin(),
	}
	jwt.Init()
	return GinServer
}
