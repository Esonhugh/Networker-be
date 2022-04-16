package router

import "github.com/gin-gonic/gin"

type GinList struct {
	MainWeb *gin.Engine
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
	return GinServer
}
