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
		Password:    user.Password,
		AccessLevel: user.AccessLevel,
		CreatedAt:   timestamppb.New(user.CreatedAt),
	}
}

func convertRoom(room db.Room) *pb.Room {
	return &pb.Room{
		RoomNameSr:         room.RoomNameSr,
		RoomNameEn:         room.RoomNameSr,
		RoomNameBg:         room.RoomNameSr,
		RoomShortdesSr:     room.RoomShortDesSr,
		RoomShortdesEn:     room.RoomShortDesEn,
		RoomShortdescBg:    room.RoomShortDesBg,
		RoomDesSr:          room.RoomDescriptionSr,
		RoomDesEn:          room.RoomDescriptionEn,
		RoomDescBg:         room.RoomDescriptionBg,
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
