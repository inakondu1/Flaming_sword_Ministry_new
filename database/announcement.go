package database

import (
	"Flaming_Sword_Ministry/models"
)

// CreateAnnouncement inserts a new announcement.
func CreateAnnouncement(title, message string) error {

	_, err := DB.Exec(
		`INSERT INTO announcements (title, message) VALUES (?, ?)`,
		title,
		message,
	)

	return err
}

// GetAllAnnouncements returns all announcements.
func GetAllAnnouncements() ([]models.Announcement, error) {

	rows, err := DB.Query(`
		SELECT id, title, message
		FROM announcements
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var announcements []models.Announcement

	for rows.Next() {

		var a models.Announcement

		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Message,
		)
		if err != nil {
			return nil, err
		}

		announcements = append(announcements, a)
	}

	// Check for errors after iterating
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return announcements, nil
}

// CountAnnouncements returns the total number of announcements.
func CountAnnouncements() (int, error) {

	var count int

	err := DB.QueryRow(
		`SELECT COUNT(*) FROM announcements`,
	).Scan(&count)

	return count, err
}
