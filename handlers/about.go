package handlers

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
