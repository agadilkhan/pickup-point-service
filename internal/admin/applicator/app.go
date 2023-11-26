package applicator

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/agadilkhan/pickup-point-service/internal/admin/admin"
	"github.com/agadilkhan/pickup-point-service/internal/admin/config"
	"github.com/agadilkhan/pickup-point-service/internal/admin/controller/http"
	"github.com/agadilkhan/pickup-point-service/internal/admin/transport"
	"go.uber.org/zap"
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

	userGrpcTransport := transport.NewUserGrpcTransport(cfg.UserGrpcTransport)

	adminService := admin.NewService(userGrpcTransport)

	endpointHandler := http.NewEndpointHandler(adminService, l)

	router := http.NewRouter(l)
	httpConfig := cfg.HTTPServer
	server, err := http.NewServer(httpConfig.Port, httpConfig.ShutdownTimeout, router, l, endpointHandler)
	if err != nil {
		l.Panicf("failed to create server: %v", err)
	}

	server.Run()

	defer func() {
		if err = server.Stop(); err != nil {
			l.Panicf("failed to close server: %v", err)
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
