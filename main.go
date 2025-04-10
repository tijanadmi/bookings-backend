package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tijanadmi/bookings_backend/api"
	db "github.com/tijanadmi/bookings_backend/db/sqlc"
	_ "github.com/tijanadmi/bookings_backend/doc/statik"
	"github.com/tijanadmi/bookings_backend/gapi"
	"github.com/tijanadmi/bookings_backend/pb"
	"github.com/tijanadmi/bookings_backend/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(connPool)

	//runGinServer(config, store)
	go runGatewayServer(config, store)
	runGrpcServer(config, store)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

// func runTaskProcessor(
// 	ctx context.Context,
// 	waitGroup *errgroup.Group,
// 	config util.Config,
// 	redisOpt asynq.RedisClientOpt,
// 	store db.Store,
// ) {
// 	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
// 	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)

// 	log.Info().Msg("start task processor")
// 	err := taskProcessor.Start()
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("failed to start task processor")
// 	}

// 	waitGroup.Go(func() error {
// 		<-ctx.Done()
// 		log.Info().Msg("graceful shutdown task processor")

// 		taskProcessor.Shutdown()
// 		log.Info().Msg("task processor is stopped")

// 		return nil
// 	})
// }

func runGrpcServer(
	config util.Config,
	store db.Store,
) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	gprcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(gprcLogger)
	pb.RegisterBookingsServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Error().Err(err).Msg("cannot start gRPC server")
	}
}

func runGatewayServer(
	config util.Config,
	store db.Store,
) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(
		jsonOption,
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			if key == "set-cookie" {
				return "Set-Cookie", true
			}
			return runtime.DefaultHeaderMatcher(key)
		}),
	)
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	err = pb.RegisterBookingsHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// Dodaj ovaj handler — ovo je novi deo
	mux.HandleFunc("/upload-room-image", server.UploadRoomImageHandler)

	// Serve static files from ./uploads directory
	fsUploads := http.FileServer(http.Dir("./uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fsUploads))


	// Serve Swagger statičke fajlove

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik fs")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer( /*http.Dir("./doc/swagger")*/ statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	// Configure CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Ovde dozvoli specifične ili sve origne
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization", "Cookie"},
		ExposedHeaders:   []string{"Content-Length", "Set-Cookie"},
		AllowCredentials: true,
	}).Handler(mux)

	// httpServer := &http.Server{
	// 	Handler: gapi.HttpLogger(mux),
	// 	Addr:    config.HTTPServerAddress,
	// }

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
	handler := gapi.HttpLogger(corsHandler)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
