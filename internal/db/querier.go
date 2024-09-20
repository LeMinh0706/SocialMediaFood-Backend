// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateComment(ctx context.Context, arg CreateCommentParams) (Post, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	DeletePost(ctx context.Context, id int64) error
	GetPost(ctx context.Context, id int64) (Post, error)
	ListPost(ctx context.Context, arg ListPostParams) ([]Post, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
}

var _ Querier = (*Queries)(nil)