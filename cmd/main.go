package main

import (
	"log"
	"os"

	"hierarchy-management/internal/db"
	"hierarchy-management/internal/handler"
	"hierarchy-management/internal/repository"
	"hierarchy-management/internal/routes"
	"hierarchy-management/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	repo := repository.NewDepartmentRepository(database)
	deptService := service.NewDepartmentService(repo)
	deptHandler := handler.NewDepartmentHandler(deptService)

	router := routes.SetupRouter(deptHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
