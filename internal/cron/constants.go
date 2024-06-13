package cron

const SELECT_MULTIPLE_POSTS_QUERY string = "select id, body, author, type, created_at from submissions where type = $1"
