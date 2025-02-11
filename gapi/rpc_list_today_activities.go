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

func (server *Server) ListTodayActivities(ctx context.Context, req *pb.ListTodayActivitiesRequest) (*pb.ListTodayActivitiesResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	var violations []*errdetails.BadRequest_FieldViolation
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	layout := "2006-01-02" // format datuma
	startDate, err := time.Parse(layout, req.GetToday())
	if err != nil {
		violations = append(violations, fieldViolation("start_date", err))
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListTodayActivityParams{
		StartDate: startDate,
		EndDate:   startDate,
	}

	reservations, err := server.store.ListTodayActivity(ctx, arg)
	if err != nil {

		return nil, status.Errorf(codes.Internal, "failed to list reservations: %s", err)
	}

	rsp := &pb.ListTodayActivitiesResponse{
		Reservations: []*pb.ReservationAll{}, // Inicijalizacija kao prazan niz
	}

	for _, reservation := range reservations {
		rsp.Reservations = append(rsp.Reservations, convertStaysAfterDate(db.ListStaysAfterDateRow(reservation)))
	}

	return rsp, nil
}
