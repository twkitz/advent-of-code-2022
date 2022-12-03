package part2

import "testing"

func TestGetBadge(t *testing.T) {
	type test struct {
		rucksack1	string
		rucksack2	string
		rucksack3	string
		want		rune
	}

	tests := []test {
		{
			rucksack1: "abcd",
			rucksack2: "AbCD",
			rucksack3: "xbyz",
			want: 'b',
		},
		{
			rucksack1: "vJrwpWtwJgWrhcsFMMfFFhFp",
			rucksack2: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			rucksack3: "PmmdzqPrVvPwwTWBwg",
			want: 'r',
		},
		{
			rucksack1: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			rucksack2: "ttgJtRGJQctTZtZT",
			rucksack3: "CrZsJsPPZsGzwwsLwLmpwMDw",
			want: 'Z',
		},
	}

	for _, tc := range tests {
		got := GetBadge(tc.rucksack1, tc.rucksack2, tc.rucksack3)
		if got != tc.want {
			t.Error("want: ", string(tc.want), " got: ", string(got))
		}
	}
}

func TestGetAnswer(t *testing.T) {
	want := 70

	got, _ := GetAnswer("../example.txt")

	if got != want {
		t.Error("want: ", want, " got: ", got)
	}
}