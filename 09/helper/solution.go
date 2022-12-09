package helper

import (
	"strconv"
	"strings"

	"github.com/twkitz/advent-of-code-2022/utils"
)

func GetAnswerPart1(filePath string) int {
	content := utils.GetFileContent(filePath)
	rope := Rope{0, 0, 0, 0}
	positionLog := make(map[string]bool)
	positionLog[rope.GetTailPosition()] = true
	for _, line := range content {
		command := strings.Split(line, " ")
		step, _ := strconv.Atoi(command[1])
		switch command[0] {
		case "U":
			for i := 0; i < step; i++ {
				rope.MoveUp()
				positionLog[rope.GetTailPosition()] = true
			}
			break
		case "D":
			for i := 0; i < step; i++ {
				rope.MoveDown()
				positionLog[rope.GetTailPosition()] = true
			}
			break
		case "L":
			for i := 0; i < step; i++ {
				rope.MoveLeft()
				positionLog[rope.GetTailPosition()] = true
			}
			break
		case "R":
			for i := 0; i < step; i++ {
				rope.MoveRight()
				positionLog[rope.GetTailPosition()] = true
			}
			break
		}
	}
	return len(positionLog)
}

func GetAnswerPart2(filePath string) int {
	content := utils.GetFileContent(filePath)
	knots := 9
	rope := make([]Rope, knots)
	positionLog := make(map[string]bool)
	positionLog[rope[0].GetTailPosition()] = true
	for _, line := range content {
		command := strings.Split(line, " ")
		step, _ := strconv.Atoi(command[1])
		for i := 0; i < step; i++ {
			switch command[0] {
			case "U":
				rope[0].MoveUp()
				break
			case "D":
				rope[0].MoveDown()
				break
			case "L":
				rope[0].MoveLeft()
				break
			case "R":
				rope[0].MoveRight()
				break
			}
			rope[1].headX = rope[0].tailX
			rope[1].headY = rope[0].tailY
			for j := 1; j < knots; j++ {
				oldTailPosition := rope[j].GetTailPosition()
				rope[j].UpdateTailPosition()
				newTailPosition := rope[j].GetTailPosition()
				if newTailPosition == oldTailPosition {
					break
				}
				if j < knots-1 {
					rope[j+1].headX = rope[j].tailX
					rope[j+1].headY = rope[j].tailY
				}
				if j == knots-1 {
					positionLog[newTailPosition] = true
				}
			}
		}
	}
	return len(positionLog)
}
