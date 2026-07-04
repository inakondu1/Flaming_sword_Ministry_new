package main

import (
	"log"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/handlers"
	"Flaming_Sword_Ministry/middleware"
)

func main() {

	// Connect Database
	database.ConnectDB()

	// Static Files
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	
	// Public Routes
	// =========================

	http.HandleFunc("/", handlers.WelcomeHandler)
	http.HandleFunc("/home", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)

	http.HandleFunc("/sermons", handlers.ViewSermonsHandler)
	http.HandleFunc("/sermon", handlers.ViewSingleSermonHandler)

	http.HandleFunc("/announcement", handlers.ViewAnnouncementsHandler)

	http.HandleFunc("/prayer", handlers.PrayerHandler)

	http.HandleFunc("/events", handlers.ViewEventsHandler)

	// Admin Routes
	http.HandleFunc("/admin",
		middleware.AdminOnly(handlers.AdminHandler),
	)
	http.HandleFunc("/admin/add-sermon",
		middleware.AdminOnly(handlers.AddSermonHandler),
	)

	http.HandleFunc("/admin/edit-sermon",
		middleware.AdminOnly(handlers.EditSermonHandler),
	)

	http.HandleFunc("/admin/users",
		middleware.AdminOnly(handlers.ViewMembersHandler),
	)

	http.HandleFunc("/admin/delete-sermon",
		middleware.AdminOnly(handlers.DeleteSermonHandler),
	)

	http.HandleFunc("/admin/delete-prayer",
		middleware.AdminOnly(handlers.DeletePrayerHandler),
	)
	http.HandleFunc("/admin/add-event",
		middleware.AdminOnly(handlers.AddEventHandler),
	)

	http.HandleFunc("/admin/delete-event",
		middleware.AdminOnly(handlers.DeleteEventHandler),
	)
	http.HandleFunc("/admin/add-announcement",
		middleware.AdminOnly(handlers.CreateAnnouncementHandler),
	)
	log.Println("🚀 Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
