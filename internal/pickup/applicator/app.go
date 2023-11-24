package applicator

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/config"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/controller/http"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/repository"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/transport"
	"go.uber.org/zap"
)

type Applicator struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewApplicator(cfg *config.Config, logger *zap.SugaredLogger) *Applicator {
	return &Applicator{
		cfg:    cfg,
		logger: logger,
	}
}

func (app *Applicator) Run() {
	var (
		cfg = app.cfg
		l   = app.logger
	)

	ctx, cancel := context.WithCancel(context.TODO())
	_ = ctx

	mainDb, err := postgres.New(cfg.Main)
	if err != nil {
		l.Panicf("cannot connect to mainDB: %v", err)
	}

	defer func() {
		if err := mainDb.Close(); err != nil {
			l.Panicf("failed to close mainDB: %v", err)
		}
		l.Info("mainDB closed")
	}()

	replicaDB, err := postgres.New(cfg.Replica)
	if err != nil {
		l.Panicf("cannot connect to replicaDB: %v", err)
	}

	defer func() {
		if err := replicaDB.Close(); err != nil {
			l.Panicf("failed to close replicaDB: %v", err)
		}
		l.Info("replicaDB closed")
	}()

	l.Info("database connection success")

	repo := repository.NewRepository(mainDb, replicaDB)

	userGrpcTransport := transport.NewUserGrpcTransport(cfg.UserGrpc)

	pickupService := pickup.NewPickupService(repo, userGrpcTransport)

	endPointHandler := http.NewEndpointHandler(pickupService, l, cfg)

	router := http.NewRouter(l)
	httpConfig := cfg.HttpServer
	server, err := http.NewServer(httpConfig.Port, httpConfig.ShutdownTimeout, router, l, endPointHandler)
	if err != nil {
		l.Panicf("failed to create HTTP server: %v", err)
	}

	server.Run()

	defer func() {
		if err := server.Stop(); err != nil {
			l.Panicf("failed to close server err: %v", err)
		}
		l.Info("server closed")
	}()

	app.gracefulShutdown(cancel)
}

func (app *Applicator) gracefulShutdown(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	signal.Stop(ch)
	cancel()
}
