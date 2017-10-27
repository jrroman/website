package main

import (
	"html/template"
	"log"
	"net/http"
)

type formData struct {
	Name    string
	Email   string
	Subject string
	Message string
}

type pageData struct {
	Title    string
	FormData formData
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
		pd.FormData.Name = r.FormValue("name")
		pd.FormData.Email = r.FormValue("email")
		pd.FormData.Subject = r.FormValue("subject")
		pd.FormData.Message = r.FormValue("message")
	}

	err := tmpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println(err)
	}
}
