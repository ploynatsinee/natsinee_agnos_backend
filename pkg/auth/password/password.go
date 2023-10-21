package password

import (
	"strconv"
	"unicode"
)

type PasswordRecommendation struct{}

func NewPasswordRecommendation() *PasswordRecommendation {
	return &PasswordRecommendation{}
}

func (pr *PasswordRecommendation) RecommendStrongPassword(initPassword string) string {
	numSteps := pr.PasswordLength(initPassword)

	if !pr.ContainsLowerUpperDigit(initPassword) {
		if !unicode.IsLower(rune(initPassword[0])) || !unicode.IsUpper(rune(initPassword[0])) {
			return "password must contain at least one lowercase and uppercase letter"
		}
		if !unicode.IsDigit(rune(initPassword[0])) {
			return "password must contain at least one digit"
		}
	}

	if !pr.NotContain3RepeatingCharacters(initPassword) {
		return "password must not contain 3 repeating characters in a row"
	}

	return strconv.Itoa(numSteps)
}

func (pr *PasswordRecommendation) PasswordLength(initPassword string) int {
	length := len(initPassword)

	if length >= 6 && length < 20 {
		return 0
	} else if length < 6 {
		return 6 - length
	} else {
		return -1
	}
}

func (pr *PasswordRecommendation) ContainsLowerUpperDigit(initPassword string) bool {
	hasLower := false
	hasUpper := false
	hasDigit := false

	for _, char := range initPassword {
		if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}

		if hasLower && hasUpper && hasDigit {
			break
		}
	}
	return hasLower && hasUpper && hasDigit
}

func (pr *PasswordRecommendation) NotContain3RepeatingCharacters(initPassword string) bool {
	for i := 0; i < len(initPassword)-2; i++ {
		if initPassword[i] == initPassword[i+1] && initPassword[i] == initPassword[i+2] {
			return false
		}
	}
	return true
}
