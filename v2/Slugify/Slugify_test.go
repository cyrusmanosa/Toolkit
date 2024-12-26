package toolkit

import "testing"

var slugTests = []struct {
	name          string
	s             string
	expected      string
	errorExpected bool
}{
	{
		name:          "valid string",
		s:             "now is the time",
		expected:      "now-is-the-time",
		errorExpected: false,
	},
	{
		name:          "empty string",
		s:             "",
		expected:      "",
		errorExpected: true,
	},
	{
		name:          "complex string",
		s:             "Now is the time for all GOOD men! + fish & such &^123",
		expected:      "now-is-the-time-for-all-good-men-fish-such-123",
		errorExpected: false,
	},
	{
		name:          "japanese string",
		s:             "こんにちは、世界",
		expected:      "",
		errorExpected: true,
	},
	{
		name:          "japanese string and roman characters",
		s:             "hello world こんにちは、世界",
		expected:      "hello-world",
		errorExpected: false,
	},
}

func TestTools_Slugify(t *testing.T) {
	for _, e := range slugTests {
		slug, err := Slugify(e.s)
		if err != nil && !e.errorExpected {
			t.Errorf("%s: error received when none expected: %s", e.name, err.Error())
		}

		if !e.errorExpected && slug != e.expected {
			t.Errorf("%s: wrong slug returned; expected %s, but got %s", e.name, e.expected, slug)
		}
	}
}
