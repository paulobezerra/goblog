package db

import "github.com/paulobezerra/goblog/src/utils"

func InitDB() {
	db := GetConnect()
	defer db.Close()

	_, errCreateUsers := db.Exec(`
	create table if not exists users(
		id integer auto_increment NOT NULL,
		username varchar(20) NOT NULL,
		password varchar(100) NOT NULL,
		firstname varchar(100) NOT NULL,
		lastname varchar(100) NOT NULL,
		CONSTRAINT users_PK PRIMARY KEY (id),
		CONSTRAINT users_UN UNIQUE KEY (username)
	);
	`)
	utils.CheckErr(errCreateUsers)

	_, errCreatePosts := db.Exec(`
	create table if not exists posts(
		id integer auto_increment,
		slug varchar(255) NOT NULL,
		title varchar(255) NOT NULL,
		content blob NOT NULL,
		tags varchar(255) NOT NULL,
		abstract varchar(500) NOT NULL,
		updated_at datetime NOT NULL,
		CONSTRAINT posts_PK PRIMARY KEY (id),
		CONSTRAINT posts_UN UNIQUE KEY (slug)
	);
	`)
	utils.CheckErr(errCreatePosts)

	adm, err3 := db.Query(`select * from users where username = ?`, "admin")
	utils.CheckErr(err3)

	if !adm.Next() {
		password, _ := utils.HashPassword("123456")
		stmt, _ := db.Prepare("insert into users(username, password, firstname, lastname) values(?, ?, ?, ?)")
		_, err4 := stmt.Exec("admin", password, "Administrador", "Site")
		utils.CheckErr(err4)
	}
}
