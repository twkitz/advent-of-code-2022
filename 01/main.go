package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getTop3(max1, max2, max3, sum int) (int, int, int) {
	if sum > max1 {
		return sum, max1, max2
	}
	if sum > max2 {
		return max1, sum, max2
	}
	if sum > max3 {
		return max1, max2, sum
	}
	return max1, max2, max3
}

func main() {
	file, err := os.Open("input.txt")
    if err != nil {
		fmt.Println("Error", err)
    }
    defer file.Close()

	max1 := 0
	max2 := 0
	max3 := 0
	sum := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		value := scanner.Text()
		if value == "" {
			max1, max2, max3 = getTop3(max1, max2, max3, sum)
			sum = 0
		} else {
			intVal, _ := strconv.Atoi(value)
			sum += intVal
		}
    }

	max1, max2, max3 = getTop3(max1, max2, max3, sum)
	fmt.Println("Max", max1, max2, max3);
	fmt.Println("Max Sum", max1 + max2 + max3)
}
