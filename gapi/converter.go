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
