package server

import (
	"time"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	middlewarehandler "github.com/TGRZiminiar/go-mc-kafka/modules/middleware/middlewareHandler"
	middlewarerepository "github.com/TGRZiminiar/go-mc-kafka/modules/middleware/middlewareRepository"
	middlewareusecase "github.com/TGRZiminiar/go-mc-kafka/modules/middleware/middlewareUsecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	server struct {
		app        *echo.Echo
		db         *mongo.Client
		cfg        *config.Config
		middleware middlewarehandler.MiddlewareHandlerService
	}
)

func newMiddleware(cfg *config.Config) middlewarehandler.MiddlewareHandlerService {
	repo := middlewarerepository.NewMiddlewareRepository()
	usecase := middlewareusecase.NewMiddlewareUsecase(repo)
	return middlewarehandler.NewMiddlewareHandler(cfg, usecase)
}

func Start(cfg *config.Config, db *mongo.Client) {
	s := &server{
		app:        echo.New(),
		db:         db,
		cfg:        cfg,
		middleware: newMiddleware(cfg),
	}

	// Basic Middleware

	// Request Timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		Timeout:      30 * time.Second,
		ErrorMessage: "Error: Request Timeout",
	}))

	// Cors
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PATCH, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Body Limit
	s.app.Use(middleware.BodyLimit("10M"))

	switch s.cfg.App.Name {
	case "auth":
	case "player":
	case "item":
	case "inventory":
	case "payment":
	}

	// Graceful Shutdown

}
