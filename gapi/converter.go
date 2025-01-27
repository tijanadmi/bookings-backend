package gapi

import (
	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
		RestrictionNameSr: restriction.RestrictionNameSr,
		RestrictionNameEn: restriction.RestrictionNameSr,
		RestrictionNameBg: restriction.RestrictionNameSr,
		CreatedAt:         timestamppb.New(restriction.CreatedAt),
	}
}

func convertReservation(reservation db.Reservation) *pb.Reservation {
	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")
	return &pb.Reservation{
		RoomId:    &reservation.RoomID,
		FirstName: &reservation.FirstName,
		LastName:  &reservation.LastName,
		Email:     &reservation.Email,
		Phone:     &reservation.Phone,
		StartDate: &startDate,
		EndDate:   &endDate,
		Processed: &reservation.Processed,
		CreatedAt: timestamppb.New(reservation.CreatedAt),
	}
}

func convertReservationAll(reservation db.AllReservationsRow) *pb.ReservationAll {
	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")

	var roomNameSr, roomNameEn, roomNameBg *string

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

	return &pb.ReservationAll{
		RoomId:     &reservation.RoomID,
		FirstName:  &reservation.FirstName,
		LastName:   &reservation.LastName,
		Email:      &reservation.Email,
		Phone:      &reservation.Phone,
		StartDate:  &startDate,
		EndDate:    &endDate,
		Processed:  &reservation.Processed,
		CreatedAt:  timestamppb.New(reservation.CreatedAt),
		RoomNameSr: roomNameSr,
		RoomNameEn: roomNameEn,
		RoomNameBg: roomNameBg,
	}
}

func convertReservationNew(reservation db.AllNewReservationsRow) *pb.ReservationAll {
	startDate := reservation.StartDate.Format("2006-01-02")
	endDate := reservation.EndDate.Format("2006-01-02")

	var roomNameSr, roomNameEn, roomNameBg *string

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

	return &pb.ReservationAll{
		RoomId:     &reservation.RoomID,
		FirstName:  &reservation.FirstName,
		LastName:   &reservation.LastName,
		Email:      &reservation.Email,
		Phone:      &reservation.Phone,
		StartDate:  &startDate,
		EndDate:    &endDate,
		Processed:  &reservation.Processed,
		CreatedAt:  timestamppb.New(reservation.CreatedAt),
		RoomNameSr: roomNameSr,
		RoomNameEn: roomNameEn,
		RoomNameBg: roomNameBg,
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

