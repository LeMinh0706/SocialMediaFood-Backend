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

func TestCreateUser(t *testing.T) {
	arg := db.CreateUserParams{
		Email:             sql.NullString{String: util.RandomEmail(), Valid: true},
		HashPashword:      "kocanpass",
		Fullname:          util.RandomString(6),
		Gender:            0,
		RoleID:            1,
		DateCreateAccount: time.Now().Unix(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashPashword, user.HashPashword)
	require.Equal(t, arg.Fullname, user.Fullname)
	require.Equal(t, arg.Gender, user.Gender)
	require.Equal(t, arg.RoleID, user.RoleID)
	require.Equal(t, arg.DateCreateAccount, user.DateCreateAccount)
	require.NotZero(t, user.ID)
}
