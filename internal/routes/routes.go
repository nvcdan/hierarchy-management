package routes

import (
	"hierarchy-management/internal/handler"
	"hierarchy-management/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(deptHandler *handler.DepartmentHandler, authHandler *handler.AuthHandler) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	router.Use(cors.New(config))

	api := router.Group("/api")
	{
		api.POST("/login", authHandler.Login)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/departments/create", deptHandler.CreateDepartment)
			protected.PUT("/departments/:id/update", deptHandler.UpdateDepartment)
			protected.DELETE("/departments/:id/delete", deptHandler.DeleteDepartment)
			protected.GET("/departments/hierarchy", deptHandler.GetDepartmentHierarchy)
			protected.GET("/departments/hierarchy/all", deptHandler.GetAllDepartmentsHierarchy)
		}
	}

	return router
}
