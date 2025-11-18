package db

import (
	"fmt"
	"strings"
	"task4/internal/config"
	"task4/internal/model"
	"task4/pkg/log"

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
		fmt.Println("connect db failed,err: %v", err)
	}
}

func CreateTable() {
	if err := MySQLDB.Debug().AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}); err != nil {
		log.Logger.Error("CreateTable error")
	}

	log.Logger.Info("CreateTable success")

}
