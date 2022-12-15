package main

import (
	"fmt"

	"github.com/twkitz/advent-of-code-2022/15/helper"
)

func main() {
	part1 := helper.GetAnswerPart1_1("input.txt", 2000000)
	fmt.Println("part1:", part1)

	part2 := helper.GetAnswerPart2("input.txt", 0, 4000000)
	fmt.Println("part2:", part2)
}
