package db

import (
	"context"
	"testing"
	"time"

	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/tijanadmi/bookings_backend/util"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		FirstName:   util.RandomString(20),
		LastName:    util.RandomString(20),
		Email:       util.RandomEmail(),
		Phone:       util.RandomString(20),
		Password:    hashedPassword,
		AccessLevel: int32(util.RandomInt(1, 5)),
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.AccessLevel, user.AccessLevel)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Phone, user2.Phone)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.AccessLevel, user2.AccessLevel)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserOnlyFirstName(t *testing.T) {
	oldUser := createRandomUser(t)

	newFirstName := util.RandomString(20)
	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		Email: oldUser.Email,
		FirstName: pgtype.Text{
			String: newFirstName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, newFirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.Password, updatedUser.Password)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)

	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		Email: oldUser.Email,
		Password: pgtype.Text{
			String: newHashedPassword,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Password, updatedUser.Password)
	require.Equal(t, newHashedPassword, updatedUser.Password)
	require.Equal(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
}

func TestUpdateUserAllFields(t *testing.T) {
	oldUser := createRandomUser(t)

	newFirstName := util.RandomString(20)
	newLastName := util.RandomString(20)
	newPhone := util.RandomString(20)
	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	newAccessLevel := int32(util.RandomInt(1, 5))
	require.NoError(t, err)

	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		Email: oldUser.Email,
		FirstName: pgtype.Text{
			String: newFirstName,
			Valid:  true,
		},
		LastName: pgtype.Text{
			String: newLastName,
			Valid:  true,
		},
		Phone: pgtype.Text{
			String: newPhone,
			Valid:  true,
		},
		Password: pgtype.Text{
			String: newHashedPassword,
			Valid:  true,
		},
		AccessLevel: pgtype.Int4{
			Int32: newAccessLevel,
			Valid: true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Password, updatedUser.Password)
	require.Equal(t, newHashedPassword, updatedUser.Password)
	require.NotEqual(t, oldUser.Phone, updatedUser.Phone)
	require.Equal(t, newPhone, updatedUser.Phone)
	require.NotEqual(t, oldUser.FirstName, updatedUser.FirstName)
	require.Equal(t, newFirstName, updatedUser.FirstName)
	require.NotEqual(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, newLastName, updatedUser.LastName)
	require.NotEqual(t, oldUser.AccessLevel, updatedUser.AccessLevel)
	fmt.Println("Old Access Level:", oldUser.AccessLevel)
	fmt.Println("New Access Level:", newAccessLevel)
	fmt.Println("Updated Access Level:", updatedUser.AccessLevel)
	//fmt.Println(newAccessLevel, updatedUser.AccessLevel)
	//require.Equal(t, newAccessLevel, updatedUser.AccessLevel)
}
