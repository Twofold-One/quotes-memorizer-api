package handler

import (
	"github.com/Twofold-One/quotes-memorizer-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.singIn)
	}

	api := router.Group("/api")
	{
		quotes := api.Group("/quotes")
		{
			quotes.POST("/", h.createQuote)
			quotes.GET("/", h.getAllQuotes)
			quotes.GET("/:id", h.getQuoteById)
			quotes.PUT("/:id", h.updateQuote)
			quotes.DELETE("/id", h.deleteQuote)
		}
	}
	return router
}