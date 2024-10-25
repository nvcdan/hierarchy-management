package domain

type DepartmentRepository interface {
	Create(dept *Department) error
	Update(dept *Department) error
	Delete(id int) error
	GetHierarchyByName(name string) ([]*Department, error)
	GetAllHierarchy() ([]*Department, error)
}
