package gapi

import (
	"context"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListNewReservations(ctx context.Context, req *pb.ListNewReservationsRequest) (*pb.ListNewReservationsResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }

	arg := db.AllNewReservationsParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	reservations, err := server.store.AllNewReservations(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list rooms: %s", err)
	}

	rsp := &pb.ListNewReservationsResponse{}
	for _, reservation := range reservations {
		rsp.ReservationNew = append(rsp.ReservationNew, convertReservationNew(reservation))
	}
	return rsp, nil
}
