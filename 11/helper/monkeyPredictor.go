package helper

import (
	"fmt"
	"strings"

	"github.com/twkitz/advent-of-code-2022/utils"
)

type MonkeyPredictor struct {
	monkeys []*Monkey
}

func (mp MonkeyPredictor) GetReport() string {
	result := []string{}
	for i := 0; i < len(mp.monkeys); i++ {
		items := utils.SliceToString(mp.monkeys[i].items)
		result = append(result, fmt.Sprintf("Monkey %d: %s", i, items))
	}
	return strings.Join(result, "\n")
}

func (mp *MonkeyPredictor) Run() {
	for _, monkey := range mp.monkeys {
		for len(monkey.items) > 0 {
			target, worryLevel := monkey.Operate()
			mp.monkeys[target].items = append(mp.monkeys[target].items, worryLevel)
			monkey.items = monkey.items[1:]
			monkey.inspectedCount++
		}
	}
}

func (mp MonkeyPredictor) GetInspectedCounts() []int {
	result := []int{}
	for _, monkey := range mp.monkeys {
		result = append(result, monkey.inspectedCount)
	}
	return result
}
