package test

import (
	"context"
	"testing"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
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
