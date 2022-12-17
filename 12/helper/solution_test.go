package helper

import "testing"

func TestGetAnswerPart1(t *testing.T) {
	want := 31
	got := GetAnswer("../example.txt", false)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestGetAnswerPart2(t *testing.T) {
	want := 29
	got := GetAnswer("../example.txt", true)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
