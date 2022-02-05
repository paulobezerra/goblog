package models

import (
	"log"

	"github.com/paulobezerra/goblog/src/configs"
	"github.com/paulobezerra/goblog/src/utils"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"-"`
}

func GetUser(id int) User {
	db := configs.GetConnect()

	var user User
	db.First(&user, id)

	return user
}

func FindOneUserByUsername(username string) User {
	db := configs.GetConnect()

	var user User
	db.Find(&user, "username = ?", username)

	return user
}

func FindAllUsers() []User {
	db := configs.GetConnect()

	var users []User

	db.Find(&users)

	return users
}

func (user *User) Create() bool {
	db := configs.GetConnect()
	user.Password, _ = utils.HashPassword(user.Password)
	if err := db.Create(&user).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (user *User) Save() bool {
	db := configs.GetConnect()

	oldUser := GetUser(user.Id)
	password := oldUser.Password
	if user.Password != "" {
		password, _ = utils.HashPassword(user.Password)
	}
	user.Password = password

	if err := db.Save(&user).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (user *User) Delete() bool {
	db := configs.GetConnect()
	if err := db.Delete(&user).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
