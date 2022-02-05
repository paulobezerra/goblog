package models

import (
	"log"

	"github.com/paulobezerra/goblog/src/configs"
)

type Category struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func GetCategory(id int) *Category {
	db := configs.GetConnect()

	var category Category
	db.First(&category, id)

	return &category
}

func FindAllCategories() []Category {
	db := configs.GetConnect()

	var categories []Category
	db.Find(&categories)

	return categories
}

func (category *Category) Create() bool {
	db := configs.GetConnect()
	if err := db.Create(&category).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (category *Category) Save() bool {
	db := configs.GetConnect()
	if err := db.Save(&category).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (category *Category) Delete() bool {
	db := configs.GetConnect()
	if err := db.Delete(&category).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
