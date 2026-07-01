package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/models"
)

// ================= PUBLIC PRAYER REQUEST =================

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
		Status:  "Pending",
	}

	err := database.CreatePrayer(prayer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/prayer", http.StatusSeeOther)
}

// ================= ADMIN VIEW PRAYERS =================

func ViewPrayersHandler(w http.ResponseWriter, r *http.Request) {

	prayers, err := database.GetAllPrayers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/admin_prayers.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, prayers)
}

// ================= DELETE PRAYER =================

func DeletePrayerHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	err := database.DeletePrayer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/prayers", http.StatusSeeOther)
}
