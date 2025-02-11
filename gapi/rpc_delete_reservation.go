package gapi

import (
	"context"

	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteReservation(ctx context.Context, req *pb.DeleteReservationRequest) (*pb.DeleteReservationResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	err = server.store.DeleteReservationTx(ctx, req.GetReservationId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete reservation: %s", err)
	}

	rsp := &pb.DeleteReservationResponse{
		Message: "the reservation was successfully deleted",
	}
	return rsp, nil
}
