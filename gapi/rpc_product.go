package gapi

import (
	"context"
	"database/sql"
	"encoding/json"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/sqlc-dev/pqtype"

	"github.com/ebukacodes21/soleluxury-server/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Image struct {
	URL string `json:"url"`
}

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

	imagesJSON, err := rawMessage(req.GetImages())
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to marshal product image")
	}

	prodArgs := db.CreateProductParams{
		Name:        req.GetName(),
		Price:       float64(req.GetPrice()),
		IsFeatured:  req.GetIsFeatured(),
		IsArchived:  req.GetIsArchived(),
		Description: req.GetDescription(),
		Images:      imagesJSON,
	}

	product, err := s.repository.CreateProduct(ctx, prodArgs)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create product")
	}

	args := db.CreateProductTxParams{
		CreateProductColorParams: db.CreateProductColorParams{
			ProductID: product.ID,
			ColorID:   req.GetColorId(),
		},
		CreateProductSizeParams: &db.CreateProductSizeParams{
			ProductID: product.ID,
			SizeID:    req.GetSizeId(),
		},
		CreateProductStoreParams: &db.CreateProductStoreParams{
			ProductID: product.ID,
			StoreID:   req.GetStoreId(),
		},
		CreateProductCategoryParams: &db.CreateProductCategoryParams{
			ProductID:  product.ID,
			CategoryID: req.GetCategoryId(),
		},
	}

	_, err = s.repository.CreateProductTx(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create joint table data")
	}

	p, err := convertProduct(product)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to convert product %s", err)
	}

	resp := &pb.CreateProductResponse{
		Product: p,
	}

	return resp, nil
}

func (s *Server) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateGetProductsRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get products")
	}

	row, err := s.repository.GetProducts(ctx, req.GetStoreId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to find products")
	}

	resp := &pb.GetProductsResponse{
		ProductRes: convertProductsRow(row),
	}

	return resp, nil
}

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	violations := validateGetProductRequest(req)
	if violations != nil {
		return nil, invalidArgs(violations)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get a product")
	}

	args := db.GetProductParams{
		StoreID:   req.GetStoreId(),
		ProductID: req.GetProductId(),
	}

	product, err := s.repository.GetProduct(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found")
	}

	resp := &pb.GetProductResponse{
		ProductRes: convertProductRow(product),
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

	images, err := rawMessage(req.GetImages())
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to marshal product image")
	}

	args := db.UpdateProductTxParams{
		UpdateProductParams: db.UpdateProductParams{
			ID: req.GetProductId(),
			Name: sql.NullString{
				Valid:  true,
				String: req.GetName(),
			},
			Price: sql.NullFloat64{
				Valid:   true,
				Float64: float64(req.GetPrice()),
			},
			IsFeatured: sql.NullBool{
				Valid: true,
				Bool:  req.GetIsFeatured(),
			},
			Description: sql.NullString{
				Valid:  true,
				String: req.GetDescription(),
			},
			Images: pqtype.NullRawMessage{
				Valid:      true,
				RawMessage: images,
			},
		},
		UpdateProductColorParams: db.UpdateProductColorParams{
			ProductID: req.ProductId,
			ColorID: sql.NullInt64{
				Valid: true,
				Int64: req.GetColorId(),
			},
		},
		UpdateProductSizeParams: &db.UpdateProductSizeParams{
			ProductID: req.ProductId,
			SizeID: sql.NullInt64{
				Valid: true,
				Int64: req.GetSizeId(),
			},
		},
		UpdateProductCategoryParams: &db.UpdateProductCategoryParams{
			ProductID: req.ProductId,
			CategoryID: sql.NullInt64{
				Valid: true,
				Int64: req.GetCategoryId(),
			},
		},
	}

	err = s.repository.UpdateProductTx(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to update a product")
	}

	resp := &pb.UpdateProductResponse{
		Message: "product update successful",
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

	err = s.repository.DeleteProductTx(ctx, req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to delete a product")
	}

	resp := &pb.DeleteProductResponse{
		Message: "product delete successful",
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

func validateGetProductRequest(req *pb.GetProductRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateId(req.GetStoreId()); err != nil {
		violations = append(violations, fieldViolation("store_id", err))
	}

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

func rawMessage(data []*pb.Item) (json.RawMessage, error) {
	imagesJSON, err := json.Marshal(data)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "error marshaling images: %v", err)
	}

	var images []Image
	for _, img := range data {
		images = append(images, Image{URL: img.GetUrl()})
	}

	imagesJSON, err = json.Marshal(images)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to marshal JSON")
	}

	return imagesJSON, nil
}
