package helper

import (
	"fmt"
	"testing"
)

func TestGetAnswerPart1(t *testing.T) {
	want := 21
	got := GetAnswerPart1("../example.txt")
	if got != want {
		t.Error(fmt.Sprintf("want: %d, got: %d", want, got))
	}
}

func TestGetAnswerPart2(t *testing.T) {
	want := 8
	got := GetAnswerPart2("../example.txt")
	if got != want {
		t.Error(fmt.Sprintf("want: %d, got: %d", want, got))
	}
}
