package dump

import (
	"context"
	"testing"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) db.RegisterRow {
	password := "kocanpass"
	hashed, err := util.HashPassword(password)
	require.NoError(t, err)
	arg := db.RegisterParams{
		Username:     util.RandomString(6),
		Email:        pgtype.Text{String: util.RandomEmail(), Valid: true},
		HashPassword: hashed,
	}
	user, err := testQueries.Register(context.Background(), arg)
	require.NotEmpty(t, user)
	require.NoError(t, err)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	return user
}

func createRandomAccount(t *testing.T, user_id int64, typeA int32) db.Account {
	gender := util.RandomGender()
	arg := db.CreateAccountsParams{
		UserID:               user_id,
		Fullname:             util.RandomString(6),
		Gender:               pgtype.Int4{Int32: gender, Valid: true},
		Type:                 typeA,
		UrlAvatar:            util.RandomAvatar(gender),
		UrlBackgroundProfile: db.GetBackground(),
	}
	account, err := testQueries.CreateAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.UserID, account.UserID)
	require.Equal(t, arg.Fullname, account.Fullname)
	require.Equal(t, gender, account.Gender.Int32)
	require.Equal(t, arg.UrlAvatar, account.UrlAvatar)
	require.Equal(t, arg.UrlBackgroundProfile, account.UrlBackgroundProfile)
	return account
}

func createRandomRegister(t *testing.T) db.Account {
	user := createRandomUser(t)
	acc := createRandomAccount(t, user.ID, 3)
	return acc
}
func TestRegister(t *testing.T) {
	createRandomRegister(t)
}

func TestGetAccount(t *testing.T) {
	user := createRandomUser(t)
	login, err := testQueries.Login(context.Background(), user.Username)
	require.NotEmpty(t, login)
	require.NoError(t, err)
	require.Equal(t, user.Username, login.Username)
	require.Equal(t, user.ID, login.ID)
	createRandomAccount(t, user.ID, 3)
	createRandomAccount(t, user.ID, 2)
	createRandomAccount(t, user.ID, 2)
	createRandomAccount(t, user.ID, 2)
	acc, err := testQueries.GetAccountByUserId(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Len(t, acc, 4)
}

func TestGetOneAccount(t *testing.T) {
	acc1 := createRandomRegister(t)
	acc2, err := testQueries.GetAccountById(context.Background(), acc1.ID)
	require.NotEmpty(t, acc2)
	require.NoError(t, err)

	require.Equal(t, acc1.UserID, acc2.UserID)
}
