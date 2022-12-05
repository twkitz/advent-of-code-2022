package main

import (
	"fmt"

	"github.com/twkitz/advent-of-code/2022/05/part1"
	"github.com/twkitz/advent-of-code/2022/05/part2"
)

func main() {
	part1, _ := part1.GetAnswer("input.txt")
	fmt.Println("part1: ", part1)

	part2, _ := part2.GetAnswer("input.txt")
	fmt.Println("part2: ", part2)
}
