package gapi

import (
	"context"
	"log"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteStore(ctx context.Context, req *pb.DeleteStoreRequest) (*pb.DeleteStoreResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}
	log.Print(req, " here")
	violations := validateDeleteStoreRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to delete store")
	}

	err = s.repository.DeleteStore(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to delete store")
	}

	resp := &pb.DeleteStoreResponse{
		Message: "store delete successful",
	}

	return resp, nil
}

func validateDeleteStoreRequest(req *pb.DeleteStoreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}
	return violations
}
