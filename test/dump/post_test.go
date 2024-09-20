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

func createRandomPost(t *testing.T) db.Post {
	arg := db.CreatePostParams{
		PostTypeID:     util.RandomType(),
		UserID:         util.RandomInt(1, 10),
		Description:    sql.NullString{String: util.RandomDescription(), Valid: true},
		DateCreatePost: time.Now().Unix(),
	}
	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.PostTypeID, post.PostTypeID)
	require.Equal(t, arg.Description, post.Description)
	require.Equal(t, arg.UserID, post.UserID)
	require.Equal(t, arg.DateCreatePost, post.DateCreatePost)
	require.NotZero(t, post.ID)

	return post
}

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}

func TestGetPost(t *testing.T) {
	post1 := createRandomPost(t)

	post2, err := testQueries.GetPost(context.Background(), post1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.PostTypeID, post2.PostTypeID)
	require.Equal(t, post1.Description, post2.Description)
	require.Equal(t, post1.PostTopID, post2.PostTopID)
	require.Equal(t, post1.UserID, post2.UserID)
	require.Equal(t, post1.DateCreatePost, post2.DateCreatePost)
}

func TestUpdatePost(t *testing.T) {
	post1 := createRandomPost(t)

	arg := db.UpdatePostParams{
		ID:          post1.ID,
		Description: sql.NullString{String: util.RandomDescription(), Valid: true},
	}

	post2, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.PostTypeID, post2.PostTypeID)
	require.Equal(t, arg.Description, post2.Description)
	require.Equal(t, post1.PostTopID, post2.PostTopID)
	require.Equal(t, post1.UserID, post2.UserID)
	require.Equal(t, post1.DateCreatePost, post2.DateCreatePost)
}

func TestDeletePost(t *testing.T) {
	post1 := createRandomPost(t)
	err := testQueries.DeletePost(context.Background(), post1.ID)
	require.NoError(t, err)

	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post2)
}

func TestListPost(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPost(t)
	}
	arg := db.ListPostParams{
		Limit:  5,
		Offset: 5,
	}
	posts, err := testQueries.ListPost(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, posts, 5)

	for _, post := range posts {
		require.NotEmpty(t, post)
	}
}
