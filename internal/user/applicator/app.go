package applicator

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/user/config"
	"github.com/agadilkhan/pickup-point-service/internal/user/controller/grpc"
	"github.com/agadilkhan/pickup-point-service/internal/user/controller/http"
	"github.com/agadilkhan/pickup-point-service/internal/user/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/user/repository"
	"github.com/agadilkhan/pickup-point-service/internal/user/user"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type Applicator struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewAplicator(logger *zap.SugaredLogger, cfg *config.Config) *Applicator {
	return &Applicator{
		logger,
		cfg,
	}
}

func (app *Applicator) Run() {
	var (
		cfg = app.cfg
		l   = app.logger
	)

	ctx, cancel := context.WithCancel(context.TODO())
	_ = ctx

	mainDB, err := postgres.New(cfg.Database.Main)
	if err != nil {
		l.Panicf("cannot connect to mainDB '%s:%d':'%v'", cfg.Database.Main.Host, cfg.Database.Main.Port, err)
	}

	defer func() {
		if err := mainDB.Close(); err != nil {
			l.Panicf("failed to close MainDB err: %v", err)
		}
		l.Infof("MainDB closed")
	}()

	replicaDB, err := postgres.New(cfg.Database.Replica)
	if err != nil {
		l.Panicf("cannot connect to replicaDB '%s:%d':'%v'", cfg.Database.Replica.Host, cfg.Database.Replica.Port, err)
	}

	defer func() {
		if err := replicaDB.Close(); err != nil {
			l.Panicf("failed to close ReplicaDB err: %v", err)
		}
		l.Infof("ReplicaDB closed")
	}()

	l.Infof("database connection success")

	repo := repository.NewRepository(mainDB, replicaDB)
	_ = repo

	userService := user.NewService(repo)

	endpointHandler := http.NewEndpointHandler(userService, l)

	router := http.NewRouter(l)
	httpCfg := cfg.HttpServer
	server, err := http.NewServer(httpCfg.Port, httpCfg.ShutdownTimeout, router, l, endpointHandler)
	if err != nil {
		l.Panicf("failed to create server err: %v", err)
	}

	grpcService := grpc.NewService(l, repo)
	grpcServer := grpc.NewServer(cfg.GrpcServer.Port, grpcService)
	err = grpcServer.Start()
	if err != nil {
		l.Panicf("failed to start grpc server err: %v", err)
	}

	defer grpcServer.Close()

	server.Run()

	defer func() {
		if err := server.Stop(); err != nil {
			l.Panicf("failed close server err: %v", err)
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
