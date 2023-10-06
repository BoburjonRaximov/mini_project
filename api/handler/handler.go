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
	redis  storage.CacheI
}

func NewHandler(strg storage.StorageI, cfg config.Config, log logger.LoggerI,redis  storage.CacheI) *Handler {
	return &Handler{strg: strg, cfg: cfg, log: log, redis: redis}
}
