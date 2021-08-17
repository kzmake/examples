package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/kelseyhightower/envconfig"
	grpc_zerolog "github.com/philip-bui/grpc-zerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"

	pb "github.com/kzmake/examples/basic-service-invocation/gen/go/hoge/v1"

	"github.com/kzmake/examples/basic-service-invocation/microservices/svc/hoge/handler"
)

type Env struct {
	Address string `default:"0.0.0.0:5050"`
	Dapr    struct {
		Address string `default:"localhost:50001"`
	}
}

const prefix = "HOGE"

var env Env
var conn grpc.ClientConnInterface

func init() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	if err := envconfig.Process(prefix, &env); err != nil {
		log.Fatal().Msgf("%+v", err)
	}

	var err error
	conn, err = grpc.Dial(env.Dapr.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("%+v", err)
	}

	log.Debug().Msgf("%+v", env)
}

func newGRPCServer() *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zerolog.NewUnaryServerInterceptorWithLogger(&log.Logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterHogeServer(s, handler.NewHoge(conn))

	return s
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	grpcS := newGRPCServer()
	g.Go(func() error {
		lis, err := net.Listen("tcp", env.Address)
		if err != nil {
			return xerrors.Errorf("failed to listen: %w", err)
		}

		return grpcS.Serve(lis)
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	select {
	case <-quit:
		break
	case <-ctx.Done():
		break
	}

	cancel()

	log.Info().Msg("Shutting down server...")

	_, timeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeout()

	grpcS.GracefulStop()

	if err := g.Wait(); err != nil {
		return xerrors.Errorf("failed to shutdown: %w", err)
	}

	log.Info().Msgf("Server exiting")

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal().Msgf("Failed to run server: %v", err)
	}
}
