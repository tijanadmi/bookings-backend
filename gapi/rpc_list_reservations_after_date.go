package gapi

import (
	"context"
	"fmt"
	"time"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListReservationsAfterDate(ctx context.Context, req *pb.ListReservationsAfterDateRequest) (*pb.ListReservationsAfterDateResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	var violations []*errdetails.BadRequest_FieldViolation
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

	if startDate.After(endDate) {
		violations = append(violations, fieldViolation("start_date", fmt.Errorf("start date cannot be after end date")))
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListReservationsAfterDateParams{
		CreatedAt:   startDate,
		CreatedAt_2: endDate,
	}

	reservations, err := server.store.ListReservationsAfterDate(ctx, arg)
	if err != nil {

		return nil, status.Errorf(codes.Internal, "failed to list reservations: %s", err)
	}

	rsp := &pb.ListReservationsAfterDateResponse{
		Reservations: []*pb.ReservationAll{}, // Inicijalizacija kao prazan niz
	}

	for _, reservation := range reservations {
		rsp.Reservations = append(rsp.Reservations, convertReservationAfterDate(reservation))
	}

	return rsp, nil
}
