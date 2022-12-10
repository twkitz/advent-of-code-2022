package helper

import (
	"strconv"
	"strings"

	"github.com/twkitz/advent-of-code-2022/utils"
)

func GetAnswerPart1(filePath string, cycles []int) int {
	content := utils.GetFileContent(filePath)
	device := Device{x: 1, interestedCycles: cycles}
	for _, line := range content {
		command := strings.Split(line, " ")
		if command[0] == "noop" {
			device.Noop()
		} else if command[0] == "addx" {
			x, _ := strconv.Atoi(command[1])
			device.AddX(x)
		}
	}
	return device.totalSignal
}

func GetAnswerPart2(filePath string) string {
	content := utils.GetFileContent(filePath)
	device := Device{x: 1}
	for _, line := range content {
		command := strings.Split(line, " ")
		if command[0] == "noop" {
			device.Noop()
		} else if command[0] == "addx" {
			x, _ := strconv.Atoi(command[1])
			device.AddX(x)
		}
	}
	return device.image
}
