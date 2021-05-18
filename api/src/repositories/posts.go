package repositories

import (
	"api/src/models"
	"database/sql"
)

type posts struct {
	db *sql.DB
}

func PostRepository(db *sql.DB) *posts {
	return &posts{db}
}

func (repo posts) Create(post models.Post) (uint64, error) {
	query := "INSERT INTO posts (title, content, author_Id) VALUES (?,?,?)"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}
