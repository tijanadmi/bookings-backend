package gapi

import (
	"context"
	"time"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateRoomRestriction(ctx context.Context, req *pb.CreateRoomRestrictionRequest) (*pb.CreateRoomRestrictionResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	violations := validateCreateRoomRestrictionRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	layout := "2006-01-02" // format datuma
	startDate, err := time.Parse(layout, req.GetStartDate())
	if err != nil {
		violations = append(violations, fieldViolation("start_date", err))
		return nil, invalidArgumentError(violations)
	}
	endDate, err := time.Parse(layout, req.GetEndDate())
	if err != nil {
		violations = append(violations, fieldViolation("end_date", err))
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateRoomRestrictionParams{
		StartDate:     startDate,
		EndDate:       endDate,
		RoomID:        req.GetRoomId(),
		ReservationID: req.GetReservationId(),
		RestrictionID: req.GetRestrictionId(),
	}
	roomRestriction, err := server.store.CreateRoomRestriction(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create reservation: %s", err)
	}

	rsp := &pb.CreateRoomRestrictionResponse{
		Restriction: convertRoomRestriction(roomRestriction),
	}
	return rsp, nil
}

func validateCreateRoomRestrictionRequest(req *pb.CreateRoomRestrictionRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	// if err := val.ValidatePhone(req.GetPhone()); err != nil {
	// 	violations = append(violations, fieldViolation("phone", err))
	// }

	return violations
}
