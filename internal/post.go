package internal

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type Pagination struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

var (
	ErrPostBodyExceedsLimit = errors.New("post body exceeds limit")
	ErrIDEmpty              = errors.New("id empty")
	ErrUUIDInvalid          = errors.New("uuid invalid")
)
