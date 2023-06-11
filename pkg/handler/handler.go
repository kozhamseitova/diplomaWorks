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
	}

	api := router.Group("/api", h.userIdentity)
	{
		works := api.Group("/works")
		{
			works.GET("/", h.getAllWorks)
			works.GET("/:id", h.getWorkById)
		}

		admin := api.Group("/admin", h.onlyAdmin)
		{
			admin.POST("/sign-up", h.onlyAdmin, h.signUp)
			adminWorks := admin.Group("/works")
			{
				adminWorks.GET("/", h.getAllWorksForAdmin)
				adminWorks.PUT("/:id", h.approveWork)
			}
		}

		students := api.Group("/students/:student_id", h.onlyStudent)
		{
			requests := students.Group("/requests")
			{
				requests.POST("/", h.createRequest)
				requests.GET("/", h.getAllRequestsByStudentId)
				requests.PUT("/:requests_id", h.closeRequest)
				requests.DELETE("/:requests_id", h.deleteRequest)
			}
		}

		instructors := api.Group("/instructors/:instructor_id", h.onlyInstructor)
		{
			instructorsWorks := instructors.Group("/works")
			{
				instructorsWorks.POST("/", h.createWork)
				instructorsWorks.GET("/", h.getWorksByInstructorId)
				instructorsWorks.PUT("/:id", h.updateWork)
				instructorsWorks.DELETE("/:id", h.deleteWork)
			}

			requests := instructors.Group("/requests")
			{
				requests.GET("/:work_id", h.getAllRequestsByWorkId)
				requests.PUT("/:requests_id", h.changeRequestStatus)
			}
		}

	}

	return router
}
