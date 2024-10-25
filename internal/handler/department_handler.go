package handler

import (
	"hierarchy-management/internal/domain"
	"hierarchy-management/internal/errors"
	"hierarchy-management/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	service service.DepartmentService
}

func NewDepartmentHandler(srv service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{srv}
}

func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var dept domain.Department
	if err := c.ShouldBindJSON(&dept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.NewValidationError("department", "Invalid JSON").Error()})
		return
	}
	err := h.service.CreateDepartment(&dept)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Department created"})
}

func (h *DepartmentHandler) UpdateDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.NewValidationError("id", "Invalid ID format").Error()})
		return
	}
	var dept domain.Department
	if err := c.ShouldBindJSON(&dept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.NewValidationError("department", "Invalid JSON").Error()})
		return
	}
	dept.ID = id
	err = h.service.UpdateDepartment(&dept)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Department updated"})
}

func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.NewValidationError("id", "Invalid ID format").Error()})
		return
	}
	err = h.service.DeleteDepartment(id)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Department deleted"})
}

func (h *DepartmentHandler) GetDepartmentHierarchy(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name parameter is required"})
		return
	}
	hierarchy, err := h.service.GetDepartmentHierarchy(name)
	if err != nil {
		handleError(c, err)
		return
	}

	var response []gin.H
	for _, dept := range hierarchy {
		response = append(response, gin.H{
			"id":          dept.ID,
			"name":        dept.Name,
			"parent_id":   dept.ParentID,
			"is_active":   dept.IsActive(),
			"is_deleted":  dept.IsDeleted(),
			"is_approved": dept.IsApproved(),
		})
	}
	c.JSON(http.StatusOK, response)
}

func (h *DepartmentHandler) GetAllDepartmentsHierarchy(c *gin.Context) {
	hierarchy, err := h.service.GetAllDepartmentsHierarchy()
	if err != nil {
		handleError(c, err)
		return
	}

	var response []gin.H
	for _, dept := range hierarchy {
		response = append(response, gin.H{
			"id":          dept.ID,
			"name":        dept.Name,
			"parent_id":   dept.ParentID,
			"is_active":   dept.IsActive(),
			"is_deleted":  dept.IsDeleted(),
			"is_approved": dept.IsApproved(),
		})
	}
	c.JSON(http.StatusOK, response)
}

func handleError(c *gin.Context, err error) {
	switch e := err.(type) {
	case *errors.InternalError:
		c.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
	case *errors.NotFoundError:
		c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
	case *errors.ValidationError:
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
	}
}
