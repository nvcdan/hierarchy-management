USE hierarchy_management;

DELIMITER $$

CREATE PROCEDURE CreateDepartment (
    IN dept_name VARCHAR(255),
    IN dept_parent_id INT,
    IN dept_flags TINYINT
)
BEGIN
    INSERT INTO departments (name, parent_id, flags)
    VALUES (dept_name, dept_parent_id, dept_flags);
END $$

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

CREATE PROCEDURE DeleteDepartment (
    IN dept_id INT
)
BEGIN
    UPDATE departments
    SET flags = flags | 2
    WHERE id = dept_id;
END $$

CREATE PROCEDURE GetDepartmentHierarchy (
    IN dept_name VARCHAR(255)
)
BEGIN
    DECLARE dept_id INT;
    SELECT id INTO dept_id FROM departments WHERE name = dept_name;

    WITH RECURSIVE dept_tree AS (
        SELECT id, name, parent_id, flags FROM departments WHERE id = dept_id
        UNION ALL
        SELECT d.id, d.name, d.parent_id, d.flags
        FROM departments d
        INNER JOIN dept_tree dt ON dt.id = d.parent_id
    )
    SELECT * FROM dept_tree;
END $$

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

DELIMITER ;