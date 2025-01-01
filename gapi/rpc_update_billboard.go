package gapi

import (
	"context"
	"database/sql"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateBillboard(ctx context.Context, req *pb.UpdateBillboardRequest) (*pb.UpdateBillboardResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateUpdateBillboardRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to update billboard")
	}

	store, err := s.repository.GetStore(ctx, req.GetStoreId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no store found")
	}

	args := db.UpdateBillboardParams{
		ID:      req.GetId(),
		StoreID: store.ID,
		Label: sql.NullString{
			Valid:  true,
			String: req.GetLabel(),
		},
		ImageUrl: sql.NullString{
			Valid:  true,
			String: req.GetImageUrl(),
		},
	}

	err = s.repository.UpdateBillboard(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to update billboard")
	}

	resp := &pb.UpdateBillboardResponse{
		Message: "Update successful",
	}

	return resp, nil
}

func validateUpdateBillboardRequest(req *pb.UpdateBillboardRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

	if err := validate.ValidateStoreName(req.GetLabel()); err != nil {
		violations = append(violations, fieldViolation("label", err))
	}

	if err := validate.ValidateUrl(req.GetImageUrl()); err != nil {
		violations = append(violations, fieldViolation("image_url", err))
	}

	return violations
}
