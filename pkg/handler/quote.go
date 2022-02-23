package handler

import (
	"net/http"

	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createQuote(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return 
	}

	var input quotes_memorizer.Quote
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Quote.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllQuotes(c *gin.Context) {
	
}

func (h *Handler) getQuoteById(c *gin.Context) {
	
}

func (h *Handler) updateQuote(c *gin.Context) {
	
}

func (h *Handler) deleteQuote(c *gin.Context) {
	
}


