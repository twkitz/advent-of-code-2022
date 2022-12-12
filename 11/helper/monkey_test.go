package helper

import (
	"testing"
)

func TestOperate(t *testing.T) {
	type test struct {
		item                 int
		operator             string
		operand              string
		testBoredOperand     int
		targetWhenBoredTrue  int
		targetWhenBoredFalse int
		wantedTarget         int
		wantedWorryLevel     int
	}

	tests := []test{
		{
			item:                 79,
			operator:             "*",
			operand:              "19",
			testBoredOperand:     23,
			targetWhenBoredTrue:  2,
			targetWhenBoredFalse: 3,
			wantedTarget:         3,
			wantedWorryLevel:     500,
		},
		{
			item:                 98,
			operator:             "*",
			operand:              "19",
			testBoredOperand:     23,
			targetWhenBoredTrue:  2,
			targetWhenBoredFalse: 3,
			wantedTarget:         3,
			wantedWorryLevel:     620,
		},
		{
			item:                 79,
			operator:             "*",
			operand:              "old",
			testBoredOperand:     13,
			targetWhenBoredTrue:  1,
			targetWhenBoredFalse: 3,
			wantedTarget:         1,
			wantedWorryLevel:     2080,
		},
	}

	for _, tc := range tests {
		monkey := Monkey{
			[]int{tc.item},
			tc.operator,
			tc.operand,
			tc.testBoredOperand,
			tc.targetWhenBoredTrue,
			tc.targetWhenBoredFalse,
			1,
		}

		gotTarget, gotWorryLevel := monkey.Operate()
		if gotTarget != tc.wantedTarget || gotWorryLevel != tc.wantedWorryLevel {
			t.Errorf("want: %d %d, got: %d %d", tc.wantedTarget, tc.wantedWorryLevel, gotTarget, gotWorryLevel)
		}
	}
}
