package gapi

import (
	"encoding/json"
	"fmt"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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

func convertStore(store db.Store) *pb.Store {
	return &pb.Store{
		StoreId:        store.ID,
		StoreName:      store.Name,
		StoreCreatedAt: timestamppb.New(store.CreatedAt),
	}
}

func convertFirstStoreRow(store db.GetFirstStoreRow) *pb.Store {
	return &pb.Store{
		StoreId:        store.StoreID,
		StoreName:      store.StoreName,
		StoreCreatedAt: timestamppb.New(store.StoreCreatedAt),
		Billboards:     mapBillboards(store.Billboards),
		Categories:     mapCategories(store.Categories),
		Sizes:          mapSizes(store.Sizes),
		Colors:         mapColors(store.Colors),
		Orders:         mapOrders(store.Orders),
		Products:       mapProducts(store.Products),
	}
}

func convertStoreRow(store db.GetStoreRow) *pb.Store {
	return &pb.Store{
		StoreId:        store.StoreID,
		StoreName:      store.StoreName,
		StoreCreatedAt: timestamppb.New(store.StoreCreatedAt),
		Billboards:     mapBillboards(store.Billboards),
		Categories:     mapCategories(store.Categories),
		Sizes:          mapSizes(store.Sizes),
		Colors:         mapColors(store.Colors),
		Orders:         mapOrders(store.Orders),
		Products:       mapProducts(store.Products),
	}
}

func convertStoresRow(stores []db.GetStoresRow) []*pb.Store {
	var pbStores []*pb.Store
	for _, store := range stores {
		pbStores = append(pbStores, &pb.Store{
			StoreId:        store.StoreID,
			StoreName:      store.StoreName,
			StoreCreatedAt: timestamppb.New(store.StoreCreatedAt),
			Billboards:     mapBillboards(store.Billboards),
			Categories:     mapCategories(store.Categories),
			Sizes:          mapSizes(store.Sizes),
			Colors:         mapColors(store.Colors),
			Orders:         mapOrders(store.Orders),
			Products:       mapProducts(store.Products),
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
		Id:          category.ID,
		Name:        category.Name,
		BillboardId: category.BillboardID,
		StoreId:     category.StoreID,
		CreatedAt:   timestamppb.New(category.CreatedAt),
	}
}

func convertCategoryRow(category db.GetCategoryRow) *pb.Category {
	return &pb.Category{
		Id:          category.CategoryID,
		Name:        category.CategoryName,
		BillboardId: category.CategoryBillboardID,
		StoreId:     category.CategoryStoreID,
		Billboard:   mapBillboards(category.Billboard),
		CreatedAt:   timestamppb.New(category.CategoryCreatedAt),
	}
}

func convertCategoriesRow(categories []db.GetCategoriesRow) []*pb.Category {
	var pbCategories []*pb.Category
	for _, category := range categories {
		pbCategories = append(pbCategories, &pb.Category{
			Id:          category.CategoryID,
			BillboardId: category.CategoryBillboardID,
			StoreId:     category.CategoryStoreID,
			Name:        category.CategoryName,
			Billboard:   mapBillboards(category.Billboards),
			CreatedAt:   timestamppb.New(category.CategoryCreatedAt),
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
		CreatedAt: timestamppb.New(size.CreatedAt),
	}
}

func convertSizes(sizes []db.Size) []*pb.Size {
	var pbSizes []*pb.Size
	for _, size := range sizes {
		pbSizes = append(pbSizes, &pb.Size{
			Id:        size.ID,
			StoreId:   size.StoreID,
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
		CreatedAt: timestamppb.New(color.CreatedAt),
	}
}

func convertColors(colors []db.Color) []*pb.Color {
	var pbColors []*pb.Color
	for _, color := range colors {
		pbColors = append(pbColors, &pb.Color{
			Id:      color.ID,
			StoreId: color.StoreID,

			Value:     color.Value,
			Name:      color.Name,
			CreatedAt: timestamppb.New(color.CreatedAt),
		})
	}

	return pbColors
}

func convertOrders(orders []db.Order) []*pb.Order {
	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id:         order.ID,
			OrderItems: mapItems(order.Items),
		})
	}

	return pbOrders
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
		Images:      mapImages(product.Images),
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

// order items
func mapItems(items json.RawMessage) []*pb.OrderItem {
	var rawItems []struct {
		Name string `json:"name"`
	}

	if err := json.Unmarshal(items, &rawItems); err != nil {
		return nil
	}

	var pbItems []*pb.OrderItem
	for _, item := range rawItems {
		pbOrderItem := &pb.OrderItem{
			Name: item.Name,
		}
		pbItems = append(pbItems, pbOrderItem)
	}

	return pbItems
}

// billboards
func mapBillboards(items json.RawMessage) []*pb.Billboard {
	var rawBillboards []struct {
		Billboard_Label    string `json:"billboard_label"`
		Billboard_ID       int64  `json:"billboard_id"`
		Billboard_ImageUrl string `json:"billboard_image_url"`
	}

	if err := json.Unmarshal(items, &rawBillboards); err != nil {
		return nil
	}

	var pbBillboards []*pb.Billboard
	for _, item := range rawBillboards {
		pbBillboard := &pb.Billboard{
			Label:    item.Billboard_Label,
			Id:       item.Billboard_ID,
			ImageUrl: item.Billboard_ImageUrl,
		}
		pbBillboards = append(pbBillboards, pbBillboard)
	}

	return pbBillboards
}

// categories
func mapCategories(items json.RawMessage) []*pb.Category {
	var rawCategories []struct {
		ID      string `json:"category_id"`
		StoreID int64  `json:"store_id"`
		Name    string `json:"category_name"`
		// Bill
	}

	if err := json.Unmarshal(items, &rawCategories); err != nil {
		return nil
	}

	var pbCategories []*pb.Category
	for _, item := range rawCategories {
		pbCategory := &pb.Category{
			Name: item.Name,
		}
		pbCategories = append(pbCategories, pbCategory)
	}

	return pbCategories
}

func mapBillboard(item json.RawMessage) *pb.Billboard {
	var rawBillboard struct {
		Billboard_Label string `json:"billboard_label"`
		Billboard_ID    int64  `json:"billboard_id"`
	}

	if err := json.Unmarshal(item, &rawBillboard); err != nil {
		return nil
	}

	return &pb.Billboard{
		Label: rawBillboard.Billboard_Label,
		Id:    rawBillboard.Billboard_ID,
	}
}

// products
func mapProducts(items json.RawMessage) []*pb.Product {
	var rawProducts []struct {
		Name string `json:"name"`
	}

	if err := json.Unmarshal(items, &rawProducts); err != nil {
		return nil
	}

	var pbProducts []*pb.Product
	for _, item := range rawProducts {
		pbProduct := &pb.Product{
			Name: item.Name,
		}

		pbProducts = append(pbProducts, pbProduct)
	}

	return pbProducts
}

// orders
func mapOrders(items json.RawMessage) []*pb.Order {
	var rawOrders []struct {
		OrderPhone string `json:"order_phone"`
	}

	if err := json.Unmarshal(items, &rawOrders); err != nil {
		return nil
	}

	var pbOrders []*pb.Order
	for _, item := range rawOrders {
		pbOrder := &pb.Order{
			OrderPhone: item.OrderPhone,
		}

		pbOrders = append(pbOrders, pbOrder)
	}

	return pbOrders
}

// colors
func mapColors(items json.RawMessage) []*pb.Color {
	var rawColors []struct {
		Name string `json:"name"`
	}

	if err := json.Unmarshal(items, &rawColors); err != nil {
		return nil
	}

	var pbColors []*pb.Color
	for _, item := range rawColors {
		pbColor := &pb.Color{
			Name: item.Name,
		}

		pbColors = append(pbColors, pbColor)
	}

	return pbColors
}

// sizes
func mapSizes(items json.RawMessage) []*pb.Size {
	var rawSizes []struct {
		Name string `json:"name"`
	}

	if err := json.Unmarshal(items, &rawSizes); err != nil {
		return nil
	}

	var pbSizes []*pb.Size
	for _, item := range rawSizes {
		pbSize := &pb.Size{
			Name: item.Name,
		}

		pbSizes = append(pbSizes, pbSize)
	}

	return pbSizes
}
