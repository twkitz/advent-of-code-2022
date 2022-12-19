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
			[3]int{0, 0, 0},
			[3]int{0, 0, 0},
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
		ls := LavaSpector{tc.droplets, [3]int{0, 0, 0}, [3]int{2, 2, 2}}
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
		ls := LavaSpector{tc.droplets, [3]int{0, 0, 0}, [3]int{3, 3, 3}}
		got := ls.GetExposedDropletSide(false)
		if got != tc.want {
			t.Errorf("want: %d, got: %d", tc.want, got)
		}
	}
}

func TestIsAirPocket(t *testing.T) {
	type test struct {
		targetValue int
		want        bool
	}

	tests := []test{
		{-1, false},
		{0, false},
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{5, false},
		{6, true},
	}

	for _, tc := range tests {
		ls := LavaSpector{
			map[[3]int]int{{0, 0, 0}: tc.targetValue},
			[3]int{0, 0, 0},
			[3]int{0, 0, 0},
		}
		got := ls.IsAirPocket([3]int{0, 0, 0})
		if got != tc.want {
			t.Errorf("want: %t, got: %t", tc.want, got)
		}
	}
}
