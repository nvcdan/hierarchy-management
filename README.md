
# hierarchy-management

`hierarchy-management` is a REST API application built in Go using the Gin framework. It manages a hierarchical structure of departments, supporting features like creation, updating, deletion, and hierarchical visualization.

## Features

- Department creation and updating with hierarchical parent-child relationships
- Logical deletion of departments
- Retrieval of department hierarchy in structured format
- Token-based authentication for secure access

## Technologies

- **Go** - main programming language
- **Gin** - REST API framework
- **MySQL** - database for department hierarchy
- **Docker** - containerization for isolated environments

## Project Structure

- `internal/domain` - Core domain interfaces and data structures
- `internal/repository` - Database interactions and data access logic
- `internal/service` - Business logic and validation
- `internal/handler` - HTTP request handling
- `cmd/` - Application entry point

## Requirements

- **Go** v1.19 or higher
- **MySQL** v8.0 or higher
- **Docker** (optional for containerized deployment)

## Configuration

Create a `.env` file in the project root with the following settings:

```plaintext
DB_USER=username
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=hierarchy_management
JWT_SECRET=your_jwt_secret
```

## Installation and Running

1. Clone the repository:

    ```bash
    git clone https://github.com/username/hierarchy-management.git
    cd hierarchy-management
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. Set up the database and tables using the SQL scripts in `scripts/`.

4. Run the application:

    ```bash
    go run cmd/main.go
    ```

5. Alternatively, use Docker:

    ```bash
    docker-compose up --build
    ```

## API Endpoints

### Authentication

All endpoints require JWT tokens in the `Authorization` header:

```
Authorization: Bearer <token>
```

### Departments

- **Create Department**  
    `POST /api/departments/create`  
    Request Body:
    ```json
    {
      "name": "Department",
      "parent_id": 1,
      "flags": 1
    }
    ```
    
- **Update Department**  
    `PUT /api/departments/{id}/update`  
    Request Body:
    ```json
    {
      "name": "Updated Department",
      "parent_id": 1,
      "flags": 2
    }
    ```

- **Delete Department**  
    `DELETE /api/departments/{id}/delete`

- **Get Hierarchy**  
    - By name: `GET /api/departments/hierarchy?name={departmentName}`
    - Full hierarchy: `GET /api/departments/hierarchy/all`

## License

Licensed under the MIT License. See `LICENSE` for details.
