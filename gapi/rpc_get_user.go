package gapi

import (
	"context"

	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	user, err := server.store.GetUser(ctx, req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get room: %s", err)
	}

	rsp := &pb.GetUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}
