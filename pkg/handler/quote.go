package handler

import (
	"net/http"
	"strconv"

	quotes_memorizer "github.com/Twofold-One/quotes-memorizer-api"
	"github.com/gin-gonic/gin"
)

// @Summary Create Quote
// @Description create new quote
// @ID create-quote
// @Tags quotes
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body quotes_memorizer.Quote true "quote info"
// @Success 200 {integer} Integer 1
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/quotes [post]
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

// @Summary Get All Quotes
// @Description get all quotes
// @ID get-all-quotes
// @Tags quotes
// @Security ApiKeyAuth	 
// @Accept json
// @Produce json
// @Success 200 {object} getAllQuotesResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/quotes [get]
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

// @Summary Get Quote By Id
// @Description get quote by id
// @ID get-quote-by-id
// @Tags quotes
// @Security ApiKeyAuth	 
// @Accept json
// @Produce json
// @Success 200 {object} quotes_memorizer.Quote
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/quotes/:id [get]
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

// @Summary Update Quote
// @Description update quote
// @ID update-quote
// @Tags quotes
// @Security ApiKeyAuth	 
// @Accept json
// @Produce json
// @Param input body quotes_memorizer.Quote true "quote info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/quotes/:id [put]
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

// @Summary Delete Quote
// @Description delete quote
// @ID delete-quote
// @Tags quotes
// @Security ApiKeyAuth	 
// @Accept json
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/quotes/:id [delete]
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


