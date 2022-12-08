package helper

import (
	"strconv"
)

type Grid struct {
	treeHeights [][]int
}

func (grid *Grid) AddRow(input string) {
	row := []int{}
	for _, char := range input {
		treeHeight, _ := strconv.Atoi(string(char))
		row = append(row, treeHeight)
	}
	grid.treeHeights = append(grid.treeHeights, row)
}

func (grid Grid) GetHeight() int {
	return len(grid.treeHeights)
}

func (grid Grid) GetWidth() int {
	return len(grid.treeHeights[0])
}

func (grid Grid) IsAtEdge(row, col int) bool {
	return row == 0 || col == 0 || row == grid.GetHeight()-1 || col == grid.GetWidth()-1
}

func (grid Grid) IsVisibleFromTop(row, col int) bool {
	for i := row - 1; i >= 0; i-- {
		if grid.treeHeights[i][col] >= grid.treeHeights[row][col] {
			return false
		}
	}
	return true
}

func (grid Grid) IsVisibleFromBottom(row, col int) bool {
	for i := row + 1; i < grid.GetHeight(); i++ {
		if grid.treeHeights[i][col] >= grid.treeHeights[row][col] {
			return false
		}
	}
	return true
}

func (grid Grid) IsVisibleFromLeft(row, col int) bool {
	for i := col - 1; i >= 0; i-- {
		if grid.treeHeights[row][i] >= grid.treeHeights[row][col] {
			return false
		}
	}
	return true
}

func (grid Grid) IsVisibleFromRight(row, col int) bool {
	for i := col + 1; i < grid.GetWidth(); i++ {
		if grid.treeHeights[row][i] >= grid.treeHeights[row][col] {
			return false
		}
	}
	return true
}

func (grid Grid) IsVisible(row, col int) bool {
	return grid.IsAtEdge(row, col) ||
		grid.IsVisibleFromTop(row, col) ||
		grid.IsVisibleFromBottom(row, col) ||
		grid.IsVisibleFromLeft(row, col) ||
		grid.IsVisibleFromRight(row, col)
}

func (grid Grid) GetViewDistanceToTop(row, col int) int {
	distance := 0
	for i := row - 1; i >= 0; i-- {
		distance++
		if grid.treeHeights[i][col] >= grid.treeHeights[row][col] {
			break
		}
	}
	return distance
}

func (grid Grid) GetViewDistanceToBottom(row, col int) int {
	distance := 0
	for i := row + 1; i < grid.GetHeight(); i++ {
		distance++
		if grid.treeHeights[i][col] >= grid.treeHeights[row][col] {
			break
		}
	}
	return distance
}

func (grid Grid) GetViewDistanceToLeft(row, col int) int {
	distance := 0
	for i := col - 1; i >= 0; i-- {
		distance++
		if grid.treeHeights[row][i] >= grid.treeHeights[row][col] {
			break
		}
	}
	return distance
}

func (grid Grid) GetViewDistanceToRight(row, col int) int {
	distance := 0
	for i := col + 1; i < grid.GetWidth(); i++ {
		distance++
		if grid.treeHeights[row][i] >= grid.treeHeights[row][col] {
			break
		}
	}
	return distance
}

func (grid Grid) GetScenaticScore(row, col int) int {
	return grid.GetViewDistanceToTop(row, col) *
		grid.GetViewDistanceToBottom(row, col) *
		grid.GetViewDistanceToLeft(row, col) *
		grid.GetViewDistanceToRight(row, col)
}
