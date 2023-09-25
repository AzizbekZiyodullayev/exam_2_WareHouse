package handler

import (
	"strconv"
	"warehouse/config"
	"warehouse/pkg/logger"
	"warehouse/storage"

	"github.com/gin-gonic/gin"
)

type handler struct {
	cfg     *config.Config
	logger  logger.LoggerI
	storage storage.StorageI
}

type Response struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

func NewHandler(cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) *handler {
	return &handler{
		cfg:     cfg,
		logger:  logger,
		storage: storage,
	}
}

func (h *handler) handlerResponse(c *gin.Context, path string, code int, message interface{}) {
	response := Response{
		Status:      code,
		Description: path,
		Data:        message,
	}

	switch {
	case code < 300:
		h.logger.Info(path, logger.Any("info", response))
	case code >= 400:
		h.logger.Error(path, logger.Any("error", response))
	}

	c.JSON(code, response)
}

func (h *handler) getOffsetQuery(offset string) (int, error) {

	if len(offset) <= 0 {
		return h.cfg.DefaultOffset, nil
	}

	return strconv.Atoi(offset)
}

func (h *handler) getLimitQuery(limit string) (int, error) {

	if len(limit) <= 0 {
		return h.cfg.DefaultLimit, nil
	}

	return strconv.Atoi(limit)
}
