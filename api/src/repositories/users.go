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
		"SELECT id, name, nick, email FROM users WHERE name LIKE ? OR nick LIKE ?",
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
			&user.Email,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (repo users) GetUserByID(userID uint64) (models.User, error) {
	result, err := repo.db.Query("SELECT id, name, nick, email FROM users WHERE ID = ?", userID)
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
			&user.Email,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil

}

func (repo users) GetUserByEmail(email string) (models.User, error) {
	result, err := repo.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer result.Close()

	var user models.User

	if result.Next() {
		if err = result.Scan(
			&user.ID,
			&user.Password,
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

func (repo users) Delete(userID uint64) error {
	stmt, err := repo.db.Prepare("DELETE FROM users WHERE ID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userID); err != nil {
		return err
	}

	return nil
}

func (repo users) Follow(userID, followerID uint64) error {
	query := "INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?,?)"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}
