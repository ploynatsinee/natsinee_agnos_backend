package main

import (
	"password_recommendation/api/handlers"
	"password_recommendation/api/middleware"
	"password_recommendation/db"

	"github.com/go-pg/pg/v9"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := db.Connect()
	defer db.Close()
	r := setupRouter(db)
	if err := r.Run(":4000"); err != nil {
		panic(err)
	}
}

func setupRouter(db *pg.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware(db))
	r.GET("/api/strong_password_steps", handlers.PasswordHandler(db))

	return r
}
