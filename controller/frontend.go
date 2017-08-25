package controller

import (
	"log"
	"net/http"
	"html/template"
)

var (
	tpl *template.Template
)

type Frontend struct {
	Title string
}

func init()  {
	tpl = template.Must(template.ParseGlob("views/*.html"))
}

func Home(res http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(res, "home.html", Frontend{Title: "Home"})
	if err != nil {
		log.Fatal(err)
		http.Error(res, "Internal server error", http.StatusInternalServerError)
	}
}