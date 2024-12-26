package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestLikeHack(t *testing.T) {
	for i := 5; i < 30; i++ {
		_, err := testQueries.CreateReact(context.Background(), db.CreateReactParams{
			AccountID: int64(i),
			PostID:    282,
			State:     1,
		})
		require.NoError(t, err)
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
	posts := []db.CreatePostParams{
		{PostTypeID: 1, AccountID: 25, Description: pgtype.Text{String: "Test nguoi la", Valid: true}},
		{PostTypeID: 1, AccountID: 27, Description: pgtype.Text{String: "Test nguoi la", Valid: true}},
		{PostTypeID: 1, AccountID: 30, Description: pgtype.Text{String: "Test nguoi la", Valid: true}},
		{PostTypeID: 1, AccountID: 1, Description: pgtype.Text{String: "Test nguoi la", Valid: true}},
	}

	for i := 1; i <= 10; i++ {
		for _, post := range posts {
			_, err := testQueries.CreatePost(context.Background(), post)
			require.NoError(t, err)
		}
	}
}

func TestGetPost(t *testing.T) {
	for i := 0; i < 10; i++ {
		post, err := testQueries.CreatePost(context.Background(), db.CreatePostParams{
			PostTypeID:     1,
			AccountID:      27,
			Description:    pgtype.Text{String: "Dia diem A", Valid: true},
			StGeomfromtext: fmt.Sprintf("POINT(%f %f)", util.RandomX(), util.RandomY()),
		})
		require.NoError(t, err)
		testQueries.AddImagePost(context.Background(), db.AddImagePostParams{UrlImage: util.RandomImage(), PostID: post.ID})
		testQueries.AddImagePost(context.Background(), db.AddImagePostParams{UrlImage: util.RandomImage(), PostID: post.ID})
		testQueries.AddImagePost(context.Background(), db.AddImagePostParams{UrlImage: util.RandomImage(), PostID: post.ID})
		testQueries.AddImagePost(context.Background(), db.AddImagePostParams{UrlImage: util.RandomImage(), PostID: post.ID})
	}
}

func TestUpgradeQueueHach(t *testing.T) {
	for i := 30; i <= 40; i++ {
		err := testQueries.UpgradeOnwerRequest(context.Background(), db.UpgradeOnwerRequestParams{
			AccountID:      int64(i),
			UpgradePriceID: 1,
		})
		require.NoError(t, err)
	}
}

func TestReport(t *testing.T) {
	for i := 30; i <= 41; i++ {
		_, err := testQueries.CreateReport(context.Background(), db.CreateReportParams{
			AccountID: int64(i),
			PostID:    283,
			IssueID:   1,
		})
		require.NoError(t, err)
	}
}
