package domain

type DepartmentRepository interface {
	Create(dept *Department) error
	IsDuplicateDepartmentName(name string) bool
	ExistsByID(id int) bool
	Update(dept *Department) error
	Delete(id int) error
	GetHierarchyByName(name string) ([]*Department, error)
	GetAllHierarchy() ([]*Department, error)
}
