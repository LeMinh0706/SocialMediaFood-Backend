// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
)

type Post struct {
	ID             int64          `json:"id"`
	PostTypeID     int64          `json:"post_type_id"`
	UserID         int64          `json:"user_id"`
	PostTopID      sql.NullInt64  `json:"post_top_id"`
	Description    sql.NullString `json:"description"`
	DateCreatePost int64          `json:"date_create_post"`
}
