package main

import (
	"html/template"
	"log"
	"net/http"
)

type pageData struct {
	Title     string
	FirstName string
	LastName  string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/contact", contact)
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

func contact(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title: "Thank You!",
	}

	if r.Method == http.MethodPost {
		pd.FirstName = r.FormValue("fname")
		pd.LastName = r.FormValue("lname")
	}

	err := tmpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println(err)
	}
}
