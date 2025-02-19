// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AllNewReservations(ctx context.Context, arg AllNewReservationsParams) ([]AllNewReservationsRow, error)
	AllProcessedReservations(ctx context.Context, arg AllProcessedReservationsParams) ([]AllProcessedReservationsRow, error)
	AllReservations(ctx context.Context, arg AllReservationsParams) ([]AllReservationsRow, error)
	CreateReservation(ctx context.Context, arg CreateReservationParams) (Reservation, error)
	CreateRestriction(ctx context.Context, arg CreateRestrictionParams) (Restriction, error)
	CreateRoom(ctx context.Context, arg CreateRoomParams) (Room, error)
	CreateRoomRestriction(ctx context.Context, arg CreateRoomRestrictionParams) (RoomRestriction, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteReservation(ctx context.Context, id int32) error
	DeleteRestriction(ctx context.Context, id int32) error
	DeleteRoom(ctx context.Context, id int32) error
	DeleteRoomRestriction(ctx context.Context, id int32) error
	DeleteRoomRestrictionForReservation(ctx context.Context, reservationID pgtype.Int4) error
	GetReservation(ctx context.Context, id int32) (Reservation, error)
	GetReservationByID(ctx context.Context, id int32) (GetReservationByIDRow, error)
	GetReservationForUpdate(ctx context.Context, id int32) (Reservation, error)
	GetRestriction(ctx context.Context, id int32) (Restriction, error)
	GetRestrictionForUpdate(ctx context.Context, id int32) (Restriction, error)
	GetRestrictionsForRoomByDate(ctx context.Context, arg GetRestrictionsForRoomByDateParams) ([]RoomRestriction, error)
	GetRoom(ctx context.Context, id int32) (Room, error)
	GetRoomForUpdate(ctx context.Context, id int32) (Room, error)
	GetRoomRestriction(ctx context.Context, id int32) (RoomRestriction, error)
	GetRoomRestrictionForUpdate(ctx context.Context, id int32) (RoomRestriction, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, email string) (User, error)
	ListReservations(ctx context.Context, arg ListReservationsParams) ([]Reservation, error)
	ListReservationsAfterDate(ctx context.Context, arg ListReservationsAfterDateParams) ([]ListReservationsAfterDateRow, error)
	ListRestrictions(ctx context.Context, arg ListRestrictionsParams) ([]Restriction, error)
	ListRoomRestrictions(ctx context.Context, arg ListRoomRestrictionsParams) ([]RoomRestriction, error)
	ListRooms(ctx context.Context, arg ListRoomsParams) ([]Room, error)
	ListStaysAfterDate(ctx context.Context, arg ListStaysAfterDateParams) ([]ListStaysAfterDateRow, error)
	ListTodayActivity(ctx context.Context, arg ListTodayActivityParams) ([]ListTodayActivityRow, error)
	SearchAvailabilityByDatesByRoomID(ctx context.Context, arg SearchAvailabilityByDatesByRoomIDParams) (int64, error)
	SearchAvailabilityForAllRooms(ctx context.Context, arg SearchAvailabilityForAllRoomsParams) ([]Room, error)
	UpdateProcessedForReservation(ctx context.Context, arg UpdateProcessedForReservationParams) (Reservation, error)
	UpdateReservation(ctx context.Context, arg UpdateReservationParams) (Reservation, error)
	UpdateRestriction(ctx context.Context, arg UpdateRestrictionParams) (Restriction, error)
	UpdateRoom(ctx context.Context, arg UpdateRoomParams) (Room, error)
	UpdateRoomRestriction(ctx context.Context, arg UpdateRoomRestrictionParams) (RoomRestriction, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
