package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {

	var err error

	DB, err = sql.Open("sqlite3", "./church.db")
	if err != nil {
		log.Fatal(err)
	}

	createUsersTable()
	createSermonsTable()
	createAnnouncementsTable()
	createPrayerTable()

	log.Println("✅ Database connected successfully.")
}

// ================= USERS =================

func createUsersTable() {

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fullname TEXT NOT NULL,
		phone TEXT NOT NULL UNIQUE,
		gender TEXT NOT NULL,
		password TEXT NOT NULL,
		role TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Users table ready.")
}

// ================= SERMONS =================

func createSermonsTable() {

	query := `
	CREATE TABLE IF NOT EXISTS sermons (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		bible_verse TEXT NOT NULL,
		scripture_references TEXT,
		content TEXT NOT NULL,
		category TEXT,
		date TEXT,
		created_by TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Sermons table ready.")
}

// ================= ANNOUNCEMENTS =================

func createAnnouncementsTable() {

	query := `
	CREATE TABLE IF NOT EXISTS announcements (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		message TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Announcements table ready.")
}
func createPrayerTable() {

	query := `
	CREATE TABLE IF NOT EXISTS prayer_requests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		phone TEXT,
		request TEXT NOT NULL,
		status TEXT DEFAULT 'Pending',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Prayer Requests table ready.")
}
