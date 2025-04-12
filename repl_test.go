package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello  world  ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "    Hello  World DOH ",
			expected: []string{"hello", "world", "doh"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length doesnt match")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %s\nactual: %s", expectedWord, word)
			}

		}
	}

}
