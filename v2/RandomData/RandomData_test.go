package toolkit

import (
	"strings"
	"testing"
)

func TestRandom_String(t *testing.T) {
	s := RandomString(10)
	if len(s) != 10 {
		t.Error("wrong length random string returned")
	}
}

func TestRandom_Email(t *testing.T) {
	s := RandomEmail()
	if len(s) != 10 && !strings.Contains(s, "@example.com") {
		t.Error("wrong length random email returned")
	}
}
