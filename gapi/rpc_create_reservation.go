package gapi

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"github.com/tijanadmi/bookings_backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateReservation(ctx context.Context, req *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	violations := validateCreateReservationRequest(req)
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

	arg := db.CreateReservationParams{
		RoomID:    req.GetRoomId(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Phone:     req.GetEmail(),
		StartDate: startDate,
		EndDate:   endDate,
		Processed: req.GetProcessed(),
		NumNights: pgtype.Int4{
			Int32: req.GetNumNights(),
			Valid: true,
		},
		NumGuests: pgtype.Int4{
			Int32: req.GetNumGuests(),
			Valid: true,
		},
		Status: req.GetStatus(),
		TotalPrice: pgtype.Int4{
			Int32: req.GetTotalPrice(),
			Valid: true,
		},
		ExtrasPrice:  req.GetExtrasPrice(),
		IsPaid:       req.GetIsPaid(),
		HasBreakfast: req.GetHasBreakfast(),
	}
	reservation, err := server.store.CreateReservationTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create reservation: %s", err)
	}

	rsp := &pb.CreateReservationResponse{
		Reservation: convertReservation(reservation),
	}
	return rsp, nil
}

func validateCreateReservationRequest(req *pb.CreateReservationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateName(req.GetFirstName()); err != nil {
		violations = append(violations, fieldViolation("first_name", err))
	}

	if err := val.ValidateName(req.GetLastName()); err != nil {
		violations = append(violations, fieldViolation("last_name", err))
	}

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if err := val.ValidatePhone(req.GetPhone()); err != nil {
		violations = append(violations, fieldViolation("phone", err))
	}

	return violations
}
