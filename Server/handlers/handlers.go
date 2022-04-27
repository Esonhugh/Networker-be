package handlers

import (
	"Network-be/Server/db"
	"Network-be/data/PO"
	"Network-be/data/VO"
	"Network-be/data/VO/getConfig"
	"Network-be/data/VO/peerinfo"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, VO.CommonResp{
		ErrorCode: "0",
		ErrorMsg:  "server up succcss",
	})
}

func GetConfig(c *gin.Context) {
	c.JSON(200, getConfig.AdminConfig())
}

func GetPeerList(c *gin.Context) {
	var peerList []peerinfo.SimplePeer
	db.DBService.MainDB.Model(&PO.Config{}).Find(&peerList)
	c.JSON(200, peerList)
}

func GetPeerInfo(c *gin.Context) {
	var peerInfo peerinfo.DetailPeer
	if db.DBService.MainDB.Model(&PO.Config{}).Where("id = ?", c.Param("id")).First(&peerInfo).Error == nil {
		c.JSON(200, peerInfo)
	} else {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "1",
			ErrorMsg:  "peer not found",
		})
	}
}

// GetMyInfo 获取自己的信息
func GetMyInfo(c *gin.Context) {
	var myInfo peerinfo.DetailPeer
	user, ok := c.Get("username")
	if ok {
		if db.DBService.MainDB.Model(&PO.Config{}).Where("username = ?", user.(string)).First(&myInfo).Error != nil {
			/*
				c.JSON(400, VO.CommonResp{
					ErrorCode: "0",
					ErrorMsg:  "You have not Create Yet",
				})
			*/
			c.JSON(200, myInfo)
		} else {
			c.JSON(200, myInfo)
		}
	} else {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "40001",
			ErrorMsg:  "Unkoown error",
		})
	}
}

func UpdatePeerInfo(c *gin.Context) {
	var NewData peerinfo.UpdateInfo
	Username, ok := c.Get("username")
	if !ok { // not Exist
		c.JSON(400, VO.CommonResp{
			ErrorCode: "40022",
			ErrorMsg:  "user not Login",
		})
		c.Abort()
		return
	}
	err := c.ShouldBindJSON(&NewData)
	if err != nil {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "40023",
			ErrorMsg:  "Bad Struct:" + err.Error(),
		})
		return
	}
	if Username.(string) != NewData.Username {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "40024",
			ErrorMsg:  "Username is not matched with login User",
		})
		c.Abort()
		return
	}
	// db.Where(User{Name: "jinzhu"}).Assign(User{Age: 20}).FirstOrCreate(&user)
	db.DBService.MainDB.Model(&PO.Config{}).Where(&PO.Config{Username: NewData.Username}).
		Assign(NewData).
		FirstOrCreate(&PO.Config{})
	// if has Username == NewData.Username update the row
	// if not equal, create a new row
	c.JSON(200, VO.CommonResp{
		ErrorCode: "0",
		ErrorMsg:  "Update Success",
	})
}
