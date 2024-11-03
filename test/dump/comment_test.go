package dump

import (
	"context"
	"testing"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomComment(t *testing.T, post_id int64) db.CreateCommentRow {
	acc := createRandomRegister(t)
	description := util.RandomDescription()
	arg := db.CreateCommentParams{
		AccountID:   acc.ID,
		PostTopID:   pgtype.Int8{Int64: post_id, Valid: true},
		Description: pgtype.Text{String: description, Valid: true},
	}
	comment, err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, acc.ID, comment.AccountID)
	require.Equal(t, post_id, comment.PostTopID.Int64)
	require.Equal(t, description, comment.Description.String)

	require.NotZero(t, comment.ID)
	return comment
}

func TestCreateComment(t *testing.T) {
	post := createPostNoPoint(t)
	createRandomComment(t, post.ID)
}

func TestListComment(t *testing.T) {
	post := createPostNoPoint(t)
	for i := 0; i < 10; i++ {
		createRandomComment(t, post.ID)
	}

	comments, err := testQueries.GetListComment(context.Background(), db.GetListCommentParams{
		PostTopID: pgtype.Int8{Int64: post.ID, Valid: true},
		Limit:     5,
		Offset:    5,
	})
	require.NoError(t, err)
	require.NotEmpty(t, comments)

	require.Len(t, comments, 5)
}

func TestUpdateComment(t *testing.T) {
	post := createPostImage(t)
	comment := createRandomComment(t, post.ID)

	update, err := testQueries.UpdateComment(context.Background(), db.UpdateCommentParams{
		ID:          comment.ID,
		Description: pgtype.Text{String: util.RandomDescription(), Valid: true},
	})

	require.NotEmpty(t, update)
	require.NoError(t, err)
}

func TestDeleteComment(t *testing.T) {
	post := createPostImage(t)
	comment := createRandomComment(t, post.ID)

	err := testQueries.DeleteComment(context.Background(), comment.ID)
	require.NoError(t, err)
	get, err := testQueries.GetComment(context.Background(), comment.ID)
	require.Error(t, err)
	require.Empty(t, get)
}
