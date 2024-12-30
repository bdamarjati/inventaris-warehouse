package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"inventory/main/db"
	"inventory/main/util"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomInventory(t *testing.T) db.Inventories {
	category := createRandomCategory(t)
	require.NotEmpty(t, category)

	status := createRandomStatus(t)
	require.NotEmpty(t, status)

	arg := createInventoryRequest{
		Name:       util.RandomText(10),
		Quantity:   int64(util.RandomNumber(1, 10)),
		CategoryId: category.ID,
		Condition:  0,
		Status:     status.ID,
	}

	data, err := json.Marshal(arg)
	require.NoError(t, err)

	url := "/inventory"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	inventory, err := requireBodyMatchInventory(recorder.Body)
	require.NoError(t, err)
	require.NotEmpty(t, inventory)

	require.Equal(t, arg.Name, inventory.Name)
	require.Equal(t, arg.Quantity, inventory.Quantity)
	require.Equal(t, arg.CategoryId, inventory.CategoryId)
	require.Equal(t, arg.Condition, inventory.Condition)
	require.Equal(t, arg.Status, inventory.Status)
	require.WithinDuration(t, inventory.CreatedAt, time.Now().Local(), time.Second)
	return inventory
}

func getInventory(id int64) (db.Inventories, error) {
	inventory := db.Inventories{}

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/inventory/%d", id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return inventory, err
	}

	err = addAuthorization(request)
	if err != nil {
		return inventory, err
	}

	testServer.router.ServeHTTP(recorder, request)
	if recorder.Code == http.StatusInternalServerError {
		return inventory, fmt.Errorf("internal testServer error: %v", err)
	}

	inventory, err = requireBodyMatchInventory(recorder.Body)
	return inventory, err
}

func requireBodyMatchInventory(body *bytes.Buffer) (db.Inventories, error) {
	var gotInventory db.Inventories

	data, err := io.ReadAll(body)
	if err != nil {
		return gotInventory, err
	}

	err = json.Unmarshal(data, &gotInventory)
	return gotInventory, err
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

func TestListInventories(t *testing.T) {
	for range 3 {
		createRandomInventory(t)
	}

	inventories := []db.Inventories{}

	arg := listInventoryRequest{
		Size: 3,
		Page: 1,
	}

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/inventories/%d/%d", arg.Size, arg.Page)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, recorder.Code, http.StatusOK)

	result, err := io.ReadAll(recorder.Body)
	require.NoError(t, err)

	err = json.Unmarshal(result, &inventories)
	require.NoError(t, err)

	require.NotEmpty(t, inventories)
	require.Len(t, inventories, arg.Size)
}

func TestUpdateInventory(t *testing.T) {
	inventory1 := createRandomInventory(t)
	require.NotEmpty(t, inventory1)

	arg := updateInventoryRequest{
		ID:         inventory1.ID,
		Name:       util.RandomText(10),
		Quantity:   int64(util.RandomNumber(1, 10)),
		CategoryId: inventory1.CategoryId,
		Condition:  1,
		Status:     inventory1.Status,
	}

	data, err := json.Marshal(arg)
	require.NoError(t, err)

	url := "/inventory"
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	inventory2, err := getInventory(arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, inventory2)

	require.Equal(t, inventory1.ID, inventory2.ID)
	require.Equal(t, arg.Name, inventory2.Name)
	require.Equal(t, arg.Quantity, inventory2.Quantity)
	require.Equal(t, arg.CategoryId, inventory2.CategoryId)
	require.Equal(t, arg.Condition, inventory2.Condition)
	require.Equal(t, arg.Status, inventory2.Status)
	require.WithinDuration(t, inventory1.CreatedAt, inventory2.CreatedAt, time.Second)
	require.WithinDuration(t, inventory1.UpdatedAt, inventory2.UpdatedAt, time.Minute)
}

func TestDeleteInventory(t *testing.T) {
	inventory1 := createRandomInventory(t)
	require.NotEmpty(t, inventory1)

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/inventory/%d", inventory1.ID)
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	inventory2, err := getInventory(inventory1.ID)
	require.Error(t, err)
	require.Empty(t, inventory2)
}
