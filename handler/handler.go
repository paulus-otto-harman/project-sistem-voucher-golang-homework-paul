package handler

import (
	"project/service"

	"go.uber.org/zap"
)

type Handler struct {
}

func NewHandler(service service.Service, logger *zap.Logger) *Handler {
	return &Handler{}
}
