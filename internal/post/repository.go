package post

import (
	"database/sql"
)

type postSQLRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postSQLRepository{
		db: db,
	}
}

func (r *postSQLRepository) InsertPost(post *Post) error {
	_, err := r.db.Exec(
		INSERT_POST_QUERY,
		post.ID,
		post.Body,
		post.Author,
		post.Type,
		post.CreatedAt,
	)

	return err
}

func (r *postSQLRepository) SelectRandomPost(postType int) (*Post, error) {
	post := new(Post)

	err := r.db.QueryRow(SELECT_RANDOM_POST_QUERY, postType).Scan(
		&post.ID,
		&post.Body,
		&post.Author,
		&post.Type,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}
