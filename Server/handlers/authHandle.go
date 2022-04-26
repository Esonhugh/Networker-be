package handlers

import (
	"Network-be/Server/db"
	"Network-be/Server/router/jwt"
	"Network-be/data/PO"
	"Network-be/data/VO"
	"Network-be/data/VO/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"strings"
	"time"
)

func AuthHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user auth.LoginRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "2001",
			ErrorMsg:  "Invalid Request",
		})
		c.Abort()
		return
	}
	// 校验用户名和密码是否正确
	var UserInDB PO.Auth
	db.DBService.MainDB.Where("username = ?", user.Username).First(&UserInDB)
	// 未验证不许登陆
	if !UserInDB.Verify {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "40033",
			ErrorMsg:  "User Unverified can't Login",
		})
		return
	}
	if UserInDB.CheckPassword(user.Password) {
		// 生成Token
		tokenString, _ := jwt.GenToken(user.Username)
		c.Header("Auth", tokenString)
		c.JSON(200, VO.CommonResp{
			ErrorCode: tokenString,
			ErrorMsg:  "Login Success",
		})
		return
	}
	c.JSON(400, VO.CommonResp{
		"2002",
		"Authorization failed",
	})
	return
}

// LogoutHandler 登出 deleted
/*
func LogoutHandler(c *gin.Context) {
	c.JSON(200, VO.CommonResp{
		ErrorCode: "0",
		ErrorMsg:  "success",
	})
	c.SetCookie("Token", "", -1, "/", "", false, false)
	return
}
*/

// RegisterHandler 注册用户
func RegisterHandler(c *gin.Context) {
	var user auth.RegRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "2001",
			ErrorMsg:  "invalid params, check your request param",
		})
		c.Abort()
		return
	}
	var UserInDB PO.Auth
	err = db.DBService.MainDB.Model(&UserInDB).
		Where("username = ?", user.Username).
		First(&UserInDB).Error
	if err == gorm.ErrRecordNotFound {
		// 找不到用户 就进行注册用户
		newUser := &PO.Auth{
			Username: user.Username,
			Email:    user.Email,
			Verify:   false,
		}
		newUser.SetPassword(user.Password)
		db.DBService.MainDB.Create(&newUser)
		err = CreateVerifyTicket(newUser.Username, newUser.Email)
		if err != nil {
			c.JSON(500, VO.CommonResp{
				ErrorCode: "5001",
				ErrorMsg:  "Internal Server Error: Mail Send Failed",
			})
		} else {
			c.JSON(200, VO.CommonResp{
				ErrorCode: "0",
				ErrorMsg:  "success, Check Your Email: " + newUser.Email,
			})
		}
	} else {
		if UserInDB.Verify == false {
			// 存在用户但是 没有验证
			if user.Email == UserInDB.Email {
				// 当前用户的邮箱和邮箱在数据库中的邮箱一致时 重新发送验证消息
				err = CreateVerifyTicket(UserInDB.Username, UserInDB.Email)
				// 邮件发送失败
				if err != nil {
					c.JSON(400, VO.CommonResp{
						ErrorCode: "5001",
						ErrorMsg:  "Internal Server Error: Mail Send Failed",
					})
					return
				}
				c.JSON(200, VO.CommonResp{
					ErrorCode: "0",
					ErrorMsg:  "success, reCheck Your Email: " + user.Email + ". We have sent you a verification email.",
				})
			} else {
				// 存在用户但是 Email 和表单中的对不上的时候 直接报错
				c.JSON(400, VO.CommonResp{
					ErrorCode: "2003",
					ErrorMsg:  "Email is not matched in database. Are You Change it? ",
				})
				log.Printf("User %v send request. \n"+
					"But Email %v is not matched in database. \n"+
					"Database Email is %v",
					user.Username, user.Email, UserInDB.Email)
			}
		} else {
			// 存在用户而且验证
			c.JSON(400, VO.CommonResp{
				ErrorCode: "2002",
				ErrorMsg:  "User already exists,if You forget your password,please contact the administrator",
			})
		}
	}
	return
}

// createVerifyTicket 创建验证邮件 将 ticket 存入 memcache 中
// username 负责定位 创建并且定位 ticket
func CreateVerifyTicket(username string, email string) error {
	ticket := GenValidateCode(32)
	db.DBService.TicketCache.SetDefault(ticket, username)
	err := SendVerifyByEmail(email, CreateContent(ticket))
	log.Println(err)
	return err
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func VerifyHandler(c *gin.Context) {
	ticket := c.Param("ticket")
	username, ok := db.DBService.TicketCache.Get(ticket)
	if ok {
		var User PO.Auth
		db.DBService.MainDB.Model(&User).Where("username = ?", username.(string)).First(&User)
		User.Verify = true
		db.DBService.MainDB.Save(&User)
		c.JSON(200, VO.CommonResp{
			ErrorCode: "0",
			ErrorMsg:  "You Have being Successful Verified",
		})
	} else {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "40025",
			ErrorMsg:  "Bad Ticket",
		})
	}
}
