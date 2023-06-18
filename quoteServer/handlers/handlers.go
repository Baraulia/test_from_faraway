package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetRandomQuote(c *gin.Context) {
	quote, err := h.QuoteService.GetRandomQuote()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, quote.QuoteText)
}
