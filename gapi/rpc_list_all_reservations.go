package gapi

import (
	"context"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListAllReservations(ctx context.Context, req *pb.ListAllReservationsRequest) (*pb.ListAllReservationsResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }

	arg := db.AllReservationsParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	reservations, err := server.store.AllReservations(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list rooms: %s", err)
	}

	rsp := &pb.ListAllReservationsResponse{}
	for _, reservation := range reservations {
		rsp.ReservationAll = append(rsp.ReservationAll, convertReservationAll(reservation))
	}
	return rsp, nil
}
