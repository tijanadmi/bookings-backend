package gapi

import (
	"context"
	"fmt"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListReservationsWithParams(ctx context.Context, req *pb.ListReservationsParamsRequest) (*pb.ListReservationsParamsResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }

	violations := validateListReservationsParamsRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ReservationsWithParams{
		OrderBy:   req.GetOrderBy(),
		OrderDir:  req.GetOrderDir(),
		Processed: req.GetProcessed(),
		Limit:     req.GetLimit(),
		Offset:    req.GetOffset(),
	}

	reservations, total, err := server.store.ListReservationsWithParams(ctx, arg)
	if err != nil {

		return nil, status.Errorf(codes.Internal, "failed to list rooms: %s", err)
	}

	// Mapiranje rezultata u gRPC odgovor
	rsp := &pb.ListReservationsParamsResponse{
		Total: total,
	}
	for _, reservation := range reservations {
		rsp.Reservations = append(rsp.Reservations, convertReservationParams(reservation))
	}

	return rsp, nil
}

func validateListReservationsParamsRequest(req *pb.ListReservationsParamsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	arg := db.ReservationsWithParams{
		OrderBy:   req.GetOrderBy(),
		OrderDir:  req.GetOrderDir(),
		Processed: req.GetProcessed(),
		Limit:     req.GetLimit(),
		Offset:    req.GetOffset(),
	}

	// Provera sigurnosti unosa (samo dozvoljena polja)
	validOrderFields := map[string]bool{
		"start_date":    true,
		"end_date":      true,
		"created_at":    true,
		"updated_at":    true,
		"room_price_en": true,
	}
	if !validOrderFields[req.GetOrderBy()] {
		violations = append(violations, fieldViolation("order_by", fmt.Errorf("invalid order by param")))
	}

	// Provera pravca sortiranja
	if req.GetOrderDir() != "ASC" && req.GetOrderDir() != "DESC" {

		violations = append(violations, fieldViolation("order_dir", fmt.Errorf("invalid order diretion param")))
	}

	if arg.Processed != "0" && arg.Processed != "1" && arg.Processed != "all" {
		violations = append(violations, fieldViolation("processed", fmt.Errorf("invalid processed param")))

	}

	return violations
}
