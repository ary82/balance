package cron

const (
	SELECT_MULTIPLE_POSTS_QUERY string = "select id, body, author, type, created_at from submissions where type = $1"
	UPDATE_TYPES_QUERY                 = "update submissions set type = $%v where id = $%v;"
	SELECT_RANDOM_POST_QUERY    string = "SELECT id, body, author, type, created_at FROM submissions WHERE type = $1 ORDER BY RANDOM() LIMIT 1"
)
