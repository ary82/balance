package database

import (
	"time"

	"github.com/google/uuid"
)

type Submission struct {
	ID        uuid.UUID `json:"id"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type SubmissionReq struct {
	Body   string `json:"body"`
	Author string `json:"author"`
	Type   int    `json:"type"`
}
