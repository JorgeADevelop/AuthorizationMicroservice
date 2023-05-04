package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMySQLDataBase() (err error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaserName := os.Getenv("DB_DATABASE")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, host, port, databaserName)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	err = migrateModels()
	return
}

func migrateModels() error {
	return Db.AutoMigrate()
}
