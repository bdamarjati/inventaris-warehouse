package db_test

import (
	"inventory/main/db"
	"inventory/main/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomStatus(t *testing.T) db.RefStatus {
	status, err := testQueries.CreateStatus(util.RandomStatus())
	require.NoError(t, err)
	require.NotEmpty(t, status)
	require.NotZero(t, status.ID)
	require.WithinDuration(t, status.CreatedAt, time.Now().Local(), time.Second)
	return status
}

func getStatus(id int64) (db.RefStatus, error) {
	status, err := testQueries.GetStatus(id)
	return status, err
}

func TestCreateStatus(t *testing.T) {
	createRandomStatus(t)
}

func TestGetStatus(t *testing.T) {
	status1 := createRandomStatus(t)
	require.NotEmpty(t, status1)

	status2, err := getStatus(status1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, status2)

	require.Equal(t, status1.ID, status2.ID)
	require.Equal(t, status1.Description, status2.Description)
	require.WithinDuration(t, status1.CreatedAt, status2.CreatedAt, time.Second)
}

func TestListStatus(t *testing.T) {
	for range 5 {
		createRandomStatus(t)
	}

	arg := db.ListStatusParams{
		Limit:  5,
		Offset: 0,
	}

	statuses, err := testQueries.ListStatus(arg)
	require.NoError(t, err)
	require.NotEmpty(t, statuses)
	require.Len(t, statuses, 5)
}

func TestUpdateStatus(t *testing.T) {
	status1 := createRandomStatus(t)
	require.NotEmpty(t, status1)

	arg := db.UpdateStatusParams{
		ID:          status1.ID,
		Description: util.RandomStatus(),
	}

	rowAffected, err := testQueries.UpdateStatus(arg)
	require.NoError(t, err)
	require.NotZero(t, rowAffected)

	status2, err := getStatus(status1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, status2)

	require.Equal(t, status1.ID, status2.ID)
	require.Equal(t, arg.Description, status2.Description)
	require.WithinDuration(t, status1.CreatedAt, status2.CreatedAt, time.Second)
}

func TestDeleteStatus(t *testing.T) {
	status1 := createRandomStatus(t)
	require.NotEmpty(t, status1)

	rowAffected, err := testQueries.DeleteStatus(status1.ID)
	require.NoError(t, err)
	require.NotZero(t, rowAffected)

	status2, err := getStatus(status1.ID)
	require.Error(t, err)
	require.Empty(t, status2)
}
