package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/cdlavacudeg/go-budget-guardian/config"
)

var DB *gorm.DB

var configDB, _ = config.GetConfig()

var (
	dbUsername = configDB.DB.Username
	dbPassword = configDB.DB.Password
	dbName     = configDB.DB.Name
	dbHost     = configDB.DB.Host
	dbPort     = configDB.DB.Port
)

func InitDB() *gorm.DB {
	DB = connectDB()
	return DB
}

func connectDB() *gorm.DB {
	var err error

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database: err=%v", err)
		return nil
	}

	return db
}
