package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/diplomaWorks/pkg/service"
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
		auth.POST("/sign-in", h.signUp)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api")
	{
		works := api.Group("/works")
		{
			works.POST("/", h.signUp)
			works.GET("/", h.signUp)
			works.GET("/:id", h.signUp)
			works.PUT("/:id", h.signUp)
			works.DELETE("/:id", h.signUp)

			requests := works.Group("/requests")
			{
				requests.POST("/", h.signUp)
				requests.GET("/", h.signUp)
				requests.GET("/:requests_id", h.signUp)
				requests.PUT("/:requests_id", h.signUp)
				requests.DELETE("/:requests_id", h.signUp)
			}
		}

	}

	return router
}
