package day1

import "testing"

var tests = []struct {
	in  string
	out string
}{
	{"-1\n-1\n-1", "-3"},
	{"-1\n+1\n-1", "-1"},
	{"+1\n+1\n-2", "0"},
	{"-1\n-2\n-3", "-6"},
}

func TestFrequency(t *testing.T) {
	solution := Solution{}
	for _, tt := range tests {
		d := solution.Solve(tt.in)
		if tt.out != d {
			t.Errorf("CountWithNextDigit(%q) => %s, want %s", tt.in, d, tt.out)
		}
	}
}
