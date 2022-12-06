package markerHelper

import "testing"

func TestIsMarker(t *testing.T) {
	type test struct {
		input string
		want  bool
	}

	tests := []test{
		{input: "abcd", want: true},
		{input: "aacd", want: false},
		{input: "abad", want: false},
		{input: "abca", want: false},
	}

	for _, tc := range tests {
		got := IsMarker(tc.input)
		if got != tc.want {
			t.Error("want:", tc.want, "got:", got)
		}
	}
}

func TestGetFirstMarker(t *testing.T) {
	type test struct {
		length  int
		message string
		want    int
	}

	tests := []test{
		{length: 4, message: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 5},
		{length: 4, message: "nppdvjthqldpwncqszvftbrmjlhg", want: 6},
		{length: 4, message: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 10},
		{length: 4, message: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 11},
		{length: 14, message: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 19},
		{length: 14, message: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 23},
		{length: 14, message: "nppdvjthqldpwncqszvftbrmjlhg", want: 23},
		{length: 14, message: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 29},
		{length: 14, message: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 26},
	}

	for _, tc := range tests {
		got := GetFirstMarker(tc.length, tc.message)
		if got != tc.want {
			t.Error("want:", tc.want, "got:", got)
		}
	}
}
