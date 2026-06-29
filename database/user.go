package database

import (
	"Flaming_Sword_Ministry/models"
)

// ================= CREATE USER =================

func CreateUser(user models.User) error {

	query := `
	INSERT INTO users
	(fullname, phone, gender, password, role)
	VALUES (?, ?, ?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		user.FullName,
		user.Phone,
		user.Gender,
		user.Password,
		user.Role,
	)

	return err
}

// ================= LOGIN =================

func GetUserByPhone(phone string) (models.User, error) {

	var user models.User

	query := `
	SELECT
		id,
		fullname,
		phone,
		gender,
		password,
		role
	FROM users
	WHERE phone = ?
	`

	err := DB.QueryRow(query, phone).Scan(
		&user.ID,
		&user.FullName,
		&user.Phone,
		&user.Gender,
		&user.Password,
		&user.Role,
	)

	return user, err
}

// ================= ALL USERS =================

func GetAllUsers() ([]models.User, error) {

	rows, err := DB.Query(`
		SELECT
			id,
			fullname,
			phone,
			gender,
			role,
			created_at
		FROM users
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Phone,
			&user.Gender,
			&user.Role,
			&user.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// ================= COUNT USERS =================

func CountUsers() (int, error) {

	var count int

	err := DB.QueryRow(
		"SELECT COUNT(*) FROM users",
	).Scan(&count)

	return count, err
}