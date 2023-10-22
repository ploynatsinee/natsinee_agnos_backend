package middleware

import (
	"io"
	"log"
	"time"

	"password_recommendation/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
)

func LoggerMiddleware(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()

		requestMethod := c.Request.Method
		requestURL := c.Request.URL.String()
		responseStatus := c.Writer.Status()
		responseTime := endTime.Sub(startTime).Milliseconds()
		requestJSON := c.Request.Body
		timestamp := time.Now()

		requestBody, err := io.ReadAll(requestJSON)
		if err != nil {
			log.Println("Error reading request body:", err)
			return
		}
		hashedPassword, err := utils.HashPassword(string(requestBody))
		if err != nil {
			log.Println("Error hashing password:", err)
			return
		}

		_, err = db.Exec(`INSERT INTO logs (method, url, request_body, response_status, timestamp, response_time)
                    VALUES (?, ?, ?, ?, ?, ?)`,
			requestMethod, requestURL, hashedPassword, responseStatus, timestamp, responseTime)

		if err != nil {
			log.Println("Error inserting log into the database:", err)
		}
	}
}
