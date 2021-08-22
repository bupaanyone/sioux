package service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/bupaanyone/sioux/server/config"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open(config.C.Service.DbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
