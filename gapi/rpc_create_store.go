package gapi

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateStore(ctx context.Context, req *pb.CreateStoreRequest) (*pb.CreateStoreResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to create store")
	}

	violations := validateCreateStoreRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	store, err := s.repository.CreateStore(ctx, req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to create store %s", err)
	}

	resp := &pb.CreateStoreResponse{
		Store: convertStore(store),
	}

	return resp, nil
}

func validateCreateStoreRequest(req *pb.CreateStoreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateStoreName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	return violations
}
