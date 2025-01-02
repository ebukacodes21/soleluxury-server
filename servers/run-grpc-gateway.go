package servers

import (
	"context"
	"log"
	"net/http"

	_ "github.com/ebukacodes21/soleluxury-server/doc/statik"
	"github.com/ebukacodes21/soleluxury-server/parser"
	"github.com/ebukacodes21/soleluxury-server/worker"

	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/gapi"
	"github.com/ebukacodes21/soleluxury-server/pb"
	"github.com/ebukacodes21/soleluxury-server/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/encoding/protojson"
)

func RunGrpcGateway(group *errgroup.Group, ctx context.Context, repository db.DatabaseContract, config utils.Config, td worker.Distributor, tp worker.Processor) {
	server, err := gapi.NewServer(repository, config, td, tp)
	if err != nil {
		log.Fatal(err)
	}

	options := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	mux := runtime.NewServeMux(runtime.SetQueryParameterParser(&parser.CustomQueryParameterParser{}), options)
	err = pb.RegisterSoleluxuryHandlerServer(ctx, mux, server)
	if err != nil {
		log.Fatal(err)
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)

	sf, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	httpMux.Handle("/swagger", http.StripPrefix("/swagger", http.FileServer(sf)))

	c := cors.New(cors.Options{
		AllowedOrigins: config.AllowedOrigins,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
		},
	})
	handler := c.Handler(gapi.HttpLogger(httpMux))

	httpServer := &http.Server{
		Handler: handler,
		Addr:    config.HTTPServerAddr,
	}

	group.Go(func() error {
		log.Print("Gateway server running on ", config.HTTPServerAddr)
		err = httpServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		log.Print("gateway gracefully shutting down...")

		err = httpServer.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		log.Print("goodbye")
		return nil
	})

}
