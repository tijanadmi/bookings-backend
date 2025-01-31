package gapi

import (
	"context"

	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteRestriction(ctx context.Context, req *pb.DeleteRestrictionRequest) (*pb.DeleteRestrictionResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }

	err := server.store.DeleteRestriction(ctx, req.GetRestrictionId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete restriction: %s", err)
	}

	rsp := &pb.DeleteRestrictionResponse{
		Message: "the restriction was successfully deleted",
	}
	return rsp, nil
}
