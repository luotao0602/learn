package db

import (
	"strings"
	"task4/internal/config"
	"task4/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQLDB *gorm.DB

func InitDB() {
	cf := config.GetConfig()
	var stringbuilder strings.Builder
	stringbuilder.WriteString(cf.MySQL.UserName)
	stringbuilder.WriteString(":")
	stringbuilder.WriteString(cf.MySQL.Password)
	stringbuilder.WriteString("@tcp(")
	stringbuilder.WriteString(cf.MySQL.Host)
	stringbuilder.WriteString(":")
	stringbuilder.WriteString(cf.MySQL.Port)
	stringbuilder.WriteString(")/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	url := stringbuilder.String()
	var err error
	MySQLDB, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		logrus.Error("connect db failed,err: %v", err)
	}
}

func CreateTable() {
	if err := MySQLDB.Debug().AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}); err != nil {
		logrus.Error("CreateTable error")
	}

	logrus.Info("CreateTable success")

}

func GetGormDB() *gorm.DB {
	return MySQLDB
}
