package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *EndpointHandler) me(ctx *gin.Context) {
	userID, ok := ctx.Value("user_id").(float64)
	if !ok {
		h.logger.Errorf("error")
		ctx.Status(http.StatusBadRequest)

		return
	}

	png, err := h.service.Me(ctx, int(userID))
	if err != nil {
		h.logger.Errorf("error")
		ctx.Status(http.StatusBadRequest)
	}

	ctx.Data(http.StatusOK, "image/png", png)
}
