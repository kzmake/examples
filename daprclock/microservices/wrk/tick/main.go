package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"

	"github.com/kzmake/examples/daprclock/microservices/wrk/tick/handler"
)

type Env struct {
	Address string `default:"0.0.0.0:5050"`
	Dapr    struct {
		Address string `default:"localhost:50001"`
	}
}

const prefix = "TICK"

var env Env
var conn *grpc.ClientConn

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

func newBindingServer() (common.Service, error) {
	s, err := daprd.NewService(env.Address)
	if err != nil {
		return nil, err
	}

	h := handler.NewTick(conn)
	if err := s.AddBindingInvocationHandler("binding-1s", h.Tick); err != nil {
		return nil, err
	}

	return s, nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	grpcS, err := newBindingServer()
	if err != nil {
		log.Fatal().Msgf("%+v", err)
	}
	g.Go(func() error {
		return grpcS.Start()
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

	_ = grpcS.Stop()

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
