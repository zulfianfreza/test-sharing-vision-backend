package app

import (
	"github.com/zulfianfreza/test-sharing-vision-backend/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	exception.PanicIfError(err)

	return db
}
