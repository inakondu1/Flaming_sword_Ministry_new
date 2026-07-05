package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/models"
)

// ================= CONTACT PAGE =================

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	// Display Contact Page
	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/contact.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Success bool
		}{
			Success: r.URL.Query().Get("success") == "1",
		}

		tmpl.Execute(w, data)
		return
	}

	// Save Contact Message
	contact := models.Contact{
		FullName: r.FormValue("fullname"),
		Phone:    r.FormValue("phone"),
		Subject:  r.FormValue("subject"),
		Message:  r.FormValue("message"),
	}

	err := database.CreateContact(contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/contact?success=1", http.StatusSeeOther)
}

// ================= VIEW CONTACTS =================

func ViewContactsHandler(w http.ResponseWriter, r *http.Request) {

	contacts, err := database.GetAllContacts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/admin_contacts.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, contacts)
}

// ================= DELETE CONTACT =================

func DeleteContactHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteContact(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/contacts", http.StatusSeeOther)
}
