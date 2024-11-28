package post

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IPostService interface {
	CreatePost(ctx context.Context, description, lng, lat string, images []string, account_id, user_id int64) (PostResponse, error)
	GetPost(ctx context.Context, account_id, id int64) (PostResponse, error)
	GetListPost(ctx context.Context, page, pageSize int32) ([]PostResponse, error)
	GetPostInLocate(ctx context.Context, dwithin, account_id int64, lng, lat string, page, pageSize int32) ([]PostResponse, error)
	GetHomePagePost(ctx context.Context, acoount_id int64, page, pageSize int32) ([]PostResponse, error)
	GetPersonPost(ctx context.Context, from, to int64, page, pageSize int32) ([]PostResponse, error)
	UpdateContentPost(ctx context.Context, desciption string, id, user_id int64, images []string) (PostResponse, error)
	GetImage(ctx context.Context, id int64) ([]db.PostImage, error)
	DeletePost(ctx context.Context, id, user_id int64) error
	DeleteImage(ctx context.Context, id, user_id int64) error
	GetListImage(ctx context.Context, page, pageSize int32) ([]db.PostImage, error)
	Backup(ctx context.Context)
}
