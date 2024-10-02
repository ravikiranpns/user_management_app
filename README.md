
# User Management Application

This is a simple user management application built using **Go**, **Gin** for routing, **PostgreSQL** for the relational database, and **MongoDB** for logging. It supports basic CRUD operations (Create, Read, Update, Delete) for user management and uses JWT-based admin authentication.

## Project Structure

```bash
.
├── config/                 # Configuration files
├── controllers/            # Contains handler functions for the routes (CRUD operations)
│   ├── userController.go   # User-related CRUD operations
│   └── userController_test.go # Tests for the userController
├── database/               # Database connection setup (PostgreSQL and MongoDB)
│   ├── connection.go       # PostgreSQL and MongoDB connection logic
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency lock file
├── main.go                 # Entry point of the application, route setup
├── middleware/             # Middleware for JWT authentication
│   └── authMiddleware.go   # JWT authentication logic
├── models/                 # Data models used in the application
│   └── user.go             # User model definition
└── routes/                 # Routes setup
    └── routes.go           # Maps URL paths to controller actions
```

### Explanation of Each Folder/File:

1. **`config/`**: 
   - Reserved for future configuration files like environment variables, database configuration, etc.

2. **`controllers/`**: 
   - Contains the core logic for user operations.
     - **`userController.go`**: Defines functions for creating, reading, updating, and deleting users, as well as the admin login logic.
     - **`userController_test.go`**: Unit tests for the user controller, verifying that the CRUD functionality works as expected.

3. **`database/`**: 
   - Contains the database connection logic for both PostgreSQL and MongoDB.
     - **`connection.go`**: Connects to both PostgreSQL (for user data) and MongoDB (for logging events).

4. **`main.go`**: 
   - This is the entry point of the application. It connects to both databases, initializes routes, and starts the server.

5. **`middleware/`**: 
   - Contains middleware functions, mainly for authentication.
     - **`authMiddleware.go`**: Handles JWT authentication, allowing only authorized users (admins) to access certain routes.

6. **`models/`**: 
   - Defines the structure of data models used in the application.
     - **`user.go`**: Defines the structure of a `User` in the system, including fields for ID, name, email, and password.

7. **`routes/`**: 
   - Sets up the routes and maps them to their corresponding controllers.
     - **`routes.go`**: Maps the endpoints (`/users`, `/admin/login`, etc.) to the respective handler functions in `userController.go`.

## Requirements

- **Go 1.19+**: The application is written in Go and requires Go installed on your system.
- **PostgreSQL**: Used to store user data.
- **MongoDB**: Used to log user-related events.
- **cURL**: To test the API endpoints.

## Installation

1. Clone the repository:

```bash
git clone git@github.com:ravikiranpns/user_management_app.git
cd user_management_app
```

2. Install the required Go modules:

```bash
go mod tidy
```

3. Ensure that both **PostgreSQL** and **MongoDB** are running on your system.

4. Update the PostgreSQL connection string in `database/connection.go` if necessary:

```go
dsn := "host=localhost user=user2 password=password2 dbname=user_management port=5432 sslmode=disable"
```

5. Start the MongoDB service if it’s not already running:

```bash
brew services start mongodb-community@6.0
```

## Usage

### 1. Running the Application

```bash
go run main.go
```

The server will start and listen on port `8080`.

### 2. API Endpoints

#### Admin Login

- **URL**: `/admin/login`
- **Method**: `POST`
- **Payload**:

  ```json
  {
    "email": "admin@example.com",
    "password": "password123"
  }
  ```

- **Response**:
  ```json
  {
    "message": "Login successful!"
  }
  ```

#### Create a User

- **URL**: `/users`
- **Method**: `POST`
- **Payload**:

  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }
  ```

#### Fetch All Users

- **URL**: `/users`
- **Method**: `GET`

#### Update a User

- **URL**: `/users/:id`
- **Method**: `PUT`
- **Payload**:

  ```json
  {
    "name": "John Updated",
    "email": "johnupdated@example.com"
  }
  ```

#### Delete a User

- **URL**: `/users/:id`
- **Method**: `DELETE`

#### MongoDB Logging

MongoDB is used to log events like user creation, fetching, updating, and deletion. The logs can be found in the MongoDB database under the `user_logs` collection.

## Example cURL Commands

1. **Admin Login**:

   ```bash
   curl -X POST http://localhost:8080/admin/login    -H "Content-Type: application/json"    -d '{"email": "admin@example.com", "password": "password123"}'
   ```

2. **Create User**:

   ```bash
   curl -X POST http://localhost:8080/users    -H "Content-Type: application/json"    -d '{"name": "John Doe", "email": "john@example.com", "password": "password123"}'
   ```

3. **Get All Users**:

   ```bash
   curl http://localhost:8080/users
   ```

4. **Update a User**:

   ```bash
   curl -X PUT http://localhost:8080/users/1    -H "Content-Type: application/json"    -d '{"name": "Shilpa Updated", "email": "shilpa_updated@example.com"}'
   ```

5. **Delete a User**:

   ```bash
   curl -X DELETE http://localhost:8080/users/1
   ```

## Conclusion

This project is a basic user management system that supports admin authentication and CRUD operations for users. The application integrates PostgreSQL for storing user data and MongoDB for logging purposes.

Feel free to contribute or report any issues!
