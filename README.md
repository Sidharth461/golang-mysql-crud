# Golang MySQL CRUD API

A simple REST API built with Go and MySQL that performs CRUD (Create, Read, Update, Delete) operations on user data.

## Features

* Create User
* Get Users
* Update User
* Delete User
* MySQL Database Integration
* Environment Variables using `.env`
* JSON Request & Response
* RESTful API

## Tech Stack

* Go (Golang)
* MySQL
* net/http
* database/sql
* godotenv
* MySQL Driver (`go-sql-driver/mysql`)

## Project Structure

```
.
├── main.go
├── go.mod
├── go.sum
├── .gitignore
├── .env (ignored)
└── README.md
```

## Database

Create a database named:

```
Users
```

Example table:

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    contact_number VARCHAR(15)
);
```

## Environment Variables

Create a `.env` file:

```
DB_USER=your_username
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=Users
```

## Run the Project

Install dependencies:

```bash
go mod tidy
```

Run:

```bash
go run main.go
```

## API Endpoints

| Method | Endpoint    | Description |
| ------ | ----------- | ----------- |
| POST   | /users      | Create User |
| GET    | /users      | Get Users   |
| PUT    | /users/1    | Update User |
| DELETE | /users/1    | Delete User |

## Author

Sidharth Kamble
