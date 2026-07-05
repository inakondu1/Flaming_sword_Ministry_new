package database

import (
	"Flaming_Sword_Ministry/models"
)

// ================= CREATE CONTACT TABLE =================

func CreateContactTable() {

	query := `
	CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fullname TEXT NOT NULL,
		phone TEXT,
		subject TEXT NOT NULL,
		message TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
}

// ================= SAVE CONTACT MESSAGE =================

func CreateContact(contact models.Contact) error {

	query := `
	INSERT INTO contacts
	(fullname, email, phone, subject, message)
	VALUES (?, ?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		contact.FullName,
		contact.Phone,
		contact.Subject,
		contact.Message,
	)

	return err
}

// ================= GET ALL CONTACTS =================

func GetAllContacts() ([]models.Contact, error) {

	rows, err := DB.Query(`
		SELECT
			id,
			fullname,
			phone,
			subject,
			message,
			created_at
		FROM contacts
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact

	for rows.Next() {

		var contact models.Contact

		err := rows.Scan(
			&contact.ID,
			&contact.FullName,
			&contact.Phone,
			&contact.Subject,
			&contact.Message,
			&contact.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}

// ================= DELETE CONTACT =================

func DeleteContact(id int) error {

	_, err := DB.Exec(
		"DELETE FROM contacts WHERE id=?",
		id,
	)

	return err
}

// ================= COUNT CONTACTS =================

func CountContacts() (int, error) {

	var count int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM contacts",
	).Scan(&count)

	return count, err
}
