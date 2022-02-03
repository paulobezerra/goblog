package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/paulobezerra/goblog/src/utils"
)

func GetConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:teste123@tcp(localhost:49153)/go_blog?parseTime=true")
	utils.CheckErr(err)
	return db
}
