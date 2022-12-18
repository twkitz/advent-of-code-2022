package helper

import (
	"testing"
)

func TestGetAnswerPart1(t *testing.T) {
	want := 64
	got := GetAnswerPart1("../example.txt")
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}