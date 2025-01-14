package gapi

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// create store
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

	store, err := s.repository.CreateStore(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to create store %s", err)
	}

	resp := &pb.CreateStoreResponse{
		Store: convertStore(store),
	}

	return resp, nil
}

// get first store
func (s *Server) GetFirstStore(ctx context.Context, req *emptypb.Empty) (*pb.GetStoreResponse, error) {
	store, err := s.repository.GetFirstStore(ctx)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to get first store %s", err)
	}

	resp := &pb.GetStoreResponse{
		Store: convertStore(store),
	}

	return resp, nil
}

// get store
func (s *Server) GetStore(ctx context.Context, req *pb.GetStoreRequest) (*pb.GetStoreResponse, error) {
	violations := validateGetStoreRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	store, err := s.repository.GetStoreByID(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to get store %s", err)
	}

	resp := &pb.GetStoreResponse{
		Store: convertStore(store),
	}

	return resp, nil
}

// get all stores
func (s *Server) GetStores(ctx context.Context, req *emptypb.Empty) (*pb.GetStoresResponse, error) {
	stores, err := s.repository.GetAllStores(ctx)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to get stores %s", err)
	}

	resp := &pb.GetStoresResponse{
		Stores: convertStores(stores),
	}

	return resp, nil
}

// update store
func (s *Server) UpdateStore(ctx context.Context, req *pb.UpdateStoreRequest) (*pb.UpdateStoreResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateUpdateStoreRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to update store")
	}

	message, err := s.repository.UpdateStore(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to update store %s ", err)
	}

	resp := &pb.UpdateStoreResponse{
		Message: message,
	}

	return resp, nil
}

// delete store
func (s *Server) DeleteStore(ctx context.Context, req *pb.DeleteStoreRequest) (*pb.DeleteStoreResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateDeleteStoreRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to delete store")
	}

	message, err := s.repository.DeleteStore(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to delete store %s ", err)
	}

	resp := &pb.DeleteStoreResponse{
		Message: message,
	}

	return resp, nil
}

// validate create store request
func validateCreateStoreRequest(req *pb.CreateStoreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	return violations
}

func validateGetStoreRequest(req *pb.GetStoreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}

func validateUpdateStoreRequest(req *pb.UpdateStoreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	if req.Name != nil {
		if err := validate.ValidateName(req.GetName()); err != nil {
			violations = append(violations, fieldViolation("name", err))
		}
	}
	return violations
}

func validateDeleteStoreRequest(req *pb.DeleteStoreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}
