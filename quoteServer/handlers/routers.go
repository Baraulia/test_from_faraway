package handlers

import (
	"github.com/gin-gonic/gin"
	"test_faraway/quoteServer/services"
)

type Handler struct {
	QuoteService services.QuoteInterface
}

func NewHandler(quoteService services.QuoteInterface) *Handler {
	return &Handler{
		QuoteService: quoteService,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(POWMiddleware)

	router.GET("/quote", h.GetRandomQuote)

	return router
}
