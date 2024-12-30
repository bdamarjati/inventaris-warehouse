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

func createRandomConsumable(t *testing.T) db.Consumables {
	category := createRandomCategory(t)
	require.NotEmpty(t, category)

	status := createRandomStatus(t)
	require.NotEmpty(t, status)

	arg := createConsumableRequest{
		Name:       util.RandomText(10),
		Quantity:   int64(util.RandomNumber(1, 10)),
		CategoryId: category.ID,
		Status:     status.ID,
	}

	data, err := json.Marshal(arg)
	require.NoError(t, err)

	url := "/consumable"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	consumable, err := requireBodyMatchConsumable(recorder.Body)
	require.NoError(t, err)
	require.NotEmpty(t, consumable)

	require.Equal(t, arg.Name, consumable.Name)
	require.Equal(t, arg.Quantity, consumable.Quantity)
	require.Equal(t, arg.CategoryId, consumable.CategoryId)
	require.Equal(t, arg.Status, consumable.Status)
	require.WithinDuration(t, consumable.CreatedAt, time.Now().Local(), time.Second)
	return consumable
}

func getConsumable(id int64) (db.Consumables, error) {
	consumable := db.Consumables{}

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/consumable/%d", id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return consumable, err
	}

	err = addAuthorization(request)
	if err != nil {
		return consumable, err
	}

	testServer.router.ServeHTTP(recorder, request)
	if recorder.Code == http.StatusInternalServerError {
		return consumable, fmt.Errorf("internal testServer error: %v", err)
	}

	consumable, err = requireBodyMatchConsumable(recorder.Body)
	return consumable, err
}

func requireBodyMatchConsumable(body *bytes.Buffer) (db.Consumables, error) {
	var gotConsumable db.Consumables

	data, err := io.ReadAll(body)
	if err != nil {
		return gotConsumable, err
	}

	err = json.Unmarshal(data, &gotConsumable)
	return gotConsumable, err
}

func TestCreateConsumable(t *testing.T) {
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

func TestListConsumables(t *testing.T) {
	for range 3 {
		createRandomConsumable(t)
	}

	consumables := []db.Consumables{}

	arg := listConsumableRequest{
		Size: 3,
		Page: 1,
	}

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/consumables/%d/%d", arg.Size, arg.Page)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, recorder.Code, http.StatusOK)

	result, err := io.ReadAll(recorder.Body)
	require.NoError(t, err)

	err = json.Unmarshal(result, &consumables)
	require.NoError(t, err)

	require.NotEmpty(t, consumables)
	require.Len(t, consumables, arg.Size)
}

func TestUpdateConsumable(t *testing.T) {
	consumable1 := createRandomConsumable(t)
	require.NotEmpty(t, consumable1)

	arg := updateConsumableRequest{
		ID:         consumable1.ID,
		Name:       util.RandomText(10),
		Quantity:   int64(util.RandomNumber(1, 10)),
		CategoryId: consumable1.CategoryId,
		Status:     consumable1.Status,
	}

	data, err := json.Marshal(arg)
	require.NoError(t, err)

	url := "/consumable"
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	consumable2, err := getConsumable(arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, consumable2)

	require.Equal(t, consumable1.ID, consumable2.ID)
	require.Equal(t, arg.Name, consumable2.Name)
	require.Equal(t, arg.Quantity, consumable2.Quantity)
	require.Equal(t, arg.CategoryId, consumable2.CategoryId)
	require.Equal(t, arg.Status, consumable2.Status)
	require.WithinDuration(t, consumable1.CreatedAt, consumable2.CreatedAt, time.Second)
	require.WithinDuration(t, consumable1.UpdatedAt, consumable2.UpdatedAt, time.Minute)
}

func TestDeleteConsumable(t *testing.T) {
	consumable1 := createRandomConsumable(t)
	require.NotEmpty(t, consumable1)

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/consumable/%d", consumable1.ID)
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	require.NoError(t, err)

	err = addAuthorization(request)
	require.NoError(t, err)

	testServer.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	consumable2, err := getConsumable(consumable1.ID)
	require.Error(t, err)
	require.Empty(t, consumable2)
}
