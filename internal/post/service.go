package post

import (
	"context"
	"errors"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/miguel-rosa/go-city-server/internal"
)

var ErrPostBodyEmpty = errors.New("post body is empty")
var ErrPostUsernameEmpty = errors.New("post username empty")
var ErrPostBodyExceedsLimit = errors.New("post body exceeds limit")
var ErrPostNotFound = errors.New("post not found")
var ErrIdEmpty = errors.New("id empty")
var ErrUUIDInvalid = errors.New("uuid invalid")

type Service struct {
	Repository Repository
}

func (p Service) Create(ctx context.Context, post internal.Post) (internal.Post, error) {
	if post.Body == "" {
		return internal.Post{}, ErrPostBodyEmpty
	}

	if post.Username == "" {
		return internal.Post{}, ErrPostUsernameEmpty
	}

	if utf8.RuneCountInString(post.Body) > 140 {
		return internal.Post{}, ErrPostBodyExceedsLimit
	}

	return p.Repository.Insert(ctx, post)
}

func (s Service) FindOneByID(ctx context.Context, id uuid.UUID) (internal.Post, error) {
	return s.Repository.FindOneByID(ctx, id)
}

func (s Service) FindAll(ctx context.Context) ([]internal.Post, error) {
	posts, err := s.Repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
