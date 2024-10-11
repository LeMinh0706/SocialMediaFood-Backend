package dump

import (
	"context"
	"database/sql"
	"testing"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/stretchr/testify/require"
)

func createReact(t *testing.T, user_id, post_id int64) db.ReactPost {
	react, err := testQueries.CreateReact(context.Background(), db.CreateReactParams{
		PostID: post_id,
		UserID: user_id,
	})

	require.NoError(t, err)
	require.NotEmpty(t, react)

	require.Equal(t, post_id, react.PostID)
	require.Equal(t, user_id, react.UserID)
	return react
}

func TestCreateReact(t *testing.T) {
	user := createRandomUser(t)
	post := createRandomPost(t)
	createReact(t, user.ID, post.ID)
}

func TestManyReact(t *testing.T) {
	post1 := createRandomPost(t)
	for i := 0; i < 10; i++ {
		user := createRandomUser(t)
		react := createReact(t, user.ID, post1.ID)
		require.NotEmpty(t, react)
		require.Equal(t, post1.ID, react.PostID)
	}

	user1 := createRandomUser(t)
	for i := 0; i < 10; i++ {
		post := createRandomPost(t)
		react := createReact(t, user1.ID, post.ID)
		require.NotEmpty(t, react)
		require.Equal(t, user1.ID, react.UserID)
	}
}

func TestDuplicateReact(t *testing.T) {
	user := createRandomUser(t)
	post := createRandomPost(t)
	react1 := createReact(t, user.ID, post.ID)
	require.NotEmpty(t, react1)
	react2, err := testQueries.CreateReact(context.Background(), db.CreateReactParams{
		PostID: post.ID,
		UserID: user.ID,
	})
	require.Error(t, err)
	require.Empty(t, react2)
}

func TestDeleteReact(t *testing.T) {
	user := createRandomUser(t)
	post := createRandomPost(t)
	react1 := createReact(t, user.ID, post.ID)

	err := testQueries.DeleteReact(context.Background(), react1.ID)
	require.NoError(t, err)
	react2, err := testQueries.GetReact(context.Background(), db.GetReactParams{
		PostID: post.ID,
		UserID: user.ID,
	})
	require.Error(t, err)
	require.Empty(t, react2)
	require.Equal(t, err, sql.ErrNoRows)
}
