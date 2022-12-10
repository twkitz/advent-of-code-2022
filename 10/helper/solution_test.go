package helper

import (
	"fmt"
	"strings"
	"testing"

	"github.com/twkitz/advent-of-code-2022/utils"
)

func TestGetAnswerPart1(t *testing.T) {
	type test struct {
		interestedCycles []int
		want             int
	}

	tests := []test{
		{[]int{20}, 420},
		{[]int{60}, 1140},
		{[]int{100}, 1800},
		{[]int{140}, 2940},
		{[]int{180}, 2880},
		{[]int{220}, 3960},
		{[]int{20, 60}, 420 + 1140},
		{[]int{20, 60, 100, 140, 180, 220}, 13140},
	}

	for _, tc := range tests {
		got := GetAnswerPart1("../example.txt", tc.interestedCycles)
		if got != tc.want {
			t.Error(fmt.Sprintf("want: %d, got: %d", tc.want, got))
		}
	}
}

func TestGetAnswerPart2(t *testing.T) {
	wantedOutput := utils.GetFileContent("../part2-example-output.txt")
	want := strings.Join(wantedOutput, "\n")
	got := strings.Trim(GetAnswerPart2("../example.txt"), "\n")
	if got != want {
		t.Error(fmt.Sprintf("\n====want====\n\n%s\n\n====got====\n\n%s\n\n===========", want, got))
	}
}
