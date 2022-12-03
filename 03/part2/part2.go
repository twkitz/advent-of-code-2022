package part2

import (
	"bufio"
	"os"
	"github.com/twkitz/advent-of-code/2022/03/part1"
)

func GetBadge(rucksack1, rucksack2, rucksack3 string) rune {
	compartment1 := []rune(rucksack1)
	compartment2 := []rune(rucksack2)
	compartment3 := []rune(rucksack3)
	for _, item1 := range compartment1 {
		// hasFound := false
		for _, item2 := range compartment2 {
			if item1 == item2 {
				for _, item3 := range compartment3 {
					if item1 == item3 {
						return item1
					}
				}
			}
		}
	}
	return '-'
}

func GetAnswer(filePath string) (int, error) {
	file, err := os.Open(filePath)
    if err != nil {
		return -1, err
    }
    defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		badge := GetBadge(line1, line2, line3)
		sum += part1.GetPoint(badge)
    }
	return sum, nil
}