package handl

import (
	"config_service/internal/service"
	"log/slog"
)

type Handler struct {
	proto.UnimplementedHandler //FIXME: fix with protos
	service                    *service.Service
	logger                     *slog.Logger
}

func NewHandler(service *service.Service, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
