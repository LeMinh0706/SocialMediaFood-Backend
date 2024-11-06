// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	AddImagePost(ctx context.Context, arg AddImagePostParams) (PostImage, error)
	CountReactPost(ctx context.Context, postID int64) (int64, error)
	CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (CreateCommentRow, error)
	CreateFollow(ctx context.Context, arg CreateFollowParams) (Follower, error)
	CreateOwnerBranch(ctx context.Context, arg CreateOwnerBranchParams) (CreateOwnerBranchRow, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (CreatePostRow, error)
	CreateReact(ctx context.Context, arg CreateReactParams) (ReactPost, error)
	DeleteComment(ctx context.Context, id int64) error
	DeleteFollow(ctx context.Context, arg DeleteFollowParams) error
	DeleteImagePost(ctx context.Context, id int64) error
	DeletePost(ctx context.Context, id int64) error
	DeleteReact(ctx context.Context, arg DeleteReactParams) error
	ForgotPassword(ctx context.Context, arg ForgotPasswordParams) (ForgotPasswordRow, error)
	GetAccountById(ctx context.Context, id int64) (Account, error)
	GetAccountByUserId(ctx context.Context, userID int64) ([]GetAccountByUserIdRow, error)
	GetComment(ctx context.Context, id int64) (GetCommentRow, error)
	GetFavorite(ctx context.Context, arg GetFavoriteParams) ([]int64, error)
	GetFollowStatus(ctx context.Context, arg GetFollowStatusParams) (Follower, error)
	GetImage(ctx context.Context, id int64) (PostImage, error)
	GetImagePost(ctx context.Context, postID int64) ([]PostImage, error)
	//comment
	GetListComment(ctx context.Context, arg GetListCommentParams) ([]int64, error)
	GetListPost(ctx context.Context, arg GetListPostParams) ([]int64, error)
	GetPost(ctx context.Context, id int64) (GetPostRow, error)
	GetReact(ctx context.Context, arg GetReactParams) (int64, error)
	GetReactPost(ctx context.Context, arg GetReactPostParams) ([]GetReactPostRow, error)
	GetUserPost(ctx context.Context, arg GetUserPostParams) ([]int64, error)
	Login(ctx context.Context, username string) (LoginRow, error)
	Register(ctx context.Context, arg RegisterParams) (RegisterRow, error)
	UpdateComment(ctx context.Context, arg UpdateCommentParams) (UpdateCommentRow, error)
	UpdateFriend(ctx context.Context, arg UpdateFriendParams) error
	UpdateName(ctx context.Context, arg UpdateNameParams) (UpdateNameRow, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (UpdatePostRow, error)
	UpdateState(ctx context.Context, arg UpdateStateParams) (ReactPost, error)
}

var _ Querier = (*Queries)(nil)
