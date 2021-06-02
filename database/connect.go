package database

import (
	"fmt"
	"log"

	"github.com/vivekv96/go-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

type MySQLConfig struct {
	Host     string
	Username string
	Password string
	Port     int
	DBName   string
}

func ConnectToMySQL(conf *MySQLConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Username, conf.Password,
		conf.Host, fmt.Sprint(conf.Port), conf.DBName)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	Gorm = db

	Gorm.Exec("CREATE DATABASE IF NOT EXISTS `admin`")
	Gorm.Exec("USE `admin`")

	if err := Gorm.AutoMigrate(&models.User{}); err != nil {
		log.Fatalln("auto-migrate failed, err: ", err)
	}

	return nil
}
