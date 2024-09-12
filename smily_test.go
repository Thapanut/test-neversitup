package main

import "testing"

func TestCountSmileyFaces(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{":)", ";(", ";}", ":-D"}, 2},       // Two valid smiley faces
		{[]string{";D", ":-(", ":-)", ";~)"}, 3},     // Three valid smiley faces
		{[]string{";]", ":[", ";*", ":$", ";-D"}, 1}, // One valid smiley face
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := CountSmileyFaces(test.input)
			if result != test.expected {
				t.Errorf("For input %v, expected %d, but got %d", test.input, test.expected, result)
			}
		})
	}
}
