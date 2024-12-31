package gapi

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetBillboard(ctx context.Context, req *pb.GetBillboardRequest) (*pb.GetBillboardResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateGetBillboardRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get billboard")
	}

	billboard, err := s.repository.GetBillboard(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get billboard")
	}

	resp := &pb.GetBillboardResponse{
		Billboard: convertBillboard(billboard),
	}

	return resp, nil
}

func validateGetBillboardRequest(req *pb.GetBillboardRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}
