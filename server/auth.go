package server

import (
	authhandler "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authHandler"
	"github.com/TGRZiminiar/go-mc-kafka/modules/auth/authRepository"
	authusecase "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authUsecase"
)

func (s *server) authService() {
	authRepo := authRepository.NewAuthRepository(s.db)
	authUsecase := authusecase.NewAuthUsecase(authRepo)
	authHttpHandler := authhandler.NewAuthHttpHandler(s.cfg, authUsecase)
	authGrpcHandler := authhandler.NewAuthGrpcHandler(authUsecase)

	_ = authHttpHandler
	_ = authGrpcHandler

	auth := s.app.Group("/auth_v1")

	// HealthCheck

	auth.GET("", s.healthCheckService)

}
