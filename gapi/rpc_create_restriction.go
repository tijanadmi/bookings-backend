package gapi

import (
	"context"

	"github.com/lib/pq"
	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"github.com/tijanadmi/bookings_backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateRestriction(ctx context.Context, req *pb.CreateRestrictionRequest) (*pb.CreateRestrictionResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }
	violations := validateCreateRestrictionRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateRestrictionParams{
		RestrictionNameSr: req.GetRestrictionNameSr(),
		RestrictionNameEn: req.GetRestrictionNameEn(),
		RestrictionNameBg: req.GetRestrictionNameBg(),
	}
	restriction, err := server.store.CreateRestriction(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "restriction name alrady exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create restriction: %s", err)
	}

	rsp := &pb.CreateRestrictionResponse{
		Restriction: convertRestriction(restriction),
	}
	return rsp, nil
}

func validateCreateRestrictionRequest(req *pb.CreateRestrictionRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetRestrictionNameSr(), 1, 50); err != nil {
		violations = append(violations, fieldViolation("restriction_name_sr", err))
	}

	if err := val.ValidateString(req.GetRestrictionNameEn(), 1, 50); err != nil {
		violations = append(violations, fieldViolation("restriction_name_en", err))
	}

	if err := val.ValidateString(req.GetRestrictionNameBg(), 1, 50); err != nil {
		violations = append(violations, fieldViolation("restriction_name_bg", err))
	}

	return violations
}
