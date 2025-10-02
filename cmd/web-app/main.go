package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("GET /", ShowHomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
