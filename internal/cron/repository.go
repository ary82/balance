package cron

import (
	"database/sql"
	"fmt"

	"github.com/ary82/balance/internal/post"
)

type cronSQLRepository struct {
	db *sql.DB
}

func NewCronRepository(db *sql.DB) CronRepository {
	return &cronSQLRepository{
		db: db,
	}
}

func (r *cronSQLRepository) SelectPosts(post_type int) ([]*post.Post, error) {
	rows, err := r.db.Query(SELECT_MULTIPLE_POSTS_QUERY, post_type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []*post.Post{}
	for rows.Next() {
		post := new(post.Post)
		err := rows.Scan(
			&post.ID,
			&post.Body,
			&post.Author,
			&post.Type,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

const q = "update submissions set type = $%v where id = $%v;"

func (r *cronSQLRepository) UpdateTypesInPosts(posts []*post.Post) error {
	query := ""
	queryArgs := []interface{}{}
	for i, v := range posts {
		query = query + fmt.Sprintf(q, 2*i, 2*i+1)
		queryArgs = append(queryArgs, v.Type)
		queryArgs = append(queryArgs, v.ID)
	}

	_, err := r.db.Exec(query, queryArgs...)
	return err
}
