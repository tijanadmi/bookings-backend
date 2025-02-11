package gapi

import (
	"context"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListProcessedReservations(ctx context.Context, req *pb.ListProcessedReservationsRequest) (*pb.ListProcessedReservationsResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	arg := db.AllProcessedReservationsParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	reservations, err := server.store.AllProcessedReservations(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list rooms: %s", err)
	}

	rsp := &pb.ListProcessedReservationsResponse{}
	for _, reservation := range reservations {
		rsp.Reservations = append(rsp.Reservations, convertReservationProcessed(reservation))
	}
	return rsp, nil
}
