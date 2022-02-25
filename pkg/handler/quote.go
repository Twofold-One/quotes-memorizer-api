package handler

import (
	"net/http"
	"strconv"

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

type getAllQuotesResponse struct {
	Data []quotes_memorizer.Quote `json:"data"`
}

func (h *Handler) getAllQuotes(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return 
	}

	quotes, err := h.services.Quote.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllQuotesResponse{
		Data: quotes,
	})

}

func (h *Handler) getQuoteById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return 
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	quote, err := h.services.Quote.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, quote)
}

func (h *Handler) updateQuote(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return 
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input quotes_memorizer.UpdateQuoteInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	if err := h.services.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteQuote(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return 
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Quote.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}


