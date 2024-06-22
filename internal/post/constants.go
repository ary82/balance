package post

const (
	AUTHOR_MIN_LENGTH int = 3
	AUTHOR_MAX_LENGTH int = 25
	BODY_MIN_LENGTH   int = 5
	BODY_MAX_LENGTH   int = 255
)

const (
	POST_MAPPING_NOT_ANALYSED int = 0
	POST_MAPPING_POSITIVE     int = 1
	POST_MAPPING_NEGATIVE     int = 2
	POST_MAPPING_UNCLEAR      int = 3
	POST_MAPPING_RACIAL       int = 4
	POST_MAPPING_SEXUAL       int = 5
	POST_MAPPING_NONE         int = 6
)

const (
	INSERT_POST_QUERY        string = "INSERT INTO submissions(id, body, author, type, created_at) VALUES($1, $2, $3, $4, $5)"
	SELECT_RANDOM_POST_QUERY string = "SELECT id, body, author, type, created_at FROM submissions WHERE type = $1 ORDER BY RANDOM() LIMIT 1"
)

const ALLOWED_CHARS string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789,.'-_ !"

const SQLSTATE_ERR_NOT_UNIQUE string = "23505"
