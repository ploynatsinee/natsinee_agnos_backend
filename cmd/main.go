package main

import (
	"log"
	"password_recommendation/api/handlers"
	"password_recommendation/db"

	"github.com/go-pg/pg/v9"

	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := db.Connect()
	defer db.Close()
	r := setupRouter(db)
	r.Use(LoggerMiddleware(db))
	r.Run()
}

func setupRouter(db *pg.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/api/strong_password_steps", handlers.PasswordHandler(db))

	if err := r.Run(":4000"); err != nil {
		panic(err)
	}

	return r
}

func LoggerMiddleware(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()
		endTime := time.Now()

		requestMethod := c.Request.Method
		requestURL := c.Request.URL.String()
		requestStatus := c.Writer.Status()
		responseTime := endTime.Sub(startTime).Milliseconds()

		// Insert the log into the PostgreSQL database
		_, err := db.Exec("INSERT INTO request_logs (method, url, status, response_time) VALUES ($1, $2, $3, $4)", requestMethod, requestURL, requestStatus, responseTime)

		if err != nil {
			log.Println("Error inserting log into the database:", err)
		}
	}
}
