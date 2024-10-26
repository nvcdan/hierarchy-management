package repository

import (
	"database/sql"
	"hierarchy-management/internal/domain"
)

type departmentRepository struct {
	db *sql.DB
}

func NewDepartmentRepository(db *sql.DB) domain.DepartmentRepository {
	return &departmentRepository{db}
}

func (r *departmentRepository) Create(dept *domain.Department) error {
	_, err := r.db.Exec("CALL CreateDepartment(?, ?, ?)", dept.Name, dept.ParentID, dept.Flags)
	return err
}

func (r *departmentRepository) IsDuplicateDepartmentName(name string) bool {
	var isDuplicate bool
	err := r.db.QueryRow("CALL IsDuplicateDepartmentName(?)", name).Scan(&isDuplicate)
	if err != nil && err != sql.ErrNoRows {
		return false
	}
	return isDuplicate
}

func (r *departmentRepository) ExistsByID(id int) bool {
	var idExists bool
	err := r.db.QueryRow("CALL ExistsByID(?)", id).Scan(&idExists)
	if err != nil && err != sql.ErrNoRows {
		return false
	}
	return idExists
}

func (r *departmentRepository) Update(dept *domain.Department) error {
	_, err := r.db.Exec("CALL UpdateDepartment(?, ?, ?, ?)", dept.ID, dept.Name, dept.ParentID, dept.Flags)
	return err
}

func (r *departmentRepository) Delete(id int) error {
	_, err := r.db.Exec("CALL DeleteDepartment(?)", id)
	return err
}

func (r *departmentRepository) GetHierarchyByName(name string) ([]*domain.Department, error) {
	rows, err := r.db.Query("CALL GetDepartmentHierarchy(?)", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []*domain.Department
	for rows.Next() {
		var dept domain.Department
		err := rows.Scan(&dept.ID, &dept.Name, &dept.ParentID, &dept.Flags)
		if err != nil {
			return nil, err
		}
		departments = append(departments, &dept)
	}
	return departments, nil
}

func (r *departmentRepository) GetAllHierarchy() ([]*domain.Department, error) {
	rows, err := r.db.Query("CALL GetAllDepartmentsHierarchy()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []*domain.Department
	for rows.Next() {
		var dept domain.Department
		err := rows.Scan(&dept.ID, &dept.Name, &dept.ParentID, &dept.Flags)
		if err != nil {
			return nil, err
		}
		departments = append(departments, &dept)
	}
	return departments, nil
}
