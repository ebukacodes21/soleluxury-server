package gapi

import (
	"context"
	"database/sql"
	"log"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateStore(ctx context.Context, req *pb.UpdateStoreRequest) (*pb.UpdateStoreResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateUpdateStoreRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	log.Print(req)
	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to update store")
	}

	args := db.UpdateStoreParams{
		ID: req.GetId(),
		Name: sql.NullString{
			Valid:  true,
			String: req.GetName(),
		},
	}

	err = s.repository.UpdateStore(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to update store")
	}

	resp := &pb.UpdateStoreResponse{
		Message: "store update successful",
	}

	return resp, nil
}

func validateUpdateStoreRequest(req *pb.UpdateStoreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	if err := validate.ValidateStoreName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	return violations
}
