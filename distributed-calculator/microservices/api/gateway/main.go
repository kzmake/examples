package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	ginlogger "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	adder "github.com/kzmake/examples/distributed-calculator/gen/go/adder/v1"
	divider "github.com/kzmake/examples/distributed-calculator/gen/go/divider/v1"
	greeter "github.com/kzmake/examples/distributed-calculator/gen/go/greeter/v1"
	multiplier "github.com/kzmake/examples/distributed-calculator/gen/go/multiplier/v1"
	subtractor "github.com/kzmake/examples/distributed-calculator/gen/go/subtractor/v1"
)

type Env struct {
	Address string `default:"0.0.0.0:8080"`
	Dapr    struct {
		Address string `default:"localhost:50001"`
	}
}

const prefix = "GATEWAY"

var env Env

func init() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	if err := envconfig.Process(prefix, &env); err != nil {
		log.Fatal().Msgf("%+v", err)
	}

	log.Debug().Msgf("%+v", env)
}

func newGatewayServer(ctx context.Context) (*http.Server, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(ginlogger.SetLogger(
		ginlogger.WithLogger(func(c *gin.Context, _ io.Writer, latency time.Duration) zerolog.Logger {
			return log.Logger.With().
				Timestamp().
				Int("status", c.Writer.Status()).
				Str("method", c.Request.Method).
				Str("path", c.Request.URL.Path).
				Str("ip", c.ClientIP()).
				Dur("latency", latency).
				Str("user_agent", c.Request.UserAgent()).
				Logger()
		}),
		ginlogger.WithSkipPath([]string{"/healthz"}),
	))
	r.Use(gin.Recovery())

	greeterMux := runtime.NewServeMux(
		runtime.WithMetadata(func(_ context.Context, _ *http.Request) metadata.MD {
			return metadata.Pairs("dapr-app-id", "svc-greeter")
		}),
	)
	if err := greeter.RegisterGreeterHandlerFromEndpoint(ctx, greeterMux, env.Dapr.Address, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		return nil, xerrors.Errorf("Failed to register handler: %w", err)
	}
	r.Group("/greeter").Any("/*any", gin.WrapH(greeterMux))

	adderMux := runtime.NewServeMux(
		runtime.WithMetadata(func(_ context.Context, _ *http.Request) metadata.MD {
			return metadata.Pairs("dapr-app-id", "svc-adder")
		}),
	)
	if err := adder.RegisterAdderHandlerFromEndpoint(ctx, adderMux, env.Dapr.Address, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		return nil, xerrors.Errorf("Failed to register handler: %w", err)
	}
	r.Group("/adder").Any("/*any", gin.WrapH(adderMux))

	subtractorMux := runtime.NewServeMux(
		runtime.WithMetadata(func(_ context.Context, _ *http.Request) metadata.MD {
			return metadata.Pairs("dapr-app-id", "svc-subtractor")
		}),
	)
	if err := subtractor.RegisterSubtractorHandlerFromEndpoint(ctx, subtractorMux, env.Dapr.Address, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		return nil, xerrors.Errorf("Failed to register handler: %w", err)
	}
	r.Group("/subtractor").Any("/*any", gin.WrapH(subtractorMux))

	multiplierMux := runtime.NewServeMux(
		runtime.WithMetadata(func(_ context.Context, _ *http.Request) metadata.MD {
			return metadata.Pairs("dapr-app-id", "svc-multiplier")
		}),
	)
	if err := multiplier.RegisterMultiplierHandlerFromEndpoint(ctx, multiplierMux, env.Dapr.Address, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		return nil, xerrors.Errorf("Failed to register handler: %w", err)
	}
	r.Group("/multiplier").Any("/*any", gin.WrapH(multiplierMux))

	dividerMux := runtime.NewServeMux(
		runtime.WithMetadata(func(_ context.Context, _ *http.Request) metadata.MD {
			return metadata.Pairs("dapr-app-id", "svc-divider")
		}),
	)
	if err := divider.RegisterDividerHandlerFromEndpoint(ctx, dividerMux, env.Dapr.Address, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		return nil, xerrors.Errorf("Failed to register handler: %w", err)
	}
	r.Group("/divider").Any("/*any", gin.WrapH(dividerMux))

	return &http.Server{Addr: env.Address, Handler: r}, nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	gatewayS, err := newGatewayServer(ctx)
	if err != nil {
		log.Fatal().Msgf("Failed to build gateway server: %v", err)
	}
	g.Go(func() error { return gatewayS.ListenAndServe() })

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

	ctx, timeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer timeout()

	if err := gatewayS.Shutdown(ctx); err != nil {
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
