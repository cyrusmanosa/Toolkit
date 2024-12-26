package toolkit

import (
	"errors"
	"regexp"
	"strings"
)

// Slugify is a very simple means of creating a slug from a string
// example: 「 Hello World! This is a test. 」 => 「 hello-world-this-is-a-test 」

func Slugify(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string not permitted")
	}

	var re = regexp.MustCompile(`[^a-z\d]+`)
	slug := strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")
	if len(slug) == 0 {
		return "", errors.New("after removing characters, slug is zero length")
	}

	return slug, nil
}
