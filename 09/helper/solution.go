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
