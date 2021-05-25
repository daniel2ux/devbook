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

func (repo posts) GetByID(postID uint64) (models.Post, error) {
	result, err := repo.db.Query(`
		SELECT p.*, u.nick
		  FROM posts p INNER JOIN users u
		    ON u.id = p.author_Id
		 WHERE p.id = ?			
	`, postID)

	if err != nil {
		return models.Post{}, err
	}

	defer result.Close()
	var post models.Post
	if result.Next() {
		if err = result.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreateAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}
