package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type routerHandler interface {
	GetHandler(endpointHandler *EndpointHandler) http.Handler
}

type server struct {
	logger          *zap.SugaredLogger
	shutdownTimeout time.Duration
	client          *http.Server
	listener        net.Listener
	isReady         bool

	EndpointHandler *EndpointHandler
}

func NewServer(
	port int,
	shutdownTimeout time.Duration,
	routerHandler routerHandler,
	logger *zap.SugaredLogger,
	endpointHandler *EndpointHandler,
) (*server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, fmt.Errorf("cannot bind HTTP server '%d':'%v'", port, err)
	}

	return &server{
		logger:          logger,
		shutdownTimeout: shutdownTimeout,
		client: &http.Server{
			Handler: routerHandler.GetHandler(endpointHandler),
		},
		listener: listener,
		isReady:  false,
	}, nil
}

func (s *server) Ready() error {
	if s.isReady {
		return nil
	}

	return errors.New("I am not ready")
}

func (s *server) Stop() error {
	s.isReady = false
	s.logger.Infof("[%s] HTTP server is stopping...", s.listener.Addr().String())

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	s.client.SetKeepAlivesEnabled(false)

	if err := s.client.Shutdown(ctx); err != nil {
		return fmt.Errorf("cannot stop HTTP server: %v", err)
	}

	s.logger.Infof("[%s] HTTP server was stopped", s.listener.Addr().String())

	return nil
}

func (s *server) Run() {
	s.logger.Infof("[%s] HTTP server is running", s.listener.Addr().String())

	go func() {
		s.isReady = true
		s.logger.Infof("[%s] HTTP server is run")

		if err := s.client.Serve(s.listener); err != nil {
			s.isReady = false
			if errors.Is(err, http.ErrServerClosed) {
				return
			}

			s.logger.Errorf("[%s] HTTP server was stopped with err: %v", err)
		}
	}()
}
