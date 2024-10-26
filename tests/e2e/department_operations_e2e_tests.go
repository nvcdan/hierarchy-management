package e2e

import (
	"bytes"
	"encoding/json"
	"hierarchy-management/internal/db"
	"hierarchy-management/internal/handler"
	"hierarchy-management/internal/repository"
	"hierarchy-management/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	_ = godotenv.Load("../../.env")
	database, _ := db.NewDB()
	repo := repository.NewDepartmentRepository(database)
	deptService := service.NewDepartmentService(repo)

	router := gin.Default()
	handler.NewDepartmentHandler(deptService)
	return router
}

func TestCreateDepartmentE2E(t *testing.T) {
	router := setupRouter()

	dept := map[string]interface{}{
		"name":      "IT Department",
		"parent_id": nil,
		"flags":     1,
	}
	jsonValue, _ := json.Marshal(dept)
	req, _ := http.NewRequest("POST", "/api/departments", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", "Bearer token")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Department created", response["message"])
}
