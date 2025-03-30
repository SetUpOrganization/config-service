package config_service

import (
	"config_service/internal/config"
	"config_service/internal/repo"
	"config_service/internal/service"
	handl "config_service/internal/transport/grpc"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

// TODO: NewServer() + Start
// for NewServer() - NewService(), for NewService() - NewClient-s() and NewHandler
type Server struct {
	grpcServer *grpc.Server
	cfg        *config.Config
	repo       *repo.Repo
	logger     *slog.Logger
}

func NewServer(cfg *config.Config, repo *repo.Repo, service *service.Service, logger *slog.Logger) (*Server, error) {
	grpcServer := grpc.NewServer()

	handler := handl.NewHandler(service, logger)
	proto.RegisterConfigServiceServer(grpcServer, handler)

	return &Server{
		grpcServer: grpcServer,
		cfg:        cfg,
		repo:       repo,
		logger:     logger,
	}, nil
}

func (s *Server) Start() error {
	s.logger.Info("Starting server")
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", s.cfg.Servers.HTTP.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	if err := s.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil

}
