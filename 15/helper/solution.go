package helper

import (
	"math"
	"regexp"
	"strconv"

	"github.com/twkitz/advent-of-code-2022/utils"
)

type Sensor struct {
	X       int
	Y       int
	BeaconX int
	BeaconY int
}

func ReadInput(input string) (int, int, int, int) {
	regex := regexp.MustCompile(`Sensor at x=(-?\d*), y=(-?\d*): closest beacon is at x=(-?\d*), y=(-?\d*)`)
	params := regex.FindStringSubmatch(input)
	sensorX, _ := strconv.Atoi(params[1])
	sensorY, _ := strconv.Atoi(params[2])
	beaconX, _ := strconv.Atoi(params[3])
	beaconY, _ := strconv.Atoi(params[4])
	return sensorX, sensorY, beaconX, beaconY
}

func GetSensorArea(s Sensor) [][2]int {
	area := [][2]int{}
	diffX := int(math.Abs(float64(s.X - s.BeaconX)))
	diffY := int(math.Abs(float64(s.Y - s.BeaconY)))
	totalDiff := diffX + diffY
	for x := s.X - totalDiff; x <= s.X+totalDiff; x++ {
		for y := s.Y - totalDiff; y <= s.Y+totalDiff; y++ {
			diff := int(math.Abs(float64(s.X-x)) + math.Abs(float64(s.Y-y)))
			if diff <= totalDiff {
				area = append(area, [2]int{x, y})
			}
		}
	}
	return area
}

func GetUnavailableInLine(s Sensor, line int) (int, int) {
	diffX := int(math.Abs(float64(s.X - s.BeaconX)))
	diffY := int(math.Abs(float64(s.Y - s.BeaconY)))
	totalDiff := diffX + diffY
	startY := s.Y - totalDiff
	endY := s.Y + totalDiff

	if line < startY || line > endY {
		return 1, 0
	}

	diffFromTargetLine := int(math.Abs(float64(s.Y - line)))
	minX := s.X - (totalDiff - diffFromTargetLine)
	maxX := s.X + (totalDiff - diffFromTargetLine)
	return minX, maxX
}

func MinifyAreas(areas [][2]int) [][2]int {
	// for _, a := range areas {
	// 	fmt.Printf("%d\t%d\n", a[0], a[1])
	// }
	// fmt.Println()
	minifiedAreas := [][2]int{areas[0]}
	for i := 1; i < len(areas); i++ {
		isMerged := false
		for j := 0; j < len(minifiedAreas); j++ {
			if (minifiedAreas[j][0] <= areas[i][0] && minifiedAreas[j][1] >= areas[i][0]) ||
				(minifiedAreas[j][0] <= areas[i][1] && minifiedAreas[j][1] >= areas[i][1]) ||
				(minifiedAreas[j][0] >= areas[i][0] && minifiedAreas[j][1] <= areas[i][1]) {
				if areas[i][0] < minifiedAreas[j][0] {
					minifiedAreas[j][0] = areas[i][0]
				}
				if areas[i][1] > minifiedAreas[j][1] {
					minifiedAreas[j][1] = areas[i][1]
				}
				isMerged = true
				break
			}
		}
		if !isMerged {
			minifiedAreas = append(minifiedAreas, areas[i])
		}
	}
	if len(areas) != len(minifiedAreas) {
		return MinifyAreas(minifiedAreas)
	}
	return minifiedAreas
}

func GetAnswerPart1(filePath string, y int) int {
	content := utils.GetFileContent(filePath)
	unavailableInLine := make(map[[2]int]bool)
	for _, input := range content {
		sX, sY, bX, bY := ReadInput(input)
		sensorArea := GetSensorArea(Sensor{sX, sY, bX, bY})
		for _, area := range sensorArea {
			if area[1] == y {
				unavailableInLine[area] = true
			}
		}
	}
	return len(unavailableInLine) - 1
}

func GetAnswerPart1_1(filePath string, y int) int {
	content := utils.GetFileContent(filePath)
	unavailableInLine := [][2]int{}
	for _, input := range content {
		sX, sY, bX, bY := ReadInput(input)
		min, max := GetUnavailableInLine(Sensor{sX, sY, bX, bY}, y)
		if min <= max {
			unavailableInLine = append(unavailableInLine, [2]int{min, max})
		}
	}
	minifiedAreas := MinifyAreas(unavailableInLine)
	sum := 0
	for _, area := range minifiedAreas {
		sum += area[1] - area[0]
	}
	return sum
}
