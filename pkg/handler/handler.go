package handler

import (
	"github.com/gin-contrib/cors"
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
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization"}

	router.Use(cors.New(config))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.userIdentity)
	{
		works := api.Group("/works")
		{
			works.POST("/", h.onlyInstructor, h.createWork)
			works.GET("/", h.getAllWorks)
			works.GET("/:id", h.signUp)
			works.PUT("/:id", h.signUp)
			works.DELETE("/:id", h.signUp)

			requests := works.Group("/requests")
			{
				requests.POST("/", h.createRequest)
				requests.GET("/", h.signUp)
				requests.GET("/:requests_id", h.signUp)
				requests.PUT("/:requests_id", h.signUp)
				requests.DELETE("/:requests_id", h.signUp)
			}
		}

	}

	return router
}
