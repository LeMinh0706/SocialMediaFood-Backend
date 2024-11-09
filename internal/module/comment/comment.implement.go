package comment

import "context"

type ICommentService interface {
	GetComment(ctx context.Context, id int64) (CommentResponse, error)
	GetListComment(ctx context.Context, page, pageSize int32, post_top_id int64) ([]CommentResponse, error)
	CreateComment(ctx context.Context, account_id, user_id, post_top_id int64, description, image string) (CommentResponse, error)
	UpdateComment(ctx context.Context, user_id, id int64, description, image string) (CommentResponse, error)
	DeleteComment(ctx context.Context, id, user_id int64) error
	Backup(ctx context.Context)
}
