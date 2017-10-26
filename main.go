package main

import (
	"html/template"
	"log"
	"net/http"
)

type pageData struct {
	Title string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title: "Index",
	}

	err := tmpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println(err)
	}
}
