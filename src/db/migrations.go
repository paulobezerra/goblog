package db

import (
	"log"

	"github.com/paulobezerra/goblog/src/configs"
	"github.com/paulobezerra/goblog/src/models"
	"github.com/paulobezerra/goblog/src/utils"
)

func Migrate() {
	db := configs.GetConnect()

	log.Println("Run users migration")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		utils.CheckErr(err)
	}

	log.Println("Run posts migration")
	if err := db.AutoMigrate(&models.Post{}); err != nil {
		utils.CheckErr(err)
	}
}
