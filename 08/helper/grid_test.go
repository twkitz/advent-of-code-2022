package helper

import (
	"fmt"
	"testing"
)

func TestAddRow(t *testing.T) {
	grid := Grid{treeHeights: [][]int{}}

	grid.AddRow("1324")

	if len(grid.treeHeights) != 1 || len(grid.treeHeights[0]) != 4 {
		t.Error(fmt.Sprintf("wantedHeight: %d, height: %d, wantedWidth: %d, width: %d", 1, 4, len(grid.treeHeights), len(grid.treeHeights[0])))
	}
}

func TestGetHeight(t *testing.T) {
	heights := []int{
		1, 2,
	}

	for _, height := range heights {
		grid := Grid{treeHeights: [][]int{}}
		for i := 0; i < height; i++ {
			grid.treeHeights = append(grid.treeHeights, []int{1})
		}

		got := grid.GetHeight()
		if got != height {
			t.Error("invalid height", height)
		}
	}
}

func TestGetWidth(t *testing.T) {
	widths := []int{
		1, 2,
	}

	for _, width := range widths {
		grid := Grid{treeHeights: [][]int{[]int{}}}
		for i := 0; i < width; i++ {
			grid.treeHeights[0] = append(grid.treeHeights[0], 1)
		}

		got := grid.GetWidth()
		if got != width {
			t.Error("invalid height", width)
		}
	}
}

func TestIsAtEdge(t *testing.T) {
	grid := Grid{treeHeights: [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}}

	type test struct {
		row  int
		col  int
		want bool
	}

	tests := []test{
		{row: 0, col: 0, want: true},
		{row: 1, col: 0, want: true},
		{row: 0, col: 1, want: true},
		{row: 2, col: 0, want: true},
		{row: 0, col: 2, want: true},
		{row: 2, col: 2, want: true},
		{row: 1, col: 1, want: false},
	}

	for _, tc := range tests {
		got := grid.IsAtEdge(tc.row, tc.col)
		if got != tc.want {
			t.Error(fmt.Sprintf("row: %d, col: %d", tc.row, tc.col))
		}
	}
}

func TestIsVisibleFromTop(t *testing.T) {
	grid := Grid{treeHeights: [][]int{
		[]int{2, 1, 2, 2},
		[]int{2, 2, 2, 2},
		[]int{2, 2, 2, 2},
	}}

	type test struct {
		row  int
		col  int
		want bool
	}

	tests := []test{
		{row: 1, col: 1, want: true},
		{row: 1, col: 2, want: false},
	}

	for _, tc := range tests {
		got := grid.IsVisibleFromTop(tc.row, tc.col)
		if got != tc.want {
			t.Error(fmt.Sprintf("row: %d, col: %d", tc.row, tc.col))
		}
	}
}

func TestIsVisibleFromBottom(t *testing.T) {
	grid := Grid{treeHeights: [][]int{
		[]int{2, 2, 2, 2},
		[]int{2, 2, 2, 2},
		[]int{2, 1, 2, 2},
	}}

	type test struct {
		row  int
		col  int
		want bool
	}

	tests := []test{
		{row: 1, col: 1, want: true},
		{row: 1, col: 2, want: false},
	}

	for _, tc := range tests {
		got := grid.IsVisibleFromBottom(tc.row, tc.col)
		if got != tc.want {
			t.Error(fmt.Sprintf("row: %d, col: %d", tc.row, tc.col))
		}
	}
}

func TestIsVisibleFromLeft(t *testing.T) {
	grid := Grid{treeHeights: [][]int{
		[]int{2, 2, 2},
		[]int{1, 2, 2},
		[]int{2, 2, 2},
		[]int{2, 2, 2},
	}}

	type test struct {
		row  int
		col  int
		want bool
	}

	tests := []test{
		{row: 1, col: 1, want: true},
		{row: 2, col: 1, want: false},
	}

	for _, tc := range tests {
		got := grid.IsVisibleFromLeft(tc.row, tc.col)
		if got != tc.want {
			t.Error(fmt.Sprintf("row: %d, col: %d", tc.row, tc.col))
		}
	}
}

func TestIsVisibleFromRight(t *testing.T) {
	grid := Grid{treeHeights: [][]int{
		[]int{2, 2, 2},
		[]int{2, 2, 1},
		[]int{2, 2, 2},
		[]int{2, 2, 2},
	}}

	type test struct {
		row  int
		col  int
		want bool
	}

	tests := []test{
		{row: 1, col: 1, want: true},
		{row: 2, col: 1, want: false},
	}

	for _, tc := range tests {
		got := grid.IsVisibleFromRight(tc.row, tc.col)
		if got != tc.want {
			t.Error(fmt.Sprintf("row: %d, col: %d", tc.row, tc.col))
		}
	}
}

func TestGetViewDistanceToTop(t *testing.T) {
	grid := Grid{treeHeights: [][]int{
		[]int{1},
		[]int{3},
		[]int{2},
		[]int{3},
	}}

	type test struct {
		row  int
		col  int
		want int
	}

	tests := []test{
		{row: 0, col: 0, want: 0},
		{row: 1, col: 0, want: 1},
		{row: 2, col: 0, want: 1},
		{row: 3, col: 0, want: 2},
	}

	for _, tc := range tests {
		got := grid.GetViewDistanceToTop(tc.row, tc.col)
		if got != tc.want {
			t.Error(fmt.Sprintf("row: %d, col: %d", tc.row, tc.col))
		}
	}
}

func TestGetViewDistanceToBottom(t *testing.T) {
	grid := Grid{treeHeights: [][]int{
		[]int{3},
		[]int{2},
		[]int{3},
		[]int{1},
	}}

	type test struct {
		row  int
		col  int
		want int
	}

	tests := []test{
		{row: 0, col: 0, want: 2},
		{row: 1, col: 0, want: 1},
		{row: 2, col: 0, want: 1},
		{row: 3, col: 0, want: 0},
	}

	for _, tc := range tests {
		got := grid.GetViewDistanceToBottom(tc.row, tc.col)
		if got != tc.want {
			t.Error(fmt.Sprintf("row: %d, col: %d", tc.row, tc.col))
		}
	}
}

func TestGetViewDistanceToLeft(t *testing.T) {
	grid := Grid{treeHeights: [][]int{
		[]int{1, 3, 2, 3},
	}}

	type test struct {
		row  int
		col  int
		want int
	}

	tests := []test{
		{row: 0, col: 0, want: 0},
		{row: 0, col: 1, want: 1},
		{row: 0, col: 2, want: 1},
		{row: 0, col: 3, want: 2},
	}

	for _, tc := range tests {
		got := grid.GetViewDistanceToLeft(tc.row, tc.col)
		if got != tc.want {
			t.Error(fmt.Sprintf("row: %d, col: %d", tc.row, tc.col))
		}
	}
}

func TestGetViewDistanceToRight(t *testing.T) {
	grid := Grid{treeHeights: [][]int{
		[]int{1, 3, 2, 3},
	}}

	type test struct {
		row  int
		col  int
		want int
	}

	tests := []test{
		{row: 0, col: 0, want: 1},
		{row: 0, col: 1, want: 2},
		{row: 0, col: 2, want: 1},
		{row: 0, col: 3, want: 0},
	}

	for _, tc := range tests {
		got := grid.GetViewDistanceToRight(tc.row, tc.col)
		if got != tc.want {
			t.Error(fmt.Sprintf("row: %d, col: %d", tc.row, tc.col))
		}
	}
}
