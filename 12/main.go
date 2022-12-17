package main

import (
	"fmt"

	"github.com/twkitz/advent-of-code-2022/12/helper"
)

func main() {
	part1 := helper.GetAnswer("input.txt", false)
	fmt.Println("part1:", part1)

	part2 := helper.GetAnswer("input.txt", true)
	fmt.Println("part2:", part2)
}
