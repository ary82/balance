package cron

import "github.com/ary82/balance/internal/post"

type CronRepository interface {
	SelectPosts(post_type int) ([]*post.Post, error)
	UpdateTypesInPosts(posts []*post.Post) error
	SelectRandomPost(postType int) (*post.Post, error)
	CountPosts(postType int) (int, error)
}

type CronService interface {
	Start() error
}
