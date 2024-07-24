package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/tijanadmi/bookings_backend/util"
)

func createRandomRestriction(t *testing.T) Restriction {

	arg := CreateRestrictionParams{
		RestrictionNameSr:  util.RandomString(5),
		RestrictionNameEn:  util.RandomString(5),
		RestrictionNameBg:  util.RandomString(5),
	}

	restriction, err := testStore.CreateRestriction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, restriction)

	require.Equal(t, arg.RestrictionNameSr, restriction.RestrictionNameSr)
	require.Equal(t, arg.RestrictionNameEn, restriction.RestrictionNameEn)
	require.Equal(t, arg.RestrictionNameBg, restriction.RestrictionNameBg)
	require.NotZero(t, restriction.ID)
	require.NotZero(t, restriction.CreatedAt)

	return restriction
}

func TestCreateRestriction(t *testing.T) {
	createRandomRestriction(t)
}

func TestGetRestriction(t *testing.T) {
	restriction1 := createRandomRestriction(t)
	restriction2, err := testStore.GetRestriction(context.Background(),restriction1.ID)
	require.NoError(t, err)
	require.NotEmpty(t,restriction2)

	require.Equal(t,restriction1.ID,restriction2.ID)
	require.Equal(t, restriction1.RestrictionNameSr, restriction2.RestrictionNameSr)
	require.Equal(t, restriction1.RestrictionNameEn, restriction2.RestrictionNameEn)
	require.Equal(t, restriction1.RestrictionNameBg, restriction2.RestrictionNameBg)
	
	require.WithinDuration(t, restriction1.CreatedAt, restriction2.CreatedAt, time.Second)
}

func TestUpdateRestrictionPriceEn(t *testing.T) {
	restriction1 := createRandomRestriction(t)

	newRestrictionNameSr:=util.RandomString(5)
	arg := UpdateRestrictionParams{
		ID:      restriction1.ID,
		RestrictionNameSr: pgtype.Text{
			String:  newRestrictionNameSr,
			Valid:  true,
		},

	}

	restriction2, err := testStore.UpdateRestriction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, restriction2)

	require.Equal(t, restriction1.ID, restriction2.ID)

	require.NotEqual(t, restriction1.RestrictionNameSr, restriction2.RestrictionNameSr)
	require.Equal(t, newRestrictionNameSr, restriction2.RestrictionNameSr)
	

	require.Equal(t, restriction1.RestrictionNameEn, restriction2.RestrictionNameEn)
	require.Equal(t, restriction1.RestrictionNameBg, restriction2.RestrictionNameBg)

	
	require.WithinDuration(t, restriction1.UpdatedAt, restriction2.UpdatedAt, time.Second)
}

func TestUpdateRestrictionAllFields(t *testing.T) {
	restriction1 := createRandomRestriction(t)

	newRestrictionNameSr:=util.RandomString(5)
	newRestrictionNameEn:=util.RandomString(5)
	newRestrictionNameBg:=util.RandomString(5)
	arg := UpdateRestrictionParams{
		ID:      restriction1.ID,
		RestrictionNameSr: pgtype.Text{
			String:  newRestrictionNameSr,
			Valid:  true,
		},
		RestrictionNameEn: pgtype.Text{
			String:  newRestrictionNameEn,
			Valid:  true,
		},
		RestrictionNameBg: pgtype.Text{
			String:  newRestrictionNameBg,
			Valid:  true,
		},

	}

	restriction2, err := testStore.UpdateRestriction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, restriction2)

	require.Equal(t, restriction1.ID, restriction2.ID)

	require.NotEqual(t, restriction1.RestrictionNameSr, restriction2.RestrictionNameSr)
	require.Equal(t, newRestrictionNameSr, restriction2.RestrictionNameSr)
	
	require.NotEqual(t, restriction1.RestrictionNameEn, restriction2.RestrictionNameEn)
	require.Equal(t, newRestrictionNameEn, restriction2.RestrictionNameEn)

	require.NotEqual(t, restriction1.RestrictionNameBg, restriction2.RestrictionNameBg)
	require.Equal(t, newRestrictionNameBg, restriction2.RestrictionNameBg)

	
	require.WithinDuration(t, restriction1.UpdatedAt, restriction2.UpdatedAt, time.Second)
}




func TestDeleteRestriction(t *testing.T) {
	restriction1 := createRandomRestriction(t)
	err := testStore.DeleteRestriction(context.Background(), restriction1.ID)
	require.NoError(t, err)

	restriction2, err := testStore.GetRestriction(context.Background(), restriction1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, restriction2)
}

func TestListRestrictions(t *testing.T) {
	arg := ListRestrictionsParams{
		Limit:  100,
		Offset: 0,
	}
	restrictions1, err := testStore.ListRestrictions(context.Background(), arg)
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		createRandomRestriction(t)
	}


	restrictions2, err := testStore.ListRestrictions(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, restrictions2)

	require.Equal(t, len(restrictions1)+5, len(restrictions2))
}