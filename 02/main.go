package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/twkitz/advent-of-code/2022/02/rockPaperScissor"
	"github.com/twkitz/advent-of-code/2022/02/rockPaperScissorPart2"
)

func main() {
	file, err := os.Open("input.txt")
    if err != nil {
		fmt.Println("Error", err)
    }
    defer file.Close()

	sum1 := 0
	sum2 := 0

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, " ")
		sum1 += rockPaperScissor.GetPoint(inputs[0], inputs[1])
		sum2 += rockPaperScissorPart2.GetPoint(inputs[0], inputs[1])
    }
	fmt.Println("sum1", sum1)
	fmt.Println("sum2", sum2)
}