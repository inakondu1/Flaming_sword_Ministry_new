package handlers

import (
	"net/http"

	"Flaming_Sword_Ministry/middleware"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := middleware.Store.Get(r, "church-session")

	session.Options.MaxAge = -1

	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
