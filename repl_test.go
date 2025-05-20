package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "well hello there ",
			expected: "well",
		},
		{
			input:    "Hello there",
			expected: "hello",
		},
		{
			input:    "POKEMON was underrated",
			expected: "pokemon",
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		expectedWord := c.expected
		if actual != expectedWord {
			t.Errorf("Expected %s, but got %s", expectedWord, actual)
		}
	}
}
