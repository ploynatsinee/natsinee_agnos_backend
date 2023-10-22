# Backend Strong Password Recommendation steps

## Assignment

1. Docker Setup: Set up a server with Docker Compose containing Nginx, a GoLang service, and a PostgreSQL database.
2. Build Backend: Implement the backend server with the specified API.
3. Database Logging: Store logs of requests and responses in a PostgreSQL database.
4. Unit Tests: Write unit tests for your application.

## Running the Application
docker compose up --build

## Running the tests
go test -v ./tests/unit_test/pkg/auth/password_test.go

## API Endpoints
`GET`
127.0.0.1:8080/api/strong_password_steps

## Project structure
```bash
.
├── Dockerfile
├── api
│   ├── handlers
│   │   └── password.go
│   └── middleware
│       └── logger.go
├── cmd
│   └── main.go
├── db
│   └── db.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── nginx
│   └── nginx.conf
├── pkg
│   └── auth
│       └── password
│           └── password.go
└── tests
    └── unit_test
        └── pkg
            └── auth
                └── password_test.go
```

## Database records
```bash
db=# SELECT * FROM logs;
 id | method |            url             | status | response_time 
----+--------+----------------------------+--------+---------------
  2 | GET    | /api/strong_password_steps |    422 |             2
  3 | GET    | /api/strong_password_steps |    200 |             1
```
