package gapi

import (
	"context"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListRooms(ctx context.Context, req *pb.ListRoomsRequest) (*pb.ListRoomsResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	arg := db.ListRoomsParams{
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	}
	rooms, err := server.store.ListRooms(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list rooms: %s", err)
	}

	rsp := &pb.ListRoomsResponse{}
	for _, room := range rooms {
		rsp.Rooms = append(rsp.Rooms, convertRoom(room))
	}
	return rsp, nil
}
