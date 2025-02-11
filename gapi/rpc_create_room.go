package gapi

import (
	"context"

	"github.com/lib/pq"
	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"github.com/tijanadmi/bookings_backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	// log.Println(req)

	violations := validateCreateRoomRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateRoomParams{
		RoomNameSr:         req.GetRoomNameSr(),
		RoomNameEn:         req.GetRoomNameEn(),
		RoomNameBg:         req.GetRoomNameBg(),
		RoomShortDesSr:     req.GetRoomShortdesSr(),
		RoomShortDesEn:     req.GetRoomShortdesEn(),
		RoomShortDesBg:     req.GetRoomShortdesBg(),
		RoomDescriptionSr:  req.GetRoomDesSr(),
		RoomDescriptionEn:  req.GetRoomDesEn(),
		RoomDescriptionBg:  req.GetRoomDesBg(),
		RoomPicturesFolder: req.GetRoomPicturesFolder(),
		RoomGuestNumber:    req.GetRoomGuestNumber(),
		RoomPriceEn:        req.GetRoomPriceEn(),
	}
	// log.Println(req)
	// log.Println(arg)

	room, err := server.store.CreateRoom(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "room name alrady exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create room: %s", err)
	}

	rsp := &pb.CreateRoomResponse{
		Room: convertRoom(room),
	}
	return rsp, nil
}

func validateCreateRoomRequest(req *pb.CreateRoomRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetRoomNameSr(), 1, 50); err != nil {
		violations = append(violations, fieldViolation("room_name_sr", err))
	}

	if err := val.ValidateString(req.GetRoomNameEn(), 1, 50); err != nil {
		violations = append(violations, fieldViolation("room_name_en", err))
	}

	if err := val.ValidateString(req.GetRoomNameBg(), 1, 50); err != nil {
		violations = append(violations, fieldViolation("room_name_bg", err))
	}

	if err := val.ValidateString(req.GetRoomShortdesSr(), 1, 200); err != nil {
		violations = append(violations, fieldViolation("room_short_des_sr", err))
	}

	if err := val.ValidateString(req.GetRoomShortdesEn(), 1, 200); err != nil {
		violations = append(violations, fieldViolation("room_short_des_en", err))
	}

	if err := val.ValidateString(req.GetRoomShortdesBg(), 1, 200); err != nil {
		violations = append(violations, fieldViolation("room_short_des_bg", err))
	}

	if err := val.ValidateString(req.GetRoomDesSr(), 1, 1000); err != nil {
		violations = append(violations, fieldViolation("room_description_sr", err))
	}

	if err := val.ValidateString(req.GetRoomDesEn(), 1, 1000); err != nil {
		violations = append(violations, fieldViolation("room_description_en", err))
	}

	if err := val.ValidateString(req.GetRoomDesBg(), 1, 1000); err != nil {
		violations = append(violations, fieldViolation("room_description_bg", err))
	}

	if err := val.ValidateString(req.GetRoomPicturesFolder(), 1, 200); err != nil {
		violations = append(violations, fieldViolation("room_pictures_folder", err))
	}

	return violations
}
