package model

import (
	"database/sql"
)

type User struct {
	ID   int
	Name string
}

var users []User

func FetchUser(db *sql.DB) (*[]User, error) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}
