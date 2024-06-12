package post

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type postService struct {
	postRepository PostRepository
}

func CreatePostService(repo PostRepository) PostService {
	return &postService{
		postRepository: repo,
	}
}

func (s *postService) CreatePost(post *Post) error {
	if len(post.Body) > BODY_MAX_LENGTH {
		return fmt.Errorf("post too long")
	}
	if len(post.Author) > AUTHOR_MAX_LENGTH {
		return fmt.Errorf("author name too long")
	}
	if len(post.Body) < BODY_MIN_LENGTH {
		return fmt.Errorf("post too short")
	}
	if len(post.Author) < AUTHOR_MIN_LENGTH {
		return fmt.Errorf("author name too short")
	}

	post.ID = uuid.New()
	post.CreatedAt = time.Now()
	post.Type = POST_MAPPING_NOT_ANALYSED

	err := s.postRepository.InsertPost(post)
	return err
}

func (s *postService) GetRandomPost(postType int) (*Post, error) {
	if postType != POST_MAPPING_POSITIVE && postType != POST_MAPPING_NEGATIVE {
		return nil, fmt.Errorf("invalid postType")
	}
	post, err := s.postRepository.SelectRandomPost(postType)
	if err != nil {
		return nil, err
	}
	return post, nil
}