package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
)

// ================= VIEW GALLERY =================

func GalleryHandler(w http.ResponseWriter, r *http.Request) {

	gallery, err := database.GetAllGallery()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/gallery.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, gallery)
}
