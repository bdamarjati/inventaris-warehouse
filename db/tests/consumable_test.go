package db_test

import (
	"inventory/main/db"
	"inventory/main/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomConsumable(t *testing.T) db.Consumables {
	category := createRandomCategory(t)
	require.NotEmpty(t, category)

	status := createRandomStatus(t)
	require.NotEmpty(t, status)

	arg := db.CreateConsumableParams{
		Name:       util.RandomName(),
		Quantity:   int64(util.RandomNumber(1, 100)),
		CategoryId: category.ID,
		Status:     status.ID,
	}

	consumable, err := testQueries.CreateConsumable(arg)
	require.NoError(t, err)
	require.NotEmpty(t, consumable)
	require.NotZero(t, consumable.ID)
	require.Equal(t, arg.Name, consumable.Name)
	require.Equal(t, arg.Quantity, consumable.Quantity)
	require.Equal(t, arg.CategoryId, consumable.CategoryId)
	require.Equal(t, arg.Status, consumable.Status)
	require.WithinDuration(t, consumable.CreatedAt, time.Now().Local(), time.Second)
	return consumable
}

func getConsumable(id int64) (db.Consumables, error) {
	consumable, err := testQueries.GetConsumable(id)
	return consumable, err
}

func TestCreateCOnsumable(t *testing.T) {
	createRandomConsumable(t)
}

func TestGetConsumable(t *testing.T) {
	consumable1 := createRandomConsumable(t)
	require.NotEmpty(t, consumable1)

	consumable2, err := getConsumable(consumable1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, consumable2)

	require.Equal(t, consumable1.ID, consumable2.ID)
	require.Equal(t, consumable1.Name, consumable2.Name)
	require.Equal(t, consumable1.Quantity, consumable2.Quantity)
	require.Equal(t, consumable1.CategoryId, consumable2.CategoryId)
	require.Equal(t, consumable1.Status, consumable2.Status)
	require.WithinDuration(t, consumable1.CreatedAt, consumable2.CreatedAt, time.Second)
}

func TestListConsumable(t *testing.T) {
	for range 3 {
		createRandomConsumable(t)
	}

	arg := db.ListConsumableParams{
		Limit:  3,
		Offset: 0,
	}

	consumables, err := testQueries.ListConsumables(arg)
	require.NoError(t, err)
	require.NotEmpty(t, consumables)
	require.Len(t, consumables, 3)
}

func TestUpdateConsumable(t *testing.T) {
	consumable1 := createRandomConsumable(t)
	require.NotEmpty(t, consumable1)

	arg := db.UpdateConsumableParams{
		ID:         consumable1.ID,
		Name:       util.RandomName(),
		Quantity:   int64(util.RandomNumber(1, 100)),
		CategoryId: consumable1.CategoryId,
		Status:     consumable1.Status,
	}

	rowAffected, err := testQueries.UpdateConsumable(arg)
	require.NoError(t, err)
	require.NotZero(t, rowAffected)

	consumable2, err := getConsumable(consumable1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, consumable2)

	require.Equal(t, consumable1.ID, consumable2.ID)
	require.Equal(t, arg.Name, consumable2.Name)
	require.Equal(t, arg.Quantity, consumable2.Quantity)
	require.Equal(t, arg.CategoryId, consumable2.CategoryId)
	require.Equal(t, arg.Status, consumable2.Status)
	require.NotEqual(t, consumable1.UpdatedAt, consumable2.UpdatedAt)
}

func TestDeleteConsumable(t *testing.T) {
	consumable1 := createRandomConsumable(t)
	require.NotEmpty(t, consumable1)

	rowAffected, err := testQueries.DeleteConsumable(consumable1.ID)
	require.NoError(t, err)
	require.NotZero(t, rowAffected)

	consumable2, err := getConsumable(consumable1.ID)
	require.Error(t, err)
	require.Empty(t, consumable2)
}
