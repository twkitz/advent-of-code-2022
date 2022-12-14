package helper

import (
	"strconv"
)

type Monkey struct {
	items                []int
	operator             string
	operand              string
	testBoredOperand     int
	targetWhenBoredTrue  int
	targetWhenBoredFalse int
	inspectedCount       int
	worryDivider         int
}

func (m Monkey) Operate() (int, int) {
	operand := 0
	newWorryLevel := 0
	oldWorryLevel := m.items[0]
	if m.operand == "old" {
		operand = oldWorryLevel
	} else {
		operand, _ = strconv.Atoi(m.operand)
	}
	switch m.operator {
	case "+":
		newWorryLevel = (oldWorryLevel + operand) / m.worryDivider
	case "*":
		newWorryLevel = (oldWorryLevel * operand) / m.worryDivider
	}
	if newWorryLevel%m.testBoredOperand == 0 {
		return m.targetWhenBoredTrue, newWorryLevel
	}
	return m.targetWhenBoredFalse, newWorryLevel
}
