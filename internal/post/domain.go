package post

import (
	"time"

	"github.com/google/uuid"
)

type PostRepository interface {
	InsertPost(post *Post) error
	SelectRandomPost(postType int) (*Post, error)
}

type PostService interface {
	CreatePost(post *Post) error
}

type Post struct {
	ID        uuid.UUID `json:"id"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
