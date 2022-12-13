package helper

import (
	"strconv"
	"strings"

	"github.com/twkitz/advent-of-code-2022/utils"
)

func GetAnswer(filePath string, roundCount, worryDivider int) int {
	content := utils.GetFileContent(filePath)
	mp := MonkeyPredictor{[]*Monkey{}, 0}
	totalTestOperand := 1
	for lineNo := 0; lineNo < len(content); lineNo++ {
		if strings.HasPrefix(content[lineNo], "Monkey ") {
			monkey := Monkey{worryDivider: worryDivider}
			lineNo++
			for lineNo < len(content) && content[lineNo] != "" {
				if strings.HasPrefix(content[lineNo], "  Starting items: ") {
					prefixLength := len("  Starting items: ")
					items := []int{}
					for _, itemRaw := range strings.Split(content[lineNo][prefixLength:], ", ") {
						item, _ := strconv.Atoi(itemRaw)
						items = append(items, item)
					}
					monkey.items = items
				} else if strings.HasPrefix(content[lineNo], "  Operation: new = old ") {
					prefixLength := len("  Operation: new = old ")
					params := strings.Split(content[lineNo][prefixLength:], " ")
					monkey.operator = params[0]
					monkey.operand = params[1]
				} else if strings.HasPrefix(content[lineNo], "  Test: divisible by ") {
					prefixLength := len("  Test: divisible by ")
					monkey.testBoredOperand, _ = strconv.Atoi(content[lineNo][prefixLength:])
				} else if strings.HasPrefix(content[lineNo], "    If true: throw to monkey ") {
					prefixLength := len("    If true: throw to monkey ")
					monkey.targetWhenBoredTrue, _ = strconv.Atoi(content[lineNo][prefixLength:])
				} else if strings.HasPrefix(content[lineNo], "    If false: throw to monkey ") {
					prefixLength := len("    If false: throw to monkey ")
					monkey.targetWhenBoredFalse, _ = strconv.Atoi(content[lineNo][prefixLength:])
				}
				lineNo++
			}
			mp.monkeys = append(mp.monkeys, &monkey)
			totalTestOperand *= monkey.testBoredOperand
		}
	}
	mp.totalTestOperand = totalTestOperand

	for i := 0; i < roundCount; i++ {
		mp.Run()
	}

	inspectedCounts := mp.GetInspectedCounts()
	top1 := 0
	top2 := 0
	for _, count := range inspectedCounts {
		if count > top1 {
			top2 = top1
			top1 = count
		} else if count > top2 {
			top2 = count
		}
	}

	return top1 * top2
}

func GetAnswerPart1(filePath string) int {
	return GetAnswer(filePath, 20, 3)
}

func GetAnswerPart2(filePath string) int {
	return GetAnswer(filePath, 10000, 1)
}
