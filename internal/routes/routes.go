package routes

import (
	"hierarchy-management/internal/handler"
	"hierarchy-management/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(deptHandler *handler.DepartmentHandler) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	router.Use(cors.New(config))
	router.Use(middleware.AuthMiddleware())

	api := router.Group("/api")
	{
		api.POST("/departments/create", deptHandler.CreateDepartment)
		api.PUT("/departments/:id/update", deptHandler.UpdateDepartment)
		api.DELETE("/departments/:id/delete", deptHandler.DeleteDepartment)
		api.GET("/departments/hierarchy", deptHandler.GetDepartmentHierarchy)
		api.GET("/departments/hierarchy/all", deptHandler.GetAllDepartmentsHierarchy)
	}

	return router
}
