package servers

import (
	"context"
	"log"
	"net"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/gapi"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGrpcServer(group *errgroup.Group, ctx context.Context, repository db.DatabaseContract) {
	server := gapi.NewServer(repository)
	// logger
	gServer := grpc.NewServer()

	pb.RegisterSoleluxuryServer(gServer, server)
	reflection.Register(gServer)

	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}

	group.Go(func() error {
		log.Print("Grpc server running on ", "0.0.0.0:8000")
		err = gServer.Serve(listener)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		log.Print("grpc gracefully shutting down...")

		gServer.GracefulStop()
		log.Print("goodbye")

		return nil
	})
}
