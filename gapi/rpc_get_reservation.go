package gapi

import (
	"context"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetReservation(ctx context.Context, req *pb.GetReservationRequest) (*pb.GetReservationResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }

	reservation, err := server.store.GetReservationByID(ctx, req.GetReservationId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get reservation: %s", err)
	}

	rsp := &pb.GetReservationResponse{
		Reservation: convertReservationAll(db.AllReservationsRow(reservation)),
	}
	return rsp, nil
}
