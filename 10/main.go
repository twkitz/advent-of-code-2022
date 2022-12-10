package main

import (
	"fmt"

	"github.com/twkitz/advent-of-code-2022/10/helper"
)

func main() {
	part1 := helper.GetAnswerPart1("input.txt", []int{20, 60, 100, 140, 180, 220})
	fmt.Println("part1: ", part1)

	part2 := helper.GetAnswerPart2("input.txt")
	fmt.Printf("part2: \n" + part2)
}
