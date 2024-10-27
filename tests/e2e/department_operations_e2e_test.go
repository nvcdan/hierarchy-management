package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hierarchy-management/internal/db"
	"hierarchy-management/internal/handler"
	"hierarchy-management/internal/repository"
	"hierarchy-management/internal/routes"
	"hierarchy-management/internal/service"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	_ = godotenv.Load("../../.env")

	os.Setenv("DB_HOST", os.Getenv("TEST_DB_HOST"))
	os.Setenv("DB_PORT", os.Getenv("TEST_DB_PORT"))
	os.Setenv("DB_USER", os.Getenv("TEST_DB_USER"))
	os.Setenv("DB_PASS", os.Getenv("TEST_DB_PASS"))
	os.Setenv("DB_NAME", os.Getenv("TEST_DB_NAME"))
	os.Setenv("JWT_TOKEN", os.Getenv("TEST_JWT_TOKEN"))

	database, _ := db.NewDB()
	deptRepo := repository.NewDepartmentRepository(database)
	userRepo := repository.NewUserRepository()

	deptService := service.NewDepartmentService(deptRepo)
	userService := service.NewUserService(userRepo)

	deptHandler := handler.NewDepartmentHandler(deptService)
	authHandler := handler.NewAuthHandler(userService)
	router := routes.SetupRouter(deptHandler, authHandler)
	return router
}

func getAuthTokenFromEnv() string {
	token := os.Getenv("JWT_TOKEN")
	if token == "" {
		fmt.Println("JWT_TOKEN not found in environment")
	}
	return "Bearer " + token
}

func TestCreateDepartmentE2E(t *testing.T) {
	router := setupRouter()
	authToken := getAuthTokenFromEnv()

	dept := map[string]interface{}{
		"name":  "IT Department",
		"flags": 1,
	}
	jsonValue, _ := json.Marshal(dept)
	req, _ := http.NewRequest("POST", "/api/departments/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Logf("Response Body: %s", w.Body.String())
	}

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, true, response["isSuccess"])
	assert.Equal(t, "Department created successfully", response["message"])
}

func TestUpdateDepartmentE2E(t *testing.T) {
	router := setupRouter()
	authToken := getAuthTokenFromEnv()

	dept := map[string]interface{}{
		"name":  "HR Department",
		"flags": 1,
	}
	jsonValue, _ := json.Marshal(dept)
	req, _ := http.NewRequest("POST", "/api/departments/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	deptID := 1

	deptUpdate := map[string]interface{}{
		"name":  "Human Resources",
		"flags": 1,
	}
	jsonValueUpdate, _ := json.Marshal(deptUpdate)
	reqUpdate, _ := http.NewRequest("PUT", "/api/departments/"+strconv.Itoa(deptID)+"/update", bytes.NewBuffer(jsonValueUpdate))
	reqUpdate.Header.Set("Authorization", authToken)
	reqUpdate.Header.Set("Content-Type", "application/json")
	wUpdate := httptest.NewRecorder()
	router.ServeHTTP(wUpdate, reqUpdate)

	if wUpdate.Code != http.StatusOK {
		t.Logf("Response Body: %s", wUpdate.Body.String())
	}

	assert.Equal(t, http.StatusOK, wUpdate.Code)

	var respUpdate map[string]interface{}
	json.Unmarshal(wUpdate.Body.Bytes(), &respUpdate)
	assert.Equal(t, true, respUpdate["isSuccess"])
	assert.Equal(t, "Department updated successfully", respUpdate["message"])
}

func TestDeleteDepartmentE2E(t *testing.T) {
	router := setupRouter()
	authToken := getAuthTokenFromEnv()

	dept := map[string]interface{}{
		"name":  "Marketing Department",
		"flags": 1,
	}
	jsonValue, _ := json.Marshal(dept)
	req, _ := http.NewRequest("POST", "/api/departments/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	deptID := 1

	reqDelete, _ := http.NewRequest("DELETE", "/api/departments/"+strconv.Itoa(deptID)+"/delete", nil)
	reqDelete.Header.Set("Authorization", authToken)
	wDelete := httptest.NewRecorder()
	router.ServeHTTP(wDelete, reqDelete)

	if wDelete.Code != http.StatusOK {
		t.Logf("Response Body: %s", wDelete.Body.String())
	}

	assert.Equal(t, http.StatusOK, wDelete.Code)

	var respDelete map[string]interface{}
	json.Unmarshal(wDelete.Body.Bytes(), &respDelete)
	assert.Equal(t, true, respDelete["isSuccess"])
	assert.Equal(t, "Department deleted successfully", respDelete["message"])
}

func TestGetDepartmentHierarchyE2E(t *testing.T) {
	router := setupRouter()
	authToken := getAuthTokenFromEnv()

	parentDept := map[string]interface{}{
		"name":  "Parent Department",
		"flags": 1,
	}
	jsonParent, _ := json.Marshal(parentDept)
	reqParent, _ := http.NewRequest("POST", "/api/departments/create", bytes.NewBuffer(jsonParent))
	reqParent.Header.Set("Authorization", authToken)
	reqParent.Header.Set("Content-Type", "application/json")
	wParent := httptest.NewRecorder()
	router.ServeHTTP(wParent, reqParent)
	assert.Equal(t, http.StatusCreated, wParent.Code)

	parentID := 1

	childDept := map[string]interface{}{
		"name":      "Child Department",
		"parent_id": parentID,
		"flags":     1,
	}
	jsonChild, _ := json.Marshal(childDept)
	reqChild, _ := http.NewRequest("POST", "/api/departments/create", bytes.NewBuffer(jsonChild))
	reqChild.Header.Set("Authorization", authToken)
	reqChild.Header.Set("Content-Type", "application/json")
	wChild := httptest.NewRecorder()
	router.ServeHTTP(wChild, reqChild)
	assert.Equal(t, http.StatusCreated, wChild.Code)

	reqGet, _ := http.NewRequest("GET", "/api/departments/hierarchy?name=Parent Department", nil)
	reqGet.Header.Set("Authorization", authToken)
	wGet := httptest.NewRecorder()
	router.ServeHTTP(wGet, reqGet)

	if wGet.Code != http.StatusOK {
		t.Logf("Response Body: %s", wGet.Body.String())
	}

	assert.Equal(t, http.StatusOK, wGet.Code)

	var respGet map[string]interface{}
	json.Unmarshal(wGet.Body.Bytes(), &respGet)
	assert.Equal(t, true, respGet["isSuccess"])
	assert.Equal(t, "Department hierarchy retrieved successfully", respGet["message"])
	assert.NotEmpty(t, respGet["data"])
}

func TestGetAllDepartmentsHierarchyE2E(t *testing.T) {
	router := setupRouter()
	authToken := getAuthTokenFromEnv()

	departments := []map[string]interface{}{
		{
			"name":  "Finance Department",
			"flags": 1,
		},
		{
			"name":  "Sales Department",
			"flags": 1,
		},
	}

	for _, dept := range departments {
		jsonValue, _ := json.Marshal(dept)
		req, _ := http.NewRequest("POST", "/api/departments/create", bytes.NewBuffer(jsonValue))
		req.Header.Set("Authorization", authToken)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}

	reqGet, _ := http.NewRequest("GET", "/api/departments/hierarchy/all", nil)
	reqGet.Header.Set("Authorization", authToken)
	wGet := httptest.NewRecorder()
	router.ServeHTTP(wGet, reqGet)

	if wGet.Code != http.StatusOK {
		t.Logf("Response Body: %s", wGet.Body.String())
	}

	assert.Equal(t, http.StatusOK, wGet.Code)

	var respGet map[string]interface{}
	json.Unmarshal(wGet.Body.Bytes(), &respGet)
	assert.Equal(t, true, respGet["isSuccess"])
	assert.Equal(t, "Department hierarchy retrieved successfully", respGet["message"])
	assert.NotEmpty(t, respGet["data"])
}
