package middleware

import "net/http"

// AdminOnly allows only logged-in admins.
func AdminOnly(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		session, err := Store.Get(r, "church-session")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		role, ok := session.Values["role"].(string)

		if !ok || role != "admin" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
}
