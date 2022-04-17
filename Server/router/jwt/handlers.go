package jwt

import (
	"Network-be/Server/db"
	"Network-be/data/PO"
	"Network-be/data/VO"
	"Network-be/data/VO/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user auth.LoginRequest
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "2001",
			ErrorMsg:  "Invalid Request",
		})
		c.Abort()
	}
	// 校验用户名和密码是否正确
	var UserInDB PO.Auth
	db.DBService.MainDB.Where("username = ?", user.Username).First(&UserInDB)
	if UserInDB.CheckPassword(user.Password) {
		// 生成Token
		tokenString, _ := GenToken(user.Username)
		c.JSON(200, VO.CommonResp{
			ErrorCode: "0",
			ErrorMsg:  "success",
		})
		c.SetCookie("Token", tokenString, 3600, "/", "", false, false)
		return
	}
	c.JSON(200, VO.CommonResp{
		"2002",
		"Authorization failed",
	})
	return
}

func LogoutHandler(c *gin.Context) {
	c.JSON(200, VO.CommonResp{
		ErrorCode: "0",
		ErrorMsg:  "success",
	})
	c.SetCookie("Token", "", -1, "/", "", false, false)
	return
}

func RegisterHandler(c *gin.Context) {
	var user auth.RegRequest
	err := c.ShouldBind(&user)

	if err != nil {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "2001",
			ErrorMsg:  "invalid params",
		})
		c.Abort()
	}

	var UserInDB PO.Auth
	err = db.DBService.MainDB.Where("username = ?", user.Username).First(&UserInDB).Error
	if err == gorm.ErrRecordNotFound {
		// 注册用户
		newUser := &PO.Auth{
			Username: user.Username,
			Email:    user.Email,
			Verify:   false,
		}
		newUser.SetPassword(user.Password)
		db.DBService.MainDB.Create(newUser)
		createVerifyTicket(newUser.Username)
		c.JSON(200, VO.CommonResp{
			ErrorCode: "0",
			ErrorMsg:  "success, Check Your Email: " + newUser.Email,
		})
	} else {
		// 存在用户但是 没有验证
		if UserInDB.Verify == false {
			createVerifyTicket(user.Username)
			c.JSON(200, VO.CommonResp{
				ErrorCode: "0",
				ErrorMsg:  "success, Check Your Email: " + user.Email,
			})
		} else { // 存在用户而且验证
			c.JSON(400, VO.CommonResp{
				ErrorCode: "2002",
				ErrorMsg:  "User already exists,if You forget your password,please contact the administrator",
			})

		}
	}
	return
}

func createVerifyTicket(username string) {
	// todo: 验证码生成
	// todo: 塞入 Redis
	// todo: 发送邮件
}

func VerifyHandler(c *gin.Context) {

}
