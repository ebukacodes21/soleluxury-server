package gapi

import "github.com/ebukacodes21/soleluxury-server/pb"

type Server struct {
	pb.UnimplementedSoleluxuryServer
}

func NewServer() *Server {
	return &Server{}
}
