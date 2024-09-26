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

func createRandomUser(t *testing.T) db.User {
	hashPassword, err := util.HashPashword("kocanpass")
	require.NoError(t, err)
	arg := db.CreateUserParams{
		Username:          util.RandomString(8),
		Email:             sql.NullString{String: util.RandomEmail(), Valid: true},
		HashPashword:      hashPassword,
		Fullname:          util.RandomString(7),
		Gender:            util.RandomGender(),
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
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Fullname, user2.Fullname)
	require.Equal(t, user1.Gender, user2.Gender)
	require.Equal(t, user1.RoleID, user2.RoleID)
	require.Equal(t, user1.DateCreateAccount, user2.DateCreateAccount)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := db.UpdateUserParams{
		ID:       user1.ID,
		Fullname: util.RandomString(6),
		Gender:   util.RandomGender(),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, arg.Fullname, user2.Fullname)
	require.Equal(t, arg.Gender, user2.Gender)
	require.Equal(t, user1.RoleID, user2.RoleID)
	require.Equal(t, user1.DateCreateAccount, user2.DateCreateAccount)
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	arg := db.GetListUserParams{
		Limit:  5,
		Offset: 5,
	}
	users, err := testQueries.GetListUser(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
