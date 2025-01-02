package gapi

import (
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
