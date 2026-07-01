package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/middleware"
)

type HomePageData struct {
	Name          string
	Role          string
	Sermons       interface{}
	Announcements interface{}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := middleware.Store.Get(r, "church-session")

	name, _ := session.Values["name"].(string)
	role, _ := session.Values["role"].(string)

	sermons, _ := database.GetAllSermons()
	announcements, _ := database.GetAllAnnouncements()

	data := HomePageData{
		Name:          name,
		Role:          role,
		Sermons:       sermons,
		Announcements: announcements,
	}

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
