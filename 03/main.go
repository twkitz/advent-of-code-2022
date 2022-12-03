package main

import (
	"fmt"
	"github.com/twkitz/advent-of-code/2022/03/part1"
	"github.com/twkitz/advent-of-code/2022/03/part2"
)

func main() {
	part1, part1Err := part1.GetAnswer("input.txt")
	if part1Err != nil {
		fmt.Println("part1 error: ", part1Err)
	}
	fmt.Println("part1: ", part1)

	part2, part2Err := part2.GetAnswer("input.txt")
	if part2Err != nil {
		fmt.Println("part2 error: ", part2Err)
	}
	fmt.Println("part2: ", part2)
}