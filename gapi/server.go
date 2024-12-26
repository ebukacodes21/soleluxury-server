package gapi

import (
	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/utils"
)

type Server struct {
	pb.UnimplementedSoleluxuryServer
	repository db.DatabaseContract
	config     utils.Config
}

func NewServer(repository db.DatabaseContract, config utils.Config) *Server {
	return &Server{repository: repository, config: config}
}
