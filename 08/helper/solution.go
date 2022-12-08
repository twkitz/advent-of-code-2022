package helper

import (
	"bufio"
	"os"
)

func GetGrid(filePath string) (*Grid, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := Grid{treeHeights: [][]int{}}

	for scanner.Scan() {
		input := scanner.Text()
		grid.AddRow(input)
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
