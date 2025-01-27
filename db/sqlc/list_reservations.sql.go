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
	RoomNameSr      string    `json:"room_name_sr"`
	RoomNameEn      string    `json:"room_name_en"`
	RoomNameBg      string    `json:"room_name_bg"`
}

// TransferTx performs a money transfer from one account to the other.
// It creates the transfer, add account entries, and update accounts' balance within a database transaction
func (store *SQLStore) ListReservationsWithParams(ctx context.Context, arg ReservationsWithParams) ([]ListReservationsResult, error) {

	// Provera sigurnosti unosa (samo dozvoljena polja)
	validOrderFields := map[string]bool{
		"start_date":    true,
		"end_date":      true,
		"created_at":    true,
		"updated_at":    true,
		"room_price_en": true,
	}
	if !validOrderFields[arg.OrderBy] {
		return nil, fmt.Errorf("invalid order by field")
	}

	// Provera pravca sortiranja
	if arg.OrderDir != "ASC" && arg.OrderDir != "DESC" {
		return nil, fmt.Errorf("invalid order direction")
	}

	// Dinamički WHERE uslov za processed
	whereClause := ""
	switch arg.Processed {
	case "0":
		whereClause = "WHERE r.processed = 0"
	case "1":
		whereClause = "WHERE r.processed = 1"
	case "all":
		whereClause = "" // Bez dodatnih uslova
	default:
		return nil, fmt.Errorf("invalid processed value")
	}

	// Dinamički SQL upit
	query := fmt.Sprintf(`
		SELECT r.id as reservation_id, rm.room_guest_number, rm.room_price_en, 
		       r.first_name, r.last_name, r.email, r.phone, 
		       r.start_date, r.end_date, r.room_id, r.created_at, r.updated_at, r.processed, 
		       rm.room_name_sr, rm.room_name_en, rm.room_name_bg
		FROM reservations r
		LEFT JOIN rooms rm ON (r.room_id = rm.id)
		%s
		ORDER BY %s %s
		LIMIT $1
		OFFSET $2;
	`, whereClause, arg.OrderBy, arg.OrderDir)

	// Izvršenje upita

	rows, err := store.db.Query(ctx, query, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListReservationsResult{}
	for rows.Next() {
		var i ListReservationsResult
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
			&i.RoomNameSr,
			&i.RoomNameEn,
			&i.RoomNameBg,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil

}
