package toolkit

import (
	"math/rand"
)

const randomString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// random characters
func RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomString)
	for i := range s {
		s[i] = r[rand.Intn(len(r))]
	}
	return string(s)
}

// random email
func RandomEmail() string {
	s, r := make([]rune, 25), []rune(randomString)
	for i := range s {
		s[i] = r[rand.Intn(len(r))]
	}
	return string(s) + "@example@example.com"
}
