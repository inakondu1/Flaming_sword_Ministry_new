package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/models"
)

// ================= ADD SERMON =================

func AddSermonHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/add_sermon.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		sermon := models.Sermon{
			Title:      r.FormValue("title"),
			BibleVerse: r.FormValue("bible_verse"),
			References: r.FormValue("references"),
			Content:    r.FormValue("content"),
			Category:   r.FormValue("category"),
			Date:       r.FormValue("date"),
			CreatedBy:  r.FormValue("created_by"),
		}

		err := database.CreateSermon(sermon)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/sermons", http.StatusSeeOther)
	}
}

// ================= VIEW ALL SERMONS =================

func ViewSermonsHandler(w http.ResponseWriter, r *http.Request) {

	sermons, err := database.GetAllSermons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/sermons.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, sermons)
}

// ================= VIEW ONE SERMON =================

func ViewSingleSermonHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid sermon ID", http.StatusBadRequest)
		return
	}

	sermon, err := database.GetSermonByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/view_sermon.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, sermon)
}

// ================= EDIT SERMON =================

func EditSermonHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid sermon ID", http.StatusBadRequest)
			return
		}

		sermon, err := database.GetSermonByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/edit_sermon.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, sermon)
		return
	}

	if r.Method == http.MethodPost {

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Invalid sermon ID", http.StatusBadRequest)
			return
		}

		sermon := models.Sermon{
			ID:         id,
			Title:      r.FormValue("title"),
			BibleVerse: r.FormValue("bible_verse"),
			References: r.FormValue("references"),
			Content:    r.FormValue("content"),
			Category:   r.FormValue("category"),
			Date:       r.FormValue("date"),
			CreatedBy:  r.FormValue("created_by"),
		}

		err = database.UpdateSermon(sermon)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/sermons", http.StatusSeeOther)
	}
}

// ================= DELETE SERMON =================

func DeleteSermonHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid sermon ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteSermon(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/sermons", http.StatusSeeOther)
}
