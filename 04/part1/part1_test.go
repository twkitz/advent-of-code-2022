package part1

import "testing"

func TestIsFullyContain(t *testing.T) {
	type test struct {
		from1 int
		to1   int
		from2 int
		to2   int
		want  bool
	}

	tests := []test{
		{
			from1: 2,
			to1:   4,
			from2: 6,
			to2:   8,
			want:  false,
		},
		{
			from1: 2,
			to1:   3,
			from2: 4,
			to2:   5,
			want:  false,
		},
		{
			from1: 2,
			to1:   8,
			from2: 3,
			to2:   7,
			want:  true,
		},
		{
			from1: 6,
			to1:   6,
			from2: 4,
			to2:   6,
			want:  true,
		},
	}

	for _, tc := range tests {
		got := IsFullyContain(tc.from1, tc.to1, tc.from2, tc.to2)
		if got != tc.want {
			t.Error("want: ", tc.want, "got:", got)
		}
	}
}

func TestGetAnswer(t *testing.T) {
	want := 2
	got, _ := GetAnswer("../example.txt")
	if got != want {
		t.Error("want: ", want, "got: ", got)
	}
}
