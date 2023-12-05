package http

import (
	"io/fs"
	"mime"
	"net/http"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/config"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"github.com/agadilkhan/pickup-point-service/swagger"
	"go.uber.org/zap"
)

type EndpointHandler struct {
	service pickup.UseCase
	logger  *zap.SugaredLogger
	cfg     *config.Config
}

func NewEndpointHandler(
	service pickup.UseCase,
	logger *zap.SugaredLogger,
	cfg *config.Config,
) *EndpointHandler {
	return &EndpointHandler{
		service: service,
		logger:  logger,
		cfg:     cfg,
	}
}

type swaggerServer struct {
	openApi http.Handler
}

func (h *EndpointHandler) Swagger() http.Handler {
	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		h.logger.Errorf("AddExtensionType mimetype err: %v", err)
	}

	openApi, err := fs.Sub(swagger.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}

	return &swaggerServer{
		openApi: http.StripPrefix("/swagger/", http.FileServer(http.FS(openApi))),
	}
}

func (sws *swaggerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sws.openApi.ServeHTTP(w, r)
}
