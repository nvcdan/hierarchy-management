USE hierarchy_management_test;

DELIMITER $$

DROP PROCEDURE IF EXISTS CreateDepartment $$
CREATE PROCEDURE CreateDepartment (
    IN dept_name VARCHAR(255),
    IN dept_parent_id INT,
    IN dept_flags TINYINT
)
BEGIN
    INSERT INTO departments (name, parent_id, flags)
    VALUES (dept_name, dept_parent_id, dept_flags);
END $$

DROP PROCEDURE IF EXISTS UpdateDepartment $$
CREATE PROCEDURE UpdateDepartment (
    IN dept_id INT,
    IN dept_name VARCHAR(255),
    IN dept_parent_id INT,
    IN dept_flags TINYINT
)
BEGIN
    UPDATE departments
    SET name = dept_name,
        parent_id = dept_parent_id,
        flags = dept_flags
    WHERE id = dept_id;
END $$

DROP PROCEDURE IF EXISTS DeleteDepartment $$
CREATE PROCEDURE DeleteDepartment (
    IN dept_id INT
)
BEGIN
    UPDATE departments
    SET flags = flags | 2
    WHERE id = dept_id;
END $$

DROP PROCEDURE IF EXISTS GetDepartmentHierarchy $$
CREATE PROCEDURE GetDepartmentHierarchy (
    IN dept_name VARCHAR(255)
)
BEGIN
    WITH RECURSIVE ancestors AS (
        SELECT id, name, parent_id, flags
        FROM departments
        WHERE name LIKE CONCAT('%', dept_name, '%')
        UNION ALL
        SELECT d.id, d.name, d.parent_id, d.flags
        FROM departments d
        INNER JOIN ancestors a ON a.parent_id = d.id
    ),
    descendants AS (
        SELECT id, name, parent_id, flags
        FROM departments
        WHERE name LIKE CONCAT('%', dept_name, '%')
        UNION ALL
        SELECT d.id, d.name, d.parent_id, d.flags
        FROM departments d
        INNER JOIN descendants ds ON d.parent_id = ds.id
    ),
    dept_tree AS (
        SELECT * FROM ancestors
        UNION
        SELECT * FROM descendants
    )
    SELECT DISTINCT * FROM dept_tree;
END $$


DROP PROCEDURE IF EXISTS GetAllDepartmentsHierarchy $$
CREATE PROCEDURE GetAllDepartmentsHierarchy()
BEGIN
    WITH RECURSIVE dept_tree AS (
        SELECT id, name, parent_id, flags
        FROM departments
        WHERE parent_id IS NULL
        UNION ALL
        SELECT d.id, d.name, d.parent_id, d.flags
        FROM departments d
        INNER JOIN dept_tree dt ON dt.id = d.parent_id
    )
    SELECT * FROM dept_tree;
END $$

DROP PROCEDURE IF EXISTS ExistsByID $$
CREATE PROCEDURE ExistsByID (
    IN dept_id INT
)
BEGIN
    SELECT EXISTS(SELECT 1 FROM departments WHERE id = dept_id) AS id_exists;
END $$

DROP PROCEDURE IF EXISTS IsDuplicateDepartmentName $$
CREATE PROCEDURE IsDuplicateDepartmentName (
    IN dept_name VARCHAR(255)
)
BEGIN
    SELECT EXISTS(SELECT 1 FROM departments WHERE name = dept_name) AS is_duplicate;
END $$


DELIMITER ;