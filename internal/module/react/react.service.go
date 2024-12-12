package react

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IReactService interface {
	Backup(ctx context.Context)
	CreateReact(ctx context.Context, username string, account_id, post_id int64, state int32) (db.ReactPost, error)
	GetReactPost(ctx context.Context, account_id, post_id int64) (ReactResponse, error)
	GetListReactPost(ctx context.Context, page, pageSize int32, post_id int64) (ListReactResponse, error)
	ChangeReactState(ctx context.Context, username string, account_id, post_id int64, state int32) (db.ReactPost, error)
	UnReaction(ctx context.Context, username string, account_id, post_id int64) error
}
