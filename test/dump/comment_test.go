package dump

import (
	"context"
	"testing"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateComment(t *testing.T) {
	user := createRandomUser(t)
	acc := createRandomAccount(t, user.ID, 3)
	post := createPostNoPoint(t)
	description := util.RandomDescription()
	arg := db.CreateCommentParams{
		AccountID:   acc.ID,
		PostTopID:   pgtype.Int8{Int64: post.ID, Valid: true},
		Description: pgtype.Text{String: description, Valid: true},
	}
	comment, err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, acc.ID, comment.AccountID)
	require.Equal(t, post.ID, comment.PostTopID.Int64)
	require.Equal(t, description, comment.Description.String)

	require.NotZero(t, comment.ID)
	require.NotZero(t, comment.CreatedAt)
}
