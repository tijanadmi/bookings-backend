package gapi

import (
	"context"

	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetRestriction(ctx context.Context, req *pb.GetRestrictionRequest) (*pb.GetRestrictionResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	restriction, err := server.store.GetRestriction(ctx, req.GetRestrictionId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get room: %s", err)
	}

	rsp := &pb.GetRestrictionResponse{
		Restriction: convertRestriction(restriction),
	}
	return rsp, nil
}
