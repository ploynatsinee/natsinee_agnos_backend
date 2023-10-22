package password

import (
	"strconv"
	"unicode"
)

func RecommendStrongPassword(initPassword string) string {
	numSteps := passwordLength(initPassword)

	if !containsLowerUpperDigit(initPassword) {
		if !unicode.IsLower(rune(initPassword[0])) || !unicode.IsUpper(rune(initPassword[0])) {
			return "password must contain at least one lowercase and uppercase letter"
		}
		if !unicode.IsDigit(rune(initPassword[0])) {
			return "password must contain at least one digit"
		}
	}

	if !notContain3RepeatingCharacters(initPassword) {
		return "password must not contain 3 repeating characters in a row"
	}

	if !containsDotOrExclamationMark(initPassword) {
		return "password must contain at least one dot or exclamation mark"
	}

	return strconv.Itoa(numSteps)
}

func passwordLength(initPassword string) int {
	length := len(initPassword)

	if length >= 6 && length < 20 {
		return 0
	} else if length < 6 {
		return 6 - length
	} else {
		return -1
	}
}

func containsLowerUpperDigit(initPassword string) bool {
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

func notContain3RepeatingCharacters(initPassword string) bool {
	for i := 0; i < len(initPassword)-2; i++ {
		if initPassword[i] == initPassword[i+1] && initPassword[i] == initPassword[i+2] {
			return false
		}
	}
	return true
}

func containsDotOrExclamationMark(initPassword string) bool {
	for _, char := range initPassword {
		if char == '.' || char == '!' {
			return true
		}
	}
	return false
}
