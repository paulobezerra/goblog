package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/paulobezerra/goblog/src/db"
	"github.com/paulobezerra/goblog/src/utils"
)

type Post struct {
	Id        int       `json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      string    `json:"tags"`
	Abstract  string    `json:"abstract"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetPost(id string) *Post {
	conn := db.GetConnect()
	defer conn.Close()

	row := conn.QueryRow("select id, slug, title, abstract, content, tags, updated_at from posts where id = ?", id)

	if row == nil {
		return nil
	}

	p := Post{}

	row.Scan(&p.Id, &p.Slug, &p.Title, &p.Abstract, &p.Content, &p.Tags, &p.UpdatedAt)

	return &p
}

func FindOnePostBySlug(slug string) *Post {
	conn := db.GetConnect()
	defer conn.Close()

	row := conn.QueryRow("select id, slug, title, abstract, content, tags, updated_at from posts where slug = ?", slug)

	if row == nil {
		return nil
	}

	p := Post{}

	row.Scan(&p.Id, &p.Slug, &p.Title, &p.Abstract, &p.Content, &p.Tags, &p.UpdatedAt)

	return &p
}

func FindAllPosts() *[]Post {
	conn := db.GetConnect()
	defer conn.Close()

	rows, err := conn.Query("select id, slug, title, abstract, tags, updated_at from posts order by id desc")
	utils.CheckErr(err)
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		p := Post{}
		rows.Scan(&p.Id, &p.Slug, &p.Title, &p.Abstract, &p.Tags, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return &posts
}

func ValidatePost(id string, slug string, title string, abstract string, content string, tags string) (map[string]string, bool) {
	var messages = map[string]string{}
	var valid = true

	if slug == "" {
		messages["Slug"] = "Slug deve ser informado"
		valid = false
	}

	if title == "" {
		messages["Title"] = "Título deve ser informado"
		valid = false
	}

	if abstract == "" {
		messages["Abstract"] = "Resumo deve ser informado"
		valid = false
	}

	if tags == "" {
		messages["Tags"] = "Tags deve ser informado"
		valid = false
	}

	post := FindOnePostBySlug(slug)
	idInt, _ := strconv.Atoi(id)
	if post.Id > 0 && post.Id != idInt {
		messages["slug"] = "Já existe um post com este slug"
		valid = false
	}

	return messages, valid
}

func CreatePost(slug string, title string, abstract string, content string, tags string) *string {
	conn := db.GetConnect()
	defer conn.Close()

	stmt, err := conn.Prepare("insert into posts (slug, title, abstract, content, tags, updated_at) values (?, ?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Print(err)
		message := "Erro ao inserir post no banco de dados."
		return &message
	}
	defer stmt.Close()
	_, errExecute := stmt.Exec(slug, title, abstract, content, tags, time.Now())
	if errExecute != nil {
		fmt.Print(errExecute.Error())
		message := "Erro ao inserir post no banco de dados."
		return &message
	}

	return nil
}

func UpdatePost(id string, slug string, title string, abstract string, content string, tags string) *string {
	conn := db.GetConnect()
	defer conn.Close()

	var query string = "update posts set slug = ?, title = ?, abstract = ?, tags = ?, updated_at = ?"
	if content != "" {
		query += ", content = ?"
	}
	query += " where id = ?"

	stmt, err := conn.Prepare(query)
	if err != nil {
		fmt.Print(err.Error())
		message := "Erro ao atualizar post no banco de dados."
		return &message
	}
	defer stmt.Close()
	var errExecute error
	if content != "" {
		_, errExecute = stmt.Exec(slug, title, abstract, tags, time.Now(), content, id)
	} else {
		_, errExecute = stmt.Exec(slug, title, abstract, tags, time.Now(), id)
	}
	if errExecute != nil {
		fmt.Print(errExecute.Error())
		message := "Erro ao atualizar post no banco de dados."
		return &message
	}

	return nil
}

func DeletePost(id string) *string {
	conn := db.GetConnect()
	defer conn.Close()

	var query string = "delete from posts where id = ?"

	stmt, err := conn.Prepare(query)
	if err != nil {
		fmt.Print(err.Error())
		message := "Erro ao excluir post no banco de dados."
		return &message
	}
	defer stmt.Close()
	_, errExecute := stmt.Exec(id)
	if errExecute != nil {
		fmt.Print(errExecute.Error())
		message := "Erro ao excluir post no banco de dados."
		return &message
	}

	return nil
}
