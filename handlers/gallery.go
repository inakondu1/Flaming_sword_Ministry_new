package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/models"
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

// ================= ADD GALLERY =================

func AddGalleryHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/add_gallery.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := header.Filename

	dst, err := os.Create(filepath.Join("uploads/gallery", filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gallery := models.Gallery{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Image:       "/uploads/gallery/" + filename,
	}

	err = database.CreateGallery(gallery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/gallery", http.StatusSeeOther)
}

// ================= DELETE GALLERY =================

func DeleteGalleryHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	err := database.DeleteGallery(id)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	http.Redirect(w, r, "/gallery", http.StatusSeeOther)
}
