package models

import (
	"fmt"

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

func GetUser(id string) *User {
	conn := db.GetConnect()
	defer conn.Close()

	row := conn.QueryRow("select id, username, firstname, lastname, password from users where id = ?", id)

	if row == nil {
		return nil
	}

	u := User{}

	row.Scan(&u.Id, &u.Username, &u.Firstname, &u.Lastname, &u.Password)

	return &u
}

func FindOneUserByUsername(username string) *User {
	conn := db.GetConnect()
	defer conn.Close()

	row := conn.QueryRow("select id, username, firstname, lastname, password from users where username = ?", username)

	if row == nil {
		return nil
	}

	u := User{}

	row.Scan(&u.Id, &u.Username, &u.Firstname, &u.Lastname, &u.Password)

	return &u
}

func FindAllUsers() *[]User {
	conn := db.GetConnect()
	defer conn.Close()

	rows, err := conn.Query("select id, username, firstname, lastname from users")
	utils.CheckErr(err)
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		u := User{}
		rows.Scan(&u.Id, &u.Username, &u.Firstname, &u.Lastname)
		users = append(users, u)
	}

	return &users
}

func ValidateUser(username string, firstname string, lastname string, password string, update bool) (map[string]string, bool) {
	var messages = map[string]string{}
	var valid = true

	if username == "" {
		messages["Username"] = "Nome do usuário deve ser informado"
		valid = false
	}

	if firstname == "" {
		messages["Firstname"] = "Primeiro nome deve ser informado"
		valid = false
	}

	if lastname == "" {
		messages["Lastname"] = "Sobrenome deve ser informado"
		valid = false
	}
	if !update {
		user := FindOneUserByUsername(username)
		if user.Id != 0 {
			messages["Username"] = "Já existe um usuário com este nome"
			valid = false
		}

		if password == "" {
			messages["Password"] = "Senha deve ser informado"
			valid = false
		} else if len(password) < 6 {
			messages["Password"] = "Senha deve conter no mínimo 6 caracteres"
			valid = false
		}
	}

	return messages, valid
}

func CreateUser(username string, firstname string, lastname string, password string) *string {
	conn := db.GetConnect()
	defer conn.Close()

	stmt, err := conn.Prepare("insert into users (username, firstname, lastname, password) values (?, ?, ?, ?);")
	if err != nil {
		fmt.Print(err.Error())
		message := "Erro ao inserir usuário no banco de dados."
		return &message
	}
	defer stmt.Close()
	passwordHash, _ := utils.HashPassword(password)
	_, errExecute := stmt.Exec(username, firstname, lastname, passwordHash)
	if errExecute != nil {
		fmt.Print(errExecute.Error())
		message := "Erro ao inserir usuário no banco de dados."
		return &message
	}

	return nil
}

func UpdateUser(id string, username string, firstname string, lastname string, password string) *string {
	conn := db.GetConnect()
	defer conn.Close()

	var query string = "update users set username = ?, firstname = ?, lastname = ?"
	fmt.Println(password)
	if password != "" {
		query += ", password = ?"
	}
	query += " where id = ?"

	stmt, err := conn.Prepare(query)
	if err != nil {
		fmt.Print(err.Error())
		message := "Erro ao inserir usuário no banco de dados."
		return &message
	}
	defer stmt.Close()
	var errExecute error
	if password != "" {
		passwordHash, _ := utils.HashPassword(password)
		_, errExecute = stmt.Exec(username, firstname, lastname, passwordHash, id)
	} else {
		_, errExecute = stmt.Exec(username, firstname, lastname, id)
	}
	if errExecute != nil {
		fmt.Print(errExecute.Error())
		message := "Erro ao atualizar usuário no banco de dados."
		return &message
	}

	return nil
}

func DeleteUser(id string) *string {
	conn := db.GetConnect()
	defer conn.Close()

	var query string = "delete from users where id = ?"

	stmt, err := conn.Prepare(query)
	if err != nil {
		fmt.Print(err.Error())
		message := "Erro ao inserir usuário no banco de dados."
		return &message
	}
	defer stmt.Close()
	_, errExecute := stmt.Exec(id)
	if errExecute != nil {
		fmt.Print(errExecute.Error())
		message := "Erro ao excluir usuário no banco de dados."
		return &message
	}

	return nil
}
