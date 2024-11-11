// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddImagePost(ctx context.Context, arg AddImagePostParams) (PostImage, error)
	CountComment(ctx context.Context, postTopID pgtype.Int8) (int64, error)
	CountFollower(ctx context.Context, fromFollow int64) (int64, error)
	CountFriend(ctx context.Context, fromFollow int64) (int64, error)
	CountReactPost(ctx context.Context, postID int64) (int64, error)
	CountRequest(ctx context.Context, fromFollow int64) (int64, error)
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
	GetAccountById(ctx context.Context, id int64) (GetAccountByIdRow, error)
	GetAccountByUserId(ctx context.Context, userID int64) ([]int64, error)
	GetComment(ctx context.Context, id int64) (GetCommentRow, error)
	GetDetailAccount(ctx context.Context, id int64) (Account, error)
	GetFavorite(ctx context.Context, arg GetFavoriteParams) ([]int64, error)
	GetFollowStatus(ctx context.Context, arg GetFollowStatusParams) (GetFollowStatusRow, error)
	GetImage(ctx context.Context, id int64) (PostImage, error)
	GetImageComment(ctx context.Context, postID int64) (PostImage, error)
	GetImagePost(ctx context.Context, postID int64) ([]PostImage, error)
	//comment
	GetListComment(ctx context.Context, arg GetListCommentParams) ([]int64, error)
	GetListPost(ctx context.Context, arg GetListPostParams) ([]int64, error)
	GetListReact(ctx context.Context, arg GetListReactParams) ([]GetListReactRow, error)
	GetPersonPost(ctx context.Context, arg GetPersonPostParams) ([]int64, error)
	GetPost(ctx context.Context, id int64) (GetPostRow, error)
	GetReact(ctx context.Context, arg GetReactParams) (ReactPost, error)
	GetYourFollower(ctx context.Context, arg GetYourFollowerParams) ([]int64, error)
	GetYourFriend(ctx context.Context, arg GetYourFriendParams) ([]int64, error)
	GetYourRequest(ctx context.Context, arg GetYourRequestParams) ([]int64, error)
	ListAccountReact(ctx context.Context, postID int64) ([]int64, error)
	Login(ctx context.Context, username string) (LoginRow, error)
	Register(ctx context.Context, arg RegisterParams) (RegisterRow, error)
	SearchingAccounts(ctx context.Context, arg SearchingAccountsParams) ([]SearchingAccountsRow, error)
	UpdateAvatar(ctx context.Context, arg UpdateAvatarParams) (UpdateAvatarRow, error)
	UpdateBackground(ctx context.Context, arg UpdateBackgroundParams) (UpdateBackgroundRow, error)
	UpdateComment(ctx context.Context, arg UpdateCommentParams) (UpdateCommentRow, error)
	UpdateFriend(ctx context.Context, arg UpdateFriendParams) error
	UpdateImagePost(ctx context.Context, arg UpdateImagePostParams) (PostImage, error)
	UpdateName(ctx context.Context, arg UpdateNameParams) (UpdateNameRow, error)
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (UpdatePasswordRow, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (UpdatePostRow, error)
	UpdateState(ctx context.Context, arg UpdateStateParams) (ReactPost, error)
}

var _ Querier = (*Queries)(nil)
