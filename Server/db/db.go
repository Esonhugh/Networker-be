package db

import (
	"Network-be/config"
	"fmt"
	"github.com/patrickmn/go-cache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type DBList struct {
	MainDB      *gorm.DB
	TicketCache *cache.Cache
}

var DBService *DBList

func CreateDB(dbInfo struct {
	Addr           string
	User           string
	Pass           string
	DB             string
	ConnectTimeout uint
}) (*gorm.DB, error) {
	//fmt.Println(CreateDSN(a))
	cfg := struct {
		Addr string
		User string
		Pass string
		DB   string
	}{dbInfo.Addr, dbInfo.User, dbInfo.Pass, dbInfo.DB}
	DB, err := gorm.Open(mysql.Open(CreateDSN(cfg)), &gorm.Config{PrepareStmt: true})
	return DB, err
}

func InitDB() *DBList {
	MainDB, err := CreateDB(struct {
		Addr           string
		User           string
		Pass           string
		DB             string
		ConnectTimeout uint
	}{
		Addr:           config.GlobalConfig.GetString("db.mainDB.Addr"),
		User:           config.GlobalConfig.GetString("db.mainDB.User"),
		Pass:           config.GlobalConfig.GetString("db.mainDB.Pass"),
		DB:             config.GlobalConfig.GetString("db.mainDB.DB"),
		ConnectTimeout: 10})
	if err != nil {
		log.Panic("connect DB error: ", err.Error())
	}
	DBService = &DBList{
		MainDB: MainDB,
	}
	DBService.TicketCache = cache.New(time.Hour*24, time.Hour*24)
	return DBService
}

func CreateDSN(dbInfo struct {
	Addr string
	User string
	Pass string
	DB   string
}) string {
	//user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbInfo.User, dbInfo.Pass, dbInfo.Addr, dbInfo.DB)
}
