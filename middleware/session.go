package middleware

import "github.com/gorilla/sessions"

// Cookie store for user sessions
var Store = sessions.NewCookieStore(
	[]byte("flaming-sword-secret-key-2026"),
)

func init() {
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
	}
}