package servers

import (
	"context"
	"log"
	"net"

	"github.com/ebukacodes21/soleluxury-server/gapi"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGrpcServer(group *errgroup.Group, ctx context.Context) {
	// instantiate a new instance of our server implementation
	server := gapi.NewServer()

	// create a new grpc server by calling the NewServer method on the grpc package. it returns a grpc server
	gServer := grpc.NewServer()

	// register the grpc server togther with its implementation
	pb.RegisterSoleluxuryServer(gServer, server)

	// this is used to dynamically discover services, methods and messages that a grpc server suppprts
	reflection.Register(gServer)

	// obtain a listener where the grpc server will listen and serve requests
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}

	// starting server in go routine
	group.Go(func() error {
		log.Print("Grpc server running on ", "0.0.0.0:8000")
		err = gServer.Serve(listener)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	// graceful shutdown
	group.Go(func() error {
		<-ctx.Done()
		log.Print("grpc gracefully shutting down...")

		gServer.GracefulStop()
		log.Print("goodbye")

		return nil
	})
}
