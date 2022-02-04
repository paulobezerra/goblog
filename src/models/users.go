package models

import (
	"github.com/paulobezerra/goblog/src/db"
	"github.com/paulobezerra/goblog/src/utils"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"-"`
}

func GetUser(id string) User {
	db := db.GetConnect()

	var user User
	db.First(&user, id)

	return user
}

func FindOneUserByUsername(username string) User {
	db := db.GetConnect()

	var user User
	db.Find(&user, "username = ?", username)

	return user
}

func FindAllUsers() []User {
	db := db.GetConnect()

	var users []User

	db.Find(&users)

	return users
}

func CreateUser(username string, firstname string, lastname string, password string) User {
	db := db.GetConnect()

	user := User{Username: username, Firstname: firstname, Lastname: lastname, Password: password}

	db.Create(&user)

	return user
}

func UpdateUser(id string, username string, firstname string, lastname string, password string) User {
	db := db.GetConnect()

	var user User
	db.First(&user, id)

	user.Username = username
	user.Firstname = firstname
	user.Lastname = lastname
	if password != "" {
		passwordHash, _ := utils.HashPassword(password)
		user.Password = passwordHash
	}
	db.Save(&user)

	return user
}

func DeleteUser(id string) {
	db := db.GetConnect()
	db.Delete(&User{}, id)
}
