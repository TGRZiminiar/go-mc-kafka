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

	// Grpc ckient
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, authGrpcHandler)

		log.Printf("Auth grpc listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = authHttpHandler
	_ = authGrpcHandler

	auth := s.app.Group("/auth_v1")

	// HealthCheck

	auth.GET("", s.healthCheckService)

}
