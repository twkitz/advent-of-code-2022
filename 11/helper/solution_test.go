package helper

import "testing"

func TestGetAnswerPart1(t *testing.T) {
	want := 10605
	got := GetAnswerPart1("../example.txt")
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestGetAnswerPart2(t *testing.T) {
	want := 2713310158
	got := GetAnswerPart2("../example.txt")
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
