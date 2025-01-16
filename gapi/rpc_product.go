package gapi

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/pb"

	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateCreateProductRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to create product")
	}

	product, err := s.repository.CreateProduct(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create product")
	}

	resp := &pb.CreateProductResponse{
		Product: convertProduct(product),
	}

	return resp, nil
}

func (s *Server) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	// payload, err := s.authGuard(ctx, []string{"user", "admin"})
	// if err != nil {
	// 	return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	// }

	violations := validateGetProductsRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	// if payload.Role == "user" {
	// 	return nil, status.Errorf(codes.PermissionDenied, "not authorized to get products")
	// }

	products, err := s.repository.GetProducts(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to find products")
	}

	resp := &pb.GetProductsResponse{
		ProductRes: convertProducts(products),
	}

	return resp, nil
}

func (s *Server) GetCategoryProducts(ctx context.Context, req *pb.GetCategoryProductsRequest) (*pb.GetCategoryProductsResponse, error) {
	// payload, err := s.authGuard(ctx, []string{"user", "admin"})
	// if err != nil {
	// 	return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	// }

	violations := validateGetCategoryProductsRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	// if payload.Role == "user" {
	// 	return nil, status.Errorf(codes.PermissionDenied, "not authorized to get products")
	// }

	products, err := s.repository.GetCategoryProducts(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to find products")
	}

	resp := &pb.GetCategoryProductsResponse{
		ProductRes: convertProducts(products),
	}

	return resp, nil
}

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	// payload, err := s.authGuard(ctx, []string{"user", "admin"})
	// if err != nil {
	// 	return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	// }

	violations := validateGetProductRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	// if payload.Role == "user" {
	// 	return nil, status.Errorf(codes.PermissionDenied, "not authorized to get a product")
	// }

	product, err := s.repository.GetProductByID(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found")
	}

	resp := &pb.GetProductResponse{
		ProductRes: convertSingleProduct(product),
	}

	return resp, nil
}

func (s *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateUpdateProductRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to update a product")
	}

	message, err := s.repository.UpdateProduct(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to update a product")
	}

	resp := &pb.UpdateProductResponse{
		Message: message,
	}

	return resp, nil
}

func (s *Server) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateDeleteProductRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to delete a product")
	}

	message, err := s.repository.DeleteProduct(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to delete a product")
	}

	resp := &pb.DeleteProductResponse{
		Message: message,
	}

	return resp, nil
}

func validateCreateProductRequest(req *pb.CreateProductRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

	if err := validate.ValidateId(req.GetCategoryId()); err != nil {
		violations = append(violations, fieldViolation("category_id", err))
	}

	if err := validate.ValidateId(req.GetColorId()); err != nil {
		violations = append(violations, fieldViolation("color_id", err))
	}

	if err := validate.ValidateId(req.GetSizeId()); err != nil {
		violations = append(violations, fieldViolation("size_id", err))
	}

	if err := validate.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	if err := validate.ValidateDescription(req.GetDescription()); err != nil {
		violations = append(violations, fieldViolation("description", err))
	}

	if err := validate.ValidatePrice(req.GetPrice()); err != nil {
		violations = append(violations, fieldViolation("price", err))
	}

	if err := validate.ValidateBool(req.GetIsArchived()); err != nil {
		violations = append(violations, fieldViolation("is_archived", err))
	}

	if err := validate.ValidateBool(req.GetIsFeatured()); err != nil {
		violations = append(violations, fieldViolation("is_featured", err))
	}

	if err := validate.ValidateUrls(req.GetImages()); err != nil {
		violations = append(violations, fieldViolation("images", err))
	}

	return violations
}

func validateGetProductsRequest(req *pb.GetProductsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

	return violations
}

func validateGetCategoryProductsRequest(req *pb.GetCategoryProductsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetCategoryId()); err != nil {
		violations = append(violations, fieldViolation("category_id", err))
	}

	return violations
}

func validateGetProductRequest(req *pb.GetProductRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetProductId()); err != nil {
		violations = append(violations, fieldViolation("product_id", err))
	}

	return violations
}

func validateUpdateProductRequest(req *pb.UpdateProductRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetProductId()); err != nil {
		violations = append(violations, fieldViolation("product_id", err))
	}

	if req.CategoryId != nil {
		if err := validate.ValidateId(req.GetCategoryId()); err != nil {
			violations = append(violations, fieldViolation("category_id", err))
		}
	}

	if req.ColorId != nil {
		if err := validate.ValidateId(req.GetColorId()); err != nil {
			violations = append(violations, fieldViolation("color_id", err))
		}
	}

	if req.SizeId != nil {
		if err := validate.ValidateId(req.GetSizeId()); err != nil {
			violations = append(violations, fieldViolation("size_id", err))
		}
	}

	if req.Name != nil {
		if err := validate.ValidateName(req.GetName()); err != nil {
			violations = append(violations, fieldViolation("name", err))
		}
	}

	if req.Description != nil {
		if err := validate.ValidateDescription(req.GetDescription()); err != nil {
			violations = append(violations, fieldViolation("description", err))
		}
	}

	if req.Price != nil {
		if err := validate.ValidatePrice(req.GetPrice()); err != nil {
			violations = append(violations, fieldViolation("price", err))
		}
	}

	if req.IsArchived != nil {
		if err := validate.ValidateBool(req.GetIsArchived()); err != nil {
			violations = append(violations, fieldViolation("is_archived", err))
		}
	}

	if req.IsFeatured != nil {
		if err := validate.ValidateBool(req.GetIsFeatured()); err != nil {
			violations = append(violations, fieldViolation("is_featured", err))
		}
	}

	if err := validate.ValidateUrls(req.GetImages()); err != nil {
		violations = append(violations, fieldViolation("images", err))
	}
	return violations
}

func validateDeleteProductRequest(req *pb.DeleteProductRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetProductId()); err != nil {
		violations = append(violations, fieldViolation("product_id", err))
	}

	if err := validate.ValidateId(req.GetProductId()); err != nil {
		violations = append(violations, fieldViolation("product_id", err))
	}

	return violations
}
