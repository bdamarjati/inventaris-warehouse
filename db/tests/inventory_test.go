package db_test

import (
	"inventory/main/db"
	"inventory/main/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomInventory(t *testing.T) db.Inventories {
	category := createRandomCategory(t)
	require.NotEmpty(t, category)

	status := createRandomStatus(t)
	require.NotEmpty(t, status)

	arg := db.CreateInventoryParams{
		Name:       util.RandomName(),
		Quantity:   int64(util.RandomNumber(1, 100)),
		CategoryId: category.ID,
		Condition:  0,
		Status:     status.ID,
	}

	inventory, err := testQueries.CreateInventory(arg)
	require.NoError(t, err)
	require.NotEmpty(t, inventory)
	require.NotZero(t, inventory.ID)
	require.WithinDuration(t, inventory.CreatedAt, time.Now().Local(), time.Second)

	return inventory
}

func getInventory(id int64) (db.Inventories, error) {
	inventory, err := testQueries.GetInventory(id)
	return inventory, err
}

func TestCreateInventory(t *testing.T) {
	createRandomInventory(t)
}

func TestGetInventory(t *testing.T) {
	inventory1 := createRandomInventory(t)
	require.NotEmpty(t, inventory1)

	inventory2, err := getInventory(inventory1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, inventory2)

	require.Equal(t, inventory1.ID, inventory2.ID)
	require.Equal(t, inventory1.Name, inventory2.Name)
	require.Equal(t, inventory1.Quantity, inventory2.Quantity)
	require.Equal(t, inventory1.CategoryId, inventory2.CategoryId)
	require.Equal(t, inventory1.Condition, inventory2.Condition)
	require.Equal(t, inventory1.Status, inventory2.Status)
	require.WithinDuration(t, inventory1.CreatedAt, inventory2.CreatedAt, time.Second)
}

func TestListInventory(t *testing.T) {
	for range 3 {
		createRandomInventory(t)
	}

	arg := db.ListInventoryParams{
		Limit:  3,
		Offset: 0,
	}

	inventories, err := testQueries.ListInventories(arg)
	require.NoError(t, err)
	require.NotEmpty(t, inventories)
	require.Len(t, inventories, 3)
}

func TestUpdateInventory(t *testing.T) {
	inventory1 := createRandomInventory(t)
	require.NotEmpty(t, inventory1)

	arg := db.UpdateInventoryParams{
		ID:         inventory1.ID,
		Name:       util.RandomName(),
		Quantity:   int64(util.RandomNumber(1, 100)),
		CategoryId: inventory1.CategoryId,
		Condition:  1,
		Status:     inventory1.Status,
	}

	rowAffected, err := testQueries.UpdateInventory(arg)
	require.NoError(t, err)
	require.NotZero(t, rowAffected)

	inventory2, err := getInventory(arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, inventory2)

	require.Equal(t, inventory1.ID, inventory2.ID)
	require.Equal(t, arg.Name, inventory2.Name)
	require.Equal(t, arg.Quantity, inventory2.Quantity)
	require.Equal(t, arg.CategoryId, inventory2.CategoryId)
	require.Equal(t, arg.Condition, inventory2.Condition)
	require.Equal(t, arg.Status, inventory2.Status)
	require.NotEqual(t, inventory1.UpdatedAt, inventory2.UpdatedAt)
}

func TestDeleteInventory(t *testing.T) {
	inventory1 := createRandomInventory(t)
	require.NotEmpty(t, inventory1)

	rowAffected, err := testQueries.DeleteInventory(inventory1.ID)
	require.NoError(t, err)
	require.NotZero(t, rowAffected)

	inventory2, err := getInventory(inventory1.ID)
	require.Error(t, err)
	require.Empty(t, inventory2)
}
