package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/middleware"
)

type HomePageData struct {
	Name          string
	Sermons       interface{}
	Announcements interface{}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// Get logged-in user (if any)
	session, _ := middleware.Store.Get(r, "church-session")

	name, _ := session.Values["name"].(string)

	// Load sermons
	sermons, _ := database.GetAllSermons()

	// Load announcements
	announcements, _ := database.GetAllAnnouncements()

	data := HomePageData{
		Name:          name,
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
