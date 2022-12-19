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

func TestGetAnswerPart2(t *testing.T) {
	want := 58
	got := GetAnswerPart2("../example.txt")
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestGetAnswerPart2_airPocket(t *testing.T) {
	want := 66
	got := GetAnswerPart2("../example2.txt")
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestGetAnswerPart2_bigAirPack(t *testing.T) {
	want := 150
	got := GetAnswerPart2("../example3.txt")
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
