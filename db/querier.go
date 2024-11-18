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
	AddToMenu(ctx context.Context, arg AddToMenuParams) (Menu, error)
	CountComment(ctx context.Context, postTopID pgtype.Int8) (int64, error)
	CountFollower(ctx context.Context, fromFollow int64) (int64, error)
	CountFriend(ctx context.Context, fromFollow int64) (int64, error)
	CountReactPost(ctx context.Context, postID int64) (int64, error)
	CountRequest(ctx context.Context, fromFollow int64) (int64, error)
	CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error)
	CreateActionNoti(ctx context.Context, arg CreateActionNotiParams) (Notification, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (CreateCommentRow, error)
	CreateFollow(ctx context.Context, arg CreateFollowParams) (Follower, error)
	CreateOwnerBranch(ctx context.Context, arg CreateOwnerBranchParams) (CreateOwnerBranchRow, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (CreatePostRow, error)
	CreatePostNoti(ctx context.Context, arg CreatePostNotiParams) (Notification, error)
	CreateReact(ctx context.Context, arg CreateReactParams) (ReactPost, error)
	CreateReport(ctx context.Context, arg CreateReportParams) (ReportPost, error)
	DeleteComment(ctx context.Context, id int64) error
	DeleteFollow(ctx context.Context, arg DeleteFollowParams) error
	DeleteImageComment(ctx context.Context, postID int64) error
	DeleteImagePost(ctx context.Context, id int64) error
	DeleteNoti(ctx context.Context, id int64) error
	DeletePost(ctx context.Context, id int64) error
	DeleteReact(ctx context.Context, arg DeleteReactParams) error
	GetAccountById(ctx context.Context, id int64) (GetAccountByIdRow, error)
	GetAccountByUserId(ctx context.Context, userID int64) ([]int64, error)
	GetCheckAction(ctx context.Context, userID int64) (ResetPassword, error)
	GetComment(ctx context.Context, id int64) (GetCommentRow, error)
	GetDetailAccount(ctx context.Context, id int64) (Account, error)
	GetDish(ctx context.Context, accountID pgtype.Int8) (GetDishRow, error)
	GetFavorite(ctx context.Context, arg GetFavoriteParams) ([]int64, error)
	GetFollowStatus(ctx context.Context, arg GetFollowStatusParams) (GetFollowStatusRow, error)
	GetHomePagePost(ctx context.Context, arg GetHomePagePostParams) ([]int64, error)
	GetImage(ctx context.Context, id int64) (PostImage, error)
	GetImageComment(ctx context.Context, postID int64) (PostImage, error)
	GetImagePost(ctx context.Context, postID int64) ([]PostImage, error)
	GetIssue(ctx context.Context, id int32) (IssuePost, error)
	//comment
	GetListComment(ctx context.Context, arg GetListCommentParams) ([]int64, error)
	GetListIssue(ctx context.Context) ([]IssuePost, error)
	GetListNoti(ctx context.Context, arg GetListNotiParams) ([]int64, error)
	GetListPost(ctx context.Context, arg GetListPostParams) ([]int64, error)
	GetListReact(ctx context.Context, arg GetListReactParams) ([]int64, error)
	GetLocation(ctx context.Context, accountID int64) ([]GetLocationRow, error)
	GetMenu(ctx context.Context, arg GetMenuParams) ([]GetMenuRow, error)
	GetNotification(ctx context.Context, accountID int64) (Notification, error)
	GetPersonPost(ctx context.Context, arg GetPersonPostParams) ([]int64, error)
	GetPost(ctx context.Context, id int64) (GetPostRow, error)
	GetReact(ctx context.Context, arg GetReactParams) (ReactPost, error)
	GetRequestByUUID(ctx context.Context, id pgtype.UUID) (ResetPassword, error)
	GetUserByEmail(ctx context.Context, email pgtype.Text) (GetUserByEmailRow, error)
	GetYourFollower(ctx context.Context, arg GetYourFollowerParams) ([]int64, error)
	GetYourFriend(ctx context.Context, arg GetYourFriendParams) ([]int64, error)
	GetYourReport(ctx context.Context, arg GetYourReportParams) ([]GetYourReportRow, error)
	GetYourRequest(ctx context.Context, arg GetYourRequestParams) ([]int64, error)
	ListAccountReact(ctx context.Context, postID int64) ([]int64, error)
	Login(ctx context.Context, username string) (LoginRow, error)
	OwnerUpdateQuanity(ctx context.Context, arg OwnerUpdateQuanityParams) (Menu, error)
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
	UpdateQuanity(ctx context.Context, arg UpdateQuanityParams) (Menu, error)
	UpdateSeen(ctx context.Context, id int64) error
	UpdateSeenAll(ctx context.Context, accountID int64) error
	UpdateState(ctx context.Context, arg UpdateStateParams) (ReactPost, error)
	UpgradeSuccess(ctx context.Context, id int64) (int64, error)
}

var _ Querier = (*Queries)(nil)
