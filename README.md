# Backend Strong Password Recommendation steps

## Assignment

1. Docker Setup: Set up a server with Docker Compose containing Nginx, a GoLang service, and a PostgreSQL database.
2. Build Backend: Implement the backend server with the specified API.
3. Database Logging: Store logs of requests and responses in a PostgreSQL database.
4. Unit Tests: Write unit tests for application.

## Running the Application
`docker compose up --build`

## Running the tests
`go test -v ./tests/unit_test/pkg/auth/password_test.go`

## API Endpoints
`GET`
`127.0.0.1:8080/api/strong_password_steps`

### example request.json
    {
    "init_password": "loskew[xs2!P"
    }

### example response
    {
    "num_of_steps": "0"
    }

## Project structure
```bash
.
├── Dockerfile
├── README.md
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
├── tests
│   └── unit_test
│       └── pkg
│           └── auth
│               └── password_test.go
└── utils
    └── hash.go
```

## Database records
```bash
db=# SELECT * FROM logs;
 log_id |           timestamp           | method |            url             |                         request_body                         | response_status | response_time 
--------+-------------------------------+--------+----------------------------+--------------------------------------------------------------+-----------------+---------------
      1 | 2023-10-22 14:43:12.763451+00 | GET    | /api/strong_password_steps | $2a$10$x3DQwFBScRzXiuUCh58qrOHzvo1YPFDFlgP2gb9tt4qS2WRr0j24W |             200 |             1
     34 | 2023-10-22 14:45:21.958982+00 | GET    | /api/strong_password_steps | $2a$10$FFgVAFwHm3v97a5Otg4lOu0eNB8UW2K4YeS8nkrN53SKoAiiYbZ2G |             422 |             0
```
