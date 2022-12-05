package part1

import (
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	type test struct {
		input string
		want  []string
	}
	tests := []test{
		{input: "[A]", want: []string{"A"}},
		{input: "[A] [B]", want: []string{"A", "B"}},
		{input: "[A]     [B]", want: []string{"A", "", "B"}},
	}
	for _, tc := range tests {
		got := ReadInput(tc.input)
		if strings.Join(got, ", ") != strings.Join(tc.want, ", ") {
			t.Error("want:", tc.want, "got:", got)
		}
	}
}

func TestPrepareInitializeData(t *testing.T) {
	input := [][]string{
		[]string{"", "C"},
		[]string{"B", "Z"},
		[]string{"A", "X", "Y"},
	}

	want := [][]string{
		[]string{"A", "B"},
		[]string{"X", "Z", "C"},
		[]string{"Y"},
	}

	got := PrepareInitializeData(input)

	for i := 0; i < len(want); i++ {
		if strings.Join(got[i], ", ") != strings.Join(want[i], ", ") {
			t.Error("want:", want[i], "got:", got[i])
		}
	}
}

func TestMove(t *testing.T) {
	before := [][]string{
		[]string{"A", "B", "C"},
		[]string{"X"},
		[]string{"Y", "Z"},
	}
	from := 1
	to := 2
	want := [][]string{
		[]string{"A", "B"},
		[]string{"X", "C"},
		[]string{"Y", "Z"},
	}

	Move(&before, from, to)

	for i := 0; i < len(want); i++ {
		if strings.Join(before[i], ", ") != strings.Join(want[i], ", ") {
			t.Error("want:", want[i], "got:", before[i])
		}
	}
}

func TestGetTops(t *testing.T) {
	input := [][]string{
		[]string{"A", "B"},
		[]string{"X"},
		[]string{"Y", "Z"},
	}
	want := "BXZ"

	got := GetTops(&input)

	if got != want {
		t.Error("want:", want, "got:", got)
	}
}

func TestGetCommand(t *testing.T) {
	input := "move 20 from 3 to 4"

	wantedCount := 20
	wantedFrom := 3
	wantedTo := 4

	count, from, to := GetCommand(input)
	if count != wantedCount || from != wantedFrom || to != wantedTo {
		t.Error("wantedCount:", wantedCount, "count:", count)
		t.Error("wantedFrom:", wantedFrom, "from:", from)
		t.Error("wantedTo:", wantedTo, "to:", to)
	}
}

func TestGetAnswer(t *testing.T) {
	want := "CMZ"
	got, _ := GetAnswer("../example.txt")
	if got != want {
		t.Error("want: ", want, "got: ", got)
	}
}
