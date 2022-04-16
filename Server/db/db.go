package db

import (
	"Network-be/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DBList struct {
	MainDB *gorm.DB
}

var DBService *DBList

func Init() {
	DBService = InitDB()
}

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
	dbList := new(DBList)
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
	// other DBs
	dbList.MainDB = MainDB
	return dbList
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
