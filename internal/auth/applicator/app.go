package applicator

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/auth/auth"
	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
	"github.com/agadilkhan/pickup-point-service/internal/auth/controller/http"
	"github.com/agadilkhan/pickup-point-service/internal/auth/database/postgres"
	"github.com/agadilkhan/pickup-point-service/internal/auth/repository"
	"github.com/agadilkhan/pickup-point-service/internal/auth/transport"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type Applicator struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewApplicator(logger *zap.SugaredLogger, cfg *config.Config) *Applicator {
	return &Applicator{
		logger: logger,
		cfg:    cfg,
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
		l.Panicf("cannot connect to mainDB '%s:%d: %v", cfg.Database.Main.Host, cfg.Database.Main.Port, err)
	}

	defer func() {
		if err := mainDB.Close(); err != nil {
			l.Panicf("failed to close MainDB err: %v", err)
		}
		l.Infof("mainDB closed")
	}()

	replicaDB, err := postgres.New(cfg.Database.Replica)
	if err != nil {
		l.Panicf("cannot connect to replicaDB '%s:%d: %v", cfg.Database.Replica.Host, cfg.Database.Replica.Port, err)
	}

	defer func() {
		if err := replicaDB.Close(); err != nil {
			l.Panicf("failed to close ReplicaDB err: %v", err)
		}
		l.Infof("replicaDB closed")
	}()

	l.Infof("db connection success")

	repo := repository.NewRepository(mainDB, replicaDB)

	userTransport := transport.NewUserTransport(cfg.User)

	authService := auth.NewAuthService(repo, cfg.Auth, userTransport)

	endPointHandler := http.NewEndpointHandler(authService, l, cfg.PasswordSecretKey)

	router := http.NewRouter(l)
	httpConfig := cfg.HttpServer
	server, err := http.NewServer(httpConfig.Port, httpConfig.ShutdownTimeout, router, l, endPointHandler)
	if err != nil {
		l.Panicf("failed to create server: %v", err)
	}

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
