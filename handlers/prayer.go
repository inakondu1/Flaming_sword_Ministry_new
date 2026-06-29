package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/models"
)

func PrayerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/prayer.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	prayer := models.Prayer{
		Name:    r.FormValue("name"),
		Request: r.FormValue("request"),
	}

	err := database.CreatePrayer(prayer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/prayer?success=1", http.StatusSeeOther)
}
