package comment

import "context"

type ICommentService interface {
	GetComment(ctx context.Context, id int64) (CommentResponse, error)
	GetListComment(ctx context.Context, page, pageSize int32, post_top_id int64) ([]CommentResponse, error)
	CreateComment(ctx context.Context, username string, account_id, post_top_id int64, description, image string) (CommentResponse, error)
	UpdateComment(ctx context.Context, id int64, username, description, image string) (CommentResponse, error)
	DeleteComment(ctx context.Context, id int64, username string) error
	Backup(ctx context.Context)
}
