package main

import ("testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello","world"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "Bulbasaur SQUiRtle CHARMANDER EeVeE",
			expected: []string{"bulbasaur","squirtle","charmander","eevee"},
		},
		{
			input: "GeoDude    MAGneMITE    Beedrill",
			expected: []string{"geodude","magnemite","beedrill"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Slice size doesn't match:\r\nexpected: %d\r\nactual: %d",len(c.expected),len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Slice content doesn't match at index %d:\r\nexpected: %s\r\nactual: %s",i, expectedWord, word)
			}
		}
	}
}