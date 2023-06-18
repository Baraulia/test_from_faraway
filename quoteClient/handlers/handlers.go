package client_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetRandomQuote(c *gin.Context) {
	quote, err := h.Service.GetQuote()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"quote": quote})
}
