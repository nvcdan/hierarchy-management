package handler

import (
	"hierarchy-management/internal/domain"
	"hierarchy-management/internal/errors"
	"hierarchy-management/internal/response"
	"hierarchy-management/internal/service"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	service service.DepartmentService
}

type DepartmentNode struct {
	ID         int               `json:"id"`
	Name       string            `json:"name"`
	ParentID   *int              `json:"parent_id,omitempty"`
	IsActive   bool              `json:"is_active"`
	IsDeleted  bool              `json:"is_deleted"`
	IsApproved bool              `json:"is_approved"`
	Children   []*DepartmentNode `json:"children,omitempty"`
}

func NewDepartmentHandler(srv service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{srv}
}

func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var dept domain.Department
	if err := c.ShouldBindJSON(&dept); err != nil {
		response.HandleError(c, errors.NewValidationError("department", "Invalid JSON"))
		return
	}
	err := h.service.CreateDepartment(&dept)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response.APIResponse{
		IsSuccess: true,
		Message:   "Department created successfully",
	})
}

func (h *DepartmentHandler) UpdateDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.HandleError(c, errors.NewValidationError("id", "Invalid ID format"))
		return
	}
	var dept domain.Department
	if err := c.ShouldBindJSON(&dept); err != nil {
		response.HandleError(c, errors.NewValidationError("department", "Invalid JSON"))
		return
	}
	dept.ID = id
	err = h.service.UpdateDepartment(&dept)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response.APIResponse{
		IsSuccess: true,
		Message:   "Department updated successfully",
	})
}

func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.HandleError(c, errors.NewValidationError("id", "Invalid ID format"))
		return
	}
	err = h.service.DeleteDepartment(id)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response.APIResponse{
		IsSuccess: true,
		Message:   "Department deleted successfully",
	})
}

func (h *DepartmentHandler) GetDepartmentHierarchy(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		response.HandleError(c, errors.NewValidationError("name", "Query parameter 'name' is required"))
		return
	}

	hierarchy, err := h.service.GetDepartmentHierarchy(name)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	deptMap := make(map[int]*DepartmentNode)
	for _, dept := range hierarchy {
		node := &DepartmentNode{
			ID:         dept.ID,
			Name:       dept.Name,
			ParentID:   dept.ParentID,
			IsActive:   dept.IsActive(),
			IsDeleted:  dept.IsDeleted(),
			IsApproved: dept.IsApproved(),
		}
		deptMap[dept.ID] = node
	}

	var roots []*DepartmentNode
	for _, node := range deptMap {
		if node.ParentID != nil {
			parentNode, ok := deptMap[*node.ParentID]
			if ok {
				parentNode.Children = append(parentNode.Children, node)
			}
		}
		if node.Name == name {
			roots = append(roots, node)
		}
	}

	if len(roots) == 0 {
		response.HandleError(c, errors.NewNotFoundError("department"))
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		IsSuccess: true,
		Message:   "Department hierarchy retrieved successfully",
		Data:      roots,
	})
}

func (h *DepartmentHandler) GetAllDepartmentsHierarchy(c *gin.Context) {
	hierarchy, err := h.service.GetAllDepartmentsHierarchy()
	if err != nil {
		response.HandleError(c, err)
		return
	}

	deptMap := make(map[int]*DepartmentNode)
	for _, dept := range hierarchy {
		node := &DepartmentNode{
			ID:         dept.ID,
			Name:       dept.Name,
			ParentID:   dept.ParentID,
			IsActive:   dept.IsActive(),
			IsDeleted:  dept.IsDeleted(),
			IsApproved: dept.IsApproved(),
		}
		deptMap[dept.ID] = node
	}

	var roots []*DepartmentNode
	for _, node := range deptMap {
		if node.ParentID != nil {
			parentNode, ok := deptMap[*node.ParentID]
			if ok {
				parentNode.Children = append(parentNode.Children, node)
			}
		} else {
			roots = append(roots, node)
		}
	}

	c.JSON(http.StatusOK, response.APIResponse{
		IsSuccess: true,
		Message:   "Department hierarchy retrieved successfully",
		Data:      roots,
	})
}
