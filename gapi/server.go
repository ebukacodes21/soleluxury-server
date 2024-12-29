package gapi

import (
	"fmt"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/token"
	"github.com/ebukacodes21/soleluxury-server/utils"
	"github.com/ebukacodes21/soleluxury-server/worker"
)

type Server struct {
	pb.UnimplementedSoleluxuryServer
	repository db.DatabaseContract
	config     utils.Config
	token      token.TokenContract
	td         worker.Distributor
	tp         worker.Processor
}

func NewServer(repository db.DatabaseContract, config utils.Config, td worker.Distributor, tp worker.Processor) (*Server, error) {
	token, err := token.NewToken(config.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("unable to create token maker %w ", err)
	}

	return &Server{
		repository: repository,
		config:     config,
		td:         td,
		tp:         tp,
		token:      token,
	}, nil
}
