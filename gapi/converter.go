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
