package handlers

import (
	"net/http"
	"text/template"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	templ := template.Must(template.ParseFiles("views/about.html"))
	templ.Execute(w, nil)
}
