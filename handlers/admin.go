package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/middleware"
)

type AdminData struct {
	Name               string
	TotalUsers         int
	TotalSermons       int
	TotalAnnouncements int
	TotalPrayers       int
	Users              interface{}
	Prayers            interface{}
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := middleware.Store.Get(r, "church-session")

	name, _ := session.Values["name"].(string)

	totalUsers, _ := database.CountUsers()
	totalSermons, _ := database.CountSermons()
	totalAnnouncements, _ := database.CountAnnouncements()

	prayers, _ := database.GetAllPrayers()
	users, _ := database.GetAllUsers()

	totalPrayers := len(prayers)

	data := AdminData{
		Name:               name,
		TotalUsers:         totalUsers,
		TotalSermons:       totalSermons,
		TotalAnnouncements: totalAnnouncements,
		TotalPrayers:       totalPrayers,
		Users:              users,
		Prayers:            prayers,
	}

	tmpl, err := template.ParseFiles("templates/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewUsersHandler(w http.ResponseWriter, r *http.Request) {

	users, err := database.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/user.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, users)
}
