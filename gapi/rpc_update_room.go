package gapi

import (
	"context"
	"time"

	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"github.com/tijanadmi/bookings_backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateRoom(ctx context.Context, req *pb.UpdateRoomRequest) (*pb.UpdateRoomResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }
	violations := validateUpdateRoomRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateRoomParams{
		ID: req.GetRoomId(),
		RoomNameSr: pgtype.Text{
			String: req.GetRoomNameSr(),
			Valid:  req.RoomNameSr != nil,
		},
		RoomNameEn: pgtype.Text{
			String: req.GetRoomNameEn(),
			Valid:  req.RoomNameEn != nil,
		},
		RoomNameBg: pgtype.Text{
			String: req.GetRoomNameBg(),
			Valid:  req.RoomNameBg != nil,
		},
		RoomShortDesSr: pgtype.Text{
			String: req.GetRoomShortdesSr(),
			Valid:  req.RoomShortdesSr != nil,
		},
		RoomShortDesEn: pgtype.Text{
			String: req.GetRoomShortdesEn(),
			Valid:  req.RoomShortdesEn != nil,
		},
		RoomShortDesBg: pgtype.Text{
			String: req.GetRoomShortdescBg(),
			Valid:  req.RoomShortdescBg != nil,
		},
		RoomDescriptionSr: pgtype.Text{
			String: req.GetRoomDesSr(),
			Valid:  req.RoomDesSr != nil,
		},
		RoomDescriptionEn: pgtype.Text{
			String: req.GetRoomDesEn(),
			Valid:  req.RoomDesEn != nil,
		},
		RoomDescriptionBg: pgtype.Text{
			String: req.GetRoomDescBg(),
			Valid:  req.RoomDescBg != nil,
		},
		RoomPicturesFolder: pgtype.Text{
			String: req.GetRoomPicturesFolder(),
			Valid:  req.RoomPicturesFolder != nil,
		},
		RoomGuestNumber: pgtype.Int4{
			Int32: req.GetRoomGuestNumber(),
			Valid: req.RoomGuestNumber != nil,
		},
		RoomPriceEn: pgtype.Int4{
			Int32: req.GetRoomPriceEn(),
			Valid: req.RoomPriceEn != nil,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}

	room, err := server.store.UpdateRoom(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "room not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update room: %s", err)
	}

	rsp := &pb.UpdateRoomResponse{
		Room: convertRoom(room),
	}
	return rsp, nil
}

func validateUpdateRoomRequest(req *pb.UpdateRoomRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.RoomNameSr != nil {
		if err := val.ValidateString(req.GetRoomNameSr(), 1, 50); err != nil {
			violations = append(violations, fieldViolation("room_name_sr", err))
		}
	}

	if req.RoomNameEn != nil {
		if err := val.ValidateString(req.GetRoomNameEn(), 1, 50); err != nil {
			violations = append(violations, fieldViolation("room_name_en", err))
		}
	}
	if req.RoomNameBg != nil {
		if err := val.ValidateString(req.GetRoomNameBg(), 1, 50); err != nil {
			violations = append(violations, fieldViolation("room_name_bg", err))
		}
	}

	if req.RoomShortdesSr != nil {
		if err := val.ValidateString(req.GetRoomShortdesSr(), 1, 200); err != nil {
			violations = append(violations, fieldViolation("room_short_des_sr", err))
		}
	}
	if req.RoomShortdesEn != nil {
		if err := val.ValidateString(req.GetRoomShortdesEn(), 1, 200); err != nil {
			violations = append(violations, fieldViolation("room_short_des_en", err))
		}
	}
	if req.RoomShortdescBg != nil {
		if err := val.ValidateString(req.GetRoomShortdescBg(), 1, 200); err != nil {
			violations = append(violations, fieldViolation("room_short_des_bg", err))
		}
	}
	if req.RoomDesSr != nil {
		if err := val.ValidateString(req.GetRoomDesSr(), 1, 1000); err != nil {
			violations = append(violations, fieldViolation("room_description_sr", err))
		}
	}
	if req.RoomDesEn != nil {
		if err := val.ValidateString(req.GetRoomDesEn(), 1, 1000); err != nil {
			violations = append(violations, fieldViolation("room_description_en", err))
		}
	}
	if req.RoomDescBg != nil {
		if err := val.ValidateString(req.GetRoomDescBg(), 1, 1000); err != nil {
			violations = append(violations, fieldViolation("room_description_bg", err))
		}
	}
	if req.RoomPicturesFolder != nil {
		if err := val.ValidateString(req.GetRoomPicturesFolder(), 1, 200); err != nil {
			violations = append(violations, fieldViolation("room_pictures_folder", err))
		}
	}

	return violations
}
