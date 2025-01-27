package gapi

import (
	"context"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListProcessedReservationsParams(ctx context.Context, req *pb.ListReservationsParamsRequest) (*pb.ListReservationsParamsResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }

	arg := db.ReservationsWithParams{
		OrderBy:   req.GetOrderBy(),
		OrderDir:  req.GetOrderDir(),
		Processed: req.GetProcessed(),
		Limit:     req.GetLimit(),
		Offset:    req.GetOffset(),
	}

	// ListReservationsWithParams(ctx context.Context, arg ReservationsWithParams) ([]ListReservationsResult, error)
	reservations, err := server.store.ListReservationsWithParams(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list rooms: %s", err)
	}

	rsp := &pb.ListReservationsParamsResponse{}
	for _, reservation := range reservations {
		rsp.Reservations = append(rsp.Reservations, convertReservationParams(reservation))
	}
	return rsp, nil
}
