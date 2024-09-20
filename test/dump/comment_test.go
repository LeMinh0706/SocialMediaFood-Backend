package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomComment(t *testing.T) db.Post {
	arg := db.CreateCommentParams{
		UserID:         util.RandomInt(1, 10),
		PostTopID:      sql.NullInt64{Int64: util.RandomInt(1, 10), Valid: true},
		Description:    sql.NullString{String: util.RandomDescription(), Valid: true},
		DateCreatePost: time.Now().Unix(),
	}
	post, err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.Description, post.Description)
	require.Equal(t, arg.PostTopID, post.PostTopID)
	require.Equal(t, arg.UserID, post.UserID)
	require.Equal(t, arg.DateCreatePost, post.DateCreatePost)
	require.NotZero(t, post.ID)

	return post
}

func TestCreateComment(t *testing.T) {
	// for i := 0; i < 10; i++ {
	// 	createRandomComment(t)
	// }
	createRandomComment(t)
}
