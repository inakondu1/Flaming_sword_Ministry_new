package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/middleware"
	"Flaming_Sword_Ministry/models"
)

type AdminPage struct {
	Name               string
	TotalUsers         int
	TotalSermons       int
	TotalAnnouncements int
	Users              []models.User
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := middleware.Store.Get(r, "church-session")

	name, _ := session.Values["name"].(string)

	users, _ := database.GetAllUsers()

	totalUsers, _ := database.CountUsers()
	totalSermons, _ := database.CountSermons()
	totalAnnouncements, _ := database.CountAnnouncements()

	data := AdminPage{
		Name:               name,
		TotalUsers:         totalUsers,
		TotalSermons:       totalSermons,
		TotalAnnouncements: totalAnnouncements,
		Users:              users,
	}

	tmpl, err := template.ParseFiles("templates/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
