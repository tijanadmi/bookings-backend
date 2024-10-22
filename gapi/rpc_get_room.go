package gapi

import (
	"context"

	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetRoom(ctx context.Context, req *pb.GetRoomRequest) (*pb.GetRoomResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }

	room, err := server.store.GetRoom(ctx, req.GetRoomId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get room: %s", err)
	}

	rsp := &pb.GetRoomResponse{
		Room: convertRoom(room),
	}
	return rsp, nil
}
