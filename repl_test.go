package main

import "testing"

func testClearInput(t *testing.T) {
	testcases := []struct {
		input    string
		expected []string
	}{{
		input:    "",
		expected: []string{},
	}, {input: "brian",
		expected: []string{"brian"},
	}, {
		input:    "brian ruhiu",
		expected: []string{"brian", "ruhiu"},
	},
	}
	for _, cases := range testcases {
		test := cleanInput(cases.input)
		if len(test) != len(cases.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", test, cases.expected)
			continue
		}
		for i := range test {
			word := test[i]
			expectedWord := cases.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", cases.input, test, cases.expected)
			}
		}
	}
}
