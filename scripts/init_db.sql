CREATE DATABASE IF NOT EXISTS hierarchy_management;

USE hierarchy_management;

CREATE TABLE departments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id INT DEFAULT NULL,
    flags TINYINT(1) NOT NULL DEFAULT 0,
    FOREIGN KEY (parent_id) REFERENCES departments(id)
);