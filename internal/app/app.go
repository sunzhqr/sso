package app

import (
	grpcapp "github.com/sunzhqr/sso/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	gRPCPort int,
	storagePath string,
	tokenTTL time.Duration,

) *App {
	// TODO: Init storage
	// TODO: Init auth service
	gRPCApp := grpcapp.New(log, gRPCPort)

	return &App{
		GRPCSrv: gRPCApp,
	}
}
