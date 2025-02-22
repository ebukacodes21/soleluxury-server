package servers

import (
	"context"
	"log"
	"net"

	db "github.com/ebukacodes21/soleluxury-server/db"
	"github.com/ebukacodes21/soleluxury-server/gapi"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/utils"
	"github.com/ebukacodes21/soleluxury-server/worker"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGrpcServer(group *errgroup.Group, ctx context.Context, repository db.DatabaseContract, config utils.Config, td worker.Distributor, tp worker.Processor) {
	server, err := gapi.NewServer(repository, config, td, tp)
	if err != nil {
		log.Fatal(err)
	}
	logger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	gServer := grpc.NewServer(logger)

	pb.RegisterSoleluxuryServer(gServer, server)
	reflection.Register(gServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	group.Go(func() error {
		log.Print("Grpc server running on ", config.GRPCServerAddr)
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
