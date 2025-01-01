package gapi

import (
	"context"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBillboard(ctx context.Context, req *pb.CreateBillboardRequest) (*pb.CreateBillboardResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateCreateBillboardRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to create billboard")
	}

	store, err := s.repository.GetStore(ctx, req.GetStoreId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no store found")
	}

	args := db.CreateBillboardParams{
		StoreID:  store.ID,
		Label:    req.GetLabel(),
		ImageUrl: req.GetImageUrl(),
	}

	billboard, err := s.repository.CreateBillboard(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create billboard")
	}

	resp := &pb.CreateBillboardResponse{
		Billboard: convertBillboard(billboard),
	}

	return resp, nil
}

func validateCreateBillboardRequest(req *pb.CreateBillboardRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
