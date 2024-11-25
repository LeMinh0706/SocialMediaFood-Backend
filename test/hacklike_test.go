package test

import (
	"context"
	"testing"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestHackLike(t *testing.T) {
	for i := 1; i < 43; i++ {
		react, _ := testQueries.CreateReact(context.Background(), db.CreateReactParams{
			AccountID: int64(i),
			PostID:    19,
			State:     1,
		})
		require.NotEmpty(t, react)
	}
}

func TestHackFollow(t *testing.T) {
	// for i := 1; i <= 26; i++ {
	// 	testQueries.CreateFollow(context.Background(), db.CreateFollowParams{
	// 		FromFollow: 27,
	// 		ToFollow:   int64(i),
	// 		Status:     "friend",
	// 	})
	// 	testQueries.CreateFollow(context.Background(), db.CreateFollowParams{
	// 		FromFollow: int64(i),
	// 		ToFollow:   27,
	// 		Status:     "friend",
	// 	})
	// }
	testQueries.CreateFollow(context.Background(), db.CreateFollowParams{
		FromFollow: 30,
		ToFollow:   20,
		Status:     "request",
	})
	testQueries.CreateFollow(context.Background(), db.CreateFollowParams{
		FromFollow: 20,
		ToFollow:   30,
		Status:     "accept",
	})
}

func TestCreatePost(t *testing.T) {
	for i := 28; i <= 40; i++ {
		_, err := testQueries.CreatePost(context.Background(), db.CreatePostParams{
			PostTypeID:  1,
			AccountID:   int64(i),
			Description: pgtype.Text{String: "Test nguoi la", Valid: true},
		})
		require.NoError(t, err)
	}
}
