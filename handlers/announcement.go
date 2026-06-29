package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
)

func ViewAnnouncementsHandler(w http.ResponseWriter, r *http.Request) {

	announcements, err := database.GetAllAnnouncements()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/announcement.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, announcements)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateAnnouncementHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/create_announcement.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	title := r.FormValue("title")
	message := r.FormValue("message")

	err := database.CreateAnnouncement(title, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
