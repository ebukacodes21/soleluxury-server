package gapi

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetBillboards(ctx context.Context, req *pb.GetBillboardsRequest) (*pb.GetBillboardsResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateGetBillboardsRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get billboards")
	}

	billboards, err := s.repository.GetBillboards(ctx, req.GetStoreId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get billboards")
	}

	reversedBillboards := convertBillboards(billboards)
	for i, j := 0, len(reversedBillboards)-1; i < j; i, j = i+1, j-1 {
		reversedBillboards[i], reversedBillboards[j] = reversedBillboards[j], reversedBillboards[i]
	}

	resp := &pb.GetBillboardsResponse{
		Billboards: reversedBillboards,
	}

	return resp, nil
}

func validateGetBillboardsRequest(req *pb.GetBillboardsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

	return violations
}
