package models

import (
	"log"

	"github.com/paulobezerra/goblog/src/configs"
)

type Tag struct {
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

	var categories []Tag
	db.Find(&categories)

	return categories
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
