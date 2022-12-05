package part2

import (
	"strings"
	"testing"
)

func TestMultipleMove(t *testing.T) {
	before := [][]string{
		[]string{"A", "B", "C"},
		[]string{"X"},
		[]string{"Y", "Z"},
	}
	count := 2
	from := 1
	to := 2
	want := [][]string{
		[]string{"A"},
		[]string{"X", "B", "C"},
		[]string{"Y", "Z"},
	}

	MultipleMove(&before, count, from, to)

	for i := 0; i < len(want); i++ {
		if strings.Join(before[i], ", ") != strings.Join(want[i], ", ") {
			t.Error("want:", want[i], "got:", before[i])
		}
	}
}

func TestGetAnswer(t *testing.T) {
	want := "MCD"
	got, _ := GetAnswer("../example.txt")
	if got != want {
		t.Error("want: ", want, "got: ", got)
	}
}
