package gapi

import (
	"context"

	"github.com/lib/pq"
	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	"github.com/tijanadmi/bookings_backend/pb"
	"github.com/tijanadmi/bookings_backend/util"
	"github.com/tijanadmi/bookings_backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		FirstName:   req.GetFirstName(),
		LastName:    req.GetLastName(),
		Email:       req.GetEmail(),
		Phone:       req.GetPhone(),
		Password:    hashedPassword,
		AccessLevel: 1,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username alrady exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}
	// arg := db.CreateUserTxParams{
	// 	CreateUserParams: db.CreateUserParams{
	// 		FirstName:   req.GetFirstName(),
	// 		LastName:    req.GetLastName(),
	// 		Email:       req.GetEmail(),
	// 		Phone:       req.GetPhone(),
	// 		Password:    hashedPassword,
	// 		AccessLevel: 1,
	// 	},
	// 	AfterCreate: func(user db.User) error {
	// 		taskPayload := &worker.PayloadSendVerifyEmail{
	// 			Username: user.Username,
	// 		}
	// 		opts := []asynq.Option{
	// 			asynq.MaxRetry(10),
	// 			asynq.ProcessIn(10 * time.Second),
	// 			asynq.Queue(worker.QueueCritical),
	// 		}

	// 		return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
	// 	},
	// }

	// txResult, err := server.store.CreateUserTx(ctx, arg)
	// if err != nil {
	// 	if db.ErrorCode(err) == db.UniqueViolation {
	// 		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	// 	}
	// 	return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	// }

	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateName(req.GetFirstName()); err != nil {
		violations = append(violations, fieldViolation("first_name", err))
	}

	if err := val.ValidateName(req.GetLastName()); err != nil {
		violations = append(violations, fieldViolation("last_name", err))
	}

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := val.ValidatePhone(req.GetPhone()); err != nil {
		violations = append(violations, fieldViolation("phone", err))
	}

	return violations
}
