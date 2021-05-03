package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) *users {
	return &users{db}
}

func (ur users) Create(user models.User) (uint64, error) {
	query := "INSERT INTO users (name,nick,email,password) VALUES (?,?,?,?)"
	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	r, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}
