package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: "Test",
		Password: "test123",
		Email:    "test@email.com",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)

	require.NotEmpty(t, user.CreatedAt)

	return user

}

func Test_CreateUser(t *testing.T) {
	createRandomUser(t)
}
