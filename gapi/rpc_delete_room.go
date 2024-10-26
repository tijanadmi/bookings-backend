package gapi

import (
	"context"

	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteRoom(ctx context.Context, req *pb.DeleteRoomRequest) (*pb.DeleteRoomResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }

	err := server.store.DeleteRoom(ctx, req.GetRoomId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete room: %s", err)
	}

	rsp := &pb.DeleteRoomResponse{
		Message: "the room was successfully deleted",
	}
	return rsp, nil
}
