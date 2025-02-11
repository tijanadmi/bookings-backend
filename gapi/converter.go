package gapi

import (
	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func logResponse(rsp proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseProtoNames:   true, // Koristi originalna proto imena umesto camelCase
		EmitUnpopulated: true, // Prikazuje čak i default vrednosti (npr. `0`, `""`, `false`)
		Indent:          "  ", // Formatira JSON sa uvučenjima
	}

	jsonData, err := marshaler.Marshal(rsp)
	if err != nil {
		return "", err // Vraća prazan string i grešku
	}

	return string(jsonData), nil
}

func convertUser(user db.User) *pb.User {
	return &pb.User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Phone:       user.Phone,
		AccessLevel: user.AccessLevel,
		CreatedAt:   timestamppb.New(user.CreatedAt),
	}
}

func convertRoom(room db.Room) *pb.Room {
	return &pb.Room{
		RoomId:             room.ID,
		RoomNameSr:         room.RoomNameSr,
		RoomNameEn:         room.RoomNameSr,
		RoomNameBg:         room.RoomNameSr,
		RoomShortdesSr:     room.RoomShortDesSr,
		RoomShortdesEn:     room.RoomShortDesEn,
		RoomShortdesBg:     room.RoomShortDesBg,
		RoomDesSr:          room.RoomDescriptionSr,
		RoomDesEn:          room.RoomDescriptionEn,
		RoomDesBg:          room.RoomDescriptionBg,
		RoomPicturesFolder: room.RoomPicturesFolder,
		RoomGuestNumber:    room.RoomGuestNumber,
		RoomPriceEn:        room.RoomPriceEn,
		CreatedAt:          timestamppb.New(room.CreatedAt),
	}
}

func convertRestriction(restriction db.Restriction) *pb.Restriction {
	return &pb.Restriction{
		RestrictionId:     restriction.ID,
		RestrictionNameSr: restriction.RestrictionNameSr,
		RestrictionNameEn: restriction.RestrictionNameEn,
		RestrictionNameBg: restriction.RestrictionNameBg,
		CreatedAt:         timestamppb.New(restriction.CreatedAt),
	}
}

func convertReservation(reservation db.Reservation) *pb.Reservation {
	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")

	var numNights, numGuests, totalPrice *int32
	// Provera i konverzija NumNights
	if reservation.NumNights.Valid {
		nn := reservation.NumNights.Int32
		numNights = &nn
	}
	// Provera i konverzija NumGuests
	if reservation.NumGuests.Valid {
		nn := reservation.NumGuests.Int32
		numGuests = &nn
	}
	// Provera i konverzija TotalPrice
	if reservation.TotalPrice.Valid {
		nn := reservation.TotalPrice.Int32
		totalPrice = &nn
	}

	return &pb.Reservation{
		RoomId:       &reservation.RoomID,
		FirstName:    &reservation.FirstName,
		LastName:     &reservation.LastName,
		Email:        &reservation.Email,
		Phone:        &reservation.Phone,
		StartDate:    &startDate,
		EndDate:      &endDate,
		Processed:    &reservation.Processed,
		NumNights:    numNights,
		NumGuests:    numGuests,
		Status:       &reservation.Status,
		TotalPrice:   totalPrice,
		ExtrasPrice:  &reservation.ExtrasPrice,
		IsPaid:       &reservation.IsPaid,
		HasBreakfast: &reservation.HasBreakfast,
		CreatedAt:    timestamppb.New(reservation.CreatedAt),
	}
}

func convertReservationAll(reservation db.AllReservationsRow) *pb.ReservationAll {
	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")

	var roomNameSr, roomNameEn, roomNameBg *string
	var roomGuestNumber, roomPriceEn *int32

	// Provera i konverzija RoomNameSr
	if reservation.RoomNameSr.Valid {
		roomName := reservation.RoomNameSr.String
		roomNameSr = &roomName
	}

	// Provera i konverzija RoomNameEn
	if reservation.RoomNameEn.Valid {
		roomName := reservation.RoomNameEn.String
		roomNameEn = &roomName
	}

	// Provera i konverzija RoomNameBg
	if reservation.RoomNameBg.Valid {
		roomName := reservation.RoomNameBg.String
		roomNameBg = &roomName
	}

	// Provera i konverzija RoomGuestNumber
	if reservation.RoomGuestNumber.Valid {
		roomGN := reservation.RoomGuestNumber.Int32
		roomGuestNumber = &roomGN
	}

	// Provera i konverzija RoomPriceEn
	if reservation.RoomPriceEn.Valid {
		roomPN := reservation.RoomPriceEn.Int32
		roomPriceEn = &roomPN
	}

	var numNights, numGuests, totalPrice *int32
	// Provera i konverzija NumNights
	if reservation.NumNights.Valid {
		nn := reservation.NumNights.Int32
		numNights = &nn
	}
	// Provera i konverzija NumGuests
	if reservation.NumGuests.Valid {
		nn := reservation.NumGuests.Int32
		numGuests = &nn
	}
	// Provera i konverzija TotalPrice
	if reservation.TotalPrice.Valid {
		nn := reservation.TotalPrice.Int32
		totalPrice = &nn
	}

	return &pb.ReservationAll{
		ReservationId:   &reservation.ReservationID,
		RoomId:          &reservation.RoomID,
		RoomGuestNumber: roomGuestNumber,
		RoomPriceEn:     roomPriceEn,
		FirstName:       &reservation.FirstName,
		LastName:        &reservation.LastName,
		Email:           &reservation.Email,
		Phone:           &reservation.Phone,
		StartDate:       &startDate,
		EndDate:         &endDate,
		Processed:       &reservation.Processed,
		NumNights:       numNights,
		NumGuests:       numGuests,
		Status:          &reservation.Status,
		TotalPrice:      totalPrice,
		ExtrasPrice:     &reservation.ExtrasPrice,
		IsPaid:          &reservation.IsPaid,
		HasBreakfast:    &reservation.HasBreakfast,
		CreatedAt:       timestamppb.New(reservation.CreatedAt),
		RoomNameSr:      roomNameSr,
		RoomNameEn:      roomNameEn,
		RoomNameBg:      roomNameBg,
	}
}

func convertReservationNew(reservation db.AllNewReservationsRow) *pb.ReservationAll {
	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")

	var roomNameSr, roomNameEn, roomNameBg *string
	var roomGuestNumber, roomPriceEn *int32

	// Provera i konverzija RoomNameSr
	if reservation.RoomNameSr.Valid {
		roomName := reservation.RoomNameSr.String
		roomNameSr = &roomName
	}

	// Provera i konverzija RoomNameEn
	if reservation.RoomNameEn.Valid {
		roomName := reservation.RoomNameEn.String
		roomNameEn = &roomName
	}

	// Provera i konverzija RoomNameBg
	if reservation.RoomNameBg.Valid {
		roomName := reservation.RoomNameBg.String
		roomNameBg = &roomName
	}

	// Provera i konverzija RoomGuestNumber
	if reservation.RoomGuestNumber.Valid {
		roomGN := reservation.RoomGuestNumber.Int32
		roomGuestNumber = &roomGN
	}

	// Provera i konverzija RoomPriceEn
	if reservation.RoomPriceEn.Valid {
		roomPN := reservation.RoomPriceEn.Int32
		roomPriceEn = &roomPN
	}

	var numNights, numGuests, totalPrice *int32
	// Provera i konverzija NumNights
	if reservation.NumNights.Valid {
		nn := reservation.NumNights.Int32
		numNights = &nn
	}
	// Provera i konverzija NumGuests
	if reservation.NumGuests.Valid {
		nn := reservation.NumGuests.Int32
		numGuests = &nn
	}
	// Provera i konverzija TotalPrice
	if reservation.TotalPrice.Valid {
		nn := reservation.TotalPrice.Int32
		totalPrice = &nn
	}

	return &pb.ReservationAll{
		ReservationId:   &reservation.ReservationID,
		RoomId:          &reservation.RoomID,
		RoomGuestNumber: roomGuestNumber,
		RoomPriceEn:     roomPriceEn,
		FirstName:       &reservation.FirstName,
		LastName:        &reservation.LastName,
		Email:           &reservation.Email,
		Phone:           &reservation.Phone,
		StartDate:       &startDate,
		EndDate:         &endDate,
		Processed:       &reservation.Processed,
		NumNights:       numNights,
		NumGuests:       numGuests,
		Status:          &reservation.Status,
		TotalPrice:      totalPrice,
		ExtrasPrice:     &reservation.ExtrasPrice,
		IsPaid:          &reservation.IsPaid,
		HasBreakfast:    &reservation.HasBreakfast,
		CreatedAt:       timestamppb.New(reservation.CreatedAt),
		RoomNameSr:      roomNameSr,
		RoomNameEn:      roomNameEn,
		RoomNameBg:      roomNameBg,
	}
}

func convertReservationProcessed(reservation db.AllProcessedReservationsRow) *pb.ReservationAll {
	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")

	var roomNameSr, roomNameEn, roomNameBg *string
	var roomGuestNumber, roomPriceEn *int32

	// Provera i konverzija RoomNameSr
	if reservation.RoomNameSr.Valid {
		roomName := reservation.RoomNameSr.String
		roomNameSr = &roomName
	}

	// Provera i konverzija RoomNameEn
	if reservation.RoomNameEn.Valid {
		roomName := reservation.RoomNameEn.String
		roomNameEn = &roomName
	}

	// Provera i konverzija RoomNameBg
	if reservation.RoomNameBg.Valid {
		roomName := reservation.RoomNameBg.String
		roomNameBg = &roomName
	}

	// Provera i konverzija RoomGuestNumber
	if reservation.RoomGuestNumber.Valid {
		roomGN := reservation.RoomGuestNumber.Int32
		roomGuestNumber = &roomGN
	}

	// Provera i konverzija RoomPriceEn
	if reservation.RoomPriceEn.Valid {
		roomPN := reservation.RoomPriceEn.Int32
		roomPriceEn = &roomPN
	}

	var numNights, numGuests, totalPrice *int32
	// Provera i konverzija NumNights
	if reservation.NumNights.Valid {
		nn := reservation.NumNights.Int32
		numNights = &nn
	}
	// Provera i konverzija NumGuests
	if reservation.NumGuests.Valid {
		nn := reservation.NumGuests.Int32
		numGuests = &nn
	}
	// Provera i konverzija TotalPrice
	if reservation.TotalPrice.Valid {
		nn := reservation.TotalPrice.Int32
		totalPrice = &nn
	}

	return &pb.ReservationAll{
		ReservationId:   &reservation.ReservationID,
		RoomId:          &reservation.RoomID,
		RoomGuestNumber: roomGuestNumber,
		RoomPriceEn:     roomPriceEn,
		FirstName:       &reservation.FirstName,
		LastName:        &reservation.LastName,
		Email:           &reservation.Email,
		Phone:           &reservation.Phone,
		StartDate:       &startDate,
		EndDate:         &endDate,
		Processed:       &reservation.Processed,
		NumNights:       numNights,
		NumGuests:       numGuests,
		Status:          &reservation.Status,
		TotalPrice:      totalPrice,
		ExtrasPrice:     &reservation.ExtrasPrice,
		IsPaid:          &reservation.IsPaid,
		HasBreakfast:    &reservation.HasBreakfast,
		CreatedAt:       timestamppb.New(reservation.CreatedAt),
		RoomNameSr:      roomNameSr,
		RoomNameEn:      roomNameEn,
		RoomNameBg:      roomNameBg,
	}
}

func convertReservationParams(reservation db.ListReservationsResult) *pb.ReservationAll {
	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")

	return &pb.ReservationAll{
		ReservationId:   &reservation.ReservationID,
		RoomId:          &reservation.RoomID,
		RoomGuestNumber: &reservation.RoomGuestNumber,
		RoomPriceEn:     &reservation.RoomPriceEn,
		FirstName:       &reservation.FirstName,
		LastName:        &reservation.LastName,
		Email:           &reservation.Email,
		Phone:           &reservation.Phone,
		StartDate:       &startDate,
		EndDate:         &endDate,
		Processed:       &reservation.Processed,
		NumNights:       &reservation.NumNights,
		NumGuests:       &reservation.NumGuests,
		Status:          &reservation.Status,
		TotalPrice:      &reservation.TotalPrice,
		ExtrasPrice:     &reservation.ExtrasPrice,
		IsPaid:          &reservation.IsPaid,
		CreatedAt:       timestamppb.New(reservation.CreatedAt),
		RoomNameSr:      &reservation.RoomNameSr,
		RoomNameEn:      &reservation.RoomNameEn,
		RoomNameBg:      &reservation.RoomNameBg,
	}
}

func convertReservationAfterDate(reservation db.ListReservationsAfterDateRow) *pb.ReservationAll {
	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")

	var roomNameSr, roomNameEn, roomNameBg *string
	var roomGuestNumber, roomPriceEn *int32

	// Provera i konverzija RoomNameSr
	if reservation.RoomNameSr.Valid {
		roomName := reservation.RoomNameSr.String
		roomNameSr = &roomName
	}

	// Provera i konverzija RoomNameEn
	if reservation.RoomNameEn.Valid {
		roomName := reservation.RoomNameEn.String
		roomNameEn = &roomName
	}

	// Provera i konverzija RoomNameBg
	if reservation.RoomNameBg.Valid {
		roomName := reservation.RoomNameBg.String
		roomNameBg = &roomName
	}

	// Provera i konverzija RoomGuestNumber
	if reservation.RoomGuestNumber.Valid {
		roomGN := reservation.RoomGuestNumber.Int32
		roomGuestNumber = &roomGN
	}

	// Provera i konverzija RoomPriceEn
	if reservation.RoomPriceEn.Valid {
		roomPN := reservation.RoomPriceEn.Int32
		roomPriceEn = &roomPN
	}

	var numNights, numGuests, totalPrice *int32
	// Provera i konverzija NumNights
	if reservation.NumNights.Valid {
		nn := reservation.NumNights.Int32
		numNights = &nn
	}
	// Provera i konverzija NumGuests
	if reservation.NumGuests.Valid {
		nn := reservation.NumGuests.Int32
		numGuests = &nn
	}
	// Provera i konverzija TotalPrice
	if reservation.TotalPrice.Valid {
		nn := reservation.TotalPrice.Int32
		totalPrice = &nn
	}

	return &pb.ReservationAll{
		ReservationId:   &reservation.ReservationID,
		RoomId:          &reservation.RoomID,
		RoomGuestNumber: roomGuestNumber,
		RoomPriceEn:     roomPriceEn,
		FirstName:       &reservation.FirstName,
		LastName:        &reservation.LastName,
		Email:           &reservation.Email,
		Phone:           &reservation.Phone,
		StartDate:       &startDate,
		EndDate:         &endDate,
		Processed:       &reservation.Processed,
		NumNights:       numNights,
		NumGuests:       numGuests,
		Status:          &reservation.Status,
		TotalPrice:      totalPrice,
		ExtrasPrice:     &reservation.ExtrasPrice,
		IsPaid:          &reservation.IsPaid,
		HasBreakfast:    &reservation.HasBreakfast,
		CreatedAt:       timestamppb.New(reservation.CreatedAt),
		RoomNameSr:      roomNameSr,
		RoomNameEn:      roomNameEn,
		RoomNameBg:      roomNameBg,
	}
}

func convertStaysAfterDate(reservation db.ListStaysAfterDateRow) *pb.ReservationAll {
	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")

	var roomNameSr, roomNameEn, roomNameBg *string
	var roomGuestNumber, roomPriceEn *int32

	// Provera i konverzija RoomNameSr
	if reservation.RoomNameSr.Valid {
		roomName := reservation.RoomNameSr.String
		roomNameSr = &roomName
	}

	// Provera i konverzija RoomNameEn
	if reservation.RoomNameEn.Valid {
		roomName := reservation.RoomNameEn.String
		roomNameEn = &roomName
	}

	// Provera i konverzija RoomNameBg
	if reservation.RoomNameBg.Valid {
		roomName := reservation.RoomNameBg.String
		roomNameBg = &roomName
	}

	// Provera i konverzija RoomGuestNumber
	if reservation.RoomGuestNumber.Valid {
		roomGN := reservation.RoomGuestNumber.Int32
		roomGuestNumber = &roomGN
	}

	// Provera i konverzija RoomPriceEn
	if reservation.RoomPriceEn.Valid {
		roomPN := reservation.RoomPriceEn.Int32
		roomPriceEn = &roomPN
	}

	var numNights, numGuests, totalPrice *int32
	// Provera i konverzija NumNights
	if reservation.NumNights.Valid {
		nn := reservation.NumNights.Int32
		numNights = &nn
	}
	// Provera i konverzija NumGuests
	if reservation.NumGuests.Valid {
		nn := reservation.NumGuests.Int32
		numGuests = &nn
	}
	// Provera i konverzija TotalPrice
	if reservation.TotalPrice.Valid {
		nn := reservation.TotalPrice.Int32
		totalPrice = &nn
	}

	return &pb.ReservationAll{
		ReservationId:   &reservation.ReservationID,
		RoomId:          &reservation.RoomID,
		RoomGuestNumber: roomGuestNumber,
		RoomPriceEn:     roomPriceEn,
		FirstName:       &reservation.FirstName,
		LastName:        &reservation.LastName,
		Email:           &reservation.Email,
		Phone:           &reservation.Phone,
		StartDate:       &startDate,
		EndDate:         &endDate,
		Processed:       &reservation.Processed,
		NumNights:       numNights,
		NumGuests:       numGuests,
		Status:          &reservation.Status,
		TotalPrice:      totalPrice,
		ExtrasPrice:     &reservation.ExtrasPrice,
		IsPaid:          &reservation.IsPaid,
		HasBreakfast:    &reservation.HasBreakfast,
		CreatedAt:       timestamppb.New(reservation.CreatedAt),
		RoomNameSr:      roomNameSr,
		RoomNameEn:      roomNameEn,
		RoomNameBg:      roomNameBg,
	}
}



func convertRoomRestriction(roomRestriction db.RoomRestriction) *pb.RoomRestriction {
	startDate := roomRestriction.StartDate.Format("2006-01-02")
	endDate := roomRestriction.EndDate.Format("2006-01-02")

	var reservation_id *int32
	// Provera i konverzija RoomNameSr
	if roomRestriction.ReservationID.Valid {
		reservationId := roomRestriction.ReservationID.Int32
		reservation_id = &reservationId
	}
	return &pb.RoomRestriction{
		StartDate:     &startDate,
		EndDate:       &endDate,
		RoomId:        &roomRestriction.RoomID,
		ReservationId: reservation_id,
		RestrictionId: &roomRestriction.RestrictionID,
		CreatedAt:     timestamppb.New(roomRestriction.CreatedAt),
	}
}
