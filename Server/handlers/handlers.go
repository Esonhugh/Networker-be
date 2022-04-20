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
	var peerList peerinfo.PeerList
	db.DBService.MainDB.Model(&PO.Auth{}).Find(&peerList)
	c.JSON(200, peerList)
}

func GetPeerInfo(c *gin.Context) {
	var peerInfo peerinfo.DetailPeer
	db.DBService.MainDB.Model(&PO.Auth{}).Where("peer_id = ?", c.Param("id")).First(&peerInfo)
	c.JSON(200, peerInfo)
}

func UpdatePeerInfo(c *gin.Context) {
	var NewData peerinfo.UpdateInfo
	c.Get("username")
	err := c.ShouldBindJSON(NewData)
	if err != nil {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "400",
			ErrorMsg:  "Bad Struct:" + err.Error(),
		})
	}
	// ToDo: Check Username is Correct?
	if NewData.Username == "" {
		c.JSON(400, VO.CommonResp{
			ErrorCode: "400",
			ErrorMsg:  "Bad User, Try Hack Others Config",
		})
	}
	// db.Where(User{Name: "jinzhu"}).Assign(User{Age: 20}).FirstOrCreate(&user)
	db.DBService.MainDB.Where(PO.Config{Username: NewData.Username}).
		Assign(NewData).
		FirstOrCreate(&PO.Config{})
	// if has Username == NewData.Username update the row
	// if not equal, create a new row
	c.JSON(200, VO.CommonResp{
		ErrorCode: "0",
		ErrorMsg:  "Update Success",
	})
}
