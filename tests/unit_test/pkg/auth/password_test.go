package tests

import (
	"testing"

	"password_recommendation/pkg/auth/password"
)

func TestLowerUpperDigit(t *testing.T) {
	input := "loskew[xs2!"
	want := "password must contain at least one lowercase and uppercase letter"
	r := password.RecommendStrongPassword(input)
	if r != want {
		t.Errorf("Expected %s, but got %s", want, r)
	}
}

func TestNotContain3RepeatingCharacters(t *testing.T) {
	input := "12poskeP[xs.pppp"
	want := "password must not contain 3 repeating characters in a row"
	r := password.RecommendStrongPassword(input)
	if r != want {
		t.Errorf("Expected %s, but got %s", want, r)
	}
}

func TestContainsDotOrExclamationMark(t *testing.T) {
	input := "12poskeP[xs"
	want := "password must contain at least one dot or exclamation mark"
	r := password.RecommendStrongPassword(input)
	if r != want {
		t.Errorf("Expected %s, but got %s", want, r)
	}
}

func TestPasswordLength(t *testing.T) {
	input := "1P!l"
	want := "2"
	r := password.RecommendStrongPassword(input)
	if r != want {
		t.Errorf("Expected %s, but got %s", want, r)
	}
}
