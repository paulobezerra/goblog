package helpers

import (
	"net/http"
	"text/template"

	"github.com/paulobezerra/goblog/src/utils"
)

var templates = template.Must(template.ParseGlob("templates/**.html"))

func RenderTemplate(w http.ResponseWriter, layout string, template string, data interface{}) {

	files := []string{
		"templates/" + layout + ".html",
		"templates/" + template + ".html",
	}
	ts, parseTemplateError := templates.ParseFiles(files...)
	utils.CheckErr(parseTemplateError)

	executeTemplateError := ts.Execute(w, data)
	utils.CheckErr(executeTemplateError)

}
