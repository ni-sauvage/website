package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("html/edit.html", "html/home.html"))

type Page struct {
	Title string
	Body  []byte
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := r.URL.Path[1 : len(r.URL.Path)-1]
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/home/", pageHandler)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
