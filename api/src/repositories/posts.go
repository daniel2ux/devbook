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

func (repo posts) GetPosts(userID uint64) ([]models.Post, error) {
	result, err := repo.db.Query(`
		SELECT DISTINCT p.*, u.nick
		  FROM posts p INNER JOIN users u
		    		      ON u.id = p.author_id
					   INNER JOIN followers f
					      ON p.author_id = f.user_id
		 WHERE u.id = ?
		    OR f.follower_id = ?
		  ORDER BY 1 DESC					
	`, userID, userID)

	if err != nil {
		return nil, err
	}

	defer result.Close()
	var posts []models.Post

	for result.Next() {
		var post models.Post
		if err = result.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreateAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo posts) Update(postID uint64, post models.Post) error {
	stmt, err := repo.db.Prepare("UPDATE posts SET title=?, content=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

func (repo posts) Delete(postID uint64) error {
	stmt, err := repo.db.Prepare("DELETE FROM posts WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repo posts) GetPostsByUser(userID uint64) ([]models.Post, error) {
	result, err := repo.db.Query(`
		SELECT p.*, u.nick
		  FROM posts p INNER JOIN users u
		    		      ON u.id = p.author_id
		 WHERE p.author_id = ?		
	`, userID)

	if err != nil {
		return nil, err
	}

	defer result.Close()
	var posts []models.Post

	for result.Next() {
		var post models.Post
		if err = result.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreateAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo posts) Like(postID uint64) error {
	stmt, err := repo.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repo posts) Dislike(postID uint64) error {
	stmt, err := repo.db.Prepare(`
		UPDATE posts SET likes = 
			CASE WHEN likes > 0 THEN likes - 1 
			ELSE likes END 
		 WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}
