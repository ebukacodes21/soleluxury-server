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

func (s *Server) CreateColor(ctx context.Context, req *pb.CreateColorRequest) (*pb.CreateColorResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateCreateColorRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to create color")
	}

	store, err := s.repository.GetStore(ctx, req.GetStoreId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no store found")
	}

	args := db.CreateColorParams{
		StoreID:   store.ID,
		StoreName: store.Name,
		Name:      req.GetName(),
		Value:     req.GetValue(),
	}

	color, err := s.repository.CreateColor(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create color %s", err)
	}

	resp := &pb.CreateColorResponse{
		Color: convertColor(color),
	}

	return resp, nil

}

func (s *Server) GetColor(ctx context.Context, req *pb.GetColorRequest) (*pb.GetColorResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateGetColorRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get color")
	}

	color, err := s.repository.GetColor(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get color %s", err)
	}

	resp := &pb.GetColorResponse{
		Color: convertColor(color),
	}

	return resp, nil
}

func (s *Server) GetColors(ctx context.Context, req *pb.GetColorsRequest) (*pb.GetColorsResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateGetColorsRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get colors")
	}

	colors, err := s.repository.GetColors(ctx, req.GetStoreId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get colors %s", err)
	}

	reversedColors := convertColors(colors)
	for i, j := 0, len(reversedColors)-1; i < j; i, j = i+1, j-1 {
		reversedColors[i], reversedColors[j] = reversedColors[j], reversedColors[i]
	}

	resp := &pb.GetColorsResponse{
		Colors: reversedColors,
	}

	return resp, nil

}

func (s *Server) UpdateColor(ctx context.Context, req *pb.UpdateColorRequest) (*pb.UpdateColorResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateUpdateColorRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to update color")
	}

	args := db.UpdateColorParams{
		ID:      req.GetId(),
		StoreID: req.GetStoreId(),
		Name: sql.NullString{
			Valid:  true,
			String: req.GetName(),
		},
		Value: sql.NullString{
			Valid:  true,
			String: req.GetValue(),
		},
	}

	err = s.repository.UpdateColor(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to update color %s", err)
	}

	resp := &pb.UpdateColorResponse{
		Message: "color update successful",
	}

	return resp, nil

}

func (s *Server) DeleteColor(ctx context.Context, req *pb.DeleteColorRequest) (*pb.DeleteColorResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateDeleteColorRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to delete color")
	}

	err = s.repository.DeleteColor(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to delete color %s", err)
	}

	resp := &pb.DeleteColorResponse{
		Message: "color delete successful",
	}

	return resp, nil

}

func validateCreateColorRequest(req *pb.CreateColorRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

	if err := validate.ValidateStoreName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	if err := validate.ValidateColorValue(req.GetValue()); err != nil {
		violations = append(violations, fieldViolation("value", err))
	}

	return violations
}

func validateGetColorRequest(req *pb.GetColorRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}

func validateGetColorsRequest(req *pb.GetColorsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

	return violations
}

func validateUpdateColorRequest(req *pb.UpdateColorRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}
	if err := validate.ValidateStoreName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	if err := validate.ValidateColorValue(req.GetValue()); err != nil {
		violations = append(violations, fieldViolation("value", err))
	}

	return violations
}

func validateDeleteColorRequest(req *pb.DeleteColorRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetId()); err != nil {
		violations = append(violations, fieldViolation("id", err))
	}

	return violations
}
