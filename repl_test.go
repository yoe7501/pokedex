package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "charmander vulpix",
			expected: []string{"charmander", "vulpix"},
		},
	}

	for _, c := range cases {
		result := cleanInput(c.input)

		if len(result) != len(c.expected) {
			t.Errorf("length of result and expected not the same. Result %v (len %d) expected %v (len %d)",
				result, len(result), c.expected, len(c.expected))
			continue // Skip the word loop if lengths don't match to avoid index out of range panics
		}

		for i := range result {
			word := result[i]
			expected := c.expected[i]

			if word != expected {
				t.Errorf("words do not match. Word %v expected %v", word, expected)
			}
		}
	}
}
