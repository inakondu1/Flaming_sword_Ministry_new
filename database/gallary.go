package database

import (
	"Flaming_Sword_Ministry/models"
)

// ================= CREATE GALLERY TABLE =================

func CreateGalleryTable() {

	query := `
	CREATE TABLE IF NOT EXISTS gallery (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		image TEXT NOT NULL,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
}

// ================= ADD IMAGE =================

func CreateGallery(gallery models.Gallery) error {

	query := `
	INSERT INTO gallery
	(title, image, description)
	VALUES (?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		gallery.Title,
		gallery.Image,
		gallery.Description,
	)

	return err
}

// ================= GET ALL IMAGES =================

func GetAllGallery() ([]models.Gallery, error) {

	rows, err := DB.Query(`
		SELECT
			id,
			title,
			image,
			description,
			created_at
		FROM gallery
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gallery []models.Gallery

	for rows.Next() {

		var img models.Gallery

		err := rows.Scan(
			&img.ID,
			&img.Title,
			&img.Image,
			&img.Description,
			&img.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		gallery = append(gallery, img)
	}

	return gallery, nil
}

// ================= DELETE IMAGE =================

func DeleteGallery(id int) error {

	_, err := DB.Exec(
		"DELETE FROM gallery WHERE id=?",
		id,
	)

	return err
}

// ================= COUNT IMAGES =================

func CountGallery() (int, error) {

	var count int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM gallery",
	).Scan(&count)

	return count, err
}
