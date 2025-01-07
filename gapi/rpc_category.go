package gapi

import (
	"context"
	"database/sql"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateCreateCategoryRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to create category")
	}

	store, err := s.repository.GetStore(ctx, req.GetStoreId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no store found")
	}

	billboard, err := s.repository.GetBillboard(ctx, req.GetBillboardId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no billboard found")
	}

	args := db.CreateCategoryParams{
		StoreID:     store.StoreID,
		BillboardID: billboard.ID,
		Name:        req.GetName(),
	}

	category, err := s.repository.CreateCategory(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to create category %s", err)
	}

	resp := &pb.CreateCategoryResponse{
		Category: convertCategory(category),
	}

	return resp, nil

}

func (s *Server) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateGetCategoryRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get category")
	}

	category, err := s.repository.GetCategory(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get category %s", err)
	}

	resp := &pb.GetCategoryResponse{
		Category: convertCategoryRow(category),
	}

	return resp, nil

}

func (s *Server) GetCategories(ctx context.Context, req *pb.GetCategoriesRequest) (*pb.GetCategoriesResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateGetCategoriesRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get categories")
	}

	categories, err := s.repository.GetCategories(ctx, req.GetStoreId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get categories %s", err)
	}

	reversedCategories := convertCategoriesRow(categories)
	for i, j := 0, len(reversedCategories)-1; i < j; i, j = i+1, j-1 {
		reversedCategories[i], reversedCategories[j] = reversedCategories[j], reversedCategories[i]
	}

	resp := &pb.GetCategoriesResponse{
		Categories: reversedCategories,
	}

	return resp, nil

}

func (s *Server) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateUpdateCategoryRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to update category")
	}

	args := db.UpdateCategoryParams{
		ID:      req.GetId(),
		StoreID: req.GetStoreId(),
		Name: sql.NullString{
			Valid:  true,
			String: req.GetName(),
		},
	}

	err = s.repository.UpdateCategory(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to update category %s", err)
	}

	resp := &pb.UpdateCategoryResponse{
		Message: "category update successful",
	}

	return resp, nil

}

func (s *Server) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateDeleteCategoryRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to delete category")
	}

	err = s.repository.DeleteCategory(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to delete category %s", err)
	}

	resp := &pb.DeleteCategoryResponse{
		Message: "category delete successful",
	}

	return resp, nil

}

func validateCreateCategoryRequest(req *pb.CreateCategoryRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}
	if err := validate.ValidateId(req.GetBillboardId()); err != nil {
		violations = append(violations, fieldViolation("billboard_id", err))
	}

	if err := validate.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	return violations
}

func validateGetCategoryRequest(req *pb.GetCategoryRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}

func validateGetCategoriesRequest(req *pb.GetCategoriesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

	return violations
}

func validateUpdateCategoryRequest(req *pb.UpdateCategoryRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}
	if err := validate.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	return violations
}

func validateDeleteCategoryRequest(req *pb.DeleteCategoryRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}
