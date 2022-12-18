package helper

import (
	"strconv"
	"strings"

	"github.com/twkitz/advent-of-code-2022/utils"
)

func GetAnswerPart1(filePath string) int {
	content := utils.GetFileContent(filePath)
	ls := LavaSpector{make(map[[3]int]int)}
	for _, line := range content {
		position := strings.Split(line, ",")
		x, _ := strconv.Atoi(position[0])
		y, _ := strconv.Atoi(position[1])
		z, _ := strconv.Atoi(position[2])
		ls.AddDroplet(x, y, z)
	}
	return ls.GetExposedDropletSide()
}
