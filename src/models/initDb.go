package models

import (
	"log"

	"github.com/paulobezerra/goblog/src/db"
	"github.com/paulobezerra/goblog/src/utils"
)

func InitDB() {
	db := db.GetConnect()

	log.Println("Run users migration")
	if err := db.AutoMigrate(&User{}); err != nil {
		utils.CheckErr(err)
	}

	log.Println("Run posts migration")
	if err := db.AutoMigrate(&Post{}); err != nil {
		utils.CheckErr(err)
	}

	var user User
	log.Println("Verify admin exists")
	db.First(&user, "username = ?", "admin")

	if user.Id == 0 {
		password, _ := utils.HashPassword("123456")
		log.Println("Create admin user")
		if err := db.Create(&User{Username: "admin", Password: password, Firstname: "Administrador", Lastname: "Site"}).Error; err != nil {
			utils.CheckErr(err)
		}
	}
}
