package database

import (
	"Flaming_Sword_Ministry/models"
)

// ================= CREATE SERMON =================

func CreateSermon(sermon models.Sermon) error {

	query := `
	INSERT INTO sermons (
		title,
		bible_verse,
		scripture_references,
		content,
		category,
		date,
		created_by
	)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		sermon.Title,
		sermon.BibleVerse,
		sermon.References,
		sermon.Content,
		sermon.Category,
		sermon.Date,
		sermon.CreatedBy,
	)

	return err
}

// ================= GET ALL SERMONS =================

// ================= GET ALL SERMONS =================

func GetAllSermons() ([]models.Sermon, error) {

	rows, err := DB.Query(`
		SELECT
			id,
			title,
			bible_verse,
			scripture_references,
			content,
			category,
			date,
			created_by,
			created_at
		FROM sermons
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sermons []models.Sermon

	for rows.Next() {

		var sermon models.Sermon

		err := rows.Scan(
			&sermon.ID,
			&sermon.Title,
			&sermon.BibleVerse,
			&sermon.References,
			&sermon.Content,
			&sermon.Category,
			&sermon.Date,
			&sermon.CreatedBy,
			&sermon.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		sermons = append(sermons, sermon)
	}

	// Check for iteration errors
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sermons, nil
}

// ================= GET ONE SERMON =================

func GetSermonByID(id int) (models.Sermon, error) {

	var sermon models.Sermon

	query := `
	SELECT
		id,
		title,
		bible_verse,
		scripture_references,
		content,
		category,
		date,
		created_by,
		created_at
	FROM sermons
	WHERE id = ?
	`

	err := DB.QueryRow(query, id).Scan(
		&sermon.ID,
		&sermon.Title,
		&sermon.BibleVerse,
		&sermon.References,
		&sermon.Content,
		&sermon.Category,
		&sermon.Date,
		&sermon.CreatedBy,
		&sermon.CreatedAt,
	)

	return sermon, err
}

// ================= UPDATE SERMON =================

func UpdateSermon(sermon models.Sermon) error {

	query := `
	UPDATE sermons
	SET
		title=?,
		bible_verse=?,
		scripture_references=?,
		content=?,
		category=?,
		date=?,
		created_by=?
	WHERE id=?
	`

	_, err := DB.Exec(
		query,
		sermon.Title,
		sermon.BibleVerse,
		sermon.References,
		sermon.Content,
		sermon.Category,
		sermon.Date,
		sermon.CreatedBy,
		sermon.ID,
	)

	return err
}

// ================= DELETE SERMON =================

func DeleteSermon(id int) error {

	_, err := DB.Exec(
		"DELETE FROM sermons WHERE id=?",
		id,
	)

	return err
}

// ================= COUNT SERMONS =================

func CountSermons() (int, error) {

	var count int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM sermons",
	).Scan(&count)

	return count, err
}
