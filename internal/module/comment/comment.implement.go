package comment

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
	"github.com/jackc/pgx/v5/pgtype"
)

type CommentService struct {
	queries *db.Queries
	post    post.IPostService
	account account.IAccountService
}

// Backup implements ICommentService.
func (c *CommentService) Backup(ctx context.Context) {
	panic("unimplemented")
}

// CreateComment implements ICommentService.
func (c *CommentService) CreateComment(ctx context.Context, account_id int64, user_id int64, post_top_id int64, description string, image string) (CommentResponse, error) {
	var res CommentResponse
	_, err := c.post.GetPost(ctx, account_id, post_top_id)
	if err != nil {
		return res, err
	}
	acc, err := c.account.GetAccountAction(ctx, account_id, user_id)
	if err != nil {
		return res, err
	}
	comment, err := c.queries.CreateComment(ctx, db.CreateCommentParams{
		AccountID:   account_id,
		PostTopID:   pgtype.Int8{Int64: post_top_id, Valid: true},
		Description: post.ConvertDescription(description),
	})
	if err != nil {
		return res, err
	}
	var img db.PostImage
	if image != "" {
		img, err = c.queries.AddImagePost(ctx, db.AddImagePostParams{
			UrlImage: image,
			PostID:   comment.ID,
		})
		if err != nil {
			return res, err
		}
	}
	res = CommentRes(comment, img, acc)
	return res, nil
}

// DeleteComment implements ICommentService.
func (c *CommentService) DeleteComment(ctx context.Context, id int64, user_id int64) error {
	comment, err := c.GetComment(ctx, id)
	if err != nil {
		return err
	}
	_, err = c.account.GetAccountAction(ctx, comment.AccountID, user_id)
	if err != nil {
		return err
	}
	err = c.queries.DeleteImageComment(ctx, id)
	if err != nil {
		return err
	}
	err = c.queries.DeleteComment(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// GetComment implements ICommentService.
func (c *CommentService) GetComment(ctx context.Context, id int64) (CommentResponse, error) {
	var res CommentResponse
	comment, err := c.queries.GetComment(ctx, id)
	if err != nil {
		return res, err
	}
	_, err = c.post.GetPost(ctx, 0, comment.PostTopID.Int64)
	if err != nil {
		return res, err
	}
	var image db.PostImage
	account, _ := c.account.GetAccountById(ctx, comment.AccountID)
	image, _ = c.queries.GetImageComment(ctx, comment.ID)
	res = CommentRes(db.CreateCommentRow(comment), image, account)
	return res, nil
}

// GetListComment implements ICommentService.
func (c *CommentService) GetListComment(ctx context.Context, page int32, pageSize int32, post_top_id int64) ([]CommentResponse, error) {
	var res []CommentResponse
	list, err := c.queries.GetListComment(ctx, db.GetListCommentParams{
		PostTopID: pgtype.Int8{Int64: post_top_id, Valid: true},
		Limit:     pageSize,
		Offset:    (page - 1) * pageSize,
	})
	if err != nil {
		return []CommentResponse{}, err
	}
	for _, element := range list {
		comment, err := c.GetComment(ctx, element)
		if err != nil {
			return nil, err
		}
		res = append(res, comment)
	}
	return res, nil
}

// UpdateComment implements ICommentService.
func (c *CommentService) UpdateComment(ctx context.Context, user_id int64, id int64, description string, image string) (CommentResponse, error) {
	var res CommentResponse
	comment, err := c.GetComment(ctx, id)
	if err != nil {
		return res, err
	}
	acc, err := c.account.GetAccountAction(ctx, comment.AccountID, user_id)
	if err != nil {
		return res, err
	}
	var img db.PostImage
	if image != "" {
		img, err = c.queries.AddImagePost(ctx, db.AddImagePostParams{
			UrlImage: image,
			PostID:   id,
		})
		if err != nil {
			return res, err
		}
	}
	update, err := c.queries.UpdateComment(ctx, db.UpdateCommentParams{
		ID:          id,
		Description: post.ConvertDescription(description),
	})
	if err != nil {
		return res, err
	}
	res = CommentRes(db.CreateCommentRow(update), img, acc)

	return res, nil
}

func NewCommentService(queries *db.Queries, ps post.IPostService, as account.IAccountService) ICommentService {
	return &CommentService{
		queries: queries,
		account: as,
		post:    ps,
	}
}
