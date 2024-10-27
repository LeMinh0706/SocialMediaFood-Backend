package dump

import (
	"context"
	"fmt"
	"testing"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createPostNoPoint(t *testing.T) db.CreatePostRow {
	user := createRandomUser(t)
	account := createRandomAccount(t, user.ID, 3)
	description := util.RandomDescription()
	arg := db.CreatePostParams{
		PostTypeID:  1,
		AccountID:   account.ID,
		Description: pgtype.Text{String: description, Valid: true},
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.PostTypeID, post.PostTypeID)
	require.Equal(t, arg.AccountID, post.AccountID)
	require.Equal(t, arg.Description, post.Description)

	require.NotZero(t, post.ID)
	require.NotZero(t, post.CreatedAt)
	return post
}

func createPostImage(t *testing.T) db.CreatePostRow {
	post := createPostNoPoint(t)

	for i := 0; i < 4; i++ {
		img, err := testQueries.AddImagePost(context.Background(), db.AddImagePostParams{
			PostID:   post.ID,
			UrlImage: util.RandomImage(),
		})
		require.NoError(t, err)
		require.NotEmpty(t, img)

		require.Equal(t, post.ID, img.PostID)
	}

	return post
}

func TestCreatePost(t *testing.T) {
	createPostNoPoint(t)
}

func TestCreatePostImage(t *testing.T) {
	createPostImage(t)
}

func TestGetImagePost(t *testing.T) {
	post := createPostImage(t)

	images, err := testQueries.GetImagePost(context.Background(), post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, images)
	require.Len(t, images, 4)
}

func TestCreatePostPosition(t *testing.T) {
	user := createRandomUser(t)
	account := createRandomAccount(t, user.ID, 3)
	description := util.RandomDescription()
	arg := db.CreatePostParams{
		PostTypeID:     1,
		AccountID:      account.ID,
		Description:    pgtype.Text{String: description, Valid: true},
		StGeomfromtext: fmt.Sprintf("POINT(%f %f)", util.RandomX(), util.RandomY()),
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NotEmpty(t, post)
	require.NoError(t, err)
}
