package dump

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomComment(t *testing.T, post_id int64) db.Post {
	user := createRandomUser(t)
	arg := db.CreateCommentParams{
		UserID:         user.ID,
		PostTopID:      sql.NullInt64{Int64: post_id, Valid: true},
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
	postTop := createRandomPost(t)

	createRandomComment(t, postTop.ID)
}

func TestGetComment(t *testing.T) {
	postTop := createRandomPost(t)
	comment1 := createRandomComment(t, postTop.ID)
	comment2, err := testQueries.GetComment(context.Background(), sql.NullInt64{Int64: postTop.ID, Valid: true})

	require.NoError(t, err)
	require.NotEmpty(t, comment2)
	require.Equal(t, comment1.ID, comment2.ID)
	require.Equal(t, comment1.PostTypeID, comment2.PostTypeID)
	require.Equal(t, comment1.Description, comment2.Description)
	require.Equal(t, comment1.PostTopID, comment2.PostTopID)
	require.Equal(t, comment1.UserID, comment2.UserID)
	require.Equal(t, comment1.DateCreatePost, comment2.DateCreatePost)

}

func TestListComment(t *testing.T) {
	postTop := createRandomPost(t)
	for i := 0; i < 10; i++ {
		createRandomComment(t, postTop.ID)
	}
	arg := db.ListCommentParams{
		PostTopID: sql.NullInt64{Int64: postTop.ID, Valid: true},
		Limit:     5,
		Offset:    5,
	}
	comments, err := testQueries.ListComment(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, comments, 5)
	for _, comment := range comments {
		require.NotEmpty(t, comment)
	}
}
