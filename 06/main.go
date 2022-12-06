package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/twkitz/advent-of-code/2022/06/markerHelper"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		part1 := markerHelper.GetFirstMarker(4, input)
		fmt.Println("part1: ", part1)

		part2 := markerHelper.GetFirstMarker(14, input)
		fmt.Println("part2: ", part2)
	}
}
