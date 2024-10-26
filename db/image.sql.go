// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: image.sql

package db

import (
	"context"
)

const addImagePost = `-- name: AddImagePost :one
INSERT INTO post_image (
    url_image,
    post_id
) VALUES (
    $1, $2
) RETURNING id, url_image, post_id
`

type AddImagePostParams struct {
	UrlImage string `json:"url_image"`
	PostID   int64  `json:"post_id"`
}

func (q *Queries) AddImagePost(ctx context.Context, arg AddImagePostParams) (PostImage, error) {
	row := q.db.QueryRow(ctx, addImagePost, arg.UrlImage, arg.PostID)
	var i PostImage
	err := row.Scan(&i.ID, &i.UrlImage, &i.PostID)
	return i, err
}
