package gapi

import (
	"context"

	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"github.com/tijanadmi/bookings_backend/util"
	"github.com/tijanadmi/bookings_backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
// 	fmt.Println(req)
// 	violations := validateLoginUserRequest(req)
// 	if violations != nil {
// 		return nil, invalidArgumentError(violations)
// 	}

// 	user, err := server.store.GetUser(ctx, req.GetUsername())
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to find user")
// 	}

// 	err = util.CheckPassword(req.Password, user.Password)
// 	if err != nil {
// 		return nil, status.Errorf(codes.NotFound, "incorrect password")
// 	}

// 	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
// 		user.Email,
// 		string(user.AccessLevel),
// 		server.config.AccessTokenDuration,
// 	)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to create access token")
// 	}

// 	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
// 		user.Email,
// 		string(user.AccessLevel),
// 		server.config.RefreshTokenDuration,
// 	)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to create refresh token")
// 	}

// 	mtdt := server.extractMetadata(ctx)
// 	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
// 		ID:           refreshPayload.ID,
// 		Username:     user.Email,
// 		RefreshToken: refreshToken,
// 		UserAgent:    mtdt.UserAgent,
// 		ClientIp:     mtdt.ClientIP,
// 		IsBlocked:    false,
// 		ExpiresAt:    refreshPayload.ExpiredAt,
// 	})
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to create session")
// 	}

// 	rsp := &pb.LoginUserResponse{
// 		SessionId:             session.ID.String(),
// 		AccessToken:           accessToken,
// 		RefreshToken:          refreshToken,
// 		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
// 		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
// 	}
// 	return rsp, nil
// }

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	// fmt.Println(req)
	violations := validateLoginUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect password")
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

	mtdt := server.extractMetadata(ctx)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Email,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session")
	}

	rsp := &pb.LoginUserResponse{
		User:         convertUser(user),
		SessionId:    session.ID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		// AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt.UTC()),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt.UTC()),
	}

	// fmt.Println("accessPayload.ExpiredAt", accessPayload.ExpiredAt)
	// fmt.Println("timestamppb.New(accessPayload.ExpiredAt.UTC())", timestamppb.New(accessPayload.ExpiredAt.UTC()))

	// fmt.Println("UTC Now:", time.Now().UTC())
	// fmt.Println("Local Now:", time.Now())

	/*** create cookie and set into metadata ***/
	refreshCookie := util.GetRefreshCookie(refreshToken, server.config.CookieName, server.config.CookiePath, server.config.CookieDomain, server.config.RefreshTokenDuration)
	// Postavite metadata sa `Set-Cookie`
	md := metadata.Pairs("Set-Cookie", refreshCookie.String())
	// fmt.Println(md)

	//grpc.SetHeader(ctx, md) // Postavite header kroz `grpc-gateway`
	// Dodavanje `Set-Cookie` u gRPC metapodatke
	grpc.SendHeader(ctx, md)
	return rsp, nil
}

func validateLoginUserRequest(req *pb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmail(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}
