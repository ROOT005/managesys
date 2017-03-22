package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB, err = gorm.Open("mysql", "root:special005@/admin?charset=utf8&parseTime=True&loc=Local")
