package part1

import "testing"

func TestPart1(t *testing.T) {
	type test struct {
        input       rune
        want        int
    }

    tests := []test{
        {input: 'a', want: 1},
        {input: 'z', want: 26},
        {input: 'A', want: 27},
        {input: 'Z', want: 52},
    }

    for _, tc := range tests {
        got := GetPoint(tc.input)
        if got != tc.want {
            t.Error("input:", tc.input, ", want:", tc.want, ", got:", got)
        }
    }
}

func TestGetAnswer(t *testing.T) {
	want := 157

	got, _ := GetAnswer("../example.txt")

	if got != want {
		t.Error("want: ", want, " got: ", got)
	}
}