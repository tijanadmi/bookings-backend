package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/tijanadmi/bookings_backend/util"
)

func createRandomRoom(t *testing.T) Room {

	arg := CreateRoomParams{
		RoomNameSr:  util.RandomString(5),
		RoomNameEn:  util.RandomString(5),
		RoomNameBg:  util.RandomString(5),
		RoomShortDesSr:  util.RandomString(20),
		RoomShortDesEn:  util.RandomString(20),
		RoomShortDesBg:  util.RandomString(20),
		RoomDescriptionSr:  util.RandomString(200),
		RoomDescriptionEn:  util.RandomString(200),
		RoomDescriptionBg:  util.RandomString(200),
		RoomPicturesFolder:  util.RandomString(20),
		RoomGuestNumber:  int32(util.RandomInt(1,4)),
		RoomPriceEn:  int32(util.RandomInt(10,40)),
	}

	room, err := testStore.CreateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room)

	require.Equal(t, arg.RoomNameSr, room.RoomNameSr)
	require.Equal(t, arg.RoomNameEn, room.RoomNameEn)
	require.Equal(t, arg.RoomNameBg, room.RoomNameBg)
	require.Equal(t, arg.RoomShortDesSr, room.RoomShortDesSr)
	require.Equal(t, arg.RoomShortDesEn, room.RoomShortDesEn)
	require.Equal(t, arg.RoomShortDesBg, room.RoomShortDesBg)
	require.Equal(t, arg.RoomDescriptionSr, room.RoomDescriptionSr)
	require.Equal(t, arg.RoomDescriptionEn, room.RoomDescriptionEn)
	require.Equal(t, arg.RoomDescriptionBg, room.RoomDescriptionBg)
	require.Equal(t, arg.RoomPicturesFolder, room.RoomPicturesFolder)
	require.Equal(t, arg.RoomGuestNumber, room.RoomGuestNumber)
	require.Equal(t, arg.RoomPriceEn, room.RoomPriceEn)
	require.NotZero(t, room.ID)
	require.NotZero(t, room.CreatedAt)

	return room
}

func TestCreateRoom(t *testing.T) {
	createRandomRoom(t)
}

func TestGetRoom(t *testing.T) {
	room1 := createRandomRoom(t)
	room2, err := testStore.GetRoom(context.Background(),room1.ID)
	require.NoError(t, err)
	require.NotEmpty(t,room2)

	require.Equal(t,room1.ID,room2.ID)
	require.Equal(t, room1.RoomNameSr, room2.RoomNameSr)
	require.Equal(t, room1.RoomNameEn, room2.RoomNameEn)
	require.Equal(t, room1.RoomNameBg, room2.RoomNameBg)
	require.Equal(t, room1.RoomShortDesSr, room2.RoomShortDesSr)
	require.Equal(t, room1.RoomShortDesEn, room2.RoomShortDesEn)
	require.Equal(t, room1.RoomShortDesBg, room2.RoomShortDesBg)
	require.Equal(t, room1.RoomDescriptionSr, room2.RoomDescriptionSr)
	require.Equal(t, room1.RoomDescriptionEn, room2.RoomDescriptionEn)
	require.Equal(t, room1.RoomDescriptionBg, room2.RoomDescriptionBg)
	require.Equal(t, room1.RoomPicturesFolder, room2.RoomPicturesFolder)
	require.Equal(t, room1.RoomGuestNumber, room2.RoomGuestNumber)
	require.Equal(t, room1.RoomPriceEn, room2.RoomPriceEn)
	
	require.WithinDuration(t, room1.CreatedAt, room2.CreatedAt, time.Second)
}

func TestUpdateRoomPriceEn(t *testing.T) {
	room1 := createRandomRoom(t)

	newRoomPriceEn:=int32(util.RandomInt(10, 40))
	arg := UpdateRoomParams{
		ID:      room1.ID,
		RoomPriceEn: pgtype.Int4{
			Int32:  newRoomPriceEn,
			Valid:  true,
		},

	}

	room2, err := testStore.UpdateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room2)

	require.Equal(t, room1.ID, room2.ID)
	require.Equal(t, room1.RoomNameSr, room2.RoomNameSr)
	require.Equal(t, room1.RoomNameEn, room2.RoomNameEn)
	require.Equal(t, room1.RoomNameBg, room2.RoomNameBg)
	require.Equal(t, room1.RoomShortDesSr, room2.RoomShortDesSr)
	require.Equal(t, room1.RoomShortDesEn, room2.RoomShortDesEn)
	require.Equal(t, room1.RoomShortDesBg, room2.RoomShortDesBg)
	require.Equal(t, room1.RoomDescriptionSr, room2.RoomDescriptionSr)
	require.Equal(t, room1.RoomDescriptionEn, room2.RoomDescriptionEn)
	require.Equal(t, room1.RoomDescriptionBg, room2.RoomDescriptionBg)
	require.Equal(t, room1.RoomPicturesFolder, room2.RoomPicturesFolder)


	require.NotEqual(t, room1.RoomPriceEn, room2.RoomPriceEn)
	require.Equal(t, newRoomPriceEn, room2.RoomPriceEn)
	require.Equal(t, room1.RoomGuestNumber, room2.RoomGuestNumber)
	
	require.WithinDuration(t, room1.UpdatedAt, room2.UpdatedAt, time.Second)
}

func TestUpdateRoomNameSr(t *testing.T) {
	room1 := createRandomRoom(t)

	newRoomNameSr:=util.RandomString(5)
	arg := UpdateRoomParams{
		ID:      room1.ID,
		RoomNameSr: pgtype.Text{
			String:  newRoomNameSr,
			Valid:  true,
		},
	}

	room2, err := testStore.UpdateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room2)

	require.Equal(t, room1.ID, room2.ID)
	
	require.Equal(t, room1.RoomNameEn, room2.RoomNameEn)
	require.Equal(t, room1.RoomNameBg, room2.RoomNameBg)
	require.Equal(t, room1.RoomShortDesSr, room2.RoomShortDesSr)
	require.Equal(t, room1.RoomShortDesEn, room2.RoomShortDesEn)
	require.Equal(t, room1.RoomShortDesBg, room2.RoomShortDesBg)
	require.Equal(t, room1.RoomDescriptionSr, room2.RoomDescriptionSr)
	require.Equal(t, room1.RoomDescriptionEn, room2.RoomDescriptionEn)
	require.Equal(t, room1.RoomDescriptionBg, room2.RoomDescriptionBg)
	require.Equal(t, room1.RoomPicturesFolder, room2.RoomPicturesFolder)
	require.Equal(t, room1.RoomPriceEn, room2.RoomPriceEn)


	require.NotEqual(t, room1.RoomNameSr, room2.RoomNameSr)
	require.Equal(t, newRoomNameSr, room2.RoomNameSr)

	require.Equal(t, room1.RoomGuestNumber, room2.RoomGuestNumber)
	
	require.WithinDuration(t, room1.UpdatedAt, room2.UpdatedAt, time.Second)
}

func TestUpdateRoomMoreThanOneField(t *testing.T) {
	room1 := createRandomRoom(t)

	newRoomNameSr:=util.RandomString(5)
	newRoomShortDesSr:=util.RandomString(25)
	newRoomDescriptionSr:=util.RandomString(100)
	newRoomPriceEn:=int32(util.RandomInt(10, 40))

	arg := UpdateRoomParams{
		ID:      room1.ID,
		RoomNameSr: pgtype.Text{
			String:  newRoomNameSr,
			Valid:  true,
		},
		RoomShortDesSr: pgtype.Text{
			String:  newRoomShortDesSr,
			Valid:  true,
		},
		RoomDescriptionSr: pgtype.Text{
			String:  newRoomDescriptionSr,
			Valid:  true,
		},
		RoomPriceEn: pgtype.Int4{
			Int32:  newRoomPriceEn,
			Valid:  true,
		},

	}

	room2, err := testStore.UpdateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room2)

	require.Equal(t, room1.ID, room2.ID)
	
	require.Equal(t, room1.RoomNameEn, room2.RoomNameEn)
	require.Equal(t, room1.RoomNameBg, room2.RoomNameBg)

	require.Equal(t, room1.RoomShortDesEn, room2.RoomShortDesEn)
	require.Equal(t, room1.RoomShortDesBg, room2.RoomShortDesBg)

	require.Equal(t, room1.RoomDescriptionEn, room2.RoomDescriptionEn)
	require.Equal(t, room1.RoomDescriptionBg, room2.RoomDescriptionBg)
	require.Equal(t, room1.RoomPicturesFolder, room2.RoomPicturesFolder)
	


	require.NotEqual(t, room1.RoomNameSr, room2.RoomNameSr)
	require.Equal(t, newRoomNameSr, room2.RoomNameSr)

	require.NotEqual(t, room1.RoomShortDesSr, room2.RoomShortDesSr)
	require.Equal(t, newRoomShortDesSr, room2.RoomShortDesSr)

	require.NotEqual(t, room1.RoomDescriptionSr, room2.RoomDescriptionSr)
	require.Equal(t, newRoomDescriptionSr, room2.RoomDescriptionSr)

	require.NotEqual(t, room1.RoomPriceEn, room2.RoomPriceEn)
	require.Equal(t, newRoomPriceEn, room2.RoomPriceEn)

	require.Equal(t, room1.RoomGuestNumber, room2.RoomGuestNumber)
	
	require.WithinDuration(t, room1.UpdatedAt, room2.UpdatedAt, time.Second)
}
func TestDeleteRoom(t *testing.T) {
	room1 := createRandomRoom(t)
	err := testStore.DeleteRoom(context.Background(), room1.ID)
	require.NoError(t, err)

	room2, err := testStore.GetRoom(context.Background(), room1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, room2)
}

func TestListRooms(t *testing.T) {
	arg := ListRoomsParams{
		Limit:  100,
		Offset: 0,
	}
	rooms1, err := testStore.ListRooms(context.Background(), arg)
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		createRandomRoom(t)
	}


	rooms2, err := testStore.ListRooms(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, rooms2)

	require.Equal(t, len(rooms1)+5, len(rooms2))
}