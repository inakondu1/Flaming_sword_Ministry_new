package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/models"
)

// ================= VIEW EVENTS =================

func ViewEventsHandler(w http.ResponseWriter, r *http.Request) {

	events, err := database.GetAllEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/events.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ================= ADD EVENT =================

func AddEventHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/add_event.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	event := models.Event{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		EventDate:   r.FormValue("event_date"),
		EventTime:   r.FormValue("event_time"),
		Venue:       r.FormValue("venue"),
	}

	err := database.CreateEvent(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/events", http.StatusSeeOther)
}

// ================= DELETE EVENT =================

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteEvent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/events", http.StatusSeeOther)
}
