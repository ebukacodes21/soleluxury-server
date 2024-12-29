package gapi

import (
	"context"

	"github.com/ebukacodes21/soleluxury-server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetStores(ctx context.Context, _ *emptypb.Empty) (*pb.GetStoresResponse, error) {
	payload, err := s.authGuard(ctx, []string{"user", "admin"})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized to access route %s ", err)
	}

	if payload.Role == "user" {
		return nil, status.Errorf(codes.PermissionDenied, "not authorized to get stores")
	}

	stores, err := s.repository.GetStores(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to get store")
	}

	reversedStores := convertStores(stores)
	for i, j := 0, len(reversedStores)-1; i < j; i, j = i+1, j-1 {
		reversedStores[i], reversedStores[j] = reversedStores[j], reversedStores[i]
	}

	resp := &pb.GetStoresResponse{
		Stores: reversedStores,
	}

	return resp, nil
}
