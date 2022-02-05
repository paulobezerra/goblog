package configs

import (
	"github.com/paulobezerra/goblog/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConnect() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:teste123@tcp(localhost:49154)/go_blog?parseTime=true"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	utils.CheckErr(err)
	return db
}
