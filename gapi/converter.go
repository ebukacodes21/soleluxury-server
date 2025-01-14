package gapi

import (
	"github.com/ebukacodes21/soleluxury-server/db"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user *db.User) *pb.User {
	return &pb.User{
		Id:               user.ID.Hex(),
		Username:         user.Username,
		Email:            user.Email,
		VerificationCode: user.VerificationCode,
		Role:             user.Role,
		IsVerified:       user.IsVerified,
		CreatedAt:        timestamppb.New(user.CreatedAt),
	}
}

func convertStore(store *db.Store) *pb.Store {
	return &pb.Store{
		Id:        store.ID.Hex(),
		Name:      store.Name,
		CreatedAt: timestamppb.New(store.CreatedAt),
	}
}

func convertStores(stores []db.Store) []*pb.Store {
	var pbStores []*pb.Store

	for _, store := range stores {
		pbStore := &pb.Store{
			Id:        store.ID.Hex(),
			Name:      store.Name,
			CreatedAt: timestamppb.New(store.CreatedAt),
		}

		pbStores = append(pbStores, pbStore)
	}
	return pbStores
}

func convertBillboard(billboard *db.Billboard) *pb.Billboard {
	return &pb.Billboard{
		Id:         billboard.ID.Hex(),
		StoreId:    billboard.StoreID.Hex(),
		Store:      convertStore(&billboard.Store),
		Label:      billboard.Label,
		ImageUrl:   billboard.ImageURL,
		Categories: convertCategories(billboard.Categories),
		CreatedAt:  timestamppb.New(billboard.CreatedAt),
	}
}

func convertBillboards(billboards []db.Billboard) []*pb.Billboard {
	var pbBillboards []*pb.Billboard

	for _, billboard := range billboards {
		pbBillboard := &pb.Billboard{
			Id:         billboard.ID.Hex(),
			StoreId:    billboard.StoreID.Hex(),
			Store:      convertStore(&billboard.Store),
			Label:      billboard.Label,
			ImageUrl:   billboard.ImageURL,
			Categories: convertCategories(billboard.Categories),
			CreatedAt:  timestamppb.New(billboard.CreatedAt),
		}

		pbBillboards = append(pbBillboards, pbBillboard)
	}
	return pbBillboards
}

func convertCategory(category *db.Category) *pb.Category {
	return &pb.Category{
		Id:          category.ID.Hex(),
		StoreId:     category.StoreID.Hex(),
		BillboardId: category.Billboard.ID.Hex(),
		Name:        category.Name,
		Billboard:   convertBillboard(&category.Billboard),
		CreatedAt:   timestamppb.New(category.CreatedAt),
	}
}

func convertCategories(categories []db.Category) []*pb.Category {
	var pbCategories []*pb.Category

	for _, category := range categories {
		pbCategory := &pb.Category{
			Id:          category.ID.Hex(),
			StoreId:     category.StoreID.Hex(),
			BillboardId: category.Billboard.ID.Hex(),
			Name:        category.Name,
			Billboard:   convertBillboard(&category.Billboard),
			CreatedAt:   timestamppb.New(category.CreatedAt),
		}

		pbCategories = append(pbCategories, pbCategory)
	}
	return pbCategories
}

func convertSize(size *db.Size) *pb.Size {
	return &pb.Size{
		Id:        size.ID.Hex(),
		Name:      size.Name,
		Value:     size.Value,
		Store:     convertStore(&size.Store),
		CreatedAt: timestamppb.New(size.CreatedAt),
	}
}

func convertSizes(sizes []db.Size) []*pb.Size {
	var pbSizes []*pb.Size

	for _, size := range sizes {
		pbSize := &pb.Size{
			Id:        size.ID.Hex(),
			Name:      size.Name,
			Value:     size.Value,
			Store:     convertStore(&size.Store),
			CreatedAt: timestamppb.New(size.CreatedAt),
		}

		pbSizes = append(pbSizes, pbSize)
	}
	return pbSizes
}

func convertColor(color *db.Color) *pb.Color {
	return &pb.Color{
		Id:        color.ID.Hex(),
		Name:      color.Name,
		Value:     color.Value,
		Store:     convertStore(&color.Store),
		CreatedAt: timestamppb.New(color.CreatedAt),
	}
}

func convertColors(colors []db.Color) []*pb.Color {
	var pbColors []*pb.Color

	for _, color := range colors {
		pbColor := &pb.Color{
			Id:        color.ID.Hex(),
			Name:      color.Name,
			Value:     color.Value,
			Store:     convertStore(&color.Store),
			CreatedAt: timestamppb.New(color.CreatedAt),
		}

		pbColors = append(pbColors, pbColor)
	}
	return pbColors
}

func convertProduct(product *db.Product) *pb.Product {
	return &pb.Product{
		Id:          product.ID.Hex(),
		Name:        product.Name,
		Price:       float32(product.Price),
		Description: product.Description,
		Images:      convertImage(product.Images),
		IsFeatured:  product.IsFeatured,
		IsArchived:  product.IsArchived,
		CreatedAt:   timestamppb.New(product.CreatedAt),
	}
}

func convertProducts(products []db.Product) []*pb.ProductResponse {
	var pbProducts []*pb.ProductResponse

	for _, product := range products {
		pbProduct := &pb.ProductResponse{
			Id:          product.ID.Hex(),
			StoreId:     product.StoreID.Hex(),
			Store:       convertStore(&product.Store),
			CategoryId:  product.CategoryID.Hex(),
			Category:    convertCategory(&product.Category),
			Name:        product.Name,
			Price:       float32(product.Price),
			IsFeatured:  product.IsFeatured,
			IsArchived:  product.IsArchived,
			SizeId:      product.SizeID.Hex(),
			Size:        convertSize(&product.Size),
			ColorId:     product.ColorID.Hex(),
			Color:       convertColor(&product.Color),
			Images:      convertImage(product.Images),
			Description: product.Description,
			CreatedAt:   timestamppb.New(product.CreatedAt),
		}

		pbProducts = append(pbProducts, pbProduct)
	}

	return pbProducts
}

func convertSingleProduct(product *db.Product) *pb.ProductResponse {
	return &pb.ProductResponse{
		Id:          product.ID.Hex(),
		StoreId:     product.StoreID.Hex(),
		Store:       convertStore(&product.Store),
		CategoryId:  product.CategoryID.Hex(),
		Category:    convertCategory(&product.Category),
		Name:        product.Name,
		Price:       float32(product.Price),
		IsFeatured:  product.IsFeatured,
		IsArchived:  product.IsArchived,
		SizeId:      product.SizeID.Hex(),
		Size:        convertSize(&product.Size),
		ColorId:     product.ColorID.Hex(),
		Color:       convertColor(&product.Color),
		Images:      convertImage(product.Images),
		Description: product.Description,
		CreatedAt:   timestamppb.New(product.CreatedAt),
	}
}

func convertImage(images []db.Image) []*pb.Item {
	var dbImages []*pb.Item
	for _, image := range images {
		dbImage := &pb.Item{
			Url: image.URL,
		}

		dbImages = append(dbImages, dbImage)
	}
	return dbImages
}
