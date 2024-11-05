package gapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/tijanadmi/bookings_backend/pb"
	"github.com/tijanadmi/bookings_backend/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	// Extract the refresh token from the cookie in metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}
	fmt.Println("Metadata:", md)

	cookies := md.Get("set-cookie")
	if len(cookies) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "refresh token not found in cookies")
	}

	fmt.Println("Cookies:", cookies)

	// Parse the cookies to extract the refresh token
	var refreshToken string
	for _, cookie := range cookies {
		// Split cookie into key-value pairs
		pairs := strings.Split(cookie, "; ")
		for _, pair := range pairs {
			parts := strings.SplitN(pair, "=", 2)
			if len(parts) == 2 && parts[0] == server.config.CookieName {
				refreshToken = parts[1]
				fmt.Println("Found refresh token:", refreshToken)
				break
			}
		}

	}

	if refreshToken == "" {
		return nil, status.Errorf(codes.Unauthenticated, "refresh token not found")
	}

	payload, err := server.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	// Retrieve user from the database

	user, err := server.store.GetUser(ctx, payload.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Email,
		string(user.AccessLevel),
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token")
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Email,
		string(user.AccessLevel),
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token")
	}

	// Set refresh cookie in metadata
	refreshCookie := util.GetRefreshCookie(refreshToken, server.config.CookieName, server.config.CookiePath, server.config.CookieDomain, server.config.RefreshTokenDuration)
	mdOut := metadata.Pairs("Set-Cookie", refreshCookie.String())
	grpc.SendHeader(ctx, mdOut)

	// Prepare and return the response
	rsp := &pb.RefreshTokenResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}
