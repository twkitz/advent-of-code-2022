package helper

import (
	"testing"

	"github.com/twkitz/advent-of-code-2022/utils"
)

func InitializeMonkeyPredictor() *MonkeyPredictor {
	mp := MonkeyPredictor{
		[]*Monkey{
			{[]int{79, 98}, "*", "19", 23, 2, 3, 0},
			{[]int{54, 65, 75, 74}, "+", "6", 19, 2, 0, 0},
			{[]int{79, 60, 97}, "*", "old", 13, 1, 3, 0},
			{[]int{74}, "+", "3", 17, 0, 1, 0},
		},
	}
	return &mp
}

func TestGetReport(t *testing.T) {
	mp := InitializeMonkeyPredictor()
	want := "Monkey 0: 79, 98\nMonkey 1: 54, 65, 75, 74\nMonkey 2: 79, 60, 97\nMonkey 3: 74"

	got := mp.GetReport()

	if got != want {
		t.Errorf("\n>>> want <<<\n%s\n\n>>> got <<<\n%s", want, got)
	}
}

func TestRun(t *testing.T) {
	type test struct {
		round int
		want  string
	}

	tests := []test{
		{1, "Monkey 0: 20, 23, 27, 26\nMonkey 1: 2080, 25, 167, 207, 401, 1046\nMonkey 2: \nMonkey 3: "},
		{2, "Monkey 0: 695, 10, 71, 135, 350\nMonkey 1: 43, 49, 58, 55, 362\nMonkey 2: \nMonkey 3: "},
		{10, "Monkey 0: 91, 16, 20, 98\nMonkey 1: 481, 245, 22, 26, 1092, 30\nMonkey 2: \nMonkey 3: "},
		{20, "Monkey 0: 10, 12, 14, 26, 34\nMonkey 1: 245, 93, 53, 199, 115\nMonkey 2: \nMonkey 3: "},
	}

	for _, tc := range tests {
		mp := InitializeMonkeyPredictor()

		for i := 0; i < tc.round; i++ {
			mp.Run()
		}
		got := mp.GetReport()

		if got != tc.want {
			t.Errorf("round: %d\n>>> want <<<\n%s\n\n>>> got <<<\n%s", tc.round, tc.want, got)
		}
	}
}

func TestGetInspectedReport(t *testing.T) {
	mp := InitializeMonkeyPredictor()
	want := utils.SliceToString([]int{101, 95, 7, 105})

	for i := 0; i < 20; i++ {
		mp.Run()
	}
	got := utils.SliceToString(mp.GetInspectedCounts())

	if got != want {
		t.Errorf("\n>>> want <<<\n%s\n\n>>> got <<<\n%s", want, got)
	}
}
