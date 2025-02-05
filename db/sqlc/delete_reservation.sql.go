package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func (store *SQLStore) DeleteReservationTx(ctx context.Context, id int32) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		reservationID := pgtype.Int4{
			Int32: id,
			Valid: true,
		}

		err = q.DeleteRoomRestrictionForReservation(ctx, reservationID)
		if err != nil {
			return err
		}

		err = q.DeleteReservation(ctx, id)
		if err != nil {
			return err
		}

		return err
	})

	return err
}
