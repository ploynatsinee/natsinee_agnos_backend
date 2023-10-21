package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"password_recommendation/pkg/auth/password"

	"github.com/gin-gonic/gin"
)

func PasswordHandler(c *gin.Context) {
	pr := password.NewPasswordRecommendation()
	r := pr.RecommendStrongPassword("12Poskew[xs")
	c.String(http.StatusOK, "num_of_steps %s", r)
}

func TestPasswordHandler(t *testing.T) {
	r := gin.Default()
	r.GET("/api/strong_password_steps", PasswordHandler)

	req, _ := http.NewRequest("GET", "/api/strong_password_steps", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	response := w.Body.String()

	if !strings.HasPrefix(response, "num_of_steps ") {
		t.Errorf("Expected response body to start with 'num_of_steps', but got %s", response)
	}
	expectedValue := "0"
	extractedValue := strings.TrimPrefix(response, "num_of_steps ")

	if extractedValue != expectedValue {
		t.Errorf("Expected response body to be %s, but got %s", expectedValue, extractedValue)
	}
}
