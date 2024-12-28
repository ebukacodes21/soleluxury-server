package gapi

import (
	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/utils"
	"github.com/ebukacodes21/soleluxury-server/worker"
)

type Server struct {
	pb.UnimplementedSoleluxuryServer
	repository db.DatabaseContract
	config     utils.Config
	td         worker.Distributor
	tp         worker.Processor
}

func NewServer(repository db.DatabaseContract, config utils.Config, td worker.Distributor, tp worker.Processor) *Server {
	return &Server{
		repository: repository,
		config:     config,
		td:         td,
		tp:         tp,
	}
}
