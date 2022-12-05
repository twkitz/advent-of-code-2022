package part1

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func ReadInput(input string) []string {
	result := []string{}
	for i := 0; i <= len(input)-3; i++ {
		crate := input[i+1 : i+2]
		if crate == " " {
			result = append(result, "")
		} else {
			result = append(result, crate)
		}
		i += 3
	}
	return result
}

func PrepareInitializeData(inputs [][]string) [][]string {
	result := [][]string{}

	for j := len(inputs) - 1; j >= 0; j-- {
		input := inputs[j]
		for i := 0; i < len(input); i++ {
			if input[i] == "" {
				continue
			}
			if i < len(result) {
				result[i] = append(result[i], input[i])
			} else {
				result = append(result, []string{input[i]})
			}
		}
	}
	return result
}

func Move(stacksAddr *[][]string, from, to int) *[][]string {
	stacks := *stacksAddr
	fromLength := len(stacks[from-1])
	if fromLength > 0 {
		stacks[to-1] = append(stacks[to-1], stacks[from-1][fromLength-1])
		stacks[from-1] = stacks[from-1][0 : fromLength-1]
	}
	return stacksAddr
}

func GetTops(stacksAddr *[][]string) string {
	result := ""
	for _, stack := range *stacksAddr {
		if len(stack) > 0 {
			result += stack[len(stack)-1]
		}
	}
	return result
}

func GetCommand(command string) (int, int, int) {
	regex := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)
	params := regex.FindStringSubmatch(command)

	count, _ := strconv.Atoi(params[1])
	from, _ := strconv.Atoi(params[2])
	to, _ := strconv.Atoi(params[3])

	return count, from, to
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
			initializedData = append(initializedData, ReadInput(line))
		}
	}

	stacks := PrepareInitializeData(initializedData)

	for scanner.Scan() {
		line := scanner.Text()
		count, from, to := GetCommand(line)
		for i := 0; i < count; i++ {
			Move(&stacks, from, to)
		}
	}
	return GetTops(&stacks), nil
}
