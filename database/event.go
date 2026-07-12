package database

import "Flaming_Sword_Ministry/models"

func CreateEvent(event models.Event) error {

	_, err := DB.Exec(`
	INSERT INTO events
	(title, description, event_date, event_time, venue)
	VALUES (?, ?, ?, ?, ?)
	`,
		event.Title,
		event.Description,
		event.EventDate,
		event.EventTime,
		event.Venue,
	)

	return err
}

func GetAllEvents() ([]models.Event, error) {

	rows, err := DB.Query(`
	SELECT
		id,
		title,
		description,
		event_date,
		event_time,
		venue,
		created_at
	FROM events
	ORDER BY event_date ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {

		var e models.Event

		err := rows.Scan(
			&e.ID,
			&e.Title,
			&e.Description,
			&e.EventDate,
			&e.EventTime,
			&e.Venue,
			&e.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func DeleteEvent(id int) error {

	_, err := DB.Exec(
		"DELETE FROM events WHERE id=?",
		id,
	)

	return err
}
