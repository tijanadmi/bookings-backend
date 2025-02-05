package db

import (
	"context"
	"fmt"
	"time"
)

// ReservationsWithParams contains the input parameters of the dinamic query for listing reservations
type ReservationsWithParams struct {
	OrderBy   string `json:"order_by"`
	OrderDir  string `json:"order_dir"`
	Processed string `json:"processed"`
	Status    string `json:"status"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

// ListReservationsResult is the result of the dinamic query for listing reservations
type ListReservationsResult struct {
	ReservationID   int32     `json:"reservation_id"`
	RoomGuestNumber int32     `json:"room_guest_number"`
	RoomPriceEn     int32     `json:"room_price_en"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	RoomID          int32     `json:"room_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Processed       int32     `json:"processed"`
	NumNights       int32     `json:"num_nights"`
	NumGuests       int32     `json:"num_guests"`
	Status          string    `json:"status"`
	TotalPrice      int32     `json:"total_price"`
	ExtrasPrice     int32     `json:"extras_price"`
	IsPaid          bool      `json:"is_paid"`
	RoomNameSr      string    `json:"room_name_sr"`
	RoomNameEn      string    `json:"room_name_en"`
	RoomNameBg      string    `json:"room_name_bg"`
}


func (store *SQLStore) ListReservationsWithParams(ctx context.Context, arg ReservationsWithParams) ([]ListReservationsResult, int64, error) {

	// Provera sigurnosti unosa (samo dozvoljena polja)
	validOrderFields := map[string]bool{
		"start_date":    true,
		"end_date":      true,
		"created_at":    true,
		"updated_at":    true,
		"room_price_en": true,
	}
	if !validOrderFields[arg.OrderBy] {
		return nil, 0, fmt.Errorf("invalid order by field")
	}

	// Provera pravca sortiranja
	if arg.OrderDir != "ASC" && arg.OrderDir != "DESC" {
		return nil, 0, fmt.Errorf("invalid order direction")
	}

	// Dinamički WHERE uslov za processed
	whereClause := ""
	if arg.Processed == "0" {
		whereClause = "WHERE r.processed = 0"
	} else if arg.Processed == "1" {
		whereClause = "WHERE r.processed = 1"
	}

	// Dodavanje uslova za status
	if arg.Status != "all" {
		if whereClause == "" {
			whereClause = "WHERE "
		} else {
			whereClause += " AND "
		}

		switch arg.Status {
		case "checked-in":
			whereClause += " r.status = 'checked-in'"
		case "checked-out":
			whereClause += " r.status = 'checked-out'"
		case "unconfirmed":
			whereClause += " r.status = 'unconfirmed'"
		default:
			return nil, 0, fmt.Errorf("invalid status value")
		}
	}

	// Dinamički SQL upit sa brojanjem ukupnog broja redova
	query := fmt.Sprintf(`
		WITH reservations_with_count AS (
			SELECT r.id as reservation_id, rm.room_guest_number, rm.room_price_en, 
			       r.first_name, r.last_name, r.email, r.phone, 
			       r.start_date, r.end_date, r.room_id, r.created_at, r.updated_at, r.processed, 
				   r.num_nights, r.num_guests, r.status, r.total_price, r.extras_price, r.is_paid,
			       rm.room_name_sr, rm.room_name_en, rm.room_name_bg,
			       COUNT(*) OVER () as total_count
			FROM reservations r
			LEFT JOIN rooms rm ON (r.room_id = rm.id)
			%s
		)
		SELECT reservation_id, room_guest_number, room_price_en, 
		       first_name, last_name, email, phone, 
		       start_date, end_date, room_id, created_at, updated_at, processed, 
			   num_nights, num_guests, status, total_price, extras_price, is_paid,
		       room_name_sr, room_name_en, room_name_bg, total_count
		FROM reservations_with_count
		ORDER BY %s %s
		LIMIT $1
		OFFSET $2;
	`, whereClause, arg.OrderBy, arg.OrderDir)

	// Izvršenje upita
	rows, err := store.db.Query(ctx, query, arg.Limit, arg.Offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := []ListReservationsResult{}
	var totalCount int64

	for rows.Next() {
		var i ListReservationsResult
		var count int64
		if err := rows.Scan(
			&i.ReservationID,
			&i.RoomGuestNumber,
			&i.RoomPriceEn,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Processed,
			&i.NumNights,
			&i.NumGuests,
			&i.Status,
			&i.TotalPrice,
			&i.ExtrasPrice,
			&i.IsPaid,
			&i.RoomNameSr,
			&i.RoomNameEn,
			&i.RoomNameBg,
			&count,
		); err != nil {
			return nil, 0, err
		}
		items = append(items, i)
		totalCount = count // totalCount će biti isti za svaki red zbog funkcije COUNT(*) OVER ()

	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return items, totalCount, nil

}
