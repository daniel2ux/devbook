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
		"SELECT id, name, nick, email, created_at FROM users WHERE name LIKE ? OR nick LIKE ?",
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
			&user.CreateAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo users) GetUserByID(userID uint64) (models.User, error) {
	result, err := repo.db.Query("SELECT id, name, nick, email, created_at FROM users WHERE ID = ?", userID)
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
			&user.CreateAt,
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

func (repo users) StopFollow(userID, followerID uint64) error {
	query := "DELETE FROM followers WHERE user_id = ? AND follower_id = ?"
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

func (repo users) GetFollowers(userID uint64) ([]models.User, error) {
	query := `SELECT u.id, u.name, u.nick, u.email
	            FROM users u INNER JOIN followers f on u.id = f.follower_id
			   WHERE f.user_id = ?`
	result, err := repo.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var followers []models.User

	for result.Next() {
		var follower models.User
		if err = result.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (repo users) GetFollows(userID uint64) ([]models.User, error) {
	query := `SELECT u.id, u.name, u.nick, u.email
	            FROM users u INNER JOIN followers f on u.id = f.user_id
			   WHERE f.follower_id = ?`
	result, err := repo.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var follows []models.User

	for result.Next() {
		var follow models.User
		if err = result.Scan(
			&follow.ID,
			&follow.Name,
			&follow.Nick,
			&follow.Email,
		); err != nil {
			return nil, err
		}

		follows = append(follows, follow)
	}

	return follows, nil
}

func (repo users) GetPassByID(userID uint64) (string, error) {
	query := `SELECT password FROM users WHERE id = ?`
	result, err := repo.db.Query(query, userID)
	if err != nil {
		return "", err
	}
	defer result.Close()

	var user models.User

	if result.Next() {
		if err = result.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repo users) UpdatePassword(userID uint64, password string) error {
	stmt, err := repo.db.Prepare("UPDATE users SET password = ? WHERE ID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(password, userID); err != nil {
		return err
	}

	return nil
}
