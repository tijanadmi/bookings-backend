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
