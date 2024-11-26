package test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
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
	for i := 1; i <= 500; i++ {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := testQueries.CreatePost(context.Background(), db.CreatePostParams{
				PostTypeID:  1,
				AccountID:   25,
				Description: pgtype.Text{String: "Test nguoi la", Valid: true},
			})
			require.NoError(t, err)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := testQueries.CreatePost(context.Background(), db.CreatePostParams{
				PostTypeID:  1,
				AccountID:   27,
				Description: pgtype.Text{String: "Test nguoi la", Valid: true},
			})
			require.NoError(t, err)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := testQueries.CreatePost(context.Background(), db.CreatePostParams{
				PostTypeID:  1,
				AccountID:   30,
				Description: pgtype.Text{String: "Test nguoi la", Valid: true},
			})
			require.NoError(t, err)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := testQueries.CreatePost(context.Background(), db.CreatePostParams{
				PostTypeID:  1,
				AccountID:   1,
				Description: pgtype.Text{String: "Test nguoi la", Valid: true},
			})
			require.NoError(t, err)
		}()
		wg.Wait()
	}
}

func TestGetPost(t *testing.T) {
	for i := 0; i < 500; i++ {
		_, err := testQueries.CreatePost(context.Background(), db.CreatePostParams{
			PostTypeID:     1,
			AccountID:      27,
			Description:    pgtype.Text{String: "Dia diem A", Valid: true},
			StGeomfromtext: fmt.Sprintf("POINT(%f %f)", util.RandomX(), util.RandomY()),
		})
		require.NoError(t, err)
	}
}
