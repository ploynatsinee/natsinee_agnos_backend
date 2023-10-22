package handlers

import (
	"net/http"
	"password_recommendation/pkg/auth/password"
	"strings"

	"github.com/go-pg/pg/v9"

	"github.com/gin-gonic/gin"
)

type PasswordRequest struct {
	InitPassword string `json:"init_password"`
}

func PasswordHandler(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request PasswordRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		initPassword := request.InitPassword
		r := password.RecommendStrongPassword(initPassword)

		if strings.HasPrefix(r, "password must") {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": r})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"num_of_steps": r,
		})
	}
}
