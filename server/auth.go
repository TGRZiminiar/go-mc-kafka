package server

import (
	"log"

	authhandler "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authHandler"
	authPb "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authPb"
	"github.com/TGRZiminiar/go-mc-kafka/modules/auth/authRepository"
	authusecase "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authUsecase"
	grpcconn "github.com/TGRZiminiar/go-mc-kafka/pkg/grpcConn"
)

func (s *server) authService() {
	authRepo := authRepository.NewAuthRepository(s.db)
	authUsecase := authusecase.NewAuthUsecase(authRepo)
	authHttpHandler := authhandler.NewAuthHttpHandler(s.cfg, authUsecase)
	authGrpcHandler := authhandler.NewAuthGrpcHandler(authUsecase)

	// Grpc client
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, authGrpcHandler)

		log.Printf("Auth grpc listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	auth := s.app.Group("/auth_v1")

	// HealthCheck

	// auth.GET("", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(s.healthCheckService, []int{1, 0})))
	auth.GET("", s.healthCheckService)
	auth.POST("/auth/login", authHttpHandler.Login)
	auth.POST("/auth/refresh-token", authHttpHandler.RefreshToken)
	auth.POST("/auth/logout", authHttpHandler.Logout)
}
