package main

import (
    "log/slog"
    "os"

    "github.com/alirezazeynali75/string-operator/api"
    "github.com/alirezazeynali75/string-operator/internal/configs"
    "github.com/alirezazeynali75/string-operator/internal/services"
    "github.com/gin-gonic/gin"
)

func main() {
    // Load configurations
    cfg, err := configs.Configure()
    if err != nil {
        slog.With("err", err.Error()).Error("failed to load configurations")
        os.Exit(1)
    }

    // Initialize logger
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

    // Initialize services
    trimSvc := services.NewTrim()
    uppercaseSvc := services.NewUppercase(trimSvc)
    lowercaseSvc := services.NewLowercase(trimSvc)
    reverseSvc := services.NewReverse()
    increaseSvc := services.NewIncrease()

    // Initialize Gin router
    router := gin.Default()

    // Initialize handlers
    handlers := api.NewHandlers(
        logger,
        trimSvc,
        uppercaseSvc,
        lowercaseSvc,
        reverseSvc,
        increaseSvc,
    )

    // Register routes
    handlers.RegisterRoutes(router)

    // Start the server
    address := cfg.Http.Address + ":" + cfg.Http.Port
    logger.Info("starting server", "address", address)
    if err := router.Run(address); err != nil {
        logger.With("err", err.Error()).Error("failed to start server")
        os.Exit(1)
    }
}