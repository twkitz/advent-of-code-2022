package helper

import (
	"fmt"
	"testing"
)

func TestUpdateDropletSide(t *testing.T) {
	type test struct {
		targetValue int
		want        int
	}

	tests := []test{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{-1, -1},
	}

	for _, tc := range tests {
		ls := LavaSpector{
			map[[3]int]int{{0, 0, 0}: tc.targetValue},
		}
		ls.UpdateDropletSide(0, 0, 0)
		if ls.droplets[[3]int{0, 0, 0}] != tc.want {
			t.Errorf("want: %d, got: %d", tc.want, ls.droplets[[3]int{0, 0, 0}])
		}
	}
}

func TestAddDroplet(t *testing.T) {
	type test struct {
		droplets   map[[3]int]int
		newDroplet [3]int
		want       map[[3]int]int
	}

	tests := []test{
		{map[[3]int]int{}, [3]int{1, 1, 1}, map[[3]int]int{
			{1, 1, 1}: -1,
			{0, 1, 1}: 1,
			{2, 1, 1}: 1,
			{1, 0, 1}: 1,
			{1, 2, 1}: 1,
			{1, 1, 0}: 1,
			{1, 1, 2}: 1,
		}},
		{
			map[[3]int]int{
				{1, 1, 1}: -1,
				{0, 1, 1}: 1,
				{2, 1, 1}: 1,
				{1, 0, 1}: 1,
				{1, 2, 1}: 1,
				{1, 1, 0}: 1,
				{1, 1, 2}: 1,
			},
			[3]int{2, 1, 1},
			map[[3]int]int{
				{1, 1, 1}: -1,
				{0, 1, 1}: 1,
				{2, 1, 1}: -1,
				{1, 0, 1}: 1,
				{1, 2, 1}: 1,
				{1, 1, 0}: 1,
				{1, 1, 2}: 1,
				{3, 1, 1}: 1,
				{2, 0, 1}: 1,
				{2, 2, 1}: 1,
				{2, 1, 0}: 1,
				{2, 1, 2}: 1,
			},
		},
		{
			map[[3]int]int{
				{1, 1, 1}: -1,
				{0, 1, 1}: 1,
				{2, 1, 1}: 1,
				{1, 0, 1}: 1,
				{1, 2, 1}: 1,
				{1, 1, 0}: 1,
				{1, 1, 2}: 1,
			},
			[3]int{2, 2, 1},
			map[[3]int]int{
				{1, 1, 1}: -1,
				{0, 1, 1}: 1,
				{2, 1, 1}: 2,
				{1, 0, 1}: 1,
				{1, 2, 1}: 2,
				{1, 1, 0}: 1,
				{1, 1, 2}: 1,
				{2, 2, 1}: -1,
				{3, 2, 1}: 1,
				{2, 3, 1}: 1,
				{2, 2, 0}: 1,
				{2, 2, 2}: 1,
			},
		},
	}

	for _, tc := range tests {
		ls := LavaSpector{tc.droplets}
		ls.AddDroplet(tc.newDroplet[0], tc.newDroplet[1], tc.newDroplet[2])
		if fmt.Sprint(ls.droplets) != fmt.Sprint(tc.want) {
			t.Errorf("\n===want===\n%s\n\n===got===\n%s\n", fmt.Sprint(tc.want), fmt.Sprint(ls.droplets))
		}
	}
}

func TestGetExposedDropletSide(t *testing.T) {
	type test struct {
		droplets map[[3]int]int
		want     int
	}

	tests := []test{
		{map[[3]int]int{
			{1, 1, 1}: -1,
			{0, 1, 1}: 1,
			{2, 1, 1}: 1,
			{1, 0, 1}: 1,
			{1, 2, 1}: 1,
			{1, 1, 0}: 1,
			{1, 1, 2}: 1,
		}, 6},
		{map[[3]int]int{
			{1, 1, 1}: -1,
			{0, 1, 1}: 1,
			{2, 1, 1}: -1,
			{1, 0, 1}: 1,
			{1, 2, 1}: 1,
			{1, 1, 0}: 1,
			{1, 1, 2}: 1,
			{3, 1, 1}: 1,
			{2, 0, 1}: 1,
			{2, 2, 1}: 1,
			{2, 1, 0}: 1,
			{2, 1, 2}: 1,
		}, 10},
		{map[[3]int]int{
			{1, 1, 1}: -1,
			{0, 1, 1}: 1,
			{2, 1, 1}: 2,
			{1, 0, 1}: 1,
			{1, 2, 1}: 2,
			{1, 1, 0}: 1,
			{1, 1, 2}: 1,
			{2, 2, 1}: -1,
			{3, 2, 1}: 1,
			{2, 3, 1}: 1,
			{2, 2, 0}: 1,
			{2, 2, 2}: 1,
		}, 12},
	}

	for _, tc := range tests {
		ls := LavaSpector{tc.droplets}
		got := ls.GetExposedDropletSide()
		if got != tc.want {
			t.Errorf("want: %d, got: %d", tc.want, got)
		}
	}
}
