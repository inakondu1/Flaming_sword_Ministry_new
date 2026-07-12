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
	TotalContacts      int
	TotalEvents        int
	TotalGallery       int

	Users    interface{}
	Prayers  interface{}
	Contacts interface{}
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := middleware.Store.Get(r, "church-session")

	name, _ := session.Values["name"].(string)

	totalUsers, _ := database.CountUsers()
	totalSermons, _ := database.CountSermons()
	totalAnnouncements, _ := database.CountAnnouncements()
	totalContacts, _ := database.CountContacts()

	// NEW
	totalEvents, _ := database.CountEvents()
	totalGallery, _ := database.CountGallery()

	prayers, _ := database.GetAllPrayers()
	contacts, _ := database.GetAllContacts()
	users, _ := database.GetAllUsers()

	totalPrayers := len(prayers)

	data := AdminData{
		Name:               name,
		TotalUsers:         totalUsers,
		TotalSermons:       totalSermons,
		TotalAnnouncements: totalAnnouncements,
		TotalPrayers:       totalPrayers,
		TotalContacts:      totalContacts,
		TotalEvents:        totalEvents,
		TotalGallery:       totalGallery,

		Users:    users,
		Prayers:  prayers,
		Contacts: contacts,
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
