package applicator

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/agadilkhan/pickup-point-service/internal/user/memory"

	"github.com/agadilkhan/pickup-point-service/internal/user/config"
	"github.com/agadilkhan/pickup-point-service/internal/user/controller/grpc"
	"github.com/agadilkhan/pickup-point-service/internal/user/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/user/repository"
	"go.uber.org/zap"
)

type Applicator struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewApplicator(logger *zap.SugaredLogger, cfg *config.Config) *Applicator {
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
		l.Panicf("cannot connect to mainDB '%s:%s':'%v'", cfg.Database.Main.Host, cfg.Database.Main.Port, err)
	}

	defer func() {
		if err := mainDB.Close(); err != nil {
			l.Panicf("failed to close MainDB err: %v", err)
		}
		l.Info("MainDB closed")
	}()

	replicaDB, err := postgres.New(cfg.Database.Replica)
	if err != nil {
		l.Panicf("cannot connect to replicaDB '%s:%s':'%v'", cfg.Database.Replica.Host, cfg.Database.Replica.Port, err)
	}

	defer func() {
		if err := replicaDB.Close(); err != nil {
			l.Panicf("failed to close ReplicaDB err: %v", err)
		}
		l.Info("ReplicaDB closed")
	}()

	l.Info("database connection success")

	repo := repository.NewRepository(mainDB, replicaDB)
	_ = repo

	userMemory := memory.NewUserMemory(l, repo, time.Second*15)

	userMemory.Run(ctx)

	grpcService := grpc.NewService(l, repo, userMemory)
	grpcServer := grpc.NewServer(cfg.GrpcServer.Port, grpcService, l)
	err = grpcServer.Start()
	if err != nil {
		l.Panicf("failed to start grpc server err: %v", err)
	}
	l.Infof("[%s] gRPC server is run", cfg.Port)

	defer grpcServer.Close()

	app.gracefulShutdown(cancel)
}

func (app *Applicator) gracefulShutdown(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	signal.Stop(ch)
	cancel()
}
