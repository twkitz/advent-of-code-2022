package helper

import (
	"math"
	"strconv"
	"strings"

	"github.com/twkitz/advent-of-code-2022/utils"
)

func GetLavaSpector(filePath string) LavaSpector {
	content := utils.GetFileContent(filePath)
	ls := LavaSpector{
		make(map[[3]int]int),
		[3]int{math.MaxInt, math.MaxInt, math.MaxInt},
		[3]int{math.MinInt, math.MinInt, math.MinInt},
	}
	for _, line := range content {
		position := strings.Split(line, ",")
		x, _ := strconv.Atoi(position[0])
		y, _ := strconv.Atoi(position[1])
		z, _ := strconv.Atoi(position[2])
		ls.AddDroplet(x, y, z)
		if x < ls.minPosition[0] {
			ls.minPosition[0] = x
		}
		if y < ls.minPosition[1] {
			ls.minPosition[1] = y
		}
		if z < ls.minPosition[2] {
			ls.minPosition[2] = z
		}
		if x > ls.maxPosition[0] {
			ls.maxPosition[0] = x
		}
		if y > ls.maxPosition[1] {
			ls.maxPosition[1] = y
		}
		if z > ls.maxPosition[2] {
			ls.maxPosition[2] = z
		}
	}
	return ls
}

func GetAnswerPart1(filePath string) int {
	ls := GetLavaSpector(filePath)
	return ls.GetExposedDropletSide(false)
}

func GetAnswerPart2(filePath string) int {
	ls := GetLavaSpector(filePath)
	return ls.GetExposedDropletSide(true)
}
