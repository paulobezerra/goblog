package db

import (
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/paulobezerra/goblog/src/utils"
	"github.com/paulobezerra/goblog/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnect() *gorm.DB {
	// db, err := sql.Open("mysql", "root:teste123@tcp(localhost:49153)/go_blog?parseTime=true")
	db, err := gorm.Open(mysql.Open("root:teste123@tcp(localhost:49154)/go_blog?parseTime=true"), &gorm.Config{})
	utils.CheckErr(err)
	return db
}
