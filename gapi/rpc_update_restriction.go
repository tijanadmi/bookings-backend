package gapi

import (
	"context"
	"time"

	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"github.com/tijanadmi/bookings_backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateRestriction(ctx context.Context, req *pb.UpdateRestrictionRequest) (*pb.UpdateRestrictionResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	violations := validateUpdateRestrictionRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateRestrictionParams{
		ID: req.GetRestrictionId(),
		RestrictionNameSr: pgtype.Text{
			String: req.GetRestrictionNameSr(),
			Valid:  req.RestrictionNameSr != nil,
		},
		RestrictionNameEn: pgtype.Text{
			String: req.GetRestrictionNameEn(),
			Valid:  req.RestrictionNameEn != nil,
		},
		RestrictionNameBg: pgtype.Text{
			String: req.GetRestrictionNameBg(),
			Valid:  req.RestrictionNameBg != nil,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}

	restriction, err := server.store.UpdateRestriction(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "restriction not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update restriction: %s", err)
	}

	rsp := &pb.UpdateRestrictionResponse{
		Restriction: convertRestriction(restriction),
	}
	return rsp, nil
}

func validateUpdateRestrictionRequest(req *pb.UpdateRestrictionRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.RestrictionNameSr != nil {
		if err := val.ValidateString(req.GetRestrictionNameSr(), 1, 50); err != nil {
			violations = append(violations, fieldViolation("restriction_name_sr", err))
		}
	}

	if req.RestrictionNameEn != nil {
		if err := val.ValidateString(req.GetRestrictionNameEn(), 1, 50); err != nil {
			violations = append(violations, fieldViolation("restriction_name_en", err))
		}
	}
	if req.RestrictionNameBg != nil {
		if err := val.ValidateString(req.GetRestrictionNameBg(), 1, 50); err != nil {
			violations = append(violations, fieldViolation("restriction_name_bg", err))
		}
	}

	return violations
}
