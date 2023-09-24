package handler

import (
	"new_project/config"
	"new_project/pkg/logger"
	"new_project/storage"
)

type Handler struct {
	strg storage.StorageI
	cfg  config.Config
	log  logger.LoggerI
}

func NewHandler(strg storage.StorageI, cfg config.Config, log logger.LoggerI) *Handler {
	return &Handler{strg: strg, cfg: cfg, log: log}
}
