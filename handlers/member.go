package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
)

func ViewMembersHandler(w http.ResponseWriter, r *http.Request) {

	users, err := database.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/members.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
