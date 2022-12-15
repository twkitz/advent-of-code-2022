package helper

import (
	"fmt"
	"testing"
)

func TestReadInput(t *testing.T) {
	type test struct {
		input         string
		wantedSensorX int
		wantedSensorY int
		wantedBeaconX int
		wantedBeaconY int
	}

	tests := []test{
		{"Sensor at x=1, y=2: closest beacon is at x=3, y=4", 1, 2, 3, 4},
		{"Sensor at x=-1, y=-2: closest beacon is at x=-3, y=-4", -1, -2, -3, -4},
	}

	for _, tc := range tests {
		gotSensorX,
			gotSensorY,
			gotBeaconX,
			gotBeaconY := ReadInput(tc.input)

		if gotSensorX != tc.wantedSensorX ||
			gotSensorY != tc.wantedSensorY ||
			gotBeaconX != tc.wantedBeaconX ||
			gotBeaconY != tc.wantedBeaconY {
			t.Errorf("want: %d, %d, %d, %d :: got: %d, %d, %d, %d", tc.wantedSensorX, tc.wantedSensorY, tc.wantedBeaconX, tc.wantedBeaconY, gotSensorX, gotSensorY, gotBeaconX, gotBeaconY)
		}
	}
}

func TestGetSensorArea(t *testing.T) {
	type test struct {
		sensor Sensor
		want   [][2]int
	}

	tests := []test{
		{Sensor{0, 0, 0, 1}, [][2]int{
			{-1, 0},
			{0, -1},
			{0, 0},
			{0, 1},
			{1, 0},
		}},
		{Sensor{0, 0, 1, 1}, [][2]int{
			{-2, 0},
			{-1, -1},
			{-1, 0},
			{-1, 1},
			{0, -2},
			{0, -1},
			{0, 0},
			{0, 1},
			{0, 2},
			{1, -1},
			{1, 0},
			{1, 1},
			{2, 0},
		}},
	}

	for _, tc := range tests {
		got := GetSensorArea(tc.sensor)
		if fmt.Sprint(got) != fmt.Sprint(tc.want) {
			t.Errorf("\nwant: %s\ngot:  %s\n", fmt.Sprint(tc.want), fmt.Sprint(got))
		}
	}
}

func TestGetUnavailableInLine(t *testing.T) {
	type test struct {
		sensor    Sensor
		line      int
		wantedMin int
		wantedMax int
	}

	tests := []test{
		{Sensor{0, 0, 0, 1}, 0, -1, 1},
		{Sensor{0, 0, 0, 1}, 1, 0, 0},
		{Sensor{0, 0, 1, 1}, -2, 0, 0},
		{Sensor{0, 0, 1, 1}, 1, -1, 1},
		{Sensor{0, 0, 1, 1}, 3, 1, 0},
		{Sensor{39, 28, 34, 26}, 20, 1, 0},
		{Sensor{3, 365, -4, 388}, 200, 1, 0},
		{Sensor{0, 11, 2, 10}, 10, -2, 2},
	}

	for _, tc := range tests {
		gotMin, gotMax := GetUnavailableInLine(tc.sensor, tc.line)
		if gotMin != tc.wantedMin || gotMax != tc.wantedMax {
			t.Errorf("%s, %d -> want: %d %d, got:  %d %d", fmt.Sprint(tc.sensor), tc.line, tc.wantedMin, tc.wantedMax, gotMin, gotMax)
		}
	}
}

func TestMinifyAreas(t *testing.T) {
	type test struct {
		areas [][2]int
		want  [][2]int
	}

	tests := []test{
		{
			[][2]int{{0, 3}, {3, 5}},
			[][2]int{{0, 5}},
		},
		{
			[][2]int{{0, 3}, {4, 5}, {3, 4}},
			[][2]int{{0, 5}},
		},
		{
			[][2]int{{0, 3}, {1, 2}},
			[][2]int{{0, 3}},
		},
		{
			[][2]int{{0, 3}, {3, 5}, {7, 8}},
			[][2]int{{0, 5}, {7, 8}},
		},
	}

	for _, tc := range tests {
		got := MinifyAreas(tc.areas)
		if fmt.Sprint(got) != fmt.Sprint(tc.want) {
			t.Errorf("%s -> want: %s, got: %s", fmt.Sprint(tc.areas), fmt.Sprint(tc.want), fmt.Sprint(got))
		}
	}
}

func TestGetAnswerPart1(t *testing.T) {
	want := 26
	got := GetAnswerPart1("../example.txt", 10)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestGetAnswerPart1_1(t *testing.T) {
	want := 26
	got := GetAnswerPart1_1("../example.txt", 10)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestGetDistressPosition(t *testing.T) {
	wantedX := 14

	gotX := GetDistressPositionX("../example.txt", 11)
	if gotX != wantedX {
		t.Errorf("want: %d, got: %d", wantedX, gotX)
	}
}

func TestGetAnswerPart2(t *testing.T) {
	want := 56000011
	got := GetAnswerPart2("../example.txt", 0, 20)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
