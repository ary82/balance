package database

import (
	"log"
	"time"

	"github.com/google/uuid"
)

func (s *service) PostSubmission(post *SubmissionReq) error {
	query := `
  INSERT INTO submissions(
  id, body, author,
  type, created_at)
  VALUES($1, $2, $3, $4, $5)
  `

	id := uuid.New()
	currentTime := time.Now()

	res, err := s.db.Exec(query, id, post.Body, post.Author, post.Type, currentTime)

	log.Println(res)
	return err
}

func (s *service) GetRandom(submissionType int) (*Submission, error) {
	query := `
  SELECT id, body, author, type, created_at FROM submissions
  WHERE type = $1
  ORDER BY RANDOM() LIMIT 1;
  `

	post := new(Submission)

	err := s.db.QueryRow(query, submissionType).Scan(
		&post.ID,
		&post.Body,
		&post.Author,
		&post.Type,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}
