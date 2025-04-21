package api

import (
	"log/slog"

	"github.com/alirezazeynali75/string-operator/internal/services"
	"github.com/gin-gonic/gin"
)


type Handlers struct {
	logger *slog.Logger
	increaseSvc *services.Increase
	trimSvc *services.Trim
	uppercaseSvc *services.Uppercase
	lowercaseSvc *services.Lowercase
	reverseSvc *services.Reverse
}


func NewHandlers(
	logger *slog.Logger,
	trimSvc *services.Trim,
	uppercaseSvc *services.Uppercase,
	lowercaseSvc *services.Lowercase,
	reverseSvc *services.Reverse,
	increaseSvc *services.Increase,
) *Handlers {
	return &Handlers{
		logger: logger.With("handlers"),
		trimSvc: trimSvc,
		uppercaseSvc: uppercaseSvc,
		lowercaseSvc: lowercaseSvc,
		reverseSvc: reverseSvc,
		increaseSvc: increaseSvc,
	}
}

func (h *Handlers) RegisterRoutes(router *gin.Engine) {
    // Register routes for all handlers
    router.POST("/trim", h.Trim)
    router.POST("/uppercase", h.Uppercase)
    router.POST("/uppercase-with-increase", h.UppercaseWithIncrease)
    router.POST("/all", h.ProcessAll)
}