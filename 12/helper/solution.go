package helper

import (
	"math"

	"github.com/twkitz/advent-of-code-2022/utils"
)

type Position struct {
	Row       int
	Col       int
	elevation rune
	MinStep   int
}

func CanGo(positions [][]Position, currentRow, currentCol, nextRow, nextCol, maxRow, maxCol int) bool {
	if nextRow < 0 || nextCol < 0 || nextCol > maxCol || nextRow > maxRow {
		return false
	}
	current := positions[currentRow][currentCol]
	next := positions[nextRow][nextCol]
	return int(current.elevation) >= int(next.elevation)-1
}

func BFS(positions [][]Position, checkingList [][2]int, targetRow, targetCol, maxRow, maxCol int) int {
	offsets := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for len(checkingList) > 0 {
		row := checkingList[0][0]
		col := checkingList[0][1]
		checkingList = checkingList[1:]

		position := positions[row][col]
		for _, offset := range offsets {
			nextRow := row + offset[0]
			nextCol := col + offset[1]
			if CanGo(positions, row, col, nextRow, nextCol, maxRow, maxCol) {
				nextPosition := &positions[nextRow][nextCol]
				if position.MinStep+1 < nextPosition.MinStep {
					nextPosition.MinStep = position.MinStep + 1
					checkingList = append(checkingList, [2]int{nextRow, nextCol})
				}
			}
		}
	}
	return positions[targetRow][targetCol].MinStep
}

func GetAnswer(filePath string, startFromAllA bool) int {
	content := utils.GetFileContent(filePath)
	positions := [][]Position{}
	checkingList := [][2]int{}
	row := 0
	col := 0
	targetRow := -1
	targetCol := -1
	for _, line := range content {
		col = 0
		rowPositions := []Position{}
		for _, char := range line {
			step := math.MaxInt
			if char == 'S' || (startFromAllA && char == 'a') {
				char = 'a'
				step = 0
				checkingList = append(checkingList, [2]int{row, col})
			} else if char == 'E' {
				char = 'z'
				targetRow = row
				targetCol = col
			}
			rowPositions = append(rowPositions, Position{Row: row, Col: col, elevation: char, MinStep: step})
			col++
		}
		positions = append(positions, rowPositions)
		row++
	}

	return BFS(positions, checkingList, targetRow, targetCol, row-1, col-1)
}
