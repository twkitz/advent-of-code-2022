package part2

import (
	"bufio"
	"os"
	"strconv"

	"github.com/twkitz/advent-of-code/2022/05/part1"
)

func MultipleMove(stacksAddr *[][]string, count, from, to int) *[][]string {
	stacks := *stacksAddr
	fromLength := len(stacks[from-1])
	if fromLength > 0 {
		stacks[to-1] = append(stacks[to-1], stacks[from-1][fromLength-count:fromLength]...)
		stacks[from-1] = stacks[from-1][0 : fromLength-count]
	}
	return stacksAddr
}

func GetAnswer(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	initializedData := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		if _, err := strconv.Atoi(line[1:2]); err != nil {
			initializedData = append(initializedData, part1.ReadInput(line))
		}
	}

	stacks := part1.PrepareInitializeData(initializedData)

	for scanner.Scan() {
		line := scanner.Text()
		count, from, to := part1.GetCommand(line)
		MultipleMove(&stacks, count, from, to)
	}
	return part1.GetTops(&stacks), nil
}
