package gapi

import (
	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
)

type Server struct {
	pb.UnimplementedSoleluxuryServer
	repository db.DatabaseContract
}

func NewServer(repository db.DatabaseContract) *Server {
	return &Server{repository: repository}
}
