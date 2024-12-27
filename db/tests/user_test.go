package db_test

import (
	"inventory/main/db"
	"inventory/main/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) (db.User, error) {
	arg := db.CreateUserParams{
		Username: util.RandomName(),
		Password: util.RandomText(8),
		Role:     "User",
	}

	user, err := testQueries.CreateUser(arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.WithinDuration(t, time.Now().Local(), user.CreatedAt, time.Second)

	return user, err
}

func getUser(username string) (db.User, error) {
	user1, err := testQueries.GetUser(username)
	return user1, err
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1, err := createRandomUser(t)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	result, err := testQueries.GetUser(user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, user1.Username, result.Username)
	require.Equal(t, user1.Password, result.Password)
	require.Equal(t, user1.Role, result.Role)
}

func TestListUser(t *testing.T) {
	for range 5 {
		_, err := createRandomUser(t)
		require.NoError(t, err)
	}

	arg := db.ListUserParams{
		Limit:  5,
		Offset: 0,
	}

	users, err := testQueries.ListUser(arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)
	require.Equal(t, len(users), 5)
}

func TestUpdateUser(t *testing.T) {
	user1, err := createRandomUser(t)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	arg := db.UpdateUserParams{
		Username: user1.Username,
		Password: util.RandomText(8),
		Role:     "Admin",
	}

	rowAffected, err := testQueries.UpdateUser(arg)
	require.NoError(t, err)
	require.NotZero(t, rowAffected)

	user2, err := getUser(user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, arg.Password, user2.Password)
	require.Equal(t, arg.Role, user2.Role)
	require.NotEmpty(t, user2.CreatedAt)
}

func TestDeleteUser(t *testing.T) {
	user1, err := createRandomUser(t)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	rowAffected, err := testQueries.DeleteUser(user1.Username)
	require.NoError(t, err)
	require.NotZero(t, rowAffected)

	user2, err := getUser(user1.Username)
	require.Error(t, err)
	require.Empty(t, user2)
}
