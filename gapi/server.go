package gapi

import (
	"fmt"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"github.com/tijanadmi/bookings_backend/token"
	"github.com/tijanadmi/bookings_backend/util"
)

// Server serves gRPC requests for booking service.
type Server struct {
	pb.UnimplementedBookingsServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey/*, config.Issuer, config.Audience, config.AccessTokenDuration, config.RefreshTokenDuration, config.CookieDomain, config.CookiePath, config.CookieName*/)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
