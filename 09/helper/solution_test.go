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
