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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) db.RefCategories {
	arg := db.RefCategories{
		Name: util.RandomCategory(),
	}

	body := gin.H{
		"name": arg.Name,
	}

	server := newTestServer(*testStore)
	recorder := httptest.NewRecorder()

	data, err := json.Marshal(body)
	require.NoError(t, err)

	url := "/category"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	category, err := requireBodyMatchCategory(recorder.Body)
	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, category.Name, arg.Name)
	require.WithinDuration(t, category.CreatedAt, time.Now().Local(), time.Second)
	return category
}

func getCategory(id int64) (db.RefCategories, error) {
	category := db.RefCategories{}

	server := newTestServer(*testStore)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/category/%d", id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return category, err
	}

	server.router.ServeHTTP(recorder, request)
	if recorder.Code == http.StatusInternalServerError {
		return category, fmt.Errorf("internal server error: %v", err)
	}

	category, err = requireBodyMatchCategory(recorder.Body)
	return category, err
}

func requireBodyMatchCategory(body *bytes.Buffer) (db.RefCategories, error) {
	var gotCategory db.RefCategories

	data, err := io.ReadAll(body)
	if err != nil {
		return gotCategory, err
	}

	err = json.Unmarshal(data, &gotCategory)
	return gotCategory, err
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	require.NotEmpty(t, category1)

	category2, err := getCategory(category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Name, category2.Name)
	require.WithinDuration(t, category1.CreatedAt, category2.CreatedAt, time.Second)
}

func TestListCategories(t *testing.T) {
	for range 3 {
		createRandomCategory(t)
	}

	categories := []db.RefCategories{}

	arg := listCategoryRequest{
		Size: 2,
		Page: 1,
	}

	server := newTestServer(*testStore)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/categories/%d/%d", arg.Size, arg.Page)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, recorder.Code, http.StatusOK)

	result, err := io.ReadAll(recorder.Body)
	require.NoError(t, err)

	err = json.Unmarshal(result, &categories)
	require.NoError(t, err)

	require.NotEmpty(t, categories)
	require.Len(t, categories, arg.Size)
}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	require.NotEmpty(t, category1)

	arg := updateCategoryRequest{
		ID:   category1.ID,
		Name: util.RandomCategory(),
	}

	body := gin.H{
		"id":   arg.ID,
		"name": arg.Name,
	}

	data, err := json.Marshal(body)
	require.NoError(t, err)

	url := "/category"
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
	require.NoError(t, err)

	server := newTestServer(*testStore)
	recorder := httptest.NewRecorder()

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	category2, err := getCategory(arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.Name, category2.Name)
	require.WithinDuration(t, category1.CreatedAt, category2.CreatedAt, time.Second)
}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	require.NotEmpty(t, category1)

	server := newTestServer(*testStore)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/category/%d", category1.ID)
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	category2, err := getCategory(category1.ID)
	require.Error(t, err)
	require.Empty(t, category2)
}
