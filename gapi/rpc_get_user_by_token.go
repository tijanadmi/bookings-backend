package gapi

import (
	"context"

	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUserByToken(ctx context.Context, req *pb.GetUserByTokenRequest) (*pb.GetUserByTokenResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }
	accessToken := req.GetAccessToken()
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid access token: %s", err)
	}

	user, err := server.store.GetUser(ctx, payload.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %s", err)
	}

	rsp := &pb.GetUserByTokenResponse{
		User: convertUser(user),
	}
	return rsp, nil
}
