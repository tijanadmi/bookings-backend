package gapi

import (
	"context"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListRestrictions(ctx context.Context, req *pb.ListRestrictionsRequest) (*pb.ListRestrictionsResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	arg := db.ListRestrictionsParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	restrictions, err := server.store.ListRestrictions(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list restrictions: %s", err)
	}

	rsp := &pb.ListRestrictionsResponse{}
	for _, restriction := range restrictions {
		rsp.Restrictions = append(rsp.Restrictions, convertRestriction(restriction))
	}
	return rsp, nil
}
