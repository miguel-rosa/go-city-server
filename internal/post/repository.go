package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/miguel-rosa/go-city-server/internal"
)

type Repository interface {
	Insert(ctx context.Context, post internal.Post) (internal.Post, error)
	FindAll(ctx context.Context, params internal.Pagination) ([]internal.Post, error)
	FindOneByID(ctx context.Context, id uuid.UUID) (internal.Post, error)
}

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}

func (r *RepositoryPostgres) Insert(ctx context.Context, post internal.Post) (internal.Post, error) {

	err := r.Conn.QueryRow(
		ctx,
		"INSERT INTO posts (username, body) VALUES ($1, $2) RETURNING id, username, body, created_at",
		post.Username,
		post.Body).Scan(&post.ID, &post.Username, &post.Body, &post.CreatedAt)

	if err != nil {
		return internal.Post{}, err
	}

	return post, nil
}

func (r *RepositoryPostgres) FindAll(ctx context.Context, params internal.Pagination) ([]internal.Post, error) {
	rows, err := r.Conn.Query(
		ctx,
		"SELECT id, username, body, created_at FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2",
		params.Limit,
		params.Offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []internal.Post

	for rows.Next() {
		var item internal.Post
		if err := rows.Scan(&item.ID, &item.Username, &item.Body, &item.CreatedAt); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *RepositoryPostgres) FindOneByID(ctx context.Context, id uuid.UUID) (internal.Post, error) {
	var post = internal.Post{ID: id}

	err := r.Conn.QueryRow(
		ctx,
		"SELECT username, body, created_at FROM posts WHERE id = $1",
		id).Scan(&post.Username, &post.Body, &post.CreatedAt)

	if err == pgx.ErrNoRows {
		return internal.Post{}, ErrPostNotFound
	}

	if err != nil {
		return internal.Post{}, err
	}

	return post, nil
}
