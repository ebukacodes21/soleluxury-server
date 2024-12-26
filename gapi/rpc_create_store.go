package gapi

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateStore(ctx context.Context, req *pb.CreateStoreRequest) (*pb.CreateStoreResponse, error) {

	// call the database
	store := &pb.Store{
		Id:        1,
		Name:      req.GetName(),
		CreatedAt: timestamppb.Now(),
	}

	resp := &pb.CreateStoreResponse{
		Store: store,
	}

	return resp, nil
}
