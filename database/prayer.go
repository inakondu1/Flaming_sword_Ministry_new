package database

import "Flaming_Sword_Ministry/models"

// ================= CREATE PRAYER REQUEST =================

// ================= CREATE PRAYER REQUEST =================

func CreatePrayer(prayer models.Prayer) error {

	_, err := DB.Exec(`
		INSERT INTO prayer_requests
		(name, request)
		VALUES (?, ?)
	`,
		prayer.Name,
		prayer.Request,
	)

	return err
}

// ================= GET ALL PRAYER REQUESTS =================

func GetAllPrayers() ([]models.Prayer, error) {

	rows, err := DB.Query(`
		SELECT
			id,
			name,
			request,
			status,
			created_at
		FROM prayer_requests
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prayers []models.Prayer

	for rows.Next() {

		var prayer models.Prayer

		err := rows.Scan(
			&prayer.ID,
			&prayer.Name,
			&prayer.Request,
			&prayer.Status,
			&prayer.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		prayers = append(prayers, prayer)
	}

	return prayers, nil
}

// ================= DELETE PRAYER =================

func DeletePrayer(id int) error {

	_, err := DB.Exec(
		"DELETE FROM prayer_requests WHERE id=?",
		id,
	)

	return err
}
