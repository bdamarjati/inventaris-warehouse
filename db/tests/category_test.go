package db_test

import (
	"inventory/main/db"
	"inventory/main/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) db.RefCategories {
	category, err := testQueries.CreateCategory(util.RandomCategory())
	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.NotZero(t, category.ID)
	require.WithinDuration(t, category.CreatedAt, time.Now().Local(), time.Second)

	return category
}

func getCategory(id int64) (db.RefCategories, error) {
	category, err := testQueries.GetCategory(id)
	return category, err
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

	require.Equal(t, category1.Name, category2.Name)
	require.WithinDuration(t, category1.CreatedAt, category2.CreatedAt, time.Second)
}

func TestListCategory(t *testing.T) {
	for range 5 {
		createRandomCategory(t)
	}

	arg := db.ListCategoriesParams{
		Limit:  5,
		Offset: 0,
	}

	categories, err := testQueries.ListCategories(arg)
	require.NoError(t, err)
	require.NotEmpty(t, categories)
	require.Len(t, categories, 5)
}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	require.NotEmpty(t, category1)

	arg := db.UpdateCategoryParams{
		ID:   category1.ID,
		Name: util.RandomCategory(),
	}

	rowAffected, err := testQueries.UpdateCategory(arg)
	require.NoError(t, err)
	require.NotZero(t, rowAffected)

	category2, err := getCategory(arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.Name, category2.Name)
	require.WithinDuration(t, category1.CreatedAt, category2.CreatedAt, time.Minute)
}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	require.NotEmpty(t, category1)

	rowAffected, err := testQueries.DeleteCategory(category1.ID)
	require.NoError(t, err)
	require.NotZero(t, rowAffected)

	category2, err := getCategory(category1.ID)
	require.Error(t, err)
	require.Empty(t, category2)
}
