package cron

import "github.com/ary82/balance/internal/post"

type CronRepository interface {
	SelectPosts(post_type int) ([]*post.Post, error)
	UpdateTypesInPosts(posts []*post.Post) error
}

type CronService interface {
	Start() error
}
