package client_handlers

import (
	"github.com/gin-gonic/gin"
	client_services "test_faraway/quoteClient/services"
)

type Handler struct {
	Service client_services.ServiceInterface
}

func NewHandler(service client_services.ServiceInterface) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/quote", h.GetRandomQuote)

	return router
}
