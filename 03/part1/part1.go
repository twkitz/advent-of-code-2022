package part1

import (
	"bufio"
	"os"
)

const offsetUpperCase = 64
const offsetLowerCase = 96

func GetPoint(character rune) int {
	ascii := int(character)
	if ascii > offsetLowerCase {
		return ascii - offsetLowerCase
	}
    return ascii - offsetUpperCase + 26
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
		line := scanner.Text()
		compartment1 := []rune(line[:len(line)/2])
		compartment2 := []rune(line[len(line)/2:])
		for _, item := range compartment2 {
			hasFound := false
			for _, item1 := range compartment1 {
				if item == item1 {
					hasFound = true
					sum += GetPoint(item)
					break
				}
			}
			if hasFound {
				break
			}
		}
    }
	return sum, nil
}