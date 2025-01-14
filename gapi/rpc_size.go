package gapi

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// create size
func (s *Server) CreateSize(ctx context.Context, req *pb.CreateSizeRequest) (*pb.CreateSizeResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateCreateSizeRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to create size")
	}

	size, err := s.repository.CreateSize(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create size %s", err)
	}

	resp := &pb.CreateSizeResponse{
		Size: convertSize(size),
	}

	return resp, nil

}

// get size
func (s *Server) GetSize(ctx context.Context, req *pb.GetSizeRequest) (*pb.GetSizeResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateGetSizeRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get size")
	}

	size, err := s.repository.GetSize(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get size %s", err)
	}

	resp := &pb.GetSizeResponse{
		Size: convertSize(size),
	}

	return resp, nil

}

// get sizes
func (s *Server) GetSizes(ctx context.Context, req *pb.GetSizesRequest) (*pb.GetSizesResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateGetSizesRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get sizes")
	}

	sizes, err := s.repository.GetAllSizes(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get sizes %s", err)
	}

	reversedSizes := convertSizes(sizes)
	for i, j := 0, len(reversedSizes)-1; i < j; i, j = i+1, j-1 {
		reversedSizes[i], reversedSizes[j] = reversedSizes[j], reversedSizes[i]
	}

	resp := &pb.GetSizesResponse{
		Sizes: reversedSizes,
	}

	return resp, nil

}

func (s *Server) UpdateSize(ctx context.Context, req *pb.UpdateSizeRequest) (*pb.UpdateSizeResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateUpdateSizeRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to update size")
	}

	message, err := s.repository.UpdateSize(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to update size %s", err)
	}

	resp := &pb.UpdateSizeResponse{
		Message: message,
	}

	return resp, nil

}

func (s *Server) DeleteSize(ctx context.Context, req *pb.DeleteSizeRequest) (*pb.DeleteSizeResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateDeleteSizeRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to delete size")
	}

	message, err := s.repository.DeleteSize(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to delete size %s", err)
	}

	resp := &pb.DeleteSizeResponse{
		Message: message,
	}

	return resp, nil

}

func validateCreateSizeRequest(req *pb.CreateSizeRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

	if err := validate.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	if err := validate.ValidateValue(req.GetValue()); err != nil {
		violations = append(violations, fieldViolation("value", err))
	}

	return violations
}

func validateGetSizeRequest(req *pb.GetSizeRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}

func validateGetSizesRequest(req *pb.GetSizesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

	return violations
}

func validateUpdateSizeRequest(req *pb.UpdateSizeRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}
	if err := validate.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	if err := validate.ValidateValue(req.GetValue()); err != nil {
		violations = append(violations, fieldViolation("value", err))
	}

	return violations
}

func validateDeleteSizeRequest(req *pb.DeleteSizeRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}
