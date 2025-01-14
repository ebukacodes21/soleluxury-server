package gapi

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// create billboard
func (s *Server) CreateBillboard(ctx context.Context, req *pb.CreateBillboardRequest) (*pb.CreateBillboardResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to create billboard")
	}

	violations := validateCreateBillboardRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	billboard, err := s.repository.CreateBillboard(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to create billboard %s", err)
	}

	resp := &pb.CreateBillboardResponse{
		Billboard: convertBillboard(billboard),
	}

	return resp, nil
}

// get billboard
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

	billboard, err := s.repository.GetBillboard(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get billboard")
	}

	resp := &pb.GetBillboardResponse{
		Billboard: convertBillboard(billboard),
	}

	return resp, nil
}

// get billboards
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

	billboards, err := s.repository.GetAllBillboards(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get billboards %s", err)
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

// update billboard
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

	message, err := s.repository.UpdateBillboard(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to update billboard")
	}

	resp := &pb.UpdateBillboardResponse{
		Message: message,
	}

	return resp, nil
}

// delete billboard
func (s *Server) DeleteBillboard(ctx context.Context, req *pb.DeleteBillboardRequest) (*pb.DeleteBillboardResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateDeleteBillboardRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to delete billboard")
	}

	message, err := s.repository.DeleteBillboard(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to delete billboard %s ", err)
	}

	resp := &pb.DeleteBillboardResponse{
		Message: message,
	}

	return resp, nil
}

func validateCreateBillboardRequest(req *pb.CreateBillboardRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

	if err := validate.ValidateName(req.GetLabel()); err != nil {
		violations = append(violations, fieldViolation("label", err))
	}

	if err := validate.ValidateUrl(req.GetImageUrl()); err != nil {
		violations = append(violations, fieldViolation("image_url", err))
	}

	return violations
}

func validateGetBillboardRequest(req *pb.GetBillboardRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}

func validateGetBillboardsRequest(req *pb.GetBillboardsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

	return violations
}

func validateUpdateBillboardRequest(req *pb.UpdateBillboardRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	if err := validate.ValidateName(req.GetLabel()); err != nil {
		violations = append(violations, fieldViolation("label", err))
	}

	if err := validate.ValidateUrl(req.GetImageUrl()); err != nil {
		violations = append(violations, fieldViolation("image_url", err))
	}

	return violations
}

func validateDeleteBillboardRequest(req *pb.DeleteBillboardRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}
	return violations
}
