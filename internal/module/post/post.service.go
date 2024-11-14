package post

import "context"

type IPostService interface {
	CreatePost(ctx context.Context, description, lng, lat string, images []string, account_id, user_id int64) (PostResponse, error)
	GetPost(ctx context.Context, id int64) (PostResponse, error)
	GetListPost(ctx context.Context, page, pageSize int32) ([]PostResponse, error)
	GetHomePagePost(ctx context.Context, acoount_id int64, page, pageSize int32) ([]PostResponse, error)
	GetPersonPost(ctx context.Context, acoount_id int64, page, pageSize int32) ([]PostResponse, error)
	UpdateContentPost(ctx context.Context, desciption string, id, user_id int64) (PostResponse, error)
	DeletePost(ctx context.Context, id, user_id int64) error
	DeleteImage(ctx context.Context, id, user_id int64) error
	Backup(ctx context.Context)
}
