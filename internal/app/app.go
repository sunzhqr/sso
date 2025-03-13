package app

import (
	grpcapp "github.com/sunzhqr/sso/internal/app/grpc"
	"github.com/sunzhqr/sso/internal/services/auth"
	"github.com/sunzhqr/sso/internal/storage/sqlite"
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
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}
	authService := auth.New(log, storage, storage, storage, tokenTTL)
	gRPCApp := grpcapp.New(log, authService, gRPCPort)

	return &App{
		GRPCSrv: gRPCApp,
	}
}
