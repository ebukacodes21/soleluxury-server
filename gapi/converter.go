package gapi

import (
	"encoding/json"
	"fmt"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertStore(store db.Store) *pb.Store {
	return &pb.Store{
		Id:        store.ID,
		Name:      store.Name,
		CreatedAt: timestamppb.New(store.CreatedAt),
	}
}

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Id:               user.ID,
		Username:         user.Username,
		Email:            user.Email,
		IsVerified:       user.IsVerified,
		VerificationCode: user.VerificationCode,
		Role:             user.Role,
		CreatedAt:        timestamppb.New(user.CreatedAt),
	}
}

func convertStores(stores []db.Store) []*pb.Store {
	var pbStores []*pb.Store
	for _, store := range stores {
		pbStores = append(pbStores, &pb.Store{
			Id:        store.ID,
			Name:      store.Name,
			CreatedAt: timestamppb.New(store.CreatedAt),
		})
	}

	return pbStores
}

func convertBillboard(billboard db.Billboard) *pb.Billboard {
	return &pb.Billboard{
		Id:        billboard.ID,
		Label:     billboard.Label,
		ImageUrl:  billboard.ImageUrl,
		StoreId:   billboard.StoreID,
		CreatedAt: timestamppb.New(billboard.CreatedAt),
	}
}

func convertBillboards(billboards []db.Billboard) []*pb.Billboard {
	var pbBillboards []*pb.Billboard
	for _, billboard := range billboards {
		pbBillboards = append(pbBillboards, &pb.Billboard{
			Id:        billboard.ID,
			Label:     billboard.Label,
			ImageUrl:  billboard.ImageUrl,
			StoreId:   billboard.StoreID,
			CreatedAt: timestamppb.New(billboard.CreatedAt),
		})
	}

	return pbBillboards
}

func convertCategory(category db.Category) *pb.Category {
	return &pb.Category{
		Id:             category.ID,
		Name:           category.Name,
		BillboardId:    category.BillboardID,
		BillboardLabel: category.BillboardLabel,
		StoreId:        category.StoreID,
		CreatedAt:      timestamppb.New(category.CreatedAt),
	}
}

func convertCategories(categories []db.Category) []*pb.Category {
	var pbCategories []*pb.Category
	for _, category := range categories {
		pbCategories = append(pbCategories, &pb.Category{
			Id:             category.ID,
			BillboardId:    category.BillboardID,
			StoreId:        category.StoreID,
			StoreName:      category.StoreName,
			BillboardLabel: category.BillboardLabel,
			Name:           category.Name,
			CreatedAt:      timestamppb.New(category.CreatedAt),
		})
	}

	return pbCategories
}

func convertSize(size db.Size) *pb.Size {
	return &pb.Size{
		Id:        size.ID,
		Name:      size.Name,
		Value:     size.Value,
		StoreId:   size.StoreID,
		StoreName: size.StoreName,
		CreatedAt: timestamppb.New(size.CreatedAt),
	}
}

func convertSizes(sizes []db.Size) []*pb.Size {
	var pbSizes []*pb.Size
	for _, size := range sizes {
		pbSizes = append(pbSizes, &pb.Size{
			Id:        size.ID,
			StoreId:   size.StoreID,
			StoreName: size.StoreName,
			Value:     size.Value,
			Name:      size.Name,
			CreatedAt: timestamppb.New(size.CreatedAt),
		})
	}

	return pbSizes
}

func convertColor(color db.Color) *pb.Color {
	return &pb.Color{
		Id:        color.ID,
		Name:      color.Name,
		Value:     color.Value,
		StoreId:   color.StoreID,
		StoreName: color.StoreName,
		CreatedAt: timestamppb.New(color.CreatedAt),
	}
}

func convertColors(colors []db.Color) []*pb.Color {
	var pbColors []*pb.Color
	for _, color := range colors {
		pbColors = append(pbColors, &pb.Color{
			Id:        color.ID,
			StoreId:   color.StoreID,
			StoreName: color.StoreName,
			Value:     color.Value,
			Name:      color.Name,
			CreatedAt: timestamppb.New(color.CreatedAt),
		})
	}

	return pbColors
}

func convertProduct(product db.Product) (*pb.Product, error) {
	var images []Image
	if len(product.Images) > 0 {
		err := json.Unmarshal(product.Images, &images)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal images: %v", err)
		}
	}

	var imageUrls []string
	for _, img := range images {
		imageUrls = append(imageUrls, img.URL)
	}

	return &pb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Price:       float32(product.Price),
		Description: product.Description,
		IsFeatured:  product.IsFeatured,
		IsArchived:  product.IsArchived,
		Images:      imageUrls,
		CreatedAt:   timestamppb.New(product.CreatedAt),
	}, nil
}

func convertProductsRow(pr []db.GetProductsRow) []*pb.ProductResponse {
	var productResponses []*pb.ProductResponse

	for _, row := range pr {
		product := &pb.ProductResponse{
			Id:           row.ProductID,
			Name:         row.ProductName,
			Description:  row.ProductDescription,
			Price:        float32(row.ProductPrice),
			IsFeatured:   row.IsFeatured,
			IsArchived:   row.IsArchived,
			Images:       mapImages(row.ProductImages),
			CategoryId:   row.CategoryID.Int64,
			CategoryName: row.CategoryName.String,
			ColorId:      row.ColorID.Int64,
			ColorValue:   row.ColorValue.String,
			SizeId:       row.SizeID.Int64,
			SizeValue:    row.SizeValue.String,
			CreatedAt:    timestamppb.New(row.ProductCreatedAt),
		}

		productResponses = append(productResponses, product)
	}

	return productResponses
}

func convertProductRow(product db.GetProductRow) *pb.ProductResponse {
	return &pb.ProductResponse{
		Id:           product.ProductID,
		Name:         product.ProductName,
		Description:  product.ProductDescription,
		Price:        float32(product.ProductPrice),
		IsFeatured:   product.IsFeatured,
		IsArchived:   product.IsArchived,
		Images:       mapImages(product.ProductImages),
		CategoryId:   product.CategoryID.Int64,
		CategoryName: product.CategoryName.String,
		ColorId:      product.ColorID.Int64,
		ColorValue:   product.ColorValue.String,
		SizeId:       product.SizeID.Int64,
		SizeValue:    product.SizeValue.String,
		CreatedAt:    timestamppb.New(product.ProductCreatedAt),
	}
}

func mapImages(images json.RawMessage) []*pb.Item {
	var rawImages []struct {
		Url string `json:"url"`
	}

	if err := json.Unmarshal(images, &rawImages); err != nil {
		return nil
	}

	var pbItems []*pb.Item
	for _, img := range rawImages {
		pbItem := &pb.Item{
			Url: img.Url,
		}
		pbItems = append(pbItems, pbItem)
	}

	return pbItems
}
