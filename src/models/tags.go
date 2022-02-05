package models

import (
	"fmt"
	"log"

	"github.com/paulobezerra/goblog/src/configs"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func GetTag(id int) *Tag {
	db := configs.GetConnect()

	var category Tag
	db.First(&category, id)

	return &category
}

func FindAllTags() []Tag {
	db := configs.GetConnect()

	var tags []Tag
	db.Find(&tags)

	return tags
}

func FindTagsById(ids []int) []Tag {
	db := configs.GetConnect()

	fmt.Println("BUSCANDO TAGS PELOS IDS", ids)

	var tags []Tag
	db.Find(&tags, "id in (?)", ids)

	return tags
}

func (category *Tag) Create() bool {
	db := configs.GetConnect()
	if err := db.Create(&category).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (category *Tag) Save() bool {
	db := configs.GetConnect()
	if err := db.Save(&category).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (category *Tag) Delete() bool {
	db := configs.GetConnect()
	if err := db.Delete(&category).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
