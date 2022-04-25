package jwt

import (
	"Network-be/data/VO"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware JWT中间件 CheckJWTtoken
func JWTAuthMiddleware(c *gin.Context) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	// authHeader, err := c.Cookie("Token")
	auth := strings.Split(c.GetHeader("Auth"), " ")
	if len(auth) != 1 {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "1",
			ErrorMsg:  "Token Error",
		})
		c.Abort()
		return
	}
	authHeader := auth[0]
	if authHeader == "" {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "1",
			ErrorMsg:  "Token Empty",
		})
		c.Abort()
		return
	}
	// 解析鉴权
	mc, err := ParseToken(authHeader)
	if err != nil {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "1",
			ErrorMsg:  "Token Valid",
		})
		c.Abort()
		return
	}
	// 将当前请求的username信息保存到请求的上下文c上
	c.Set("username", mc.Username)
	c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
}
