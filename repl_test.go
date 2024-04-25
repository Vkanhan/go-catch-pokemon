package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string 
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input : "Hello World",
			expected: []string{"hello","world",},
		},
		{
			input : "HELLO World",
			expected: []string{"hello","world",},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v not equal %v", expectedWord, actualWord)
			}
		}
	}
}

