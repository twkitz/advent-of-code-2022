package helper

import (
	"github.com/twkitz/advent-of-code-2022/utils"
)

func GetGrid(filePath string) (*Grid, error) {
	content := utils.GetFileContent(filePath)
	grid := Grid{treeHeights: [][]int{}}

	for _, line := range content {
		grid.AddRow(line)
	}

	return &grid, nil
}

func GetAnswerPart1(filePath string) int {
	grid, _ := GetGrid(filePath)
	sum := 0
	for row := 0; row < grid.GetHeight(); row++ {
		for col := 0; col < grid.GetWidth(); col++ {
			if grid.IsVisible(row, col) {
				sum++
			}
		}
	}

	return sum
}

func GetAnswerPart2(filePath string) int {
	grid, _ := GetGrid(filePath)
	max := 0
	for row := 0; row < grid.GetHeight(); row++ {
		for col := 0; col < grid.GetWidth(); col++ {
			scenicScore := grid.GetScenaticScore(row, col)
			if scenicScore > max {
				max = scenicScore
			}
		}
	}

	return max
}
