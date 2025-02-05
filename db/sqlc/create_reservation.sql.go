package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func (store *SQLStore) CreateReservationTx(ctx context.Context, arg CreateReservationParams) (Reservation, error) {
	var reservation = Reservation{}

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		count, err := q.SearchAvailabilityByDatesByRoomID(ctx, SearchAvailabilityByDatesByRoomIDParams{
			RoomID:    arg.RoomID,
			EndDate:   arg.StartDate,
			StartDate: arg.EndDate,
		})

		fmt.Println(count)

		if err != nil {
			return err
		}
		if count != 0 {
			return fmt.Errorf("for the given period, a reservation already exists for the room")
		}

		reservation, err = q.CreateReservation(ctx, arg)
		if err != nil {
			return err
		}
		reservationId := pgtype.Int4{
			Int32: reservation.ID,
			Valid: true,
		}

		_, err = q.CreateRoomRestriction(ctx, CreateRoomRestrictionParams{
			StartDate:     arg.StartDate,
			EndDate:       arg.EndDate,
			RoomID:        arg.RoomID,
			ReservationID: reservationId,
			RestrictionID: 1})
		if err != nil {
			return err
		}

		return err
	})

	return reservation, err
}
