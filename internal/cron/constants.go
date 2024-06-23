package cron

const (
	SELECT_MULTIPLE_POSTS_QUERY string = "SELECT id, body, author, type, created_at FROM submissions WHERE type = $1"
	SELECT_RANDOM_POST_QUERY    string = "SELECT id, body, author, type, created_at FROM submissions WHERE type = $1 ORDER BY RANDOM() LIMIT 1"
	UPDATE_QUERY                string = "UPDATE submissions SET type = $1 WHERE id = $2;"
	COUNT_POSTS_QUERY           string = "SELECT COUNT(1) FROM submissions WHERE type =$1"
)
