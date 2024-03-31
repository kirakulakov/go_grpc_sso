package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "github.com/kirakulakov/go_grpc_sso/internal/grpc/auth"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// MustRun run gRPC server. If error occurred, panics
func (a *App) MustRun() {
	const op = "grpcapp.MustRun"

	log := a.log.With(
		slog.String("op", op))

	log.Info("starting gRPC server ...")

	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	// у каждого лога будет эта строчка (строчки)
	log := a.log.With(
		slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer l.Close()

	log.Info("gRPC server is starting", slog.String("address", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil

}

func (a *App) Stop() error {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op))
	a.log.Info("stopping gRPC server")
	a.gRPCServer.GracefulStop()
	return nil
}
