package part2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func IsContain(from1, to1, from2, to2 int) bool {
	if from1 >= from2 && from1 <= to2 {
		return true
	}
	if to1 >= from2 && to1 <= to2 {
		return true
	}
	if from2 >= from1 && from2 <= to1 {
		return true
	}
	if to2 >= from1 && to2 <= to1 {
		return true
	}
	return false
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
		elves := strings.Split(line, ",")

		elf1 := strings.Split(elves[0], "-")
		from1, _ := strconv.Atoi(elf1[0])
		to1, _ := strconv.Atoi(elf1[1])

		elf2 := strings.Split(elves[1], "-")
		from2, _ := strconv.Atoi(elf2[0])
		to2, _ := strconv.Atoi(elf2[1])

		if IsContain(from1, to1, from2, to2) {
			sum++
		}
	}
	return sum, nil
}
