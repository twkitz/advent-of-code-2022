package helper

import (
	"fmt"
	"testing"
)

func TestGetAnswerPart1(t *testing.T) {
	want := 13
	got := GetAnswerPart1("../example.txt")
	if got != want {
		t.Error(fmt.Sprintf("want: %d, got: %d", want, got))
	}
}

func TestGetAnswerPart2(t *testing.T) {
	type test struct {
		filePath string
		want     int
	}

	tests := []test{
		{"../example.txt", 1},
		{"../example2.txt", 36},
	}

	for _, tc := range tests {
		got := GetAnswerPart2(tc.filePath)
		if got != tc.want {
			t.Error(fmt.Sprintf("want: %d, got: %d", tc.want, got))
		}
	}
}
