package db

import (
	"log"

	"github.com/paulobezerra/goblog/src/configs"
	"github.com/paulobezerra/goblog/src/models"
	"github.com/paulobezerra/goblog/src/utils"
)

func Seed() {
	db := configs.GetConnect()

	var user models.User
	log.Println("Verify admin exists")
	db.First(&user, "username = ?", "admin")

	if user.Id == 0 {
		password, _ := utils.HashPassword("123456")
		log.Println("Create admin user")
		if err := db.Create(&models.User{Username: "admin", Password: password, Firstname: "Administrador", Lastname: "Site"}).Error; err != nil {
			utils.CheckErr(err)
		}
	}
}
