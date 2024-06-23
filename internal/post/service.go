package post

import (
	"fmt"
	"strings"
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

type checker interface {
	SQLState() string
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

	if strings.ReplaceAll(post.Author, " ", "") == RESERVED_NAME {
		return fmt.Errorf("that's my name.")
	}

	for i, v := range post.Body {
		if !strings.ContainsRune(ALLOWED_CHARS, v) {
			return fmt.Errorf("body contains invalid characters at index: %d", i)
		}
	}
	for i, v := range post.Author {
		if !strings.ContainsRune(ALLOWED_CHARS, v) {
			return fmt.Errorf("author contains invalid characters at index: %d", i)
		}
	}

	post.ID = uuid.New()
	post.CreatedAt = time.Now().UTC()
	post.Type = POST_MAPPING_NOT_ANALYSED

	err := s.postRepository.InsertPost(post)
	if err != nil {
		pe, ok := err.(checker)
		if !ok {
			return fmt.Errorf("server error")
		}
		s := pe.SQLState()
		if s == SQLSTATE_ERR_NOT_UNIQUE {
			return fmt.Errorf("not a unique post. try changing your post body or name")
		}
	}
	return err
}
