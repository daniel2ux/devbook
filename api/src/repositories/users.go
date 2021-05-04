package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repo users) Create(user models.User) (uint64, error) {
	query := "INSERT INTO users (name,nick,email,password) VALUES (?,?,?,?)"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

func (repo users) GetUsers(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	result, err := repo.db.Query(
		"SELECT id, name, nick, created_at FROM users WHERE name LIKE ? OR nick LIKE ?",
		nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var users []models.User

	for result.Next() {
		var user models.User
		if err = result.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Created_at,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (repo users) GetUserByID(userID uint64) (models.User, error) {
	result, err := repo.db.Query("SELECT id, name, nick, created_at FROM users WHERE ID = ?", userID)
	if err != nil {
		return models.User{}, err
	}
	defer result.Close()

	var user models.User

	if result.Next() {
		if err = result.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Created_at,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil

}

func (repo users) Update(userID uint64, user models.User) error {
	stmt, err := repo.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE ID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Name, user.Nick, user.Email, userID); err != nil {
		return err
	}

	return nil

}