package main

import (
	"testing"
)

type testCase struct {
	input		string
	expected	[]string
}

func TestCleanInput(t *testing.T) {
	cases := []testCase {
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "",
			expected: []string{""},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		
		if len(actual) != len(c.expected) {
			t.Fatalf("the lengths of cases and cleaned inputs don't match")
		}

		for i, _ := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Fatalf("expected: %s and got: %s", word, expectedWord)
			}
		}
	}
}
