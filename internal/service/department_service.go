package service

import (
	"hierarchy-management/internal/domain"
)

type DepartmentService interface {
	CreateDepartment(dept *domain.Department) error
	UpdateDepartment(dept *domain.Department) error
	DeleteDepartment(id int) error
	GetDepartmentHierarchy(name string) ([]*domain.Department, error)
	GetAllDepartmentsHierarchy() ([]*domain.Department, error)
}

type departmentService struct {
	repo domain.DepartmentRepository
}

func NewDepartmentService(repo domain.DepartmentRepository) DepartmentService {
	return &departmentService{repo}
}

func (s *departmentService) CreateDepartment(dept *domain.Department) error {
	return s.repo.Create(dept)
}

func (s *departmentService) UpdateDepartment(dept *domain.Department) error {
	return s.repo.Update(dept)
}

func (s *departmentService) DeleteDepartment(id int) error {
	return s.repo.Delete(id)
}

func (s *departmentService) GetDepartmentHierarchy(name string) ([]*domain.Department, error) {
	return s.repo.GetHierarchyByName(name)
}

func (s *departmentService) GetAllDepartmentsHierarchy() ([]*domain.Department, error) {
	return s.repo.GetAllHierarchy()
}
