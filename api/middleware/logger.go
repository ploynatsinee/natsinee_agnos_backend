package middleware

import (
	"log"
	"time"

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
		requestStatus := c.Writer.Status()
		responseTime := endTime.Sub(startTime).Milliseconds()

		_, err := db.Exec("INSERT INTO logs (method, url, status, response_time) VALUES (?,?,?,?)", requestMethod, requestURL, requestStatus, responseTime)

		if err != nil {
			log.Println("Error inserting log into the database:", err)
		}
	}
}
