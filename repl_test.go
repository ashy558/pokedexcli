package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	type test struct {
		name string
		input string
		expected []string
	}
	cases := []test {
		{
			name: "double spaces",
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name: "all titles",
			input: "Charmander Bulbasaur Pikachu",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			name: "all capitals",
			input: "CHARMELEON IVYSAUR RAICHU",
			expected: []string{"charmeleon", "ivysaur", "raichu"},
		},
		{
			name: "empty string",
			input: "",
			expected: []string{},
		},
		{
			name: "single space",
			input: " ",
			expected: []string{},
		},
		{
			name: "mixed casing",
			input: "charizard VENUSAUR Pichu",
			expected: []string{"charizard", "venusaur", "pichu"},
		},
		{
			name: "spongebob casing",
			input: "sQuIrTlE",
			expected: []string{"squirtle"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLength := len(actual)
		expectedLength := len(c.expected)
		if actualLength != expectedLength {
			t.Errorf("%s: expected: %d, actual: %d", c.name, expectedLength, actualLength)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s: expected: %s, actual: %s", c.name, expectedWord, word)
			}
		}
	}
}
