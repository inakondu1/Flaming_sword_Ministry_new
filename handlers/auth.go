package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/middleware"
	"Flaming_Sword_Ministry/models"
)

// ================= REGISTER =================

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/register.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		user := models.User{
			FullName: r.FormValue("fullname"),
			Phone:    r.FormValue("phone"),
			Gender:   r.FormValue("gender"),
			Password: r.FormValue("password"),

			// Change this to "member" later if you don't
			// want every new account to be an admin.
			Role: "member",
		}

		err := database.CreateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// ================= LOGIN =================

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		phone := r.FormValue("phone")
		password := r.FormValue("password")

		user, err := database.GetUserByPhone(phone)
		if err != nil {
			http.Error(w, "Invalid phone or password", http.StatusUnauthorized)
			return
		}

		if user.Password != password {
			http.Error(w, "Invalid phone or password", http.StatusUnauthorized)
			return
		}

		session, _ := middleware.Store.Get(r, "church-session")

		session.Values["user_id"] = user.ID
		session.Values["name"] = user.FullName
		session.Values["role"] = user.Role

		println("================================")
		println("Name:", user.FullName)
		println("Role:", user.Role)
		println("================================")

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user.Role == "admin" {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}
