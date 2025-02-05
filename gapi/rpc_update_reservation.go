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

func (server *Server) UpdateReservation(ctx context.Context, req *pb.UpdateReservationRequest) (*pb.UpdateReservationResponse, error) {
	// _, err := server.authorizeUser(ctx)
	// if err != nil {
	// 	return nil, unauthenticatedError(err)
	// }
	violations := validateUpdateReservationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	// layout := "2006-01-02" // format datuma
	// startDate, err := time.Parse(layout, req.GetStartDate())
	// if err != nil {
	// 	violations = append(violations, fieldViolation("start_date", err))
	// 	return nil, invalidArgumentError(violations)
	// }
	// endDate, err := time.Parse(layout, req.GetEndDate())
	// if err != nil {
	// 	violations = append(violations, fieldViolation("end_date", err))
	// 	return nil, invalidArgumentError(violations)
	// }

	arg := db.UpdateReservationParams{
		ID: req.GetReservationId(),
		RoomID: pgtype.Int4{
			Int32: req.GetRoomId(),
			Valid: req.RoomId != nil,
		},
		FirstName: pgtype.Text{
			String: req.GetFirstName(),
			Valid:  req.FirstName != nil,
		},
		LastName: pgtype.Text{
			String: req.GetLastName(),
			Valid:  req.LastName != nil,
		},
		Email: pgtype.Text{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		Phone: pgtype.Text{
			String: req.GetPhone(),
			Valid:  req.Phone != nil,
		},
		// StartDate: pgtype.Date{
		// 	Time:  startDate, // Konverzija u pgtype.Date
		// 	Valid: true,
		// },
		// EndDate: pgtype.Date{
		// 	Time:  endDate, // Konverzija u pgtype.Date
		// 	Valid: true,
		// },
		Processed: pgtype.Int4{
			Int32: req.GetProcessed(),
			Valid: req.Processed != nil,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
		NumNights: pgtype.Int4{
			Int32: req.GetNumNights(),
			Valid: req.NumNights != nil,
		},
		NumGuests: pgtype.Int4{
			Int32: req.GetNumGuests(),
			Valid: req.NumGuests != nil,
		},
		Status: pgtype.Text{
			String: req.GetStatus(),
			Valid:  req.Status != nil,
		},
		TotalPrice: pgtype.Int4{
			Int32: req.GetTotalPrice(),
			Valid: req.TotalPrice != nil,
		},
		ExtrasPrice: pgtype.Int4{
			Int32: req.GetExtrasPrice(),
			Valid: req.ExtrasPrice != nil,
		},
		IsPaid: pgtype.Bool{
			Bool:  req.GetIsPaid(),
			Valid: req.IsPaid != nil,
		},
		HasBreakfast: pgtype.Bool{
			Bool:  req.GetHasBreakfast(),
			Valid: req.HasBreakfast != nil,
		},
	}

	layout := "2006-01-02" // format datuma
	if req.StartDate != nil {

		startDate, err := time.Parse(layout, req.GetStartDate())
		if err != nil {
			violations = append(violations, fieldViolation("start_date", err))
			return nil, invalidArgumentError(violations)
		}
		arg.StartDate = pgtype.Date{
			Time:  startDate, // Konverzija u pgtype.Date
			Valid: true,
		}
	}
	if req.StartDate != nil {
		endDate, err := time.Parse(layout, req.GetEndDate())
		if err != nil {
			violations = append(violations, fieldViolation("end_date", err))
			return nil, invalidArgumentError(violations)
		}
		arg.EndDate = pgtype.Date{
			Time:  endDate, // Konverzija u pgtype.Date
			Valid: true,
		}
	}

	reservation, err := server.store.UpdateReservation(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "reservation not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update reservation: %s", err)
	}

	rsp := &pb.UpdateReservationResponse{
		Reservation: convertReservation(reservation),
	}
	return rsp, nil
}

func validateUpdateReservationRequest(req *pb.UpdateReservationRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.FirstName != nil {
		if err := val.ValidateName(req.GetFirstName()); err != nil {
			violations = append(violations, fieldViolation("first_name", err))
		}
	}

	if req.LastName != nil {
		if err := val.ValidateName(req.GetLastName()); err != nil {
			violations = append(violations, fieldViolation("last_name", err))
		}
	}

	if req.Email != nil {
		if err := val.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	if req.Phone != nil {
		if err := val.ValidatePhone(req.GetPhone()); err != nil {
			violations = append(violations, fieldViolation("phone", err))
		}
	}
	return violations

}
