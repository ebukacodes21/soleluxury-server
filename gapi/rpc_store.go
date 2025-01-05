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
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *Server) GetStore(ctx context.Context, req *pb.GetStoreRequest) (*pb.GetStoreResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateGetStoreRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get store")
	}

	store, err := s.repository.GetStore(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get store")
	}

	resp := &pb.GetStoreResponse{
		Store: convertStore(store),
	}

	return resp, nil
}

func (s *Server) GetStores(ctx context.Context, _ *emptypb.Empty) (*pb.GetStoresResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get stores")
	}

	stores, err := s.repository.GetStores(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get store")
	}

	reversedStores := convertStores(stores)
	for i, j := 0, len(reversedStores)-1; i < j; i, j = i+1, j-1 {
		reversedStores[i], reversedStores[j] = reversedStores[j], reversedStores[i]
	}

	resp := &pb.GetStoresResponse{
		Stores: reversedStores,
	}

	return resp, nil
}

func (s *Server) GetFirstStore(ctx context.Context, _ *emptypb.Empty) (*pb.GetStoreResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get store")
	}

	store, err := s.repository.GetFirstStore(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get store")
	}

	resp := &pb.GetStoreResponse{
		Store: convertStore(store),
	}

	return resp, nil
}

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

	err = s.repository.DeleteStore(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to delete store")
	}

	resp := &pb.DeleteStoreResponse{
		Message: "store delete successful",
	}

	return resp, nil
}

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

	if err := validate.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	return violations
}

func validateDeleteStoreRequest(req *pb.DeleteStoreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}
	return violations
}
