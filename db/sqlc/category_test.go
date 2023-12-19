package db

import (
	"context"
	"go-finances/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	user := createRandomUser(t)
	arg := CreateCategoryParams{
		UserID:      user.ID,
		Title:       util.RandomString(12),
		Type:        "debit",
		Description: util.RandomString(20),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.UserID, category.UserID)
	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.Description, category.Description)

	require.NotEmpty(t, category.CreatedAt)

	require.NotEmpty(t, user.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category.ID)

	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category.ID, category2.ID)
	require.Equal(t, category.Title, category2.Title)
	require.Equal(t, category.Type, category2.Type)
	require.Equal(t, category.UserID, category2.UserID)

	require.NotEmpty(t, category2.CreatedAt)

}

func TestDeleteCategory(t *testing.T) {
	category := createRandomCategory(t)
	err := testQueries.DeleteCategories(context.Background(), category.ID)

	require.NoError(t, err)

}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	arg := UpdateCategoriesParams{
		ID:          category1.ID,
		Title:       util.RandomString(12),
		Description: util.RandomString(20),
	}

	category, err := testQueries.UpdateCategories(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.Description, category.Description)
	require.Equal(t, category1.ID, category.ID)

}

func TestListCategories(t *testing.T) {

	lastCategory := createRandomCategory(t)

	arg := GetCategoriesParams{
		UserID:      lastCategory.UserID,
		Type:        lastCategory.Type,
		Title:       lastCategory.Title,
		Description: lastCategory.Description,
	}

	categories, err := testQueries.GetCategories(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, categories)

	for _, category := range categories {
		require.Equal(t, lastCategory.Title, category.Title)
		require.Equal(t, lastCategory.Description, category.Description)
		require.Equal(t, lastCategory.ID, category.ID)
		require.Equal(t, lastCategory.UserID, category.UserID)
	}

}
