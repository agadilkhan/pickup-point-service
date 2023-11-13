package applicator

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/config"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/controller/http"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/repository"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
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

	err = mainDb.AutoMigrate(&entity.Order{}, &entity.OrderItem{}, &entity.Customer{}, &entity.OrderPickup{}, &entity.Company{}, &entity.PickupPoint{}, &entity.Product{})
	if err != nil {
		l.Panicf("AutoMigrate err: %v", err)
	}

	l.Info("database connection success")

	repo := repository.NewRepository(mainDb, replicaDB)
	deps := pickup.NewDeps(repo, cfg)
	orderService := pickup.NewService(deps)

	endPointHandler := http.NewEndpointHandler(orderService, l)

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
